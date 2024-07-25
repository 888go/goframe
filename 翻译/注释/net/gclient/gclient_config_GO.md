
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// SetBrowserMode enables browser mode of the client.
// When browser mode is enabled, it automatically saves and sends cookie content
// from and to server.
<原文结束>

# <翻译开始>
// SetBrowserMode 启用客户端的浏览器模式。
// 当浏览器模式启用时，它会自动保存并从客户端向服务器发送cookie内容以及从服务器接收cookie内容。 md5:00d8775b045e9c09
# <翻译结束>


<原文开始>
// SetHeader sets a custom HTTP header pair for the client.
<原文结束>

# <翻译开始>
// SetHeader 为客户端设置自定义的 HTTP 头部对。 md5:adc9509c3dab54ca
# <翻译结束>


<原文开始>
// SetHeaderMap sets custom HTTP headers with map.
<原文结束>

# <翻译开始>
// SetHeaderMap 使用映射设置自定义HTTP头。 md5:466373137e3ccd66
# <翻译结束>


<原文开始>
// SetAgent sets the User-Agent header for client.
<原文结束>

# <翻译开始>
// SetAgent 设置客户端的 User-Agent 头部。 md5:1ec87db52b5537ba
# <翻译结束>


<原文开始>
// SetContentType sets HTTP content type for the client.
<原文结束>

# <翻译开始>
// SetContentType 为客户端设置HTTP内容类型。 md5:063d3cafd0626b0a
# <翻译结束>


<原文开始>
// SetHeaderRaw sets custom HTTP header using raw string.
<原文结束>

# <翻译开始>
// SetHeaderRaw 使用原始字符串设置自定义HTTP头。 md5:e15c66308baf6cd5
# <翻译结束>


<原文开始>
// SetCookie sets a cookie pair for the client.
<原文结束>

# <翻译开始>
// SetCookie 为客户端设置一个 cookie 对。 md5:656700fcca56fb72
# <翻译结束>


<原文开始>
// SetCookieMap sets cookie items with map.
<原文结束>

# <翻译开始>
// SetCookieMap 使用映射设置Cookie项。 md5:3abd18bc89684efb
# <翻译结束>


<原文开始>
// SetPrefix sets the request server URL prefix.
<原文结束>

# <翻译开始>
// SetPrefix 设置请求服务器的URL前缀。 md5:945a0fd6f4acac16
# <翻译结束>


<原文开始>
// SetTimeout sets the request timeout for the client.
<原文结束>

# <翻译开始>
// SetTimeout 设置客户端的请求超时时间。 md5:ce4f874cd14c1c2d
# <翻译结束>


<原文开始>
// SetBasicAuth sets HTTP basic authentication information for the client.
<原文结束>

# <翻译开始>
// SetBasicAuth 为客户端设置HTTP基本认证信息。 md5:22c36a5363199cd0
# <翻译结束>


<原文开始>
// SetRetry sets retry count and interval.
// TODO removed.
<原文结束>

# <翻译开始>
// SetRetry 设置重试次数和间隔。
// TODO：移除。 md5:1089293b9f9371f0
# <翻译结束>


<原文开始>
// SetRedirectLimit limits the number of jumps.
<原文结束>

# <翻译开始>
// SetRedirectLimit 限制跳转次数。 md5:14e010f8e3d003b5
# <翻译结束>


<原文开始>
// SetNoUrlEncode sets the mark that do not encode the parameters before sending request.
<原文结束>

# <翻译开始>
// SetNoUrlEncode 设置标记，表示在发送请求之前不编码参数。 md5:6dd55f5543918206
# <翻译结束>


<原文开始>
// SetProxy set proxy for the client.
// This func will do nothing when the parameter `proxyURL` is empty or in wrong pattern.
// The correct pattern is like `http://USER:PASSWORD@IP:PORT` or `socks5://USER:PASSWORD@IP:PORT`.
// Only `http` and `socks5` proxies are supported currently.
<原文结束>

# <翻译开始>
// SetProxy 为客户端设置代理。
// 当参数 `proxyURL` 为空或格式不正确时，此函数将不会执行任何操作。
// 正确的格式应为 `http://用户名:密码@IP:端口` 或 `socks5://用户名:密码@IP:端口`。
// 目前仅支持 `http` 和 `socks5` 类型的代理。 md5:aa3f2b21308c7bec
# <翻译结束>


<原文开始>
// refer to the source code, error is always nil
<原文结束>

# <翻译开始>
		// 参考源代码，错误始终为nil. md5:43df5b2c264029cb
# <翻译结束>


<原文开始>
// c.SetTimeout(10*time.Second)
<原文结束>

# <翻译开始>
		// 设置超时时间为10秒. md5:ee88d389b4a64b4a
# <翻译结束>


<原文开始>
// SetTLSKeyCrt sets the certificate and key file for TLS configuration of client.
<原文结束>

# <翻译开始>
// SetTLSKeyCrt 设置客户端TLS配置的证书和密钥文件。 md5:48b3322243e8e691
# <翻译结束>


<原文开始>
// SetTLSConfig sets the TLS configuration of client.
<原文结束>

# <翻译开始>
// SetTLSConfig 设置客户端的TLS配置。 md5:f1882ff235302c92
# <翻译结束>


<原文开始>
// SetBuilder sets the load balance builder for client.
<原文结束>

# <翻译开始>
// SetBuilder 设置客户端的负载均衡构建器。 md5:1f374a9a600309bb
# <翻译结束>


<原文开始>
// SetDiscovery sets the load balance builder for client.
<原文结束>

# <翻译开始>
// SetDiscovery 为客户端设置负载均衡构建器。 md5:0ea9a7eaf5c235e7
# <翻译结束>

