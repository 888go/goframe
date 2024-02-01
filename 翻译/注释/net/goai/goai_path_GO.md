
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
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Path is specified by OpenAPI/Swagger standard version 3.0.
<原文结束>

# <翻译开始>
// Path 是由 OpenAPI/Swagger 标准版本 3.0 规定的。
# <翻译结束>


<原文开始>
// Paths are specified by OpenAPI/Swagger standard version 3.0.
<原文结束>

# <翻译开始>
// 路径由OpenAPI/Swagger标准版本3.0指定。
# <翻译结束>


<原文开始>
// Create instance according input/output types.
<原文结束>

# <翻译开始>
// 根据输入/输出类型创建实例。
# <翻译结束>


<原文开始>
		// Path and Operation are not the same thing, so it is necessary to copy a Meta for Path from Operation and edit it.
		// And you know, we set the Summary and Description for Operation, not for Path, so we need to remove them.
<原文结束>

# <翻译开始>
// Path（路径）和 Operation（操作）不是同一概念，因此有必要从 Operation 中为 Path 复制一份 Meta 信息并进行编辑。
// 另外需要注意的是，我们在 Operation 上设置 Summary（摘要）和 Description（描述），而不是在 Path 上设置，所以我们需要将它们移除。
# <翻译结束>


<原文开始>
	// path security
	// note: the security schema type only support http and apiKey;not support oauth2 and openIdConnect.
	// multi schema separate with comma, e.g. `security: apiKey1,apiKey2`
<原文结束>

# <翻译开始>
// 路径安全
// 注：安全模式类型仅支持http和apiKey，不支持oauth2和openIdConnect。
// 多种模式之间用逗号分隔，例如 `security: apiKey1,apiKey2`
# <翻译结束>


<原文开始>
	// =================================================================================================================
	// Request Parameter.
	// =================================================================================================================
<原文结束>

# <翻译开始>
// =================================================================================================================
// 请求参数.
// =================================================================================================================
# <翻译结束>


<原文开始>
	// =================================================================================================================
	// Request Body.
	// =================================================================================================================
<原文结束>

# <翻译开始>
// =================================================================================================================
// 请求体.
// =================================================================================================================
# <翻译结束>


<原文开始>
// Supported mime types of request.
<原文结束>

# <翻译开始>
// 请求支持的MIME类型。
# <翻译结束>


<原文开始>
	// =================================================================================================================
	// Response.
	// =================================================================================================================
<原文结束>

# <翻译开始>
// =================================================================================================================
// 响应.
// =================================================================================================================
# <翻译结束>


<原文开始>
// Supported mime types of response.
<原文结束>

# <翻译开始>
// 响应支持的MIME类型。
# <翻译结束>


<原文开始>
// If customized response mime type, it then ignores common response feature.
<原文结束>

# <翻译开始>
// 如果指定了自定义的响应MIME类型，则会忽略通用的响应特性。
# <翻译结束>


<原文开始>
// Remove operation body duplicated properties.
<原文结束>

# <翻译开始>
// 移除操作体中重复的属性。
# <翻译结束>


<原文开始>
// Assign to certain operation attribute.
<原文结束>

# <翻译开始>
// 给特定操作属性赋值。
# <翻译结束>


<原文开始>
// GET operations cannot have a requestBody.
<原文结束>

# <翻译开始>
// GET方法不能包含请求体。
# <翻译结束>


<原文开始>
// DELETE operations cannot have a requestBody.
<原文结束>

# <翻译开始>
// DELETE操作不能包含请求体。
# <翻译结束>


<原文开始>
// Check operation request body have common request data field.
<原文结束>

# <翻译开始>
// 检查操作请求体中是否包含通用请求数据字段。
# <翻译结束>


<原文开始>
// Check request body schema ref.
<原文结束>

# <翻译开始>
// 检查请求体schema引用。
# <翻译结束>


<原文开始>
// Check the Value public field for the request body.
<原文结束>

# <翻译开始>
// 检查请求体中Value公共字段的值。
# <翻译结束>


<原文开始>
// Check request body schema value.
<原文结束>

# <翻译开始>
// 检查请求体中 schema 值。
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
# <翻译结束>


<原文开始>
// To prevent JSON marshal recursion error.
<原文结束>

# <翻译开始>
// 为防止JSON序列化时出现递归错误
# <翻译结束>


<原文开始>
// Precise route path.
<原文结束>

# <翻译开始>
// 精确路由路径。
# <翻译结束>


<原文开始>
// Route path prefix.
<原文结束>

# <翻译开始>
// 路由路径前缀。
# <翻译结束>


<原文开始>
// Uniformed function.
<原文结束>

# <翻译开始>
// 统一化函数
# <翻译结束>


<原文开始>
// Allowed request mime.
<原文结束>

# <翻译开始>
// 允许的请求MIME类型。
# <翻译结束>


<原文开始>
// Nothing to do for Connect.
<原文结束>

# <翻译开始>
// 对于Connect无需执行任何操作。
# <翻译结束>

