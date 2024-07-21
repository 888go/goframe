
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
// Write writes `content` to the response buffer.
<原文结束>

# <翻译开始>
// Write 将 `content` 写入响应缓冲区。 md5:2f656734fbf8eab6
# <翻译结束>


<原文开始>
// WriteExit writes `content` to the response buffer and exits executing of current handler.
// The "Exit" feature is commonly used to replace usage of return statements in the handler,
// for convenience.
<原文结束>

# <翻译开始>
// WriteExit 将`content`写入响应缓冲区并退出当前处理器的执行。
// "Exit"功能常用于替代处理器中的返回语句，以提供便利。
// md5:afcb2dda1beb9358
# <翻译结束>


<原文开始>
// WriteOver overwrites the response buffer with `content`.
<原文结束>

# <翻译开始>
// WriteOver 将响应缓冲区用`content`重写。 md5:d68f13dc57329ab0
# <翻译结束>


<原文开始>
// WriteOverExit overwrites the response buffer with `content` and exits executing
// of current handler. The "Exit" feature is commonly used to replace usage of return
// statements in the handler, for convenience.
<原文结束>

# <翻译开始>
// WriteOverExit 将响应缓冲区用 `content` 替换，然后退出当前处理程序的执行。"Exit" 功能通常用于方便地替换处理程序中的返回语句。
// md5:968d387aea44eeab
# <翻译结束>


<原文开始>
// Writef writes the response with fmt.Sprintf.
<原文结束>

# <翻译开始>
// Writef 使用 fmt.Sprintf 格式化字符串并写出响应。 md5:15163b759bd146b8
# <翻译结束>


<原文开始>
// WritefExit writes the response with fmt.Sprintf and exits executing of current handler.
// The "Exit" feature is commonly used to replace usage of return statements in the handler,
// for convenience.
<原文结束>

# <翻译开始>
// WritefExit使用fmt.Sprintf格式化响应内容并退出当前处理器的执行。
// "Exit"功能常用于替代处理器中的返回语句，以提供便利。
// md5:01275db804fa4029
# <翻译结束>


<原文开始>
// Writeln writes the response with `content` and new line.
<原文结束>

# <翻译开始>
// Writeln 使用`content`和换行符写入响应。 md5:574e18a271a92e20
# <翻译结束>


<原文开始>
// WritelnExit writes the response with `content` and new line and exits executing
// of current handler. The "Exit" feature is commonly used to replace usage of return
// statements in the handler, for convenience.
<原文结束>

# <翻译开始>
// WritelnExit 写入包含`content`和换行符的响应，然后退出当前处理器的执行。"Exit"特性通常用于方便地替换处理器中的返回语句。
// md5:bb5f123bedaec380
# <翻译结束>


<原文开始>
// Writefln writes the response with fmt.Sprintf and new line.
<原文结束>

# <翻译开始>
// Writefln 使用 fmt.Sprintf 格式化输出并将结果作为响应写入，同时添加换行符。 md5:154e6f5f52878f00
# <翻译结束>


<原文开始>
// WriteflnExit writes the response with fmt.Sprintf and new line and exits executing
// of current handler. The "Exit" feature is commonly used to replace usage of return
// statement in the handler, for convenience.
<原文结束>

# <翻译开始>
// WriteflnExit使用fmt.Sprintf格式化响应并添加换行符，然后退出当前处理程序的执行。
// "Exit"功能常用于替代处理器中的返回语句，以提供便利。
// md5:ee5288e61cdea4b2
# <翻译结束>


<原文开始>
// WriteJson writes `content` to the response with JSON format.
<原文结束>

# <翻译开始>
// WriteJson 将`content`以JSON格式写入到响应中。 md5:0ca8d5da1805456f
# <翻译结束>


<原文开始>
// If given string/[]byte, response it directly to the client.
<原文结束>

# <翻译开始>
// 如果给定字符串/字节切片，直接将其返回给客户端。 md5:e14783864a1068a9
# <翻译结束>


<原文开始>
// Else use json.Marshal function to encode the parameter.
<原文结束>

