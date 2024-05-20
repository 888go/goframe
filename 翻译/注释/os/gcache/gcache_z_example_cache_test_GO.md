
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
	// Create a cache object,
	// Of course, you can also easily use the gcache package method directly.
<原文结束>

# <翻译开始>
// 创建一个缓存对象，
// 当然，你也可以直接轻松地使用gcache包提供的方法。
// md5:ffacc88538a18ac1
# <翻译结束>


<原文开始>
// Set cache without expiration
<原文结束>

# <翻译开始>
// 设置不带过期时间的缓存. md5:10e925a877b589df
# <翻译结束>


<原文开始>
// Does the specified key name exist in the cache
<原文结束>

# <翻译开始>
// 指定的键名是否存在于缓存中. md5:9f18fb652f082e8d
# <翻译结束>


<原文开始>
// Delete and return the deleted key value
<原文结束>

# <翻译开始>
// 删除并返回被删除的键值对. md5:01f96ef0f0eae0e3
# <翻译结束>


<原文开始>
// Close the cache object and let the GC reclaim resources
<原文结束>

# <翻译开始>
// 关闭缓存对象，允许GC回收资源. md5:e6035a9d9a9583ce
# <翻译结束>


<原文开始>
// Write when the key name does not exist, and set the expiration time to 1000 milliseconds
<原文结束>

# <翻译开始>
// 如果键名不存在，则写入，并设置过期时间为1000毫秒. md5:41f1844c720c4e5f
# <翻译结束>


<原文开始>
// Returns false when the key name already exists
<原文结束>

# <翻译开始>
// 当键名已存在时返回false. md5:db29c3756eb62f45
# <翻译结束>


<原文开始>
// Print the current list of key values
<原文结束>

# <翻译开始>
// 打印当前的键值对列表. md5:9b92acd8c3138f30
# <翻译结束>


<原文开始>
// It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.
<原文结束>

# <翻译开始>
// 如果`duration`等于0，它不会过期。如果`duration`小于0或给定的`value`为nil，它将删除`key`。. md5:a0794bc140ecff80
# <翻译结束>


<原文开始>
// Wait 1.5 second for K1: V1 to expire automatically
<原文结束>

# <翻译开始>
// 等待1.5秒，让K1: V1自动过期. md5:58d8a37819ac1a31
# <翻译结束>


<原文开始>
// Print the current key value pair again and find that K1: V1 has expired
<原文结束>

# <翻译开始>
// 再次打印当前的键值对，发现K1: V1 已过期. md5:50e60452bc46b7a6
# <翻译结束>


<原文开始>
// map[interface{}]interface{}
<原文结束>

# <翻译开始>
// 一个映射，其键和值都是接口类型. md5:5e21ae79c908b4df
# <翻译结束>


<原文开始>
	// Sets batch sets cache with key-value pairs by `data`, which is expired after `duration`.
	// It does not expire if `duration` == 0. It deletes the keys of `data` if `duration` < 0 or given `value` is nil.
<原文结束>

# <翻译开始>
// 使用键值对`data`设置批处理缓存，过期时间为`duration`。
// 如果`duration`为0，则不会过期。如果`duration`小于0或给定的`value`为nil，将删除`data`中的键。
// md5:b2f121999e39c24d
# <翻译结束>


<原文开始>
// Gets the specified key value
<原文结束>

# <翻译开始>
// 获取指定的键值. md5:78fc35a6610c5179
# <翻译结束>


<原文开始>
// Add 10 elements without expiration
<原文结束>

# <翻译开始>
// 添加10个不带过期的元素. md5:cbefc995139f6ed9
# <翻译结束>


<原文开始>
// Size returns the number of items in the cache.
<原文结束>

# <翻译开始>
// Size 返回缓存中的项目数量。. md5:2122f80de9340261
# <翻译结束>


<原文开始>
// Update updates the value of `key` without changing its expiration and returns the old value.
<原文结束>

# <翻译开始>
// Update 更新 `key` 对应的值，而不改变其过期时间，并返回旧值。. md5:1e7dc1ae84b2f449
# <翻译结束>


<原文开始>
	// The returned value `exist` is false if the `key` does not exist in the cache.
	// It does nothing if `key` does not exist in the cache.
<原文结束>

