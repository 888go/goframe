
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
// Package gfsnotify provides a platform-independent interface for file system notifications.
<原文结束>

# <翻译开始>
// 包gfsnotify提供了一个与平台无关的接口，用于文件系统通知。 md5:85a6a9e7b52e09e5
# <翻译结束>


<原文开始>
// Watcher is the monitor for file changes.
<原文结束>

# <翻译开始>
// Watcher是文件更改的监视器。 md5:82c8c6038aefc543
# <翻译结束>


<原文开始>
// Underlying fsnotify object.
<原文结束>

# <翻译开始>
// 基础的fsnotify对象。 md5:33936d13041bf0ad
# <翻译结束>


<原文开始>
// Used for internal event management.
<原文结束>

# <翻译开始>
// 用于内部事件管理。 md5:3b850c6c87220b73
# <翻译结束>


<原文开始>
// Used for repeated event filter.
<原文结束>

# <翻译开始>
// 用于重复事件过滤。 md5:bc4edc625e5962e4
# <翻译结束>


<原文开始>
// Used for AddOnce feature.
<原文结束>

# <翻译开始>
// 用于添加一次功能。 md5:840c98179b4051aa
# <翻译结束>


<原文开始>
// Path(file/folder) to callbacks mapping.
<原文结束>

# <翻译开始>
// 回调函数的路径（文件/文件夹）映射。 md5:ffdb0824b6a6c3e3
# <翻译结束>


<原文开始>
// Used for watcher closing notification.
<原文结束>

# <翻译开始>
// 用于监视器关闭的通知。 md5:bd08433f373aee64
# <翻译结束>


<原文开始>
// Callback is the callback function for Watcher.
<原文结束>

# <翻译开始>
// Callback是Watcher的回调函数。 md5:fa67ea2e048fd039
# <翻译结束>


<原文开始>
// Unique id for callback object.
<原文结束>

# <翻译开始>
// 回调对象的唯一标识符。 md5:5e0b273b79bf867f
# <翻译结束>


<原文开始>
// Bound file path (absolute).
<原文结束>

# <翻译开始>
// 限定的文件路径（绝对）。 md5:300376a734dfb30f
# <翻译结束>


<原文开始>
// Registered name for AddOnce.
<原文结束>

# <翻译开始>
// AddOnce的注册名称。 md5:eb134382190fde28
# <翻译结束>


<原文开始>
// Element in the callbacks of watcher.
<原文结束>

# <翻译开始>
// 观察者回调中的元素。 md5:abe4631c343e0396
# <翻译结束>


<原文开始>
// Is bound to path recursively or not.
<原文结束>

# <翻译开始>
// 是否递归地绑定到路径。 md5:b5e23c5fc9be130a
# <翻译结束>


<原文开始>
// Event is the event produced by underlying fsnotify.
<原文结束>

# <翻译开始>
// Event 是底层 fsnotify 产生的事件。 md5:cf38e0981dbdfa84
# <翻译结束>


<原文开始>
// Op is the bits union for file operations.
<原文结束>

# <翻译开始>
// Op 是文件操作的位标志联合。 md5:b9c63ca71272a7d3
# <翻译结束>


<原文开始>
// internalPanic is the custom panic for internal usage.
<原文结束>

# <翻译开始>
// internalPanic 是用于内部使用的自定义恐慌。 md5:287806e552654f1d
# <翻译结束>


<原文开始>
// Duration for repeated event filter.
<原文结束>

# <翻译开始>
// 重复事件过滤器的持续时间。 md5:2f6fec9f3496777a
# <翻译结束>


<原文开始>
// Custom exit event for internal usage.
<原文结束>

# <翻译开始>
// 用于内部使用的自定义退出事件。 md5:7c86231249e45c6a
# <翻译结束>


<原文开始>
// Mutex for concurrent safety of defaultWatcher.
<原文结束>

# <翻译开始>
// 用于并发安全的defaultWatcher的Mutex。 md5:0b64dbe940db78a8
# <翻译结束>


<原文开始>
// Id to callback mapping.
<原文结束>

# <翻译开始>
// ID到回调函数的映射。 md5:641a269397a47639
# <翻译结束>


<原文开始>
// Atomic id generator for callback.
<原文结束>

# <翻译开始>
// 用于回调的原子性ID生成器。 md5:2caf00d0d805af7b
# <翻译结束>


<原文开始>
// New creates and returns a new watcher.
// Note that the watcher number is limited by the file handle setting of the system.
// Eg: fs.inotify.max_user_instances system variable in linux systems.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的观察者。
// 注意，观察者的数量受系统文件描述符限制。
// 例如：在 Linux 系统中，fs.inotify.max_user_instances 系统变量。 md5:a2587b1623329074
# <翻译结束>


<原文开始>
// Add monitors `path` using default watcher with callback function `callbackFunc`.
// The optional parameter `recursive` specifies whether monitoring the `path` recursively, which is true in default.
<原文结束>

# <翻译开始>
// 使用默认的观察者(`watcher`)监控路径`path`，并调用回调函数`callbackFunc`。
// 可选参数`recursive`指定是否递归地监控路径`path`，默认为true。 md5:e660326b83136bd1
# <翻译结束>


<原文开始>
// AddOnce monitors `path` using default watcher with callback function `callbackFunc` only once using unique name `name`.
// If AddOnce is called multiple times with the same `name` parameter, `path` is only added to monitor once. It returns error
// if it's called twice with the same `name`.
//
// The optional parameter `recursive` specifies whether monitoring the `path` recursively, which is true in default.
<原文结束>

# <翻译开始>
// AddOnce 使用唯一名称 `name` 及回调函数 `callbackFunc`，仅使用默认监视器监控 `path` 一次。
// 如果多次调用 AddOnce 并传入相同的 `name` 参数，`path` 仅会被添加监控一次。如果使用相同的 `name` 调用两次，它将返回错误。
//
// 可选参数 `recursive` 指定是否递归监控 `path`，默认为 true。 md5:c28c83d5a2230d07
# <翻译结束>


<原文开始>
// Remove removes all monitoring callbacks of given `path` from watcher recursively.
<原文结束>

# <翻译开始>
// Remove 递归地从监视器中删除给定`path`的所有监控回调。 md5:63888786f53ffca5
# <翻译结束>


<原文开始>
// RemoveCallback removes specified callback with given id from watcher.
<原文结束>

# <翻译开始>
// RemoveCallback 从观察者中移除具有给定ID的指定回调。 md5:af906f3547f93046
# <翻译结束>


<原文开始>
// Exit is only used in the callback function, which can be used to remove current callback
// of itself from the watcher.
<原文结束>

# <翻译开始>
// Exit 只在回调函数中使用，可以用于从观察者中移除当前的回调。 md5:697f4cd00adc082e
# <翻译结束>


<原文开始>
// getDefaultWatcher creates and returns the default watcher.
// This is used for lazy initialization purpose.
<原文结束>

# <翻译开始>
// getDefaultWatcher 创建并返回默认的监视器。
// 这用于惰性初始化的目的。 md5:c1a7b4f4102130c0
# <翻译结束>

