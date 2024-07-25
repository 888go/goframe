
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
// xmlHeaderBytes is the most common XML format header.
<原文结束>

# <翻译开始>
	// xmlHeaderBytes是最常见的XML格式头。 md5:a1aeea32c5f6c441
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
// Parse 是最常用的函数，它将请求参数转换为结构体或结构体切片。同时，根据结构体上的验证标签，自动对结构体或结构体切片的每个元素进行验证。
//
// 参数 `pointer` 可以是以下类型之一：*struct/*struct/*[]struct/*[]*struct。
//
// 它支持单个和多个结构体的转换：
// 1. 单个结构体时，请求内容格式如：{"id":1, "name":"john"} 或 ?id=1&name=john
// 2. 多个结构体时，请求内容格式如：[{"id":1, "name":"john"}, {"id":, "name":"smith"}]
//
// 待办事项：通过减少跨包对同一变量的重复反射使用，来提升性能。 md5:ad971f0fee54e93d
# <翻译结束>


<原文开始>
// ParseQuery performs like function Parse, but only parses the query parameters.
<原文结束>

# <翻译开始>
// ParseQuery 的行为类似于 Parse 函数，但只解析查询参数。 md5:4104abbe70053960
# <翻译结束>


<原文开始>
// ParseForm performs like function Parse, but only parses the form parameters or the body content.
<原文结束>

# <翻译开始>
// ParseForm 类似于 Parse 函数，但只解析表单参数或主体内容。 md5:c384eb18ba068958
# <翻译结束>


<原文开始>
// doParse parses the request data to struct/structs according to request type.
<原文结束>

# <翻译开始>
// doParse 根据请求类型解析请求数据到结构体/结构体中。 md5:82daab462d052004
# <翻译结束>


<原文开始>
	// Single struct, post content like:
	// 1. {"id":1, "name":"john"}
	// 2. ?id=1&name=john
<原文结束>

# <翻译开始>
	// 单个结构体，帖子内容格式如下：
	// 1. {"id":1, "name":"john"}
	// 2. ?id=1&name=john md5:968f64e28941480c
# <翻译结束>


<原文开始>
		// TODO: https://github.com/gogf/gf/pull/2450
		// Validation.
<原文结束>

# <翻译开始>
		// 待办事项: https:		//github.com/gogf/gf/pull/2450
		// 验证。 md5:ec24b1494dabb977
# <翻译结束>


<原文开始>
	// Multiple struct, it only supports JSON type post content like:
	// [{"id":1, "name":"john"}, {"id":, "name":"smith"}]
<原文结束>

# <翻译开始>
	// 多个结构体，它只支持像这样的JSON类型POST内容：
	// [{"id":1, "name":"john"}, {"id":2, "name":"smith"}] md5:b759870b71d2ffab
# <翻译结束>


<原文开始>
		// If struct slice conversion, it might post JSON/XML/... content,
		// so it uses `gjson` for the conversion.
<原文结束>

# <翻译开始>
		// 如果是结构体切片转换，可能会包含JSON/XML等内容，因此它使用`gjson`进行转换。 md5:e60fd34347047253
# <翻译结束>


<原文开始>
// Get is alias of GetRequest, which is one of the most commonly used functions for
// retrieving parameter.
// See r.GetRequest.
<原文结束>

# <翻译开始>
// Get 是 GetRequest 的别名，它是用于检索参数的最常用函数之一。
// 请参见 r.GetRequest。 md5:80825e01a3c06041
# <翻译结束>


<原文开始>
// GetBody retrieves and returns request body content as bytes.
// It can be called multiple times retrieving the same body content.
<原文结束>

# <翻译开始>
// GetBody 读取并返回请求体内容为字节。
// 可以多次调用，每次都返回相同的正文内容。 md5:be66d2484fd786ca
# <翻译结束>


<原文开始>
// MakeBodyRepeatableRead marks the request body could be repeatedly readable or not.
// It also returns the current content of the request body.
<原文结束>

# <翻译开始>
// MakeBodyRepeatableRead 标记请求体是否可以重复读取。它还会返回当前请求体的内容。 md5:3cda0a2da5c712d7
# <翻译结束>


<原文开始>
// GetBodyString retrieves and returns request body content as string.
// It can be called multiple times retrieving the same body content.
<原文结束>

# <翻译开始>
// GetBodyString 用于检索并返回请求体内容作为字符串。可以多次调用以获取相同的内容。 md5:503c28317dc909ca
# <翻译结束>


<原文开始>
// GetJson parses current request content as JSON format, and returns the JSON object.
// Note that the request content is read from request BODY, not from any field of FORM.
<原文结束>

# <翻译开始>
// GetJson 将当前请求内容解析为JSON格式，并返回JSON对象。
// 注意，请求内容是从请求体(BODY)中读取的，而不是从表单的任何字段中读取。 md5:166af4b89b6a5a68
# <翻译结束>


<原文开始>
// GetMap is an alias and convenient function for GetRequestMap.
// See GetRequestMap.
<原文结束>

# <翻译开始>
// GetMap 是 GetRequestMap 函数的别名，提供便利的使用方式。
// 参考 GetRequestMap。 md5:395e8bbf3fea416a
# <翻译结束>


