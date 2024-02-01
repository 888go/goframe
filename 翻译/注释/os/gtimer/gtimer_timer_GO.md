
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
// New creates and returns a Timer.
<原文结束>

# <翻译开始>
// New 创建并返回一个 Timer。
# <翻译结束>


<原文开始>
// Add adds a timing job to the timer, which runs in interval of `interval`.
<原文结束>

# <翻译开始>
// Add 向定时器添加一个计时任务，该任务以 `interval` 为间隔运行。
# <翻译结束>


<原文开始>
// AddEntry adds a timing job to the timer with detailed parameters.
//
// The parameter `interval` specifies the running interval of the job.
//
// The parameter `singleton` specifies whether the job running in singleton mode.
// There's only one of the same job is allowed running when it's a singleton mode job.
//
// The parameter `times` specifies limit for the job running times, which means the job
// exits if its run times exceeds the `times`.
//
// The parameter `status` specifies the job status when it's firstly added to the timer.
<原文结束>

# <翻译开始>
// AddEntry 向定时器添加一个具有详细参数的定时任务。
//
// 参数 `interval` 指定了该任务的运行间隔时间。
//
// 参数 `singleton` 指定了任务是否以单例模式运行。如果是单例模式的任务，同一任务在运行时只允许存在一个实例。
//
// 参数 `times` 指定了任务运行次数的限制，意味着当任务运行次数超过 `times` 时，该任务将结束运行。
//
// 参数 `status` 指定了任务首次添加到定时器时的状态。
# <翻译结束>


<原文开始>
// AddSingleton is a convenience function for add singleton mode job.
<原文结束>

# <翻译开始>
// AddSingleton 是一个用于添加单例模式任务的便捷函数。
# <翻译结束>


<原文开始>
// AddOnce is a convenience function for adding a job which only runs once and then exits.
<原文结束>

# <翻译开始>
// AddOnce 是一个便捷函数，用于添加一个仅运行一次然后退出的任务。
# <翻译结束>


<原文开始>
// AddTimes is a convenience function for adding a job which is limited running times.
<原文结束>

# <翻译开始>
// AddTimes 是一个便捷函数，用于添加有一定运行次数限制的任务。
# <翻译结束>


<原文开始>
// DelayAdd adds a timing job after delay of `delay` duration.
// Also see Add.
<原文结束>

# <翻译开始>
// DelayAdd 在 `delay` 延迟时间后添加一个定时任务。
// 也可参考 Add。
# <翻译结束>


<原文开始>
// DelayAddEntry adds a timing job after delay of `delay` duration.
// Also see AddEntry.
<原文结束>

# <翻译开始>
// DelayAddEntry 在`delay`延迟时间后添加一个定时任务。
// 也可参考 AddEntry。
# <翻译结束>


<原文开始>
// DelayAddSingleton adds a timing job after delay of `delay` duration.
// Also see AddSingleton.
<原文结束>

# <翻译开始>
// DelayAddSingleton在`delay`延迟时间后添加一个定时任务。
// 也请参阅AddSingleton。
# <翻译结束>


<原文开始>
// DelayAddOnce adds a timing job after delay of `delay` duration.
// Also see AddOnce.
<原文结束>

# <翻译开始>
// DelayAddOnce在`delay`延迟时间之后添加一个定时任务。
// 另请参阅AddOnce。
# <翻译结束>


<原文开始>
// DelayAddTimes adds a timing job after delay of `delay` duration.
// Also see AddTimes.
<原文结束>

# <翻译开始>
// DelayAddTimes 在`delay`延迟时间后添加一个定时任务。
// 另请参阅 AddTimes。
# <翻译结束>

















<原文开始>
// createEntry creates and adds a timing job to the timer.
<原文结束>

# <翻译开始>
// createEntry 创建并添加一个定时任务到计时器中。
# <翻译结束>


<原文开始>
		// If the given interval is lesser than the one of the wheel,
		// then sets it to one tick, which means it will be run in one interval.
<原文结束>

# <翻译开始>
// 如果给定的间隔小于轮子的间隔，
// 则将其设置为一个刻度，这意味着它将在一个间隔后运行。
# <翻译结束>


<原文开始>
		// If the quick mode is enabled, which means it will be run right now.
		// Don't need to wait for the first interval.
<原文结束>

# <翻译开始>
// 如果启用了快速模式，意味着它将立即运行。
// 不需要等待第一个间隔时间。
# <翻译结束>


<原文开始>
// Start starts the timer.
<原文结束>

# <翻译开始>
// Start 开始计时器。
# <翻译结束>


<原文开始>
// Stop stops the timer.
<原文结束>

# <翻译开始>
// Stop 停止定时器。
# <翻译结束>


<原文开始>
// Close closes the timer.
<原文结束>

# <翻译开始>
// Close 关闭计时器。
# <翻译结束>

