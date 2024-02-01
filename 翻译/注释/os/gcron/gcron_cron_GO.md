
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
// Used for unique name generation.
<原文结束>

# <翻译开始>
// 用于生成唯一名称。
# <翻译结束>


<原文开始>
// Timed task status(0: Not Start; 1: Running; 2: Stopped; -1: Closed)
<原文结束>

# <翻译开始>
// 定时任务状态(0: 未开始; 1: 运行中; 2: 已停止; -1: 关闭)
# <翻译结束>


<原文开始>
// Logger, it is nil in default.
<原文结束>

# <翻译开始>
// Logger，默认情况下为nil。
# <翻译结束>


<原文开始>
// New returns a new Cron object with default settings.
<原文结束>

# <翻译开始>
// New 函数返回一个使用默认设置的新 Cron 对象。
# <翻译结束>


<原文开始>
// SetLogger sets the logger for cron.
<原文结束>

# <翻译开始>
// SetLogger 用于设置cron的日志记录器。
# <翻译结束>


<原文开始>
// GetLogger returns the logger in the cron.
<原文结束>

# <翻译开始>
// GetLogger 返回cron中的日志器。
# <翻译结束>


<原文开始>
// AddEntry creates and returns a new Entry object.
<原文结束>

# <翻译开始>
// AddEntry 创建并返回一个新的 Entry 对象。
# <翻译结束>


<原文开始>
// Add adds a timed task.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// Add 添加一个定时任务。
// 可以用唯一的 `name` 与定时任务进行绑定。
// 如果 `name` 已经被使用，则返回错误。
# <翻译结束>


<原文开始>
// AddSingleton adds a singleton timed task.
// A singleton timed task is that can only be running one single instance at the same time.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// AddSingleton 添加一个单例定时任务。
// 单例定时任务是指在同一时刻只能运行一个实例的任务。
// 可以使用唯一的 `name` 与定时任务绑定。
// 如果 `name` 已被使用，则返回错误。
# <翻译结束>


<原文开始>
// AddTimes adds a timed task which can be run specified times.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// AddTimes 添加一个可运行指定次数的定时任务。
// 可以使用唯一的 `name` 与定时任务关联绑定。
// 如果 `name` 已被使用，则返回错误。
# <翻译结束>


<原文开始>
// AddOnce adds a timed task which can be run only once.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// AddOnce 添加一个仅能运行一次的定时任务。
// 可以为定时任务绑定一个唯一的 `name`。
// 如果 `name` 已被使用，则返回错误。
# <翻译结束>


<原文开始>
// DelayAddEntry adds a timed task after `delay` time.
<原文结束>

# <翻译开始>
// DelayAddEntry 在 `delay` 时间后添加一个定时任务。
# <翻译结束>


<原文开始>
// DelayAdd adds a timed task after `delay` time.
<原文结束>

# <翻译开始>
// DelayAdd 在 `delay` 时间后添加一个定时任务。
# <翻译结束>


<原文开始>
// DelayAddSingleton adds a singleton timed task after `delay` time.
<原文结束>

# <翻译开始>
// DelayAddSingleton 在`delay`时间后添加一个单例定时任务。
# <翻译结束>


<原文开始>
// DelayAddOnce adds a timed task after `delay` time.
// This timed task can be run only once.
<原文结束>

# <翻译开始>
// DelayAddOnce在`delay`时间后添加一个定时任务。
// 这个定时任务只能运行一次。
# <翻译结束>


<原文开始>
// DelayAddTimes adds a timed task after `delay` time.
// This timed task can be run specified times.
<原文结束>

# <翻译开始>
// DelayAddTimes 在 `delay` 时间后添加一个定时任务。
// 此定时任务可以运行指定次数。
# <翻译结束>


<原文开始>
// Search returns a scheduled task with the specified `name`.
// It returns nil if not found.
<原文结束>

# <翻译开始>
// Search 返回具有指定`name`的已计划任务。
// 如果未找到，则返回nil。
# <翻译结束>


<原文开始>
// Start starts running the specified timed task named `name`.
// If no`name` specified, it starts the entire cron.
<原文结束>

# <翻译开始>
// Start 开始运行指定的定时任务，名为 `name`。
// 如果未指定 `name`，则启动整个 cron。
# <翻译结束>


<原文开始>
// Stop stops running the specified timed task named `name`.
// If no`name` specified, it stops the entire cron.
<原文结束>

# <翻译开始>
// Stop 停止运行指定的定时任务，该任务名为 `name`。
// 如果未指定 `name`，则停止整个 cron。
# <翻译结束>


<原文开始>
// Remove deletes scheduled task which named `name`.
<原文结束>

# <翻译开始>
// Remove 删除名为 `name` 的已计划任务。
# <翻译结束>


<原文开始>
// Close stops and closes current cron.
<原文结束>

# <翻译开始>
// Close 停止并关闭当前的cron任务。
# <翻译结束>


<原文开始>
// Size returns the size of the timed tasks.
<原文结束>

# <翻译开始>
// Size 返回定时任务的数量。
# <翻译结束>


<原文开始>
// Entries return all timed tasks as slice(order by registered time asc).
<原文结束>

# <翻译开始>
// Entries 返回所有按注册时间升序排列的定时任务切片。
# <翻译结束>


<原文开始>
// All timed task entries.
<原文结束>

# <翻译开始>
// 所有定时任务条目。
# <翻译结束>

