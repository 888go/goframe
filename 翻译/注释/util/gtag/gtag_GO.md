
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
// Package gtag providing tag content storing for struct.
//
// Note that calling functions of this package is not concurrently safe,
// which means you cannot call them in runtime but in boot procedure.
<原文结束>

# <翻译开始>
// Package gtag 提供了为结构体存储标签内容的功能。
//
// 注意：该包中的函数调用不保证并发安全，
// 这意味着你不能在运行时调用它们，而只能在启动过程中调用。
// 这段注释翻译成中文后为：
// ```go
// 包gtag 提供了用于为结构体提供标签内容存储的服务。
//
// 需要注意的是，本包中提供的函数并非线程安全的，
// 即这些函数不能在程序运行时并发调用，只能在初始化或启动阶段调用。
# <翻译结束>


<原文开始>
// Default value tag of struct field for receiving parameters from HTTP request.
<原文结束>

# <翻译开始>
// 结构体字段的默认值标签，用于从HTTP请求接收参数。
# <翻译结束>


<原文开始>
// Parameter name for converting certain parameter to specified struct field.
<原文结束>

# <翻译开始>
// 将特定参数转换为指定结构体字段的参数名称。
# <翻译结束>


<原文开始>
// Validation rule tag for struct of field.
<原文结束>

# <翻译开始>
// 结构体字段的验证规则标签
# <翻译结束>


<原文开始>
// No validation for specified struct/field.
<原文结束>

# <翻译开始>
// 对指定的结构体/字段不进行验证。
# <翻译结束>


<原文开始>
// ORM tag for ORM feature, which performs different features according scenarios.
<原文结束>

# <翻译开始>
// ORM标签用于ORM特性，根据不同的场景执行不同的功能。
# <翻译结束>


<原文开始>
// Arg tag for struct, usually for command argument option.
<原文结束>

# <翻译开始>
// Arg标签用于结构体，通常用于命令行参数选项。
# <翻译结束>


<原文开始>
// Brief tag for struct, usually be considered as summary.
<原文结束>

# <翻译开始>
// 结构体的简短标签，通常被视为概述或摘要。
# <翻译结束>


<原文开始>
// Root tag for struct, usually for nested commands management.
<原文结束>

# <翻译开始>
// Root标签用于结构体，通常用于嵌套命令的管理。
# <翻译结束>


<原文开始>
// Additional tag for struct, usually for additional description of command.
<原文结束>

# <翻译开始>
// 结构体的附加标签，通常用于命令的额外描述。
# <翻译结束>


<原文开始>
// Route path for HTTP request.
<原文结束>

# <翻译开始>
// HTTP请求的路由路径。
# <翻译结束>


<原文开始>
// Route method for HTTP request.
<原文结束>

# <翻译开始>
// Route 方法用于处理 HTTP 请求。
# <翻译结束>


<原文开始>
// Route domain for HTTP request.
<原文结束>

# <翻译开始>
// 为HTTP请求路由域名。
# <翻译结束>


<原文开始>
// MIME type for HTTP request/response.
<原文结束>

# <翻译开始>
// HTTP请求/响应的MIME类型。
# <翻译结束>


<原文开始>
// Summary for struct, usually for OpenAPI in request struct.
<原文结束>

# <翻译开始>
// Summary 用于结构体，通常用于OpenAPI中的请求结构体。
# <翻译结束>


<原文开始>
// Description for struct, usually for OpenAPI in request struct.
<原文结束>

# <翻译开始>
// 结构体描述，通常用于OpenAPI中的请求结构体。
# <翻译结束>


<原文开始>
// Example for struct, usually for OpenAPI in request struct.
<原文结束>

# <翻译开始>
// 示例：用于结构体，通常用于OpenAPI中的请求结构体。
# <翻译结束>


<原文开始>
// Examples for struct, usually for OpenAPI in request struct.
<原文结束>

# <翻译开始>
// 示例代码展示结构体的用法，通常用于OpenAPI中的请求结构体。
# <翻译结束>


<原文开始>
// External docs for struct, always for OpenAPI in request struct.
<原文结束>

# <翻译开始>
// 结构体外部文档，始终用于OpenAPI中的请求结构体。
# <翻译结束>


<原文开始>
// GConv defines the converting target name for specified struct field.
<原文结束>

# <翻译开始>
// GConv 定义了为特定结构体字段指定的转换目标名称。
# <翻译结束>


<原文开始>
// Json tag is supported by stdlib.
<原文结束>

# <翻译开始>
// Json标签受到stdlib的支持。
# <翻译结束>


<原文开始>
// Security defines scheme for authentication. Detail to see https://swagger.io/docs/specification/authentication/
<原文结束>

# <翻译开始>
// Security 定义了身份验证方案。详情请参阅 https://swagger.io/docs/specification/authentication/
# <翻译结束>


<原文开始>
// Swagger distinguishes between the following parameter types based on the parameter location. Detail to see https://swagger.io/docs/specification/describing-parameters/
<原文结束>

# <翻译开始>
// Swagger 根据参数位置区分以下几种参数类型。详情请参阅 https://swagger.io/docs/specification/describing-parameters/
// 在Swagger中，根据参数所在的位置，对参数进行如下类型的区别定义
# <翻译结束>


<原文开始>
// Short name of Default.
<原文结束>

# <翻译开始>
// Default的简称。
# <翻译结束>


<原文开始>
// Short name of Param.
<原文结束>

# <翻译开始>
// Param的简称。
# <翻译结束>


<原文开始>
// Short name of Valid.
<原文结束>

# <翻译开始>
// Valid的简称
# <翻译结束>


<原文开始>
// Short name of Additional.
<原文结束>

# <翻译开始>
// Additional的简称。
# <翻译结束>


<原文开始>
// MIME type for HTTP request.
<原文结束>

# <翻译开始>
// HTTP请求的MIME类型。
# <翻译结束>


<原文开始>
// Short name of Summary.
<原文结束>

# <翻译开始>
// Summary的简称。
# <翻译结束>


<原文开始>
// Short name of Description.
<原文结束>

# <翻译开始>
// 描述的简称。
# <翻译结束>


<原文开始>
// Short name of Example.
<原文结束>

# <翻译开始>
// Example 的简称。
# <翻译结束>


<原文开始>
// Short name of Examples.
<原文结束>

# <翻译开始>
// 示例的简称。
# <翻译结束>


<原文开始>
// Short name of ExternalDocs.
<原文结束>

# <翻译开始>
// ExternalDocs 的简称。
# <翻译结束>