# <翻译开始>
// 返回的值`exist`如果`key`在缓存中不存在，则为false。
// 如果`key`不在缓存中，它将不执行任何操作。
// md5:1ecdac9a4397de58
# <翻译结束>


<原文开始>
	// UpdateExpire updates the expiration of `key` and returns the old expiration duration value.
	// It returns -1 and does nothing if the `key` does not exist in the cache.
<原文结束>

# <翻译开始>
// UpdateExpire 更新键`key`的过期时间，并返回旧的过期持续时间值。如果缓存中不存在`key`，则返回-1并什么都不做。
// md5:8a59f61ead71c844
# <翻译结束>


<原文开始>
	// May Output:
	// 1s
	// 500ms
<原文结束>

# <翻译开始>
	// May Output:
	// 1s
	// 500ms
# <翻译结束>


<原文开始>
	// c.Set(ctx, "k2", "Here is Value2", 0)
	// c.Set(ctx, "k3", 111, 0)
<原文结束>

# <翻译开始>
// 将键为 "k2" 的值设置为 "Here is Value2"，过期时间为 0（默认不设置过期时间）
// 将键为 "k3" 的值设置为整数 111，过期时间为 0（默认不设置过期时间）
// md5:b8ee9984ed9da43d
# <翻译结束>


<原文开始>
// Values returns all values in the cache as slice.
<原文结束>

# <翻译开始>
// Values 返回缓存中所有的值作为切片。. md5:dc00b32eb8913e9b
# <翻译结束>


<原文开始>
	// May Output:
	// [map[k1:v1 k2:v2]]
<原文结束>

# <翻译开始>
	// May Output:
	// [map[k1:v1 k2:v2]]
# <翻译结束>


<原文开始>
// Close closes the cache if necessary.
<原文结束>

# <翻译开始>
// Close如果有必要，关闭缓存。. md5:f9a73a30e4b4b396
# <翻译结束>


<原文开始>
	// Contains returns true if `key` exists in the cache, or else returns false.
	// return true
<原文结束>

# <翻译开始>
// Contains 如果`key`存在于缓存中，则返回true，否则返回false。
// md5:370eab2c5835bfb1
# <翻译结束>


<原文开始>
	// Get retrieves and returns the associated value of given `key`.
	// It returns nil if it does not exist, its value is nil or it's expired.
<原文结束>

# <翻译开始>
// Get 获取并返回给定`key`关联的值。
// 如果键不存在、其值为nil或已过期，它将返回nil。
// md5:2999106994454771
# <翻译结束>


<原文开始>
	// GetExpire retrieves and returns the expiration of `key` in the cache.
	// It returns 0 if the `key` does not expire. It returns -1 if the `key` does not exist in the cache.
<原文结束>

# <翻译开始>
// GetExpire 从缓存中检索并返回`key`的过期时间。如果`key`不过期，它将返回0。如果`key`在缓存中不存在，它将返回-1。
// md5:a60a46e9632013e1
# <翻译结束>


<原文开始>
	// GetOrSet retrieves and returns the value of `key`, or sets `key`-`value` pair and returns `value`
	// if `key` does not exist in the cache.
<原文结束>

# <翻译开始>
// GetOrSet 从缓存中检索并返回`key`对应的值，如果`key`在缓存中不存在，则设置`key-value`对，并返回`value`。
// md5:f1f24272b9b4a43c
# <翻译结束>


<原文开始>
	// GetOrSetFunc retrieves and returns the value of `key`, or sets `key` with result of function `f`
	// and returns its result if `key` does not exist in the cache.
<原文结束>

# <翻译开始>
// GetOrSetFunc 方法尝试获取并返回`key`对应的值，如果`key`在缓存中不存在，则使用函数`f`的结果设置`key`的值，
// 并返回该函数的结果。
// md5:6a21b8ed72969e95
# <翻译结束>


<原文开始>
// If func returns nil, no action is taken
<原文结束>

# <翻译开始>
// 如果函数返回nil，不执行任何操作. md5:a7c277c9df4048d6
# <翻译结束>


<原文开始>
// Modify locking Note that the function `f` should be executed within writing mutex lock for concurrent safety purpose.
<原文结束>

# <翻译开始>
// 修改锁定注意，为了并发安全，函数 `f` 应该在写入锁的保护下执行。. md5:a86de4ea66d58271
# <翻译结束>


<原文开始>
// KeyStrings returns all keys in the cache as string slice.
<原文结束>

