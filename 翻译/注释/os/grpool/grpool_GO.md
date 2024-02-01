
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
// Package grpool implements a goroutine reusable pool.
<原文结束>

# <翻译开始>
// 包grpool实现了goroutine可复用池。
# <翻译结束>


<原文开始>
// Func is the pool function which contains context parameter.
<原文结束>

# <翻译开始>
// Func 是一个包含上下文参数的池函数。
# <翻译结束>


<原文开始>
// RecoverFunc is the pool runtime panic recover function which contains context parameter.
<原文结束>

# <翻译开始>
// RecoverFunc 是包含上下文参数的运行时 panic 恢复函数，用于池化组件。
# <翻译结束>


<原文开始>
// Pool manages the goroutines using pool.
<原文结束>

# <翻译开始>
// Pool 通过使用 goroutine 池来管理 goroutines。
# <翻译结束>







<原文开始>
// Current running goroutine count.
<原文结束>

# <翻译开始>
// 当前运行的 goroutine 数量
# <翻译结束>


<原文开始>
// List for asynchronous job adding purpose.
<原文结束>

# <翻译开始>
// 用于异步任务添加目的的列表。
# <翻译结束>







<原文开始>
// localPoolItem is the job item storing in job list.
<原文结束>

# <翻译开始>
// localPoolItem 是存储在任务列表中的工作项。
# <翻译结束>







<原文开始>
// New creates and returns a new goroutine pool object.
// The parameter `limit` is used to limit the max goroutine count,
// which is not limited in default.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的 goroutine 池对象。
// 参数 `limit` 用于限制最大goroutine数量，默认情况下不限制。
# <翻译结束>


<原文开始>
// Add pushes a new job to the default goroutine pool.
// The job will be executed asynchronously.
<原文结束>

# <翻译开始>
// Add 向默认的goroutine池中添加一个新的任务。
// 该任务将会被异步执行。
# <翻译结束>


<原文开始>
// AddWithRecover pushes a new job to the default pool with specified recover function.
//
// The optional `recoverFunc` is called when any panic during executing of `userFunc`.
// If `recoverFunc` is not passed or given nil, it ignores the panic from `userFunc`.
// The job will be executed asynchronously.
<原文结束>

# <翻译开始>
// AddWithRecover 将带有指定恢复函数的新任务推送到默认工作池。
//
// 可选的 `recoverFunc` 参数在执行 `userFunc` 函数期间发生任何 panic 时被调用。
// 如果未传递或给定 nil，则忽略来自 `userFunc` 的 panic。
// 该任务将被异步执行。
# <翻译结束>


<原文开始>
// Size returns current goroutine count of default goroutine pool.
<原文结束>

# <翻译开始>
// Size 返回默认协程池中当前的协程数量。
# <翻译结束>


<原文开始>
// Jobs returns current job count of default goroutine pool.
<原文结束>

# <翻译开始>
// Jobs 返回默认goroutine池当前的任务数量。
# <翻译结束>


<原文开始>
// Max goroutine count limit.
<原文结束>

# <翻译开始>
// 最大goroutine数量限制。
# <翻译结束>


<原文开始>
// Is pool closed or not.
<原文结束>

# <翻译开始>
// 是否已关闭池。
# <翻译结束>


<原文开始>
// Default goroutine pool.
<原文结束>

# <翻译开始>
// 默认的 goroutine 池。
# <翻译结束>

