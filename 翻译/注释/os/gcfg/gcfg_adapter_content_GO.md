
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
// AdapterContent implements interface Adapter using content.
// The configuration content supports the coding types as package `gjson`.
<原文结束>

# <翻译开始>
// AdapterContent 使用内容实现 Adapter 接口。
// 配置内容支持的编码类型与 `gjson` 包中的编码类型相同。
// md5:81e39ba9e6de51fa
# <翻译结束>


<原文开始>
// The pared JSON object for configuration content, type: *gjson.Json.
<原文结束>

# <翻译开始>
// 配置内容的修剪过的 JSON 对象，类型：*gjson.Json。. md5:379162a7a5ad528f
# <翻译结束>


<原文开始>
// NewAdapterContent returns a new configuration management object using custom content.
// The parameter `content` specifies the default configuration content for reading.
<原文结束>

# <翻译开始>
// NewAdapterContent 返回一个使用自定义内容的新配置管理对象。
// 参数 `content` 指定用于读取的默认配置内容。
// md5:efafcabf61d7087b
# <翻译结束>


<原文开始>
// SetContent sets customized configuration content for specified `file`.
// The `file` is unnecessary param, default is DefaultConfigFile.
<原文结束>

# <翻译开始>
// SetContent 为指定的`file`设置自定义配置内容。
// `file`是可选参数，默认值为DefaultConfigFile。
// md5:49ae38cf671e3b96
# <翻译结束>


<原文开始>
// Available checks and returns the backend configuration service is available.
// The optional parameter `resource` specifies certain configuration resource.
//
// Note that this function does not return error as it just does simply check for
// backend configuration service.
<原文结束>

# <翻译开始>
// 可用性检查并返回后端配置服务是否可用。
// 可选参数 `resource` 指定特定的配置资源。
// 
// 请注意，此函数不会返回错误，因为它只是简单地检查后端配置服务。
// md5:79f955eb2fcdd137
# <翻译结束>


<原文开始>
// Get retrieves and returns value by specified `pattern` in current resource.
// Pattern like:
// "x.y.z" for map item.
// "x.0.y" for slice item.
<原文结束>

# <翻译开始>
// Get 通过当前资源中指定的`pattern`获取并返回值。
// `pattern`示例：
// "x.y.z" 用于map中的条目。
// "x.0.y" 用于切片中的条目。
// md5:39b9171603468968
# <翻译结束>


<原文开始>
// Data retrieves and returns all configuration data in current resource as map.
// Note that this function may lead lots of memory usage if configuration data is too large,
// you can implement this function if necessary.
<原文结束>

# <翻译开始>
// Data 获取并以映射的形式返回当前资源中的所有配置数据。
// 注意，如果配置数据量过大，此函数可能会占用大量内存。
// 如有需要，你可以根据实际情况实现这个函数。
// md5:19dfa88d9aa6ece5
# <翻译结束>

