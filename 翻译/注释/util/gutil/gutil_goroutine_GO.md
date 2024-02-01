
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
// Go creates a new asynchronous goroutine function with specified recover function.
//
// The parameter `recoverFunc` is called when any panic during executing of `goroutineFunc`.
// If `recoverFunc` is given nil, it ignores the panic from `goroutineFunc` and no panic will
// throw to parent goroutine.
//
// But, note that, if `recoverFunc` also throws panic, such panic will be thrown to parent goroutine.
<原文结束>

# <翻译开始>
// Go 创建了一个具有指定恢复函数的新的异步 goroutine 函数。
//
// 参数 `recoverFunc` 在执行 `goroutineFunc` 期间发生任何 panic 时被调用。
// 如果 `recoverFunc` 被赋予 nil，它将忽略来自 `goroutineFunc` 的 panic，并且不会向父级 goroutine 抛出 panic。
//
// 但是请注意，如果 `recoverFunc` 也抛出了 panic，这样的 panic 将会被抛给父级 goroutine。
# <翻译结束>

