
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
// AdapterMemory is an adapter implements using memory.
<原文结束>

# <翻译开始>
// AdapterMemory 是一个使用内存实现的适配器。
# <翻译结束>


<原文开始>
	// cap limits the size of the cache pool.
	// If the size of the cache exceeds the cap,
	// the cache expiration process performs according to the LRU algorithm.
	// It is 0 in default which means no limits.
<原文结束>

# <翻译开始>
// cap 限制缓存池的大小。
// 如果缓存的大小超过 cap，
// 缓存过期过程将根据 LRU 算法执行。
// 默认值为 0，表示无限制。
# <翻译结束>


<原文开始>
// data is the underlying cache data which is stored in a hash table.
<原文结束>

# <翻译开始>
// data 是底层缓存数据，存储在一个哈希表中。
# <翻译结束>


<原文开始>
// expireTimes is the expiring key to its timestamp mapping, which is used for quick indexing and deleting.
<原文结束>

# <翻译开始>
// expireTimes 是一个过期键与其时间戳的映射，用于快速索引和删除。
# <翻译结束>


<原文开始>
// expireSets is the expiring timestamp to its key set mapping, which is used for quick indexing and deleting.
<原文结束>

# <翻译开始>
// expireSets 是一个映射表，用于存储即将过期的时间戳及其对应的键集合。这个映射表用于快速索引和删除操作。
# <翻译结束>


<原文开始>
// lru is the LRU manager, which is enabled when attribute cap > 0.
<原文结束>

# <翻译开始>
// lru 是 LRU（最近最少使用）管理器，当属性 cap 大于 0 时启用。
# <翻译结束>


<原文开始>
// lruGetList is the LRU history according to Get function.
<原文结束>

# <翻译开始>
// lruGetList 是根据 Get 函数实现的 LRU（最近最少使用）历史记录列表。
# <翻译结束>


<原文开始>
// eventList is the asynchronous event list for internal data synchronization.
<原文结束>

# <翻译开始>
// eventList 是用于内部数据同步的异步事件列表。
# <翻译结束>


<原文开始>
// closed controls the cache closed or not.
<原文结束>

# <翻译开始>
// closed 控制缓存是否关闭
# <翻译结束>







<原文开始>
// Expire timestamp in milliseconds.
<原文结束>

# <翻译开始>
// 过期时间戳（毫秒）
# <翻译结束>







<原文开始>
// Expire time in milliseconds.
<原文结束>

# <翻译开始>
// 过期时间（以毫秒为单位）
# <翻译结束>


<原文开始>
	// defaultMaxExpire is the default expire time for no expiring items.
	// It equals to math.MaxInt64/1000000.
<原文结束>

# <翻译开始>
// defaultMaxExpire 是未设置过期时间项目的默认过期时间。
// 它等于 math.MaxInt64/1000000。
# <翻译结束>


<原文开始>
// NewAdapterMemory creates and returns a new memory cache object.
<原文结束>

# <翻译开始>
// NewAdapterMemory 创建并返回一个新的内存缓存对象。
# <翻译结束>


<原文开始>
// Set sets cache with `key`-`value` pair, which is expired after `duration`.
//
// It does not expire if `duration` == 0.
// It deletes the keys of `data` if `duration` < 0 or given `value` is nil.
<原文结束>

# <翻译开始>
// Set 通过 `key`-`value` 对设置缓存，该对在 `duration` 后过期。
//
// 如果 `duration` == 0，则永不过期。
// 如果 `duration` < 0 或提供的 `value` 为 nil，则删除 `data` 的键。
# <翻译结束>


<原文开始>
// SetMap batch sets cache with key-value pairs by `data` map, which is expired after `duration`.
//
// It does not expire if `duration` == 0.
// It deletes the keys of `data` if `duration` < 0 or given `value` is nil.
<原文结束>

