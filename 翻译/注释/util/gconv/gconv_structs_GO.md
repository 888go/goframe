
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
// Structs converts any slice to given struct slice.
// Also see Scan, Struct.
<原文结束>

# <翻译开始>
// Structs 将任何切片转换为给定结构体类型的切片。
// 另请参见 Scan, Struct。
// md5:9f4251433eeb586c
# <翻译结束>


<原文开始>
// SliceStruct is alias of Structs.
<原文结束>

# <翻译开始>
// SliceStruct 是 Structs 的别名。. md5:844cd0606fb4edf0
# <翻译结束>


<原文开始>
// StructsTag acts as Structs but also with support for priority tag feature, which retrieves the
// specified tags for `params` key-value items to struct attribute names mapping.
// The parameter `priorityTag` supports multiple tags that can be joined with char ','.
<原文结束>

# <翻译开始>
// StructsTag 作为 Structs 的功能增强版本，还支持优先级标签特性。它根据 `params` 键值对获取指定的标签，并将其映射到结构体属性名称上。
// 参数 `priorityTag` 支持多个标签，这些标签可以使用逗号分隔。
// md5:ddc344beca5956a8
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
// 如果 `params` 是字符串或[]byte，它会自动检查并将其转换为[]map。
//
// 参数 `pointer` 应该是指向结构体切片的指针。注意，如果 `pointer` 是指向结构体切片的另一个指针的指针，
// 它会在内部创建结构体/指针。
// md5:0bddadd6970c6b0b
# <翻译结束>


<原文开始>
// Catch the panic, especially the reflection operation panics.
<原文结束>

# <翻译开始>
// 捕获panic，尤其是反射操作引发的panic。. md5:dd183bf8028f513a
# <翻译结束>


<原文开始>
// Converting `params` to map slice.
<原文结束>

# <翻译开始>
// 将`params`转换为映射切片。. md5:d0685c4290b475fe
# <翻译结束>


<原文开始>
// If `params` is an empty slice, no conversion.
<原文结束>

# <翻译开始>
// 如果`params`是一个空切片，则不进行转换。. md5:c2aa546ea7052f08
# <翻译结束>

