
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
// MapToMap converts any map type variable `params` to another map type variable `pointer`
// using reflect.
// See doMapToMap.
<原文结束>

# <翻译开始>
// MapToMap 通过反射将任何map类型变量`params`转换为另一个map类型变量`pointer`。
// 参考 doMapToMap。
// md5:8fbdb048d4cad524
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
// doMapToMap 将任何类型的映射变量 `params` 转换为另一个映射类型变量 `pointer`。
//
// 参数 `params` 可以是任何类型的映射，例如：map[string]string, map[string]struct, map[string]*struct, reflect.Value 等。
//
// 参数 `pointer` 应该是 *map 类型，例如：map[int]string, map[string]struct, map[string]*struct, reflect.Value 等。
//
// 可选参数 `mapping` 用于结构体属性到映射键的映射，只有当原始映射 `params` 的项是结构体类型时，这个参数才有意义。
// md5:08b8fa82edaf8b08
# <翻译结束>


<原文开始>
// Empty params map, no need continue.
<原文结束>

# <翻译开始>
// 空参数映射，无需继续。. md5:7734e4bea4d21319
# <翻译结束>


<原文开始>
// Catch the panic, especially the reflection operation panics.
<原文结束>

# <翻译开始>
// 捕获panic，尤其是反射操作引发的panic。. md5:dd183bf8028f513a
# <翻译结束>


<原文开始>
// Retrieve the true element type of target map.
<原文结束>

# <翻译开始>
// 获取目标映射的真正元素类型。. md5:7e93cce5ee0c27e1
# <翻译结束>

