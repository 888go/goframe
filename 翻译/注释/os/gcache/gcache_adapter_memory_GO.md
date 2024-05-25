
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
// AdapterMemory is an adapter implements using memory.
<原文结束>

# <翻译开始>
// AdapterMemory是一个适配器，它实现了使用内存。 md5:1058c2331fc6bbaa
# <翻译结束>


<原文开始>
	// cap limits the size of the cache pool.
	// If the size of the cache exceeds the cap,
	// the cache expiration process performs according to the LRU algorithm.
	// It is 0 in default which means no limits.
<原文结束>

# <翻译开始>
// cap 限制了缓存池的大小。
// 如果缓存的大小超过了 cap，
// 则按照 LRU（最近最少使用）算法进行缓存淘汰过程。
// 默认值为 0，表示没有大小限制。
// md5:70436dcd07b73070
# <翻译结束>


<原文开始>
// data is the underlying cache data which is stored in a hash table.
<原文结束>

# <翻译开始>
// data 是底层的缓存数据，它存储在一个哈希表中。 md5:7cfaf636328aa0e7
# <翻译结束>


<原文开始>
// expireTimes is the expiring key to its timestamp mapping, which is used for quick indexing and deleting.
<原文结束>

# <翻译开始>
// expireTimes是过期键到其时间戳的映射，用于快速索引和删除。 md5:5e7fa0cd3e17ed6c
# <翻译结束>


<原文开始>
// expireSets is the expiring timestamp to its key set mapping, which is used for quick indexing and deleting.
<原文结束>

# <翻译开始>
// expireSets 是过期时间戳到其键集合的映射，用于快速索引和删除。 md5:d2c25eb345e1ea19
# <翻译结束>


<原文开始>
// lru is the LRU manager, which is enabled when attribute cap > 0.
<原文结束>

# <翻译开始>
// lru 是 LRU（Least Recently Used）管理器，当属性 cap 大于 0 时启用。 md5:182c6471c0b4b317
# <翻译结束>


<原文开始>
// lruGetList is the LRU history according to Get function.
<原文结束>

# <翻译开始>
// lruGetList是根据Get函数的LRU历史记录。 md5:0ad54aeec8e8c762
# <翻译结束>


<原文开始>
// eventList is the asynchronous event list for internal data synchronization.
<原文结束>

# <翻译开始>
// eventList 是内部数据同步的异步事件列表。 md5:48cbe56e8d02ee7f
# <翻译结束>


<原文开始>
// closed controls the cache closed or not.
<原文结束>

# <翻译开始>
// closed 控制缓存是否已关闭。 md5:8ebf4858be3c0e42
# <翻译结束>


<原文开始>
// Expire timestamp in milliseconds.
<原文结束>

# <翻译开始>
// 过期时间戳，单位为毫秒。 md5:d7096ed51593fa59
# <翻译结束>


<原文开始>
// Expire time in milliseconds.
<原文结束>

# <翻译开始>
// 过期时间，以毫秒为单位。 md5:baebc3abd37be203
# <翻译结束>


<原文开始>
	// defaultMaxExpire is the default expire time for no expiring items.
	// It equals to math.MaxInt64/1000000.
<原文结束>

# <翻译开始>
// defaultMaxExpire是不设置过期时间的默认过期时间。
// 它等于math.MaxInt64除以1000000。
// md5:75ccaa3b4b490a54
# <翻译结束>


<原文开始>
// NewAdapterMemory creates and returns a new memory cache object.
<原文结束>

# <翻译开始>
// NewAdapterMemory 创建并返回一个新的内存缓存对象。 md5:188f107c550c0b2e
# <翻译结束>


<原文开始>
	// Here may be a "timer leak" if adapter is manually changed from memory adapter.
	// Do not worry about this, as adapter is less changed, and it does nothing if it's not used.
<原文结束>

# <翻译开始>
// 如果适配器手动从内存适配器更改，这里可能存在“计时器泄露”。
// 但不必担心这个问题，因为适配器的变更较少，并且如果未被使用，它也不会做什么。
// md5:0d85b615ef8507fb
# <翻译结束>


<原文开始>
// Set sets cache with `key`-`value` pair, which is expired after `duration`.
//
// It does not expire if `duration` == 0.
// It deletes the keys of `data` if `duration` < 0 or given `value` is nil.
<原文结束>

