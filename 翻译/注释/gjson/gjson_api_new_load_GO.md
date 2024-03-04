
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
// New creates a Json object with any variable type of `data`, but `data` should be a map
// or slice for data access reason, or it will make no sense.
//
// The parameter `safe` specifies whether using this Json object in concurrent-safe context,
// which is false in default.
<原文结束>

# <翻译开始>
// New 函数通过任意类型的 `data` 创建一个 Json 对象，但 `data` 应为 map 或 slice 类型以保证数据可访问性，否则创建此对象将无实际意义。
//
// 参数 `safe` 指定是否在并发安全的上下文中使用此 Json 对象，默认情况下 `safe` 为 false。
# <翻译结束>


<原文开始>
// NewWithTag creates a Json object with any variable type of `data`, but `data` should be a map
// or slice for data access reason, or it will make no sense.
//
// The parameter `tags` specifies priority tags for struct conversion to map, multiple tags joined
// with char ','.
//
// The parameter `safe` specifies whether using this Json object in concurrent-safe context, which
// is false in default.
<原文结束>

# <翻译开始>
// NewWithTag 创建一个Json对象，其数据类型可以是任意的 `data`，但为了方便数据访问，`data` 应该是一个 map 或 slice，否则将失去意义。
//
// 参数 `tags` 指定了在结构体转为 map 时使用的优先级标签，多个标签使用字符 ',' 连接。
//
// 参数 `safe` 指定了是否在并发安全的上下文中使用此 Json 对象，默认情况下为 false。
# <翻译结束>


<原文开始>
// NewWithOptions creates a Json object with any variable type of `data`, but `data` should be a map
// or slice for data access reason, or it will make no sense.
<原文结束>

# <翻译开始>
// NewWithOptions 创建一个Json对象，其变量类型可以是 `data` 的任意类型，但为了能够访问数据，`data` 应该是一个 map 或 slice，否则将毫无意义。
# <翻译结束>


<原文开始>
// Load loads content from specified file `path`, and creates a Json object from its content.
<原文结束>

# <翻译开始>
// Load 从指定的文件路径`path`加载内容，并根据其内容创建一个Json对象。
# <翻译结束>


<原文开始>
// LoadWithOptions creates a Json object from given JSON format content and options.
<原文结束>

# <翻译开始>
// LoadWithOptions 根据给定的JSON格式内容和选项，创建一个Json对象。
# <翻译结束>


<原文开始>
// LoadJson creates a Json object from given JSON format content.
<原文结束>

# <翻译开始>
// LoadJson 从给定的 JSON 格式内容创建一个 Json 对象。
# <翻译结束>


<原文开始>
// LoadXml creates a Json object from given XML format content.
<原文结束>

# <翻译开始>
// LoadXml 从给定的 XML 格式内容创建一个 Json 对象。
# <翻译结束>


<原文开始>
// LoadIni creates a Json object from given INI format content.
<原文结束>

# <翻译开始>
// LoadIni 从给定的 INI 格式内容创建一个 Json 对象。
# <翻译结束>


<原文开始>
// LoadYaml creates a Json object from given YAML format content.
<原文结束>

# <翻译开始>
// LoadYaml 从给定的YAML格式内容创建一个Json对象。
# <翻译结束>


<原文开始>
// LoadToml creates a Json object from given TOML format content.
<原文结束>

# <翻译开始>
// LoadToml 从给定的TOML格式内容创建一个Json对象。
# <翻译结束>


<原文开始>
// LoadProperties creates a Json object from given TOML format content.
<原文结束>

# <翻译开始>
// LoadProperties 从给定的TOML格式内容创建一个Json对象。
# <翻译结束>


<原文开始>
// LoadContent creates a Json object from given content, it checks the data type of `content`
// automatically, supporting data content type as follows:
// JSON, XML, INI, YAML and TOML.
<原文结束>

# <翻译开始>
// LoadContent 从给定的内容创建一个 Json 对象，它会自动检查 `content` 的数据类型，
// 支持以下数据内容类型：
// JSON、XML、INI、YAML 和 TOML。
# <翻译结束>


<原文开始>
// LoadContentType creates a Json object from given type and content,
// supporting data content type as follows:
// JSON, XML, INI, YAML and TOML.
<原文结束>

# <翻译开始>
// LoadContentType 从给定的类型和内容创建一个 Json 对象，
// 支持以下数据内容类型：
// JSON、XML、INI、YAML 和 TOML。
# <翻译结束>


<原文开始>
// IsValidDataType checks and returns whether given `dataType` a valid data type for loading.
<原文结束>

# <翻译开始>
// IsValidDataType 检查并返回给定的 `dataType` 是否为有效载入数据类型。
# <翻译结束>


<原文开始>
// doLoadContent creates a Json object from given content.
// It supports data content type as follows:
// JSON, XML, INI, YAML and TOML.
<原文结束>

# <翻译开始>
// doLoadContent 从给定的内容创建一个 Json 对象。
// 它支持以下数据内容类型：
// JSON、XML、INI、YAML 和 TOML。
# <翻译结束>


<原文开始>
// checkDataType automatically checks and returns the data type for `content`.
// Note that it uses regular expression for loose checking, you can use LoadXXX/LoadContentType
// functions to load the content for certain content type.
<原文结束>

# <翻译开始>
// checkDataType 自动检查并返回 `content` 的数据类型。
// 注意，它使用正则表达式进行宽松的检查，你可以使用 LoadXXX/LoadContentType
// 函数来按特定内容类型加载内容。
# <翻译结束>


<原文开始>
// Must contain "[xxx]" section.
<原文结束>

# <翻译开始>
// 必须包含 "[xxx]" 部分。
# <翻译结束>






