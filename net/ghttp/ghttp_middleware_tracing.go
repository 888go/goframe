// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	
	"github.com/888go/goframe"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/httputil"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/net/gtrace"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	)
const (
	tracingInstrumentName                       = "github.com/gogf/gf/v2/net/ghttp.Server"
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

// internalMiddlewareServerTracing 是一个服务器中间件，它利用 OpenTelemetry 的标准启用追踪功能。
func internalMiddlewareServerTracing(r *Request) {
	var (
		ctx = r.Context()
	)
// 标记该请求已被服务器追踪中间件处理，
// 以避免被同一中间件重复处理。
	if ctx.Value(tracingMiddlewareHandled) != nil {
		r.Middleware.Next()
		return
	}

	ctx = context.WithValue(ctx, tracingMiddlewareHandled, 1)
	var (
		span trace.Span
		tr   = otel.GetTracerProvider().Tracer(
			tracingInstrumentName,
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

	// 注入追踪上下文。
	r.SetCtx(ctx)

	// 如果当前正在使用默认的追踪提供者，则不执行复杂的追踪任务。
	if gtrace.IsUsingDefaultProvider() {
		r.Middleware.Next()
		return
	}

	// 请求内容日志记录。
	reqBodyContentBytes, err := io.ReadAll(r.Body)
	if err != nil {
		r.SetError(gerror.Wrap(err, `read request body failed`))
		span.SetStatus(codes.Error, fmt.Sprintf(`%+v`, err))
		return
	}
	r.Body = utils.NewReadCloser(reqBodyContentBytes, false)

	span.AddEvent(tracingEventHttpRequest, trace.WithAttributes(
		attribute.String(tracingEventHttpRequestUrl, r.URL.String()),
		attribute.String(tracingEventHttpRequestHeaders, gconv.String(httputil.HeaderToMap(r.Header))),
		attribute.String(tracingEventHttpRequestBaggage, gtrace.GetBaggageMap(ctx).String()),
		attribute.String(tracingEventHttpRequestBody, gstr.StrLimit(
			string(reqBodyContentBytes),
			gtrace.MaxContentLogSize(),
			"...",
		)),
	))

	// 继续执行。
	r.Middleware.Next()

	// Error logging.
	if err = r.GetError(); err != nil {
		span.SetStatus(codes.Error, fmt.Sprintf(`%+v`, err))
	}
	// 响应内容日志记录。
	var resBodyContent = gstr.StrLimit(r.Response.BufferString(), gtrace.MaxContentLogSize(), "...")
	if gzipAccepted(r.Response.Header()) {
		reader, err := gzip.NewReader(strings.NewReader(r.Response.BufferString()))
		if err != nil {
			span.SetStatus(codes.Error, fmt.Sprintf(`read gzip response err:%+v`, err))
		}
		defer reader.Close()
		uncompressed, err := io.ReadAll(reader)
		if err != nil {
			span.SetStatus(codes.Error, fmt.Sprintf(`get uncompress value err:%+v`, err))
		}
		resBodyContent = gstr.StrLimit(string(uncompressed), gtrace.MaxContentLogSize(), "...")
	}

	span.AddEvent(tracingEventHttpResponse, trace.WithAttributes(
		attribute.String(tracingEventHttpResponseHeaders, gconv.String(httputil.HeaderToMap(r.Response.Header()))),
		attribute.String(tracingEventHttpResponseBody, resBodyContent),
	))
}

// gzipAccepted 返回一个布尔值，表示客户端是否接受 gzip 压缩编码的内容。
func gzipAccepted(header http.Header) bool {
	a := header.Get("Content-Encoding")
	parts := strings.Split(a, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "gzip" || strings.HasPrefix(part, "gzip;") {
			return true
		}
	}
	return false
}
