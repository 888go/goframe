
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
// IsEmpty checks and returns whether `r` is empty.
<原文结束>

# <翻译开始>
// IsEmpty 检查并返回 `r` 是否为空。
# <翻译结束>


<原文开始>
// Len returns the length of result list.
<原文结束>

# <翻译开始>
// Len 返回结果列表的长度。
# <翻译结束>


<原文开始>
// Size is alias of function Len.
<原文结束>

# <翻译开始>
// Size 是函数 Len 的别名。
# <翻译结束>


<原文开始>
// Chunk splits a Result into multiple Results,
// the size of each array is determined by `size`.
// The last chunk may contain less than size elements.
<原文结束>

# <翻译开始>
// Chunk 将一个 Result 切分成多个 Result，
// 每个数组的大小由 `size` 决定。
// 最后一块切片可能包含少于 size 个元素。
# <翻译结束>


<原文开始>
// Json converts `r` to JSON format content.
<原文结束>

# <翻译开始>
// Json 将 `r` 转换为 JSON 格式的内容。
# <翻译结束>


<原文开始>
// Xml converts `r` to XML format content.
<原文结束>

# <翻译开始>
// Xml 将`r`转换为XML格式的内容。
# <翻译结束>


<原文开始>
// List converts `r` to a List.
<原文结束>

# <翻译开始>
// List 将 `r` 转换为一个 List。
# <翻译结束>


<原文开始>
// Array retrieves and returns specified column values as slice.
// The parameter `field` is optional is the column field is only one.
// The default `field` is the first field name of the first item in `Result` if parameter `field` is not given.
<原文结束>

# <翻译开始>
// Array 根据指定列字段获取并返回其值作为一个切片。
// 当列字段只有一个时，参数 `field` 可选。
// 如果未给出参数 `field`，则默认的 `field` 为 `Result` 中第一项的第一个字段名。
# <翻译结束>


<原文开始>
// MapKeyValue converts `r` to a map[string]Value of which key is specified by `key`.
// Note that the item value may be type of slice.
<原文结束>

# <翻译开始>
// MapKeyValue 将 `r` 转换为一个 map[string]Value，其中键由 `key` 指定。
// 注意，项值可能为切片类型。
# <翻译结束>


<原文开始>
// MapKeyStr converts `r` to a map[string]Map of which key is specified by `key`.
<原文结束>

# <翻译开始>
// MapKeyStr 将 `r` 转换为一个 map[string]Map 类型的映射，其中的键由 `key` 指定。
# <翻译结束>


<原文开始>
// MapKeyInt converts `r` to a map[int]Map of which key is specified by `key`.
<原文结束>

# <翻译开始>
// MapKeyInt 将 `r` 转换为一个映射 map[int]Map，其中键由 `key` 指定。
// （注：这里可能需要上下文信息，对于 `Map` 类型没有明确说明，所以翻译时假设它是一个已知的类型名。如果 `Map` 是自定义类型或有特殊含义，请替换为实际含义。）
# <翻译结束>


<原文开始>
// MapKeyUint converts `r` to a map[uint]Map of which key is specified by `key`.
<原文结束>

# <翻译开始>
// MapKeyUint 将`r`转换为一个map[uint]Map类型，其中键由`key`指定。
# <翻译结束>


<原文开始>
// RecordKeyStr converts `r` to a map[string]Record of which key is specified by `key`.
<原文结束>

# <翻译开始>
// RecordKeyStr 将 `r` 转换为一个 map[string]Record 类型的映射，其中键由 `key` 指定。
# <翻译结束>


<原文开始>
// RecordKeyInt converts `r` to a map[int]Record of which key is specified by `key`.
<原文结束>

# <翻译开始>
// RecordKeyInt 将 `r` 转换为一个 map[int]Record 类型的映射，其中键由 `key` 指定。
# <翻译结束>


<原文开始>
// RecordKeyUint converts `r` to a map[uint]Record of which key is specified by `key`.
<原文结束>

# <翻译开始>
// RecordKeyUint 将 `r` 转换为一个 map[uint]Record 类型的映射，其中键由 `key` 指定。
# <翻译结束>


<原文开始>
// Structs converts `r` to struct slice.
// Note that the parameter `pointer` should be type of *[]struct/*[]*struct.
<原文结束>

# <翻译开始>
// Structs 将 `r` 转换为结构体切片。
// 注意参数 `pointer` 应该是指向结构体切片的指针类型，即 *[]struct 或 *[]*struct。
# <翻译结束>


<原文开始>
// If the result is empty and the target pointer is not empty, it returns error.
<原文结束>

# <翻译开始>
// 如果结果为空且目标指针不为空，则返回错误。
# <翻译结束>

