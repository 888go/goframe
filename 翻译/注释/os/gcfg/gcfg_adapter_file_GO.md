
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
// AdapterFile implements interface Adapter using file.
<原文结束>

# <翻译开始>
// AdapterFile 实现了使用文件的 Adapter 接口。. md5:c0f0e0b1d4b217fd
# <翻译结束>


<原文开始>
// Default configuration file name.
<原文结束>

# <翻译开始>
// 默认配置文件名。. md5:30af4aed147cf623
# <翻译结束>


<原文开始>
// The pared JSON objects for configuration files.
<原文结束>

# <翻译开始>
// 配置文件中使用的简化JSON对象。. md5:dc2b385b92cd7edc
# <翻译结束>


<原文开始>
// Whether it does violence check in value index searching. It affects the performance when set true(false in default).
<原文结束>

# <翻译开始>
// 是否在值索引搜索中进行暴力检查。设置为true（默认为false）会影响性能。. md5:b2ea5ca1ded97be3
# <翻译结束>


<原文开始>
// commandEnvKeyForFile is the configuration key for command argument or environment configuring file name.
<原文结束>

# <翻译开始>
// commandEnvKeyForFile是用于命令参数或环境配置文件名的配置键。. md5:32e2ea36b81b5269
# <翻译结束>


<原文开始>
// commandEnvKeyForPath is the configuration key for command argument or environment configuring directory path.
<原文结束>

# <翻译开始>
// commandEnvKeyForPath 是用于命令参数或配置目录路径的配置键。. md5:0e1e71d5290a8c3c
# <翻译结束>


<原文开始>
// All supported file types suffixes.
<原文结束>

# <翻译开始>
// 所支持的文件类型后缀。. md5:3609c8928b780170
# <翻译结束>


<原文开始>
// Instances map containing configuration instances.
<原文结束>

# <翻译开始>
// Instances映射，其中包含配置实例。. md5:df7e552f8e970f97
# <翻译结束>


<原文开始>
// Customized configuration content.
<原文结束>

# <翻译开始>
// 定制化配置内容。. md5:e408d212ab61e310
# <翻译结束>


<原文开始>
// Prefix array for trying searching in resource manager.
<原文结束>

# <翻译开始>
// 用于在资源管理器中尝试搜索的前缀数组。. md5:f69485b110ee7be3
# <翻译结束>


<原文开始>
// Prefix array for trying searching in local system.
<原文结束>

# <翻译开始>
// 前缀数组，用于在本地系统中尝试搜索。. md5:51a8f1255f95f3fc
# <翻译结束>


<原文开始>
// NewAdapterFile returns a new configuration management object.
// The parameter `file` specifies the default configuration file name for reading.
<原文结束>

# <翻译开始>
// NewAdapterFile 返回一个新的配置管理对象。
// 参数 `file` 指定了默认的配置文件读取名称。
// md5:52ab633a98562ceb
# <翻译结束>


<原文开始>
// Custom default configuration file name from command line or environment.
<原文结束>

# <翻译开始>
// 从命令行或环境变量中获取自定义的默认配置文件名。. md5:d43279fee761ac4d
# <翻译结束>


<原文开始>
// Customized dir path from env/cmd.
<原文结束>

# <翻译开始>
// 从环境变量或命令行自定义的目录路径。. md5:8cfcbca968e23c5b
# <翻译结束>


<原文开始>
		// ================================================================================
		// Automatic searching directories.
		// It does not affect adapter object cresting if these directories do not exist.
		// ================================================================================
<原文结束>

# <翻译开始>
// =================================================================================
// 自动搜索目录。
// 如果这些目录不存在，不影响适配器对象的创建。
// =================================================================================
// md5:08a226598ce0311e
# <翻译结束>


<原文开始>
// Dir path of working dir.
<原文结束>

# <翻译开始>
// Dir 是工作目录的路径。. md5:0fba211853ea97a0
# <翻译结束>


<原文开始>
// Dir path of main package.
<原文结束>

# <翻译开始>
// 主包的目录路径。. md5:a4d2802779172abe
# <翻译结束>


<原文开始>
// SetViolenceCheck sets whether to perform hierarchical conflict checking.
// This feature needs to be enabled when there is a level symbol in the key name.
// It is off in default.
//
// Note that, turning on this feature is quite expensive, and it is not recommended
// allowing separators in the key names. It is best to avoid this on the application side.
<原文结束>

# <翻译开始>
// SetViolenceCheck 设置是否执行层次冲突检查。
// 当键名中包含级别符号时，需要启用此功能。默认情况下禁用。
// 
// 注意，开启此功能的开销较大，并不建议在键名中使用分隔符。最好在应用层面上避免这种情况。
// md5:5009f694ccd4efc0
# <翻译结束>


<原文开始>
// SetFileName sets the default configuration file name.
<原文结束>

# <翻译开始>
// SetFileName 设置默认的配置文件名。. md5:b540171ead70ddf8
# <翻译结束>


<原文开始>
// GetFileName returns the default configuration file name.
<原文结束>

# <翻译开始>
// GetFileName 返回默认的配置文件名。. md5:d13e3bd27526f03d
# <翻译结束>


