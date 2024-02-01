
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
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly.
<原文结束>

# <翻译开始>
// 创建一个缓存对象，
// 当然，你也可以很方便地直接使用gcache包的方法。
# <翻译结束>


<原文开始>
// Set cache without expiration
<原文结束>

# <翻译开始>
// 设置缓存，不设置过期时间
# <翻译结束>


<原文开始>
// Does the specified key name exist in the cache
<原文结束>

# <翻译开始>
// 指定的键名是否存在于缓存中
# <翻译结束>


<原文开始>
// Delete and return the deleted key value
<原文结束>

# <翻译开始>
// 删除并返回已删除的键值对
# <翻译结束>


<原文开始>
// Close the cache object and let the GC reclaim resources
<原文结束>

# <翻译开始>
// 关闭缓存对象，让垃圾回收器回收资源
# <翻译结束>


<原文开始>
// Write when the key name does not exist, and set the expiration time to 1000 milliseconds
<原文结束>

# <翻译开始>
// 当键名不存在时写入，并设置过期时间为1000毫秒
# <翻译结束>


<原文开始>
// Returns false when the key name already exists
<原文结束>

# <翻译开始>
// 当键名已存在时返回false
# <翻译结束>


<原文开始>
// Print the current list of key values
<原文结束>

# <翻译开始>
// 打印当前键值对列表
# <翻译结束>


<原文开始>
// It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.
<原文结束>

# <翻译开始>
// 如果 `duration` 等于 0，则它不会过期。如果 `duration` 小于 0 或给定的 `value` 为 nil，则它会删除 `key`。
# <翻译结束>


<原文开始>
// Wait 1.5 second for K1: V1 to expire automatically
<原文结束>

# <翻译开始>
// 等待1.5秒，直至K1: V1自动过期
# <翻译结束>


<原文开始>
// Print the current key value pair again and find that K1: V1 has expired
<原文结束>

# <翻译开始>
// 再次打印当前键值对，会发现 K1: V1 已经过期
# <翻译结束>







<原文开始>
	// Sets batch sets cache with key-value pairs by `data`, which is expired after `duration`.
	// It does not expire if `duration` == 0. It deletes the keys of `data` if `duration` < 0 or given `value` is nil.
<原文结束>

# <翻译开始>
// SetsBatch 函数通过 `data` 设置缓存中的键值对，该键值对在 `duration` 后过期。
// 如果 `duration` 等于 0，则不会过期。如果 `duration` 小于 0 或提供的 `value` 为 nil，则会删除 `data` 中的键。
# <翻译结束>


<原文开始>
// Gets the specified key value
<原文结束>

# <翻译开始>
// 获取指定键的值
# <翻译结束>


<原文开始>
// Add 10 elements without expiration
<原文结束>

# <翻译开始>
// 添加10个无过期时间的元素
# <翻译结束>


<原文开始>
// Size returns the number of items in the cache.
<原文结束>

# <翻译开始>
// Size 返回缓存中的项目数量。
# <翻译结束>


<原文开始>
// Update updates the value of `key` without changing its expiration and returns the old value.
<原文结束>

# <翻译开始>
// Update 更新键 `key` 的值，但不会改变其过期时间，并返回旧的值。
# <翻译结束>


<原文开始>
	// The returned value `exist` is false if the `key` does not exist in the cache.
	// It does nothing if `key` does not exist in the cache.
<原文结束>

# <翻译开始>
// 若`key`在缓存中不存在，返回的值`exist`为false。
// 若`key`在缓存中不存在，则此操作不做任何处理。
# <翻译结束>


<原文开始>
	// UpdateExpire updates the expiration of `key` and returns the old expiration duration value.
	// It returns -1 and does nothing if the `key` does not exist in the cache.
<原文结束>

# <翻译开始>
// UpdateExpire 更新键 `key` 的过期时间，并返回旧的过期持续时间值。
// 如果 `key` 不存在于缓存中，则返回 -1 并不做任何操作。
# <翻译结束>


<原文开始>
	// c.Set(ctx, "k2", "Here is Value2", 0)
	// c.Set(ctx, "k3", 111, 0)
<原文结束>

# <翻译开始>
// c.Set(ctx, "k2", "Here is Value2", 0)
// 在给定的上下文ctx中，将键为"k2"的值设置为"Here is Value2"，并设置过期时间为0（表示永不过期）
// c.Set(ctx, "k3", 111, 0)
// 在给定的上下文ctx中，将键为"k3"的值设置为整数111，并设置过期时间为0（表示永不过期）
// 在上述代码中，`c` 应该是一个具有缓存功能的对象，`Set` 方法用于设置缓存项，参数包括操作的上下文、键名和对应的值以及可选的过期时间。这里设置的过期时间是0，通常意味着缓存项永不过期。
# <翻译结束>


<原文开始>
// Values returns all values in the cache as slice.
<原文结束>

# <翻译开始>
// Values 返回缓存中的所有值作为一个切片。
# <翻译结束>


<原文开始>
// Close closes the cache if necessary.
<原文结束>

# <翻译开始>
// Close 在必要时关闭缓存。
# <翻译结束>


<原文开始>
	// Contains returns true if `key` exists in the cache, or else returns false.
	// return true
<原文结束>

# <翻译开始>
// Contains 返回 true 如果 `key` 存在于缓存中，否则返回 false。
# <翻译结束>


<原文开始>
	// Get retrieves and returns the associated value of given `key`.
	// It returns nil if it does not exist, its value is nil or it's expired.
<原文结束>

# <翻译开始>
// Get 方法用于获取并返回给定`key`关联的值。
// 如果该键不存在，其值为 nil 或已过期，则返回 nil。
# <翻译结束>


