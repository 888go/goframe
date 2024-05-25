
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
// stack represents a stack of program counters.
<原文结束>

# <翻译开始>
// stack 表示一个程序计数器栈。 md5:974ffc07f2ddbe11
# <翻译结束>


<原文开始>
// maxStackDepth marks the max stack depth for error back traces.
<原文结束>

# <翻译开始>
// maxStackDepth 标记错误回溯的最大堆栈深度。 md5:e4020e5e182a92f7
# <翻译结束>


<原文开始>
// Cause returns the root cause error of `err`.
<原文结束>

# <翻译开始>
// Cause 返回 `err` 的根本原因错误。 md5:3719c97ae5cd6a94
# <翻译结束>


<原文开始>
// Stack returns the stack callers as string.
// It returns the error string directly if the `err` does not support stacks.
<原文结束>

# <翻译开始>
// Stack 返回调用堆栈作为字符串。
// 如果 `err` 不支持堆栈信息，它将直接返回错误字符串。
// md5:bcba4c4f440cb0a7
# <翻译结束>


<原文开始>
// Current creates and returns the current level error.
// It returns nil if current level error is nil.
<原文结束>

# <翻译开始>
// Current 创建并返回当前级别的错误。如果当前级别错误为 nil，则返回 nil。
// md5:d8b26e22ec63a837
# <翻译结束>


<原文开始>
// Unwrap returns the next level error.
// It returns nil if current level error or the next level error is nil.
<原文结束>

# <翻译开始>
// Unwrap返回下一级错误。
// 如果当前级别或下一级错误为nil，则返回nil。
// md5:97894c3cda2b0c4e
# <翻译结束>


<原文开始>
// HasStack checks and reports whether `err` implemented interface `gerror.IStack`.
<原文结束>

# <翻译开始>
// HasStack 检查并报告 `err` 是否实现了接口 `gerror.IStack`。 md5:f8ab57042234eea7
# <翻译结束>


<原文开始>
// Equal reports whether current error `err` equals to error `target`.
// Please note that, in default comparison logic for `Error`,
// the errors are considered the same if both the `code` and `text` of them are the same.
<原文结束>

# <翻译开始>
// Equal 检查当前错误 `err` 是否等于错误 `target`。
// 请注意，在默认的 `Error` 比较逻辑中，如果两个错误的 `code` 和 `text` 都相同，则认为它们是相同的。
// md5:adaa63023ba44f79
# <翻译结束>


<原文开始>
// Is reports whether current error `err` has error `target` in its chaining errors.
// It is just for implements for stdlib errors.Is from Go version 1.17.
<原文结束>

# <翻译开始>
// Is 判断当前错误 `err` 是否在其嵌套错误中包含目标错误 `target`。这是为了实现从 Go 1.17 版本开始的标准库中的 errors.Is 接口。
// md5:dfc92c8d3ba58133
# <翻译结束>


<原文开始>
// HasError is alias of Is, which more easily understanding semantics.
<原文结束>

# <翻译开始>
// HasError是Is的别名，它更易于理解语义。 md5:d84dea05563aadb6
# <翻译结束>


<原文开始>
// callers returns the stack callers.
// Note that it here just retrieves the caller memory address array not the caller information.
<原文结束>

# <翻译开始>
// callers 返回调用栈的调用者。
// 注意，这里只是获取调用者内存地址数组，并非调用者的具体信息。
// md5:9c65fc07c2395a65
# <翻译结束>

