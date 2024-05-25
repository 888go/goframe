// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gclient

import (
	"crypto/tls"
	"net/http/httptrace"
	"net/textproto"
)

type clientTracer struct {
	*httptrace.ClientTrace
}

// newClientTracer 创建并返回一个httptrace.ClientTrace对象。. md5:7c929ccddc1401b4
func newClientTracer(baseClientTracer *httptrace.ClientTrace) *httptrace.ClientTrace {
	c := &clientTracer{
		ClientTrace: baseClientTracer,
	}
	return &httptrace.ClientTrace{
		GetConn:              c.GetConn,
		GotConn:              c.GotConn,
		PutIdleConn:          c.PutIdleConn,
		GotFirstResponseByte: c.GotFirstResponseByte,
		Got100Continue:       c.Got100Continue,
		Got1xxResponse:       c.Got1xxResponse,
		DNSStart:             c.DNSStart,
		DNSDone:              c.DNSDone,
		ConnectStart:         c.ConnectStart,
		ConnectDone:          c.ConnectDone,
		TLSHandshakeStart:    c.TLSHandshakeStart,
		TLSHandshakeDone:     c.TLSHandshakeDone,
		WroteHeaderField:     c.WroteHeaderField,
		WroteHeaders:         c.WroteHeaders,
		Wait100Continue:      c.Wait100Continue,
		WroteRequest:         c.WroteRequest,
	}
}

// 在创建或从空闲池中检索连接之前，会调用GetConn。hostPort是目标或代理的"主机:端口"。即使已经有空闲缓存的连接可用，也会调用GetConn。
// md5:de9f40c04cc79ecf
func (ct *clientTracer) GetConn(hostPort string) {
	ct.ClientTrace.GetConn(hostPort)
}

// GotConn 在成功建立连接后被调用。如果无法获取连接，没有相应的钩子；相反，应使用 Transport.RoundTrip 的错误。
// md5:939ef1676a19b2be
func (ct *clientTracer) GotConn(info httptrace.GotConnInfo) {
	ct.ClientTrace.GotConn(info)
}

// PutIdleConn 在连接被返回到空闲池时被调用。
// 如果 err 为 nil，表示连接已成功归还到空闲池；
// 如果 err 非 nil，则表示连接未能归还的原因。
// 如果通过 Transport.DisableKeepAlives 禁用了连接复用，那么不会调用 PutIdleConn。
// PutIdleConn 在调用者的 Response.Body.Close 方法返回之前被调用。
// 对于 HTTP/2，当前此钩子尚未使用。
// md5:fe7db4b427b3ca9a
func (ct *clientTracer) PutIdleConn(err error) {
	ct.ClientTrace.PutIdleConn(err)
}

// GotFirstResponseByte 在响应头部的第一个字节可用时被调用。
// md5:9c6cad13171a6268
func (ct *clientTracer) GotFirstResponseByte() {
	ct.ClientTrace.GotFirstResponseByte()
}

// 如果服务器回复"100继续"响应，将调用Got100Continue。
// md5:e3d8588650aed762
func (ct *clientTracer) Got100Continue() {
	ct.ClientTrace.Got100Continue()
}

// Got1xxResponse 在返回最终非1xx响应之前，为每个1xx信息性响应头被调用。即使也定义了Got100Continue，它也会为"100 Continue"响应被调用。如果它返回一个错误，客户端请求将使用该错误值终止。
// md5:4150bad8ebc631f8
func (ct *clientTracer) Got1xxResponse(code int, header textproto.MIMEHeader) error {
	return ct.ClientTrace.Got1xxResponse(code, header)
}

// DNSStart 在开始进行DNS查询时被调用。. md5:83ac5e2521d4c870
func (ct *clientTracer) DNSStart(info httptrace.DNSStartInfo) {
	ct.ClientTrace.DNSStart(info)
}

// DNSDone 在DNS查询结束时被调用。. md5:89b2144439dd0bf7
func (ct *clientTracer) DNSDone(info httptrace.DNSDoneInfo) {
	ct.ClientTrace.DNSDone(info)
}

// ConnectStart 在新的连接开始时被调用。
// 如果开启了 net.Dialer.DualStack（IPv6“快乐眼睛”）支持，这可能会多次被调用。
// md5:3d2fb8f45b11e3f6
func (ct *clientTracer) ConnectStart(network, addr string) {
	ct.ClientTrace.ConnectStart(network, addr)
}

// ConnectDone 在新的连接调用 Dial 完成时被调用。提供的 err 表示连接是否成功完成。
// 如果启用了 net.Dialer.DualStack（"快乐眼睛"）支持，这可能会多次调用。
// md5:d12afb735b1333e9
func (ct *clientTracer) ConnectDone(network, addr string, err error) {
	ct.ClientTrace.ConnectDone(network, addr, err)
}

// TLSHandshakeStart 在TLS握手开始时被调用。当通过HTTP代理连接到HTTPS站点时，握手发生在代理处理CONNECT请求之后。
// md5:6fc9716e1e4ce6ab
func (ct *clientTracer) TLSHandshakeStart() {
	ct.ClientTrace.TLSHandshakeStart()
}

// TLSHandshakeDone 在TLS握手完成后被调用，参数可能是成功的握手连接状态，或者在握手失败时传入一个非nil的错误。
// md5:8d12c1d5d2e869b2
func (ct *clientTracer) TLSHandshakeDone(state tls.ConnectionState, err error) {
	ct.ClientTrace.TLSHandshakeDone(state, err)
}

// WroteHeaderField 在Transport 写入每个请求头后被调用。在该调用时，这些值可能已经被缓冲，尚未写入网络。
// md5:8eab4965618e56fc
func (ct *clientTracer) WroteHeaderField(key string, value []string) {
	ct.ClientTrace.WroteHeaderField(key, value)
}

// WroteHeaders在Transport写入所有请求头后被调用。
// md5:57022d5551d5b8c3
func (ct *clientTracer) WroteHeaders() {
	ct.ClientTrace.WroteHeaders()
}

// Wait100Continue 在以下情况下被调用：请求中指定了 "Expect: 100-continue"，并且传输层已经写入了请求头，但在写入请求体之前正在等待服务器返回 "100 Continue"。
// md5:5e03081e885d021d
func (ct *clientTracer) Wait100Continue() {
	ct.ClientTrace.Wait100Continue()
}

// WroteRequest 在写入请求和任何主体后被调用，可能会在重试请求时多次调用。
// md5:0b40d83812c12716
func (ct *clientTracer) WroteRequest(info httptrace.WroteRequestInfo) {
	ct.ClientTrace.WroteRequest(info)
}