# <翻译开始>
// SetMap 批量设置缓存，通过 `data` 参数中的键值对进行设置，并在 `duration` 时间后过期。
//
// 如果 `duration` == 0，则表示永不过期。
// 如果 `duration` < 0 或者给定的 `value` 为 nil，则会删除 `data` 中的键。
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
// SetIfNotExist 若`key`不存在于缓存中，则设置带有`key`-`value`对的缓存，该对在`duration`后过期。
// 如果`key`在缓存中不存在，它将返回true，并成功将`value`设置到缓存中，否则返回false。
//
// 如果`duration` == 0，则不会设置过期时间。
// 如果`duration` < 0 或给定的`value`为nil，则删除`key`。
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
// SetIfNotExistFunc 函数用于设置 `key` 为函数 `f` 的计算结果，并在 `key` 不存在于缓存中时返回 true，
// 否则如果 `key` 已存在，则不做任何操作并返回 false。
//
// 参数 `value` 可以是类型 `func() interface{}`，但如果其结果为 nil，则该函数不会执行任何操作。
//
// 如果 `duration` == 0，则不设置过期时间。
// 如果 `duration` < 0 或给定的 `value` 为 nil，则会删除 `key`。
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
// SetIfNotExistFuncLock 将通过函数 `f` 计算的结果设置为 `key` 的值，并在以下情况下返回 true：
// 1. 如果 `key` 不存在于缓存中，则设置并返回 true。
// 2. 否则，如果 `key` 已经存在，则不做任何操作并返回 false。
// 若 `duration` 等于 0，则不设置过期时间。
// 若 `duration` 小于 0 或提供的 `value` 为 nil，则删除 `key`。
// 注意，此方法与函数 `SetIfNotExistFunc` 的不同之处在于，
// 函数 `f` 在写入互斥锁保护下执行，以确保并发安全。
# <翻译结束>


<原文开始>
// Get retrieves and returns the associated value of given `key`.
// It returns nil if it does not exist, or its value is nil, or it's expired.
// If you would like to check if the `key` exists in the cache, it's better using function Contains.
<原文结束>

# <翻译开始>
// Get 方法通过给定的 `key` 获取并返回关联的值。
// 若该 `key` 对应的值不存在，或者其值为 nil，或已过期，则返回 nil。
// 如果你想检查 `key` 是否存在于缓存中，最好使用 Contains 函数。
# <翻译结束>


<原文开始>
// Adding to LRU history if LRU feature is enabled.
<原文结束>

# <翻译开始>
// 如果启用了LRU功能，则添加到LRU历史记录中。
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
// GetOrSet 获取并返回键`key`的值，如果`key`在缓存中不存在，则设置`key`-`value`对并返回`value`。
// 键值对在`duration`时间后过期。
//
// 如果`duration` == 0，则不会过期。
// 如果`duration` < 0 或者给定的`value`为nil，则删除`key`，但如果`value`是一个函数且函数结果为nil，则不做任何操作。
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
// GetOrSetFunc 函数用于获取并返回 `key` 对应的值，如果 `key` 不存在于缓存中，则使用函数 `f` 的结果设置 `key` 并返回其结果。
// 这对键值在 `duration` 时间后将自动过期。
//
// 如果 `duration` 等于 0，则表示该键值对永不过期。
// 如果 `duration` 小于 0 或者给定的 `value` 为 nil，则会删除 `key`，但如果 `value` 是一个函数且函数结果为 nil，则不做任何操作。
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
// GetOrSetFuncLock 从缓存中获取并返回`key`的值，如果`key`不存在，则使用函数`f`的结果设置`key`并返回其结果。键值对在`duration`时间后过期。
// 如果`duration`为0，则它不会过期。
// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但如果`value`是一个函数且函数结果为nil，则不做任何操作。
// 注意，该方法与函数`GetOrSetFunc`的不同之处在于，为了保证并发安全，函数`f`在写入互斥锁内执行。
# <翻译结束>


<原文开始>
// Contains checks and returns true if `key` exists in the cache, or else returns false.
<原文结束>

# <翻译开始>
// Contains 检查并返回 true，如果 `key` 存在于缓存中；否则返回 false。
# <翻译结束>


<原文开始>
// GetExpire retrieves and returns the expiration of `key` in the cache.
//
// Note that,
// It returns 0 if the `key` does not expire.
// It returns -1 if the `key` does not exist in the cache.
<原文结束>

# <翻译开始>
// GetExpire 从缓存中检索并返回`key`的过期时间。
//
// 注意：
// 如果`key`永不过期，则返回0。
// 如果`key`在缓存中不存在，则返回-1。
# <翻译结束>


<原文开始>
// Remove deletes one or more keys from cache, and returns its value.
// If multiple keys are given, it returns the value of the last deleted item.
<原文结束>

# <翻译开始>
// Remove 从缓存中删除一个或多个键，并返回其对应的值。
// 如果提供了多个键，它将返回最后一个被删除项的值。
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
// Size returns the size of the cache.
<原文结束>

