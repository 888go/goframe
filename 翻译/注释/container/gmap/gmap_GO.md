
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。
# <翻译结束>


<原文开始>
// Package gmap provides most commonly used map container which also support concurrent-safe/unsafe switch feature.
<原文结束>

# <翻译开始>
// 包gmap提供了最常用的映射容器，同时支持并发安全/非安全切换功能。
# <翻译结束>







<原文开始>
// HashMap is alias of AnyAnyMap.
<原文结束>

# <翻译开始>
// HashMap 是 AnyAnyMap 的别名。
# <翻译结束>


<原文开始>
// New creates and returns an empty hash map.
// The parameter `safe` is used to specify whether using map in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// New 创建并返回一个空的哈希表。
// 参数`safe`用于指定是否使用线程安全的map，默认为false。
# <翻译结束>


<原文开始>
// NewFrom creates and returns a hash map from given map `data`.
// Note that, the param `data` map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewFrom 创建并返回一个由给定的 `data` 地图生成的哈希映射。
// 注意，参数 `data` 中的地图将被设置为底层数据地图（无深度复制），
// 当在外部修改该映射时，可能会存在一些并发安全问题。
// 参数 `safe` 用于指定是否在并发环境下使用安全的树结构，默认情况下为 false。
# <翻译结束>


<原文开始>
// NewHashMap creates and returns an empty hash map.
// The parameter `safe` is used to specify whether using map in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewHashMap 创建并返回一个空的哈希映射。
// 参数 `safe` 用于指定是否使用线程安全的 map，其默认值为 false。
# <翻译结束>


<原文开始>
// NewHashMapFrom creates and returns a hash map from given map `data`.
// Note that, the param `data` map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewHashMapFrom 通过给定的 map `data` 创建并返回一个哈希映射。
// 注意，参数 `data` 中的地图将被设置为底层数据映射（非深度拷贝），如果在外部修改此映射，可能会存在一些并发安全问题。
// 参数 `safe` 用于指定是否在并发安全场景下使用树结构，默认情况下为 false。
# <翻译结束>


<原文开始>
// Map is alias of AnyAnyMap.
<原文结束>

# <翻译开始>
// Map 是 AnyAnyMap 的别名。
# <翻译结束>

