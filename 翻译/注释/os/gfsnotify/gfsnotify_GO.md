
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
// Package gfsnotify provides a platform-independent interface for file system notifications.
<原文结束>

# <翻译开始>
// Package gfsnotify 提供了一个跨平台的文件系统通知接口。
# <翻译结束>


<原文开始>
// Watcher is the monitor for file changes.
<原文结束>

# <翻译开始>
// Watcher 是用于文件变化监测的监视器。
# <翻译结束>


<原文开始>
// Underlying fsnotify object.
<原文结束>

# <翻译开始>
// 底层 fsnotify 对象。
# <翻译结束>


<原文开始>
// Used for internal event management.
<原文结束>

# <翻译开始>
// 用于内部事件管理。
# <翻译结束>


<原文开始>
// Used for repeated event filter.
<原文结束>

# <翻译开始>
// 用于重复事件过滤。
# <翻译结束>







<原文开始>
// Path(file/folder) to callbacks mapping.
<原文结束>

# <翻译开始>
// Path(文件/文件夹)到回调函数的映射。
# <翻译结束>


<原文开始>
// Used for watcher closing notification.
<原文结束>

# <翻译开始>
// 用于通知观察者关闭。
# <翻译结束>


<原文开始>
// Callback is the callback function for Watcher.
<原文结束>

# <翻译开始>
// Callback是Watcher的回调函数。
# <翻译结束>


<原文开始>
// Unique id for callback object.
<原文结束>

# <翻译开始>
// 唯一标识回调对象的ID。
# <翻译结束>







<原文开始>
// Bound file path (absolute).
<原文结束>

# <翻译开始>
// 绑定文件路径（绝对路径）。
# <翻译结束>


<原文开始>
// Registered name for AddOnce.
<原文结束>

# <翻译开始>
// AddOnce的注册名称。
# <翻译结束>


<原文开始>
// Element in the callbacks of watcher.
<原文结束>

# <翻译开始>
// watcher回调中的元素。
# <翻译结束>


<原文开始>
// Is bound to path recursively or not.
<原文结束>

# <翻译开始>
// 是否递归绑定到路径
# <翻译结束>


<原文开始>
// Event is the event produced by underlying fsnotify.
<原文结束>

# <翻译开始>
// Event 是由底层 fsnotify 产生的事件。
# <翻译结束>







<原文开始>
// Op is the bits union for file operations.
<原文结束>

# <翻译开始>
// Op 是用于文件操作的位集合（联合体）。
# <翻译结束>


<原文开始>
// internalPanic is the custom panic for internal usage.
<原文结束>

# <翻译开始>
// internalPanic 是用于内部使用的自定义恐慌函数。
# <翻译结束>


<原文开始>
// Duration for repeated event filter.
<原文结束>

# <翻译开始>
// 重复事件过滤的持续时间。
# <翻译结束>


<原文开始>
// Custom exit event for internal usage.
<原文结束>

# <翻译开始>
// 自定义退出事件，用于内部使用。
# <翻译结束>


<原文开始>
// Mutex for concurrent safety of defaultWatcher.
<原文结束>

# <翻译开始>
// 用于保护defaultWatcher在并发环境下的安全性的互斥锁。
# <翻译结束>







<原文开始>
// Atomic id generator for callback.
<原文结束>

# <翻译开始>
// 原子式ID生成器，用于回调。
# <翻译结束>


<原文开始>
// New creates and returns a new watcher.
// Note that the watcher number is limited by the file handle setting of the system.
// Eg: fs.inotify.max_user_instances system variable in linux systems.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的监视器。
// 注意，监视器的数量受系统文件句柄设置的限制。
// 例如：在Linux系统中，fs.inotify.max_user_instances系统变量。
# <翻译结束>


<原文开始>
// Add monitors `path` using default watcher with callback function `callbackFunc`.
// The optional parameter `recursive` specifies whether monitoring the `path` recursively, which is true in default.
<原文结束>

# <翻译开始>
// Add 使用默认观察器监控`path`，并使用回调函数`callbackFunc`。
// 可选参数`recursive`指定了是否递归地监控`path`，默认为true。
# <翻译结束>


<原文开始>
// AddOnce monitors `path` using default watcher with callback function `callbackFunc` only once using unique name `name`.
// If AddOnce is called multiple times with the same `name` parameter, `path` is only added to monitor once. It returns error
// if it's called twice with the same `name`.
//
// The optional parameter `recursive` specifies whether monitoring the `path` recursively, which is true in default.
<原文结束>

# <翻译开始>
// AddOnce 使用默认的监视器，通过回调函数 `callbackFunc` 仅对 `path` 进行一次性监控，并使用唯一的名称 `name` 标识。
// 如果多次调用 AddOnce 并传入相同的 `name` 参数，`path` 只会被添加监控一次。如果两次调用时使用了相同的 `name`，将会返回错误。
//
// 可选参数 `recursive` 指定是否递归地监控 `path`，默认情况下其值为 true。
# <翻译结束>


<原文开始>
// Remove removes all monitoring callbacks of given `path` from watcher recursively.
<原文结束>

# <翻译开始>
// Remove 递归地从 watcher 中移除指定 `path` 的所有监控回调函数。
# <翻译结束>


<原文开始>
// RemoveCallback removes specified callback with given id from watcher.
<原文结束>

# <翻译开始>
// RemoveCallback 从观察者中移除具有给定id的指定回调函数。
# <翻译结束>


<原文开始>
// Exit is only used in the callback function, which can be used to remove current callback
// of itself from the watcher.
<原文结束>

# <翻译开始>
// Exit 仅在回调函数中使用，可用于从监视器中移除自身的当前回调。
# <翻译结束>


<原文开始>
// getDefaultWatcher creates and returns the default watcher.
// This is used for lazy initialization purpose.
<原文结束>

# <翻译开始>
// getDefaultWatcher 创建并返回默认的监视器。
// 这是为了实现延迟初始化的目的。
# <翻译结束>






















<原文开始>
// Used for AddOnce feature.
<原文结束>

# <翻译开始>
// 用于AddOnce功能。
# <翻译结束>


<原文开始>
// Callback function.
<原文结束>

# <翻译开始>
// 回调函数。
# <翻译结束>


<原文开始>
// Absolute file path.
<原文结束>

# <翻译开始>
// 绝对文件路径。
# <翻译结束>


<原文开始>
// Id to callback mapping.
<原文结束>

# <翻译开始>
// Id 到回调函数的映射。
# <翻译结束>

