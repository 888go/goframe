
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
// watchLoop starts the loop for event listening from underlying inotify monitor.
<原文结束>

# <翻译开始>
// watchLoop 启动循环以从底层inotify监控器监听事件。 md5:a057c294cb3f7186
# <翻译结束>


<原文开始>
// Filter the repeated event in custom duration.
<原文结束>

# <翻译开始>
// 过滤自定义持续时间内的重复事件。 md5:f7b5d987e84f8092
# <翻译结束>


<原文开始>
// eventLoop is the core event handler.
<原文结束>

# <翻译开始>
// eventLoop是核心事件处理器。 md5:7adb3bf9c821349b
# <翻译结束>


<原文开始>
// If there's no any callback of this path, it removes it from monitor.
<原文结束>

# <翻译开始>
// 如果该路径没有任何回调，就从监控中移除它。 md5:1d18925e16d1ccb5
# <翻译结束>


<原文开始>
					// It should check again the existence of the path.
					// It adds it back to the monitor if it still exists.
<原文结束>

# <翻译开始>
					// 它应该再次检查路径的存在。
					// 如果该路径仍然存在，它会将其重新添加到监视器中。
					// md5:216ebbce200ac7a4
# <翻译结束>


<原文开始>
						// It adds the path back to monitor.
						// We need no worry about the repeat adding.
<原文结束>

# <翻译开始>
						// 将路径重新添加到监控列表。
						// 我们不需要担心重复添加的问题。
						// md5:4487198f5d35bb60
# <翻译结束>


<原文开始>
// Change the event to RENAME, which means it renames itself to its origin name.
<原文结束>

# <翻译开始>
// 将事件更改为 RENAME，这意味着它将自己重命名为原始名称。 md5:7e6fbf14f9528be7
# <翻译结束>


<原文开始>
					// It should check again the existence of the path.
					// It adds it back to the monitor if it still exists.
					// Especially Some editors might do RENAME and then CHMOD when it's editing file.
<原文结束>

# <翻译开始>
					// 应该再次检查路径是否存在。
					// 如果它仍然存在，将其重新添加到监视器中。
					// 特别是有些编辑器在编辑文件时可能会进行RENAME和CHMOD操作。
					// md5:b3ab149aa0d38e0c
# <翻译结束>


<原文开始>
						// It might lost the monitoring for the path, so we add the path back to monitor.
						// We need no worry about the repeat adding.
<原文结束>

# <翻译开始>
						// 可能会丢失对路径的监控，因此我们需将路径重新添加到监控中。
						// 我们无需担心重复添加的问题。
						// md5:d6dd87eba165d9e7
# <翻译结束>


<原文开始>
// Change the event to CHMOD.
<原文结束>

# <翻译开始>
// 将事件更改为CHMOD。 md5:84c563944c4dfa07
# <翻译结束>


<原文开始>
					// =========================================
					// Note that it here just adds the path to monitor without any callback registering,
					// because its parent already has the callbacks.
					// =========================================
<原文结束>

# <翻译开始>
					// =========================================
					// 注意，这里只是添加了要监控的路径，而不需要注册回调，
					// 因为它的父级已经具有了回调。
					// =========================================
					// md5:2b5f1f3849c5ccff
# <翻译结束>


<原文开始>
// If it's a folder, it then does adding recursively to monitor.
<原文结束>

# <翻译开始>
// 如果这是一个文件夹，它会递归地添加以进行监控。 md5:3b1a61cf45e4cf3a
# <翻译结束>


<原文开始>
// If it's a file, it directly adds it to monitor.
<原文结束>

# <翻译开始>
// 如果它是一个文件，就直接将其添加到监控中。 md5:18b66bfd2946b42e
# <翻译结束>


<原文开始>
// Calling the callbacks in order.
<原文结束>

# <翻译开始>
// 按顺序调用回调函数。 md5:426b787bf42f20fa
# <翻译结束>


<原文开始>
// getCallbacks searches and returns all callbacks with given `path`.
// It also searches its parents for callbacks if they're recursive.
<原文结束>

# <翻译开始>
// getCallbacks 搜索并返回所有具有给定 `path` 的回调。如果它们是递归的，还会在其父级中搜索回调。
// md5:abe3c32241868912
# <翻译结束>


<原文开始>
// Firstly add the callbacks of itself.
<原文结束>

# <翻译开始>
// 首先，添加自身的回调。 md5:474fe6dbf371de56
# <翻译结束>


<原文开始>
	// Secondly searches its direct parent for callbacks.
	// It is special handling here, which is the different between `recursive` and `not recursive` logic
	// for direct parent folder of `path` that events are from.
<原文结束>

# <翻译开始>
	// 其次，在其直接父级中搜索回调。
	// 这里有特殊的处理逻辑，区别于“递归”与“非递归”的处理方式，
	// 特指针对来源于`path`的直接父级目录的事件。
	// md5:4e4cd99683eb9f66
# <翻译结束>


<原文开始>
// Lastly searches all the parents of directory of `path` recursively for callbacks.
<原文结束>

# <翻译开始>
// 最后，递归地搜索`path`目录的所有父级以查找回调函数。 md5:24dea4c80a5e5c6d
# <翻译结束>