# <翻译开始>
// Size 返回缓存的大小。
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
// Clear clears all data of the cache.
// Note that this function is sensitive and should be carefully used.
<原文结束>

# <翻译开始>
// Clear 清除缓存中的所有数据。
// 注意：此函数较为敏感，应谨慎使用。
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
// doSetWithLockCheck 函数用于在缓存中设置键值对 `key`-`value`，如果 `key` 不存在于缓存中且设置了过期时间 `duration`。
//
// 如果 `duration` 等于0，则不会设置过期时间。
// 参数 `value` 可以是 <func() interface{}> 类型，但如果函数结果为 nil，则不做任何操作。
//
// 在设置缓存前，它会通过互斥写锁对 `key` 是否存在于缓存进行双重检查。
# <翻译结束>


<原文开始>
// getInternalExpire converts and returns the expiration time with given expired duration in milliseconds.
<原文结束>

# <翻译开始>
// getInternalExpire 将给定的以毫秒为单位的过期时长转换并返回其对应的过期时间。
# <翻译结束>


<原文开始>
// makeExpireKey groups the `expire` in milliseconds to its according seconds.
<原文结束>

# <翻译开始>
// makeExpireKey 将以毫秒为单位的 `expire` 分组到相应的秒数。
# <翻译结束>


<原文开始>
// syncEventAndClearExpired does the asynchronous task loop:
// 1. Asynchronously process the data in the event list,
// and synchronize the results to the `expireTimes` and `expireSets` properties.
// 2. Clean up the expired key-value pair data.
<原文结束>

# <翻译开始>
// syncEventAndClearExpired 执行异步任务循环:
// 1. 异步处理事件列表中的数据，
// 并将处理结果同步到 `expireTimes` 和 `expireSets` 属性上。
// 2. 清理已过期的键值对数据。
# <翻译结束>


<原文开始>
	// ========================
	// Data Synchronization.
	// ========================
<原文结束>

# <翻译开始>
// ========================================
// 数据同步.
// ========================================
# <翻译结束>


<原文开始>
// Fetching the old expire set.
<原文结束>

# <翻译开始>
// 获取旧的过期集合。
# <翻译结束>


<原文开始>
// Calculating the new expiration time set.
<原文结束>

# <翻译开始>
// 计算新设置的过期时间
# <翻译结束>


<原文开始>
// Updating the expired time for <event.k>.
<原文结束>

# <翻译开始>
// 更新<event.k>的过期时间。
# <翻译结束>


<原文开始>
// Adding the key the LRU history by writing operations.
<原文结束>

# <翻译开始>
// 通过写操作将键添加到LRU历史记录中。
# <翻译结束>


<原文开始>
// Processing expired keys from LRU.
<原文结束>

# <翻译开始>
// 处理LRU中已过期的键。
# <翻译结束>


<原文开始>
	// ========================
	// Data Cleaning up.
	// ========================
<原文结束>

# <翻译开始>
// ==================================
// 数据清理
// ==================================
# <翻译结束>


<原文开始>
// Iterating the set to delete all keys in it.
<原文结束>

# <翻译开始>
// 遍历集合以删除其中的所有键。
# <翻译结束>


<原文开始>
// Deleting the set after all of its keys are deleted.
<原文结束>

# <翻译开始>
// 在其所有键都被删除后，删除该集合。
# <翻译结束>


<原文开始>
// clearByKey deletes the key-value pair with given `key`.
// The parameter `force` specifies whether doing this deleting forcibly.
<原文结束>

# <翻译开始>
// clearByKey 通过给定的 `key` 删除键值对。
// 参数 `force` 指定是否强制执行此删除操作。
# <翻译结束>


<原文开始>
// Doubly check before really deleting it from cache.
<原文结束>

# <翻译开始>
// 在真正从缓存中删除之前进行双重检查。
# <翻译结束>


<原文开始>
// Deleting its expiration time from `expireTimes`.
<原文结束>

# <翻译开始>
// 从`expireTimes`中删除其过期时间。
# <翻译结束>







<原文开始>
// Internal cache item.
<原文结束>

# <翻译开始>
// 内部缓存项。
# <翻译结束>


<原文开始>
// Internal event item.
<原文结束>

# <翻译开始>
// 内部事件项
# <翻译结束>


<原文开始>
// Close closes the cache.
<原文结束>

# <翻译开始>
// Close 关闭缓存。
# <翻译结束>


<原文开始>
// Deleting it from LRU.
<原文结束>

# <翻译开始>
// 从LRU中删除它
# <翻译结束>

