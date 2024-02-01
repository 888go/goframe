
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
// Parse parses the string into map[string]interface{}.
//
// v1=m&v2=n           -> map[v1:m v2:n]
// v[a]=m&v[b]=n       -> map[v:map[a:m b:n]]
// v[a][a]=m&v[a][b]=n -> map[v:map[a:map[a:m b:n]]]
// v[]=m&v[]=n         -> map[v:[m n]]
// v[a][]=m&v[a][]=n   -> map[v:map[a:[m n]]]
// v[][]=m&v[][]=n     -> map[v:[map[]]] // Currently does not support nested slice.
// v=m&v[a]=n          -> error
// a .[[b=c            -> map[a___[b:c]
<原文结束>

# <翻译开始>
// Parse 将字符串解析为 map[string]interface{} 类型。
//
// v1=m&v2=n           -> 解析得到的映射：map[v1:m v2:n]
// v[a]=m&v[b]=n       -> 解析得到的映射：map[v:map[a:m b:n]]
// v[a][a]=m&v[a][b]=n -> 解析得到的映射：map[v:map[a:map[a:m b:n]]]
// v[]=m&v[]=n         -> 解析得到的映射：map[v:[m n]]
// v[a][]=m&v[a][]=n   -> 解析得到的映射：map[v:map[a:[m n]]]
// v[][]=m&v[][]=n     -> 解析得到的映射：map[v:[map[]]]  // 目前不支持嵌套切片
// v=m&v[a]=n          -> 报错
// a .[[b=c            -> 解析得到的映射：map[a___[b:c]]
// 注意，上述代码注释描述了一个将查询字符串形式的数据解析成 Go 语言中的 map 的功能。在处理嵌套结构时，它会根据键名包含的中括号 `[]` 和方括号 `[]` 来构建嵌套的 map 或 slice。不过需要注意的是，对于 "v[][]=m&v[][]=n" 这种情况，当前实现并不支持嵌套的 slice 结构。
# <翻译结束>












<原文开始>
// The end is slice. like f[], f[a][]
<原文结束>

# <翻译开始>
// 结尾是切片。例如 f[]，f[a][]
# <翻译结束>







<原文开始>
// The end is slice + map. like v[][a]
<原文结束>

# <翻译开始>
// 结尾是切片加映射，形式如 v[][a]
# <翻译结束>












<原文开始>
// split into multiple keys
<原文结束>

# <翻译开始>
// 将其拆分为多个键
# <翻译结束>


<原文开始>
// map, like v[a], v[a][b]
<原文结束>

# <翻译开始>
// map，类似于 v[a]、v[a][b] 的用法
# <翻译结束>

