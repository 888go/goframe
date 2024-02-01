
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
// It searches file internally with high performance in order by the directory adding sequence.
// Note that:
// If caching feature enabled, there would be a searching delay after adding/deleting files.
<原文结束>

# <翻译开始>
// Package gspath 实现了文件索引和目录搜索功能。
//
// 它按照目录添加的顺序，内部高效地进行文件搜索。
// 注意：
// 如果启用了缓存功能，在添加或删除文件后，会有一个搜索延迟。
# <翻译结束>


<原文开始>
// SPath manages the path searching feature.
<原文结束>

# <翻译开始>
// SPath 管理路径搜索功能。
# <翻译结束>


<原文开始>
// The searching directories array.
<原文结束>

# <翻译开始>
// 搜索目录的数组。
# <翻译结束>


<原文开始>
// SPathCacheItem is a cache item for searching.
<原文结束>

# <翻译开始>
// SPathCacheItem 是一个用于搜索的缓存项。
# <翻译结束>


<原文开始>
// Path to searching object mapping, used for instance management.
<原文结束>

# <翻译开始>
// 对象映射的搜索路径，用于实例管理。
# <翻译结束>


<原文开始>
// New creates and returns a new path searching manager.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的路径搜索管理器。
# <翻译结束>


<原文开始>
// Get creates and returns an instance of searching manager for given path.
// The parameter `cache` specifies whether using cache feature for this manager.
// If cache feature is enabled, it asynchronously and recursively scans the path
// and updates all sub files/folders to the cache using package gfsnotify.
<原文结束>

# <翻译开始>
// Get根据给定路径创建并返回一个搜索管理器实例。
// 参数`cache`用于指定此管理器是否启用缓存功能。
// 如果启用了缓存功能，它会异步地递归扫描该路径，
// 并使用gfsnotify包更新所有子文件/文件夹到缓存中。
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
// 参数 `root` 应为一个绝对路径。出于性能考虑，它不会自动将 `root` 转换为绝对路径。
// 可选参数 `indexFiles` 指定了在结果是目录时要搜索的索引文件列表。
// 例如，如果结果 `filePath` 是一个目录，并且 `indexFiles` 是 [index.html, main.html]，那么它还会在 `filePath` 下搜索 [index.html, main.html]。
// 如果找到其中任意一个文件，它将返回该文件的绝对路径，否则返回 `filePath`。
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
// SearchWithCache 在启用缓存功能的情况下，搜索路径`root`下的文件`name`。
// 参数`root`应为绝对路径。出于性能考虑，它不会自动将`root`转换为绝对路径。
// 可选参数`indexFiles`用于指定当结果是目录时要搜索的索引文件。
// 例如，如果结果`filePath`是一个目录，并且`indexFiles`为[index.html, main.html]，那么它还会在`filePath`下搜索[index.html, main.html]。
// 如果找到其中任何一个文件，则返回该绝对文件路径，否则返回`filePath`。
# <翻译结束>


<原文开始>
// Set deletes all other searching directories and sets the searching directory for this manager.
<原文结束>

# <翻译开始>
// Set 删除所有其他搜索目录，并为此管理器设置搜索目录。
# <翻译结束>


<原文开始>
// The set path must be a directory.
<原文结束>

# <翻译开始>
// 设置的路径必须是一个目录。
# <翻译结束>


<原文开始>
// Add adds more searching directory to the manager.
// The manager will search file in added order.
<原文结束>

# <翻译开始>
// Add 向管理器添加更多搜索目录。
// 管理器将按照添加顺序搜索文件。
# <翻译结束>


<原文开始>
// The added path must be a directory.
<原文结束>

# <翻译开始>
// 添加的路径必须是一个目录。
# <翻译结束>


<原文开始>
		// fmt.Println("gspath:", realPath, sp.paths.Search(realPath))
		// It will not add twice for the same directory.
<原文结束>

# <翻译开始>
// fmt.Println("gspath:", realPath, sp.paths.Search(realPath)) // 输出gspath:（realPath的值），以及sp.paths在realPath路径下搜索的结果
// 同一目录不会被重复添加两次。
# <翻译结束>


<原文开始>
// Search searches file `name` in the manager.
// The optional parameter `indexFiles` specifies the searching index files when the result is a directory.
// For example, if the result `filePath` is a directory, and `indexFiles` is [index.html, main.html], it will also
// search [index.html, main.html] under `filePath`. It returns the absolute file path if any of them found,
// or else it returns `filePath`.
<原文结束>

# <翻译开始>
// Search 在manager中搜索文件`name`。
// 可选参数`indexFiles`指定了当结果为目录时需要搜索的索引文件。
// 例如，如果结果`filePath`是一个目录，并且`indexFiles`是[index.html, main.html]，
// 它还会在`filePath`下搜索[index.html, main.html]。如果有任何一项被找到，它将返回该绝对文件路径，
// 否则返回`filePath`。
# <翻译结束>


<原文开始>
// Security check: the result file path must be under the searching directory.
<原文结束>

# <翻译开始>
// 安全检查：结果文件路径必须在搜索目录下。
# <翻译结束>


<原文开始>
// Remove deletes the `path` from cache files of the manager.
// The parameter `path` can be either an absolute path or just a relative file name.
<原文结束>

# <翻译开始>
// Remove 从管理器的缓存文件中删除指定`path`。
// 参数`path`可以是绝对路径，也可以只是一个相对文件名。
# <翻译结束>


<原文开始>
// Paths returns all searching directories.
<原文结束>

# <翻译开始>
// Paths 返回所有搜索目录。
# <翻译结束>


<原文开始>
// AllPaths returns all paths cached in the manager.
<原文结束>

# <翻译开始>
// AllPaths 返回缓存在manager中的所有路径。
# <翻译结束>


<原文开始>
// Size returns the count of the searching directories.
<原文结束>

# <翻译开始>
// Size 返回搜索目录的数量。
# <翻译结束>


<原文开始>
// Absolute path for file/dir.
<原文结束>

# <翻译开始>
// 绝对路径（文件/目录）
# <翻译结束>


<原文开始>
// Is directory or not.
<原文结束>

# <翻译开始>
// 是否是目录。
# <翻译结束>


<原文开始>
// Using cache feature.
<原文结束>

# <翻译开始>
// 使用缓存功能。
# <翻译结束>

