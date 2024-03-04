
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
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Package gpool provides object-reusable concurrent-safe pool.
<原文结束>

# <翻译开始>
// 包gpool提供了一个对象可重用的并发安全池。
# <翻译结束>


<原文开始>
// Pool is an Object-Reusable Pool.
<原文结束>

# <翻译开始>
// Pool 是一个对象可重用池。
# <翻译结束>


<原文开始>
// Time To Live for pool items.
<原文结束>

# <翻译开始>
// Time To Live for pool items. // 池中项目（对象）的生存时间。
# <翻译结束>


<原文开始>
// Callback function to create pool item.
<原文结束>

# <翻译开始>
// 回调函数，用于创建池中的项目。
# <翻译结束>


<原文开始>
	// ExpireFunc is the for expired items destruction.
	// This function needs to be defined when the pool items
	// need to perform additional destruction operations.
	// Eg: net.Conn, os.File, etc.
<原文结束>

# <翻译开始>
// ExpireFunc 是用于过期项销毁的函数。
// 当池中的项需要执行额外的销毁操作时，需要定义此函数。
// 例如：net.Conn（网络连接），os.File（操作系统文件）等。
# <翻译结束>


<原文开始>
// Expire timestamp in milliseconds.
<原文结束>

# <翻译开始>
// 过期时间戳（毫秒）
# <翻译结束>


<原文开始>
// NewFunc Creation function for object.
<原文结束>

# <翻译开始>
// NewFunc 创建对象的构造函数
# <翻译结束>


<原文开始>
// ExpireFunc Destruction function for object.
<原文结束>

# <翻译开始>
// ExpireFunc 对象的销毁函数。
# <翻译结束>


<原文开始>
// New creates and returns a new object pool.
// To ensure execution efficiency, the expiration time cannot be modified once it is set.
//
// Note the expiration logic:
// ttl = 0 : not expired;
// ttl < 0 : immediate expired after use;
// ttl > 0 : timeout expired;
<原文结束>

# <翻译开始>
// New 创建并返回一个新的对象池。
// 为了确保执行效率，一旦设置过期时间则不可再修改。
//
// 注意过期逻辑：
// ttl = 0 : 不过期；
// ttl < 0 : 使用后立即过期；
// ttl > 0 : 超时后过期；
# <翻译结束>


<原文开始>
		// As for Golang version < 1.13, there's no method Milliseconds for time.Duration.
		// So we need calculate the milliseconds using its nanoseconds value.
<原文结束>

# <翻译开始>
// 对于 Go 语言版本小于 1.13 的情况，time.Duration 类型没有内置的 Milliseconds 方法。
// 因此我们需要通过其纳秒值来计算毫秒值。
# <翻译结束>


<原文开始>
// MustPut puts an item to pool, it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustPut 将一个项目放入池中，如果发生任何错误，它会引发panic。
# <翻译结束>


<原文开始>
// Clear clears pool, which means it will remove all items from pool.
<原文结束>

# <翻译开始>
// Clear 清空 pool，这意味着它会从 pool 中移除所有项目。
# <翻译结束>


<原文开始>
// Get picks and returns an item from pool. If the pool is empty and NewFunc is defined,
// it creates and returns one from NewFunc.
<原文结束>

# <翻译开始>
// Get 从池中获取并返回一个项目。如果池为空且定义了 NewFunc，
// 则会通过 NewFunc 创建并返回一个新的项目。
# <翻译结束>


<原文开始>
// TODO: move expire function calling asynchronously out from `Get` operation.
<原文结束>

# <翻译开始>
// TODO: 将过期函数调用异步移出 `Get` 操作。
# <翻译结束>


<原文开始>
// Size returns the count of available items of pool.
<原文结束>

# <翻译开始>
// Size 返回 pool 中可用项目的数量。
# <翻译结束>


<原文开始>
// Close closes the pool. If `p` has ExpireFunc,
// then it automatically closes all items using this function before it's closed.
// Commonly you do not need to call this function manually.
<原文结束>

# <翻译开始>
// Close 关闭连接池。如果 `p` 拥有 ExpireFunc，
// 则在关闭前会自动使用该函数关闭所有项目。
// 通常情况下，你不需要手动调用这个函数。
# <翻译结束>


<原文开始>
// checkExpire removes expired items from pool in every second.
<原文结束>

# <翻译开始>
// checkExpire 每隔一秒从池中移除已过期的项目。
# <翻译结束>


<原文开始>
		// If p has ExpireFunc,
		// then it must close all items using this function.
<原文结束>

# <翻译开始>
// 如果p拥有ExpireFunc，
// 那么它必须使用这个函数关闭所有项。
# <翻译结束>


<原文开始>
// The latest item expire timestamp in milliseconds.
<原文结束>

# <翻译开始>
// 最近一项过期时间戳（以毫秒为单位）
# <翻译结束>


<原文开始>
	// Retrieve the current timestamp in milliseconds, it expires the items
	// by comparing with this timestamp. It is not accurate comparison for
	// every item expired, but high performance.
<原文结束>

# <翻译开始>
// 获取当前时间戳（毫秒级），通过与此时间戳进行比较来决定缓存项是否过期。这种方法并非对每一个缓存项的过期判断都精确，但具有高性能的特点。
# <翻译结束>


<原文开始>
// TODO improve the auto-expiration mechanism of the pool.
<原文结束>

# <翻译开始>
// TODO: 改进池的自动过期机制。
# <翻译结束>


<原文开始>
// Available/idle items list.
<原文结束>

# <翻译开始>
// 可用/空闲项目列表。
# <翻译结束>


<原文开始>
// Whether the pool is closed.
<原文结束>

# <翻译开始>
// 是否池已关闭。
# <翻译结束>


<原文开始>
// Put puts an item to pool.
<原文结束>

# <翻译开始>
// Put 将一个项目放入池中。
# <翻译结束>


<原文开始>
// All items do not expire.
<原文结束>

# <翻译开始>
// 所有项目永不过期。
# <翻译结束>

