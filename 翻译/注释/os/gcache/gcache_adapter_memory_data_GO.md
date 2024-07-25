
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
// dataMu ensures the concurrent safety of underlying data map.
<原文结束>

# <翻译开始>
// dataMu 确保底层数据映射的并发安全性。 md5:ddcd414a151f3cf2
# <翻译结束>


<原文开始>
// data is the underlying cache data which is stored in a hash table.
<原文结束>

# <翻译开始>
// data 是底层的缓存数据，它存储在一个哈希表中。 md5:7cfaf636328aa0e7
# <翻译结束>


<原文开始>
// Update updates the value of `key` without changing its expiration and returns the old value.
// The returned value `exist` is false if the `key` does not exist in the cache.
//
// It deletes the `key` if given `value` is nil.
// It does nothing if `key` does not exist in the cache.
<原文结束>

# <翻译开始>
// Update 更新`key`的值，不改变其过期时间，并返回旧的值。
// 如果`key`在缓存中不存在，返回的值`exist`为false。
//
// 如果给定的`value`为nil，它会删除`key`。
// 如果`key`不在缓存中，它不会做任何操作。
// md5:6d92816db5b1d3bd
# <翻译结束>


<原文开始>
// UpdateExpire updates the expiration of `key` and returns the old expiration duration value.
//
// It returns -1 and does nothing if the `key` does not exist in the cache.
// It deletes the `key` if `duration` < 0.
<原文结束>

# <翻译开始>
// UpdateExpire 更新键`key`的过期时间，并返回旧的过期持续时间值。
//
// 如果`key`在缓存中不存在，它将返回-1并什么都不做。如果`duration`小于0，它会删除`key`。
// md5:b974907dd46b44be
# <翻译结束>


<原文开始>
// Remove deletes the one or more keys from cache, and returns its value.
// If multiple keys are given, it returns the value of the deleted last item.
<原文结束>

# <翻译开始>
// Remove 从缓存中删除一个或多个键，并返回其值。
// 如果给出了多个键，它将返回最后删除项的值。
// md5:b3f23906b769df08
# <翻译结束>


<原文开始>
// Data returns a copy of all key-value pairs in the cache as map type.
<原文结束>

# <翻译开始>
// Data 返回一个缓存中所有键值对的副本，以映射类型表示。 md5:d88afdf7cfc66604
# <翻译结束>


<原文开始>
// Keys returns all keys in the cache as slice.
<原文结束>

# <翻译开始>
// Keys 返回缓存中所有键的切片。 md5:7ebd9dba01282dc2
# <翻译结束>


<原文开始>
// Values returns all values in the cache as slice.
<原文结束>

# <翻译开始>
// Values 返回缓存中所有的值作为切片。 md5:dc00b32eb8913e9b
# <翻译结束>


<原文开始>
// Size returns the size of the cache.
<原文结束>

# <翻译开始>
// Size 返回缓存的大小。 md5:c939a4ed87cd79ce
# <翻译结束>


<原文开始>
// Clear clears all data of the cache.
// Note that this function is sensitive and should be carefully used.
<原文结束>

# <翻译开始>
// Clear 清空缓存中的所有数据。
// 注意，此函数涉及敏感操作，应谨慎使用。
// md5:9212cab88870d3df
# <翻译结束>


<原文开始>
// SetMap batch sets cache with key-value pairs by `data`, which is expired after `duration`.
//
// It does not expire if `duration` == 0.
// It deletes the keys of `data` if `duration` < 0 or given `value` is nil.
<原文结束>

# <翻译开始>
// SetMap 通过 `data` 批量设置缓存键值对，这些缓存在 `duration` 时间后过期。
//
// 如果 `duration` == 0，则不设置过期时间。
// 如果 `duration` < 0 或者给定的 `value` 为 nil，则删除 `data` 中的键。
// md5:cc6156a6df071b21
# <翻译结束>


<原文开始>
// Compatible with raw function value.
<原文结束>

# <翻译开始>
		// 与原始函数值兼容。 md5:b6980bd817389e7f
# <翻译结束>


<原文开始>
// Doubly check before really deleting it from cache.
<原文结束>

# <翻译开始>
	// 在从缓存中真正删除之前，再双检查一次。 md5:53767fc86cbfbf5e
# <翻译结束>

