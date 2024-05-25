
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
// Package gcfg provides reading, caching and managing for configuration.
<原文结束>

# <翻译开始>
// 包gcfg提供了配置的读取、缓存和管理功能。 md5:5ae504d1379cd99a
# <翻译结束>


<原文开始>
// Config is the configuration management object.
<原文结束>

# <翻译开始>
// Config 是配置管理对象。 md5:1ef57338c678e400
# <翻译结束>


<原文开始>
// DefaultName is the default instance name for instance usage.
<原文结束>

# <翻译开始>
// DefaultName 是实例使用的默认实例名称。 md5:4736f3b4285b6846
# <翻译结束>


<原文开始>
// DefaultConfigFile is the default configuration file name.
<原文结束>

# <翻译开始>
// DefaultConfigFile 是默认的配置文件名。 md5:b558e9c92a774f9a
# <翻译结束>


<原文开始>
// New creates and returns a Config object with default adapter of AdapterFile.
<原文结束>

# <翻译开始>
// New 创建并返回一个 Config 对象，其默认适配器为 AdapterFile。 md5:52cd678118524272
# <翻译结束>


<原文开始>
// NewWithAdapter creates and returns a Config object with given adapter.
<原文结束>

# <翻译开始>
// NewWithAdapter使用给定的适配器创建并返回一个Config对象。 md5:9ddaae0ddb0e0297
# <翻译结束>


<原文开始>
// Instance returns an instance of Config with default settings.
// The parameter `name` is the name for the instance. But very note that, if the file "name.toml"
// exists in the configuration directory, it then sets it as the default configuration file. The
// toml file type is the default configuration file type.
<原文结束>

# <翻译开始>
// Instance 返回一个具有默认设置的 Config 实例。
// 参数 `name` 是该实例的名称。但请注意，如果配置目录中存在文件 "name.toml"，
// 则将其设置为默认配置文件。TOML 文件类型是默认的配置文件类型。
// md5:4164ff567a8c8c31
# <翻译结束>


<原文开始>
// SetAdapter sets the adapter of current Config object.
<原文结束>

# <翻译开始>
// SetAdapter 设置当前 Config 对象的适配器。 md5:8d00d377baafeb01
# <翻译结束>


<原文开始>
// GetAdapter returns the adapter of current Config object.
<原文结束>

# <翻译开始>
// GetAdapter 返回当前Config对象的适配器。 md5:46c003ab367518d8
# <翻译结束>


<原文开始>
// Available checks and returns the configuration service is available.
// The optional parameter `pattern` specifies certain configuration resource.
//
// It returns true if configuration file is present in default AdapterFile, or else false.
// Note that this function does not return error as it just does simply check for backend configuration service.
<原文结束>

# <翻译开始>
// 可用性检查并返回配置服务是否可用。
// 可选参数 `pattern` 指定某些配置资源。
// 
// 如果默认AdapterFile中存在配置文件，则返回true，否则返回false。
// 请注意，此函数不会返回错误，因为它只是简单地检查后端配置服务。
// md5:771d98d194158bc1
# <翻译结束>


<原文开始>
// Get retrieves and returns value by specified `pattern`.
// It returns all values of current Json object if `pattern` is given empty or string ".".
// It returns nil if no value found by `pattern`.
//
// It returns a default value specified by `def` if value for `pattern` is not found.
<原文结束>

# <翻译开始>
// Get 通过指定的`pattern`获取并返回值。
// 如果`pattern`为空字符串或"."，它将返回当前Json对象的所有值。
// 如果根据`pattern`没有找到值，它将返回nil。
//
// 如果没有为`pattern`找到值，它将返回由`def`指定的默认值。
// md5:b10a106fb9d6af41
# <翻译结束>


<原文开始>
// GetWithEnv returns the configuration value specified by pattern `pattern`.
// If the configuration value does not exist, then it retrieves and returns the environment value specified by `key`.
// It returns the default value `def` if none of them exists.
//
// Fetching Rules: Environment arguments are in uppercase format, eg: GF_PACKAGE_VARIABLE.
<原文结束>

# <翻译开始>
// GetWithEnv 根据模式`pattern`返回配置值。
// 如果配置值不存在，那么它会获取并返回由`key`指定的环境变量值。
// 如果两者都不存在，则返回默认值`def`。
//
// 获取规则：环境变量参数以大写格式表示，例如：GF_PACKAGE_VARIABLE。
// md5:d533293fbfbf6350
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
// 如果找不到配置值，它将检索并返回由 `key` 指定的命令行选项。
// 如果它们都不存在，则返回默认值 `def`。
// 
// 获取规则：命令行参数采用小写格式，例如：gf.package.variable。
// md5:2a77887f42041d88
# <翻译结束>


<原文开始>
// Data retrieves and returns all configuration data as map type.
<原文结束>

# <翻译开始>
// Data 获取并以映射类型返回所有配置数据。 md5:2a92e8bbe7388f01
# <翻译结束>


<原文开始>
// MustGet acts as function Get, but it panics if error occurs.
<原文结束>

# <翻译开始>
// MustGet 行为类似于函数 Get，但如果发生错误时会引发 panic。 md5:b1d3af83a52fd248
# <翻译结束>


<原文开始>
// MustGetWithEnv acts as function GetWithEnv, but it panics if error occurs.
<原文结束>

# <翻译开始>
// MustGetWithEnv 作为 GetWithEnv 函数的行为，但如果发生错误，它会引发恐慌。 md5:9f816c41440b51cf
# <翻译结束>


<原文开始>
// MustGetWithCmd acts as function GetWithCmd, but it panics if error occurs.
<原文结束>

# <翻译开始>
// MustGetWithCmd 的行为类似于 GetWithCmd 函数，但如果发生错误，它会直接 panic。 md5:683d24a1f4aceb7b
# <翻译结束>


<原文开始>
// MustData acts as function Data, but it panics if error occurs.
<原文结束>

# <翻译开始>
// MustData 行为类似于函数 Data，但如果发生错误则会引发恐慌。 md5:eb72c1ce036d70b6
# <翻译结束>

