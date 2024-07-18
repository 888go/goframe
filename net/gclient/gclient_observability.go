// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/internal/httputil"
	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gmetric"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	instrumentName                              = "github.com/gogf/gf/v2/net/gclient.Client"
	tracingAttrHttpAddressRemote                = "http.address.remote"
	tracingAttrHttpAddressLocal                 = "http.address.local"
	tracingAttrHttpDnsStart                     = "http.dns.start"
	tracingAttrHttpDnsDone                      = "http.dns.done"
	tracingAttrHttpConnectStart                 = "http.connect.start"
	tracingAttrHttpConnectDone                  = "http.connect.done"
	tracingEventHttpRequest                     = "http.request"
	tracingEventHttpRequestHeaders              = "http.request.headers"
	tracingEventHttpRequestBaggage              = "http.request.baggage"
	tracingEventHttpRequestBody                 = "http.request.body"
	tracingEventHttpResponse                    = "http.response"
	tracingEventHttpResponseHeaders             = "http.response.headers"
	tracingEventHttpResponseBody                = "http.response.body"
	tracingMiddlewareHandled        gctx.StrKey = `MiddlewareClientTracingHandled`
)

// internalMiddlewareObservability 是一个客户端中间件，用于启用可观测性特性。 md5:3e8ca5815bee6c43
func internalMiddlewareObservability(c *Client, r *http.Request) (response *Response, err error) {
	var ctx = r.Context()
// 标记此请求已由服务器跟踪中间件处理，以避免被相同的中间件重复处理。
// md5:0ca4c50f5a9f8851
	if ctx.Value(tracingMiddlewareHandled) != nil {
		return c.Next(r)
	}

	ctx = context.WithValue(ctx, tracingMiddlewareHandled, 1)
	tr := otel.GetTracerProvider().Tracer(
		instrumentName,
		trace.WithInstrumentationVersion(gf.VERSION),
	)
	ctx, span := tr.Start(ctx, r.URL.String(), trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	span.SetAttributes(gtrace.CommonLabels()...)

	// 将跟踪内容注入到HTTP头中。 md5:1e1cdfd3fa2c105a
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))

	// 将ClientTrace注入到HTTP请求的上下文中。 md5:2eb24a7227c63b0f
	var (
		httpClientTracer       *httptrace.ClientTrace
		baseClientTracer       = newClientTracerNoop()
		isUsingDefaultProvider = gtrace.IsUsingDefaultProvider()
	)
	// Tracing.
	if !isUsingDefaultProvider {
		baseClientTracer = newClientTracerTracing(ctx, span, r)
	}
	// Metrics.
	if gmetric.IsEnabled() {
		baseClientTracer = newClientTracerMetrics(r, baseClientTracer)
	}
	httpClientTracer = newClientTracer(baseClientTracer)
	r = r.WithContext(
		httptrace.WithClientTrace(
			ctx, httpClientTracer,
		),
	)
	response, err = c.Next(r)

	// 如果当前正在使用默认的跟踪提供程序，那么它将不执行复杂的跟踪任务。 md5:27e9b5e1b834aa5d
	if isUsingDefaultProvider {
		return
	}

	if err != nil {
		span.SetStatus(codes.Error, fmt.Sprintf(`%+v`, err))
	}
	if response == nil || response.Response == nil {
		return
	}

	reqBodyContentBytes, _ := io.ReadAll(response.Body)
	response.Body = utils.NewReadCloser(reqBodyContentBytes, false)

	resBodyContent, err := gtrace.SafeContentForHttp(reqBodyContentBytes, response.Header)
	if err != nil {
		span.SetStatus(codes.Error, fmt.Sprintf(`converting safe content failed: %s`, err.Error()))
	}

	span.AddEvent(tracingEventHttpResponse, trace.WithAttributes(
		attribute.String(tracingEventHttpResponseHeaders, gconv.String(httputil.HeaderToMap(response.Header))),
		attribute.String(tracingEventHttpResponseBody, resBodyContent),
	))
	return
}
