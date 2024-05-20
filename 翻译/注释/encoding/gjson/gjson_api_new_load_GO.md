
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
// New creates a Json object with any variable type of `data`, but `data` should be a map
// or slice for data access reason, or it will make no sense.
//
// The parameter `safe` specifies whether using this Json object in concurrent-safe context,
// which is false in default.
<原文结束>

# <翻译开始>
// New 使用任何类型的`data`创建一个Json对象，但为了数据访问的原因，`data`应该是map或slice，
// 否则将失去意义。
//
// 参数`safe`指定是否在并发安全的上下文中使用此Json对象，默认值为false。
// md5:b84f401db24e69d8
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
// NewWithTag 创建一个Json对象，可以包含任何类型的`data`，但出于数据访问的原因，`data`应该是一个map或切片，否则将没有意义。
// 
// 参数`tags`用于指定结构体转换为map的优先标签，多个标签之间用逗号分隔。
// 
// 参数`safe`表示是否在并发安全上下文中使用这个Json对象，默认为false。
// md5:2558f08f4f082a16
# <翻译结束>


<原文开始>
// NewWithOptions creates a Json object with any variable type of `data`, but `data` should be a map
// or slice for data access reason, or it will make no sense.
<原文结束>

# <翻译开始>
// NewWithOptions使用任何类型的'data'创建一个Json对象，但出于数据访问的原因，`data`应该是map或切片，否则将没有意义。
// md5:48be1828a6556518
# <翻译结束>


<原文开始>
// Load loads content from specified file `path`, and creates a Json object from its content.
<原文结束>

# <翻译开始>
// Load 从指定的文件`path`加载内容，并根据其内容创建一个Json对象。. md5:fc26d8aa3d537173
# <翻译结束>


<原文开始>
// LoadWithOptions creates a Json object from given JSON format content and options.
<原文结束>

# <翻译开始>
// LoadWithOptions 根据给定的 JSON 格式内容和选项创建一个 Json 对象。. md5:77290b5f994f3ff1
# <翻译结束>


<原文开始>
// LoadJson creates a Json object from given JSON format content.
<原文结束>

# <翻译开始>
// LoadJson 从给定的JSON格式内容创建一个Json对象。. md5:1f41cbc0a35bd390
# <翻译结束>


<原文开始>
// LoadXml creates a Json object from given XML format content.
<原文结束>

# <翻译开始>
// LoadXml 从给定的XML格式内容创建一个Json对象。. md5:a170d56aa371a2bb
# <翻译结束>


<原文开始>
// LoadIni creates a Json object from given INI format content.
<原文结束>

# <翻译开始>
// LoadIni 从给定的INI格式内容创建一个Json对象。. md5:bf3225da0be4c26b
# <翻译结束>


<原文开始>
// LoadYaml creates a Json object from given YAML format content.
<原文结束>

# <翻译开始>
// LoadYaml 根据给定的 YAML 格式内容创建一个 Json 对象。. md5:d810aac213716b5a
# <翻译结束>


<原文开始>
// LoadToml creates a Json object from given TOML format content.
<原文结束>

# <翻译开始>
// LoadToml 从给定的TOML格式内容创建一个Json对象。. md5:a27ac84d2a7e5a70
# <翻译结束>


<原文开始>
// LoadProperties creates a Json object from given TOML format content.
<原文结束>

# <翻译开始>
// LoadProperties 从给定的TOML格式内容创建一个Json对象。. md5:aacff07e57605d82
# <翻译结束>


<原文开始>
// LoadContent creates a Json object from given content, it checks the data type of `content`
// automatically, supporting data content type as follows:
// JSON, XML, INI, YAML and TOML.
<原文结束>

# <翻译开始>
// LoadContent 根据给定的内容创建一个Json对象，它会自动检查`content`的数据类型，
// 支持如下数据内容类型：
// JSON、XML、INI、YAML和TOML。
// md5:e930374f4ac3b32e
# <翻译结束>


<原文开始>
// LoadContentType creates a Json object from given type and content,
// supporting data content type as follows:
// JSON, XML, INI, YAML and TOML.
<原文结束>

# <翻译开始>
// LoadContentType 根据给定的类型和内容创建一个 JSON 对象，支持以下数据内容类型：
// JSON, XML, INI, YAML 和 TOML.
// md5:7db5bd0b429fea01
# <翻译结束>


<原文开始>
// IsValidDataType checks and returns whether given `dataType` a valid data type for loading.
<原文结束>

# <翻译开始>
// IsValidDataType 检查并返回给定的 `dataType` 是否是用于加载的有效数据类型。. md5:3cc6cab5a2631a3e
# <翻译结束>


<原文开始>
// doLoadContent creates a Json object from given content.
// It supports data content type as follows:
// JSON, XML, INI, YAML and TOML.
<原文结束>

# <翻译开始>
// doLoadContent 从给定内容创建一个Json对象。
// 它支持以下数据内容类型：
// JSON、XML、INI、YAML和TOML。
// md5:a1daf6666c64b0bc
# <翻译结束>


<原文开始>
// checkDataType automatically checks and returns the data type for `content`.
// Note that it uses regular expression for loose checking, you can use LoadXXX/LoadContentType
// functions to load the content for certain content type.
<原文结束>

# <翻译开始>
// checkDataType 会自动检查并返回`content`的数据类型。
// 注意，它使用正则表达式进行宽松的检查，你可以根据需要使用LoadXXX/LoadContentType
// 函数来为特定内容类型加载内容。
// md5:faa69696c8f02af2
# <翻译结束>


<原文开始>
// Must contain "[xxx]" section.
<原文结束>

# <翻译开始>
// 必须包含"xxx"部分。. md5:6dc6d0a6d417b6a6
# <翻译结束>