<原文开始>
// GetMapStrStr is an alias and convenient function for GetRequestMapStrStr.
// See GetRequestMapStrStr.
<原文结束>

# <翻译开始>
// GetMapStrStr是GetRequestMapStrStr的别名，提供便捷的功能。详情请参阅GetRequestMapStrStr。 md5:1828f3886ccd906d
# <翻译结束>


<原文开始>
// GetStruct is an alias and convenient function for GetRequestStruct.
// See GetRequestStruct.
<原文结束>

# <翻译开始>
// GetStruct 是 GetRequestStruct 的别名和便捷函数。详情请参阅 GetRequestStruct。 md5:c558debb875b77cd
# <翻译结束>


<原文开始>
// parseQuery parses query string into r.queryMap.
<原文结束>

# <翻译开始>
// parseQuery 将查询字符串解析到 r.queryMap 中。 md5:9a26b305dc518866
# <翻译结束>


<原文开始>
// parseBody parses the request raw data into r.rawMap.
// Note that it also supports JSON data from client request.
<原文结束>

# <翻译开始>
// parseBody 将请求的原始数据解析到 r.rawMap 中。
// 请注意，它还支持从客户端请求的 JSON 数据。 md5:f8f001deccef59e6
# <翻译结束>


<原文开始>
// There's no data posted.
<原文结束>

# <翻译开始>
	// 没有提交任何数据。 md5:cf70840053024c2b
# <翻译结束>


<原文开始>
// Trim space/new line characters.
<原文结束>

# <翻译开始>
		// 去除空格和换行符。 md5:0cf77adc8fee1e9a
# <翻译结束>


<原文开始>
// Default parameters decoding.
<原文结束>

# <翻译开始>
		// 默认参数解码。 md5:941d9de3ebb46554
# <翻译结束>


<原文开始>
// parseForm parses the request form for HTTP method PUT, POST, PATCH.
// The form data is pared into r.formMap.
//
// Note that if the form was parsed firstly, the request body would be cleared and empty.
<原文结束>

# <翻译开始>
// parseForm 解析HTTP方法PUT，POST，PATCH的请求表单。
// 表单数据被解析到r.formMap中。
//
// 请注意，如果已经先解析了表单，那么请求体将会被清空。 md5:97f04aa06758375b
# <翻译结束>


<原文开始>
			// To avoid big memory consuming.
			// The `multipart/` type form always contains binary data, which is not necessary read twice.
<原文结束>

# <翻译开始>
			// 为了避免大量消耗内存。
			// `multipart/` 类型的表单始终包含二进制数据，没有必要读取两次。 md5:d95befcac4fa7fd0
# <翻译结束>


<原文开始>
// multipart/form-data, multipart/mixed
<原文结束>

# <翻译开始>
			// 这两个注释是在描述MIME类型。`multipart/form-data`通常用于通过HTTP发送表单数据，如文件上传。`multipart/mixed`则用于包含多个部分的混合内容，每个部分可以是不同的MIME类型，常用于邮件或API请求中包含多种类型的附件或数据。 md5:5f5a1e86722f47ec
# <翻译结束>


<原文开始>
// application/x-www-form-urlencoded
<原文结束>

# <翻译开始>
			// 应用程序/x-www-form-urlencoded. md5:6de553b2a7019beb
# <翻译结束>


<原文开始>
// Parse the form data using united parsing way.
<原文结束>

# <翻译开始>
			// 使用统一的解析方式解析表单数据。 md5:21f3f94370e18b5d
# <翻译结束>


<原文开始>
				// Invalid parameter name.
				// Only allow chars of: '\w', '[', ']', '-'.
<原文结束>

# <翻译开始>
				// 非法的参数名称。
				// 只允许使用以下字符：`\w`，`[`，`]`，`-`。 md5:72a7ff7f2d38a973
# <翻译结束>


<原文开始>
// It might be JSON/XML content.
<原文结束>

# <翻译开始>
					// 它可能是JSON或XML内容。 md5:105b844bbc2857c0
# <翻译结束>


<原文开始>
// It parses the request body without checking the Content-Type.
<原文结束>

# <翻译开始>
	// 它解析请求体，而不检查Content-Type。 md5:89cfec67836d4575
# <翻译结束>


<原文开始>
// GetMultipartForm parses and returns the form as multipart forms.
<原文结束>

# <翻译开始>
// GetMultipartForm 解析并返回表单为多部分形式。 md5:c80c641ed3887bea
# <翻译结束>


<原文开始>
// GetMultipartFiles parses and returns the post files array.
// Note that the request form should be type of multipart.
<原文结束>

# <翻译开始>
// GetMultipartFiles 解析并返回表单中的文件数组。
// 请注意，请求表单的类型应该是multipart。 md5:33503fc76a60c149
# <翻译结束>


<原文开始>
// Support "name[]" as array parameter.
<原文结束>

# <翻译开始>
	// 支持" name[]"作为数组参数。 md5:f1460d96fee37609
# <翻译结束>


<原文开始>
// Support "name[0]","name[1]","name[2]", etc. as array parameter.
<原文结束>

# <翻译开始>
	// 支持将"name[0]","name[1]","name[2]"等作为数组参数使用。 md5:a9545b3b88169505
# <翻译结束>

