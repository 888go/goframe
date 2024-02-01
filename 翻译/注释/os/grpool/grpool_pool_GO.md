
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
// Add pushes a new job to the pool.
// The job will be executed asynchronously.
<原文结束>

# <翻译开始>
// Add 向任务池中添加一个新的任务。
// 该任务将会被异步执行。
# <翻译结束>







<原文开始>
// AddWithRecover pushes a new job to the pool with specified recover function.
//
// The optional `recoverFunc` is called when any panic during executing of `userFunc`.
// If `recoverFunc` is not passed or given nil, it ignores the panic from `userFunc`.
// The job will be executed asynchronously.
<原文结束>

# <翻译开始>
// AddWithRecover 向池中添加一个带有指定恢复函数的新任务。
//
// 可选的 `recoverFunc` 参数会在执行 `userFunc` 函数期间发生 panic 时被调用。
// 如果未传递 `recoverFunc` 或其值为 nil，则会忽略来自 `userFunc` 的 panic。
// 该任务将被异步执行。
# <翻译结束>


<原文开始>
// Cap returns the capacity of the pool.
// This capacity is defined when pool is created.
// It returns -1 if there's no limit.
<原文结束>

# <翻译开始>
// Cap 返回池的容量。
// 这个容量在创建池时定义。
// 如果没有限制，则返回 -1。
# <翻译结束>


<原文开始>
// Size returns current goroutine count of the pool.
<原文结束>

# <翻译开始>
// Size 返回当前 goroutine 池中的goroutine数量。
# <翻译结束>


<原文开始>
// Jobs returns current job count of the pool.
// Note that, it does not return worker/goroutine count but the job/task count.
<原文结束>

# <翻译开始>
// Jobs 返回当前工作池中的任务数量。
// 注意，它返回的是任务/任务计数，而不是工作者/协程计数。
# <翻译结束>


<原文开始>
// IsClosed returns if pool is closed.
<原文结束>

# <翻译开始>
// IsClosed 返回 pool 是否已关闭。
# <翻译结束>


<原文开始>
// Close closes the goroutine pool, which makes all goroutines exit.
<原文结束>

# <翻译开始>
// Close 关闭 goroutine 池，这将使所有 goroutine 退出。
# <翻译结束>


<原文开始>
// checkAndForkNewGoroutineWorker checks and creates a new goroutine worker.
// Note that the worker dies if the job function panics and the job has no recover handling.
<原文结束>

# <翻译开始>
// checkAndForkNewGoroutineWorker 检查并创建一个新的 goroutine 工作者。
// 注意，如果 job 函数出现 panic，并且该任务未进行 recover 处理，则该工作者会终止。
# <翻译结束>


<原文开始>
// Check whether fork new goroutine or not.
<原文结束>

# <翻译开始>
// 检查是否需要新建一个goroutine。
# <翻译结束>


<原文开始>
// No need fork new goroutine.
<原文结束>

# <翻译开始>
// 无需创建新的goroutine。
# <翻译结束>


<原文开始>
// Use CAS to guarantee atomicity.
<原文结束>

# <翻译开始>
// 使用CAS（Compare And Swap）来保证原子性。
# <翻译结束>


<原文开始>
// Create job function in goroutine.
<原文结束>

# <翻译开始>
// 在goroutine中创建任务函数。
# <翻译结束>


<原文开始>
// Harding working, one by one, job never empty, worker never die.
<原文结束>

# <翻译开始>
// 哈丁辛勤工作，逐个完成任务，工作永不停歇，工作者永不消亡。
# <翻译结束>


<原文开始>
// Check and fork new worker.
<原文结束>

# <翻译开始>
// 检查并创建新的工作进程。
# <翻译结束>

