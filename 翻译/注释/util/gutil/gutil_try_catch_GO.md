
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
// Throw throws out an exception, which can be caught be TryCatch or recover.
<原文结束>

# <翻译开始>
// Throw 抛出一个异常，这个异常可以被 TryCatch 结构或 recover 语句捕获。. md5:44e69b1d2fded57f
# <翻译结束>


<原文开始>
// Try implements try... logistics using internal panic...recover.
// It returns error if any exception occurs, or else it returns nil.
<原文结束>

# <翻译开始>
// Try 使用内部的 panic...recover 实现 try...catch 逻辑。如果发生任何异常，它会返回错误；否则返回 nil。
// md5:7c85aa857bb16fca
# <翻译结束>


<原文开始>
// TryCatch implements `try...catch..`. logistics using internal `panic...recover`.
// It automatically calls function `catch` if any exception occurs and passes the exception as an error.
// If `catch` is given nil, it ignores the panic from `try` and no panic will throw to parent goroutine.
//
// But, note that, if function `catch` also throws panic, the current goroutine will panic.
<原文结束>

# <翻译开始>
// TryCatch 实现了类似于 `try...catch...` 的异常处理机制，利用内部的 `panic...recover` 逻辑。
// 当在 `try` 块中发生任何异常时，它会自动调用函数 `catch` 并将异常传递给错误参数。
// 如果 `catch` 函数为 nil，则忽略 `try` 中的 panic，不会向父 goroutine 抛出 panic。
//
// 但是，需要注意的是，如果 `catch` 函数自身也抛出 panic，当前 goroutine 将会 panic。
// md5:6cbe568aa0940b07
# <翻译结束>

