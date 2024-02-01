
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
// xmlHeaderBytes is the most common XML format header.
<原文结束>

# <翻译开始>
// xmlHeaderBytes 是最常见的 XML 格式头部。
# <翻译结束>


<原文开始>
// Parse is the most commonly used function, which converts request parameters to struct or struct
// slice. It also automatically validates the struct or every element of the struct slice according
// to the validation tag of the struct.
//
// The parameter `pointer` can be type of: *struct/**struct/*[]struct/*[]*struct.
//
// It supports single and multiple struct converting:
// 1. Single struct, post content like: {"id":1, "name":"john"} or ?id=1&name=john
// 2. Multiple struct, post content like: [{"id":1, "name":"john"}, {"id":, "name":"smith"}]
//
// TODO: Improve the performance by reducing duplicated reflect usage on the same variable across packages.
<原文结束>

# <翻译开始>
// Parse 是最常用的函数，用于将请求参数转换为结构体或结构体切片。
// 同时，它也会根据结构体上的验证标签自动校验结构体或结构体切片中的每个元素。
//
// 参数 `pointer` 可以是以下类型：*struct/**struct/*[]struct/*[]*struct。
//
// 它支持单个和多个结构体的转换：
// 1. 单个结构体，POST 内容如：{"id":1, "name":"john"} 或 ?id=1&name=john
// 2. 多个结构体，POST 内容如：[{"id":1, "name":"john"}, {"id":, "name":"smith"}]
//
// TODO: 通过减少在不同包中对同一变量重复使用 reflect 来提高性能。
# <翻译结束>


<原文开始>
// ParseQuery performs like function Parse, but only parses the query parameters.
<原文结束>

# <翻译开始>
// ParseQuery 类似于函数 Parse，但它只解析查询参数。
# <翻译结束>


<原文开始>
// ParseForm performs like function Parse, but only parses the form parameters or the body content.
<原文结束>

# <翻译开始>
// ParseForm执行类似于函数Parse的功能，但只解析表单参数或主体内容。
# <翻译结束>


<原文开始>
// doParse parses the request data to struct/structs according to request type.
<原文结束>

# <翻译开始>
// doParse 根据请求类型将请求数据解析到结构体/结构体数组中。
# <翻译结束>


<原文开始>
	// Single struct, post content like:
	// 1. {"id":1, "name":"john"}
	// 2. ?id=1&name=john
<原文结束>

# <翻译开始>
// Single 结构体，用于处理如下的POST内容：
// 1. {"id":1, "name":"john"} 
// 2. ?id=1&name=john
# <翻译结束>


<原文开始>
		// TODO: https://github.com/gogf/gf/pull/2450
		// Validation.
<原文结束>

# <翻译开始>
// TODO: 参考 https://github.com/gogf/gf/pull/2450
// 验证功能。
# <翻译结束>


<原文开始>
	// Multiple struct, it only supports JSON type post content like:
	// [{"id":1, "name":"john"}, {"id":, "name":"smith"}]
<原文结束>

# <翻译开始>
// 多条数据结构，它仅支持类似以下JSON格式的POST内容：
// [{"id":1, "name":"john"}, {"id":, "name":"smith"}]
# <翻译结束>


<原文开始>
		// If struct slice conversion, it might post JSON/XML/... content,
		// so it uses `gjson` for the conversion.
<原文结束>

# <翻译开始>
// 如果是结构体切片转换，它可能会发布 JSON/XML/... 格式的内容，
// 因此这里使用 `gjson` 进行转换。
# <翻译结束>


<原文开始>
// Get is alias of GetRequest, which is one of the most commonly used functions for
// retrieving parameter.
// See r.GetRequest.
<原文结束>

# <翻译开始>
// Get 是 GetRequest 的别名，它是用于获取参数的最常用函数之一。
// 请参阅 r.GetRequest。
# <翻译结束>


<原文开始>
// GetBody retrieves and returns request body content as bytes.
// It can be called multiple times retrieving the same body content.
<原文结束>

# <翻译开始>
// GetBody 用于检索并返回请求正文内容作为字节。
// 它可以被多次调用，获取相同正文内容。
# <翻译结束>


<原文开始>
// MakeBodyRepeatableRead marks the request body could be repeatedly readable or not.
// It also returns the current content of the request body.
<原文结束>

# <翻译开始>
// MakeBodyRepeatableRead 标记请求体是否可以被重复读取。
// 同时，它还会返回当前请求体的内容。
# <翻译结束>


<原文开始>
// GetBodyString retrieves and returns request body content as string.
// It can be called multiple times retrieving the same body content.
<原文结束>

# <翻译开始>
// GetBodyString 用于获取并返回请求体内容作为字符串。
// 它可以被多次调用，以获取相同的请求体内容。
# <翻译结束>


<原文开始>
// GetJson parses current request content as JSON format, and returns the JSON object.
// Note that the request content is read from request BODY, not from any field of FORM.
<原文结束>

# <翻译开始>
// GetJson 将当前请求内容解析为JSON格式，并返回JSON对象。
// 注意：请求内容是从request BODY中读取，而不是从FORM的任何字段中读取。
# <翻译结束>


<原文开始>
// GetMap is an alias and convenient function for GetRequestMap.
// See GetRequestMap.
<原文结束>

