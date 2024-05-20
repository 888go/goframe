
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
// Package gpool provides object-reusable concurrent-safe pool.
<原文结束>

# <翻译开始>
// 包gpool提供了对象可重用的并发安全池。. md5:d111530cd572ede7
# <翻译结束>


<原文开始>
// Pool is an Object-Reusable Pool.
<原文结束>

# <翻译开始>
// Pool是一个对象可重用池。. md5:08c256ba80594199
# <翻译结束>


<原文开始>
// Available/idle items list.
<原文结束>

# <翻译开始>
// 可用/闲置项目列表。. md5:f93c8d617cafe97f
# <翻译结束>


<原文开始>
// Whether the pool is closed.
<原文结束>

# <翻译开始>
// 是否关闭了连接池。. md5:73ea5526318af92f
# <翻译结束>


<原文开始>
// Time To Live for pool items.
<原文结束>

# <翻译开始>
// 对象池中项目的生存时间。. md5:d9c944077d869281
# <翻译结束>


<原文开始>
// Callback function to create pool item.
<原文结束>

# <翻译开始>
// 创建池项的回调函数。. md5:f37bfc92a2188739
# <翻译结束>


<原文开始>
	// ExpireFunc is the function for expired items destruction.
	// This function needs to be defined when the pool items
	// need to perform additional destruction operations.
	// Eg: net.Conn, os.File, etc.
<原文结束>

# <翻译开始>
// ExpireFunc 是用于过期项目销毁的函数。
// 当池中的项目需要执行额外销毁操作时，需要定义这个函数。
// 例如：net.Conn、os.File 等。
// md5:f09911de2780aeaa
# <翻译结束>


<原文开始>
// Expire timestamp in milliseconds.
<原文结束>

# <翻译开始>
// 过期时间戳，单位为毫秒。. md5:d7096ed51593fa59
# <翻译结束>


<原文开始>
// NewFunc Creation function for object.
<原文结束>

# <翻译开始>
// NewFunc 对象的创建函数。. md5:245f622ac151f3ff
# <翻译结束>


<原文开始>
// ExpireFunc Destruction function for object.
<原文结束>

# <翻译开始>
// ExpireFunc 对象的销毁函数。. md5:cd0e3912eae30a98
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
// 为了确保执行效率，一旦设置，过期时间将不能修改。
// 
// 注意过期逻辑：
// ttl = 0：未过期；
// ttl < 0：使用后立即过期；
// ttl > 0：超时过期。
// md5:9f724382dd2313e7
# <翻译结束>


<原文开始>
// Put puts an item to pool.
<原文结束>

# <翻译开始>
// Put 将一个项目放入池中。. md5:d7b57780f7e8f1cc
# <翻译结束>


<原文开始>
		// As for Golang version < 1.13, there's no method Milliseconds for time.Duration.
		// So we need calculate the milliseconds using its nanoseconds value.
<原文结束>

# <翻译开始>
// 对于Golang版本小于1.13的，time.Duration没有Milliseconds方法。
// 因此我们需要使用其纳秒值来计算毫秒数。
// md5:87b516a9573fac98
# <翻译结束>


<原文开始>
// MustPut puts an item to pool, it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustPut 将一个项目放入池中，如果发生任何错误，它将引发恐慌。. md5:10206f4587a99039
# <翻译结束>


<原文开始>
// Clear clears pool, which means it will remove all items from pool.
<原文结束>

# <翻译开始>
// Clear 清空池子，这意味着它将从池中移除所有项目。. md5:c141b6e6c215bc68
# <翻译结束>


<原文开始>
// Get picks and returns an item from pool. If the pool is empty and NewFunc is defined,
// it creates and returns one from NewFunc.
<原文结束>

# <翻译开始>
// Get 从池中选取并返回一个项目。如果池是空的并且已定义了NewFunc，
// 则会使用NewFunc创建并返回一个项目。
// md5:7782b49d380b807b
# <翻译结束>


<原文开始>
// TODO: move expire function calling asynchronously out from `Get` operation.
<原文结束>

# <翻译开始>
// 待办事项：将过期功能调用异步移出`Get`操作。. md5:13d59efb4d92da03
# <翻译结束>


<原文开始>
// Size returns the count of available items of pool.
<原文结束>

# <翻译开始>
// Size 返回池中可用项目的数量。. md5:2b8a683e177e1586
# <翻译结束>


<原文开始>
// Close closes the pool. If `p` has ExpireFunc,
// then it automatically closes all items using this function before it's closed.
// Commonly you do not need to call this function manually.
<原文结束>

# <翻译开始>
// Close 关闭池。如果 `p` 有 ExpireFunc，那么在关闭之前，它会自动使用这个函数关闭所有项目。通常情况下，你不需要手动调用这个函数。
// md5:368c18d44115f9cc
# <翻译结束>


<原文开始>
// checkExpire removes expired items from pool in every second.
<原文结束>

# <翻译开始>
// checkExpire 每秒从池中移除过期的项目。. md5:1177ab8b3e8a341e
# <翻译结束>


<原文开始>
		// If p has ExpireFunc,
		// then it must close all items using this function.
<原文结束>

# <翻译开始>
// 如果p具有ExpireFunc，
// 则必须使用此函数关闭所有项。
// md5:8ec38193671c9632
# <翻译结束>


<原文开始>
// All items do not expire.
<原文结束>

# <翻译开始>
// 所有项目都不会过期。. md5:9c3b9311c20c9c20
# <翻译结束>


<原文开始>
// The latest item expire timestamp in milliseconds.
<原文结束>

# <翻译开始>
// 最新的项目过期时间（以毫秒为单位）。. md5:46946a9b5c1228ca
# <翻译结束>


<原文开始>
	// Retrieve the current timestamp in milliseconds, it expires the items
	// by comparing with this timestamp. It is not accurate comparison for
	// every item expired, but high performance.
<原文结束>

# <翻译开始>
// 获取当前时间戳（毫秒），使用这个时间戳来判断项目是否过期。
// 对于每个项目过期的判断不是非常精确，但性能较高。
// md5:5dc686eec927131e
# <翻译结束>


<原文开始>
// TODO improve the auto-expiration mechanism of the pool.
<原文结束>

# <翻译开始>
// TODO 优化池的自动过期机制。. md5:b4e2c483478d7737
# <翻译结束>

