
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
// AdapterFile implements interface Adapter using file.
<原文结束>

# <翻译开始>
// AdapterFile实现了使用文件的Adapter接口。
# <翻译结束>


<原文开始>
// Default configuration file name.
<原文结束>

# <翻译开始>
// 默认配置文件名称。
# <翻译结束>







<原文开始>
// The pared JSON objects for configuration files.
<原文结束>

# <翻译开始>
// 这是用于配置文件的解析后的JSON对象。
# <翻译结束>


<原文开始>
// Whether it does violence check in value index searching. It affects the performance when set true(false in default).
<原文结束>

# <翻译开始>
// 是否在值索引搜索时进行暴力检查。当设置为true（默认为false）时，会影响性能。
# <翻译结束>


<原文开始>
// commandEnvKeyForFile is the configuration key for command argument or environment configuring file name.
<原文结束>

# <翻译开始>
// commandEnvKeyForFile 是用于配置命令行参数或环境配置文件名的配置键。
# <翻译结束>


<原文开始>
// commandEnvKeyForPath is the configuration key for command argument or environment configuring directory path.
<原文结束>

# <翻译开始>
// commandEnvKeyForPath 是用于配置命令参数或环境目录路径的配置键。
# <翻译结束>


<原文开始>
// All supported file types suffixes.
<原文结束>

# <翻译开始>
// 所有支持的文件类型后缀。
# <翻译结束>


<原文开始>
// Instances map containing configuration instances.
<原文结束>

# <翻译开始>
// Instances：包含配置实例的映射（map）。
# <翻译结束>


<原文开始>
// Customized configuration content.
<原文结束>

# <翻译开始>
// 自定义配置内容
# <翻译结束>


<原文开始>
// Prefix array for trying searching in resource manager.
<原文结束>

# <翻译开始>
// 前缀数组，用于在资源管理器中尝试搜索。
# <翻译结束>


<原文开始>
// Prefix array for trying searching in local system.
<原文结束>

# <翻译开始>
// 前缀数组，用于尝试在本地系统中搜索。
# <翻译结束>


<原文开始>
// NewAdapterFile returns a new configuration management object.
// The parameter `file` specifies the default configuration file name for reading.
<原文结束>

# <翻译开始>
// NewAdapterFile 返回一个新的配置管理对象。
// 参数`file`指定了用于读取的默认配置文件名。
# <翻译结束>


<原文开始>
// Custom default configuration file name from command line or environment.
<原文结束>

# <翻译开始>
// 从命令行或环境变量自定义默认配置文件名。
# <翻译结束>


<原文开始>
// Customized dir path from env/cmd.
<原文结束>

# <翻译开始>
// 从环境变量/命令行自定义目录路径。
# <翻译结束>


<原文开始>
		// ================================================================================
		// Automatic searching directories.
		// It does not affect adapter object cresting if these directories do not exist.
		// ================================================================================
<原文结束>

# <翻译开始>
// ================================================================================
// 自动搜索目录
// 如果这些目录不存在，也不会影响适配器对象的创建。
// ================================================================================
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
// SetViolenceCheck 设置是否进行层级冲突检查。
// 当键名中存在层级符号时，需要开启此功能。默认情况下，该功能是关闭的。
//
// 注意，开启此特性代价较高，且不建议在键名中允许分隔符。最好在应用层面避免这种情况。
# <翻译结束>


<原文开始>
// SetFileName sets the default configuration file name.
<原文结束>

# <翻译开始>
// SetFileName 设置默认的配置文件名。
# <翻译结束>


<原文开始>
// GetFileName returns the default configuration file name.
<原文结束>

# <翻译开始>
// GetFileName 返回默认配置文件名称。
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
// Get 方法通过指定的`pattern`获取并返回值。
// 如果`pattern`为空字符串或"."，则返回当前Json对象的所有值。
// 若通过`pattern`未找到任何值，则返回nil。
//
// 我们还可以通过在`pattern`中使用索引号访问切片元素，例如：
// "list.10", "array.0.name", "array.0.1.id"。
//
// 如果根据`pattern`未能找到对应的值，则返回由`def`指定的默认值。
# <翻译结束>