# <翻译开始>
// Set 使用键值对 `key`-`value` 设置缓存，该缓存在 `duration` 时间后过期。
//
// 如果 `duration` 等于 0，则不会过期。
// 如果 `duration` 小于 0 或者给定的 `value` 为 nil，它将删除 `data` 中的键。
// md5:7faea7b643bffd7c
# <翻译结束>


<原文开始>
// SetMap batch sets cache with key-value pairs by `data` map, which is expired after `duration`.
//
// It does not expire if `duration` == 0.
// It deletes the keys of `data` if `duration` < 0 or given `value` is nil.
<原文结束>

# <翻译开始>
// SetMap 批量设置缓存，使用 `data` 映射（键值对）的方式，其在 `duration` 后过期。
//
// 如果 `duration` 等于 0，则不会过期。
// 如果 `duration` 小于 0 或给定的 `value` 为 `nil`，则会删除 `data` 中的键。
// md5:a09a11cd5d9d21e6
# <翻译结束>


<原文开始>
// SetIfNotExist sets cache with `key`-`value` pair which is expired after `duration`
// if `key` does not exist in the cache. It returns true the `key` does not exist in the
// cache, and it sets `value` successfully to the cache, or else it returns false.
//
// It does not expire if `duration` == 0.
// It deletes the `key` if `duration` < 0 or given `value` is nil.
<原文结束>

# <翻译开始>
// SetIfNotExist 如果缓存中不存在`key`，则设置过期时间为`duration`的`key`-`value`对。如果成功将`value`设置到缓存中，它会返回`true`，表示`key`在缓存中不存在；否则返回`false`。
// 
// 如果`duration`为0，缓存不会过期。
// 如果`duration`小于0或给定的`value`为`nil`，它会删除`key`。
// md5:38aa90beb53ed441
# <翻译结束>


<原文开始>
// SetIfNotExistFunc sets `key` with result of function `f` and returns true
// if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.
//
// The parameter `value` can be type of `func() interface{}`, but it does nothing if its
// result is nil.
//
// It does not expire if `duration` == 0.
// It deletes the `key` if `duration` < 0 or given `value` is nil.
<原文结束>

# <翻译开始>
// SetIfNotExistFunc 如果`key`不存在于缓存中，则使用函数`f`的结果设置`key`并返回true。
// 否则，如果`key`已存在，则不做任何操作并返回false。
//
// 参数`value`可以是类型为`func() interface{}`的函数，
// 但如果其结果为nil，则不会执行任何操作。
//
// 如果`duration`等于0，则不设置过期时间。
// 如果`duration`小于0或给定的`value`为nil，则删除该`key`。
// md5:8300c80b9bab735d
# <翻译结束>


<原文开始>
// SetIfNotExistFuncLock sets `key` with result of function `f` and returns true
// if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.
//
// It does not expire if `duration` == 0.
// It deletes the `key` if `duration` < 0 or given `value` is nil.
//
// Note that it differs from function `SetIfNotExistFunc` is that the function `f` is executed within
// writing mutex lock for concurrent safety purpose.
<原文结束>

# <翻译开始>
// SetIfNotExistFuncLock 当`key`在缓存中不存在时，使用函数`f`的结果设置`key`，并返回true。
// 如果`key`已经存在，则不执行任何操作并返回false。
//
// 如果`duration`等于0，则不会过期。
// 如果`duration`小于0或给定的`value`为nil，将删除`key`。
//
// 注意，它与函数`SetIfNotExistFunc`的区别在于，函数`f`在写入互斥锁内部执行，以保证并发安全性。
// md5:629e13ace9eaf720
# <翻译结束>


<原文开始>
// Get retrieves and returns the associated value of given `key`.
// It returns nil if it does not exist, or its value is nil, or it's expired.
// If you would like to check if the `key` exists in the cache, it's better using function Contains.
<原文结束>

# <翻译开始>
// Get 从缓存中检索并返回给定 `key` 的关联值。如果不存在、值为nil或已过期，它将返回nil。如果你想检查`key`是否存在于缓存中，建议使用Contains函数。
// md5:f78c30f8338ce106
# <翻译结束>


<原文开始>
// Adding to LRU history if LRU feature is enabled.
<原文结束>

# <翻译开始>
// 如果启用了LRU功能，则将其添加到LRU历史记录中。 md5:01c169ae5b2999b0
# <翻译结束>


