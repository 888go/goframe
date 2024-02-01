
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
// Package gtimer implements timer for interval/delayed jobs running and management.
//
// This package is designed for management for millions of timing jobs. The differences
// between gtimer and gcron are as follows:
//  1. package gcron is implemented based on package gtimer.
//  2. gtimer is designed for high performance and for millions of timing jobs.
//  3. gcron supports configuration pattern grammar like linux crontab, which is more manually
//     readable.
//  4. gtimer's benchmark OP is measured in nanoseconds, and gcron's benchmark OP is measured
//     in microseconds.
//
// ALSO VERY NOTE the common delay of the timer: https://github.com/golang/go/issues/14410
<原文结束>

# <翻译开始>
// Package gtimer 实现了用于处理和管理间隔/延迟任务的定时器。
//
// 该包设计用于高效管理数百万计的定时任务。gtimer 和 gcron 包之间的区别如下：
//  1. 包 gcron 是基于 gtimer 包实现的。
//  2. gtimer 针对高性能场景设计，适用于处理数百万级别的定时任务。
//  3. gcron 支持类似 Linux crontab 的配置模式语法，更便于人工阅读。
//  4. gtimer 的基准操作（OP）性能以纳秒为单位衡量，而 gcron 的基准操作性能则以微秒为单位衡量。
//
// 另外，请特别注意定时器常见的延迟问题：https://github.com/golang/go/issues/14410
# <翻译结束>


<原文开始>
// Timer is the timer manager, which uses ticks to calculate the timing interval.
<原文结束>

# <翻译开始>
// Timer 是计时器管理器，它利用滴答（ticks）来计算定时间隔。
# <翻译结束>


<原文开始>
// queue is a priority queue based on heap structure.
<原文结束>

# <翻译开始>
// queue 是一个基于堆结构的优先队列。
# <翻译结束>


<原文开始>
// status is the current timer status.
<原文结束>

# <翻译开始>
// status 是当前计时器的状态。
# <翻译结束>


<原文开始>
// ticks is the proceeded interval number by the timer.
<原文结束>

# <翻译开始>
// ticks 是定时器已处理的间隔数。
# <翻译结束>


<原文开始>
// timer options is used for timer configuration.
<原文结束>

# <翻译开始>
// timer options 用于配置定时器。
# <翻译结束>


<原文开始>
// TimerOptions is the configuration object for Timer.
<原文结束>

# <翻译开始>
// TimerOptions 是 Timer 的配置对象。
# <翻译结束>


<原文开始>
// (optional) Interval is the underlying rolling interval tick of the timer.
<原文结束>

# <翻译开始>
// (可选) Interval 是该定时器底层的滚动间隔刻度。
# <翻译结束>


<原文开始>
// Quick is used for quick timer, which means the timer will not wait for the first interval to be elapsed.
<原文结束>

# <翻译开始>
// Quick 用于快速计时器，这意味着计时器不会等待第一个间隔时间过去。
# <翻译结束>


<原文开始>
// internalPanic is the custom panic for internal usage.
<原文结束>

# <翻译开始>
// internalPanic 是用于内部使用的自定义恐慌函数。
# <翻译结束>


<原文开始>
// Job or Timer is ready for running.
<原文结束>

# <翻译开始>
// Job 或 Timer 准备就绪，可以开始运行。
# <翻译结束>


<原文开始>
// Job or Timer is already running.
<原文结束>

# <翻译开始>
// 任务或计时器已经在运行中。
# <翻译结束>







<原文开始>
// Job or Timer is closed and waiting to be deleted.
<原文结束>

# <翻译开始>
// 作业或计时器已关闭，正在等待被删除。
# <翻译结束>


<原文开始>
// panicExit is used for custom job exit with panic.
<原文结束>

# <翻译开始>
// panicExit 用于通过 panic 进行自定义作业退出。
# <翻译结束>


<原文开始>
// defaultTimerInterval is the default timer interval in milliseconds.
<原文结束>

