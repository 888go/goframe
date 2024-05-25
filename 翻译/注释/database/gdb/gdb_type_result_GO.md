
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
// IsEmpty checks and returns whether `r` is empty.
<原文结束>

# <翻译开始>
// IsEmpty 检查 `r` 是否为空，然后返回结果。 md5:4ee28a47e769cceb
# <翻译结束>


<原文开始>
// Len returns the length of result list.
<原文结束>

# <翻译开始>
// Len 返回结果列表的长度。 md5:9abccfc01a850f4f
# <翻译结束>


<原文开始>
// Size is alias of function Len.
<原文结束>

# <翻译开始>
// Size 是函数 Len 的别名。 md5:4cfc93cb64eff9b5
# <翻译结束>


<原文开始>
// Chunk splits a Result into multiple Results,
// the size of each array is determined by `size`.
// The last chunk may contain less than size elements.
<原文结束>

# <翻译开始>
// Chunk 将一个 Result 分割成多个 Result，
// 每个数组的大小由 `size` 决定。
// 最后一块可能包含少于 size 个元素。
// md5:e1e9bbb7e5ba1969
# <翻译结束>


<原文开始>
// Json converts `r` to JSON format content.
<原文结束>

# <翻译开始>
// Json 将 `r` 转换为JSON格式的内容。 md5:60a0626b0a333d14
# <翻译结束>


<原文开始>
// Xml converts `r` to XML format content.
<原文结束>

# <翻译开始>
// Xml 将 `r` 转换为 XML 格式的内容。 md5:31a335fedb874d26
# <翻译结束>


<原文开始>
// List converts `r` to a List.
<原文结束>

# <翻译开始>
// List 将 `r` 转换为一个 List。 md5:ee79a42f10af264e
# <翻译结束>


<原文开始>
// Array retrieves and returns specified column values as slice.
// The parameter `field` is optional is the column field is only one.
// The default `field` is the first field name of the first item in `Result` if parameter `field` is not given.
<原文结束>

# <翻译开始>
// Array 用于获取并返回指定列的值作为切片。
// 参数 `field` 是可选的，如果列字段只有一个。如果未给定 `field` 参数，其默认值为 `Result` 中第一条项的第一个字段名。
// md5:f3e0a3bab6043d80
# <翻译结束>


<原文开始>
// MapKeyValue converts `r` to a map[string]Value of which key is specified by `key`.
// Note that the item value may be type of slice.
<原文结束>

# <翻译开始>
// MapKeyValue 将 `r` 转换为一个 map[string]Value，其中的键由 `key` 指定。
// 注意，项目值可能为切片类型。
// md5:0c805cb25cfa56ff
# <翻译结束>


<原文开始>
// MapKeyStr converts `r` to a map[string]Map of which key is specified by `key`.
<原文结束>

# <翻译开始>
// MapKeyStr 将 `r` 转换为一个键为指定字符串的 map[string]Map。 md5:43bb233c139ab262
# <翻译结束>


<原文开始>
// MapKeyInt converts `r` to a map[int]Map of which key is specified by `key`.
<原文结束>

# <翻译开始>
// MapKeyInt 将 `r` 转换为一个 map[int]Map，其中的键由 `key` 指定。 md5:2c63f4f80e8a0e1b
# <翻译结束>


<原文开始>
// MapKeyUint converts `r` to a map[uint]Map of which key is specified by `key`.
<原文结束>

# <翻译开始>
// MapKeyUint 将 `r` 转换为一个 map，其中的键是通过 `key` 指定的 uint 类型。 md5:0597073b149b7e00
# <翻译结束>


<原文开始>
// RecordKeyStr converts `r` to a map[string]Record of which key is specified by `key`.
<原文结束>

# <翻译开始>
// RecordKeyStr 将 `r` 转换为一个 map[string]Record，其中的键由 `key` 指定。 md5:6eaa1193e5507d8a
# <翻译结束>


<原文开始>
// RecordKeyInt converts `r` to a map[int]Record of which key is specified by `key`.
<原文结束>

# <翻译开始>
// RecordKeyInt 将 `r` 转换为一个映射[int]Record，其中键由 `key` 指定。 md5:0ebe0554d495cbae
# <翻译结束>


<原文开始>
// RecordKeyUint converts `r` to a map[uint]Record of which key is specified by `key`.
<原文结束>

# <翻译开始>
// RecordKeyUint 将 `r` 转换为一个以指定的 `key` 作为键的 [uint]Record 映射。 md5:26ce469215f5d9c2
# <翻译结束>


<原文开始>
// Structs converts `r` to struct slice.
// Note that the parameter `pointer` should be type of *[]struct/*[]*struct.
<原文结束>

# <翻译开始>
// Structs 将 `r` 转换为结构体切片。
// 注意参数 `pointer` 的类型应该是 `*[]struct` 或 `*[]*struct`。
// md5:fef766b4997dca03
# <翻译结束>


<原文开始>
// If the result is empty and the target pointer is not empty, it returns error.
<原文结束>

# <翻译开始>
// 如果结果为空且目标指针不为空，则返回错误。 md5:74dffcb96270ed89
# <翻译结束>

