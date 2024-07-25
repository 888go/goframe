
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
// gtimer 包实现了定时/延迟任务的执行与管理。
//
// 本包旨在管理数百万级别的定时任务。gtimer 与 gcron 的区别在于：
// 1. gcron 包是基于 gtimer 包实现的。
// 2. gtimer 侧重于高性能，适用于处理百万级的定时任务。
// 3. gcron 支持类似 Linux crontab 的配置模式语法，更加便于人工阅读。
// 4. gtimer 的基准操作时间单位为纳秒，而 gcron 的基准操作时间单位为微秒。
//
// 另外，特别注意定时器常见的延迟问题：https://github.com/golang/go/issues/14410 md5:c0dc9837a603eb26
# <翻译结束>


<原文开始>
// Timer is the timer manager, which uses ticks to calculate the timing interval.
<原文结束>

# <翻译开始>
// Timer 是计时器管理器，它使用ticks来计算时间间隔。 md5:eab16bf4737f12a8
# <翻译结束>


<原文开始>
// queue is a priority queue based on heap structure.
<原文结束>

# <翻译开始>
// queue 是基于堆结构的优先队列。 md5:6393b80eaafadb06
# <翻译结束>


<原文开始>
// status is the current timer status.
<原文结束>

# <翻译开始>
// status是当前计时器的状态。 md5:86cd8448a559a7ec
# <翻译结束>


<原文开始>
// ticks is the proceeded interval number by the timer.
<原文结束>

# <翻译开始>
// ticks 是计时器经过的间隔数量。 md5:6162dcb7888ba5fd
# <翻译结束>


<原文开始>
// timer options is used for timer configuration.
<原文结束>

# <翻译开始>
// timer options 用于配置定时器的选项。 md5:a922957501fcab8b
# <翻译结束>


<原文开始>
// TimerOptions is the configuration object for Timer.
<原文结束>

# <翻译开始>
// TimerOptions是Timer的配置对象。 md5:9d2942910cd72ea4
# <翻译结束>


<原文开始>
// (optional) Interval is the underlying rolling interval tick of the timer.
<原文结束>

# <翻译开始>
// (可选) Interval是计时器的基础滚动间隔刻度。 md5:4c933d7fc9ae2121
# <翻译结束>


<原文开始>
// Quick is used for quick timer, which means the timer will not wait for the first interval to be elapsed.
<原文结束>

# <翻译开始>
// Quick 用于快速计时器，这意味着计时器不会等待第一个间隔结束便立即触发。 md5:7b719aee1cbeb308
# <翻译结束>


<原文开始>
// internalPanic is the custom panic for internal usage.
<原文结束>

# <翻译开始>
// internalPanic 是用于内部使用的自定义恐慌。 md5:287806e552654f1d
# <翻译结束>


<原文开始>
// Job or Timer is ready for running.
<原文结束>

# <翻译开始>
// 作业或定时器已准备好运行。 md5:043261776a379433
# <翻译结束>


<原文开始>
// Job or Timer is already running.
<原文结束>

# <翻译开始>
// 已经有任务或定时器在运行。 md5:148ed9b5d89215fe
# <翻译结束>


<原文开始>
// Job or Timer is stopped.
<原文结束>

# <翻译开始>
// 任务或定时器已停止。 md5:3cc8479a2061f7db
# <翻译结束>


<原文开始>
// Job or Timer is closed and waiting to be deleted.
<原文结束>

# <翻译开始>
// 作业或计时器已关闭，正在等待被删除。 md5:9b775a18fd7fe5b9
# <翻译结束>


<原文开始>
// panicExit is used for custom job exit with panic.
<原文结束>

# <翻译开始>
// panicExit 用于在发生恐慌时自定义作业退出。 md5:b22834505d9c74ec
# <翻译结束>


<原文开始>
// defaultTimerInterval is the default timer interval in milliseconds.
<原文结束>

# <翻译开始>
// defaultTimerInterval 是默认的定时器间隔（以毫秒为单位）。 md5:8b2b5568f4526000
# <翻译结束>


<原文开始>
// commandEnvKeyForInterval is the key for command argument or environment configuring default interval duration for timer.
<原文结束>

