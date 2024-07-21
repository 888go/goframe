
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
// JobFunc is the timing called job function in cron.
<原文结束>

# <翻译开始>
// JobFunc 是在cron中定时被调用的作业函数。 md5:476072dbc1ca96ff
# <翻译结束>


<原文开始>
// Entry is timing task entry.
<原文结束>

# <翻译开始>
// Entry 是定时任务的入口。 md5:ebc0ffec0c819fa5
# <翻译结束>


<原文开始>
// Cron object belonged to.
<原文结束>

# <翻译开始>
// 属于的Cron对象。 md5:b0d31cd3bc86924f
# <翻译结束>


<原文开始>
// Associated timer Entry.
<原文结束>

# <翻译开始>
// 相关的定时器条目。 md5:cc0c92f42e2f28c0
# <翻译结束>


<原文开始>
// Callback function name(address info).
<原文结束>

# <翻译开始>
// 回调函数名称（地址信息）。 md5:7f130ad66e787420
# <翻译结束>


<原文开始>
// Name names this entry for manual control.
<原文结束>

# <翻译开始>
// Name 为手动控制此条目提供名称。 md5:ab6add69f0d6cead
# <翻译结束>


<原文开始>
// Job is the callback function for timed task execution.
<原文结束>

# <翻译开始>
// Job是定时任务执行的回调函数。 md5:7faad71757692791
# <翻译结束>


<原文开始>
// The context for the job.
<原文结束>

# <翻译开始>
// 作业的上下文。 md5:505a832c1cca7c0e
# <翻译结束>


<原文开始>
// Times specifies the running limit times for the entry.
<原文结束>

# <翻译开始>
// Times 指定了条目的运行限制次数。 md5:bb45eebc85ac5fcc
# <翻译结束>


<原文开始>
// Pattern is the crontab style string for scheduler.
<原文结束>

# <翻译开始>
// Pattern 是用于调度的 crontab 风格的字符串。 md5:2076c5de0b0a0288
# <翻译结束>


<原文开始>
// Singleton specifies whether timed task executing in singleton mode.
<原文结束>

# <翻译开始>
// Singleton 指定是否以单例模式执行定时任务。 md5:f98d6fb605fc538f
# <翻译结束>


<原文开始>
// Infinite specifies whether this entry is running with no times limit.
<原文结束>

# <翻译开始>
// Infinite表示此条目是否无限运行，不受时间限制。 md5:14267dbbffb84b38
# <翻译结束>


<原文开始>
// doAddEntry creates and returns a new Entry object.
<原文结束>

# <翻译开始>
// doAddEntry 创建并返回一个新的 Entry 对象。 md5:c29537f522377c29
# <翻译结束>


<原文开始>
// No limit for `times`, for timer checking scheduling every second.
<原文结束>

# <翻译开始>
// 对于`times`没有限制，用于每秒检查调度的计时器。 md5:dce371ba28b68c21
# <翻译结束>


<原文开始>
	// When you add a scheduled task, you cannot allow it to run.
	// It cannot start running when added to timer.
	// It should start running after the entry is added to the Cron entries map, to avoid the task
	// from running during adding where the entries do not have the entry information, which might cause panic.
<原文结束>

# <翻译开始>
	// 当你添加一个定时任务时，不能让它立即运行。
	// 它在添加到计时器时不能开始执行。
	// 应该在将条目添加到Cron条目映射后开始运行，以防止在添加过程中（此时条目信息可能不完整）任务运行，从而可能导致恐慌。
	// md5:e2b503aef8166c84
# <翻译结束>


<原文开始>
// IsSingleton return whether this entry is a singleton timed task.
<原文结束>

# <翻译开始>
// IsSingleton 判断这个条目是否是定时单例任务。 md5:171967c731b60f88
# <翻译结束>


<原文开始>
// SetSingleton sets the entry running in singleton mode.
<原文结束>

# <翻译开始>
// SetSingleton 设置以单例模式运行的条目。 md5:6c81a3a09d0ef0b1
# <翻译结束>


<原文开始>
// SetTimes sets the times which the entry can run.
<原文结束>

# <翻译开始>
// SetTimes 设置条目可以运行的时间。 md5:663af054d5aab5e9
# <翻译结束>


<原文开始>
// Status returns the status of entry.
<原文结束>

# <翻译开始>
// Status 返回条目的状态。 md5:6a9d3438dc575881
# <翻译结束>


<原文开始>
// SetStatus sets the status of the entry.
<原文结束>

# <翻译开始>
// SetStatus 设置条目的状态。 md5:ea0ecb4171f3f017
# <翻译结束>


<原文开始>
// Start starts running the entry.
<原文结束>

# <翻译开始>
// Start 开始运行入口函数。 md5:aa729d73eb626ca1
# <翻译结束>


<原文开始>
// Stop stops running the entry.
<原文结束>

# <翻译开始>
// Stop 停止运行条目。 md5:06d53148d6536ce9
# <翻译结束>


<原文开始>
// Close stops and removes the entry from cron.
<原文结束>

# <翻译开始>
// Close 停止并从 cron 中移除条目。 md5:a2a5eee9228cd918
# <翻译结束>


<原文开始>
// checkAndRun is the core timing task check logic.
// This function is called every second.
<原文结束>

# <翻译开始>
// checkAndRun是核心定时任务检查逻辑。
// 这个函数每秒被调用一次。
// md5:e4a94345d01fd7df
# <翻译结束>


<原文开始>
// Exception caught, it logs the error content to logger in default behavior.
<原文结束>

# <翻译开始>
// 捕获到异常，默认行为是将错误内容记录到日志中。 md5:9af43ee64f5795bf
# <翻译结束>

