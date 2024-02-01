
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
// Throw throws out an exception, which can be caught be TryCatch or recover.
<原文结束>

# <翻译开始>
// Throw 抛出一个异常，该异常可以被 TryCatch 或 recover 捕获。
# <翻译结束>


<原文开始>
// Try implements try... logistics using internal panic...recover.
// It returns error if any exception occurs, or else it returns nil.
<原文结束>

# <翻译开始>
// Try 使用内部 panic...recover 实现 try...逻辑。
// 如果发生任何异常，它将返回错误，否则返回 nil。
# <翻译结束>


<原文开始>
// TryCatch implements `try...catch..`. logistics using internal `panic...recover`.
// It automatically calls function `catch` if any exception occurs and passes the exception as an error.
// If `catch` is given nil, it ignores the panic from `try` and no panic will throw to parent goroutine.
//
// But, note that, if function `catch` also throws panic, the current goroutine will panic.
<原文结束>

# <翻译开始>
// TryCatch 实现了类似 `try...catch...` 的逻辑，通过内部使用 `panic...recover`。
// 如果发生任何异常，它会自动调用函数 `catch` 并将异常作为 error 传递给 `catch` 函数。
// 若传入的 `catch` 为 nil，则忽略来自 `try` 的 panic，并且不会向父级 goroutine 抛出 panic。
//
// 但是请注意，如果函数 `catch` 本身也抛出了 panic，则当前 goroutine 将会 panic。
# <翻译结束>

