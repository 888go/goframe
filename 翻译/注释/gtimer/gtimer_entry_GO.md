
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
// The context for the job, for READ ONLY.
<原文结束>

# <翻译开始>
// 该上下文为作业的上下文，仅用于 READ ONLY（只读）。
# <翻译结束>

















<原文开始>
// JobFunc is the timing called job function in timer.
<原文结束>

# <翻译开始>
// JobFunc 是在定时器中被定时调用的任务函数。
# <翻译结束>


<原文开始>
// Status returns the status of the job.
<原文结束>

# <翻译开始>
// Status 返回作业的状态。
# <翻译结束>


<原文开始>
// Run runs the timer job asynchronously.
<原文结束>

# <翻译开始>
// Run 启动计时器任务并异步执行。
# <翻译结束>


<原文开始>
// It checks its running times exceeding.
<原文结束>

# <翻译开始>
// 它检查运行时间是否超过限制。
# <翻译结束>


<原文开始>
// doCheckAndRunByTicks checks the if job can run in given timer ticks,
// it runs asynchronously if the given `currentTimerTicks` meets or else
// it increments its ticks and waits for next running check.
<原文结束>

# <翻译开始>
// doCheckAndRunByTicks 根据给定的定时器刻度检查任务是否可以运行，
// 如果给定的 `currentTimerTicks` 满足条件，则异步执行该任务；
// 否则，增加其刻度并等待下次运行检查。
# <翻译结束>












<原文开始>
// SetStatus custom sets the status for the job.
<原文结束>

# <翻译开始>
// SetStatus 自定义设置任务的状态。
# <翻译结束>












<原文开始>
// Close closes the job, and then it will be removed from the timer.
<原文结束>

# <翻译开始>
// Close 关闭作业，随后该作业将从计时器中移除。
# <翻译结束>


<原文开始>
// Reset resets the job, which resets its ticks for next running.
<原文结束>

# <翻译开始>
// Reset 重置作业，这将重置其下一次运行的滴答次数。
# <翻译结束>


<原文开始>
// IsSingleton checks and returns whether the job in singleton mode.
<原文结束>

# <翻译开始>
// IsSingleton 检查并返回该任务是否处于单例模式。
# <翻译结束>


<原文开始>
// SetSingleton sets the job singleton mode.
<原文结束>

# <翻译开始>
// SetSingleton 设置作业单例模式。
# <翻译结束>


<原文开始>
// Job returns the job function of this job.
<原文结束>

# <翻译开始>
// Job 返回此任务的工作函数。
# <翻译结束>


<原文开始>
// Ctx returns the initialized context of this job.
<原文结束>

# <翻译开始>
// Ctx 返回此任务初始化后的上下文。
# <翻译结束>


<原文开始>
// SetTimes sets the limit running times for the job.
<原文结束>

# <翻译开始>
// SetTimes 设置作业的最大运行次数。
# <翻译结束>

















<原文开始>
// Entry is the timing job.
<原文结束>

# <翻译开始>
// Entry 是定时任务。
# <翻译结束>


<原文开始>
// The job runs every tick.
<原文结束>

# <翻译开始>
// 该任务在每个时间间隔运行一次。
# <翻译结束>


<原文开始>
// Limit running times.
<原文结束>

# <翻译开始>
// 限制运行次数。
# <翻译结束>


<原文开始>
// Next run ticks of the job.
<原文结束>

# <翻译开始>
// 下一次运行该任务的时间刻度
# <翻译结束>


<原文开始>
// Perform job checking.
<原文结束>

# <翻译开始>
// 执行任务检查。
# <翻译结束>


<原文开始>
// Perform job running.
<原文结束>

# <翻译开始>
// 执行任务运行。
# <翻译结束>


<原文开始>
// Start starts the job.
<原文结束>

# <翻译开始>
// Start 启动任务。
# <翻译结束>


<原文开始>
// Stop stops the job.
<原文结束>

# <翻译开始>
// Stop 停止作业。
# <翻译结束>

