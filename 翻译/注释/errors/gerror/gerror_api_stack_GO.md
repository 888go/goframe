
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
// stack represents a stack of program counters.
<原文结束>

# <翻译开始>
// stack代表一个程序计数器的堆栈。
# <翻译结束>


<原文开始>
// maxStackDepth marks the max stack depth for error back traces.
<原文结束>

# <翻译开始>
// maxStackDepth 标记了错误回溯时的最大堆栈深度。
# <翻译结束>


<原文开始>
// Cause returns the root cause error of `err`.
<原文结束>

# <翻译开始>
// Cause 返回 `err` 的根本原因错误。
# <翻译结束>


<原文开始>
// Stack returns the stack callers as string.
// It returns the error string directly if the `err` does not support stacks.
<原文结束>

# <翻译开始>
// Stack 返回调用栈信息作为字符串。
// 如果`err`不支持堆栈信息，则直接返回错误字符串。
# <翻译结束>


<原文开始>
// Current creates and returns the current level error.
// It returns nil if current level error is nil.
<原文结束>

# <翻译开始>
// Current函数创建并返回当前层级的错误信息。
// 如果当前层级的错误信息为nil，则该函数返回nil。
# <翻译结束>


<原文开始>
// Unwrap returns the next level error.
// It returns nil if current level error or the next level error is nil.
<原文结束>

# <翻译开始>
// Unwrap 返回下一层级的错误。
// 如果当前层级的错误或下一层级的错误为 nil，它将返回 nil。
# <翻译结束>


<原文开始>
// HasStack checks and reports whether `err` implemented interface `gerror.IStack`.
<原文结束>

# <翻译开始>
// HasStack 检查并报告 `err` 是否实现了接口 `gerror.IStack`。
# <翻译结束>


<原文开始>
// Equal reports whether current error `err` equals to error `target`.
// Please note that, in default comparison logic for `Error`,
// the errors are considered the same if both the `code` and `text` of them are the same.
<原文结束>

# <翻译开始>
// Equal 判断当前错误 `err` 是否等于目标错误 `target`。
// 请注意，在默认的 `Error` 比较逻辑中，
// 如果两个错误的 `code` 和 `text` 都相同，则认为它们是相同的错误。
# <翻译结束>


<原文开始>
// Is reports whether current error `err` has error `target` in its chaining errors.
// It is just for implements for stdlib errors.Is from Go version 1.17.
<原文结束>

# <翻译开始>
// Is 报告当前错误 `err` 在其链式错误中是否包含错误 `target`。
// 这只是为了实现从 Go 1.17 版本开始的stdlib errors.Is功能。
# <翻译结束>


<原文开始>
// HasError is alias of Is, which more easily understanding semantics.
<原文结束>

# <翻译开始>
// HasError 是 Is 的别名，它具有更易于理解的语义。
# <翻译结束>


<原文开始>
// callers returns the stack callers.
// Note that it here just retrieves the caller memory address array not the caller information.
<原文结束>

# <翻译开始>
// callers 返回调用栈的调用者信息。
// 注意，此处它仅获取调用者内存地址数组，并非调用者详细信息。
# <翻译结束>

