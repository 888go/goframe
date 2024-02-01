
<原文开始>
// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame gf 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
# <翻译结束>


<原文开始>
// Package gcode provides universal error code definition and common error codes implements.
<原文结束>

# <翻译开始>
// Package gcode 提供通用错误码定义及常见错误码实现。
# <翻译结束>


<原文开始>
// Code is universal error code interface definition.
<原文结束>

# <翻译开始>
// Code 是通用错误码接口的定义。
# <翻译结束>


<原文开始>
// Code returns the integer number of current error code.
<原文结束>

# <翻译开始>
// Code 返回当前错误代码的整数值。
# <翻译结束>


<原文开始>
// Message returns the brief message for current error code.
<原文结束>

# <翻译开始>
// Message 返回当前错误代码的简短消息。
# <翻译结束>


<原文开始>
	// Detail returns the detailed information of current error code,
	// which is mainly designed as an extension field for error code.
<原文结束>

# <翻译开始>
// Detail 返回当前错误代码的详细信息，
// 主要设计为错误代码的扩展字段。
# <翻译结束>


<原文开始>
// ================================================================================================================
// Common error code definition.
// There are reserved internal error code by framework: code < 1000.
// ================================================================================================================
<原文结束>

# <翻译开始>
// ================================================================================================================
// 公共错误码定义。
// 框架内部预留了以下错误码：code < 1000。
// ================================================================================================================
# <翻译结束>







<原文开始>
// An error occurred internally.
<原文结束>

# <翻译开始>
// 发生了内部错误。
# <翻译结束>












<原文开始>
// The given parameter for current operation is invalid.
<原文结束>

# <翻译开始>
// 当前操作给定的参数无效。
# <翻译结束>


<原文开始>
// Parameter for current operation is missing.
<原文结束>

# <翻译开始>
// 当前操作缺少参数。
# <翻译结束>


<原文开始>
// The function cannot be used like this.
<原文结束>

# <翻译开始>
// 该函数不能这样使用。
# <翻译结束>


<原文开始>
// The configuration is invalid for current operation.
<原文结束>

# <翻译开始>
// 当前操作的配置无效。
# <翻译结束>


<原文开始>
// The configuration is missing for current operation.
<原文结束>

# <翻译开始>
// 当前操作缺少配置。
# <翻译结束>


<原文开始>
// The operation is not implemented yet.
<原文结束>

# <翻译开始>
// 该操作尚未实现。
# <翻译结束>


<原文开始>
// The operation is not supported yet.
<原文结束>

# <翻译开始>
// 该操作尚未被支持。
# <翻译结束>


<原文开始>
// I tried, but I cannot give you what you want.
<原文结束>

# <翻译开始>
// 我尝试了，但我无法给你你想要的东西。
# <翻译结束>


<原文开始>
// Server is busy, please try again later.
<原文结束>

# <翻译开始>
// 服务器繁忙，请稍后再试。
# <翻译结束>







<原文开始>
// It needs necessary package import.
<原文结束>

# <翻译开始>
// 它需要必要的包导入。
# <翻译结束>


<原文开始>
// An panic occurred internally.
<原文结束>

# <翻译开始>
// 发生了内部 panic
# <翻译结束>


<原文开始>
// Business validation failed.
<原文结束>

# <翻译开始>
// 业务验证失败。
# <翻译结束>


<原文开始>
// New creates and returns an error code.
// Note that it returns an interface object of Code.
<原文结束>

# <翻译开始>
// New 创建并返回一个错误代码。
// 注意，它返回的是 Code 接口对象。
# <翻译结束>


<原文开始>
// WithCode creates and returns a new error code based on given Code.
// The code and message is from given `code`, but the detail if from given `detail`.
<原文结束>

# <翻译开始>
// WithCode根据给定的Code创建并返回一个新的错误代码。
// 代码和消息来自给定的`code`，但详情来自给定的`detail`。
# <翻译结束>

















<原文开始>
// No error code specified.
<原文结束>

# <翻译开始>
// 未指定错误代码。
# <翻译结束>


<原文开始>
// Data validation failed.
<原文结束>

# <翻译开始>
// 数据验证失败。
# <翻译结束>


<原文开始>
// Database operation error.
<原文结束>

# <翻译开始>
// 数据库操作错误。
# <翻译结束>


<原文开始>
// Resource does not exist.
<原文结束>

# <翻译开始>
// 资源不存在。
# <翻译结束>

