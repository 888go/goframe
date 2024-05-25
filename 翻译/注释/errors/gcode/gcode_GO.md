
<原文开始>
// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457
# <翻译结束>


<原文开始>
// Package gcode provides universal error code definition and common error codes implements.
<原文结束>

# <翻译开始>
// 包gcode提供了通用的错误代码定义和常见的错误代码实现。 md5:cb91541987c67096
# <翻译结束>


<原文开始>
// Code is universal error code interface definition.
<原文结束>

# <翻译开始>
// Code 是一个通用错误代码接口的定义。 md5:bc72f9cd69a9f042
# <翻译结束>


<原文开始>
// Code returns the integer number of current error code.
<原文结束>

# <翻译开始>
// Code 返回当前错误代码的整数值。 md5:75b8de0b4b9fa0a7
# <翻译结束>


<原文开始>
// Message returns the brief message for current error code.
<原文结束>

# <翻译开始>
// Message返回当前错误代码的简要消息。 md5:e0440d2d9a5b929c
# <翻译结束>


<原文开始>
	// Detail returns the detailed information of current error code,
	// which is mainly designed as an extension field for error code.
<原文结束>

# <翻译开始>
// Detail返回当前错误代码的详细信息，主要用于作为错误代码的扩展字段。
// md5:b363ac7e7695be15
# <翻译结束>


<原文开始>
// ================================================================================================================
// Common error code definition.
// There are reserved internal error code by framework: code < 1000.
// ================================================================================================================
<原文结束>

# <翻译开始>
//==============================================================================================================
// 公共错误码定义。
// 框架保留了内部错误码的使用范围：代码小于1000。
//==============================================================================================================
// md5:aeebc2e4a8ad2666
# <翻译结束>


<原文开始>
// No error code specified.
<原文结束>

# <翻译开始>
// 没有指定错误代码。 md5:f3402e31e47f29a9
# <翻译结束>


<原文开始>
// An error occurred internally.
<原文结束>

# <翻译开始>
// 发生了内部错误。 md5:68452eba157c4f37
# <翻译结束>


<原文开始>
// Data validation failed.
<原文结束>

# <翻译开始>
// 数据验证失败。 md5:9bd6126b3a2cb386
# <翻译结束>


<原文开始>
// Database operation error.
<原文结束>

# <翻译开始>
// 数据库操作错误。 md5:67c037697b9e335d
# <翻译结束>


<原文开始>
// The given parameter for current operation is invalid.
<原文结束>

# <翻译开始>
// 给定的当前操作参数无效。 md5:ca885036e7406885
# <翻译结束>


<原文开始>
// Parameter for current operation is missing.
<原文结束>

# <翻译开始>
// 当前操作缺少参数。 md5:1ce758fa97191ebc
# <翻译结束>


<原文开始>
// The function cannot be used like this.
<原文结束>

# <翻译开始>
// 这个函数不能这样使用。 md5:a02d2635b1d2a487
# <翻译结束>


<原文开始>
// The configuration is invalid for current operation.
<原文结束>

# <翻译开始>
// 当前操作的配置无效。 md5:babdd505987f15c5
# <翻译结束>


<原文开始>
// The configuration is missing for current operation.
<原文结束>

# <翻译开始>
// 当前操作的配置缺失。 md5:8f05e88006bb7f7f
# <翻译结束>


<原文开始>
// The operation is not implemented yet.
<原文结束>

# <翻译开始>
// 此操作尚未实现。 md5:5277696d372ccedc
# <翻译结束>


<原文开始>
// The operation is not supported yet.
<原文结束>

# <翻译开始>
// 此操作尚不支持。 md5:90cc232b1a9aa21e
# <翻译结束>


<原文开始>
// I tried, but I cannot give you what you want.
<原文结束>

# <翻译开始>
// 我尝试了，但是我无法给你你想要的。 md5:a8cb7ffbfd6211e4
# <翻译结束>


<原文开始>
// Server is busy, please try again later.
<原文结束>

# <翻译开始>
// 服务器正忙，请稍后重试。 md5:474334c09e329e2d
# <翻译结束>


<原文开始>
// Resource does not exist.
<原文结束>

# <翻译开始>
// 资源不存在。 md5:4e9493277f9141d8
# <翻译结束>


<原文开始>
// It needs necessary package import.
<原文结束>

# <翻译开始>
// 它需要必要的包导入。 md5:bd34126e0df110ff
# <翻译结束>


<原文开始>
// An panic occurred internally.
<原文结束>

# <翻译开始>
// 内部发生了恐慌。 md5:f12430dbb6bb9ee9
# <翻译结束>


<原文开始>
// Business validation failed.
<原文结束>

# <翻译开始>
// 业务验证失败。 md5:816812c09b9bed71
# <翻译结束>


<原文开始>
// New creates and returns an error code.
// Note that it returns an interface object of Code.
<原文结束>

# <翻译开始>
// New 创建并返回一个错误代码。
// 注意，它返回一个 Code 接口对象。
// md5:a3d7ec3807589165
# <翻译结束>


<原文开始>
// WithCode creates and returns a new error code based on given Code.
// The code and message is from given `code`, but the detail if from given `detail`.
<原文结束>

# <翻译开始>
// WithCode 根据给定的`Code`创建并返回一个新的错误代码。
// 该错误代码的码和消息来自`code`，但详细信息来自`detail`。
// md5:6f2355f302e9ea32
# <翻译结束>

