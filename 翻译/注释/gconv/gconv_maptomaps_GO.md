
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
// MapToMaps converts any slice type variable `params` to another map slice type variable `pointer`.
// See doMapToMaps.
<原文结束>

# <翻译开始>
// MapToMaps 将任意切片类型变量 `params` 转换为另一种映射切片类型变量 `pointer`。
// 请参考 doMapToMaps 函数。
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
// doMapToMaps 将任意类型的 map 变量 `params` 转换为另一个 map 切片变量 `pointer`。
//
// 参数 `params` 可以为 []map、[]*map、[]struct 或 []*struct 类型。
//
// 参数 `pointer` 应为 []map 或 []*map 类型。
//
// 可选参数 `mapping` 用于 struct 属性到 map 键的映射，仅当 `params` 的元素类型为 struct 时才有意义。
// 这段代码注释翻译成中文后如下：
// ```go
// doMapToMaps 函数将任何类型的 map 变量 `params` 转换为另一种 map 切片变量 `pointer`。
//
// 参数 `params` 支持以下类型：[]map、[]*map、[]struct、([]*struct。
//
// 参数 `pointer` 必须是 []map 或 []*map 类型。
//
// 可选参数 `mapping` 用于进行 struct 属性到 map 键的映射，这个参数只有在 `params` 中的项为 struct 类型时才起作用。
# <翻译结束>


<原文开始>
// If given `params` is JSON, it then uses json.Unmarshal doing the converting.
<原文结束>

# <翻译开始>
// 如果给定的`params`是JSON格式，那么它将使用json.Unmarshal进行转换。
# <翻译结束>


<原文开始>
// Params and its element type check.
<原文结束>

# <翻译开始>
// 参数及其元素类型检查
# <翻译结束>


<原文开始>
// Empty slice, no need continue.
<原文结束>

# <翻译开始>
// 空切片，无需继续。
# <翻译结束>


<原文开始>
// Pointer and its element type check.
<原文结束>

# <翻译开始>
// 指针及其元素类型检查。
# <翻译结束>


<原文开始>
// Catch the panic, especially the reflection operation panics.
<原文结束>

# <翻译开始>
// 捕获 panic，特别是反射操作引发的 panic。
# <翻译结束>

