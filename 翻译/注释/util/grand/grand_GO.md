
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
// Package grand provides high performance random bytes/number/string generation functionality.
<原文结束>

# <翻译开始>
// 包grand提供了高性能的随机字节/数字/字符串生成功能。. md5:b787416316af6730
# <翻译结束>


<原文开始>
// Intn returns an int number which is between 0 and max: [0, max).
//
// Note that:
// 1. The `max` can only be greater than 0, or else it returns `max` directly;
// 2. The result is greater than or equal to 0, but less than `max`;
// 3. The result number is 32bit and less than math.MaxUint32.
<原文结束>

# <翻译开始>
// Intn 返回一个位于0和max之间的整数：[0, max)。
//
// 注意事项：
// 1. `max` 必须大于0，否则将直接返回 `max`；
// 2. 结果大于等于0，但小于 `max`；
// 3. 结果数字是32位的，并且小于 math.MaxUint32。
// md5:a6e33dbab8d41e7e
# <翻译结束>


<原文开始>
// B retrieves and returns random bytes of given length `n`.
<原文结束>

# <翻译开始>
// B 函数获取并返回给定长度 `n` 的随机字节。. md5:bf6fb575b93930c3
# <翻译结束>


<原文开始>
// N returns a random int between min and max: [min, max].
// The `min` and `max` also support negative numbers.
<原文结束>

# <翻译开始>
// N 返回一个在`min`和`max`之间的随机整数：[min, max]。`min`和`max`也支持负数。
// md5:7512aa05c9a381dc
# <翻译结束>


<原文开始>
	// As `Intn` dose not support negative number,
	// so we should first shift the value to right,
	// then call `Intn` to produce the random number,
	// and finally shift the result back to left.
<原文结束>

# <翻译开始>
// 由于`Intn`不支持负数，
// 所以我们首先需要将值向右移动，
// 然后调用`Intn`生成随机数，
// 最后再将结果向左移动回去。
// md5:4ce39d9c2b1e67a4
# <翻译结束>


<原文开始>
// S returns a random string which contains digits and letters, and its length is `n`.
// The optional parameter `symbols` specifies whether the result could contain symbols,
// which is false in default.
<原文结束>

# <翻译开始>
// S 返回一个随机字符串，该字符串包含数字和字母，其长度为 `n`。
// 可选参数 `symbols` 指定结果中是否可以包含符号，默认为 false。
// md5:9b705df69354173e
# <翻译结束>


<原文开始>
// D returns a random time.Duration between min and max: [min, max].
<原文结束>

# <翻译开始>
// D 返回一个在min和max之间的随机时间.Duration：[min, max]。. md5:ef8c3a046e8eb92a
# <翻译结束>


<原文开始>
// Str randomly picks and returns `n` count of chars from given string `s`.
// It also supports unicode string like Chinese/Russian/Japanese, etc.
<原文结束>

# <翻译开始>
// Str 随机从给定字符串 `s` 中选择并返回 `n` 个字符。它也支持Unicode字符串，如中文、俄文、日文等。
// md5:49cb0972916cd378
# <翻译结束>


<原文开始>
// Digits returns a random string which contains only digits, and its length is `n`.
<原文结束>

# <翻译开始>
// Digits 返回一个只包含数字的随机字符串，其长度为 `n`。. md5:b0370aa46ce5a9b0
# <翻译结束>


<原文开始>
// Letters returns a random string which contains only letters, and its length is `n`.
<原文结束>

# <翻译开始>
// Letters 返回一个只包含字母的随机字符串，其长度为 `n`。. md5:42f482efe6a2cdc2
# <翻译结束>


<原文开始>
// Symbols returns a random string which contains only symbols, and its length is `n`.
<原文结束>

# <翻译开始>
// Symbols 返回一个只包含符号的随机字符串，其长度为 `n`。. md5:7478be0329d79c20
# <翻译结束>


<原文开始>
// Perm returns, as a slice of n int numbers, a pseudo-random permutation of the integers [0,n).
// TODO performance improving for large slice producing.
<原文结束>

# <翻译开始>
// Perm返回一个整数切片，它是[0,n)范围内的伪随机排列。该排列作为长度为n的int类型的元素组成。// TODO：优化对于生成大切片时的性能。
// md5:2a47a879b72568c7
# <翻译结束>


<原文开始>
// Meet randomly calculate whether the given probability `num`/`total` is met.
<原文结束>

# <翻译开始>
// 随机计算给定的概率 `num`/`total` 是否满足。. md5:faa82da76f9b9e2a
# <翻译结束>


<原文开始>
// MeetProb randomly calculate whether the given probability is met.
<原文结束>

# <翻译开始>
// MeetProb 随机计算给定的概率是否满足条件。. md5:8443c7abcac61392
# <翻译结束>