<原文开始>
	// GetExpire retrieves and returns the expiration of `key` in the cache.
	// It returns 0 if the `key` does not expire. It returns -1 if the `key` does not exist in the cache.
<原文结束>

# <翻译开始>
// GetExpire 从缓存中获取并返回`key`的过期时间。
// 如果`key`永不过期，则返回0。如果`key`在缓存中不存在，则返回-1。
# <翻译结束>


<原文开始>
	// GetOrSet retrieves and returns the value of `key`, or sets `key`-`value` pair and returns `value`
	// if `key` does not exist in the cache.
<原文结束>

# <翻译开始>
// GetOrSet 获取并返回 `key` 对应的值，如果 `key` 不存在于缓存中，则设置 `key`-`value` 键值对并返回 `value`。
# <翻译结束>


<原文开始>
	// GetOrSetFunc retrieves and returns the value of `key`, or sets `key` with result of function `f`
	// and returns its result if `key` does not exist in the cache.
<原文结束>

# <翻译开始>
// GetOrSetFunc 函数用于获取并返回 `key` 对应的值，如果 `key` 不存在于缓存中，则使用函数 `f` 的结果设置 `key` 并返回其执行结果。
# <翻译结束>


<原文开始>
// If func returns nil, no action is taken
<原文结束>

# <翻译开始>
// 如果函数返回nil，则不执行任何操作
# <翻译结束>


<原文开始>
// Modify locking Note that the function `f` should be executed within writing mutex lock for concurrent safety purpose.
<原文结束>

# <翻译开始>
// 修改锁定：请注意，为了保证并发安全，函数`f`应当在写入互斥锁保护下执行。
# <翻译结束>







<原文开始>
// KeyStrings returns all keys in the cache as string slice.
<原文结束>

# <翻译开始>
// KeyStrings 返回缓存中的所有键，以字符串切片形式。
# <翻译结束>


<原文开始>
	// Remove deletes one or more keys from cache, and returns its value.
	// If multiple keys are given, it returns the value of the last deleted item.
<原文结束>

# <翻译开始>
// Remove 从缓存中删除一个或多个键，并返回其对应的值。
// 如果提供了多个键，它将返回最后被删除项的值。
# <翻译结束>


<原文开始>
	// Intercepting panic exception information
	// err is empty, so panic is not performed
<原文结束>

# <翻译开始>
// 拦截 panic 异常信息
// err 为空，因此不执行 panic
# <翻译结束>


<原文开始>
// MustGet acts like Get, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGet 行为类似于 Get，但是当出现任何错误时，它会触发panic（异常）。
# <翻译结束>


<原文开始>
// MustGetOrSet acts like GetOrSet, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGetOrSet 行为类似于 GetOrSet，但当发生任何错误时，它会触发panic（异常）。
# <翻译结束>


<原文开始>
// MustGetOrSetFunc acts like GetOrSetFunc, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGetOrSetFunc 行为类似于 GetOrSetFunc，但当出现任何错误时，它会触发panic（异常）。
# <翻译结束>


<原文开始>
// MustGetOrSetFuncLock acts like GetOrSetFuncLock, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGetOrSetFuncLock 行为类似于 GetOrSetFuncLock，但是当发生任何错误时，它会引发panic（恐慌）。
# <翻译结束>


<原文开始>
	// MustContains returns true if `key` exists in the cache, or else returns false.
	// return true
<原文结束>

# <翻译开始>
// MustContains 返回一个布尔值，如果 `key` 存在于缓存中则返回 true，否则返回 false。
# <翻译结束>


<原文开始>
// MustGetExpire acts like GetExpire, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGetExpire 行为类似于 GetExpire，但是当发生任何错误时它会触发panic（异常）。
# <翻译结束>


<原文开始>
// MustKeys acts like Keys, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustKeys 行为类似于 Keys，但如果出现任何错误，它会引发 panic。
# <翻译结束>


<原文开始>
	// MustKeyStrings returns all keys in the cache as string slice.
	// MustKeyStrings acts like KeyStrings, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustKeyStrings 返回缓存中的所有键作为字符串切片。
// MustKeyStrings 类似于 KeyStrings，但在出现任何错误时会触发 panic。
# <翻译结束>


<原文开始>
// MustValues returns all values in the cache as slice.
<原文结束>

# <翻译开始>
// MustValues 返回缓存中所有值的切片。
# <翻译结束>


<原文开始>
// Create redis client object.
<原文结束>

# <翻译开始>
// 创建Redis客户端对象。
# <翻译结束>


<原文开始>
// Create redis cache adapter and set it to cache object.
<原文结束>

# <翻译开始>
// 创建Redis缓存适配器并将它设置为缓存对象。
# <翻译结束>


<原文开始>
// Set and Get using cache object.
<原文结束>

# <翻译开始>
// 使用缓存对象进行设置和获取操作。
# <翻译结束>






 
<原文开始>
// map[interface{}]interface{}
<原文结束>

# <翻译开始>
// map[interface{}]interface{}：这是一个Go语言中的映射（map）类型，它的键和值都是接口类型（interface{}）。这意味着这个映射可以存储任意类型的键值对，因为interface{}可以表示任何类型。在实际使用中，这种类型的映射通常用于需要处理多种不同类型数据的场景，但需要注意，由于go的静态类型特性，在取值时需要进行类型断言转换。
# <翻译结束>


<原文开始>
// Modification failed
<原文结束>

# <翻译开始>
// 修改失败
# <翻译结束>


<原文开始>
// Get using redis client.
<原文结束>

# <翻译开始>
// 使用redis客户端获取。
# <翻译结束>

