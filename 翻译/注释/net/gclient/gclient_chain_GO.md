
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
// Prefix is a chaining function,
// which sets the URL prefix for next request of this client.
// Eg:
// Prefix("http://127.0.0.1:8199/api/v1")
// Prefix("http://127.0.0.1:8199/api/v2")
<原文结束>

# <翻译开始>
// Prefix是一个链式函数，
// 它为该客户端的下一个请求设置URL前缀。例如：
// Prefix("http://127.0.0.1:8199/api/v1")
// Prefix("http://127.0.0.1:8199/api/v2")
// md5:223b00b33201dec6
# <翻译结束>


<原文开始>
// Header is a chaining function,
// which sets custom HTTP headers with map for next request.
<原文结束>

# <翻译开始>
// Header是一个链式函数，它为下一个请求设置自定义HTTP头，使用映射（map）来存储头信息。
// md5:e9d44530a2916792
# <翻译结束>


<原文开始>
// HeaderRaw is a chaining function,
// which sets custom HTTP header using raw string for next request.
<原文结束>

# <翻译开始>
// HeaderRaw 是一个链式函数，
// 用于为下一个请求使用原始字符串设置自定义的HTTP头。
// md5:dff868fea39b9738
# <翻译结束>


<原文开始>
// Discovery is a chaining function, which sets the discovery for client.
// You can use `Discovery(nil)` to disable discovery feature for current client.
<原文结束>

# <翻译开始>
// Discovery 是一个链式函数，用于设置客户端的发现机制。
// 你可以使用 `Discovery(nil)` 来禁用当前客户端的发现功能。
// md5:75d0bac47755ed2a
# <翻译结束>


<原文开始>
// Cookie is a chaining function,
// which sets cookie items with map for next request.
<原文结束>

# <翻译开始>
// Cookie是一个链式函数，它使用映射为下一个请求设置Cookie项。
// md5:458381c5ded10dae
# <翻译结束>


<原文开始>
// ContentType is a chaining function,
// which sets HTTP content type for the next request.
<原文结束>

# <翻译开始>
// ContentType是一个链式函数，用于为下一个请求设置HTTP内容类型。
// md5:48b017c4d22a94ee
# <翻译结束>


<原文开始>
// ContentJson is a chaining function,
// which sets the HTTP content type as "application/json" for the next request.
//
// Note that it also checks and encodes the parameter to JSON format automatically.
<原文结束>

# <翻译开始>
// ContentJson 是一个链式函数，
// 用于为下一个请求将HTTP内容类型设置为 "application/json"。
//
// 注意，它还会自动检查并以JSON格式对参数进行编码。
// md5:3b265101262bf8be
# <翻译结束>


<原文开始>
// ContentXml is a chaining function,
// which sets the HTTP content type as "application/xml" for the next request.
//
// Note that it also checks and encodes the parameter to XML format automatically.
<原文结束>

# <翻译开始>
// ContentXml是一个链式函数，
// 它将下一个请求的HTTP内容类型设置为"application/xml"。
//
// 请注意，它还会自动检查并把参数编码为XML格式。
// md5:b6e6ad51561fa7a6
# <翻译结束>


<原文开始>
// Timeout is a chaining function,
// which sets the timeout for next request.
<原文结束>

# <翻译开始>
// Timeout是一个链式函数，它为下一个请求设置超时时间。
// md5:8198ec107ce0113f
# <翻译结束>


<原文开始>
// BasicAuth is a chaining function,
// which sets HTTP basic authentication information for next request.
<原文结束>

# <翻译开始>
// BasicAuth是一个链式函数，为下一个请求设置HTTP基本认证信息。
// md5:c0ca33e41692898b
# <翻译结束>


<原文开始>
// Retry is a chaining function,
// which sets retry count and interval when failure for next request.
// TODO removed.
<原文结束>

# <翻译开始>
// Retry 是一个链式函数，
// 用于设置在下次请求失败时的重试次数和间隔时间。
// TODO 待移除。
// md5:2206b70379c77ed7
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
// 用于为下一个请求设置代理。
// 确保传递正确的 `proxyURL`。
// 正确的格式应为 `http://USER:PASSWORD@IP:PORT` 或 `socks5://USER:PASSWORD@IP:PORT`。
// 目前仅支持 `http` 和 `socks5` 类型的代理。
// md5:d8c660f7a12e72ea
# <翻译结束>


<原文开始>
// RedirectLimit is a chaining function,
// which sets the redirect limit the number of jumps for the request.
<原文结束>

# <翻译开始>
// RedirectLimit 是一个链式函数，用于设置请求的重定向限制次数。
// md5:ecac398510aa92bb
# <翻译结束>


<原文开始>
// NoUrlEncode sets the mark that do not encode the parameters before sending request.
<原文结束>

# <翻译开始>
// NoUrlEncode 设置一个标记，表示在发送请求之前不进行参数编码。. md5:0f78cc83f0909b0e
# <翻译结束>