# <翻译开始>
// defaultTimerInterval 是默认计时器间隔，单位为毫秒。
# <翻译结束>


<原文开始>
// commandEnvKeyForInterval is the key for command argument or environment configuring default interval duration for timer.
<原文结束>

# <翻译开始>
// commandEnvKeyForInterval 是用于配置定时器默认间隔持续时间的命令行参数或环境变量的键。
# <翻译结束>


<原文开始>
// DefaultOptions creates and returns a default options object for Timer creation.
<原文结束>

# <翻译开始>
// DefaultOptions 创建并返回一个用于创建Timer的默认选项对象。
# <翻译结束>


<原文开始>
// SetTimeout runs the job once after duration of `delay`.
// It is like the one in javascript.
<原文结束>

# <翻译开始>
// SetTimeout 在 `delay` 延迟时间过后执行一次任务。
// 它类似于 JavaScript 中的 SetTimeout。
# <翻译结束>


<原文开始>
// SetInterval runs the job every duration of `delay`.
// It is like the one in javascript.
<原文结束>

# <翻译开始>
// SetInterval 每隔 `delay` 时间间隔运行 job。
// 它类似于 JavaScript 中的 SetInterval。
# <翻译结束>


<原文开始>
// Add adds a timing job to the default timer, which runs in interval of `interval`.
<原文结束>

# <翻译开始>
// Add 将一个定时任务添加到默认计时器，该计时器以`interval`为间隔运行。
# <翻译结束>


<原文开始>
// AddEntry adds a timing job to the default timer with detailed parameters.
//
// The parameter `interval` specifies the running interval of the job.
//
// The parameter `singleton` specifies whether the job running in singleton mode.
// There's only one of the same job is allowed running when its a singleton mode job.
//
// The parameter `times` specifies limit for the job running times, which means the job
// exits if its run times exceeds the `times`.
//
// The parameter `status` specifies the job status when it's firstly added to the timer.
<原文结束>

# <翻译开始>
// AddEntry 向默认计时器添加一个具有详细参数的定时任务。
//
// 参数 `interval` 指定了任务的运行间隔时间。
//
// 参数 `singleton` 指定了任务是否以单例模式运行。如果任务是单例模式，那么同一任务在运行时只允许有一个实例。
//
// 参数 `times` 指定了任务运行次数的限制，这意味着当任务运行次数超过 `times` 时，该任务将退出。
//
// 参数 `status` 指定了任务首次添加到计时器时的状态。
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
// DelayAdd adds a timing job after delay of `interval` duration.
// Also see Add.
<原文结束>

# <翻译开始>
// DelayAdd 在 `interval` 延迟时间后添加一个定时任务。
// 另请参阅 Add。
# <翻译结束>


<原文开始>
// DelayAddEntry adds a timing job after delay of `interval` duration.
// Also see AddEntry.
<原文结束>

# <翻译开始>
// DelayAddEntry 在 `interval` 延迟后添加一个定时任务。
// 也可参考 AddEntry。
# <翻译结束>


<原文开始>
// DelayAddSingleton adds a timing job after delay of `interval` duration.
// Also see AddSingleton.
<原文结束>

# <翻译开始>
// DelayAddSingleton 在`interval`延迟时间后添加一个定时任务。
// 另请参阅 AddSingleton。
# <翻译结束>


<原文开始>
// DelayAddOnce adds a timing job after delay of `interval` duration.
// Also see AddOnce.
<原文结束>

# <翻译开始>
// DelayAddOnce在`interval`延迟时间后添加一个定时任务。
// 另请参阅AddOnce。
# <翻译结束>


<原文开始>
// DelayAddTimes adds a timing job after delay of `interval` duration.
// Also see AddTimes.
<原文结束>

# <翻译开始>
// DelayAddTimes 在 `interval` 延迟后添加一个定时任务。
// 另请参阅 AddTimes。
# <翻译结束>


<原文开始>
// Job or Timer is stopped.
<原文结束>

# <翻译开始>
// Job 或 Timer 已停止。
# <翻译结束>

