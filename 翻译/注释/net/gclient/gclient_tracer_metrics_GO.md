
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// newClientTracerMetrics creates and returns object of httptrace.ClientTrace.
<原文结束>

# <翻译开始>
// newClientTracerMetrics 创建并返回一个httptrace.ClientTrace对象。. md5:ff0ba189cef9c033
# <翻译结束>


<原文开始>
// GetConn is called before a connection is created or
// retrieved from an idle pool. The hostPort is the
// "host:port" of the target or proxy. GetConn is called even
// if there's already an idle cached connection available.
<原文结束>

# <翻译开始>
// 在创建或从空闲池中检索连接之前，会调用GetConn。hostPort是目标或代理的"主机:端口"。即使已经有空闲缓存的连接可用，也会调用GetConn。
// md5:de9f40c04cc79ecf
# <翻译结束>


<原文开始>
// GotConn is called after a successful connection is
// obtained. There is no hook for failure to obtain a
// connection; instead, use the error from
// Transport.RoundTrip.
<原文结束>

# <翻译开始>
// GotConn 在成功建立连接后被调用。如果无法获取连接，没有相应的钩子；相反，应使用 Transport.RoundTrip 的错误。
// md5:939ef1676a19b2be
# <翻译结束>


<原文开始>
// PutIdleConn is called when the connection is returned to
// the idle pool. If err is nil, the connection was
// successfully returned to the idle pool. If err is non-nil,
// it describes why not. PutIdleConn is not called if
// connection reuse is disabled via Transport.DisableKeepAlives.
// PutIdleConn is called before the caller's Response.Body.Close
// call returns.
// For HTTP/2, this hook is not currently used.
<原文结束>

# <翻译开始>
// PutIdleConn 在连接被返回到空闲池时被调用。
// 如果 err 为 nil，表示连接已成功归还到空闲池；
// 如果 err 非 nil，则表示连接未能归还的原因。
// 如果通过 Transport.DisableKeepAlives 禁用了连接复用，那么不会调用 PutIdleConn。
// PutIdleConn 在调用者的 Response.Body.Close 方法返回之前被调用。
// 对于 HTTP/2，当前此钩子尚未使用。
// md5:fe7db4b427b3ca9a
# <翻译结束>


<原文开始>
// GotFirstResponseByte is called when the first byte of the response
// headers is available.
<原文结束>

# <翻译开始>
// GotFirstResponseByte 在响应头部的第一个字节可用时被调用。
// md5:9c6cad13171a6268
# <翻译结束>


<原文开始>
// Got100Continue is called if the server replies with a "100
// Continue" response.
<原文结束>

# <翻译开始>
// 如果服务器回复"100继续"响应，将调用Got100Continue。
// md5:e3d8588650aed762
# <翻译结束>


<原文开始>
// Got1xxResponse is called for each 1xx informational response header
// returned before the final non-1xx response. Got1xxResponse is called
// for "100 Continue" responses, even if Got100Continue is also defined.
// If it returns an error, the client request is aborted with that error value.
<原文结束>

# <翻译开始>
// Got1xxResponse 在返回最终非1xx响应之前，为每个1xx信息性响应头被调用。即使也定义了Got100Continue，它也会为"100 Continue"响应被调用。如果它返回一个错误，客户端请求将使用该错误值终止。
// md5:4150bad8ebc631f8
# <翻译结束>


<原文开始>
// DNSStart is called when a DNS lookup begins.
<原文结束>

# <翻译开始>
// DNSStart 在开始进行DNS查询时被调用。. md5:83ac5e2521d4c870
# <翻译结束>


<原文开始>
// DNSDone is called when a DNS lookup ends.
<原文结束>

# <翻译开始>
// DNSDone 在DNS查询结束时被调用。. md5:89b2144439dd0bf7
# <翻译结束>


<原文开始>
// ConnectStart is called when a new connection's Dial begins.
// If net.Dialer.DualStack (IPv6 "Happy Eyeballs") support is
// enabled, this may be called multiple times.
<原文结束>

# <翻译开始>
// ConnectStart 在新的连接开始时被调用。
// 如果开启了 net.Dialer.DualStack（IPv6“快乐眼睛”）支持，这可能会多次被调用。
// md5:3d2fb8f45b11e3f6
# <翻译结束>


<原文开始>
// ConnectDone is called when a new connection's Dial
// completes. The provided err indicates whether the
// connection completed successfully.
// If net.Dialer.DualStack ("Happy Eyeballs") support is
// enabled, this may be called multiple times.
<原文结束>

# <翻译开始>
// ConnectDone 在新的连接调用 Dial 完成时被调用。提供的 err 表示连接是否成功完成。
// 如果启用了 net.Dialer.DualStack（"快乐眼睛"）支持，这可能会多次调用。
// md5:d12afb735b1333e9
# <翻译结束>


<原文开始>
// TLSHandshakeStart is called when the TLS handshake is started. When
// connecting to an HTTPS site via an HTTP proxy, the handshake happens
// after the CONNECT request is processed by the proxy.
<原文结束>

# <翻译开始>
// TLSHandshakeStart 在TLS握手开始时被调用。当通过HTTP代理连接到HTTPS站点时，握手发生在代理处理CONNECT请求之后。
// md5:6fc9716e1e4ce6ab
# <翻译结束>


<原文开始>
// TLSHandshakeDone is called after the TLS handshake with either the
// successful handshake's connection state, or a non-nil error on handshake
// failure.
<原文结束>

# <翻译开始>
// TLSHandshakeDone 在TLS握手完成后被调用，参数可能是成功的握手连接状态，或者在握手失败时传入一个非nil的错误。
// md5:8d12c1d5d2e869b2
# <翻译结束>


<原文开始>
// WroteHeaderField is called after the Transport has written
// each request header. At the time of this call the values
// might be buffered and not yet written to the network.
<原文结束>

# <翻译开始>
// WroteHeaderField 在Transport 写入每个请求头后被调用。在该调用时，这些值可能已经被缓冲，尚未写入网络。
// md5:8eab4965618e56fc
# <翻译结束>


<原文开始>
// WroteHeaders is called after the Transport has written
// all request headers.
<原文结束>

# <翻译开始>
// WroteHeaders在Transport写入所有请求头后被调用。
// md5:57022d5551d5b8c3
# <翻译结束>


<原文开始>
// Wait100Continue is called if the Request specified
// "Expect: 100-continue" and the Transport has written the
// request headers but is waiting for "100 Continue" from the
// server before writing the request body.
<原文结束>

# <翻译开始>
// Wait100Continue 在以下情况下被调用：请求中指定了 "Expect: 100-continue"，并且传输层已经写入了请求头，但在写入请求体之前正在等待服务器返回 "100 Continue"。
// md5:5e03081e885d021d
# <翻译结束>


<原文开始>
// WroteRequest is called with the result of writing the
// request and any body. It may be called multiple times
// in the case of retried requests.
<原文结束>

# <翻译开始>
// WroteRequest 在写入请求和任何主体后被调用，可能会在重试请求时多次调用。
// md5:0b40d83812c12716
# <翻译结束>

