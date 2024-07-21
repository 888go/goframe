
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
// Adapter is the core adapter for cache features implements.
//
// Note that the implementer itself should guarantee the concurrent safety of these functions.
<原文结束>

# <翻译开始>
// Adapter是缓存功能的核心适配器。
// 
// 注意，实现者本身应确保这些函数的并发安全性。
// md5:cd91041442c2fdbf
# <翻译结束>


<原文开始>
	// Set sets cache with `key`-`value` pair, which is expired after `duration`.
	//
	// It does not expire if `duration` == 0.
	// It deletes the keys of `data` if `duration` < 0 or given `value` is nil.
<原文结束>

# <翻译开始>
	// Set 使用 `key`-`value` 对设置缓存，该缓存在 `duration` 时间后过期。
	//
	// 如果 `duration` == 0，则不设置过期时间。
	// 如果 `duration` < 0 或给定的 `value` 为 nil，则删除 `data` 的键。
	// md5:3f5918d3cc5c36fd
# <翻译结束>


<原文开始>
	// SetMap batch sets cache with key-value pairs by `data` map, which is expired after `duration`.
	//
	// It does not expire if `duration` == 0.
	// It deletes the keys of `data` if `duration` < 0 or given `value` is nil.
<原文结束>

# <翻译开始>
	// SetMap 批量设置缓存，使用 `data` 映射中的键值对，这些缓存在 `duration` 时间后过期。
	//
	// 如果 `duration` == 0，则不会过期。
	// 如果 `duration` < 0 或者给定的 `value` 为 nil，将删除 `data` 中的键。
	// md5:029757e42001dd48
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
	// md5:a442e240e2ddb849
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
	// SetIfNotExistFunc 如果缓存中不存在`key`，则使用函数`f`的结果设置`key`，并返回true。如果`key`已存在，则不做任何操作，返回false。
	//
	// 参数`value`可以是`func() interface{}`类型，但如果其结果为nil，则不执行任何操作。
	//
	// 如果`duration`为0，表示永不过期。如果`duration`小于0或给定的`value`为nil，则删除`key`。
	// md5:33f0e2bb534c4ac4
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
	// SetIfNotExistFuncLock 如果`key`在缓存中不存在，则使用函数`f`的结果设置`key`并返回true。
	// 如果`key`已存在，则不做任何操作并返回false。
	//
	// 如果`duration`为0，则不设置过期时间。
	// 如果`duration`小于0或给定的`value`为nil，则删除`key`。
	//
	// 注意，它与函数`SetIfNotExistFunc`的不同之处在于，为了并发安全，函数`f`在写入互斥锁内部执行。
	// md5:906879fb08827346
# <翻译结束>


<原文开始>
	// Get retrieves and returns the associated value of given `key`.
	// It returns nil if it does not exist, or its value is nil, or it's expired.
	// If you would like to check if the `key` exists in the cache, it's better using function Contains.
<原文结束>

# <翻译开始>
	// Get 获取并返回给定`key`关联的值。
	// 如果键不存在、其值为nil或已过期，它将返回nil。
	// 如果你想检查`key`是否在缓存中存在，最好使用Contains函数。
	// md5:a04abebd42f9db26
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
	// GetOrSet 从缓存中获取并返回`key`的值，如果`key`不存在，则设置`key-value`对，并返回`value`。缓存中的键值对在`duration`后过期。
	// 如果`duration`为0，则不会过期。
	// 如果`duration`小于0或给定的`value`为nil，它会删除`key`；但如果`value`是一个函数且函数结果为nil，则不做任何操作。
	// md5:a9525aacd8a5324e
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
	// GetOrSetFunc 从缓存中获取并返回键`key`的值，如果键不存在，则使用函数`f`的结果设置键并返回该结果。键值对在`duration`后过期。
	// 
	// 如果`duration`为0，表示永不过期。
	// 如果`duration`小于0或给定的`value`为nil，它会删除键`key`。但如果`value`是一个函数并且函数结果为nil，它不会做任何操作。
	// md5:57a987bd75623802
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
	// GetOrSetFuncLock 获取并返回`key`的值，如果`key`在缓存中不存在，则使用函数`f`的结果设置`key`并返回该结果。
	// 键值对将在`duration`时间后过期。
	//
	// 如果`duration`为0，则不设置过期时间。
	// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但若`value`是一个函数且函数结果为nil时，它不做任何操作。
	//
	// 需要注意的是，此函数与`GetOrSetFunc`的区别在于，函数`f`是在写入互斥锁内部执行的，以确保并发安全。
	// md5:b0a08f256bf6fcfc
# <翻译结束>


<原文开始>
// Contains checks and returns true if `key` exists in the cache, or else returns false.
<原文结束>

# <翻译开始>
// Contains 检查并返回如果 `key` 在缓存中存在则为真，否则为假。 md5:4ff234995709b9ab
# <翻译结束>


<原文开始>
// Size returns the number of items in the cache.
<原文结束>

# <翻译开始>
// Size 返回缓存中的项目数量。 md5:2122f80de9340261
# <翻译结束>


<原文开始>
	// Data returns a copy of all key-value pairs in the cache as map type.
	// Note that this function may lead lots of memory usage, you can implement this function
	// if necessary.
<原文结束>

# <翻译开始>
	// Data返回缓存中所有键值对的副本，以map类型。
	// 注意，此函数可能会导致大量内存使用。如果需要，您可以实现这个函数。
	// md5:96cf9c57d77ba2dd
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
	// md5:28635aef7c0fc7a9
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
	// md5:f1bb94e5134bebed
# <翻译结束>


<原文开始>
	// GetExpire retrieves and returns the expiration of `key` in the cache.
	//
	// Note that,
	// It returns 0 if the `key` does not expire.
	// It returns -1 if the `key` does not exist in the cache.
<原文结束>

# <翻译开始>
	// GetExpire 获取并返回缓存中 `key` 的过期时间。
	//
	// 注意：
	// 如果 `key` 没有设置过期时间，它将返回 0。
	// 如果 `key` 在缓存中不存在，它将返回 -1。
	// md5:6a059254c0534a31
# <翻译结束>


<原文开始>
	// Remove deletes one or more keys from cache, and returns its value.
	// If multiple keys are given, it returns the value of the last deleted item.
<原文结束>

# <翻译开始>
	// Remove 从缓存中删除一个或多个键，并返回其值。
	// 如果提供了多个键，它将返回最后一个被删除项的值。
	// md5:6e5f157befbc08c2
# <翻译结束>


<原文开始>
	// Clear clears all data of the cache.
	// Note that this function is sensitive and should be carefully used.
<原文结束>

# <翻译开始>
	// Clear 清空缓存中的所有数据。
	// 注意，此函数涉及敏感操作，应谨慎使用。
	// md5:8f66f62d0fce831a
# <翻译结束>


<原文开始>
// Close closes the cache if necessary.
<原文结束>

# <翻译开始>
// Close如果有必要，关闭缓存。 md5:f9a73a30e4b4b396
# <翻译结束>

