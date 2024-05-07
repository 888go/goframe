
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
// SetPath sets the configuration `directory` path for file search.
// The parameter `path` can be absolute or relative `directory` path,
// but absolute `directory` path is strongly recommended.
//
// Note that this parameter is a path to a directory not a file.
<原文结束>

# <翻译开始>
// SetPath 设置配置文件搜索的 `directory` 路径。
// 参数 `path` 可以是绝对路径或相对 `directory` 路径，
// 但强烈建议使用绝对 `directory` 路径。
//
// 注意，此参数是一个目录而非文件的路径。
# <翻译结束>












<原文开始>
// AddPath adds an absolute or relative `directory` path to the search paths.
//
// Note that this parameter is paths to a directories not files.
<原文结束>

# <翻译开始>
// AddPath 将绝对或相对的 `directory` 路径添加到搜索路径中。
//
// 注意，此参数是目录而非文件的路径。
# <翻译结束>


<原文开始>
// doAddPath adds an absolute or relative `directory` path to the search paths.
<原文结束>

# <翻译开始>
// doAddPath 将绝对或相对 `directory` 路径添加到搜索路径中。
# <翻译结束>


<原文开始>
	// It firstly checks the resource manager,
	// and then checks the filesystem for the path.
<原文结束>

# <翻译开始>
// 首先检查资源管理器，
// 然后在文件系统中检查路径。
# <翻译结束>


<原文开始>
// GetPaths returns the searching directory path array of current configuration manager.
<原文结束>

# <翻译开始>
// GetPaths 返回当前配置管理器的搜索目录路径切片。
# <翻译结束>


<原文开始>
// doGetFilePath returns the absolute configuration file path for the given filename by `file`.
// If `file` is not passed, it returns the configuration file path of the default name.
// It returns an empty `path` string and an error if the given `file` does not exist.
<原文结束>

# <翻译开始>
// doGetFilePath 根据传入的文件名 `file` 返回其绝对配置文件路径。
// 若未传入 `file`，则返回默认文件名的配置文件路径。
// 若给定的 `file` 不存在，则返回空字符串 `path` 及错误信息。
# <翻译结束>


<原文开始>
// Searching resource manager.
<原文结束>

# <翻译开始>
// 搜索资源管理器
# <翻译结束>


<原文开始>
// Searching local file system.
<原文结束>

# <翻译开始>
// 正在搜索本地文件系统。
# <翻译结束>


<原文开始>
// GetFilePath returns the absolute configuration file path for the given filename by `file`.
// If `file` is not passed, it returns the configuration file path of the default name.
// It returns an empty `path` string and an error if the given `file` does not exist.
<原文结束>

# <翻译开始>
// GetFilePath 函数根据传入的文件名`file`返回该文件的绝对配置文件路径。
// 若未传递`file`参数，则返回默认文件名的配置文件路径。
// 如果给定的`file`不存在，则返回一个空字符串`path`及错误信息。
# <翻译结束>


<原文开始>
		// If it's not using default configuration or its configuration file is not available,
		// it searches the possible configuration file according to the name and all supported
		// file types.
<原文结束>

# <翻译开始>
// 如果不是使用默认配置，或者其配置文件不可用，
// 则根据名称和所有支持的文件类型搜索可能的配置文件。
# <翻译结束>


<原文开始>
// If it cannot find the filePath of `file`, it formats and returns a detailed error.
<原文结束>

# <翻译开始>
// 如果无法找到`file`的filePath，它将格式化并返回一个详细的错误。
# <翻译结束>


<原文开始>
// Should be a directory.
<原文结束>

# <翻译开始>
// 应该是一个目录。
# <翻译结束>


<原文开始>
// Repeated path check.
<原文结束>

# <翻译开始>
// 重复路径检查。
# <翻译结束>

