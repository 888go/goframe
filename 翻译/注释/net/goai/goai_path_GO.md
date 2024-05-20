
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
// Path is specified by OpenAPI/Swagger standard version 3.0.
<原文结束>

# <翻译开始>
// Path遵循OpenAPI/Swagger标准版本3.0。. md5:26b252ebd7fb17bd
# <翻译结束>


<原文开始>
// Paths are specified by OpenAPI/Swagger standard version 3.0.
<原文结束>

# <翻译开始>
// 路径按照OpenAPI/Swagger标准版本3.0进行指定。. md5:77c53887ba9bfc0f
# <翻译结束>


<原文开始>
// Create instance according input/output types.
<原文结束>

# <翻译开始>
// 根据输入/输出类型创建实例。. md5:f07c2f3124391e08
# <翻译结束>


<原文开始>
		// Path and Operation are not the same thing, so it is necessary to copy a Meta for Path from Operation and edit it.
		// And you know, we set the Summary and Description for Operation, not for Path, so we need to remove them.
<原文结束>

# <翻译开始>
// 路径（Path）和操作（Operation）不是同一概念，因此需要从操作中复制一个元信息（Meta）到路径，并进行编辑。
// 你知道的，我们是在操作上设置Summary和Description，而不是在路径上，所以我们需要将它们移除。
// md5:82d486896b1d65b3
# <翻译结束>


<原文开始>
	// path security
	// note: the security schema type only support http and apiKey;not support oauth2 and openIdConnect.
	// multi schema separate with comma, e.g. `security: apiKey1,apiKey2`
<原文结束>

# <翻译开始>
// 路径安全
// 注意：安全模式类型仅支持http和apiKey；不支持oauth2和openIdConnect。
// 多个模式使用逗号分隔，例如：`security: apiKey1,apiKey2`
// md5:b64ffa4261f0711d
# <翻译结束>


<原文开始>
	// =================================================================================================================
	// Request Parameter.
	// =================================================================================================================
<原文结束>

# <翻译开始>
// =================================================================================================================
// 请求参数。
// =================================================================================================================
// md5:c70d5376eecf5c01
# <翻译结束>


<原文开始>
	// =================================================================================================================
	// Request Body.
	// =================================================================================================================
<原文结束>

# <翻译开始>
// =================================================================================================================
// 请求体
// =================================================================================================================
// md5:c70baaeba9963b54
# <翻译结束>


<原文开始>
// Supported mime types of request.
<原文结束>

# <翻译开始>
// 支持的请求MIME类型。. md5:fd32e8079c221b58
# <翻译结束>


<原文开始>
	// =================================================================================================================
	// Response.
	// =================================================================================================================
<原文结束>

# <翻译开始>
// =======================================================
// 响应。
// =======================================================
// md5:ceb9c442cfbdefa1
# <翻译结束>


<原文开始>
// Supported mime types of response.
<原文结束>

# <翻译开始>
// 支持的响应MIME类型。. md5:aefcf019c3abea83
# <翻译结束>


<原文开始>
// If customized response mime type, it then ignores common response feature.
<原文结束>

# <翻译开始>
// 如果指定了自定义的响应MIME类型，则会忽略通用响应特性。. md5:c0c25e2bd38f6d7b
# <翻译结束>


<原文开始>
// Remove operation body duplicated properties.
<原文结束>

# <翻译开始>
// 移除操作体中重复的属性。. md5:976053e0b8002715
# <翻译结束>


<原文开始>
// Assign to certain operation attribute.
<原文结束>

# <翻译开始>
// 为特定操作分配属性。. md5:2e40ddbde8a1317e
# <翻译结束>


<原文开始>
// GET operations cannot have a requestBody.
<原文结束>

# <翻译开始>
// GET 操作不能有请求体。. md5:efd94c634a1773f9
# <翻译结束>


<原文开始>
// DELETE operations cannot have a requestBody.
<原文结束>

# <翻译开始>
// DELETE操作不能有requestBody。. md5:29660405e268d3ca
# <翻译结束>


<原文开始>
// Nothing to do for Connect.
<原文结束>

# <翻译开始>
// 对于Connect，无需执行任何操作。. md5:200e0639d4f11b33
# <翻译结束>


<原文开始>
// Check operation request body have common request data field.
<原文结束>

# <翻译开始>
// 检查操作请求体是否包含通用请求数据字段。. md5:3e4ccc578046cc45
# <翻译结束>


<原文开始>
// Check request body schema
<原文结束>

# <翻译开始>
// 检查请求体架构. md5:dab7ff5a79f31000
# <翻译结束>


<原文开始>
// Check request body schema ref.
<原文结束>

# <翻译开始>
// 检查请求体架构引用。. md5:0b89b0d73335c7ae
# <翻译结束>


<原文开始>
// Check the Value public field for the request body.
<原文结束>

# <翻译开始>
// 检查请求体中的 Value 公共字段。. md5:dd0253ff15259e4b
# <翻译结束>


<原文开始>
// Check request body schema value.
<原文结束>

# <翻译开始>
// 检查请求体的schema值。. md5:549791cc92e82372
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。. md5:43c3b36e60a18f9a
# <翻译结束>


<原文开始>
// To prevent JSON marshal recursion error.
<原文结束>

# <翻译开始>
// 为了防止JSON序列化时的递归错误。. md5:add9f5a47e638cc5
# <翻译结束>

