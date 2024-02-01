
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
// Package grand provides high performance random bytes/number/string generation functionality.
<原文结束>

# <翻译开始>
// 包grand提供了高性能的随机字节/数字/字符串生成功能。
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
// Intn 返回一个在0和max之间的随机整数: [0, max)。
//
// 注意事项：
// 1. 参数`max`必须大于0，否则函数直接返回`max`；
// 2. 返回结果大于等于0且小于`max`；
// 3. 返回的随机数是32位整数，并且小于math.MaxUint32。
# <翻译结束>


<原文开始>
// B retrieves and returns random bytes of given length `n`.
<原文结束>

# <翻译开始>
// B 函数获取并返回指定长度 `n` 的随机字节。
# <翻译结束>


<原文开始>
// N returns a random int between min and max: [min, max].
// The `min` and `max` also support negative numbers.
<原文结束>

# <翻译开始>
// N 返回一个在 min 和 max 范围内的随机整数：[min, max]。
// 其中 `min` 和 `max` 也支持负数。
# <翻译结束>


<原文开始>
	// As `Intn` dose not support negative number,
	// so we should first shift the value to right,
	// then call `Intn` to produce the random number,
	// and finally shift the result back to left.
<原文结束>

# <翻译开始>
// 由于`Intn`函数不支持负数，
// 因此我们首先需要将数值右移，
// 然后调用`Intn`生成随机数，
// 最后将结果左移回来。
# <翻译结束>


<原文开始>
// S returns a random string which contains digits and letters, and its length is `n`.
// The optional parameter `symbols` specifies whether the result could contain symbols,
// which is false in default.
<原文结束>

# <翻译开始>
// S 返回一个随机字符串，该字符串包含数字和字母，其长度为 `n`。
// 可选参数 `symbols` 指定生成的字符串是否可以包含符号，默认情况下为 false（不包含）。
# <翻译结束>


<原文开始>
// D returns a random time.Duration between min and max: [min, max].
<原文结束>

# <翻译开始>
// D 返回一个在min和max之间随机的time.Duration值：[min, max]。
# <翻译结束>


<原文开始>
// Str randomly picks and returns `n` count of chars from given string `s`.
// It also supports unicode string like Chinese/Russian/Japanese, etc.
<原文结束>

# <翻译开始>
// Str 从给定的字符串 `s` 中随机选取并返回 `n` 个字符。
// 它还支持Unicode字符串，如中文、俄文、日文等。
# <翻译结束>


<原文开始>
// Digits returns a random string which contains only digits, and its length is `n`.
<原文结束>

# <翻译开始>
// Digits 返回一个随机字符串，其中只包含数字，并且其长度为 `n`。
# <翻译结束>


<原文开始>
// Letters returns a random string which contains only letters, and its length is `n`.
<原文结束>

# <翻译开始>
// Letters 返回一个随机字符串，该字符串只包含字母，并且长度为 `n`。
# <翻译结束>


<原文开始>
// Symbols returns a random string which contains only symbols, and its length is `n`.
<原文结束>

# <翻译开始>
// Symbols 返回一个随机字符串，其中只包含符号，并且其长度为 `n`。
# <翻译结束>


<原文开始>
// Perm returns, as a slice of n int numbers, a pseudo-random permutation of the integers [0,n).
// TODO performance improving for large slice producing.
<原文结束>

# <翻译开始>
// Perm 函数返回一个整数切片，其中包含了从 [0, n) 范围内随机排列的 n 个整数。
// TODO：优化处理大规模切片生成时的性能。
# <翻译结束>


<原文开始>
// Meet randomly calculate whether the given probability `num`/`total` is met.
<原文结束>

# <翻译开始>
// Meet 随机计算给定的概率 `num`/`total` 是否满足条件。
# <翻译结束>


<原文开始>
// MeetProb randomly calculate whether the given probability is met.
<原文结束>

# <翻译开始>
// MeetProb 随机计算给定概率是否满足。
# <翻译结束>

