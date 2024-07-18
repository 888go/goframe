// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包grand提供了高性能的随机字节/数字/字符串生成功能。 md5:b787416316af6730
package grand//bm:随机类

import (
	"encoding/binary"
	"time"
)

var (
	letters    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 52
	symbols    = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"                   // 32
	digits     = "0123456789"                                           // 10
	characters = letters + digits + symbols                             // 94
)

// Intn 返回一个位于0和max之间的整数：[0, max)。
//
// 注意事项：
// 1. `max` 必须大于0，否则将直接返回 `max`；
// 2. 结果大于等于0，但小于 `max`；
// 3. 结果数字是32位的，并且小于 math.MaxUint32。
// md5:a6e33dbab8d41e7e
// ff:整数
// max:最大值
func Intn(max int) int {
	if max <= 0 {
		return max
	}
	n := int(binary.LittleEndian.Uint32(<-bufferChan)) % max
	if (max > 0 && n < 0) || (max < 0 && n > 0) {
		return -n
	}
	return n
}

// B 函数获取并返回给定长度 `n` 的随机字节。 md5:bf6fb575b93930c3
// ff:字节集
// n:长度
func B(n int) []byte {
	if n <= 0 {
		return nil
	}
	i := 0
	b := make([]byte, n)
	for {
		copy(b[i:], <-bufferChan)
		i += 4
		if i >= n {
			break
		}
	}
	return b
}

// N 返回一个在`min`和`max`之间的随机整数：[min, max]。`min`和`max`也支持负数。
// md5:7512aa05c9a381dc
// ff:区间整数
// min:最小值
// max:最大值
func N(min, max int) int {
	if min >= max {
		return min
	}
	if min >= 0 {
		return Intn(max-min+1) + min
	}
// 由于`Intn`不支持负数，
// 所以我们首先需要将值向右移动，
// 然后调用`Intn`生成随机数，
// 最后再将结果向左移动回去。
// md5:4ce39d9c2b1e67a4
	return Intn(max+(0-min)+1) - (0 - min)
}

// S 返回一个随机字符串，该字符串包含数字和字母，其长度为 `n`。
// 可选参数 `symbols` 指定结果中是否可以包含符号，默认为 false。
// md5:9b705df69354173e
// ff:文本
// n:长度
// symbols:包含特殊字符
func S(n int, symbols ...bool) string {
	if n <= 0 {
		return ""
	}
	var (
		b           = make([]byte, n)
		numberBytes = B(n)
	)
	for i := range b {
		if len(symbols) > 0 && symbols[0] {
			b[i] = characters[numberBytes[i]%94]
		} else {
			b[i] = characters[numberBytes[i]%62]
		}
	}
	return string(b)
}

// D 返回一个在min和max之间的随机时间.Duration：[min, max]。 md5:ef8c3a046e8eb92a
// ff:时长
// min:最小值
// max:最大值
func D(min, max time.Duration) time.Duration {
	multiple := int64(1)
	if min != 0 {
		for min%10 == 0 {
			multiple *= 10
			min /= 10
			max /= 10
		}
	}
	n := int64(N(int(min), int(max)))
	return time.Duration(n * multiple)
}

// Str 随机从给定字符串 `s` 中选择并返回 `n` 个字符。它也支持Unicode字符串，如中文、俄文、日文等。
// md5:49cb0972916cd378
// ff:从文本生成文本
// s:给定文本
// n:长度
func Str(s string, n int) string {
	if n <= 0 {
		return ""
	}
	var (
		b     = make([]rune, n)
		runes = []rune(s)
	)
	if len(runes) <= 255 {
		numberBytes := B(n)
		for i := range b {
			b[i] = runes[int(numberBytes[i])%len(runes)]
		}
	} else {
		for i := range b {
			b[i] = runes[Intn(len(runes))]
		}
	}
	return string(b)
}

// Digits 返回一个只包含数字的随机字符串，其长度为 `n`。 md5:b0370aa46ce5a9b0
// ff:数字文本
// n:长度
func Digits(n int) string {
	if n <= 0 {
		return ""
	}
	var (
		b           = make([]byte, n)
		numberBytes = B(n)
	)
	for i := range b {
		b[i] = digits[numberBytes[i]%10]
	}
	return string(b)
}

// Letters 返回一个只包含字母的随机字符串，其长度为 `n`。 md5:42f482efe6a2cdc2
// ff:字母文本
// n:长度
func Letters(n int) string {
	if n <= 0 {
		return ""
	}
	var (
		b           = make([]byte, n)
		numberBytes = B(n)
	)
	for i := range b {
		b[i] = letters[numberBytes[i]%52]
	}
	return string(b)
}

// Symbols 返回一个只包含符号的随机字符串，其长度为 `n`。 md5:7478be0329d79c20
// ff:特殊字符文本
// n:长度
func Symbols(n int) string {
	if n <= 0 {
		return ""
	}
	var (
		b           = make([]byte, n)
		numberBytes = B(n)
	)
	for i := range b {
		b[i] = symbols[numberBytes[i]%32]
	}
	return string(b)
}

// Perm返回一个整数切片，它是[0,n)范围内的伪随机排列。该排列作为长度为n的int类型的元素组成。// TODO：优化对于生成大切片时的性能。
// md5:2a47a879b72568c7
// ff:整数切片
// n:长度
func Perm(n int) []int {
	m := make([]int, n)
	for i := 0; i < n; i++ {
		j := Intn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}

// 随机计算给定的概率 `num`/`total` 是否满足。 md5:faa82da76f9b9e2a
// ff:
// num:
// total:
func Meet(num, total int) bool {
	return Intn(total) < num
}

// MeetProb 随机计算给定的概率是否满足条件。 md5:8443c7abcac61392
// ff:
// prob:
func MeetProb(prob float32) bool {
	return Intn(1e7) < int(prob*1e7)
}
