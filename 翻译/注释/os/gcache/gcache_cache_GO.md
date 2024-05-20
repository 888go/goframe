
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
// localAdapter is alias of Adapter, for embedded attribute purpose only.
<原文结束>

# <翻译开始>
// localAdapter 是 Adapter 的别名，仅用于嵌入属性的目的。. md5:be3fc375883fa166
# <翻译结束>


<原文开始>
// New creates and returns a new cache object using default memory adapter.
// Note that the LRU feature is only available using memory adapter.
<原文结束>

# <翻译开始>
// New 使用默认的内存适配器创建并返回一个新的缓存对象。
// 请注意，LRU（最近最少使用）功能仅在使用内存适配器时可用。
// md5:658995a71d08fbbe
# <翻译结束>


<原文开始>
// NewWithAdapter creates and returns a Cache object with given Adapter implements.
<原文结束>

# <翻译开始>
// NewWithAdapter 使用给定的实现了Adapter接口的适配器创建并返回一个Cache对象。. md5:0c92c6f9af030ccb
# <翻译结束>


<原文开始>
// SetAdapter changes the adapter for this cache.
// Be very note that, this setting function is not concurrent-safe, which means you should not call
// this setting function concurrently in multiple goroutines.
<原文结束>

# <翻译开始>
// SetAdapter 更改此缓存的适配器。
// 非常注意，这个设置函数不是并发安全的，这意味着你不应该在多个goroutine中并发调用此设置函数。
// md5:5f950a554baddc2c
# <翻译结束>


<原文开始>
// GetAdapter returns the adapter that is set in current Cache.
<原文结束>

# <翻译开始>
// GetAdapter 返回当前缓存中设置的适配器。. md5:e93da9e47a8b0c21
# <翻译结束>


<原文开始>
// Removes deletes `keys` in the cache.
<原文结束>

# <翻译开始>
// 从缓存中删除`keys`。. md5:370028bf9f2e1d24
# <翻译结束>


<原文开始>
// KeyStrings returns all keys in the cache as string slice.
<原文结束>

# <翻译开始>
// KeyStrings返回缓存中的所有键作为字符串切片。. md5:3b0126221389825e
# <翻译结束>

