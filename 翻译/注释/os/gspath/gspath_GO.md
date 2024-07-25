
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
// It searches file internally with high performance in order by the directory adding sequence.
// Note that:
// If caching feature enabled, there would be a searching delay after adding/deleting files.
<原文结束>

# <翻译开始>
// 包gspath实现了文件夹的索引和搜索功能。
//
// 它按照目录添加顺序，内部高效地进行文件搜索。请注意：
// 如果启用了缓存功能，添加或删除文件后可能会有搜索延迟。 md5:626b2e878f4df376
# <翻译结束>


<原文开始>
// SPath manages the path searching feature.
<原文结束>

# <翻译开始>
// SPath 管理路径搜索功能。 md5:703d47a59dea75cf
# <翻译结束>


<原文开始>
// The searching directories array.
<原文结束>

# <翻译开始>
// 搜索目录数组。 md5:1bb898235e976652
# <翻译结束>


<原文开始>
// SPathCacheItem is a cache item for searching.
<原文结束>

# <翻译开始>
// SPathCacheItem 是用于搜索的缓存项。 md5:9b45b61130ff4d97
# <翻译结束>


<原文开始>
// Absolute path for file/dir.
<原文结束>

# <翻译开始>
// 文件或目录的绝对路径。 md5:8adda623344c67c8
# <翻译结束>


<原文开始>
// Path to searching object mapping, used for instance management.
<原文结束>

# <翻译开始>
	// 用于实例管理的对象映射的搜索路径。 md5:1eefbdf17ed9097e
# <翻译结束>


<原文开始>
// New creates and returns a new path searching manager.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的路径搜索管理器。 md5:4a9d5d03b9c2c8be
# <翻译结束>


<原文开始>
// Get creates and returns an instance of searching manager for given path.
// The parameter `cache` specifies whether using cache feature for this manager.
// If cache feature is enabled, it asynchronously and recursively scans the path
// and updates all sub files/folders to the cache using package gfsnotify.
<原文结束>

# <翻译开始>
// Get 创建并返回一个针对给定路径的搜索管理器实例。
// 参数 `cache` 指定是否为此管理器启用缓存功能。
// 如果启用了缓存功能，它将异步且递归地扫描该路径，
// 并使用 gfsnotify 包将所有子文件/文件夹更新到缓存中。 md5:db411c9b09cbef91
# <翻译结束>


<原文开始>
// Search searches file `name` under path `root`.
// The parameter `root` should be an absolute path. It will not automatically
// convert `root` to absolute path for performance reason.
// The optional parameter `indexFiles` specifies the searching index files when the result is a directory.
// For example, if the result `filePath` is a directory, and `indexFiles` is [index.html, main.html], it will also
// search [index.html, main.html] under `filePath`. It returns the absolute file path if any of them found,
// or else it returns `filePath`.
<原文结束>

# <翻译开始>
// Search 在路径 `root` 下搜索文件 `name`。
// 参数 `root` 应为绝对路径。出于性能考虑，它不会自动将 `root` 转换为绝对路径。
// 可选参数 `indexFiles` 指定在结果为目录时要搜索的索引文件。
// 例如，如果结果 `filePath` 是一个目录，并且 `indexFiles` 是 [index.html, main.html]，它将在 `filePath` 下也搜索这两个文件。
// 如果找到其中任何一个文件，它将返回该文件的绝对路径，否则返回 `filePath`。 md5:dfc991bd35d2d178
# <翻译结束>


<原文开始>
// SearchWithCache searches file `name` under path `root` with cache feature enabled.
// The parameter `root` should be an absolute path. It will not automatically
// convert `root` to absolute path for performance reason.
// The optional parameter `indexFiles` specifies the searching index files when the result is a directory.
// For example, if the result `filePath` is a directory, and `indexFiles` is [index.html, main.html], it will also
// search [index.html, main.html] under `filePath`. It returns the absolute file path if any of them found,
// or else it returns `filePath`.
<原文结束>

