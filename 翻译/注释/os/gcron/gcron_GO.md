
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
// Package gcron implements a cron pattern parser and job runner.
<原文结束>

# <翻译开始>
// 包gcron实现了cron模式解析器和任务执行器。
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
// Add adds a timed task to default cron object.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// Add 向默认的 cron 对象添加一个定时任务。
// 可以用唯一的 `name` 与定时任务关联。
// 如果 `name` 已经被使用，则返回错误。
# <翻译结束>


<原文开始>
// AddSingleton adds a singleton timed task, to default cron object.
// A singleton timed task is that can only be running one single instance at the same time.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// AddSingleton 添加单例定时任务到默认的cron对象。
// 单例定时任务是指在同一时间只能运行一个实例的任务。
// 可以通过唯一的`name`与定时任务关联绑定。
// 如果`name`已经被使用，则返回错误。
# <翻译结束>


<原文开始>
// AddOnce adds a timed task which can be run only once, to default cron object.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// AddOnce 向默认的cron对象添加一个仅能执行一次的定时任务。
// 可以为定时任务绑定一个唯一的`name`标识。
// 如果`name`已被使用，则返回错误。
# <翻译结束>


<原文开始>
// AddTimes adds a timed task which can be run specified times, to default cron object.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// AddTimes 向默认的 cron 对象添加一个可运行指定次数的定时任务。
// 可以为定时任务绑定一个唯一的 `name`。
// 如果 `name` 已经被使用，将会返回错误。
# <翻译结束>


<原文开始>
// DelayAdd adds a timed task to default cron object after `delay` time.
<原文结束>

# <翻译开始>
// DelayAdd 在 `delay` 时间后向默认的 cron 对象添加一个定时任务。
# <翻译结束>


<原文开始>
// DelayAddSingleton adds a singleton timed task after `delay` time to default cron object.
<原文结束>

# <翻译开始>
// DelayAddSingleton 在`delay`时间后向默认的cron对象添加一个单例定时任务。
# <翻译结束>


<原文开始>
// DelayAddOnce adds a timed task after `delay` time to default cron object.
// This timed task can be run only once.
<原文结束>

# <翻译开始>
// DelayAddOnce在`delay`时间后向默认的cron对象添加一个定时任务。
// 这个定时任务只能运行一次。
# <翻译结束>


<原文开始>
// DelayAddTimes adds a timed task after `delay` time to default cron object.
// This timed task can be run specified times.
<原文结束>

# <翻译开始>
// DelayAddTimes 在`delay`时间后向默认的cron对象添加一个定时任务。
// 这个定时任务可以运行指定次数。
# <翻译结束>


<原文开始>
// Search returns a scheduled task with the specified `name`.
// It returns nil if no found.
<原文结束>

# <翻译开始>
// Search 返回具有指定名称 `name` 的已计划任务。
// 如果未找到，则返回 nil。
# <翻译结束>


<原文开始>
// Remove deletes scheduled task which named `name`.
<原文结束>

# <翻译开始>
// Remove 删除名为 `name` 的已计划任务。
# <翻译结束>


<原文开始>
// Size returns the size of the timed tasks of default cron.
<原文结束>

# <翻译开始>
// Size 返回默认Cron定时任务的数量。
# <翻译结束>


<原文开始>
// Entries return all timed tasks as slice.
<原文结束>

# <翻译开始>
// Entries 返回所有定时任务作为一个切片。
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
// Default cron object.
<原文结束>

# <翻译开始>
// 默认的cron对象。
# <翻译结束>

