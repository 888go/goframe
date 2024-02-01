// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

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
	
	"github.com/888go/goframe"
	"github.com/888go/goframe/internal/httputil"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/net/gtrace"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	)
const (
	tracingInstrumentName                       = "github.com/gogf/gf/v2/net/gclient.Client"
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

// internalMiddlewareTracing 是一个客户端中间件，它利用 OpenTelemetry 的标准启用追踪功能。
func internalMiddlewareTracing(c *Client, r *http.Request) (response *Response, err error) {
	var ctx = r.Context()
// 标记该请求已被服务器追踪中间件处理，
// 以避免被同一中间件重复处理。
	if ctx.Value(tracingMiddlewareHandled) != nil {
		return c.Next(r)
	}

	ctx = context.WithValue(ctx, tracingMiddlewareHandled, 1)
	tr := otel.GetTracerProvider().Tracer(
		tracingInstrumentName,
		trace.WithInstrumentationVersion(gf.VERSION),
	)
	ctx, span := tr.Start(ctx, r.URL.String(), trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	span.SetAttributes(gtrace.CommonLabels()...)

	// 向HTTP头中注入追踪内容
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))

	// 如果当前正在使用默认的追踪提供者，则不执行任何复杂的追踪任务。
	if gtrace.IsUsingDefaultProvider() {
		response, err = c.Next(r)
		return
	}

	// 继续执行客户端处理器。
	response, err = c.Next(
		r.WithContext(
			httptrace.WithClientTrace(
				ctx, newClientTrace(ctx, span, r),
			),
		),
	)
	if err != nil {
		span.SetStatus(codes.Error, fmt.Sprintf(`%+v`, err))
	}
	if response == nil || response.Response == nil {
		return
	}

	reqBodyContentBytes, _ := io.ReadAll(response.Body)
	response.Body = utils.NewReadCloser(reqBodyContentBytes, false)

	span.AddEvent(tracingEventHttpResponse, trace.WithAttributes(
		attribute.String(tracingEventHttpResponseHeaders, gconv.String(httputil.HeaderToMap(response.Header))),
		attribute.String(tracingEventHttpResponseBody, gstr.StrLimit(
			string(reqBodyContentBytes),
			gtrace.MaxContentLogSize(),
			"...",
		)),
	))
	return
}