<原文开始>
// GetOrSet retrieves and returns the value of `key`, or sets `key`-`value` pair and
// returns `value` if `key` does not exist in the cache. The key-value pair expires
// after `duration`.
//
// It does not expire if `duration` == 0.
// It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing
// if `value` is a function and the function result is nil.
<原文结束>

# <翻译开始>
// GetOrSet 获取并返回`key`对应的值，如果`key`在缓存中不存在，则设置`key`-`value`对并返回`value`。
// 这对键值将在指定的`duration`后过期。
//
// 如果`duration`为0，则不会过期。
// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但若`value`是一个函数且函数结果为nil，它则不做任何操作。
// md5:b8646fcb99c81de9
# <翻译结束>


<原文开始>
// GetOrSetFunc retrieves and returns the value of `key`, or sets `key` with result of
// function `f` and returns its result if `key` does not exist in the cache. The key-value
// pair expires after `duration`.
//
// It does not expire if `duration` == 0.
// It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing
// if `value` is a function and the function result is nil.
<原文结束>

# <翻译开始>
// GetOrSetFunc 获取并返回`key`的值，如果缓存中不存在`key`，则使用函数`f`的结果设置`key`并返回该结果。键值对在`duration`时间后过期。
//
// 如果`duration`等于0，则不会过期。
// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但若`value`是一个函数且其结果为nil，则不执行任何操作。
// md5:822486c86baa87d1
# <翻译结束>


<原文开始>
// GetOrSetFuncLock retrieves and returns the value of `key`, or sets `key` with result of
// function `f` and returns its result if `key` does not exist in the cache. The key-value
// pair expires after `duration`.
//
// It does not expire if `duration` == 0.
// It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing
// if `value` is a function and the function result is nil.
//
// Note that it differs from function `GetOrSetFunc` is that the function `f` is executed within
// writing mutex lock for concurrent safety purpose.
<原文结束>

# <翻译开始>
// GetOrSetFuncLock 获取并返回键`key`的值，或者如果`key`在缓存中不存在，则使用函数`f`的结果设置`key`，并返回其结果。键值对在`duration`后过期。
// 
// 如果`duration`为0，它不会过期。
// 如果`duration`小于0或给定的`value`为nil，它会删除`key`；但如果`value`是一个函数并且函数结果为nil，它将不执行任何操作。
// 
// 注意，它与`GetOrSetFunc`函数不同，函数`f`是在写入互斥锁保护下执行的，以确保并发安全。
// md5:3e49c54e5e0c2857
# <翻译结束>


<原文开始>
// Contains checks and returns true if `key` exists in the cache, or else returns false.
<原文结束>

# <翻译开始>
// Contains 检查并返回如果 `key` 在缓存中存在则为真，否则为假。 md5:4ff234995709b9ab
# <翻译结束>


<原文开始>
// GetExpire retrieves and returns the expiration of `key` in the cache.
//
// Note that,
// It returns 0 if the `key` does not expire.
// It returns -1 if the `key` does not exist in the cache.
<原文结束>

# <翻译开始>
// GetExpire 从缓存中检索并返回 `key` 的过期时间。
// 
// 注意，
// 如果 `key` 没有过期，它将返回 0。
// 如果 `key` 不在缓存中，它将返回 -1。
// md5:d80ce12df8668b97
# <翻译结束>


<原文开始>
// Remove deletes one or more keys from cache, and returns its value.
// If multiple keys are given, it returns the value of the last deleted item.
<原文结束>

# <翻译开始>
// Remove 从缓存中删除一个或多个键，并返回其值。
// 如果给出了多个键，它将返回最后删除项的值。
// md5:d3b1c8af168b0ebf
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
// Size returns the size of the cache.
<原文结束>

# <翻译开始>
// Size 返回缓存的大小。 md5:c939a4ed87cd79ce
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
// Clear clears all data of the cache.
// Note that this function is sensitive and should be carefully used.
<原文结束>

# <翻译开始>
// Clear 清空缓存中的所有数据。
// 注意，此函数涉及敏感操作，应谨慎使用。
// md5:9212cab88870d3df
# <翻译结束>


<原文开始>
// Close closes the cache.
<原文结束>

# <翻译开始>
// Close 关闭缓存。 md5:c1a9d7a347be93a8
# <翻译结束>


<原文开始>
// doSetWithLockCheck sets cache with `key`-`value` pair if `key` does not exist in the
// cache, which is expired after `duration`.
//
// It does not expire if `duration` == 0.
// The parameter `value` can be type of <func() interface{}>, but it does nothing if the
// function result is nil.
//
// It doubly checks the `key` whether exists in the cache using mutex writing lock
// before setting it to the cache.
<原文结束>

