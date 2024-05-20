
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
// Add pushes a new job to the pool.
// The job will be executed asynchronously.
<原文结束>

# <翻译开始>
// Add 将一个新任务添加到池中。
// 该任务将会异步执行。
// md5:69389d53e280086b
# <翻译结束>


<原文开始>
// Check and fork new worker.
<原文结束>

# <翻译开始>
// 检查并 fork 新的 worker。. md5:d3acb042c3373fa4
# <翻译结束>


<原文开始>
// AddWithRecover pushes a new job to the pool with specified recover function.
//
// The optional `recoverFunc` is called when any panic during executing of `userFunc`.
// If `recoverFunc` is not passed or given nil, it ignores the panic from `userFunc`.
// The job will be executed asynchronously.
<原文结束>

# <翻译开始>
// AddWithRecover 将指定的恢复函数推送到队列中执行新任务。
// 
// 可选的 `recoverFunc` 在执行 `userFunc` 时发生任何 panic 时被调用。如果未传递或给定 nil，它将忽略来自 `userFunc` 的 panic。任务将异步执行。
// md5:764d1260466b9a5d
# <翻译结束>


<原文开始>
// Cap returns the capacity of the pool.
// This capacity is defined when pool is created.
// It returns -1 if there's no limit.
<原文结束>

# <翻译开始>
// Cap 返回池的容量。
// 这个容量在创建池时定义。
// 如果没有限制，则返回-1。
// md5:1c6cae16429df1b2
# <翻译结束>


<原文开始>
// Size returns current goroutine count of the pool.
<原文结束>

# <翻译开始>
// Size 返回当前池中的goroutine数量。. md5:247eb1685633ccc3
# <翻译结束>


<原文开始>
// Jobs returns current job count of the pool.
// Note that, it does not return worker/goroutine count but the job/task count.
<原文结束>

# <翻译开始>
// Jobs 返回池中的当前任务数。
// 注意，它返回的不是工作器/goroutine的数量，而是任务的数量。
// md5:c82d92b33047974c
# <翻译结束>


<原文开始>
// IsClosed returns if pool is closed.
<原文结束>

# <翻译开始>
// IsClosed 返回池是否已关闭。. md5:85755176347bcfea
# <翻译结束>


<原文开始>
// Close closes the goroutine pool, which makes all goroutines exit.
<原文结束>

# <翻译开始>
// Close 关闭goroutine池，导致所有goroutines退出。. md5:3d9c73ed9b0f4643
# <翻译结束>


<原文开始>
// checkAndForkNewGoroutineWorker checks and creates a new goroutine worker.
// Note that the worker dies if the job function panics and the job has no recover handling.
<原文结束>

# <翻译开始>
// checkAndForkNewGoroutineWorker 检查并创建一个新的goroutine工作进程。
// 请注意，如果工作函数出现恐慌且该工作没有恢复处理，那么该工作进程将会死亡。
// md5:242a912451066181
# <翻译结束>


<原文开始>
// Check whether fork new goroutine or not.
<原文结束>

# <翻译开始>
// 检查是否需要在新的goroutine中 fork。. md5:20ef20b082ef0b86
# <翻译结束>


<原文开始>
// No need fork new goroutine.
<原文结束>

# <翻译开始>
// 不需要启动新的goroutine。. md5:a4d7257aa086311e
# <翻译结束>


<原文开始>
// Use CAS to guarantee atomicity.
<原文结束>

# <翻译开始>
// 使用CAS（比较并交换）来保证操作的原子性。. md5:2337a31243acf132
# <翻译结束>


<原文开始>
// Create job function in goroutine.
<原文结束>

# <翻译开始>
// 在goroutine中创建任务函数。. md5:e0e70df051fd0b1a
# <翻译结束>


<原文开始>
// Harding working, one by one, job never empty, worker never die.
<原文结束>

# <翻译开始>
// 哈丁工作，一个接一个，任务永无止境，工人永不消亡。. md5:625670ae6a926602
# <翻译结束>

