
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
// Cron stores all the cron job entries.
<原文结束>

# <翻译开始>
// Cron 存储所有的cron作业项。 md5:3a72b04261d69d0e
# <翻译结束>


<原文开始>
// Used for unique name generation.
<原文结束>

# <翻译开始>
// 用于生成唯一名称。 md5:6a2c8b5e748394aa
# <翻译结束>


<原文开始>
// Timed task status(0: Not Start; 1: Running; 2: Stopped; -1: Closed)
<原文结束>

# <翻译开始>
// 定时任务状态 (0: 未开始; 1: 运行中; 2: 已停止; -1: 已关闭). md5:c07ab3c74e0378b5
# <翻译结束>


<原文开始>
// All timed task entries.
<原文结束>

# <翻译开始>
// 所有定时任务条目。 md5:a74f7440115592e5
# <翻译结束>


<原文开始>
// Logger, it is nil in default.
<原文结束>

# <翻译开始>
// Logger，默认情况下为nil。 md5:49f083354c677555
# <翻译结束>


<原文开始>
// New returns a new Cron object with default settings.
<原文结束>

# <翻译开始>
// New 返回一个具有默认设置的新 Cron 对象。 md5:d0ca0545e6ef9b46
# <翻译结束>


<原文开始>
// SetLogger sets the logger for cron.
<原文结束>

# <翻译开始>
// SetLogger 为 cron 设置日志记录器。 md5:87e9222eac80e2a1
# <翻译结束>


<原文开始>
// GetLogger returns the logger in the cron.
<原文结束>

# <翻译开始>
// GetLogger 在cron中返回日志记录器。 md5:014719534223048e
# <翻译结束>


<原文开始>
// AddEntry creates and returns a new Entry object.
<原文结束>

# <翻译开始>
// AddEntry 创建并返回一个新的 Entry 对象。 md5:1a7d5709c2867437
# <翻译结束>


<原文开始>
// Add adds a timed task.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// Add 添加一个定时任务。
// 一个唯一的`name`可以与定时任务绑定。
// 如果`name`已经被使用，它将返回一个错误。 md5:850ebd654a2e3695
# <翻译结束>


<原文开始>
// AddSingleton adds a singleton timed task.
// A singleton timed task is that can only be running one single instance at the same time.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// AddSingleton 添加一个单例定时任务。
// 单例定时任务是指在同一时间只能运行一个实例的任务。
// 可以通过一个唯一的`name`与定时任务绑定。
// 如果`name`已经被使用，它将返回错误。 md5:9e0e86c2aee09877
# <翻译结束>


<原文开始>
// AddTimes adds a timed task which can be run specified times.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// AddTimes 添加一个定时任务，可以指定执行次数。
// 可以为定时任务绑定一个唯一的`name`。
// 如果`name`已存在，将返回一个错误。 md5:b01e5695f9cc54d5
# <翻译结束>


<原文开始>
// AddOnce adds a timed task which can be run only once.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// AddOnce 添加一个定时任务，该任务只能运行一次。
// 可以为定时任务绑定一个唯一的`name`。
// 如果`name`已使用，将返回一个错误。 md5:fd5fb4f131e1f937
# <翻译结束>


<原文开始>
// DelayAddEntry adds a timed task after `delay` time.
<原文结束>

# <翻译开始>
// DelayAddEntry 在延迟`delay`时间后添加一个定时任务。 md5:a45391b1d2daacd5
# <翻译结束>


<原文开始>
// DelayAdd adds a timed task after `delay` time.
<原文结束>

# <翻译开始>
// DelayAdd 在`delay`时间后添加一个定时任务。 md5:027e39a4b8e3b167
# <翻译结束>


<原文开始>
// DelayAddSingleton adds a singleton timed task after `delay` time.
<原文结束>

# <翻译开始>
// DelayAddSingleton 在`delay`时间后添加一个单例计时任务。 md5:c56847cf6733a3e4
# <翻译结束>


<原文开始>
// DelayAddOnce adds a timed task after `delay` time.
// This timed task can be run only once.
<原文结束>

# <翻译开始>
// DelayAddOnce 在 `delay` 时间后添加一个定时任务。
// 这个定时任务只能运行一次。 md5:34aa0df8fb8e5477
# <翻译结束>


<原文开始>
// DelayAddTimes adds a timed task after `delay` time.
// This timed task can be run specified times.
<原文结束>

# <翻译开始>
// DelayAddTimes 在延迟`delay`时间后添加一个定时任务。
// 该定时任务可以指定运行次数。 md5:5ed58fb7650ed0bb
# <翻译结束>


<原文开始>
// Search returns a scheduled task with the specified `name`.
// It returns nil if not found.
<原文结束>

# <翻译开始>
// Search 返回指定名称的计划任务。
// 如果未找到，则返回 nil。 md5:b0da4b1e1203c6c7
# <翻译结束>


<原文开始>
// Start starts running the specified timed task named `name`.
// If no`name` specified, it starts the entire cron.
<原文结束>

# <翻译开始>
// Start 启动指定名为 `name` 的定时任务。如果没有指定 `name`，则启动整个 cron。 md5:d573745c6d8edaac
# <翻译结束>


<原文开始>
// Stop stops running the specified timed task named `name`.
// If no`name` specified, it stops the entire cron.
<原文结束>

# <翻译开始>
// Stop 停止运行指定的定时任务，任务名为 `name`。如果未指定 `name`，则停止整个cron（cron job）。 md5:68ed27359d633f5e
# <翻译结束>


<原文开始>
// Remove deletes scheduled task which named `name`.
<原文结束>

# <翻译开始>
// Remove 删除名为`name`的计划任务。 md5:bc96b6bdb0bac57b
# <翻译结束>


<原文开始>
// Close stops and closes current cron.
<原文结束>

# <翻译开始>
// Close 停止并关闭当前的cron。 md5:95a4276ef94fb50c
# <翻译结束>


<原文开始>
// Size returns the size of the timed tasks.
<原文结束>

# <翻译开始>
// Size返回定时任务的大小。 md5:a282381f7ca9bf53
# <翻译结束>


<原文开始>
// Entries return all timed tasks as slice(order by registered time asc).
<原文结束>

# <翻译开始>
// Entries 返回所有定时任务作为切片（按注册时间升序排列）。 md5:67b4f559a25d411e
# <翻译结束>