# <翻译开始>
// doSetWithLockCheck 如果缓存中不存在键为`key`的项，将`key-value`对设置到缓存中，且该项的过期时间为`duration`。
//
// 如果`duration`为0，则不过期。参数`value`可以是类型为`func() interface{}`的函数，但如果函数结果为nil，则不执行任何操作。
//
// 在将`key-value`对设置到缓存之前，它会使用写入锁双重检查`key`是否已存在于缓存中。
// md5:17967ab63e2b200c
# <翻译结束>


<原文开始>
// getInternalExpire converts and returns the expiration time with given expired duration in milliseconds.
<原文结束>

# <翻译开始>
// getInternalExpire 将给定的过期毫秒数转换并返回过期时间。 md5:176ebdcfb2a89f78
# <翻译结束>


<原文开始>
// makeExpireKey groups the `expire` in milliseconds to its according seconds.
<原文结束>

# <翻译开始>
// makeExpireKey 将毫秒级的 `expire` 值归类到其对应的秒级单位。 md5:40d29c22e827fc9e
# <翻译结束>


<原文开始>
// syncEventAndClearExpired does the asynchronous task loop:
// 1. Asynchronously process the data in the event list,
// and synchronize the results to the `expireTimes` and `expireSets` properties.
// 2. Clean up the expired key-value pair data.
<原文结束>

# <翻译开始>
// syncEventAndClearExpired 执行异步任务循环：
// 1. 异步处理事件列表中的数据，
// 并将结果同步到 `expireTimes` 和 `expireSets` 属性。
// 2. 清理过期的键值对数据。
// md5:ce52abd32c5f232e
# <翻译结束>


<原文开始>
	// ========================
	// Data Synchronization.
	// ========================
<原文结束>

# <翻译开始>
// ========================
// 数据同步。
// ========================
// md5:a7203ea428e10983
# <翻译结束>


<原文开始>
// Fetching the old expire set.
<原文结束>

# <翻译开始>
// 获取旧的过期集合。 md5:e6633f31f39e1499
# <翻译结束>


<原文开始>
// Calculating the new expiration time set.
<原文结束>

# <翻译开始>
// 计算新的过期时间设置。 md5:57b48d53f5270f91
# <翻译结束>


<原文开始>
// Updating the expired time for <event.k>.
<原文结束>

# <翻译开始>
// 更新<event.k>的过期时间。 md5:f04ccde84655d99f
# <翻译结束>


<原文开始>
// Adding the key the LRU history by writing operations.
<原文结束>

# <翻译开始>
// 通过写操作将键添加到LRU历史中。 md5:ca17e775d3b31310
# <翻译结束>


<原文开始>
// Processing expired keys from LRU.
<原文结束>

# <翻译开始>
// 从最近最少使用（Least Recently Used，LRU）缓存中处理过期的键。 md5:c555319093b1296e
# <翻译结束>


<原文开始>
	// ========================
	// Data Cleaning up.
	// ========================
<原文结束>

# <翻译开始>
// ========================
// 数据清理。
// ========================
// md5:c845ec8cb41f31ac
# <翻译结束>


<原文开始>
// Iterating the set to delete all keys in it.
<原文结束>

# <翻译开始>
// 遍历集合以删除其中的所有键。 md5:de77c90f243260c0
# <翻译结束>


<原文开始>
// Deleting the set after all of its keys are deleted.
<原文结束>

# <翻译开始>
// 在删除所有键之后，删除集合。 md5:d34b6cd2767c7800
# <翻译结束>


<原文开始>
// clearByKey deletes the key-value pair with given `key`.
// The parameter `force` specifies whether doing this deleting forcibly.
<原文结束>

# <翻译开始>
// clearByKey 删除给定`key`的键值对。参数`force`指定是否强制执行删除操作。
// md5:5b26398959f735ad
# <翻译结束>


<原文开始>
// Doubly check before really deleting it from cache.
<原文结束>

# <翻译开始>
// 在从缓存中真正删除之前，再双检查一次。 md5:53767fc86cbfbf5e
# <翻译结束>


<原文开始>
// Deleting its expiration time from `expireTimes`.
<原文结束>

# <翻译开始>
// 从`expireTimes`中删除其过期时间。 md5:d2320f7b4a5f1c26
# <翻译结束>

