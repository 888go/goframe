
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
// Package gcache provides kinds of cache management for process.
//
// It provides a concurrent-safe in-memory cache adapter for process in default.
<原文结束>

# <翻译开始>
// Package gcache 提供了多种用于进程的缓存管理功能。
//
// 默认情况下，它提供了一个线程安全的内存缓存适配器，用于进程中的缓存管理。
# <翻译结束>


<原文开始>
// Func is the cache function that calculates and returns the value.
<原文结束>

# <翻译开始>
// Func 是缓存函数，用于计算并返回值。
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
// Removes deletes `keys` in the cache.
<原文结束>

# <翻译开始>
// 删除缓存中的`keys`。
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
// Size returns the number of items in the cache.
<原文结束>

# <翻译开始>
// Size 返回缓存中的项目数量。
# <翻译结束>


<原文开始>
// Data returns a copy of all key-value pairs in the cache as map type.
// Note that this function may lead lots of memory usage, you can implement this function
// if necessary.
<原文结束>

# <翻译开始>
// Data 返回缓存中所有键值对的副本，类型为 map。注意，此函数可能会导致大量内存使用，
// 因此请按需实现该函数。
# <翻译结束>


<原文开始>
// Keys returns all keys in the cache as slice.
<原文结束>

# <翻译开始>
// Keys 返回缓存中的所有键作为切片。
# <翻译结束>


<原文开始>
// KeyStrings returns all keys in the cache as string slice.
<原文结束>

# <翻译开始>
// KeyStrings 返回缓存中的所有键，以字符串切片的形式。
# <翻译结束>


<原文开始>
// Values returns all values in the cache as slice.
<原文结束>

# <翻译开始>
// Values 返回缓存中的所有值作为一个切片。
# <翻译结束>


<原文开始>
// MustGet acts like Get, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGet 行为类似于 Get，但当发生任何错误时，它会触发panic。
# <翻译结束>


<原文开始>
// MustGetOrSet acts like GetOrSet, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGetOrSet 行为类似于 GetOrSet，但是当发生任何错误时，它会触发panic（异常）。
# <翻译结束>


<原文开始>
// MustGetOrSetFunc acts like GetOrSetFunc, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGetOrSetFunc 行为类似于 GetOrSetFunc，但当发生任何错误时它会触发panic。
# <翻译结束>


<原文开始>
// MustGetOrSetFuncLock acts like GetOrSetFuncLock, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGetOrSetFuncLock 类似于 GetOrSetFuncLock，但如果发生任何错误，它会触发 panic。
# <翻译结束>


<原文开始>
// MustContains acts like Contains, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustContains 行为类似于 Contains，但当发生任何错误时，它会触发panic（异常退出）。
# <翻译结束>


<原文开始>
// MustGetExpire acts like GetExpire, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGetExpire 的行为类似于 GetExpire，但是当发生任何错误时，它会触发panic。
# <翻译结束>


<原文开始>
// MustSize acts like Size, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustSize 的行为类似于 Size，但如果发生任何错误，它会触发 panic。
# <翻译结束>


<原文开始>
// MustData acts like Data, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustData 的行为类似于 Data，但是当发生任何错误时它会触发panic（异常）。
# <翻译结束>


<原文开始>
// MustKeys acts like Keys, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustKeys 行为类似 Keys，但当发生任何错误时会触发 panic。
# <翻译结束>


<原文开始>
// MustKeyStrings acts like KeyStrings, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustKeyStrings 行为类似 KeyStrings，但当发生任何错误时，它会引发 panic。
# <翻译结束>


<原文开始>
// MustValues acts like Values, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustValues 行为类似 Values，但是当发生任何错误时它会触发panic（异常）。
# <翻译结束>


<原文开始>
// Expire duration that never expires.
<原文结束>

# <翻译开始>
// Expire 永不过期的持续时间。
# <翻译结束>


<原文开始>
// Default cache object.
<原文结束>

# <翻译开始>
// 默认缓存对象。
# <翻译结束>

