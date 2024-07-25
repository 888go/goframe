
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
// SetPath sets the configuration `directory` path for file search.
// The parameter `path` can be absolute or relative `directory` path,
// but absolute `directory` path is strongly recommended.
//
// Note that this parameter is a path to a directory not a file.
<原文结束>

# <翻译开始>
// SetPath 设置文件搜索的配置`目录`路径。
// 参数 `path` 可以是绝对或相对的`目录`路径，
// 但强烈建议使用绝对`目录`路径。
//
// 注意，此参数是一个指向目录的路径，而不是指向文件的路径。
// md5:56f162e4bbfc634d
# <翻译结束>


<原文开始>
// AddPath adds an absolute or relative `directory` path to the search paths.
//
// Note that this parameter is paths to a directories not files.
<原文结束>

# <翻译开始>
// AddPath 向搜索路径中添加一个绝对或相对的`目录`路径。
//
// 请注意，此参数是目录路径，而不是文件路径。
// md5:25c79c7444dc4e16
# <翻译结束>


<原文开始>
// doAddPath adds an absolute or relative `directory` path to the search paths.
<原文结束>

# <翻译开始>
// doAddPath 将绝对或相对的 `directory` 路径添加到搜索路径中。 md5:43115dba5403276a
# <翻译结束>


<原文开始>
	// It firstly checks the resource manager,
	// and then checks the filesystem for the path.
<原文结束>

# <翻译开始>
	// 首先检查资源管理器，然后在文件系统中查找路径。
	// md5:deb5a0d060375b57
# <翻译结束>


<原文开始>
// GetPaths returns the searching directory path array of current configuration manager.
<原文结束>

# <翻译开始>
// GetPaths 返回当前配置管理器的搜索目录路径数组。 md5:c77738d1ef96cc99
# <翻译结束>


<原文开始>
// doGetFilePath returns the absolute configuration file path for the given filename by `file`.
// If `file` is not passed, it returns the configuration file path of the default name.
// It returns an empty `path` string and an error if the given `file` does not exist.
<原文结束>

# <翻译开始>
// doGetFilePath 根据`file`返回绝对配置文件路径。
// 如果未传递`file`，则返回默认名称的配置文件路径。
// 如果给定的`file`不存在，它将返回一个空的`path`字符串和一个错误。
// md5:4044ef5a7532d997
# <翻译结束>


<原文开始>
// Searching resource manager.
<原文结束>

# <翻译开始>
	// 在搜索资源管理器。 md5:52083f8252a4c319
# <翻译结束>


<原文开始>
// Searching local file system.
<原文结束>

# <翻译开始>
	// 在本地文件系统中搜索。 md5:a557bf6cadf8eec7
# <翻译结束>


<原文开始>
// GetFilePath returns the absolute configuration file path for the given filename by `file`.
// If `file` is not passed, it returns the configuration file path of the default name.
// It returns an empty `path` string and an error if the given `file` does not exist.
<原文结束>

# <翻译开始>
// GetFilePath 通过 `file` 参数返回给定文件名的绝对配置文件路径。
// 如果没有传递 `file`，则返回默认名称的配置文件路径。
// 如果给定的 `file` 不存在，它将返回一个空的 `path` 字符串和一个错误。
// md5:b116b9d063e12bc9
# <翻译结束>


<原文开始>
		// If it's not using default configuration or its configuration file is not available,
		// it searches the possible configuration file according to the name and all supported
		// file types.
<原文结束>

# <翻译开始>
		// 如果它没有使用默认配置，或者其配置文件不可用，
		// 它将根据名称和所有支持的文件类型搜索可能的配置文件。
		// md5:421551127aec1652
# <翻译结束>


<原文开始>
// If it cannot find the filePath of `file`, it formats and returns a detailed error.
<原文结束>

# <翻译开始>
	// 如果无法找到`file`的filePath，它会格式化并返回一个详细的错误。 md5:4aed299684f45971
# <翻译结束>

