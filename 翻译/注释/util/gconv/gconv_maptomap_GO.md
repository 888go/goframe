
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
// MapToMap converts any map type variable `params` to another map type variable `pointer`
// using reflect.
// See doMapToMap.
<原文结束>

# <翻译开始>
// MapToMap 使用 reflect 将任意类型的 map 变量 `params` 转换为另一种 map 类型变量 `pointer`
// 详细实现请参考 doMapToMap 函数。
# <翻译结束>


<原文开始>
// doMapToMap converts any map type variable `params` to another map type variable `pointer`.
//
// The parameter `params` can be any type of map, like:
// map[string]string, map[string]struct, map[string]*struct, reflect.Value, etc.
//
// The parameter `pointer` should be type of *map, like:
// map[int]string, map[string]struct, map[string]*struct, reflect.Value, etc.
//
// The optional parameter `mapping` is used for struct attribute to map key mapping, which makes
// sense only if the items of original map `params` is type struct.
<原文结束>

# <翻译开始>
// doMapToMap 将任意类型的 map 变量 `params` 转换为另一种 map 类型变量 `pointer`。
//
// 参数 `params` 可以是任意类型的 map，例如：
// map[string]string, map[string]struct, map[string]*struct, reflect.Value 等等。
//
// 参数 `pointer` 应该是指向 map 的类型，例如：
// *map[int]string, *map[string]struct, *map[string]*struct, *reflect.Value 等等。
//
// 可选参数 `mapping` 用于 struct 属性到 map 键的映射，只有当原 map `params` 中的项为 struct 类型时才有意义。
# <翻译结束>


<原文开始>
// If given `params` is JSON, it then uses json.Unmarshal doing the converting.
<原文结束>

# <翻译开始>
// 如果给定的`params`是JSON格式，那么它将使用json.Unmarshal进行转换。
# <翻译结束>


<原文开始>
// Empty params map, no need continue.
<原文结束>

# <翻译开始>
// 空参数映射，无需继续。
# <翻译结束>


<原文开始>
// Catch the panic, especially the reflection operation panics.
<原文结束>

# <翻译开始>
// 捕获 panic，特别是反射操作引发的 panic。
# <翻译结束>


<原文开始>
// Retrieve the true element type of target map.
<原文结束>

# <翻译开始>
// 获取目标映射的真正元素类型。
# <翻译结束>

