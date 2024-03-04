
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
// Structs converts any slice to given struct slice.
// Also see Scan, Struct.
<原文结束>

# <翻译开始>
// Structs 将任何切片转换为给定的结构体切片。
// 也可以参考 Scan, Struct。
# <翻译结束>


<原文开始>
// SliceStruct is alias of Structs.
<原文结束>

# <翻译开始>
// SliceStruct 是 Structs 的别名。
# <翻译结束>


<原文开始>
// StructsTag acts as Structs but also with support for priority tag feature, which retrieves the
// specified tags for `params` key-value items to struct attribute names mapping.
// The parameter `priorityTag` supports multiple tags that can be joined with char ','.
<原文结束>

# <翻译开始>
// StructsTag的行为类似于Structs，但增加了对优先级标签功能的支持，该功能用于获取指定的标签，以便将`params`键值对映射到结构体属性名称。
// 参数`priorityTag`支持多个标签，这些标签可以使用字符','连接。
# <翻译结束>


<原文开始>
// doStructs converts any slice to given struct slice.
//
// It automatically checks and converts json string to []map if `params` is string/[]byte.
//
// The parameter `pointer` should be type of pointer to slice of struct.
// Note that if `pointer` is a pointer to another pointer of type of slice of struct,
// it will create the struct/pointer internally.
<原文结束>

# <翻译开始>
// doStructs 将任何切片转换为给定的结构体切片。
//
// 如果`params`是字符串或[]byte，它会自动检查并转换为json字符串到[]map。
//
// 参数`pointer`应为指向结构体切片的指针类型。
// 注意，如果`pointer`是指向另一个结构体切片类型的指针的指针，
// 它将内部创建结构体/指针。
// 以下是更详细的中文注释：
// ```go
// doStructs 函数用于将任意类型的切片转换为目标结构体切片。
//
// 当传入参数 `params` 为字符串或字节切片（[]byte）时，函数会自动检测并将其转换成 JSON 字符串，进一步解析为 []map 类型的数据。
//
// 参数 `pointer` 需要是指向结构体切片的指针类型。特别地，如果 `pointer` 是一个指向结构体切片指针的指针，该函数将在内部自行创建所需结构体及相应的指针。
# <翻译结束>


<原文开始>
// If `params` is nil, no conversion.
<原文结束>

# <翻译开始>
// 如果`params`为nil，则不进行转换。
# <翻译结束>


<原文开始>
// Catch the panic, especially the reflection operation panics.
<原文结束>

# <翻译开始>
// 捕获 panic，特别是反射操作引发的 panic。
# <翻译结束>


<原文开始>
// If given `params` is JSON, it then uses json.Unmarshal doing the converting.
<原文结束>

# <翻译开始>
// 如果给定的`params`是JSON格式，那么它将使用json.Unmarshal进行转换。
# <翻译结束>







<原文开始>
// Converting `params` to map slice.
<原文结束>

# <翻译开始>
// 将`params`转换为映射切片。
# <翻译结束>


<原文开始>
// If `params` is an empty slice, no conversion.
<原文结束>

# <翻译开始>
// 如果`params`是一个空切片，则不进行转换。
# <翻译结束>







<原文开始>
// doStructsByDirectReflectSet do the converting directly using reflect Set.
// It returns true if success, or else false.
<原文结束>

# <翻译开始>
// doStructsByDirectReflectSet 直接使用 reflect.Set 进行转换操作。
// 如果转换成功，返回 true，否则返回 false。
# <翻译结束>







<原文开始>
// Pointer type check.
<原文结束>

# <翻译开始>
// 指针类型检查。
# <翻译结束>