# <翻译开始>
// KeyStrings返回缓存中的所有键作为字符串切片。. md5:3b0126221389825e
# <翻译结束>


<原文开始>
	// May Output:
	// [k1 k2]
<原文结束>

# <翻译开始>
	// May Output:
	// [k1 k2]
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
// clears all data of the cache.
<原文结束>

# <翻译开始>
// 清空缓存的所有数据。. md5:13010db2c416938b
# <翻译结束>


<原文开始>
	// Intercepting panic exception information
	// err is empty, so panic is not performed
<原文结束>

# <翻译开始>
// 拦截恐慌异常信息
// err 为空，因此不执行恐慌操作
// md5:aa899aa9abc889f7
# <翻译结束>


<原文开始>
// MustGet acts like Get, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGet 的行为就像 Get 一样，但如果发生任何错误，它会引发 panic。. md5:9004545d221e9637
# <翻译结束>


<原文开始>
// MustGetOrSet acts like GetOrSet, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGetOrSet 的行为类似于 GetOrSet，但是如果发生任何错误，它会直接 panic。. md5:684c6b06451a2f6f
# <翻译结束>


<原文开始>
// MustGetOrSetFunc acts like GetOrSetFunc, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGetOrSetFunc 行为类似于 GetOrSetFunc，但如果发生任何错误，则会引发 panic。. md5:07fd1ef2dbfce0b4
# <翻译结束>


<原文开始>
// MustGetOrSetFuncLock acts like GetOrSetFuncLock, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGetOrSetFuncLock 行为与 GetOrSetFuncLock 类似，但如果发生任何错误，它将引发恐慌。. md5:7f84f54a71da5305
# <翻译结束>


<原文开始>
	// MustContains returns true if `key` exists in the cache, or else returns false.
	// return true
<原文结束>

# <翻译开始>
// MustContains 返回true如果`key`在缓存中存在，否则返回false。
// 返回true
// md5:226a8dda1fb50b87
# <翻译结束>


<原文开始>
// MustGetExpire acts like GetExpire, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGetExpire 的行为类似于 GetExpire，但如果发生任何错误，它会直接 panic。. md5:c97fa5941bbc47a3
# <翻译结束>


<原文开始>
	// May Output:
	// map[k1:v1 k2:v2]
<原文结束>

# <翻译开始>
	// May Output:
	// map[k1:v1 k2:v2]
# <翻译结束>


<原文开始>
// MustKeys acts like Keys, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustKeys 行为与 Keys 类似，但如果发生任何错误，它将引发 panic。. md5:7f7801d0cd170166
# <翻译结束>


<原文开始>
	// MustKeyStrings returns all keys in the cache as string slice.
	// MustKeyStrings acts like KeyStrings, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustKeyStrings 返回缓存中的所有键作为字符串切片。
// MustKeyStrings 的行为类似于 KeyStrings，但如果发生任何错误，它将引发恐慌。
// md5:e647a507f2385601
# <翻译结束>


<原文开始>
// MustValues returns all values in the cache as slice.
<原文结束>

# <翻译开始>
// MustValues 返回缓存中的所有值作为切片。. md5:8b4269e238366b51
# <翻译结束>


<原文开始>
// Create redis client object.
<原文结束>

# <翻译开始>
// 创建Redis客户端对象。. md5:f412dd032940c79e
# <翻译结束>


<原文开始>
// Create redis cache adapter and set it to cache object.
<原文结束>

# <翻译开始>
// 创建Redis缓存适配器，并将其设置到缓存对象中。. md5:fe080f47e7881b0a
# <翻译结束>


<原文开始>
// Set and Get using cache object.
<原文结束>

# <翻译开始>
// 使用缓存对象进行设置和获取。. md5:27779d48bc4565ef
# <翻译结束>


<原文开始>
// Get using redis client.
<原文结束>

# <翻译开始>
// 使用redis客户端获取。. md5:413fbbedeb694205
# <翻译结束>


<原文开始>
	// May Output:
	// value
	// value
<原文结束>

# <翻译开始>
	// May Output:
	// value
	// value
# <翻译结束>


<原文开始>
	// May Output:
	// value
	// <nil>
	// value
<原文结束>

# <翻译开始>
	// May Output:
	// value
	// <nil>
	// value
# <翻译结束>

