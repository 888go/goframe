
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
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// JobFunc is the timing called job function in cron.
<原文结束>

# <翻译开始>
// JobFunc 是 cron 中被定时调用的任务函数。
# <翻译结束>


<原文开始>
// Callback function name(address info).
<原文结束>

# <翻译开始>
// 回调函数名称(地址信息)。
# <翻译结束>


<原文开始>
// Name names this entry for manual control.
<原文结束>

# <翻译开始>
// Name 为该条目设置名称以便进行手动控制。
# <翻译结束>


<原文开始>
// Job is the callback function for timed task execution.
<原文结束>

# <翻译开始>
// Job 是用于定时任务执行的回调函数。
# <翻译结束>


<原文开始>
// Times specifies the running limit times for the entry.
<原文结束>

# <翻译开始>
// Times 指定该条目运行的限制次数。
# <翻译结束>


<原文开始>
// Pattern is the crontab style string for scheduler.
<原文结束>

# <翻译开始>
// Pattern 是用于调度器的 crontab 风格字符串。
# <翻译结束>


<原文开始>
// Singleton specifies whether timed task executing in singleton mode.
<原文结束>

# <翻译开始>
// Singleton 指定定时任务是否以单例模式执行。
# <翻译结束>


<原文开始>
// Infinite specifies whether this entry is running with no times limit.
<原文结束>

# <翻译开始>
// Infinite 指定此条目是否在无时间限制的情况下运行。
# <翻译结束>


<原文开始>
// doAddEntry creates and returns a new Entry object.
<原文结束>

# <翻译开始>
// doAddEntry 创建并返回一个新的 Entry 对象。
# <翻译结束>


<原文开始>
// No limit for `times`, for timer checking scheduling every second.
<原文结束>

# <翻译开始>
// 对于 `times` 没有限制，表示该定时器用于每秒检查调度。
# <翻译结束>


<原文开始>
	// When you add a scheduled task, you cannot allow it to run.
	// It cannot start running when added to timer.
	// It should start running after the entry is added to the Cron entries map, to avoid the task
	// from running during adding where the entries do not have the entry information, which might cause panic.
<原文结束>

# <翻译开始>
// 当你添加一个计划任务时，不能允许它立即运行。
// 它不能在添加到timer时就开始运行。
// 应该在该任务条目被添加到Cron的entries映射中之后才开始运行，以避免在添加过程中由于entries尚未拥有该条目的信息而运行任务，这可能会导致panic。
# <翻译结束>


<原文开始>
// IsSingleton return whether this entry is a singleton timed task.
<原文结束>

# <翻译开始>
// IsSingleton 返回该条目是否为单例定时任务。
# <翻译结束>


<原文开始>
// SetSingleton sets the entry running in singleton mode.
<原文结束>

# <翻译开始>
// SetSingleton 设置入口以单例模式运行。
# <翻译结束>


<原文开始>
// SetTimes sets the times which the entry can run.
<原文结束>

# <翻译开始>
// SetTimes 设置条目可以运行的时间。
# <翻译结束>


<原文开始>
// Status returns the status of entry.
<原文结束>

# <翻译开始>
// Status 返回 entry 的状态。
# <翻译结束>


<原文开始>
// SetStatus sets the status of the entry.
<原文结束>

# <翻译开始>
// SetStatus 设置条目的状态。
# <翻译结束>


<原文开始>
// Start starts running the entry.
<原文结束>

# <翻译开始>
// Start 启动运行入口。
# <翻译结束>


<原文开始>
// Stop stops running the entry.
<原文结束>

# <翻译开始>
// Stop 停止运行 entry。
# <翻译结束>


<原文开始>
// Close stops and removes the entry from cron.
<原文结束>

# <翻译开始>
// Close 停止该任务并从 cron 中移除该条目。
# <翻译结束>


<原文开始>
// checkAndRun is the core timing task check logic.
<原文结束>

# <翻译开始>
// checkAndRun 是核心的定时任务检查逻辑。
# <翻译结束>


<原文开始>
// Exception caught, it logs the error content to logger in default behavior.
<原文结束>

# <翻译开始>
// 捕获到异常，按照默认行为将错误内容记录到日志器中。
# <翻译结束>


<原文开始>
// Entry is timing task entry.
<原文结束>

# <翻译开始>
// Entry 是定时任务的入口。
# <翻译结束>


<原文开始>
// Cron object belonged to.
<原文结束>

# <翻译开始>
// Cron对象所属的。
# <翻译结束>


<原文开始>
// Associated timer Entry.
<原文结束>

# <翻译开始>
// 关联的定时器条目。
# <翻译结束>


<原文开始>
// Timed schedule object.
<原文结束>

# <翻译开始>
// 定时调度对象
# <翻译结束>


<原文开始>
// Running times limit.
<原文结束>

# <翻译开始>
// 运行次数限制。
# <翻译结束>


<原文开始>
// Callback function.
<原文结束>

# <翻译开始>
// 回调函数。
# <翻译结束>


<原文开始>
// The context for the job.
<原文结束>

# <翻译开始>
// 该作业的上下文。
# <翻译结束>


<原文开始>
// Running times check.
<原文结束>

# <翻译开始>
// 运行时间检查。
# <翻译结束>

