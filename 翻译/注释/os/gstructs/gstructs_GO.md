
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
// Package gstructs provides functions for struct information retrieving.
<原文结束>

# <翻译开始>
// 包gstructs提供了用于获取结构体信息的函数。 md5:ff6813ae5e3453ba
# <翻译结束>


<原文开始>
// Type wraps reflect.Type for additional features.
<原文结束>

# <翻译开始>
// Type 是 reflect.Type 的增强版本，提供了更多功能。 md5:8ebe2d126efacb49
# <翻译结束>


<原文开始>
// Field contains information of a struct field .
<原文结束>

# <翻译开始>
// Field包含一个结构字段的信息。 md5:937dc69e9da6620a
# <翻译结束>


<原文开始>
// The underlying value of the field.
<原文结束>

# <翻译开始>
// 字段的基础值。 md5:30c17ef0d945aeca
# <翻译结束>


<原文开始>
// The underlying field of the field.
<原文结束>

# <翻译开始>
// 字段的底层字段。 md5:5d18dc4427e59bd7
# <翻译结束>


<原文开始>
// Retrieved tag name. It depends TagValue.
<原文结束>

# <翻译开始>
// 从TagValue中获取标签名。这取决于TagValue。 md5:2abab33cf7d9aa7a
# <翻译结束>


<原文开始>
	// Retrieved tag value.
	// There might be more than one tags in the field,
	// but only one can be retrieved according to calling function rules.
<原文结束>

# <翻译开始>
// 获取标签值。
// 字段中可能有多个标签，但根据调用函数的规则，只能获取一个。
// md5:45a4365044272532
# <翻译结束>


<原文开始>
// FieldsInput is the input parameter struct type for function Fields.
<原文结束>

# <翻译开始>
// FieldsInput 是函数 Fields 的输入参数结构体类型。 md5:dea3d8d32792f883
# <翻译结束>


<原文开始>
	// Pointer should be type of struct/*struct.
	// TODO this attribute name is not suitable, which would make confuse.
<原文结束>

# <翻译开始>
// Pointer 应该是 struct 类型的指针/*指向结构体的指针类型*/。
// TODO 这个属性名称不合适，可能会引起混淆。
// md5:7115141b84d46b4a
# <翻译结束>


<原文开始>
	// RecursiveOption specifies the way retrieving the fields recursively if the attribute
	// is an embedded struct. It is RecursiveOptionNone in default.
<原文结束>

# <翻译开始>
// RecursiveOption 定义了如果属性是一个嵌入的结构体，如何递归地检索字段。默认情况下为 RecursiveOptionNone。
// md5:ad0b9ef1d1f1f1e5
# <翻译结束>


<原文开始>
// FieldMapInput is the input parameter struct type for function FieldMap.
<原文结束>

# <翻译开始>
// FieldMapInput是FieldMap函数的输入参数结构体类型。 md5:6265e8efb4329ab9
# <翻译结束>


<原文开始>
	// PriorityTagArray specifies the priority tag array for retrieving from high to low.
	// If it's given `nil`, it returns map[name]Field, of which the `name` is attribute name.
<原文结束>

# <翻译开始>
// PriorityTagArray 用于指定优先级标签数组，按从高到低的顺序检索。
// 如果传入 `nil`，则返回 map[name]Field，其中 `name` 是属性名称。
// md5:454af14097a1e0a3
# <翻译结束>


<原文开始>
// No recursively retrieving fields as map if the field is an embedded struct.
<原文结束>

# <翻译开始>
// 如果字段是一个嵌入的结构体，不递归地将其字段作为映射获取。 md5:7e5b33b2b6192298
# <翻译结束>


<原文开始>
// Recursively retrieving fields as map if the field is an embedded struct.
<原文结束>

# <翻译开始>
// 如果字段是一个嵌入的结构体，递归地获取其字段作为map。 md5:5c2576800c0efe83
# <翻译结束>


<原文开始>
// Recursively retrieving fields as map if the field is an embedded struct and the field has no tag.
<原文结束>

# <翻译开始>
// 如果字段是嵌入的结构并且字段没有标签，递归地获取字段作为映射。 md5:24a441218c457b09
# <翻译结束>


<原文开始>
// Fields retrieves and returns the fields of `pointer` as slice.
<原文结束>

# <翻译开始>
// Fields 获取并以切片形式返回 `pointer` 的字段。 md5:7856c4ee9e72f56f
# <翻译结束>


<原文开始>
// The current level fields can overwrite the sub-struct fields with the same name.
<原文结束>

# <翻译开始>
// 当前级别字段可以覆盖具有相同名称的子结构体字段。 md5:e9bd19d3abe6f7e5
# <翻译结束>


<原文开始>
// FieldMap retrieves and returns struct field as map[name/tag]Field from `pointer`.
//
// The parameter `pointer` should be type of struct/*struct.
//
// The parameter `priority` specifies the priority tag array for retrieving from high to low.
// If it's given `nil`, it returns map[name]Field, of which the `name` is attribute name.
//
// The parameter `recursive` specifies whether retrieving the fields recursively if the attribute
// is an embedded struct.
//
// Note that it only retrieves the exported attributes with first letter upper-case from struct.
<原文结束>

# <翻译开始>
// FieldMap 从 `pointer` 获取并返回结构体字段，作为 map[name/tag]Field。
//
// 参数 `pointer` 应为 struct 或 *struct 类型。
//
// 参数 `priority` 指定了用于检索的优先级标签数组，从高到低。如果为空（`nil`），则返回 map[name]Field，其中 `name` 为属性名称。
//
// 参数 `recursive` 指定是否递归检索嵌入的结构体中的字段。
//
// 注意，它仅从结构体中检索首字母大写的导出属性。
// md5:deef4c5e31602259
# <翻译结束>


<原文开始>
// Only retrieve exported attributes.
<原文结束>

# <翻译开始>
// 只检索导出的属性。 md5:d8185f07060feffb
# <翻译结束>


<原文开始>
// StructType retrieves and returns the struct Type of specified struct/*struct.
// The parameter `object` should be either type of struct/*struct/[]struct/[]*struct.
<原文结束>

# <翻译开始>
// StructType 获取并返回指定结构体的类型。
// 参数 `object` 应为结构体类型、指向结构体的指针类型、结构体切片类型或指向结构体的切片类型。
// md5:023b27218d435b61
# <翻译结束>


<原文开始>
// If pointer is type of *struct and nil, then automatically create a temporary struct.
<原文结束>

# <翻译开始>
// 如果指针是*struct类型且为nil，那么会自动创建一个临时的struct。 md5:23b5ebc131739e7d
# <翻译结束>

