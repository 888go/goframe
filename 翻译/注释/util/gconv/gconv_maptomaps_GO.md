
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
// MapToMaps converts any slice type variable `params` to another map slice type variable `pointer`.
// See doMapToMaps.
<原文结束>

# <翻译开始>
// MapToMaps 将任何切片类型变量 `params` 转换为另一个映射切片类型变量 `pointer`。
// 参见 doMapToMaps。
// md5:70b6d6cf0e63da31
# <翻译结束>


<原文开始>
// doMapToMaps converts any map type variable `params` to another map slice variable `pointer`.
//
// The parameter `params` can be type of []map, []*map, []struct, []*struct.
//
// The parameter `pointer` should be type of []map, []*map.
//
// The optional parameter `mapping` is used for struct attribute to map key mapping, which makes
// sense only if the item of `params` is type struct.
<原文结束>

# <翻译开始>
// doMapToMaps 将任何类型的映射变量`params`转换为另一个映射切片变量`pointer`。
//
// 参数`params`可以是[]map, []*map, []struct, []*struct类型。
//
// 参数`pointer`应该是[]map, []*map类型。
//
// 可选参数`mapping`用于结构体属性到映射键的映射，只有当`params`的元素类型为struct时才有意义。
// md5:e5da204851e0f1b9
# <翻译结束>


<原文开始>
// Params and its element type check.
<原文结束>

# <翻译开始>
	// 检查参数及其元素类型。 md5:9678a18f11496e59
# <翻译结束>


<原文开始>
// Empty slice, no need continue.
<原文结束>

# <翻译开始>
	// 空切片，无需继续。 md5:3e185b94ae24e0b3
# <翻译结束>


<原文开始>
// Pointer and its element type check.
<原文结束>

# <翻译开始>
	// 指针及其元素类型的检查。 md5:b460debe108087f5
# <翻译结束>


<原文开始>
// Catch the panic, especially the reflection operation panics.
<原文结束>

# <翻译开始>
		// 捕获panic，尤其是反射操作引发的panic。 md5:dd183bf8028f513a
# <翻译结束>

