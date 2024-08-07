
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
// Go creates a new asynchronous goroutine function with specified recover function.
//
// The parameter `recoverFunc` is called when any panic during executing of `goroutineFunc`.
// If `recoverFunc` is given nil, it ignores the panic from `goroutineFunc` and no panic will
// throw to parent goroutine.
//
// But, note that, if `recoverFunc` also throws panic, such panic will be thrown to parent goroutine.
<原文结束>

# <翻译开始>
// Go 创建一个新的异步 goroutine 函数，并指定了恢复函数。
//
// 参数 `recoverFunc` 在 `goroutineFunc` 执行过程中发生任何 panic 时被调用。
// 如果 `recoverFunc` 为 nil，它将忽略 `goroutineFunc` 中的 panic，且不会向父 goroutine 抛出 panic。
//
// 但是要注意，如果 `recoverFunc` 自身也抛出 panic，这个 panic 将会被抛给父 goroutine。
// md5:3820395064a9e843
# <翻译结束>

