
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
// dataMu ensures the concurrent safety of underlying data map.
<原文结束>

# <翻译开始>
// dataMu 用于确保底层数据映射的并发安全性。
# <翻译结束>


<原文开始>
// data is the underlying cache data which is stored in a hash table.
<原文结束>

# <翻译开始>
// data 是底层缓存数据，存储在一个哈希表中。
# <翻译结束>


<原文开始>
// Update updates the value of `key` without changing its expiration and returns the old value.
// The returned value `exist` is false if the `key` does not exist in the cache.
//
// It deletes the `key` if given `value` is nil.
// It does nothing if `key` does not exist in the cache.
<原文结束>

# <翻译开始>
// Update 更新`key`的值，但不改变其过期时间，并返回旧值。
// 返回的布尔值`exist`，如果`key`在缓存中不存在，则为false。
//
// 如果给定的`value`为nil，则删除`key`。
// 若`key`在缓存中不存在，则不做任何操作。
# <翻译结束>


<原文开始>
// UpdateExpire updates the expiration of `key` and returns the old expiration duration value.
//
// It returns -1 and does nothing if the `key` does not exist in the cache.
// It deletes the `key` if `duration` < 0.
<原文结束>

# <翻译开始>
// UpdateExpire 更新键 `key` 的过期时间，并返回旧的过期持续时长值。
//
// 若 `key` 不存在于缓存中，则返回 -1 并不做任何操作。
// 若 `duration` 小于 0，则删除 `key`。
# <翻译结束>


<原文开始>
// Remove deletes the one or more keys from cache, and returns its value.
// If multiple keys are given, it returns the value of the deleted last item.
<原文结束>

# <翻译开始>
// Remove 从缓存中删除一个或多个键，并返回其对应的值。
// 如果提供了多个键，它将返回被删除的最后一个项目的值。
# <翻译结束>


<原文开始>
// Data returns a copy of all key-value pairs in the cache as map type.
<原文结束>

# <翻译开始>
// Data 返回缓存中所有键值对的副本，类型为 map。
# <翻译结束>


<原文开始>
// Keys returns all keys in the cache as slice.
<原文结束>

# <翻译开始>
// Keys 返回缓存中的所有键作为切片。
# <翻译结束>


<原文开始>
// Values returns all values in the cache as slice.
<原文结束>

# <翻译开始>
// Values 返回缓存中的所有值作为一个切片。
# <翻译结束>


<原文开始>
// Size returns the size of the cache.
<原文结束>

# <翻译开始>
// Size 返回缓存的大小。
# <翻译结束>


<原文开始>
// Clear clears all data of the cache.
// Note that this function is sensitive and should be carefully used.
<原文结束>

# <翻译开始>
// Clear 清除缓存中的所有数据。
// 注意：此函数较为敏感，应谨慎使用。
# <翻译结束>


<原文开始>
// SetMap batch sets cache with key-value pairs by `data`, which is expired after `duration`.
//
// It does not expire if `duration` == 0.
// It deletes the keys of `data` if `duration` < 0 or given `value` is nil.
<原文结束>

# <翻译开始>
// SetMap 通过 `data` 批量设置缓存键值对，缓存将在 `duration` 后过期。
//
// 如果 `duration` == 0，则表示永不过期。
// 若 `duration` < 0 或提供的 `value` 为 nil，则删除 `data` 中的键。
# <翻译结束>


<原文开始>
// Compatible with raw function value.
<原文结束>

# <翻译开始>
// 与原始函数值兼容。
# <翻译结束>


<原文开始>
// Doubly check before really deleting it from cache.
<原文结束>

# <翻译开始>
// 在真正从缓存中删除之前进行双重检查。
# <翻译结束>

