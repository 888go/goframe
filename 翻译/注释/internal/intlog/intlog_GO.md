
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
// Package intlog provides internal logging for GoFrame development usage only.
<原文结束>

# <翻译开始>
// Package intlog 提供内部日志功能，仅用于 GoFrame 开发使用。
# <翻译结束>


<原文开始>
// Print prints `v` with newline using fmt.Println.
// The parameter `v` can be multiple variables.
<原文结束>

# <翻译开始>
// Print 通过 fmt.Println 打印 `v`（并附带换行）。参数 `v` 可以是多个变量。
# <翻译结束>


<原文开始>
// Printf prints `v` with format `format` using fmt.Printf.
// The parameter `v` can be multiple variables.
<原文结束>

# <翻译开始>
// Printf 使用 fmt.Printf 格式化并打印 `v`。其中参数 `v` 可以是多个变量。
// ```go
// Printf 函数通过 fmt.Printf 格式化输出 `v`。
// 注意，这里的 `v` 参数可以接受多个变量。
# <翻译结束>


<原文开始>
// Error prints `v` with newline using fmt.Println.
// The parameter `v` can be multiple variables.
<原文结束>

# <翻译开始>
// Error 使用 fmt.Println 打印 `v`（并附带换行）。参数 `v` 可以是多个变量。
# <翻译结束>


<原文开始>
// Errorf prints `v` with format `format` using fmt.Printf.
<原文结束>

# <翻译开始>
// Errorf使用fmt.Printf格式化方式打印变量v，格式字符串为format。
# <翻译结束>


<原文开始>
// PrintFunc prints the output from function `f`.
// It only calls function `f` if debug mode is enabled.
<原文结束>

# <翻译开始>
// PrintFunc 用于打印函数 `f` 的输出结果。
// 只有在调试模式开启时，才会调用函数 `f`。
# <翻译结束>


<原文开始>
// ErrorFunc prints the output from function `f`.
// It only calls function `f` if debug mode is enabled.
<原文结束>

# <翻译开始>
// ErrorFunc 用于打印函数 `f` 的输出结果。
// 只有在调试模式开启的情况下，才会调用函数 `f`。
# <翻译结束>


<原文开始>
// traceIdStr retrieves and returns the trace id string for logging output.
<原文结束>

# <翻译开始>
// traceIdStr 用于获取并返回日志输出的追踪ID字符串。
# <翻译结束>


<原文开始>
// file returns caller file name along with its line number.
<原文结束>

# <翻译开始>
// file 返回调用文件名及其行号。
# <翻译结束>

