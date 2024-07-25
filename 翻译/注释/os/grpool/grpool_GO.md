
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
// Package grpool implements a goroutine reusable pool.
<原文结束>

# <翻译开始>
// 包grpool实现了可重用的goroutine池。 md5:8908f4659795e87e
# <翻译结束>


<原文开始>
// Func is the pool function which contains context parameter.
<原文结束>

# <翻译开始>
// Func 是包含上下文参数的池函数。 md5:6974cc4f941bf840
# <翻译结束>


<原文开始>
// RecoverFunc is the pool runtime panic recover function which contains context parameter.
<原文结束>

# <翻译开始>
// RecoverFunc 是一个带有上下文参数的池运行时panic恢复函数。 md5:745651d2d0aad841
# <翻译结束>


<原文开始>
// Pool manages the goroutines using pool.
<原文结束>

# <翻译开始>
// Pool 使用池管理goroutine。 md5:b07df489dc176e1c
# <翻译结束>


<原文开始>
// Max goroutine count limit.
<原文结束>

# <翻译开始>
// 最大goroutine计数限制。 md5:d8f9a54c9665e042
# <翻译结束>


<原文开始>
// Current running goroutine count.
<原文结束>

# <翻译开始>
// 当前正在运行的goroutine数量。 md5:ab841a9a1dadf1a0
# <翻译结束>


<原文开始>
// List for asynchronous job adding purpose.
<原文结束>

# <翻译开始>
// 用于异步添加任务的列表。 md5:65d48b67f1f6833d
# <翻译结束>


<原文开始>
// localPoolItem is the job item storing in job list.
<原文结束>

# <翻译开始>
// localPoolItem是存储在作业列表中的作业项。 md5:2b97e4fa0813a9cc
# <翻译结束>


<原文开始>
// Default goroutine pool.
<原文结束>

# <翻译开始>
// 默认的goroutine池。 md5:a834f6aca53d4225
# <翻译结束>


<原文开始>
// New creates and returns a new goroutine pool object.
// The parameter `limit` is used to limit the max goroutine count,
// which is not limited in default.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的 goroutine 池对象。
// 参数 `limit` 用于限制最大 goroutine 数量，
// 默认情况下不限制。
// md5:523f5833750663c7
# <翻译结束>


<原文开始>
// Add pushes a new job to the default goroutine pool.
// The job will be executed asynchronously.
<原文结束>

# <翻译开始>
// Add 将一个新任务添加到默认的 Goroutine 池中。
// 该任务将会异步执行。
// md5:2edd63bb852da48c
# <翻译结束>


<原文开始>
// AddWithRecover pushes a new job to the default pool with specified recover function.
//
// The optional `recoverFunc` is called when any panic during executing of `userFunc`.
// If `recoverFunc` is not passed or given nil, it ignores the panic from `userFunc`.
// The job will be executed asynchronously.
<原文结束>

# <翻译开始>
// AddWithRecover 将一个新的任务推送到默认池中，指定一个恢复函数。
// 
// 可选的 `recoverFunc` 在执行 `userFunc` 时遇到任何 panic 时被调用。如果未传递或给定 `nil`，则忽略来自 `userFunc` 的 panic。任务将异步执行。
// md5:4b448b4fd7caa604
# <翻译结束>


<原文开始>
// Size returns current goroutine count of default goroutine pool.
<原文结束>

# <翻译开始>
// Size 返回默认goroutine池中的当前goroutine数量。 md5:f66351deb380810c
# <翻译结束>


<原文开始>
// Jobs returns current job count of default goroutine pool.
<原文结束>

# <翻译开始>
// Jobs 返回默认goroutine池当前的任务数量。 md5:d9b300d28d86f6c3
# <翻译结束>

