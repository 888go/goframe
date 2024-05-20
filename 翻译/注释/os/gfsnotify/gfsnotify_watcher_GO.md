
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
// Add monitors `path` with callback function `callbackFunc` to the watcher.
// The optional parameter `recursive` specifies whether monitoring the `path` recursively,
// which is true in default.
<原文结束>

# <翻译开始>
// 将监控器添加到观察者，监控的路径为`path`，回调函数为`callbackFunc`。
// 可选参数`recursive`指定是否递归监控`path`，默认值为真。
// md5:35e0c4a9c0901ef8
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
// AddOnce 使用唯一的名称 `name` 监控 `path`，并调用回调函数 `callbackFunc` 只一次。如果多次使用相同的 `name` 参数调用 AddOnce，`path` 将只被添加一次监控。
// 
// 如果两次使用相同的 `name` 调用，将返回错误。
// 
// 可选参数 `recursive` 指定是否递归地监控 `path`，默认为 true。
// md5:6ead1d3d4bff4432
# <翻译结束>


<原文开始>
// Firstly add the path to watcher.
<原文结束>

# <翻译开始>
// 首先，添加路径到观察者。. md5:8830f7aece4dab2e
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
// 如果是递归添加，那么它会将所有子文件夹都添加到监视中。
// 注意：
// 1. 它仅递归地向监视器添加**文件夹**，而不添加文件，
//    因为如果文件夹被监视，其下属文件自然也会被监视。
// 2. 它不对这些文件夹绑定回调函数，因为如果有任何事件产生，
//    它会从父级开始递归地查找对应的回调函数。
// md5:ebaf807e7d2c1bb2
# <翻译结束>


<原文开始>
// addWithCallbackFunc adds the path to underlying monitor, creates and returns a callback object.
// Very note that if it calls multiple times with the same `path`, the latest one will overwrite the previous one.
<原文结束>

# <翻译开始>
// addWithCallbackFunc 将路径添加到底层监控器中，创建并返回一个回调对象。
// 请注意，如果多次调用该函数并传入相同的`path`，最新的调用将覆盖之前的设置。
// md5:bec1b4834bd3126d
# <翻译结束>


<原文开始>
// Check and convert the given path to absolute path.
<原文结束>

# <翻译开始>
// 检查并转换给定的路径为绝对路径。. md5:a7b26b31d4dc4d54
# <翻译结束>


<原文开始>
// Create callback object.
<原文结束>

# <翻译开始>
// 创建回调对象。. md5:35c1374926d9f0ab
# <翻译结束>


<原文开始>
// Register the callback to watcher.
<原文结束>

# <翻译开始>
// 向监视器注册回调函数。. md5:803a60e7f5c04013
# <翻译结束>


<原文开始>
// Add the path to underlying monitor.
<原文结束>

# <翻译开始>
// 将路径添加到基础监控器中。. md5:3136f0f0e2cd9407
# <翻译结束>


<原文开始>
// Add the callback to global callback map.
<原文结束>

# <翻译开始>
// 将回调添加到全局回调映射中。. md5:32fee24607f18b97
# <翻译结束>


<原文开始>
// Close closes the watcher.
<原文结束>

# <翻译开始>
// Close 关闭监听器。. md5:c20cd2708e199b34
# <翻译结束>


<原文开始>
// Remove removes monitor and all callbacks associated with the `path` recursively.
<原文结束>

# <翻译开始>
// Remove 递归地移除与"path"关联的监视器和所有回调。. md5:e48d059cb96966c1
# <翻译结束>


<原文开始>
// Firstly remove the callbacks of the path.
<原文结束>

# <翻译开始>
// 首先移除路径上的回调函数。. md5:15ba778318ad7bb9
# <翻译结束>


<原文开始>
// Secondly remove monitor of all sub-files which have no callbacks.
<原文结束>

# <翻译开始>
// 其次，移除所有没有回调函数的子文件的监控。. md5:477e906da20585f9
# <翻译结束>


<原文开始>
// Lastly remove the monitor of the path from underlying monitor.
<原文结束>

# <翻译开始>
// 最后，从底层监视器中删除路径的监视器。. md5:d85b24985f7de449
# <翻译结束>


<原文开始>
// checkPathCanBeRemoved checks whether the given path have no callbacks bound.
<原文结束>

# <翻译开始>
// checkPathCanBeRemoved 检查给定路径是否绑定了没有回调。. md5:affaad498733b441
# <翻译结束>


<原文开始>
// Firstly check the callbacks in the watcher directly.
<原文结束>

# <翻译开始>
// 首先直接检查监视器中的回调函数。. md5:7167bd75166cef8a
# <翻译结束>


<原文开始>
// Secondly check its parent whether has callbacks.
<原文结束>

# <翻译开始>
// 其次，检查其父级是否具有回调函数。. md5:e913cd965e3ce822
# <翻译结束>


<原文开始>
// Recursively check its parent.
<原文结束>

# <翻译开始>
// 递归检查其父节点。. md5:4557bd1d1f1bec8a
# <翻译结束>


<原文开始>
// RemoveCallback removes callback with given callback id from watcher.
<原文结束>

# <翻译开始>
// RemoveCallback 从观察者中移除具有给定回调ID的回调。. md5:78b678cca3a84b90
# <翻译结束>