<原文开始>
// Get retrieves and returns value by specified `pattern`.
// It returns all values of current Json object if `pattern` is given empty or string ".".
// It returns nil if no value found by `pattern`.
//
// We can also access slice item by its index number in `pattern` like:
// "list.10", "array.0.name", "array.0.1.id".
//
// It returns a default value specified by `def` if value for `pattern` is not found.
<原文结束>

# <翻译开始>
// Get通过指定的`pattern`获取并返回值。如果`pattern`为空或为"."，则返回当前Json对象的所有值。如果没有找到匹配`pattern`的值，它将返回nil。
// 
// 我们也可以通过在`pattern`中使用索引来访问切片项，例如："list.10"，"array.0.name"，"array.0.1.id"。
// 
// 如果没有找到与`pattern`匹配的值，它将返回由`def`指定的默认值。
// md5:8a88d01912ac6218
# <翻译结束>


<原文开始>
// Set sets value with specified `pattern`.
// It supports hierarchical data access by char separator, which is '.' in default.
// It is commonly used for updates certain configuration value in runtime.
// Note that, it is not recommended using `Set` configuration at runtime as the configuration would be
// automatically refreshed if underlying configuration file changed.
<原文结束>

# <翻译开始>
// Set 使用指定的 `pattern` 设置值。
// 它支持通过字符分隔符（默认为`.`）进行层次数据访问。
// 这通常用于在运行时更新特定配置值。
// 请注意，不建议在运行时使用 `Set` 配置，因为如果底层配置文件更改，配置会自动刷新。
// md5:65992c2815af747e
# <翻译结束>


<原文开始>
// Data retrieves and returns all configuration data as map type.
<原文结束>

# <翻译开始>
// Data 获取并以映射类型返回所有配置数据。. md5:2a92e8bbe7388f01
# <翻译结束>


<原文开始>
// MustGet acts as function Get, but it panics if error occurs.
<原文结束>

# <翻译开始>
// MustGet 行为类似于函数 Get，但如果发生错误时会引发 panic。. md5:b1d3af83a52fd248
# <翻译结束>


<原文开始>
// Clear removes all parsed configuration files content cache,
// which will force reload configuration content from file.
<原文结束>

# <翻译开始>
// Clear 清除所有解析的配置文件内容缓存，这将强制重新从文件加载配置内容。
// md5:5868c636ce62cb14
# <翻译结束>


<原文开始>
// Dump prints current Json object with more manually readable.
<原文结束>

# <翻译开始>
// Dump 打印当前的Json对象，使其更便于人工阅读。. md5:c8c6bbdb40fa6383
# <翻译结束>


<原文开始>
// Available checks and returns whether configuration of given `file` is available.
<原文结束>

# <翻译开始>
// 可用检查并返回给定`file`的配置是否可用。. md5:d915d3cb575cbd5b
# <翻译结束>


<原文开始>
// Custom configuration content exists.
<原文结束>

# <翻译开始>
// 存在自定义配置内容。. md5:50d226a12b07427d
# <翻译结束>


<原文开始>
// Configuration file exists in system path.
<原文结束>

# <翻译开始>
// 配置文件存在于系统路径中。. md5:a32283fd4eff7ddf
# <翻译结束>


<原文开始>
// autoCheckAndAddMainPkgPathToSearchPaths automatically checks and adds directory path of package main
// to the searching path list if it's currently in development environment.
<原文结束>

# <翻译开始>
// autoCheckAndAddMainPkgPathToSearchPaths 自动检查并添加当前开发环境中的"main"包目录路径到搜索路径列表中。
// md5:4a1366fa2d1d98ab
# <翻译结束>


<原文开始>
// getJson returns a *gjson.Json object for the specified `file` content.
// It would print error if file reading fails. It returns nil if any error occurs.
<原文结束>

# <翻译开始>
// getJson 为指定的`file`内容返回一个*gjson.Json*对象。
// 如果文件读取失败，它会打印错误。如果发生任何错误，它将返回nil。
// md5:ffbc3e1a6ff12753
# <翻译结束>


<原文开始>
// It uses json map to cache specified configuration file content.
<原文结束>

# <翻译开始>
// 它使用json映射来缓存指定的配置文件内容。. md5:70b9eac1f3ac38b4
# <翻译结束>


<原文开始>
// The configured content can be any kind of data type different from its file type.
<原文结束>

# <翻译开始>
// 配置的内容可以是与其文件类型不同的任何数据类型。. md5:11fb8ecd6511ef10
# <翻译结束>


<原文开始>
// Note that the underlying configuration json object operations are concurrent safe.
<原文结束>

# <翻译开始>
// 注意，底层的配置json对象操作是并发安全的。. md5:2cd371ca691286f9
# <翻译结束>


<原文开始>
		// Add monitor for this configuration file,
		// any changes of this file will refresh its cache in Config object.
<原文结束>

# <翻译开始>
// 为这个配置文件添加监控，
// 该文件的任何更改都会刷新Config对象中的缓存。
// md5:8520fe419f2d8cc1
# <翻译结束>

