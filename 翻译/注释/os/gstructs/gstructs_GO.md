
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
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Package gstructs provides functions for struct information retrieving.
<原文结束>

# <翻译开始>
// 包gstructs提供了用于获取结构体信息的函数。
// 这个不翻译了, 这是一个偏底层组件，一般业务上很少会用到，在框架、基础库、中间件编写中用到。
# <翻译结束>


<原文开始>
// Type wraps reflect.Type for additional features.
<原文结束>

# <翻译开始>
// Type 类型封装了 reflect.Type，提供了额外的功能。
# <翻译结束>


<原文开始>
// Field contains information of a struct field .
<原文结束>

# <翻译开始>
// Field 包含结构体字段的信息。
# <翻译结束>


<原文开始>
// The underlying value of the field.
<原文结束>

# <翻译开始>
// 字段的底层值
# <翻译结束>


<原文开始>
// The underlying field of the field.
<原文结束>

# <翻译开始>
// 字段的底层字段。
# <翻译结束>


<原文开始>
// Retrieved tag name. It depends TagValue.
<原文结束>

# <翻译开始>
// 获取到的标签名称，其依赖于TagValue。
# <翻译结束>


<原文开始>
	// Retrieved tag value.
	// There might be more than one tags in the field,
	// but only one can be retrieved according to calling function rules.
<原文结束>

# <翻译开始>
// 获取到的标签值。
// 在该字段中可能包含多个标签，
// 但根据调用函数规则，只能获取其中一个。
# <翻译结束>


<原文开始>
// FieldsInput is the input parameter struct type for function Fields.
<原文结束>

# <翻译开始>
// FieldsInput 是函数 Fields 的输入参数结构体类型。
# <翻译结束>


<原文开始>
	// Pointer should be type of struct/*struct.
	// TODO this attribute name is not suitable, which would make confuse.
<原文结束>

# <翻译开始>
// Pointer 应为结构体类型。/*struct. */
// TODO 这个属性名称不合适，可能会引起混淆。
# <翻译结束>


<原文开始>
	// RecursiveOption specifies the way retrieving the fields recursively if the attribute
	// is an embedded struct. It is RecursiveOptionNone in default.
<原文结束>

# <翻译开始>
// RecursiveOption 指定当属性是一个嵌入式结构体时，以何种方式递归获取其字段。默认情况下，其值为 RecursiveOptionNone。
// ```go
// RecursiveOption 指定了在遇到嵌入式结构体属性时，如何进行递归地获取其字段的选项。默认设置为 RecursiveOptionNone。
# <翻译结束>


<原文开始>
// FieldMapInput is the input parameter struct type for function FieldMap.
<原文结束>

# <翻译开始>
// FieldMapInput 是函数 FieldMap 的输入参数结构体类型。
# <翻译结束>


<原文开始>
	// PriorityTagArray specifies the priority tag array for retrieving from high to low.
	// If it's given `nil`, it returns map[name]Field, of which the `name` is attribute name.
<原文结束>

# <翻译开始>
// PriorityTagArray 指定从高到低检索的优先级标签数组。
// 如果传入 `nil`，则返回 map[name]Field，其中 `name` 是属性名。
# <翻译结束>


<原文开始>
// No recursively retrieving fields as map if the field is an embedded struct.
<原文结束>

# <翻译开始>
// 如果字段是嵌入式结构体，则不递归地以 map 形式获取其字段。
# <翻译结束>


<原文开始>
// Recursively retrieving fields as map if the field is an embedded struct.
<原文结束>

# <翻译开始>
// 如果字段是一个嵌入式结构体，则递归地将其字段作为映射获取。
# <翻译结束>


<原文开始>
// Recursively retrieving fields as map if the field is an embedded struct and the field has no tag.
<原文结束>

# <翻译开始>
// 如果字段是一个嵌入式结构体且该字段没有标签，则递归获取其字段并以映射形式返回。
# <翻译结束>


<原文开始>
// Fields retrieves and returns the fields of `pointer` as slice.
<原文结束>

# <翻译开始>
// Fields 函数检索并返回 `pointer` 的字段作为一个切片。
# <翻译结束>


<原文开始>
// The current level fields can overwrite the sub-struct fields with the same name.
<原文结束>

# <翻译开始>
// 当前层级字段可以覆盖具有相同名称的子结构体字段。
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
// FieldMap 从`pointer`获取并返回结构体字段为map[name/tag]Field。
//
// 参数`pointer`应为struct/*struct类型。
//
// 参数`priority`指定了按从高到低优先级获取的标签数组。如果给出`nil`，则返回map[name]Field，其中`name`是属性名称。
//
// 参数`recursive`指定了当属性是一个嵌入式结构体时，是否递归地获取其字段。
//
// 注意，它仅从结构体中获取首字母大写的导出属性（即公开属性）。
# <翻译结束>


<原文开始>
// Only retrieve exported attributes.
<原文结束>

# <翻译开始>
// 仅获取导出的属性。
# <翻译结束>


<原文开始>
// StructType retrieves and returns the struct Type of specified struct/*struct.
// The parameter `object` should be either type of struct/*struct/[]struct/[]*struct.
<原文结束>

# <翻译开始>
// StructType 函数检索并返回指定结构体的结构体类型。
// 参数 `object` 应为 struct 类型、指针到 struct 类型、struct 切片或指针到 struct 切片类型。
# <翻译结束>


<原文开始>
// If pointer is type of *struct and nil, then automatically create a temporary struct.
<原文结束>

# <翻译开始>
// 如果指针是结构体类型且为nil，则自动创建一个临时结构体。
# <翻译结束>

