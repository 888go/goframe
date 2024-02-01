
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
// watchLoop starts the loop for event listening from underlying inotify monitor.
<原文结束>

# <翻译开始>
// watchLoop 启动一个循环，用于从底层 inotify 监视器监听事件。
# <翻译结束>







<原文开始>
// Filter the repeated event in custom duration.
<原文结束>

# <翻译开始>
// 在自定义时间段内过滤重复事件。
# <翻译结束>


<原文开始>
// eventLoop is the core event handler.
<原文结束>

# <翻译开始>
// eventLoop 是核心事件处理器。
# <翻译结束>


<原文开始>
// If there's no any callback of this path, it removes it from monitor.
<原文结束>

# <翻译开始>
// 如果此路径没有任何回调函数，它将从监视器中移除该路径。
# <翻译结束>


<原文开始>
					// It should check again the existence of the path.
					// It adds it back to the monitor if it still exists.
<原文结束>

# <翻译开始>
// 应该再次检查路径是否存在。
// 如果路径仍然存在，则将其重新添加到监视器中。
# <翻译结束>


<原文开始>
						// It adds the path back to monitor.
						// We need no worry about the repeat adding.
<原文结束>

# <翻译开始>
// 它将路径添加回监控中。
// 我们无需担心重复添加的问题。
# <翻译结束>


<原文开始>
// Change the event to RENAME, which means it renames itself to its origin name.
<原文结束>

# <翻译开始>
// 将事件更改为 RENAME，这意味着它将自身重命名为原始名称。
# <翻译结束>


<原文开始>
					// It should check again the existence of the path.
					// It adds it back to the monitor if it still exists.
					// Especially Some editors might do RENAME and then CHMOD when it's editing file.
<原文结束>

# <翻译开始>
// 它应当再次检查路径是否存在。
// 如果路径仍然存在，则将其重新添加到监视器中。
// 特别需要注意的是，某些编辑器在编辑文件时可能会先执行重命名（RENAME）操作，然后进行权限更改（CHMOD）操作。
# <翻译结束>


<原文开始>
						// It might lost the monitoring for the path, so we add the path back to monitor.
						// We need no worry about the repeat adding.
<原文结束>

# <翻译开始>
// 可能会丢失对路径的监控，所以我们将路径重新添加回监控中。
// 我们无需担心重复添加的问题。
# <翻译结束>







<原文开始>
					// =========================================
					// Note that it here just adds the path to monitor without any callback registering,
					// because its parent already has the callbacks.
					// =========================================
<原文结束>

# <翻译开始>
// =========================================
// 注意，这里仅添加要监视的路径而无需注册任何回调函数，
// 因为其父级已经拥有这些回调函数。
// =========================================
# <翻译结束>


<原文开始>
// If it's a folder, it then does adding recursively to monitor.
<原文结束>

# <翻译开始>
// 如果这是一个文件夹，那么它会递归地添加到监控中。
# <翻译结束>


<原文开始>
// If it's a file, it directly adds it to monitor.
<原文结束>

# <翻译开始>
// 如果它是一个文件，就直接将其添加到监控中。
# <翻译结束>


<原文开始>
// Calling the callbacks in order.
<原文结束>

# <翻译开始>
// 按照顺序调用回调函数。
# <翻译结束>


<原文开始>
// getCallbacks searches and returns all callbacks with given `path`.
// It also searches its parents for callbacks if they're recursive.
<原文结束>

# <翻译开始>
// getCallbacks 搜索并返回给定 `path` 的所有回调函数。
// 如果它们是递归的，还会在其父级中搜索回调函数。
# <翻译结束>


<原文开始>
// Firstly add the callbacks of itself.
<原文结束>

# <翻译开始>
// 首先添加自身的回调函数。
# <翻译结束>


<原文开始>
	// Secondly searches its direct parent for callbacks.
	// It is special handling here, which is the different between `recursive` and `not recursive` logic
	// for direct parent folder of `path` that events are from.
<原文结束>

# <翻译开始>
// 第二步，在其直接父级中搜索回调函数。
// 这里有特殊处理，这是`递归`和`非递归`逻辑之间的区别，
// 特指从`path`的直接父目录传递过来的事件。
# <翻译结束>


<原文开始>
// Lastly searches all the parents of directory of `path` recursively for callbacks.
<原文结束>

# <翻译开始>
// 最后，递归地在`path`目录的所有父级目录中搜索回调函数。
# <翻译结束>


<原文开始>
// Change the event to CHMOD.
<原文结束>

# <翻译开始>
// 将事件更改为 CHMOD。
# <翻译结束>

