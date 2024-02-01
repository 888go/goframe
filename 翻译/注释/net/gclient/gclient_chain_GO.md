
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
// Prefix is a chaining function,
// which sets the URL prefix for next request of this client.
// Eg:
// Prefix("http://127.0.0.1:8199/api/v1")
// Prefix("http://127.0.0.1:8199/api/v2")
<原文结束>

# <翻译开始>
// Prefix 是一个链式函数，
// 它用于设置该客户端下一次请求的URL前缀。
// 示例：
// Prefix("http://127.0.0.1:8199/api/v1") // 设置URL前缀为v1版本
// Prefix("http://127.0.0.1:8199/api/v2") // 设置URL前缀为v2版本
# <翻译结束>


<原文开始>
// Header is a chaining function,
// which sets custom HTTP headers with map for next request.
<原文结束>

# <翻译开始>
// Header 是一个链式函数，
// 用于为下一个请求设置自定义HTTP头部，参数为映射(map)类型。
# <翻译结束>


<原文开始>
// HeaderRaw is a chaining function,
// which sets custom HTTP header using raw string for next request.
<原文结束>

# <翻译开始>
// HeaderRaw 是一个链式函数，
// 它用于设置下一次请求的自定义HTTP头部，使用原始字符串形式。
# <翻译结束>


<原文开始>
// Discovery is a chaining function, which sets the discovery for client.
// You can use `Discovery(nil)` to disable discovery feature for current client.
<原文结束>

# <翻译开始>
// Discovery 是一个链式函数，用于为客户端设置发现机制。
// 你可以使用 `Discovery(nil)` 来禁用当前客户端的发现功能。
# <翻译结束>


<原文开始>
// Cookie is a chaining function,
// which sets cookie items with map for next request.
<原文结束>

# <翻译开始>
// Cookie 是一个链式函数，
// 用于为下一次请求通过映射设置cookie项。
# <翻译结束>


<原文开始>
// ContentType is a chaining function,
// which sets HTTP content type for the next request.
<原文结束>

# <翻译开始>
// ContentType 是一个链式函数，
// 用于为下一个请求设置HTTP内容类型。
# <翻译结束>


<原文开始>
// ContentJson is a chaining function,
// which sets the HTTP content type as "application/json" for the next request.
//
// Note that it also checks and encodes the parameter to JSON format automatically.
<原文结束>

# <翻译开始>
// ContentJson 是一个链式函数，
// 它用于为下一个请求设置HTTP内容类型为 "application/json"。
//
// 注意，它还会自动检查并以JSON格式对参数进行编码。
# <翻译结束>


<原文开始>
// ContentXml is a chaining function,
// which sets the HTTP content type as "application/xml" for the next request.
//
// Note that it also checks and encodes the parameter to XML format automatically.
<原文结束>

# <翻译开始>
// ContentXml 是一个链式函数，
// 它为下一次请求设置HTTP内容类型为"application/xml"。
//
// 注意，它还会自动检查并以XML格式对参数进行编码。
# <翻译结束>


<原文开始>
// Timeout is a chaining function,
// which sets the timeout for next request.
<原文结束>

# <翻译开始>
// Timeout 是一个链式函数，
// 用于设置下一次请求的超时时间。
# <翻译结束>


<原文开始>
// BasicAuth is a chaining function,
// which sets HTTP basic authentication information for next request.
<原文结束>

# <翻译开始>
// BasicAuth 是一个链式函数，
// 用于为下一次请求设置HTTP基本认证信息。
# <翻译结束>


<原文开始>
// Retry is a chaining function,
// which sets retry count and interval when failure for next request.
<原文结束>

# <翻译开始>
// Retry 是一个链式函数，
// 它在下一次请求失败时设置重试次数和间隔。
# <翻译结束>


<原文开始>
// Proxy is a chaining function,
// which sets proxy for next request.
// Make sure you pass the correct `proxyURL`.
// The correct pattern is like `http://USER:PASSWORD@IP:PORT` or `socks5://USER:PASSWORD@IP:PORT`.
// Only `http` and `socks5` proxies are supported currently.
<原文结束>

# <翻译开始>
// Proxy 是一个链式函数，
// 用于设置下一次请求的代理。
// 请确保传递正确的 `proxyURL`。
// 正确的格式应为 `http://USER:PASSWORD@IP:PORT` 或 `socks5://USER:PASSWORD@IP:PORT`。
// 目前仅支持 `http` 和 `socks5` 类型的代理。
# <翻译结束>


<原文开始>
// RedirectLimit is a chaining function,
// which sets the redirect limit the number of jumps for the request.
<原文结束>

# <翻译开始>
// RedirectLimit 是一个链式函数，
// 用于设置请求的重定向限制次数。
# <翻译结束>


<原文开始>
// NoUrlEncode sets the mark that do not encode the parameters before sending request.
<原文结束>

# <翻译开始>
// NoUrlEncode 设置标记，表示在发送请求前不应对参数进行URL编码。
# <翻译结束>

