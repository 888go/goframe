
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//
# <翻译结束>


<原文开始>
// Response is the http response manager.
// Note that it implements the http.ResponseWriter interface with buffering feature.
<原文结束>

# <翻译开始>
// Response 是HTTP响应管理器。
// 请注意，它实现了带有缓冲功能的http.ResponseWriter接口。
# <翻译结束>

















<原文开始>
// newResponse creates and returns a new Response object.
<原文结束>

# <翻译开始>
// newResponse 创建并返回一个新的 Response 对象。
# <翻译结束>


<原文开始>
// ServeFile serves the file to the response.
<原文结束>

# <翻译开始>
// ServeFile 将文件发送至响应。
// 会自动识别文件格式，如果是目录或者文本内容将会直接展示文件内容。
// 如果path参数为目录，那么第二个参数allowIndex控制是否可以展示目录下的文件列表。
# <翻译结束>


<原文开始>
// ServeFileDownload serves file downloading to the response.
<原文结束>

# <翻译开始>
// ServeFileDownload 为响应提供文件下载服务。
// 用于直接引导客户端下载指定路径的文件，并可以重新给定下载的文件名称。
// ServeFileDownload方法采用的是流式下载控制，对内存占用较少。
// 使用示例，我们把示例中的ServeFile方法改为ServeFileDownload方法：
// func main() {
// 	s := g.Server()
// 	s.BindHandler("/", func(r *ghttp.Request) {
// 		r.Response.ServeFileDownload("test.txt")
// 	})
// 	s.SetPort(8999)
// 	s.Run()
// }
# <翻译结束>


<原文开始>
// RedirectTo redirects the client to another location.
// The optional parameter `code` specifies the http status code for redirecting,
// which commonly can be 301 or 302. It's 302 in default.
<原文结束>

# <翻译开始>
// RedirectTo 重定向客户端到另一个位置。
// 可选参数 `code` 指定了用于重定向的 HTTP 状态码，
// 通常可以是 301 或 302，默认为 302。
# <翻译结束>


<原文开始>
// RedirectBack redirects the client back to referer.
// The optional parameter `code` specifies the http status code for redirecting,
// which commonly can be 301 or 302. It's 302 in default.
<原文结束>

# <翻译开始>
// RedirectBack 重定向客户端返回到referer页面。
// 可选参数 `code` 指定用于重定向的http状态码，通常可以是301或302，默认为302。
# <翻译结束>


<原文开始>
// Buffer returns the buffered content as []byte.
<原文结束>

# <翻译开始>
// Buffer返回缓冲区中的内容作为[]byte。
# <翻译结束>


<原文开始>
// BufferString returns the buffered content as string.
<原文结束>

# <翻译开始>
// BufferString 返回缓冲区中的内容作为字符串。
# <翻译结束>


<原文开始>
// BufferLength returns the length of the buffered content.
<原文结束>

# <翻译开始>
// BufferLength 返回缓冲区内容的长度。
# <翻译结束>


<原文开始>
// SetBuffer overwrites the buffer with `data`.
<原文结束>

# <翻译开始>
// SetBuffer 将`data`覆盖写入缓冲区。
# <翻译结束>


<原文开始>
// ClearBuffer clears the response buffer.
<原文结束>

# <翻译开始>
// 清空缓冲区 ClearBuffer 用于清空响应缓冲区。
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
// ServeContent 函数通过提供的 ReadSeeker 中的内容回复请求。与 io.Copy 相比，ServeContent 的主要优点在于它能妥善处理 Range 请求，设置 MIME 类型，并正确处理 If-Match, If-Unmodified-Since, If-None-Match, If-Modified-Since 以及 If-Range 等请求。
//
// 参见 http.ServeContent
# <翻译结束>


<原文开始>
// Flush outputs the buffer content to the client and clears the buffer.
<原文结束>

# <翻译开始>
// Flush 将缓冲区内容输出到客户端并清空缓冲区。
# <翻译结束>


<原文开始>
// Underlying ResponseWriter.
<原文结束>

# <翻译开始>
// 基础的 ResponseWriter。
# <翻译结束>


<原文开始>
// Alias of ResponseWriter.
<原文结束>

# <翻译开始>
// ResponseWriter的别名。
# <翻译结束>


<原文开始>
// According request.
<原文结束>

# <翻译开始>
// 根据请求。
# <翻译结束>

