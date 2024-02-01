
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// SetBrowserMode enables browser mode of the client.
// When browser mode is enabled, it automatically saves and sends cookie content
// from and to server.
<原文结束>

# <翻译开始>
// SetBrowserMode启用客户端的浏览器模式。
// 当浏览器模式被启用后，它会自动保存并从服务器发送、接收cookie内容。
# <翻译结束>


<原文开始>
// SetHeader sets a custom HTTP header pair for the client.
<原文结束>

# <翻译开始>
// SetHeader 为客户端设置自定义HTTP头键值对。
# <翻译结束>


<原文开始>
// SetHeaderMap sets custom HTTP headers with map.
<原文结束>

# <翻译开始>
// SetHeaderMap 通过映射设置自定义HTTP头。
# <翻译结束>


<原文开始>
// SetAgent sets the User-Agent header for client.
<原文结束>

# <翻译开始>
// SetAgent 设置客户端的 User-Agent 头部信息。
# <翻译结束>


<原文开始>
// SetContentType sets HTTP content type for the client.
<原文结束>

# <翻译开始>
// SetContentType 为客户端设置HTTP内容类型。
# <翻译结束>


<原文开始>
// SetHeaderRaw sets custom HTTP header using raw string.
<原文结束>

# <翻译开始>
// SetHeaderRaw 通过原始字符串设置自定义HTTP头。
# <翻译结束>


<原文开始>
// SetCookie sets a cookie pair for the client.
<原文结束>

# <翻译开始>
// SetCookie 为客户端设置一个cookie对。
# <翻译结束>


<原文开始>
// SetCookieMap sets cookie items with map.
<原文结束>

# <翻译开始>
// SetCookieMap 通过map设置cookie项目。
# <翻译结束>


<原文开始>
// SetPrefix sets the request server URL prefix.
<原文结束>

# <翻译开始>
// SetPrefix 设置请求服务器 URL 前缀。
# <翻译结束>


<原文开始>
// SetTimeout sets the request timeout for the client.
<原文结束>

# <翻译开始>
// SetTimeout 设置客户端的请求超时时间。
# <翻译结束>


<原文开始>
// SetBasicAuth sets HTTP basic authentication information for the client.
<原文结束>

# <翻译开始>
// SetBasicAuth为客户端设置HTTP基础认证信息。
# <翻译结束>


<原文开始>
// SetRetry sets retry count and interval.
<原文结束>

# <翻译开始>
// SetRetry 设置重试次数和间隔。
# <翻译结束>


<原文开始>
// SetRedirectLimit limits the number of jumps.
<原文结束>

# <翻译开始>
// SetRedirectLimit 限制跳转次数。
# <翻译结束>


<原文开始>
// SetNoUrlEncode sets the mark that do not encode the parameters before sending request.
<原文结束>

# <翻译开始>
// SetNoUrlEncode 设置标记，表示在发送请求前不对参数进行URL编码。
# <翻译结束>


<原文开始>
// SetProxy set proxy for the client.
// This func will do nothing when the parameter `proxyURL` is empty or in wrong pattern.
// The correct pattern is like `http://USER:PASSWORD@IP:PORT` or `socks5://USER:PASSWORD@IP:PORT`.
// Only `http` and `socks5` proxies are supported currently.
<原文结束>

# <翻译开始>
// SetProxy 为客户端设置代理。
// 当参数`proxyURL`为空或者格式不正确时，此函数将不做任何操作。
// 正确的格式应如 `http://USER:PASSWORD@IP:PORT` 或 `socks5://USER:PASSWORD@IP:PORT`。
// 目前仅支持 `http` 和 `socks5` 类型的代理。
# <翻译结束>


<原文开始>
// refer to the source code, error is always nil
<原文结束>

# <翻译开始>
// 参考源代码，error 值始终为 nil
# <翻译结束>







<原文开始>
// SetTLSKeyCrt sets the certificate and key file for TLS configuration of client.
<原文结束>

# <翻译开始>
// SetTLSKeyCrt 用于设置客户端TLS配置所需的证书和密钥文件。
# <翻译结束>


<原文开始>
// SetTLSConfig sets the TLS configuration of client.
<原文结束>

# <翻译开始>
// SetTLSConfig 设置客户端的 TLS 配置。
# <翻译结束>


<原文开始>
// SetBuilder sets the load balance builder for client.
<原文结束>

# <翻译开始>
// SetBuilder 为客户端设置负载均衡器生成器。
# <翻译结束>


<原文开始>
// SetDiscovery sets the load balance builder for client.
<原文结束>

# <翻译开始>
// SetDiscovery 设置客户端的负载均衡器生成器。
# <翻译结束>


<原文开始>
// c.SetTimeout(10*time.Second)
<原文结束>

# <翻译开始>
// c.SetTimeout(10*time.Second) // 设置c的超时时间为10秒钟
# <翻译结束>

