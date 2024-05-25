
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b
# <翻译结束>


<原文开始>
// Package gmap provides most commonly used map container which also support concurrent-safe/unsafe switch feature.
<原文结束>

# <翻译开始>
// gmap 包提供了最常用的地图容器，同时支持并发安全/不安全切换特性。 md5:1f468a4fc387a466
# <翻译结束>


<原文开始>
// Map is alias of AnyAnyMap.
<原文结束>

# <翻译开始>
// Map 是 AnyAnyMap 的别名。 md5:5055001ecc89b987
# <翻译结束>


<原文开始>
// HashMap is alias of AnyAnyMap.
<原文结束>

# <翻译开始>
// HashMap 是 AnyAnyMap 的别名。 md5:3a5e44b2149d96c7
# <翻译结束>


<原文开始>
// New creates and returns an empty hash map.
// The parameter `safe` is used to specify whether using map in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// New 创建并返回一个空的哈希映射。
// 参数 `safe` 用于指定是否在并发安全模式下使用映射，默认为 false。
// md5:fca522578c694911
# <翻译结束>


<原文开始>
// NewFrom creates and returns a hash map from given map `data`.
// Note that, the param `data` map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewFrom 根据给定的映射 `data` 创建并返回一个哈希映射。
// 注意，参数 `data` 映射将被设置为底层数据映射（非深度复制），
// 因此，在外部修改该映射时可能会存在并发安全问题。
// 参数 `safe` 用于指定是否使用并发安全的树结构，默认为 false。
// md5:f596b726a77cdf08
# <翻译结束>


<原文开始>
// NewHashMap creates and returns an empty hash map.
// The parameter `safe` is used to specify whether using map in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewHashMap 创建并返回一个空的哈希映射。
// 参数 `safe` 用于指定是否在并发安全环境下使用映射，默认值为 false。
// md5:3d312812ffecae59
# <翻译结束>


<原文开始>
// NewHashMapFrom creates and returns a hash map from given map `data`.
// Note that, the param `data` map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewHashMapFrom 从给定的映射 `data` 创建并返回一个哈希映射。
// 注意，参数 `data` 映射将被设置为底层数据映射（不进行深拷贝），
// 在外部修改映射时可能会存在并发安全问题。
// 参数 `safe` 用于指定是否在并发安全中使用树，其默认值为 false。
// md5:0e21655091039f16
# <翻译结束>

