
<原文开始>
// Copyright 2020 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 (c) 2020 gogf 作者(https://github.com/gogf/gf)。保留所有权利。
//
// 本源代码形式遵循MIT许可协议。若未随此文件一同分发MIT许可协议副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:7063305469ff40c2
# <翻译结束>


<原文开始>
// AdapterRedis is the gcache adapter implements using Redis server.
<原文结束>

# <翻译开始>
// AdapterRedis 是使用 Redis 服务器实现的 gcache 适配器。. md5:7ac226ec6d59930e
# <翻译结束>


<原文开始>
// NewAdapterRedis creates and returns a new memory cache object.
<原文结束>

# <翻译开始>
// NewAdapterRedis 创建并返回一个新的内存缓存对象。. md5:ac9ad598fcd2adbb
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
// Execute the function and retrieve the result.
<原文结束>

# <翻译开始>
// 执行函数并获取结果。. md5:1443cd3171693ec8
# <翻译结束>


<原文开始>
// Compatible with raw function value.
<原文结束>

# <翻译开始>
// 与原始函数值兼容。. md5:b6980bd817389e7f
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
// Get retrieves and returns the associated value of given <key>.
// It returns nil if it does not exist or its value is nil.
<原文结束>

# <翻译开始>
// Get 通过给定的 <key> 获取并返回关联的值。如果不存在或其值为 nil，则返回 nil。
// md5:ecb61eca16fb4324
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
// Contains 检查并返回如果 `key` 在缓存中存在则为真，否则为假。. md5:4ff234995709b9ab
# <翻译结束>


<原文开始>
// Size returns the number of items in the cache.
<原文结束>

# <翻译开始>
// Size 返回缓存中的项目数量。. md5:2122f80de9340261
# <翻译结束>


<原文开始>
// Data returns a copy of all key-value pairs in the cache as map type.
// Note that this function may lead lots of memory usage, you can implement this function
// if necessary.
<原文结束>

# <翻译开始>
// Data 返回缓存中所有键值对的副本，以映射类型形式呈现。
// 注意：此函数可能会占用大量内存，请根据需要决定是否实现该功能。
// md5:c44cdbd9b10ab98f
# <翻译结束>


<原文开始>
// Keys returns all keys in the cache as slice.
<原文结束>

# <翻译开始>
// Keys 返回缓存中所有键的切片。. md5:7ebd9dba01282dc2
# <翻译结束>


<原文开始>
// Values returns all values in the cache as slice.
<原文结束>

# <翻译开始>
// Values 返回缓存中所有的值作为切片。. md5:dc00b32eb8913e9b
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
// update ttl -> pttl(millisecond)
<原文结束>

# <翻译开始>
// update ttl -> 更新时间戳到毫秒级的pttl. md5:a9616c495a46fa50
# <翻译结束>


<原文开始>
// It does not exist or expired.
<原文结束>

# <翻译开始>
// 它不存在或已过期。. md5:a51ac96e5909ca59
# <翻译结束>


<原文开始>
		// update SetEX -> SET PX Option(millisecond)
		// Starting with Redis version 2.6.12: Added the EX, PX, NX and XX options.
<原文结束>

# <翻译开始>
// 更新 SetEX -> 设置PX选项（毫秒）
// 从Redis版本2.6.12开始：添加了EX、PX、NX和XX选项。
// md5:490be86df7cc2df5
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
// Remove deletes the one or more keys from cache, and returns its value.
// If multiple keys are given, it returns the value of the deleted last item.
<原文结束>

# <翻译开始>
// Remove 从缓存中删除一个或多个键，并返回其值。
// 如果给出了多个键，它将返回最后删除项的值。
// md5:b3f23906b769df08
# <翻译结束>


<原文开始>
// Retrieves the last key value.
<原文结束>

# <翻译开始>
// 获取最后一个键值。. md5:c348d395d5ea0c9f
# <翻译结束>


<原文开始>
// Deletes all given keys.
<原文结束>

# <翻译开始>
// 删除所有给定的键。. md5:5c8528683a62a6e5
# <翻译结束>


<原文开始>
// Clear clears all data of the cache.
// Note that this function is sensitive and should be carefully used.
// It uses `FLUSHDB` command in redis server, which might be disabled in server.
<原文结束>

# <翻译开始>
// Clear 清空缓存中的所有数据。
// 注意，此函数具有敏感性，应谨慎使用。
// 它使用了 Redis 服务器中的 `FLUSHDB` 命令，但该命令可能在服务器中被禁用。
// md5:e9b895cf3a7760c0
# <翻译结束>


<原文开始>
// The "FLUSHDB" may not be available.
<原文结束>

# <翻译开始>
// "FLUSHDB"可能不可用。. md5:95fb09eb47c6baab
# <翻译结束>


<原文开始>
// Close closes the cache.
<原文结束>

# <翻译开始>
// Close 关闭缓存。. md5:c1a9d7a347be93a8
# <翻译结束>

