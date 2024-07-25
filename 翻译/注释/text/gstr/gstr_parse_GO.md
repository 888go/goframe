
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
// v1=m&v2=n           -> map[v1:m, v2:n]
// v[a]=m&v[b]=n       -> map[v:map[a:m, b:n]]
// v[a][a]=m&v[a][b]=n -> map[v:map[a:map[a:m, b:n]]]
// v[]=m&v[]=n         -> map[v:[m, n]]
// v[a][]=m&v[a][]=n   -> map[v:map[a:[m, n]]]
// v[][]=m&v[][]=n     -> map[v:[map[]]] // 当前不支持嵌套切片。
// v=m&v[a]=n          -> 错误
// a .[[b=c            -> 无法解析，缺少有效的键值对格式。
// md5:28f985708060eab0
# <翻译结束>


<原文开始>
// split into multiple keys
<原文结束>

# <翻译开始>
		// 分割成多个键. md5:3bdb5e68a953321c
# <翻译结束>


<原文开始>
// The end is slice. like f[], f[a][]
<原文结束>

# <翻译开始>
	// "end" 是一个切片，类似于 f[] 或者 f[a][]. md5:41e332252e9d2da1
# <翻译结束>


<原文开始>
// The end is slice + map. like v[][a]
<原文结束>

# <翻译开始>
	// 结束是切片和映射。就像 v[][][a]. md5:79444379ed8ddfc4
# <翻译结束>


<原文开始>
// map, like v[a], v[a][b]
<原文结束>

# <翻译开始>
	// 类似于 v[a]，v[a][b] 的映射. md5:e8aa555b3543c9ea
# <翻译结束>

