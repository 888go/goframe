
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
// Package gspath implements file index and search for folders.
//
<原文结束>

# <翻译开始>
// Package gspath 实现了文件索引和对文件夹的搜索功能。
# <翻译结束>


<原文开始>
// updateCacheByPath adds all files under `path` recursively.
<原文结束>

# <翻译开始>
// updateCacheByPath 递归地将`path`路径下所有文件添加到缓存中。
# <翻译结束>


<原文开始>
// formatCacheName formats `name` with following rules:
// 1. The separator is unified to char '/'.
// 2. The name should be started with '/' (similar as HTTP URI).
<原文结束>

# <翻译开始>
// formatCacheName 根据以下规则格式化 `name`：
// 1. 统一分隔符为字符 '/'。
// 2. 名称应以 '/' 开头（类似于 HTTP URI）。
# <翻译结束>


<原文开始>
// nameFromPath converts `filePath` to cache name.
<原文结束>

# <翻译开始>
// nameFromPath将`filePath`转换为缓存名称。
# <翻译结束>


<原文开始>
// makeCacheValue formats `filePath` to cache value.
<原文结束>

# <翻译开始>
// makeCacheValue将`filePath`格式化为缓存值。
# <翻译结束>


<原文开始>
// parseCacheValue parses cache value to file path and type.
<原文结束>

# <翻译开始>
// parseCacheValue 将缓存值解析为文件路径和类型。
# <翻译结束>


<原文开始>
// addToCache adds an item to cache.
// If `filePath` is a directory, it also adds its all sub files/directories recursively
// to the cache.
<原文结束>

# <翻译开始>
// addToCache 将一个项目添加到缓存中。
// 如果 `filePath` 是一个目录，它还会递归地将该目录下的所有子文件和子目录
// 一并添加到缓存中。
# <翻译结束>


<原文开始>
// If it's a directory, it adds all of its sub files/directories.
<原文结束>

# <翻译开始>
// 如果它是一个目录，那么它会添加其所有子文件/子目录。
# <翻译结束>


<原文开始>
// fmt.Println("gspath add to cache:", filePath, files)
<原文结束>

# <翻译开始>
// 输出到控制台：gspath添加到缓存中: filePath, files
# <翻译结束>


<原文开始>
// addMonitorByPath adds gfsnotify monitoring recursively.
// When the files under the directory are updated, the cache will be updated meanwhile.
// Note that since the listener is added recursively, if you delete a directory, the files (including the directory)
// under the directory will also generate delete events, which means it will generate N+1 events in total
// if a directory deleted and there're N files under it.
<原文结束>

# <翻译开始>
// addMonitorByPath 递归地添加gfsnotify监控。
// 当目录下的文件被更新时，缓存会同时得到更新。
// 注意，由于监听器是递归添加的，如果你删除了一个目录，那么该目录下的所有文件（包括目录本身）
// 也会生成删除事件。这意味着，如果一个目录被删除且其下有N个文件，则总共会产生N+1个事件。
# <翻译结束>


<原文开始>
// removeMonitorByPath removes gfsnotify monitoring of `path` recursively.
<原文结束>

# <翻译开始>
// removeMonitorByPath 递归地移除对`path`的gfsnotify监控。
# <翻译结束>


<原文开始>
// Add itself firstly.
<原文结束>

# <翻译开始>
// 首先对其自身进行加法操作。
# <翻译结束>


<原文开始>
// glog.Debug(event.String())
<原文结束>

# <翻译开始>
// glog.Debug(event.String()) // 使用glog库输出debug级别的日志，内容为event对象转换为字符串后的结果
# <翻译结束>

