
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
// Write writes `content` to the response buffer.
<原文结束>

# <翻译开始>
// Write将`content`写入响应缓冲区。
# <翻译结束>


<原文开始>
// WriteExit writes `content` to the response buffer and exits executing of current handler.
// The "Exit" feature is commonly used to replace usage of return statements in the handler,
// for convenience.
<原文结束>

# <翻译开始>
// WriteExit 将`content`写入响应缓冲区并退出当前处理器的执行。
// “Exit”特性通常用于为了方便起见，在处理器中替换 return 语句的使用。
# <翻译结束>


<原文开始>
// WriteOver overwrites the response buffer with `content`.
<原文结束>

# <翻译开始>
// WriteOver 将`content`覆盖写入响应缓冲区。
# <翻译结束>


<原文开始>
// WriteOverExit overwrites the response buffer with `content` and exits executing
// of current handler. The "Exit" feature is commonly used to replace usage of return
// statements in the handler, for convenience.
<原文结束>

# <翻译开始>
// WriteOverExit 用 `content` 覆盖响应缓冲区并退出当前处理器的执行。  
// "Exit" 特性通常用于为了方便起见，在处理器中替换 return 语句的使用。
# <翻译结束>


<原文开始>
// Writef writes the response with fmt.Sprintf.
<原文结束>

# <翻译开始>
// Writef 使用 fmt.Sprintf 方法写入响应内容。
# <翻译结束>


<原文开始>
// WritefExit writes the response with fmt.Sprintf and exits executing of current handler.
// The "Exit" feature is commonly used to replace usage of return statements in the handler,
// for convenience.
<原文结束>

# <翻译开始>
// WritefExit 通过 fmt.Sprintf 写入响应，并退出当前处理器的执行。
// "Exit" 特性通常用于为了方便，在处理器中替代 return 语句的使用。
# <翻译结束>


<原文开始>
// Writeln writes the response with `content` and new line.
<原文结束>

# <翻译开始>
// Writeln将`content`内容和换行符一起写入响应。
# <翻译结束>


<原文开始>
// WritelnExit writes the response with `content` and new line and exits executing
// of current handler. The "Exit" feature is commonly used to replace usage of return
// statements in the handler, for convenience.
<原文结束>

# <翻译开始>
// WritelnExit 将`content`内容及换行符写入响应，并终止当前处理器的执行。
// "Exit"特性通常用于为了方便起见，替代处理器中return语句的使用。
# <翻译结束>


<原文开始>
// Writefln writes the response with fmt.Sprintf and new line.
<原文结束>

# <翻译开始>
// Writefln 使用 fmt.Sprintf 格式化输出并将内容与换行符写入响应。
# <翻译结束>


<原文开始>
// WriteflnExit writes the response with fmt.Sprintf and new line and exits executing
// of current handler. The "Exit" feature is commonly used to replace usage of return
// statement in the handler, for convenience.
<原文结束>

# <翻译开始>
// WriteflnExit 通过 fmt.Sprintf 和换行符写出响应，并退出当前处理器的执行。
// "Exit" 特性通常用于为了方便，而替换处理器中 return 语句的使用。
# <翻译结束>


<原文开始>
// WriteJson writes `content` to the response with JSON format.
<原文结束>

# <翻译开始>
// WriteJson 将`content`以JSON格式写入响应中。
# <翻译结束>


<原文开始>
// If given string/[]byte, response it directly to the client.
<原文结束>

# <翻译开始>
// 如果给定的是字符串或[]byte，直接将响应发送回客户端。
# <翻译结束>


<原文开始>
// Else use json.Marshal function to encode the parameter.
<原文结束>

# <翻译开始>
// 否则使用json.Marshal函数对参数进行编码。
# <翻译结束>


<原文开始>
// WriteJsonExit writes `content` to the response with JSON format and exits executing
// of current handler if success. The "Exit" feature is commonly used to replace usage of
// return statements in the handler, for convenience.
<原文结束>

# <翻译开始>
// WriteJsonExit将`content`以JSON格式写入响应，并在成功时退出当前处理器的执行。
// “Exit”特性通常用于为了方便起见，替换处理器中return语句的使用。
# <翻译结束>


<原文开始>
// WriteJsonP writes `content` to the response with JSONP format.
//
// Note that there should be a "callback" parameter in the request for JSONP format.
<原文结束>

# <翻译开始>
// WriteJsonP 将`content`以JSONP格式写入响应中。
//
// 注意：对于JSONP格式，请求中应包含一个"callback"参数。
# <翻译结束>


<原文开始>
// If given string/[]byte, response it directly to client.
<原文结束>

# <翻译开始>
// 如果给定的是字符串或[]byte，直接将响应发送给客户端。
# <翻译结束>


<原文开始>
// r.Header().Set("Content-Type", "application/json")
<原文结束>

# <翻译开始>
// 设置HTTP响应头的"Content-Type"字段为"application/json"
# <翻译结束>


<原文开始>
// WriteJsonPExit writes `content` to the response with JSONP format and exits executing
// of current handler if success. The "Exit" feature is commonly used to replace usage of
// return statements in the handler, for convenience.
//
// Note that there should be a "callback" parameter in the request for JSONP format.
<原文结束>

# <翻译开始>
// WriteJsonPExit 将`content`以JSONP格式写入响应，并在成功时退出当前处理器的执行。
// “Exit”特性通常用于替换处理器中return语句的使用，以便于简化代码。
//
// 注意，请求中应包含一个“callback”参数以适应JSONP格式。
# <翻译结束>


<原文开始>
// WriteXml writes `content` to the response with XML format.
<原文结束>

# <翻译开始>
// WriteXml 将`content`以XML格式写入响应。
# <翻译结束>


<原文开始>
// If given string/[]byte, response it directly to clients.
<原文结束>

# <翻译开始>
// 如果给定的是字符串或[]byte，直接将其响应给客户端。
# <翻译结束>


<原文开始>
// WriteXmlExit writes `content` to the response with XML format and exits executing
// of current handler if success. The "Exit" feature is commonly used to replace usage
// of return statements in the handler, for convenience.
<原文结束>

# <翻译开始>
// WriteXmlExit将`content`以XML格式写入响应，并在成功时退出当前处理器的执行。
// "Exit"特性通常用于为了方便起见，在处理器中替换return语句的使用。
# <翻译结束>


<原文开始>
// WriteStatus writes HTTP `status` and `content` to the response.
// Note that it does not set a Content-Type header here.
<原文结束>

# <翻译开始>
// WriteStatus将HTTP状态码`status`和内容`content`写入响应中。
// 注意，这里没有设置Content-Type头信息。
# <翻译结束>


<原文开始>
// WriteStatusExit writes HTTP `status` and `content` to the response and exits executing
// of current handler if success. The "Exit" feature is commonly used to replace usage of
// return statements in the handler, for convenience.
<原文结束>

# <翻译开始>
// WriteStatusExit 将HTTP状态码`status`和内容`content`写入响应，并在成功时退出当前处理器的执行。
// "Exit"特性通常用于为了方便，替换处理器中return语句的使用。
# <翻译结束>

