// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gclient

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"net/textproto"
	"strings"
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"
)

// clientTracerTracing 用于实现 httptrace.ClientTrace。 md5:06ba597d73062dd5
type clientTracerTracing struct {
	context.Context
	span        trace.Span
	request     *http.Request
	requestBody []byte
	headers     map[string]interface{}
	mtx         sync.Mutex
}

// newClientTracerTracing 创建并返回一个 httptrace.ClientTrace 对象。 md5:3b73f25116451d66
func newClientTracerTracing(
	ctx context.Context,
	span trace.Span,
	request *http.Request,
) *httptrace.ClientTrace {
	ct := &clientTracerTracing{
		Context: ctx,
		span:    span,
		request: request,
		headers: make(map[string]interface{}),
	}

	reqBodyContent, _ := io.ReadAll(ct.request.Body)
	ct.requestBody = reqBodyContent
	ct.request.Body = utils.NewReadCloser(reqBodyContent, false)

	return &httptrace.ClientTrace{
		GetConn:              ct.GetConn,
		GotConn:              ct.GotConn,
		PutIdleConn:          ct.PutIdleConn,
		GotFirstResponseByte: ct.GotFirstResponseByte,
		Got100Continue:       ct.Got100Continue,
		Got1xxResponse:       ct.Got1xxResponse,
		DNSStart:             ct.DNSStart,
		DNSDone:              ct.DNSDone,
		ConnectStart:         ct.ConnectStart,
		ConnectDone:          ct.ConnectDone,
		TLSHandshakeStart:    ct.TLSHandshakeStart,
		TLSHandshakeDone:     ct.TLSHandshakeDone,
		WroteHeaderField:     ct.WroteHeaderField,
		WroteHeaders:         ct.WroteHeaders,
		Wait100Continue:      ct.Wait100Continue,
		WroteRequest:         ct.WroteRequest,
	}
}

// 在创建或从空闲池中检索连接之前，会调用GetConn。hostPort是目标或代理的"主机:端口"。即使已经有空闲缓存的连接可用，也会调用GetConn。
// md5:de9f40c04cc79ecf
// ff:
// ct:
// host:
func (ct *clientTracerTracing) GetConn(host string) {}

// GotConn 在成功建立连接后被调用。如果无法获取连接，没有相应的钩子；相反，应使用 Transport.RoundTrip 的错误。
// md5:939ef1676a19b2be
// ff:
// ct:
// info:
func (ct *clientTracerTracing) GotConn(info httptrace.GotConnInfo) {
	remoteAddr := ""
	if info.Conn.RemoteAddr() != nil {
		remoteAddr = info.Conn.RemoteAddr().String()
	}
	localAddr := ""
	if info.Conn.LocalAddr() != nil {
		localAddr = info.Conn.LocalAddr().String()
	}
	ct.span.SetAttributes(
		attribute.String(tracingAttrHttpAddressRemote, remoteAddr),
		attribute.String(tracingAttrHttpAddressLocal, localAddr),
	)
}

// PutIdleConn 在连接被返回到空闲池时被调用。
// 如果 err 为 nil，表示连接已成功归还到空闲池；
// 如果 err 非 nil，则表示连接未能归还的原因。
// 如果通过 Transport.DisableKeepAlives 禁用了连接复用，那么不会调用 PutIdleConn。
// PutIdleConn 在调用者的 Response.Body.Close 方法返回之前被调用。
// 对于 HTTP/2，当前此钩子尚未使用。
// md5:fe7db4b427b3ca9a
// ff:
// ct:
// err:
func (ct *clientTracerTracing) PutIdleConn(err error) {
	if err != nil {
		ct.span.SetStatus(codes.Error, fmt.Sprintf(`%+v`, err))
	}
}

// GotFirstResponseByte 在响应头部的第一个字节可用时被调用。
// md5:9c6cad13171a6268
// ff:
// ct:
func (ct *clientTracerTracing) GotFirstResponseByte() {}

// 如果服务器回复"100继续"响应，将调用Got100Continue。
// md5:e3d8588650aed762
// ff:
// ct:
func (ct *clientTracerTracing) Got100Continue() {}

// Got1xxResponse 在返回最终非1xx响应之前，为每个1xx信息性响应头被调用。即使也定义了Got100Continue，它也会为"100 Continue"响应被调用。如果它返回一个错误，客户端请求将使用该错误值终止。
// md5:4150bad8ebc631f8
// ff:
// ct:
// code:
// header:
func (ct *clientTracerTracing) Got1xxResponse(code int, header textproto.MIMEHeader) error {
	return nil
}

