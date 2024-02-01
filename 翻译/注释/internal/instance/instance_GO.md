
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
// Package instance provides instances management.
//
// Note that this package is not used for cache, as it has no cache expiration.
<原文结束>

# <翻译开始>
// Package instance 提供了实例管理功能。
//
// 注意，此包并不用于缓存，因为它没有缓存过期机制。
# <翻译结束>


<原文开始>
// Get returns the instance by given name.
<原文结束>

# <翻译开始>
// Get通过给定的名称返回实例。
# <翻译结束>


<原文开始>
// Set sets an instance to the instance manager with given name.
<原文结束>

# <翻译开始>
// Set 将具有给定名称的实例设置到实例管理器中。
# <翻译结束>


<原文开始>
// GetOrSet returns the instance by name,
// or set instance to the instance manager if it does not exist and returns this instance.
<原文结束>

# <翻译开始>
// GetOrSet 函数通过名称获取实例，
// 如果实例不存在，则将其设置到实例管理器中并返回该实例。
# <翻译结束>


<原文开始>
// GetOrSetFunc returns the instance by name,
// or sets instance with returned value of callback function `f` if it does not exist
// and then returns this instance.
<原文结束>

# <翻译开始>
// GetOrSetFunc 函数通过名称返回实例，
// 如果实例不存在，则使用回调函数 `f` 返回的值设置该实例，
// 然后返回这个已设置的实例。
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
// GetOrSetFuncLock 通过名称返回实例，
// 如果实例不存在，则使用回调函数 `f` 返回的值设置该实例，
// 然后返回这个实例。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，
// 它在哈希映射的 mutex.Lock 保护下执行函数 `f`。
# <翻译结束>


<原文开始>
// SetIfNotExist sets `instance` to the map if the `name` does not exist, then returns true.
// It returns false if `name` exists, and `instance` would be ignored.
<原文结束>

# <翻译开始>
// SetIfNotExist 如果`name`不存在，则将`instance`设置到map中，并返回true。
// 若`name`已存在，则返回false，同时`instance`将被忽略。
# <翻译结束>


<原文开始>
// Clear deletes all instances stored.
<原文结束>

# <翻译开始>
// Clear 删除所有已存储的实例。
# <翻译结束>

