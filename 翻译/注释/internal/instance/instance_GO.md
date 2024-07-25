
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
// Package instance provides instances management.
//
// Note that this package is not used for cache, as it has no cache expiration.
<原文结束>

# <翻译开始>
// Package instance 提供实例管理功能。
// 
// 注意，此包不用于缓存，因为它没有缓存过期。
// md5:9cde92d483190e72
# <翻译结束>


<原文开始>
// Get returns the instance by given name.
<原文结束>

# <翻译开始>
// Get 根据给定的名称返回实例。 md5:a44f9ed4c07f4bd7
# <翻译结束>


<原文开始>
// Set sets an instance to the instance manager with given name.
<原文结束>

# <翻译开始>
// Set 将给定名称的实例设置到实例管理器中。 md5:b2ea0ff086c307ba
# <翻译结束>


<原文开始>
// GetOrSet returns the instance by name,
// or set instance to the instance manager if it does not exist and returns this instance.
<原文结束>

# <翻译开始>
// GetOrSet 通过名称获取实例，
// 如果不存在，则将其设置到实例管理器中并返回该实例。
// md5:6e30e1788811bdcf
# <翻译结束>


<原文开始>
// GetOrSetFunc returns the instance by name,
// or sets instance with returned value of callback function `f` if it does not exist
// and then returns this instance.
<原文结束>

# <翻译开始>
// GetOrSetFunc 通过名称获取实例，
// 如果不存在，它将使用回调函数 `f` 返回的值设置实例，
// 然后返回这个实例。
// md5:3e2dff7c2a8267b6
# <翻译结束>


<原文开始>
// GetOrSetFuncLock returns the instance by name,
// or sets instance with returned value of callback function `f` if it does not exist
// and then returns this instance.
//
// GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f`
// with mutex.Lock of the hash map.
<原文结束>

# <翻译开始>
// GetOrSetFuncLock 通过名称获取实例，
// 如果该实例不存在，则使用回调函数 `f` 的返回值设置实例，
// 并随后返回这个实例。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，
// 它在执行函数 `f` 时会对哈希映射加锁（mutex.Lock）。
// md5:d7adba14d37045fa
# <翻译结束>


<原文开始>
// SetIfNotExist sets `instance` to the map if the `name` does not exist, then returns true.
// It returns false if `name` exists, and `instance` would be ignored.
<原文结束>

# <翻译开始>
// SetIfNotExist 如果`name`不存在，则将`instance`设置到地图中，然后返回true。
// 如果`name`已经存在，则忽略`instance`并返回false。
// md5:0eb14110f7286ae3
# <翻译结束>


<原文开始>
// Clear deletes all instances stored.
<原文结束>

# <翻译开始>
// Clear 删除所有存储的实例。 md5:19c1efdd76e32ce6
# <翻译结束>

