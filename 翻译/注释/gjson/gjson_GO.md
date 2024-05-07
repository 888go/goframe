
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
// Package gjson provides convenient API for JSON/XML/INI/YAML/TOML data handling.
<原文结束>

# <翻译开始>
// Package gjson 提供了方便的 API 用于处理 JSON/XML/INI/YAML/TOML 数据。
# <翻译结束>


<原文开始>
// Separator char for hierarchical data access.
<原文结束>

# <翻译开始>
// 分隔符字符，用于分层数据访问。
# <翻译结束>


<原文开始>
// Json is the customized JSON struct.
<原文结束>

# <翻译开始>
// Json 是自定义的 JSON 结构体。
# <翻译结束>


<原文开始>
// Violence Check(false in default), which is used to access data when the hierarchical data key contains separator char.
<原文结束>

# <翻译开始>
// 暴力检查（默认为false），用于在分层数据键包含分隔符字符时访问数据。
# <翻译结束>


<原文开始>
// Options for Json object creating/loading.
<原文结束>

# <翻译开始>
// Json对象创建/加载的选项。
# <翻译结束>


<原文开始>
// Mark this object is for in concurrent-safe usage. This is especially for Json object creating.
<原文结束>

# <翻译开始>
// 标记此对象适用于并发安全的使用场景。这尤其针对Json对象的创建。
# <翻译结束>


<原文开始>
// Type specifies the data content type, eg: json, xml, yaml, toml, ini.
<原文结束>

# <翻译开始>
// Type 指定数据内容类型，例如：json、xml、yaml、toml、ini。
# <翻译结束>


<原文开始>
// StrNumber causes the Decoder to unmarshal a number into an interface{} as a string instead of as a float64.
<原文结束>

# <翻译开始>
// StrNumber 使 Decoder 在将数字反序列化到 interface{} 时，将其作为字符串处理而非 float64。
# <翻译结束>


<原文开始>
// iInterfaces is used for type assert api for Interfaces().
<原文结束>

# <翻译开始>
// iInterfaces 用于对 Interfaces() 方法进行类型断言。
# <翻译结束>


<原文开始>
// iMapStrAny is the interface support for converting struct parameter to map.
<原文结束>

# <翻译开始>
// iMapStrAny 是支持将结构体参数转换为映射的接口。
# <翻译结束>


<原文开始>
// iVal is the interface for underlying interface{} retrieving.
<原文结束>

# <翻译开始>
// iVal 是用于获取底层 interface{} 的接口。
# <翻译结束>


<原文开始>
// setValue sets `value` to `j` by `pattern`.
// Note:
// 1. If value is nil and removed is true, means deleting this value;
// 2. It's quite complicated in hierarchical data search, node creating and data assignment;
<原文结束>

# <翻译开始>
// setValue 通过 `pattern` 将 `value` 设置为 `j`。
// 注意：
// 1. 如果 value 为 nil 且 removed 为 true，表示删除这个值；
// 2. 在层次数据搜索、节点创建和数据赋值过程中较为复杂。
# <翻译结束>












<原文开始>
// If the key does not exit in the map.
<原文结束>

# <翻译开始>
// 如果键在映射中不存在。
# <翻译结束>
































<原文开始>
		// If the variable pointed to by the `pointer` is not of a reference type,
		// then it modifies the variable via its the parent, ie: pparent.
<原文结束>

# <翻译开始>
// 如果通过`pointer`指向的变量不是引用类型，
// 那么它将通过其父级（即：pparent）来修改该变量。
// 在Go语言中，这段注释描述了如果给定的指针`pointer`不指向一个引用类型，那么对变量的修改会通过其上级父级指针`pparent`间接完成。
# <翻译结束>


<原文开始>
// convertValue converts `value` to map[string]interface{} or []interface{},
// which can be supported for hierarchical data access.
<原文结束>

# <翻译开始>
// convertValue 将 `value` 转换为 map[string]interface{} 或 []interface{}，
// 这样就可以支持对层级数据的访问。
# <翻译结束>


<原文开始>
// setPointerWithValue sets `key`:`value` to `pointer`, the `key` may be a map key or slice index.
// It returns the pointer to the new value set.
<原文结束>

# <翻译开始>
// setPointerWithValue 将 `key`:`value` 设置到 `pointer` 中，其中 `key` 可能是 map 的键或 slice 的索引。
// 它返回指向新设置值的指针。
# <翻译结束>


<原文开始>
// getPointerByPattern returns a pointer to the value by specified `pattern`.
<原文结束>

# <翻译开始>
// getPointerByPattern 根据指定的`pattern`返回一个指向该值的指针。
# <翻译结束>


<原文开始>
// getPointerByPatternWithViolenceCheck returns a pointer to the value of specified `pattern` with violence check.
<原文结束>

# <翻译开始>
// getPointerByPatternWithViolenceCheck 函数通过暴力检查的方式，返回指定 `pattern` 的值的指针。
# <翻译结束>


<原文开始>
// It returns nil if pattern is empty.
<原文结束>

# <翻译开始>
// 如果模式为空，则返回nil。
# <翻译结束>


<原文开始>
// It returns all if pattern is ".".
<原文结束>

# <翻译开始>
// 如果模式为"."，则返回所有内容。
# <翻译结束>


<原文开始>
// Get the position for next separator char.
<原文结束>

# <翻译开始>
// 获取下一个分隔符字符的位置。
# <翻译结束>


<原文开始>
// getPointerByPatternWithoutViolenceCheck returns a pointer to the value of specified `pattern`, with no violence check.
<原文结束>

# <翻译开始>
// getPointerByPatternWithoutViolenceCheck 根据指定的`pattern`返回一个指向其值的指针，且不进行暴力检查。
# <翻译结束>


<原文开始>
// checkPatternByPointer checks whether there's value by `key` in specified `pointer`.
// It returns a pointer to the value.
<原文结束>

# <翻译开始>
// checkPatternByPointer 检查在指定的 `pointer` 中是否存在通过 `key` 访问的值。
// 它会返回该值的指针。
# <翻译结束>


<原文开始>
// Char separator('.' in default).
<原文结束>

# <翻译开始>
// 字符分隔符（默认为'.'）
# <翻译结束>


<原文开始>
// Custom priority tags for decoding, eg: "json,yaml,MyTag". This is especially for struct parsing into Json object.
<原文结束>

# <翻译开始>
// 自定义解析优先级标签，例如："json,yaml,MyTag"。这主要用于将结构体解析为Json对象时。
# <翻译结束>












<原文开始>
// Initialization checks.
<原文结束>

# <翻译开始>
// 初始化检查。
# <翻译结束>


<原文开始>
// Delete item from map.
<原文结束>

# <翻译开始>
// 从map中删除项目。
# <翻译结束>


<原文开始>
// Creating new node.
<原文结束>

# <翻译开始>
// 创建新的节点。
# <翻译结束>


<原文开始>
// Creating array node.
<原文结束>

# <翻译开始>
// 创建切片节点
# <翻译结束>


<原文开始>
// Creating map node.
<原文结束>

# <翻译开始>
// 创建映射节点。
# <翻译结束>


<原文开始>
// It is the root node.
<原文结束>

# <翻译开始>
// 这是根节点。
# <翻译结束>


<原文开始>
// It is not the root node.
<原文结束>

# <翻译开始>
// 它不是根节点。
# <翻译结束>

