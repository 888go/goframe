// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文本类

import (
	"strings"
)

// X查找 函数返回 `needle` 在 `haystack` 中从 `startOffset` 开始的第一次出现的位置，
// 区分大小写。如果未找到，则返回 -1。
// md5:5f4ec91dd8b1819b
func X查找(文本, 查找文本 string, 可选查找起点 ...int) int {
	length := len(文本)
	offset := 0
	if len(可选查找起点) > 0 {
		offset = 可选查找起点[0]
	}
	if length == 0 || offset > length || -offset > length {
		return -1
	}
	if offset < 0 {
		offset += length
	}
	pos := strings.Index(文本[offset:], 查找文本)
	if pos == NotFoundIndex {
		return NotFoundIndex
	}
	return pos + offset
}

// X查找Unicode 行为类似于函数 Pos，但将 `haystack` 和 `needle` 视为Unicode字符串。 md5:8b42bd50a9b68293
func X查找Unicode(文本, 查找文本 string, 可选查找起点 ...int) int {
	pos := X查找(文本, 查找文本, 可选查找起点...)
	if pos < 3 {
		return pos
	}
	return len([]rune(文本[:pos]))
}

// X查找并忽略大小写 函数从 `startOffset` 开始在 `haystack` 中按大小写不敏感的方式返回 `needle` 的第一个出现位置。
// 如果未找到，它将返回 -1。
// md5:9643b4a9d5243ed9
func X查找并忽略大小写(文本, 查找文本 string, 可选查找起点 ...int) int {
	length := len(文本)
	offset := 0
	if len(可选查找起点) > 0 {
		offset = 可选查找起点[0]
	}
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(strings.ToLower(文本[offset:]), strings.ToLower(查找文本))
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// X查找并忽略大小写Unicode函数的行为类似于PosI函数，但它将`haystack`和`needle`视为Unicode字符串。 md5:23449f37c0061789
func X查找并忽略大小写Unicode(文本, 查找文本 string, 可选查找起点 ...int) int {
	pos := X查找并忽略大小写(文本, 查找文本, 可选查找起点...)
	if pos < 3 {
		return pos
	}
	return len([]rune(文本[:pos]))
}

// X倒找 返回 `needle` 在 `haystack` 中从 `startOffset` 开始的最后一次出现的位置，
// 区分大小写。如果未找到，则返回 -1。
// md5:ceee68bb3f71cbed
func X倒找(文本, 查找文本 string, 可选查找起点 ...int) int {
	offset := 0
	if len(可选查找起点) > 0 {
		offset = 可选查找起点[0]
	}
	pos, length := 0, len(文本)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		文本 = 文本[:offset+length+1]
	} else {
		文本 = 文本[offset:]
	}
	pos = strings.LastIndex(文本, 查找文本)
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

// X倒找Unicode 行为类似于函数 PosR，但它将`haystack`和`needle`视为Unicode字符串。 md5:ef6eb08c3f0bdfdf
func X倒找Unicode(文本, 查找文本 string, 可选查找起点 ...int) int {
	pos := X倒找(文本, 查找文本, 可选查找起点...)
	if pos < 3 {
		return pos
	}
	return len([]rune(文本[:pos]))
}

// X倒找并忽略大小写 在从 `startOffset` 开始的 `haystack` 中，不区分大小写地返回 `needle` 的最后一次出现的位置。
// 如果未找到，它将返回 -1。
// md5:5ce310967675e981
func X倒找并忽略大小写(文本, 查找文本 string, 可选查找起点 ...int) int {
	offset := 0
	if len(可选查找起点) > 0 {
		offset = 可选查找起点[0]
	}
	pos, length := 0, len(文本)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		文本 = 文本[:offset+length+1]
	} else {
		文本 = 文本[offset:]
	}
	pos = strings.LastIndex(strings.ToLower(文本), strings.ToLower(查找文本))
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

// X倒找并忽略大小写Unicode 行为类似于 PosRI 函数，但它将 `haystack` 和 `needle` 视为 Unicode 字符串。 md5:5dc327c84c5271a6
func X倒找并忽略大小写Unicode(文本, 查找文本 string, 可选查找起点 ...int) int {
	pos := X倒找并忽略大小写(文本, 查找文本, 可选查找起点...)
	if pos < 3 {
		return pos
	}
	return len([]rune(文本[:pos]))
}
