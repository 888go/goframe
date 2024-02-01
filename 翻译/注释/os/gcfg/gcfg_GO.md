
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
// Package gcfg provides reading, caching and managing for configuration.
<原文结束>

# <翻译开始>
// Package gcfg 提供了配置的读取、缓存和管理功能。
# <翻译结束>


<原文开始>
// Config is the configuration management object.
<原文结束>

# <翻译开始>
// Config 是配置管理对象。
# <翻译结束>


<原文开始>
// DefaultName is the default instance name for instance usage.
<原文结束>

# <翻译开始>
// DefaultName 是用于实例使用的默认实例名称。
# <翻译结束>


<原文开始>
// DefaultConfigFile is the default configuration file name.
<原文结束>

# <翻译开始>
// DefaultConfigFile 是默认的配置文件名称。
# <翻译结束>


<原文开始>
// New creates and returns a Config object with default adapter of AdapterFile.
<原文结束>

# <翻译开始>
// New 创建并返回一个 Config 对象，其默认适配器为 AdapterFile。
# <翻译结束>


<原文开始>
// NewWithAdapter creates and returns a Config object with given adapter.
<原文结束>

# <翻译开始>
// NewWithAdapter 使用给定的适配器创建并返回一个Config对象。
# <翻译结束>


<原文开始>
// Instance returns an instance of Config with default settings.
// The parameter `name` is the name for the instance. But very note that, if the file "name.toml"
// exists in the configuration directory, it then sets it as the default configuration file. The
// toml file type is the default configuration file type.
<原文结束>

# <翻译开始>
// Instance 返回一个使用默认设置的 Config 实例。
// 参数 `name` 是该实例的名称。但请注意，如果配置目录中存在名为 "name.toml" 的文件，
// 则将其设置为默认配置文件。toml 文件类型是默认的配置文件类型。
# <翻译结束>


<原文开始>
// SetAdapter sets the adapter of current Config object.
<原文结束>

# <翻译开始>
// SetAdapter 设置当前 Config 对象的适配器。
# <翻译结束>


<原文开始>
// GetAdapter returns the adapter of current Config object.
<原文结束>

# <翻译开始>
// GetAdapter 返回当前 Config 对象的适配器。
# <翻译结束>


<原文开始>
// Available checks and returns the configuration service is available.
// The optional parameter `pattern` specifies certain configuration resource.
//
// It returns true if configuration file is present in default AdapterFile, or else false.
// Note that this function does not return error as it just does simply check for backend configuration service.
<原文结束>

# <翻译开始>
// Available 检查并返回配置服务是否可用。
// 可选参数 `pattern` 指定了特定的配置资源。
//
// 如果配置文件存在于默认的 AdapterFile 中，则返回 true，否则返回 false。
// 注意，此函数不会返回错误，因为它只是简单地检查后端配置服务是否存在。
# <翻译结束>


<原文开始>
// Get retrieves and returns value by specified `pattern`.
// It returns all values of current Json object if `pattern` is given empty or string ".".
// It returns nil if no value found by `pattern`.
//
// It returns a default value specified by `def` if value for `pattern` is not found.
<原文结束>

# <翻译开始>
// Get 方法通过指定的`pattern`获取并返回值。
// 如果给定的`pattern`为空字符串或"."，则返回当前Json对象的所有值。
// 若未找到通过`pattern`匹配到的值，则返回nil。
//
// 当通过`pattern`未能找到对应的值时，将返回由`def`指定的默认值。
# <翻译结束>


<原文开始>
// GetWithEnv returns the configuration value specified by pattern `pattern`.
// If the configuration value does not exist, then it retrieves and returns the environment value specified by `key`.
// It returns the default value `def` if none of them exists.
//
// Fetching Rules: Environment arguments are in uppercase format, eg: GF_PACKAGE_VARIABLE.
<原文结束>

# <翻译开始>
// GetWithEnv 返回通过模式`pattern`指定的配置值。
// 如果配置值不存在，则获取并返回由`key`指定的环境变量值。
// 若两者都不存在，则返回默认值 `def`。
//
// 获取规则：环境变量参数采用大写格式，例如：GF_PACKAGE_VARIABLE。
# <翻译结束>


<原文开始>
// GetWithCmd returns the configuration value specified by pattern `pattern`.
// If the configuration value does not exist, then it retrieves and returns the command line option specified by `key`.
// It returns the default value `def` if none of them exists.
//
// Fetching Rules: Command line arguments are in lowercase format, eg: gf.package.variable.
<原文结束>

# <翻译开始>
// GetWithCmd 根据模式 `pattern` 返回配置值。
// 如果配置值不存在，则获取并返回由 `key` 指定的命令行选项。
// 若两者都不存在，则返回默认值 `def`。
//
// 获取规则：命令行参数采用小写格式，例如：gf.package.variable。
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
// MustGetWithEnv acts as function GetWithEnv, but it panics if error occurs.
<原文结束>

# <翻译开始>
// MustGetWithEnv 行为类似于函数 GetWithEnv，但是当发生错误时它会触发panic。
# <翻译结束>


<原文开始>
// MustGetWithCmd acts as function GetWithCmd, but it panics if error occurs.
<原文结束>

# <翻译开始>
// MustGetWithCmd 的行为与函数 GetWithCmd 相同，但当出现错误时它会触发 panic。
# <翻译结束>


<原文开始>
// MustData acts as function Data, but it panics if error occurs.
<原文结束>

# <翻译开始>
// MustData 的行为与函数 Data 相同，但是当发生错误时，它会引发 panic（异常）。
# <翻译结束>