# <翻译开始>
// SearchWithCache 在启用缓存的情况下，在路径`root`下搜索文件`name`。参数`root`应为绝对路径，出于性能考虑，它不会自动将`root`转换为绝对路径。
// 可选参数`indexFiles`用于指定当结果为目录时的索引文件。例如，如果结果`filePath`是一个目录，并且`indexFiles`为`[index.html, main.html]`，它将在`filePath`下搜索`[index.html, main.html)`。如果有找到任何文件，则返回其绝对文件路径，否则返回`filePath`。 md5:f0b25342a4685319
# <翻译结束>


<原文开始>
// Set deletes all other searching directories and sets the searching directory for this manager.
<原文结束>

# <翻译开始>
// Set删除所有其他搜索目录，并为这个管理器设置搜索目录。 md5:6bb092ed0381b154
# <翻译结束>


<原文开始>
// The set path must be a directory.
<原文结束>

# <翻译开始>
	// 设置的路径必须是一个目录。 md5:a1d52d7d0583a6ef
# <翻译结束>


<原文开始>
// Add adds more searching directory to the manager.
// The manager will search file in added order.
<原文结束>

# <翻译开始>
// Add 向管理器添加更多的搜索目录。
// 管理器将按照添加的顺序查找文件。 md5:b27b49ecc2f1758a
# <翻译结束>


<原文开始>
// The added path must be a directory.
<原文结束>

# <翻译开始>
	// 添加的路径必须是一个目录。 md5:3e2662f535c6872c
# <翻译结束>


<原文开始>
		// fmt.Println("gspath:", realPath, sp.paths.Search(realPath))
		// It will not add twice for the same directory.
<原文结束>

# <翻译开始>
		// fmt.Println("gspath:", realPath, sp.paths.Search(realPath))
		// 对于同一个目录，它不会重复添加。 md5:701deef87cf571aa
# <翻译结束>


<原文开始>
// Search searches file `name` in the manager.
// The optional parameter `indexFiles` specifies the searching index files when the result is a directory.
// For example, if the result `filePath` is a directory, and `indexFiles` is [index.html, main.html], it will also
// search [index.html, main.html] under `filePath`. It returns the absolute file path if any of them found,
// or else it returns `filePath`.
<原文结束>

# <翻译开始>
// Search 在管理器中搜索文件`name`。
// 可选参数`indexFiles`指定了当搜索结果为目录时，需要在该目录下查找的索引文件列表。
// 例如，如果搜索结果`filePath`是一个目录，并且`indexFiles`是[index.html, main.html]，那么它还会
// 在`filePath`目录下查找[index.html, main.html]。如果找到了其中任何一个文件，就返回其绝对路径；
// 否则，直接返回`filePath`。 md5:8210196c6e2ae787
# <翻译结束>


<原文开始>
// Security check: the result file path must be under the searching directory.
<原文结束>

# <翻译开始>
					// 安全检查：结果文件路径必须在搜索目录下。 md5:b83726d297546baa
# <翻译结束>


<原文开始>
// Remove deletes the `path` from cache files of the manager.
// The parameter `path` can be either an absolute path or just a relative file name.
<原文结束>

# <翻译开始>
// Remove 从管理器的缓存文件中删除`path`。参数`path`可以是绝对路径或仅仅是相对文件名。 md5:30f46aaaf75a1da8
# <翻译结束>


<原文开始>
// Paths returns all searching directories.
<原文结束>

# <翻译开始>
// Paths返回所有搜索目录。 md5:0e02e3a85da8e197
# <翻译结束>


<原文开始>
// AllPaths returns all paths cached in the manager.
<原文结束>

# <翻译开始>
// AllPaths 返回存储在管理器中的所有路径。 md5:75157edfcae7d2a0
# <翻译结束>


<原文开始>
// Size returns the count of the searching directories.
<原文结束>

# <翻译开始>
// Size 返回搜索目录的数量。 md5:e115dd584d3351a2
# <翻译结束>

