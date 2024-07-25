
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
	// Cache for regex object.
	// Note that:
	// 1. It uses sync.RWMutex ensuring the concurrent safety.
	// 2. There's no expiring logic for this map.
<原文结束>

# <翻译开始>
	// 正则表达式对象的缓存。
	// 注意：
	// 1. 使用 sync.RWMutex 确保并发安全性。
	// 2. 这个映射表中没有过期逻辑。 md5:645e245ad93c001d
# <翻译结束>


<原文开始>
// getRegexp returns *regexp.Regexp object with given `pattern`.
// It uses cache to enhance the performance for compiling regular expression pattern,
// which means, it will return the same *regexp.Regexp object with the same regular
// expression pattern.
//
// It is concurrent-safe for multiple goroutines.
<原文结束>

# <翻译开始>
// getRegexp 使用给定的 `pattern` 返回一个 *regexp.Regexp 对象。
// 它使用缓存来提升正则表达式模式编译的性能，
// 即，对于相同的正则表达式模式，它会返回同一个 *regexp.Regexp 对象。
//
// 它是多线程安全的，适用于多个goroutine。 md5:7c16df93c3eeb2b1
# <翻译结束>


<原文开始>
// Retrieve the regular expression object using reading lock.
<原文结束>

# <翻译开始>
	// 使用读取锁获取正则表达式对象。 md5:8d1b5f1036b66cce
# <翻译结束>


<原文开始>
	// If it does not exist in the cache,
	// it compiles the pattern and creates one.
<原文结束>

# <翻译开始>
	// 如果该模式不存在于缓存中，
	// 则编译该模式并创建一个。 md5:16abd6a4a92df88a
# <翻译结束>


<原文开始>
// Cache the result object using writing lock.
<原文结束>

# <翻译开始>
	// 使用写入锁缓存结果对象。 md5:4d4db9dbdc7391d7
# <翻译结束>

