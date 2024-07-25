
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Entry is the timing job.
<原文结束>

# <翻译开始>
// Entry是定时任务。 md5:50a66c0cedad73c2
# <翻译结束>


<原文开始>
// The context for the job, for READ ONLY.
<原文结束>

# <翻译开始>
// 作业的上下文，只读。 md5:f8c9c52b497bf322
# <翻译结束>


<原文开始>
// The job runs every tick.
<原文结束>

# <翻译开始>
// 任务每滴答一声就执行一次。 md5:ee9f167eaedbe210
# <翻译结束>


<原文开始>
// Next run ticks of the job.
<原文结束>

# <翻译开始>
// 下一次运行任务的ticks。 md5:41046b30b7471bf3
# <翻译结束>


<原文开始>
// JobFunc is the timing called job function in timer.
<原文结束>

# <翻译开始>
// JobFunc 是定时器中调用的定时任务函数。 md5:8958cfb2ccc06eff
# <翻译结束>


<原文开始>
// Status returns the status of the job.
<原文结束>

# <翻译开始>
// Status 返回作业的状态。 md5:2147922a20ade271
# <翻译结束>


<原文开始>
// Run runs the timer job asynchronously.
<原文结束>

# <翻译开始>
// Run 异步运行计时器任务。 md5:11fb4e6232736ab7
# <翻译结束>


<原文开始>
// It checks its running times exceeding.
<原文结束>

# <翻译开始>
		// 检查其运行时间是否超时。 md5:040ad43c2af11b3d
# <翻译结束>


<原文开始>
// doCheckAndRunByTicks checks the if job can run in given timer ticks,
// it runs asynchronously if the given `currentTimerTicks` meets or else
// it increments its ticks and waits for next running check.
<原文结束>

# <翻译开始>
// doCheckAndRunByTicks 检查在给定的计时器周期内，任务是否可以运行。如果当前的`currentTimerTicks`满足条件，它会异步运行；否则，它会增加其周期并等待下一次运行检查。 md5:44d5223afb2e4f9d
# <翻译结束>


<原文开始>
// SetStatus custom sets the status for the job.
<原文结束>

# <翻译开始>
// SetStatus 自定义设置作业的状态。 md5:c143d90d99990f2c
# <翻译结束>


<原文开始>
// Close closes the job, and then it will be removed from the timer.
<原文结束>

# <翻译开始>
// Close 方法关闭任务，随后该任务将从计时器中移除。 md5:f499b51290bff676
# <翻译结束>


<原文开始>
// Reset resets the job, which resets its ticks for next running.
<原文结束>

# <翻译开始>
// Reset 重置作业，这将为下次运行重置其计数器。 md5:5a5ab5e4b73a76fe
# <翻译结束>


<原文开始>
// IsSingleton checks and returns whether the job in singleton mode.
<原文结束>

# <翻译开始>
// IsSingleton 检查并返回当前任务是否处于单例模式。 md5:a380e519564eb9da
# <翻译结束>


<原文开始>
// SetSingleton sets the job singleton mode.
<原文结束>

# <翻译开始>
// SetSingleton 设置单例模式。 md5:3fd379a01f57d11e
# <翻译结束>


<原文开始>
// Job returns the job function of this job.
<原文结束>

# <翻译开始>
// Job 返回此任务的工作函数。 md5:38a44e496baf9d51
# <翻译结束>


<原文开始>
// Ctx returns the initialized context of this job.
<原文结束>

# <翻译开始>
// Ctx 返回此任务的初始化上下文。 md5:c2a3cb7932f8bf8a
# <翻译结束>


<原文开始>
// SetTimes sets the limit running times for the job.
<原文结束>

# <翻译开始>
// SetTimes 设置作业的运行次数限制。 md5:812717e2b2bcce7c
# <翻译结束>

