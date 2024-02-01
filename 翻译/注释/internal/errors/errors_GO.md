
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
// Package errors provides functionalities to manipulate errors for internal usage purpose.
<原文结束>

# <翻译开始>
// Package errors 提供用于内部使用目的的错误处理功能。
# <翻译结束>


<原文开始>
// StackMode is the mode that printing stack information in StackModeBrief or StackModeDetail mode.
<原文结束>

# <翻译开始>
// StackMode 是打印堆栈信息的模式，可在 StackModeBrief（简要模式）或 StackModeDetail（详细模式）中选择。
# <翻译结束>


<原文开始>
	// commandEnvKeyForBrief is the command environment name for switch key for brief error stack.
	// Deprecated: use commandEnvKeyForStackMode instead.
<原文结束>

# <翻译开始>
// commandEnvKeyForBrief 是用于简短错误堆栈切换键的命令环境名称。
// 已弃用：请改用 commandEnvKeyForStackMode。
# <翻译结束>


<原文开始>
// commandEnvKeyForStackMode is the command environment name for switch key for brief error stack.
<原文结束>

# <翻译开始>
// commandEnvKeyForStackMode 是用于简略错误堆栈模式切换键的命令环境名称。
# <翻译结束>


<原文开始>
// StackModeBrief specifies all error stacks printing no framework error stacks.
<原文结束>

# <翻译开始>
// StackModeBrief 指定错误堆栈打印时，不包含框架内部的错误堆栈信息。
# <翻译结束>


<原文开始>
// StackModeDetail specifies all error stacks printing detailed error stacks including framework stacks.
<原文结束>

# <翻译开始>
// StackModeDetail 指定打印所有错误堆栈，包括框架堆栈的详细错误堆栈。
# <翻译结束>


<原文开始>
	// stackModeConfigured is the configured error stack mode variable.
	// It is brief stack mode in default.
<原文结束>

# <翻译开始>
// stackModeConfigured 是配置的错误堆栈模式变量。
// 默认情况下，它是简洁堆栈模式。
# <翻译结束>


<原文开始>
// The error stack mode is configured using command line arguments or environments.
<原文结束>

# <翻译开始>
// 错误堆栈模式通过命令行参数或环境变量进行配置。
# <翻译结束>


<原文开始>
// IsStackModeBrief returns whether current error stack mode is in brief mode.
<原文结束>

# <翻译开始>
// IsStackModeBrief 返回当前错误堆栈模式是否处于简洁模式。
# <翻译结束>

