
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
// Package gcron implements a cron pattern parser and job runner.
<原文结束>

# <翻译开始>
// gcron 包实现了cron模式解析器和任务运行器。 md5:56d461ee2c75e1f7
# <翻译结束>


<原文开始>
// SetLogger sets the global logger for cron.
<原文结束>

# <翻译开始>
// SetLogger 设置 cron 的全局日志记录器。 md5:b0a9c7514a8c8b29
# <翻译结束>


<原文开始>
// GetLogger returns the global logger in the cron.
<原文结束>

# <翻译开始>
// GetLogger 返回 cron 中的全局日志记录器。 md5:908af8c408f25d3d
# <翻译结束>


<原文开始>
// Add adds a timed task to default cron object.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// Add 向默认的 cron 对象添加一个定时任务。
// 一个唯一的 `name` 可以与定时任务绑定。
// 如果 `name` 已经被使用，它将返回一个错误。
// md5:0f46b08a5a96144a
# <翻译结束>


<原文开始>
// AddSingleton adds a singleton timed task, to default cron object.
// A singleton timed task is that can only be running one single instance at the same time.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// AddSingleton 向默认的 cron 对象添加一个单例定时任务。
// 单例定时任务是指在同一时间只能运行一个实例的任务。
// 可以将一个唯一的 `name` 与定时任务绑定。
// 如果 `name` 已经被使用，它将返回错误。
// md5:96268d3e8373359e
# <翻译结束>


<原文开始>
// AddOnce adds a timed task which can be run only once, to default cron object.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// AddOnce 向默认cron对象添加一个仅能运行一次的定时任务。可以为定时任务绑定一个唯一的`name`。
// 如果`name`已存在，它将返回一个错误。
// md5:9701998ce952852d
# <翻译结束>


<原文开始>
// AddTimes adds a timed task which can be run specified times, to default cron object.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
<原文结束>

# <翻译开始>
// AddTimes 将一个定时任务添加到默认cron对象中，该任务可以执行指定次数。可以为定时任务绑定一个唯一的`name`。
// 如果`name`已经存在，它将返回一个错误。
// md5:505a2bbf10529705
# <翻译结束>


<原文开始>
// DelayAdd adds a timed task to default cron object after `delay` time.
<原文结束>

# <翻译开始>
// DelayAdd 在延迟'delay'时间后，向默认的cron对象添加一个定时任务。 md5:7c28aadbf64d1362
# <翻译结束>


<原文开始>
// DelayAddSingleton adds a singleton timed task after `delay` time to default cron object.
<原文结束>

# <翻译开始>
// DelayAddSingleton 在默认的cron对象中延迟`delay`时间后添加一个单例定时任务。 md5:d6c26c1edc16c19a
# <翻译结束>


<原文开始>
// DelayAddOnce adds a timed task after `delay` time to default cron object.
// This timed task can be run only once.
<原文结束>

# <翻译开始>
// DelayAddOnce 在 `delay` 时间后向默认cron对象添加一个定时任务。
// 这个定时任务只能运行一次。
// md5:56e7e748e4d4c63d
# <翻译结束>


<原文开始>
// DelayAddTimes adds a timed task after `delay` time to default cron object.
// This timed task can be run specified times.
<原文结束>

# <翻译开始>
// DelayAddTimes 在`delay`时间后向默认cron对象添加一个定时任务。这个定时任务可以运行指定次数。
// md5:b4ecf68ee4e86408
# <翻译结束>


<原文开始>
// Search returns a scheduled task with the specified `name`.
// It returns nil if no found.
<原文结束>

# <翻译开始>
// Search 返回具有指定`name`的计划任务。
// 如果未找到，返回nil。
// md5:06da84fc0766d888
# <翻译结束>


<原文开始>
// Remove deletes scheduled task which named `name`.
<原文结束>

# <翻译开始>
// Remove 删除名为`name`的计划任务。 md5:bc96b6bdb0bac57b
# <翻译结束>


<原文开始>
// Size returns the size of the timed tasks of default cron.
<原文结束>

# <翻译开始>
// Size 返回默认cron中定时任务的数量。 md5:e87d680e31031739
# <翻译结束>


<原文开始>
// Entries return all timed tasks as slice.
<原文结束>

# <翻译开始>
// Entries 返回所有定时任务作为切片。 md5:16823f1ebafbd9cc
# <翻译结束>


<原文开始>
// Start starts running the specified timed task named `name`.
// If no`name` specified, it starts the entire cron.
<原文结束>

# <翻译开始>
// Start 启动指定名为 `name` 的定时任务。如果没有指定 `name`，则启动整个 cron。
// md5:d573745c6d8edaac
# <翻译结束>


<原文开始>
// Stop stops running the specified timed task named `name`.
// If no`name` specified, it stops the entire cron.
<原文结束>

# <翻译开始>
// Stop 停止运行指定的定时任务，任务名为 `name`。如果未指定 `name`，则停止整个cron（cron job）。
// md5:68ed27359d633f5e
# <翻译结束>