// DNSStart 在开始进行DNS查询时被调用。 md5:83ac5e2521d4c870
// ff:
// ct:
// info:
func (ct *clientTracerTracing) DNSStart(info httptrace.DNSStartInfo) {
	ct.span.SetAttributes(
		attribute.String(tracingAttrHttpDnsStart, info.Host),
	)
}

// DNSDone 在DNS查询结束时被调用。 md5:89b2144439dd0bf7
// ff:
// ct:
// info:
func (ct *clientTracerTracing) DNSDone(info httptrace.DNSDoneInfo) {
	var buffer strings.Builder
	for _, v := range info.Addrs {
		if buffer.Len() != 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(v.String())
	}
	if info.Err != nil {
		ct.span.SetStatus(codes.Error, fmt.Sprintf(`%+v`, info.Err))
	}
	ct.span.SetAttributes(
		attribute.String(tracingAttrHttpDnsDone, buffer.String()),
	)
}

// ConnectStart 在新的连接开始时被调用。
// 如果开启了 net.Dialer.DualStack（IPv6“快乐眼睛”）支持，这可能会多次被调用。
// md5:3d2fb8f45b11e3f6
// ff:
// ct:
// network:
// addr:
func (ct *clientTracerTracing) ConnectStart(network, addr string) {
	ct.span.SetAttributes(
		attribute.String(tracingAttrHttpConnectStart, network+"@"+addr),
	)
}

// ConnectDone 在新的连接调用 Dial 完成时被调用。提供的 err 表示连接是否成功完成。
// 如果启用了 net.Dialer.DualStack（"快乐眼睛"）支持，这可能会多次调用。
// md5:d12afb735b1333e9
// ff:
// ct:
// network:
// addr:
// err:
func (ct *clientTracerTracing) ConnectDone(network, addr string, err error) {
	if err != nil {
		ct.span.SetStatus(codes.Error, fmt.Sprintf(`%+v`, err))
	}
	ct.span.SetAttributes(
		attribute.String(tracingAttrHttpConnectDone, network+"@"+addr),
	)
}

// TLSHandshakeStart 在TLS握手开始时被调用。当通过HTTP代理连接到HTTPS站点时，握手发生在代理处理CONNECT请求之后。
// md5:6fc9716e1e4ce6ab
// ff:
// ct:
func (ct *clientTracerTracing) TLSHandshakeStart() {}

// TLSHandshakeDone 在TLS握手完成后被调用，参数可能是成功的握手连接状态，或者在握手失败时传入一个非nil的错误。
// md5:8d12c1d5d2e869b2
// ff:
// ct:
// _:
// err:
func (ct *clientTracerTracing) TLSHandshakeDone(_ tls.ConnectionState, err error) {
	if err != nil {
		ct.span.SetStatus(codes.Error, fmt.Sprintf(`%+v`, err))
	}
}

// WroteHeaderField 在Transport 写入每个请求头后被调用。在该调用时，这些值可能已经被缓冲，尚未写入网络。
// md5:8eab4965618e56fc
// ff:
// ct:
// k:
// v:
func (ct *clientTracerTracing) WroteHeaderField(k string, v []string) {
	if len(v) > 1 {
		ct.headers[k] = v
	} else if len(v) == 1 {
		ct.headers[k] = v[0]
	}
}

// WroteHeaders在Transport写入所有请求头后被调用。
// md5:57022d5551d5b8c3
// ff:
// ct:
func (ct *clientTracerTracing) WroteHeaders() {}

// Wait100Continue 在以下情况下被调用：请求中指定了 "Expect: 100-continue"，并且传输层已经写入了请求头，但在写入请求体之前正在等待服务器返回 "100 Continue"。
// md5:5e03081e885d021d
// ff:
// ct:
func (ct *clientTracerTracing) Wait100Continue() {}

// WroteRequest 在写入请求和任何主体后被调用，可能会在重试请求时多次调用。
// md5:0b40d83812c12716
// ff:
// ct:
// info:
func (ct *clientTracerTracing) WroteRequest(info httptrace.WroteRequestInfo) {
	if info.Err != nil {
		ct.span.SetStatus(codes.Error, fmt.Sprintf(`%+v`, info.Err))
	}

	reqBodyContent, err := gtrace.SafeContentForHttp(ct.requestBody, ct.request.Header)
	if err != nil {
		ct.span.SetStatus(codes.Error, fmt.Sprintf(`converting safe content failed: %s`, err.Error()))
	}

	ct.span.AddEvent(tracingEventHttpRequest, trace.WithAttributes(
		attribute.String(tracingEventHttpRequestHeaders, gconv.String(ct.headers)),
		attribute.String(tracingEventHttpRequestBaggage, gtrace.GetBaggageMap(ct.Context).String()),
		attribute.String(tracingEventHttpRequestBody, reqBodyContent),
	))
}
