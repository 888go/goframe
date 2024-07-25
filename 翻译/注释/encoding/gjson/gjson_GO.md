
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
// Package gjson provides convenient API for JSON/XML/INI/YAML/TOML data handling.
<原文结束>

# <翻译开始>
// 包gjson提供了处理JSON/XML/INI/YAML/TOML数据的便捷API。 md5:ddbf6ad5d309a49c
# <翻译结束>


<原文开始>
// Separator char for hierarchical data access.
<原文结束>

# <翻译开始>
// 用于层次数据访问的分隔符字符。 md5:3020966f087a4732
# <翻译结束>


<原文开始>
// Json is the customized JSON struct.
<原文结束>

# <翻译开始>
// Json 是自定义的JSON结构体。 md5:764b883e7cf79da7
# <翻译结束>


<原文开始>
// Char separator('.' in default).
<原文结束>

# <翻译开始>
// 字符分隔符（默认为'.'）。 md5:15307025b7ed9ae7
# <翻译结束>


<原文开始>
// Violence Check(false in default), which is used to access data when the hierarchical data key contains separator char.
<原文结束>

# <翻译开始>
// 暴力检查（默认为false），用于在层次数据键包含分隔符字符时访问数据。 md5:465e099ccbdc4ca3
# <翻译结束>


<原文开始>
// Options for Json object creating/loading.
<原文结束>

# <翻译开始>
// 创建/加载Json对象的选项。 md5:d8614ea5dc358e89
# <翻译结束>


<原文开始>
// Mark this object is for in concurrent-safe usage. This is especially for Json object creating.
<原文结束>

# <翻译开始>
// 标记此对象适用于并发安全使用。这尤其适用于 Json 对象的创建。 md5:59d439559fecdc34
# <翻译结束>


<原文开始>
// Custom priority tags for decoding, eg: "json,yaml,MyTag". This is especially for struct parsing into Json object.
<原文结束>

# <翻译开始>
// 自定义解码优先级标签，例如："json,yaml,MyTag"。这主要用于将结构体解析为Json对象。 md5:486cf257ddd06463
# <翻译结束>


<原文开始>
// Type specifies the data content type, eg: json, xml, yaml, toml, ini.
<原文结束>

# <翻译开始>
// 类型指定了数据内容类型，例如：json、xml、yaml、toml、ini等。 md5:afbae78560edde30
# <翻译结束>


<原文开始>
// StrNumber causes the Decoder to unmarshal a number into an interface{} as a string instead of as a float64.
<原文结束>

# <翻译开始>
// StrNumber 使得 Decoder 将数字解码为 interface{}` 作为字符串，而不是 float64。 md5:32e44e32c3cc37cc
# <翻译结束>


<原文开始>
// iInterfaces is used for type assert api for Interfaces().
<原文结束>

# <翻译开始>
// iInterfaces 用于类型断言接口，用于 Interfaces() 方法。 md5:711dc755f9cd4979
# <翻译结束>


<原文开始>
// iMapStrAny is the interface support for converting struct parameter to map.
<原文结束>

# <翻译开始>
// iMapStrAny 是一个接口，支持将结构体参数转换为映射。 md5:cfd4642c77fca6ec
# <翻译结束>


<原文开始>
// iVal is the interface for underlying interface{} retrieving.
<原文结束>

# <翻译开始>
// iVal是用于获取底层interface{}的接口。 md5:2915e3bd3d7e4f43
# <翻译结束>


<原文开始>
// setValue sets `value` to `j` by `pattern`.
// Note:
// 1. If value is nil and removed is true, means deleting this value;
// 2. It's quite complicated in hierarchical data search, node creating and data assignment;
<原文结束>

# <翻译开始>
// setValue 将`value`设置为`j`，按照`pattern`。
// 注意：
// 1. 如果`value`为nil且`removed`为true，表示删除这个值；
// 2. 在层次数据搜索、节点创建和数据赋值方面相当复杂。
// md5:6aca091405b9da40
# <翻译结束>


<原文开始>
// If the key does not exit in the map.
<原文结束>

# <翻译开始>
				// 如果键在映射中不存在。 md5:ba2af475e1347525
# <翻译结束>


<原文开始>
// It is not the root node.
<原文结束>

# <翻译开始>
						// 它不是根节点。 md5:b90762478c5a92c6
# <翻译结束>


<原文开始>
		// If the variable pointed to by the `pointer` is not of a reference type,
		// then it modifies the variable via its the parent, ie: pparent.
<原文结束>

# <翻译开始>
		// 如果`pointer`指向的变量不是引用类型，那么它会通过其父对象（pparent）来修改该变量。
		// md5:aa59525c846686ce
# <翻译结束>


<原文开始>
// convertValue converts `value` to map[string]interface{} or []interface{},
// which can be supported for hierarchical data access.
<原文结束>

# <翻译开始>
// convertValue将"value"转换为map[string]interface{}或[]interface{}，这样可以支持层级数据访问。
// md5:089e6e9291ed7aab
# <翻译结束>


<原文开始>
// setPointerWithValue sets `key`:`value` to `pointer`, the `key` may be a map key or slice index.
// It returns the pointer to the new value set.
<原文结束>

# <翻译开始>
// setPointerWithValue 将 `key`:`value` 设置到 `pointer` 中，其中 `key` 可能是映射的键或切片的索引。
// 它返回新设置值的指针。
// md5:2642aca0fd23f46c
# <翻译结束>


<原文开始>
// getPointerByPattern returns a pointer to the value by specified `pattern`.
<原文结束>

# <翻译开始>
// getPointerByPattern 根据指定的 `pattern` 返回值的指针。 md5:e5422879dc2c9285
# <翻译结束>


<原文开始>
// getPointerByPatternWithViolenceCheck returns a pointer to the value of specified `pattern` with violence check.
<原文结束>

# <翻译开始>
// getPointerByPatternWithViolenceCheck 通过暴力检查返回指定`pattern`的值的指针。 md5:4ac204b4633753dc
# <翻译结束>


<原文开始>
// It returns nil if pattern is empty.
<原文结束>

# <翻译开始>
	// 如果pattern为空，它将返回nil。 md5:8e2a6f56affd353a
# <翻译结束>


<原文开始>
// It returns all if pattern is ".".
<原文结束>

# <翻译开始>
	// 如果pattern是"."，则返回所有。 md5:1f0d65d517f332bd
# <翻译结束>


<原文开始>
// Get the position for next separator char.
<原文结束>

# <翻译开始>
			// 获取下一个分隔符字符的位置。 md5:7268bb1b6598460b
# <翻译结束>


<原文开始>
// getPointerByPatternWithoutViolenceCheck returns a pointer to the value of specified `pattern`, with no violence check.
<原文结束>

# <翻译开始>
// getPointerByPatternWithoutViolenceCheck 返回指定`pattern`值的指针，不进行暴力检查。 md5:fd58f2cfd08f8751
# <翻译结束>


<原文开始>
// checkPatternByPointer checks whether there's value by `key` in specified `pointer`.
// It returns a pointer to the value.
<原文结束>

# <翻译开始>
// checkPatternByPointer 检查指定`pointer`中是否存在键为`key`的值。它返回该值的指针。
// md5:10f17307c0c6e052
# <翻译结束>

