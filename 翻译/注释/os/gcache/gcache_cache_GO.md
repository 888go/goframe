
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
// localAdapter is alias of Adapter, for embedded attribute purpose only.
<原文结束>

# <翻译开始>
// localAdapter 是 Adapter 的别名，仅用于嵌入式属性的目的。
# <翻译结束>


<原文开始>
// New creates and returns a new cache object using default memory adapter.
// Note that the LRU feature is only available using memory adapter.
<原文结束>

# <翻译开始>
// New 函数创建并返回一个使用默认内存适配器的新缓存对象。
// 注意，LRU（最近最少使用）特性仅在使用内存适配器时可用。
# <翻译结束>


<原文开始>
	// Here may be a "timer leak" if adapter is manually changed from memory adapter.
	// Do not worry about this, as adapter is less changed, and it does nothing if it's not used.
<原文结束>

# <翻译开始>
// 如果适配器手动从内存适配器更改，这里可能存在“计时器泄漏”的问题。
// 不必担心这一点，因为适配器很少更改，并且如果未使用则不会造成任何影响。
# <翻译结束>


<原文开始>
// NewWithAdapter creates and returns a Cache object with given Adapter implements.
<原文结束>

# <翻译开始>
// NewWithAdapter 使用给定的已实现Adapter接口的对象创建并返回一个Cache对象。
# <翻译结束>


<原文开始>
// SetAdapter changes the adapter for this cache.
// Be very note that, this setting function is not concurrent-safe, which means you should not call
// this setting function concurrently in multiple goroutines.
<原文结束>

# <翻译开始>
// SetAdapter 更改此缓存的适配器。
// 非常需要注意的是，这个设置函数不是并发安全的，这意味着你不应该在多个goroutine中并发调用这个设置函数。
# <翻译结束>


<原文开始>
// GetAdapter returns the adapter that is set in current Cache.
<原文结束>

# <翻译开始>
// GetAdapter 返回当前 Cache 中设置的适配器。
# <翻译结束>


<原文开始>
// Removes deletes `keys` in the cache.
<原文结束>

# <翻译开始>
// 删除缓存中的`keys`。
# <翻译结束>


<原文开始>
// KeyStrings returns all keys in the cache as string slice.
<原文结束>

# <翻译开始>
// KeyStrings 返回缓存中的所有键，以字符串切片的形式。
# <翻译结束>

