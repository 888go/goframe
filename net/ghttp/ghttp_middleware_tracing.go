// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"context"
	"fmt"
	"io"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/httputil"
	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	instrumentName                              = "github.com/gogf/gf/v2/net/ghttp.Server"
	tracingEventHttpRequest                     = "http.request"
	tracingEventHttpRequestHeaders              = "http.request.headers"
	tracingEventHttpRequestBaggage              = "http.request.baggage"
	tracingEventHttpRequestBody                 = "http.request.body"
	tracingEventHttpResponse                    = "http.response"
	tracingEventHttpResponseHeaders             = "http.response.headers"
	tracingEventHttpResponseBody                = "http.response.body"
	tracingEventHttpRequestUrl                  = "http.request.url"
	tracingMiddlewareHandled        gctx.StrKey = `MiddlewareServerTracingHandled`
)

// internalMiddlewareServerTracing 是一个服务器中间件，它使用 OpenTelemetry 标准启用追踪功能。 md5:935f3728f0dd44c3
func internalMiddlewareServerTracing(r *Request) {
	var (
		ctx = r.Context()
	)
	// 标记此请求已由服务器跟踪中间件处理，以避免被相同的中间件重复处理。
	// md5:0ca4c50f5a9f8851
	if ctx.Value(tracingMiddlewareHandled) != nil {
		r.Middleware.Next()
		return
	}

	ctx = context.WithValue(ctx, tracingMiddlewareHandled, 1)
	var (
		span trace.Span
		tr   = otel.GetTracerProvider().Tracer(
			instrumentName,
			trace.WithInstrumentationVersion(gf.VERSION),
		)
	)
	ctx, span = tr.Start(
		otel.GetTextMapPropagator().Extract(
			ctx,
			propagation.HeaderCarrier(r.Header),
		),
		r.URL.Path,
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	span.SetAttributes(gtrace.CommonLabels()...)

		// 注入追踪上下文。 md5:97547f0c7d05fa84
	r.SetCtx(ctx)

		// 如果现在它正在使用默认的追踪提供者，那么它就不会执行复杂的追踪任务。 md5:5a8a3f90fc875a4f
	if gtrace.IsUsingDefaultProvider() {
		r.Middleware.Next()
		return
	}

		// 请求内容日志记录。 md5:b2d7ffe537e9751f
	reqBodyContentBytes, err := io.ReadAll(r.Body)
	if err != nil {
		r.SetError(gerror.Wrap(err, `read request body failed`))
		span.SetStatus(codes.Error, fmt.Sprintf(`%+v`, err))
		return
	}
	r.Body = utils.NewReadCloser(reqBodyContentBytes, false)
	reqBodyContent, err := gtrace.SafeContentForHttp(reqBodyContentBytes, r.Header)
	if err != nil {
		span.SetStatus(codes.Error, fmt.Sprintf(`converting safe content failed: %s`, err.Error()))
	}

	span.AddEvent(tracingEventHttpRequest, trace.WithAttributes(
		attribute.String(tracingEventHttpRequestUrl, r.URL.String()),
		attribute.String(tracingEventHttpRequestHeaders, gconv.String(httputil.HeaderToMap(r.Header))),
		attribute.String(tracingEventHttpRequestBaggage, gtrace.GetBaggageMap(ctx).String()),
		attribute.String(tracingEventHttpRequestBody, reqBodyContent),
	))

	// Continue executing.
	r.Middleware.Next()

		// 在将路由设置为span名称后进行解析. md5:9bc1405ad003a5df
	if handler := r.GetServeHandler(); handler != nil && handler.Handler.Router != nil {
		span.SetName(handler.Handler.Router.Uri)
	}

	// Error logging.
	if err = r.GetError(); err != nil {
		span.SetStatus(codes.Error, fmt.Sprintf(`%+v`, err))
	}

		// 响应内容日志记录。 md5:77aef723834bb158
	resBodyContent, err := gtrace.SafeContentForHttp(r.Response.Buffer(), r.Response.Header())
	if err != nil {
		span.SetStatus(codes.Error, fmt.Sprintf(`converting safe content failed: %s`, err.Error()))
	}

	span.AddEvent(tracingEventHttpResponse, trace.WithAttributes(
		attribute.String(tracingEventHttpResponseHeaders, gconv.String(httputil.HeaderToMap(r.Response.Header()))),
		attribute.String(tracingEventHttpResponseBody, resBodyContent),
	))
}