# <翻译开始>
// GetMap 是一个别名，也是一个方便获取请求映射的函数。
// 请参阅 GetRequestMap。
# <翻译结束>


<原文开始>
// GetMapStrStr is an alias and convenient function for GetRequestMapStrStr.
// See GetRequestMapStrStr.
<原文结束>

# <翻译开始>
// GetMapStrStr 是 GetRequestMapStrStr 的别名和便捷函数。
// 请参阅 GetRequestMapStrStr。
# <翻译结束>


<原文开始>
// GetStruct is an alias and convenient function for GetRequestStruct.
// See GetRequestStruct.
<原文结束>

# <翻译开始>
// GetStruct 是 GetRequestStruct 的别名和便捷函数。
// 请参阅 GetRequestStruct。
# <翻译结束>


<原文开始>
// parseQuery parses query string into r.queryMap.
<原文结束>

# <翻译开始>
// parseQuery 将查询字符串解析到 r.queryMap 中。
# <翻译结束>


<原文开始>
// parseBody parses the request raw data into r.rawMap.
// Note that it also supports JSON data from client request.
<原文结束>

# <翻译开始>
// parseBody 将请求原始数据解析到 r.rawMap 中。
// 注意，它还支持从客户端请求的 JSON 数据。
# <翻译结束>







<原文开始>
// Trim space/new line characters.
<原文结束>

# <翻译开始>
// 去除空格/换行符。
# <翻译结束>












<原文开始>
// Default parameters decoding.
<原文结束>

# <翻译开始>
// 默认参数解码
# <翻译结束>


<原文开始>
// parseForm parses the request form for HTTP method PUT, POST, PATCH.
// The form data is pared into r.formMap.
//
// Note that if the form was parsed firstly, the request body would be cleared and empty.
<原文结束>

# <翻译开始>
// parseForm 用于解析HTTP方法PUT、POST、PATCH的请求表单。
// 表单数据将被解析并存储到r.formMap中。
//
// 注意：如果表单首先被解析，请求体将会被清空，变成空内容。
# <翻译结束>


<原文开始>
// multipart/form-data, multipart/mixed
<原文结束>

# <翻译开始>
// multipart/form-data：这是一种HTTP内容类型，用于编码同一条请求中包含多种不同类型数据（如文本、文件等）的表单数据。在上传文件时尤为常见。
// multipart/mixed：这也是一种HTTP内容类型，用于表示消息体包含多个独立的部分，各个部分可以是不同的数据类型，且每个部分都有自己的Content-Type和边界标识符。它通常用于混合多部分消息，比如在一个HTTP请求中同时发送文本信息和附件。
# <翻译结束>







<原文开始>
// Parse the form data using united parsing way.
<原文结束>

# <翻译开始>
// 使用统一解析方式解析表单数据。
# <翻译结束>


<原文开始>
				// Invalid parameter name.
				// Only allow chars of: '\w', '[', ']', '-'.
<原文结束>

# <翻译开始>
// 无效的参数名称。
// 只允许包含字符：'\w'（代表字母、数字或下划线）、'['、']' 和 '-'。
# <翻译结束>


<原文开始>
// It might be JSON/XML content.
<原文结束>

# <翻译开始>
// 这可能是一个JSON/XML内容。
# <翻译结束>


<原文开始>
// It parses the request body without checking the Content-Type.
<原文结束>

# <翻译开始>
// 它解析请求体，但不检查 Content-Type。
# <翻译结束>


<原文开始>
// GetMultipartForm parses and returns the form as multipart forms.
<原文结束>

# <翻译开始>
// GetMultipartForm 解析并返回表单为多部分表单形式。
# <翻译结束>


<原文开始>
// GetMultipartFiles parses and returns the post files array.
// Note that the request form should be type of multipart.
<原文结束>

# <翻译开始>
// GetMultipartFiles 解析并返回 POST 请求中的文件数组。
// 注意，请求表单的类型应当为 multipart。
# <翻译结束>


<原文开始>
// Support "name[]" as array parameter.
<原文结束>

# <翻译开始>
// 支持 "name[]" 作为数组参数。
# <翻译结束>


<原文开始>
// Support "name[0]","name[1]","name[2]", etc. as array parameter.
<原文结束>

# <翻译开始>
// 支持 "name[0]"、"name[1]"、"name[2]" 等形式的数组参数。
# <翻译结束>


<原文开始>
// application/x-www-form-urlencoded
<原文结束>

# <翻译开始>
// application/x-www-form-urlencoded 是一种常见的HTTP内容类型，用于表示URL编码的表单数据。在Go语言中，通常在网络请求或表单提交时使用这种格式对键值对进行编码。
// ```go
// 这是HTTP请求Header中Content-Type的一种常见取值
// 表示请求体中的数据采用了"application/x-www-form-urlencoded"编码格式
// 该格式将表单字段名和值连接成键值对，并且对特殊字符进行URL编码
# <翻译结束>


<原文开始>
// There's no data posted.
<原文结束>

# <翻译开始>
// 没有提交任何数据。
# <翻译结束>


<原文开始>
// JSON format checks.
<原文结束>

# <翻译开始>
// JSON格式检查。
# <翻译结束>


<原文开始>
// XML format checks.
<原文结束>

# <翻译开始>
// XML格式检查。
# <翻译结束>

