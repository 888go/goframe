
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31
# <翻译结束>


<原文开始>
// Response is the http response manager.
// Note that it implements the http.ResponseWriter interface with buffering feature.
<原文结束>

# <翻译开始>
// Response 是HTTP响应管理器。
// 注意它实现了带有缓冲功能的http.ResponseWriter接口。
// md5:897398e62eaf56fc
# <翻译结束>


<原文开始>
// Underlying ResponseWriter.
<原文结束>

# <翻译开始>
// 基础的 ResponseWriter。 md5:edecebd8a0d4cf02
# <翻译结束>


<原文开始>
// newResponse creates and returns a new Response object.
<原文结束>

# <翻译开始>
// newResponse 创建并返回一个新的 Response 对象。 md5:b2d8b0e3f410571c
# <翻译结束>


<原文开始>
// ServeFile serves the file to the response.
<原文结束>

# <翻译开始>
// ServeFile 向响应中发送文件。 md5:e5a83a4dd0cadaf6
# <翻译结束>


<原文开始>
// ServeFileDownload serves file downloading to the response.
<原文结束>

# <翻译开始>
// ServeFileDownload 用于将文件下载服务响应到请求。 md5:b5e9e8b76f0afca0
# <翻译结束>


<原文开始>
// RedirectTo redirects the client to another location.
// The optional parameter `code` specifies the http status code for redirecting,
// which commonly can be 301 or 302. It's 302 in default.
<原文结束>

# <翻译开始>
// RedirectTo 将客户端重定向到另一个位置。
// 可选参数 `code` 指定重定向的HTTP状态码，通常可以是301或302。默认为302。
// md5:ba008c02151efa61
# <翻译结束>


<原文开始>
// RedirectBack redirects the client back to referer.
// The optional parameter `code` specifies the http status code for redirecting,
// which commonly can be 301 or 302. It's 302 in default.
<原文结束>

# <翻译开始>
// RedirectBack 将客户端重定向回引荐来源。
// 可选参数 `code` 指定了用于重定向的HTTP状态码，
// 常见的可选值有301或302，默认情况下使用302。
// md5:b52d05fd1d742c11
# <翻译结束>


<原文开始>
// ServeContent replies to the request using the content in the
// provided ReadSeeker. The main benefit of ServeContent over io.Copy
// is that it handles Range requests properly, sets the MIME type, and
// handles If-Match, If-Unmodified-Since, If-None-Match, If-Modified-Since,
// and If-Range requests.
//
// See http.ServeContent
<原文结束>

# <翻译开始>
// ServeContent 使用提供的 ReadSeeker 中的内容回复请求。ServeContent 相较于 io.Copy 的主要优点是它能正确处理范围请求，设置 MIME 类型，并处理 If-Match, If-Unmodified-Since, If-None-Match, If-Modified-Since 和 If-Range 请求。
//
// 参考 http.ServeContent
// md5:935db9add8e4232c
# <翻译结束>


<原文开始>
// Flush outputs the buffer content to the client and clears the buffer.
<原文结束>

# <翻译开始>
// Flush 将缓冲区的内容输出到客户端并清空缓冲区。 md5:16e9c330d696be4e
# <翻译结束>