# <翻译开始>
// 否则，使用json.Marshal函数对参数进行编码。 md5:b140f4be3fab1fa1
# <翻译结束>


<原文开始>
// WriteJsonExit writes `content` to the response with JSON format and exits executing
// of current handler if success. The "Exit" feature is commonly used to replace usage of
// return statements in the handler, for convenience.
<原文结束>

# <翻译开始>
// WriteJsonExit 将`content`以JSON格式写入响应，并在成功时退出当前处理器的执行。"Exit"功能常用于替代处理器中的返回语句，以提供便利。
// md5:0714d99528fcb93e
# <翻译结束>


<原文开始>
// WriteJsonP writes `content` to the response with JSONP format.
//
// Note that there should be a "callback" parameter in the request for JSONP format.
<原文结束>

# <翻译开始>
// WriteJsonP 将`content`按照JSONP格式写入响应。
// 
// 注意，对于JSONP格式，请求中应该包含一个名为"callback"的参数。
// md5:32a3e4fa6b4e92b0
# <翻译结束>


<原文开始>
// If given string/[]byte, response it directly to client.
<原文结束>

# <翻译开始>
// 如果给定字符串/字节切片，直接将其返回给客户端。 md5:ff82edfcddee9a78
# <翻译结束>


<原文开始>
// r.Header().Set("Content-Type", "application/json")
<原文结束>

# <翻译开始>
// 设置HTTP响应头中的"Content-Type"为"application/json". md5:8e0be5eb3c232d44
# <翻译结束>


<原文开始>
// WriteJsonPExit writes `content` to the response with JSONP format and exits executing
// of current handler if success. The "Exit" feature is commonly used to replace usage of
// return statements in the handler, for convenience.
//
// Note that there should be a "callback" parameter in the request for JSONP format.
<原文结束>

# <翻译开始>
// WriteJsonPExit 将 `content` 以 JSONP 格式写入响应，并在成功时退出当前处理器的执行。"Exit" 功能常用于替代处理器中的返回语句，以提供便利。
//
// 请注意，为了使用 JSONP 格式，请求中应该包含一个 "callback" 参数。
// md5:6c959e76945e075a
# <翻译结束>


<原文开始>
// WriteXml writes `content` to the response with XML format.
<原文结束>

# <翻译开始>
// WriteXml 将`content`以XML格式写入响应。 md5:850a872cf25d6a70
# <翻译结束>


<原文开始>
// If given string/[]byte, response it directly to clients.
<原文结束>

# <翻译开始>
// 如果给定字符串/字节切片，直接将其返回给客户端。 md5:4fc9d6cf062e5bf9
# <翻译结束>


<原文开始>
// WriteXmlExit writes `content` to the response with XML format and exits executing
// of current handler if success. The "Exit" feature is commonly used to replace usage
// of return statements in the handler, for convenience.
<原文结束>

# <翻译开始>
// WriteXmlExit 将`content`以XML格式写入响应，并在成功时退出当前处理器的执行。
// “退出”功能常用于便捷地替代处理器中return语句的使用。
// md5:12a9a20328b00f55
# <翻译结束>


<原文开始>
// WriteStatus writes HTTP `status` and `content` to the response.
// Note that it does not set a Content-Type header here.
<原文结束>

# <翻译开始>
// WriteStatus 将HTTP状态码`status`和内容`content`写入响应。
// 请注意，它不会在这里设置Content-Type头。
// md5:8b7195f02ad8ced0
# <翻译结束>


<原文开始>
// WriteStatusExit writes HTTP `status` and `content` to the response and exits executing
// of current handler if success. The "Exit" feature is commonly used to replace usage of
// return statements in the handler, for convenience.
<原文结束>

# <翻译开始>
// WriteStatusExit 将HTTP状态码`status`和内容`content`写入响应，并在成功时退出当前处理程序的执行。"Exit"特性通常用于方便地替代处理程序中返回语句的使用。
// md5:2e5cbf96316a0c3c
# <翻译结束>

