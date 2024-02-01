
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
// Add monitors `path` with callback function `callbackFunc` to the watcher.
// The optional parameter `recursive` specifies whether monitoring the `path` recursively,
// which is true in default.
<原文结束>

# <翻译开始>
// 添加监控，将回调函数`callbackFunc`添加到watcher中，并监控路径`path`。
// 可选参数`recursive`指定了是否递归地监控路径`path`，默认为true。
# <翻译结束>


<原文开始>
// AddOnce monitors `path` with callback function `callbackFunc` only once using unique name
// `name` to the watcher. If AddOnce is called multiple times with the same `name` parameter,
// `path` is only added to monitor once.
//
// It returns error if it's called twice with the same `name`.
//
// The optional parameter `recursive` specifies whether monitoring the `path` recursively,
// which is true in default.
<原文结束>

# <翻译开始>
// AddOnce 通过唯一名称 `name` 使用回调函数 `callbackFunc` 仅对 `path` 进行一次性监控。
// 如果多次调用 AddOnce 并传入相同的 `name` 参数，`path` 只会被添加一次进行监控。
//
// 若同一 `name` 被调用两次，将会返回错误。
//
// 可选参数 `recursive` 指定是否递归地监控 `path`，默认情况下其值为 true。
# <翻译结束>


<原文开始>
// Firstly add the path to watcher.
<原文结束>

# <翻译开始>
// 首先将路径添加到监视器中。
# <翻译结束>


<原文开始>
		// If it's recursive adding, it then adds all sub-folders to the monitor.
		// NOTE:
		// 1. It only recursively adds **folders** to the monitor, NOT files,
		//    because if the folders are monitored and their sub-files are also monitored.
		// 2. It bounds no callbacks to the folders, because it will search the callbacks
		//    from its parent recursively if any event produced.
<原文结束>

# <翻译开始>
// 如果是递归添加，则将所有子文件夹添加到监视器中。
// 注意：
// 1. 它仅递归地向监视器添加**文件夹**，而不添加文件，
//    因为如果监视了文件夹及其子文件也会被监视。
// 2. 它没有给这些文件夹绑定任何回调函数，因为在产生任何事件时，它会从其父级开始递归地搜索回调函数。
# <翻译结束>


<原文开始>
// addWithCallbackFunc adds the path to underlying monitor, creates and returns a callback object.
// Very note that if it calls multiple times with the same `path`, the latest one will overwrite the previous one.
<原文结束>

# <翻译开始>
// addWithCallbackFunc 将路径添加到底层监视器中，创建并返回一个回调对象。
// 非常注意，如果对同一`path`调用多次，最新的一次将会覆盖之前的所有内容。
# <翻译结束>


<原文开始>
// Check and convert the given path to absolute path.
<原文结束>

# <翻译开始>
// 检查并把给定的路径转换为绝对路径。
# <翻译结束>







<原文开始>
// Register the callback to watcher.
<原文结束>

# <翻译开始>
// 将回调函数注册到监视器中。
# <翻译结束>


<原文开始>
// Add the path to underlying monitor.
<原文结束>

# <翻译开始>
// 将路径添加到基础监视器中。
# <翻译结束>


<原文开始>
// Add the callback to global callback map.
<原文结束>

# <翻译开始>
// 将回调函数添加到全局回调映射中。
# <翻译结束>







<原文开始>
// Remove removes monitor and all callbacks associated with the `path` recursively.
<原文结束>

# <翻译开始>
// Remove 递归地移除与`path`关联的监视器及其所有回调。
# <翻译结束>


<原文开始>
// Firstly remove the callbacks of the path.
<原文结束>

# <翻译开始>
// 首先移除该路径的回调函数。
# <翻译结束>


<原文开始>
// Secondly remove monitor of all sub-files which have no callbacks.
<原文结束>

# <翻译开始>
// 其次，移除所有无回调函数的子文件的监控
# <翻译结束>


<原文开始>
// Lastly remove the monitor of the path from underlying monitor.
<原文结束>

# <翻译开始>
// 最后，从底层监视器中移除该路径的监控。
# <翻译结束>


<原文开始>
// checkPathCanBeRemoved checks whether the given path have no callbacks bound.
<原文结束>

# <翻译开始>
// checkPathCanBeRemoved 检查给定路径是否未绑定任何回调函数。
# <翻译结束>


<原文开始>
// Firstly check the callbacks in the watcher directly.
<原文结束>

# <翻译开始>
// 首先检查watcher中直接的回调函数。
# <翻译结束>


<原文开始>
// Secondly check its parent whether has callbacks.
<原文结束>

# <翻译开始>
// 第二步检查其父级是否具有回调函数。
# <翻译结束>


<原文开始>
// Recursively check its parent.
<原文结束>

# <翻译开始>
// 递归检查其父级
# <翻译结束>


<原文开始>
// RemoveCallback removes callback with given callback id from watcher.
<原文结束>

# <翻译开始>
// RemoveCallback 从 watcher 中移除具有给定回调 id 的回调函数。
# <翻译结束>


<原文开始>
// Create callback object.
<原文结束>

# <翻译开始>
// 创建回调对象。
# <翻译结束>


<原文开始>
// Close closes the watcher.
<原文结束>

# <翻译开始>
// Close 关闭监视器。
# <翻译结束>

