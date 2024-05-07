// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包grand提供了高性能的随机字节/数字/字符串生成功能。
package 随机类

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

// Intn 返回一个在0和max之间的随机整数: [0, max)。
//
// 注意事项：
// 1. 参数`max`必须大于0，否则函数直接返回`max`；
// 2. 返回结果大于等于0且小于`max`；
// 3. 返回的随机数是32位整数，并且小于math.MaxUint32。
func X整数(最大值 int) int {
	if 最大值 <= 0 {
		return 最大值
	}
	n := int(binary.LittleEndian.Uint32(<-bufferChan)) % 最大值
	if (最大值 > 0 && n < 0) || (最大值 < 0 && n > 0) {
		return -n
	}
	return n
}

// B 函数获取并返回指定长度 `n` 的随机字节。
func X字节集(长度 int) []byte {
	if 长度 <= 0 {
		return nil
	}
	i := 0
	b := make([]byte, 长度)
	for {
		copy(b[i:], <-bufferChan)
		i += 4
		if i >= 长度 {
			break
		}
	}
	return b
}

// N 返回一个在 min 和 max 范围内的随机整数：[min, max]。
// 其中 `min` 和 `max` 也支持负数。
func X区间整数(最小值, 最大值 int) int {
	if 最小值 >= 最大值 {
		return 最小值
	}
	if 最小值 >= 0 {
		return X整数(最大值-最小值+1) + 最小值
	}
	// 由于`Intn`函数不支持负数，
	// 因此我们首先需要将数值右移，
	// 然后调用`Intn`生成随机数，
	// 最后将结果左移回来。
	return X整数(最大值+(0-最小值)+1) - (0 - 最小值)
}

// S 返回一个随机字符串，该字符串包含数字和字母，其长度为 `n`。
// 可选参数 `symbols` 指定生成的字符串是否可以包含符号，默认情况下为 false（不包含）。
func X文本(长度 int, 包含特殊字符 ...bool) string {
	if 长度 <= 0 {
		return ""
	}
	var (
		b           = make([]byte, 长度)
		numberBytes = X字节集(长度)
	)
	for i := range b {
		if len(包含特殊字符) > 0 && 包含特殊字符[0] {
			b[i] = characters[numberBytes[i]%94]
		} else {
			b[i] = characters[numberBytes[i]%62]
		}
	}
	return string(b)
}

// D 返回一个在min和max之间随机的time.Duration值：[min, max]。
func X时长(最小值, 最大值 time.Duration) time.Duration {
	multiple := int64(1)
	if 最小值 != 0 {
		for 最小值%10 == 0 {
			multiple *= 10
			最小值 /= 10
			最大值 /= 10
		}
	}
	n := int64(X区间整数(int(最小值), int(最大值)))
	return time.Duration(n * multiple)
}

// Str 从给定的字符串 `s` 中随机选取并返回 `n` 个字符。
// 它还支持Unicode字符串，如中文、俄文、日文等。
func X从文本生成文本(给定文本 string, 长度 int) string {
	if 长度 <= 0 {
		return ""
	}
	var (
		b     = make([]rune, 长度)
		runes = []rune(给定文本)
	)
	if len(runes) <= 255 {
		numberBytes := X字节集(长度)
		for i := range b {
			b[i] = runes[int(numberBytes[i])%len(runes)]
		}
	} else {
		for i := range b {
			b[i] = runes[X整数(len(runes))]
		}
	}
	return string(b)
}

// Digits 返回一个随机字符串，其中只包含数字，并且其长度为 `n`。
func X数字文本(长度 int) string {
	if 长度 <= 0 {
		return ""
	}
	var (
		b           = make([]byte, 长度)
		numberBytes = X字节集(长度)
	)
	for i := range b {
		b[i] = digits[numberBytes[i]%10]
	}
	return string(b)
}

// Letters 返回一个随机字符串，该字符串只包含字母，并且长度为 `n`。
func X字母文本(长度 int) string {
	if 长度 <= 0 {
		return ""
	}
	var (
		b           = make([]byte, 长度)
		numberBytes = X字节集(长度)
	)
	for i := range b {
		b[i] = letters[numberBytes[i]%52]
	}
	return string(b)
}

// Symbols 返回一个随机字符串，其中只包含符号，并且其长度为 `n`。
func X特殊字符文本(长度 int) string {
	if 长度 <= 0 {
		return ""
	}
	var (
		b           = make([]byte, 长度)
		numberBytes = X字节集(长度)
	)
	for i := range b {
		b[i] = symbols[numberBytes[i]%32]
	}
	return string(b)
}

// Perm 函数返回一个整数切片，其中包含了从 [0, n) 范围内随机排列的 n 个整数。
// TODO：优化处理大规模切片生成时的性能。
func X整数切片(长度 int) []int {
	m := make([]int, 长度)
	for i := 0; i < 长度; i++ {
		j := X整数(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}

// Meet 随机计算给定的概率 `num`/`total` 是否满足条件。
func Meet(num, total int) bool {
	return X整数(total) < num
}

// MeetProb 随机计算给定概率是否满足。
func MeetProb(prob float32) bool {
	return X整数(1e7) < int(prob*1e7)
}
