
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
// New creates and returns a Timer.
<原文结束>

# <翻译开始>
// New 创建并返回一个计时器。 md5:3db3235abce3957a
# <翻译结束>


<原文开始>
// Add adds a timing job to the timer, which runs in interval of `interval`.
<原文结束>

# <翻译开始>
// Add 向定时器添加一个定时任务，该任务以 `interval` 为间隔运行。 md5:358decaca6313fa2
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
// AddEntry 向计时器添加一个具有详细参数的定时任务。
//
// 参数 `interval` 指定任务的运行间隔。
//
// 参数 `singleton` 指定任务是否以单例模式运行。当任务为单例模式时，只允许运行一个相同的任务。
//
// 参数 `times` 指定任务的最大运行次数，即如果任务的运行次数超过 `times`，则退出。
//
// 参数 `status` 指定任务初次添加到计时器时的状态。
// md5:22c21ed21d95479c
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
// DelayAdd adds a timing job after delay of `delay` duration.
// Also see Add.
<原文结束>

# <翻译开始>
// DelayAdd 在延迟`delay`持续时间后添加一个定时任务。
// 参阅 Add。
// md5:5db03c917e923b89
# <翻译结束>


<原文开始>
// DelayAddEntry adds a timing job after delay of `delay` duration.
// Also see AddEntry.
<原文结束>

# <翻译开始>
// DelayAddEntry 在延迟`delay`持续时间后，添加一个定时任务。
// 参见AddEntry。
// md5:6f230211b81dca10
# <翻译结束>


<原文开始>
// DelayAddSingleton adds a timing job after delay of `delay` duration.
// Also see AddSingleton.
<原文结束>

# <翻译开始>
// DelayAddSingleton 在延迟`delay`时间后添加一个定时任务。
// 另请参阅AddSingleton。
// md5:4df66d0755ab0371
# <翻译结束>


<原文开始>
// DelayAddOnce adds a timing job after delay of `delay` duration.
// Also see AddOnce.
<原文结束>

# <翻译开始>
// DelayAddOnce 在延迟 `delay` 持续时间后添加一个定时任务。
// 参见 AddOnce。
// md5:613532ca791628bf
# <翻译结束>


<原文开始>
// DelayAddTimes adds a timing job after delay of `delay` duration.
// Also see AddTimes.
<原文结束>

# <翻译开始>
// DelayAddTimes 在`delay`持续时间后添加一个定时任务。
// 参阅 AddTimes。
// md5:8a53452ea0989047
# <翻译结束>


<原文开始>
// Start starts the timer.
<原文结束>

# <翻译开始>
// Start 开始计时器。 md5:3330d3e996e68d8f
# <翻译结束>


<原文开始>
// Close closes the timer.
<原文结束>

# <翻译开始>
// Close 关闭定时器。 md5:cc2ed98b62fc4904
# <翻译结束>


<原文开始>
// createEntry creates and adds a timing job to the timer.
<原文结束>

# <翻译开始>
// createEntry 创建并添加一个计时任务到计时器中。 md5:8d0d62888bb8b536
# <翻译结束>


<原文开始>
		// If the given interval is lesser than the one of the wheel,
		// then sets it to one tick, which means it will be run in one interval.
<原文结束>

# <翻译开始>
// 如果给定的间隔小于车轮的间隔，那么将其设置为一个刻度，这意味着它将在一个间隔内运行。
// md5:4f1ce7e56c107c6e
# <翻译结束>


<原文开始>
		// If the quick mode is enabled, which means it will be run right now.
		// Don't need to wait for the first interval.
<原文结束>

# <翻译开始>
// 如果启用了快速模式，意味着它将立即执行。
// 无需等待第一个时间间隔。
// md5:6d9ecc987797b1ba
# <翻译结束>