# <翻译开始>
	// commandEnvKeyForInterval 是用于命令参数或环境配置中，定时器的默认间隔持续时间的键。 md5:baf94b4095117907
# <翻译结束>


<原文开始>
// DefaultOptions creates and returns a default options object for Timer creation.
<原文结束>

# <翻译开始>
// DefaultOptions 创建并返回一个用于创建计时器的默认选项对象。 md5:67460fb8c6e56969
# <翻译结束>


<原文开始>
// SetTimeout runs the job once after duration of `delay`.
// It is like the one in javascript.
<原文结束>

# <翻译开始>
// SetTimeout 在`delay`时间间隔后执行一次该任务。
// 它的作用类似于JavaScript中的同名函数。 md5:e0477460fecac4d8
# <翻译结束>


<原文开始>
// SetInterval runs the job every duration of `delay`.
// It is like the one in javascript.
<原文结束>

# <翻译开始>
// SetInterval 每隔 `delay` 的时间运行一次任务。
// 类似于 JavaScript 中的 SetInterval。 md5:1a47e33b3567d415
# <翻译结束>


<原文开始>
// Add adds a timing job to the default timer, which runs in interval of `interval`.
<原文结束>

# <翻译开始>
// Add 将一个计时任务添加到默认计时器中，该计时器以`interval`的间隔运行。 md5:222cab00f4afd790
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
// AddEntry 向默认计时器添加一个带有详细参数的定时任务。
//
// 参数 `interval` 指定任务的运行间隔。
//
// 参数 `singleton` 指定任务是否以单例模式运行。
// 当任务为单例模式时，同一任务只允许有一个实例在运行。
//
// 参数 `times` 限制了任务的运行次数，意味着如果任务运行次数超过 `times` 就会退出。
//
// 参数 `status` 指定任务首次添加到计时器时的状态。 md5:0f65b7fd26b5f483
# <翻译结束>


<原文开始>
// AddSingleton is a convenience function for add singleton mode job.
<原文结束>

# <翻译开始>
// AddSingleton 是一个方便的函数，用于添加单例模式任务。 md5:8530c92e0f700eff
# <翻译结束>


<原文开始>
// AddOnce is a convenience function for adding a job which only runs once and then exits.
<原文结束>

# <翻译开始>
// AddOnce 是一个方便函数，用于添加一个只运行一次然后退出的工作。 md5:7674bfdda1236b76
# <翻译结束>


<原文开始>
// AddTimes is a convenience function for adding a job which is limited running times.
<原文结束>

# <翻译开始>
// AddTimes 是一个方便的函数，用于添加一个有限运行次数的任务。 md5:98b9f39d64b7906c
# <翻译结束>


<原文开始>
// DelayAdd adds a timing job after delay of `interval` duration.
// Also see Add.
<原文结束>

# <翻译开始>
// DelayAdd 在延迟`interval`持续时间后，添加一个定时任务。
// 参见Add。 md5:e1bb93aeff16693d
# <翻译结束>


<原文开始>
// DelayAddEntry adds a timing job after delay of `interval` duration.
// Also see AddEntry.
<原文结束>

# <翻译开始>
// DelayAddEntry 在 `interval` 时长后添加一个定时任务。
// 另请参阅 AddEntry。 md5:e6e85b15472aaf98
# <翻译结束>


<原文开始>
// DelayAddSingleton adds a timing job after delay of `interval` duration.
// Also see AddSingleton.
<原文结束>

# <翻译开始>
// DelayAddSingleton 在延迟`interval`时长后添加一个定时任务。
// 参阅 AddSingleton。 md5:80f70090fa17a370
# <翻译结束>


<原文开始>
// DelayAddOnce adds a timing job after delay of `interval` duration.
// Also see AddOnce.
<原文结束>

# <翻译开始>
// DelayAddOnce 在延迟 `interval` 持续时间后添加一个定时任务。
// 参见 AddOnce。 md5:71c388c8096e5e48
# <翻译结束>


<原文开始>
// DelayAddTimes adds a timing job after delay of `interval` duration.
// Also see AddTimes.
<原文结束>

# <翻译开始>
// DelayAddTimes 在延迟`interval`持续时间后，添加一个定时任务。
// 参见AddTimes。 md5:62463bf6c56709b2
# <翻译结束>