<原文开始>
// Set sets value with specified `pattern`.
// It supports hierarchical data access by char separator, which is '.' in default.
// It is commonly used for updates certain configuration value in runtime.
// Note that, it is not recommended using `Set` configuration at runtime as the configuration would be
// automatically refreshed if underlying configuration file changed.
<原文结束>

# <翻译开始>
// Set 通过指定的`pattern`设置值。
// 它支持使用字符分隔符（默认为'. '）进行层级数据访问。
// 通常用于在运行时更新特定配置值。
// 注意，不建议在运行时使用`Set`方法来配置，因为如果底层配置文件发生更改，
// 配置将会自动刷新。因此，直接运行时设置可能不会持久生效。
# <翻译结束>


<原文开始>
// Data retrieves and returns all configuration data as map type.
<原文结束>

# <翻译开始>
// Data 函数获取并以 map 类型返回所有配置数据。
# <翻译结束>


<原文开始>
// MustGet acts as function Get, but it panics if error occurs.
<原文结束>

# <翻译开始>
// MustGet 行为类似于函数 Get，但在发生错误时会触发 panic。
# <翻译结束>


<原文开始>
// Clear removes all parsed configuration files content cache,
// which will force reload configuration content from file.
<原文结束>

# <翻译开始>
// 清除所有已解析的配置文件内容缓存，
// 这将强制从文件重新加载配置内容。
# <翻译结束>


<原文开始>
// Dump prints current Json object with more manually readable.
<原文结束>

# <翻译开始>
// Dump 打印当前Json对象，使其更易于人工阅读。
# <翻译结束>


<原文开始>
// Available checks and returns whether configuration of given `file` is available.
<原文结束>

# <翻译开始>
// Available 检查并返回给定 `file` 配置是否可用。
# <翻译结束>


<原文开始>
// Custom configuration content exists.
<原文结束>

# <翻译开始>
// 自定义配置内容存在。
# <翻译结束>


<原文开始>
// Configuration file exists in system path.
<原文结束>

# <翻译开始>
// 配置文件存在于系统路径中。
# <翻译结束>


<原文开始>
// autoCheckAndAddMainPkgPathToSearchPaths automatically checks and adds directory path of package main
// to the searching path list if it's currently in development environment.
<原文结束>

# <翻译开始>
// autoCheckAndAddMainPkgPathToSearchPaths 自动检测并添加 main 包的目录路径到搜索路径列表中，
// 如果当前处于开发环境的话。
# <翻译结束>


<原文开始>
// getJson returns a *gjson.Json object for the specified `file` content.
// It would print error if file reading fails. It returns nil if any error occurs.
<原文结束>

# <翻译开始>
// getJson 函数返回指定 `file` 内容对应的 *gjson.Json 对象。
// 如果文件读取失败，会打印错误信息。若发生任何错误，将返回 nil。
# <翻译结束>


<原文开始>
// It uses json map to cache specified configuration file content.
<原文结束>

# <翻译开始>
// 它使用json映射来缓存指定配置文件的内容。
# <翻译结束>


<原文开始>
// The configured content can be any kind of data type different from its file type.
<原文结束>

# <翻译开始>
// 配置的内容可以是与文件类型不同的任何数据类型。
# <翻译结束>


<原文开始>
// Note that the underlying configuration json object operations are concurrent safe.
<原文结束>

# <翻译开始>
// 注意，底层的配置json对象操作是线程安全的。
# <翻译结束>


<原文开始>
		// Add monitor for this configuration file,
		// any changes of this file will refresh its cache in Config object.
<原文结束>

# <翻译开始>
// 添加对这个配置文件的监控，
// 当该文件有任何变化时，都会在Config对象中刷新其缓存。
# <翻译结束>


<原文开始>
// Searching path array.
<原文结束>

# <翻译开始>
// 搜索路径数组
# <翻译结束>


<原文开始>
// Dir path of working dir.
<原文结束>

# <翻译开始>
// Dir：工作目录的路径。
# <翻译结束>


<原文开始>
// Dir path of main package.
<原文结束>

# <翻译开始>
// Dir：主包的路径。
# <翻译结束>


<原文开始>
// Dir path of binary.
<原文结束>

# <翻译开始>
// Dir 二进制文件的路径。
# <翻译结束>

