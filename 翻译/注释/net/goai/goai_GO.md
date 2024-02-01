
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
// Package goai implements and provides document generating for OpenApi specification.
//
// https://editor.swagger.io/
<原文结束>

# <翻译开始>
// Package goai 实现并提供了针对 OpenApi 规范的文档生成功能。
//
// 参考链接：https://editor.swagger.io/
# <翻译结束>


<原文开始>
// OpenApiV3 is the structure defined from:
// https://swagger.io/specification/
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md
<原文结束>

# <翻译开始>
// OpenApiV3 是从以下网址定义的结构体：
// https://swagger.io/specification/
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md
// （译：OpenApiV3 结构体是根据以下链接中定义的 OpenAPI 3.0 规范实现的：）
// （https://swagger.io/specification/，以及 OpenAPI 3.0 的具体版本规范：）
// （https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md）
# <翻译结束>


<原文开始>
// New creates and returns an OpenApiV3 implements object.
<原文结束>

# <翻译开始>
// New 创建并返回一个实现了 OpenApiV3 的对象。
# <翻译结束>


<原文开始>
// AddInput is the structured parameter for function OpenApiV3.Add.
<原文结束>

# <翻译开始>
// AddInput 是函数 OpenApiV3.Add 的结构化参数。
# <翻译结束>


<原文开始>
// Path specifies the custom path if this is not configured in Meta of struct tag.
<原文结束>

# <翻译开始>
// Path 指定自定义路径，如果在结构体标签的 Meta 中未配置此路径，则使用该指定路径。
# <翻译结束>


<原文开始>
// Prefix specifies the custom route path prefix, which will be added with the path tag in Meta of struct tag.
<原文结束>

# <翻译开始>
// Prefix 指定自定义路由路径前缀，它将与结构体标签中 Meta 的 path 标签相结合。
# <翻译结束>


<原文开始>
// Method specifies the custom HTTP method if this is not configured in Meta of struct tag.
<原文结束>

# <翻译开始>
// Method 指定自定义的 HTTP 方法，如果在结构体标签的 Meta 中未配置该方法时使用。
# <翻译结束>


<原文开始>
// Object can be an instance of struct or a route function.
<原文结束>

# <翻译开始>
// Object 可以是结构体实例或路由函数。
# <翻译结束>


<原文开始>
// Add adds an instance of struct or a route function to OpenApiV3 definition implements.
<原文结束>

# <翻译开始>
// Add 将一个结构体实例或路由函数添加到OpenApiV3定义实现中。
# <翻译结束>


<原文开始>
// golangTypeToOAIFormat converts and returns OpenAPI parameter format for given golang type `t`.
// Note that it does not return standard OpenAPI parameter format but custom format in golang type.
<原文结束>

# <翻译开始>
// golangTypeToOAIFormat 将给定的Golang类型`t`转换并返回OpenAPI参数格式。
// 注意，它不会返回标准的OpenAPI参数格式，而是返回Golang类型的自定义格式。
# <翻译结束>


<原文开始>
// Pointer type has no PkgPath.
<原文结束>

# <翻译开始>
// 指针类型没有 PkgPath。
# <翻译结束>

