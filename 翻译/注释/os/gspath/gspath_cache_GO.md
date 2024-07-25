
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
// Package gspath implements file index and search for folders.
//
<原文结束>

# <翻译开始>
// gspath 包实现了对文件夹的索引和搜索功能。 md5:04299e0152ee648b
# <翻译结束>


<原文开始>
// updateCacheByPath adds all files under `path` recursively.
<原文结束>

# <翻译开始>
// updateCacheByPath 递归地在`path`下添加所有文件。 md5:ef869f8f30af135a
# <翻译结束>


<原文开始>
// formatCacheName formats `name` with following rules:
// 1. The separator is unified to char '/'.
// 2. The name should be started with '/' (similar as HTTP URI).
<原文结束>

# <翻译开始>
// formatCacheName 根据以下规则格式化`name`：
// 1. 分隔符统一为字符'/'。
// 2. 名称应以'/'开头（类似于HTTP URI）。 md5:ed5316ca14ce4d4c
# <翻译结束>


<原文开始>
// nameFromPath converts `filePath` to cache name.
<原文结束>

# <翻译开始>
// nameFromPath 将 `filePath` 转换为缓存名称。 md5:5e0f623421b9d54d
# <翻译结束>


<原文开始>
// makeCacheValue formats `filePath` to cache value.
<原文结束>

# <翻译开始>
// makeCacheValue 将 `filePath` 格式化为缓存值。 md5:ac703cee872ac9d4
# <翻译结束>


<原文开始>
// parseCacheValue parses cache value to file path and type.
<原文结束>

# <翻译开始>
// parseCacheValue 解析缓存值为文件路径和类型。 md5:c7d4d6cc498a746f
# <翻译结束>


<原文开始>
// addToCache adds an item to cache.
// If `filePath` is a directory, it also adds its all sub files/directories recursively
// to the cache.
<原文结束>

# <翻译开始>
// addToCache 将一个项目添加到缓存中。
// 如果 `filePath` 是一个目录，它还会递归地将所有子文件/目录添加到缓存中。 md5:836028ec6822544d
# <翻译结束>


<原文开始>
// If it's a directory, it adds all of its sub files/directories.
<原文结束>

# <翻译开始>
	// 如果是一个目录，它会添加其下所有的子文件和子目录。 md5:d133c73c85e80b5b
# <翻译结束>


<原文开始>
// fmt.Println("gspath add to cache:", filePath, files)
<原文结束>

# <翻译开始>
			// fmt.Println("将文件路径", filePath, "和文件列表添加到缓存:", files). md5:787f23087852cffe
# <翻译结束>


<原文开始>
// addMonitorByPath adds gfsnotify monitoring recursively.
// When the files under the directory are updated, the cache will be updated meanwhile.
// Note that since the listener is added recursively, if you delete a directory, the files (including the directory)
// under the directory will also generate delete events, which means it will generate N+1 events in total
// if a directory deleted and there're N files under it.
<原文结束>

# <翻译开始>
// addMonitorByPath 递归地添加 gfsnotify 监控。
// 当目录下的文件被更新时，缓存也会同时被更新。
// 注意，由于监听是递归添加的，如果你删除一个目录，该目录下的所有文件（包括目录本身）都会触发删除事件，这意味着如果删除了一个包含 N 个文件的目录，总共会产生 N+1 个事件。 md5:0142c351fd8dd58f
# <翻译结束>


<原文开始>
// glog.Debug(event.String())
<原文结束>

# <翻译开始>
		// 这个Go语言注释的中文翻译是：使用glog库的Debug级别记录event的字符串表示。 md5:9b7b21454414b499
# <翻译结束>


<原文开始>
// removeMonitorByPath removes gfsnotify monitoring of `path` recursively.
<原文结束>

# <翻译开始>
// removeMonitorByPath 递归地移除对 `path` 的 gfsnotify 监控。 md5:8401f0941bb9504c
# <翻译结束>

