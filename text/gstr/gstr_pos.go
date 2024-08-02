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

// Pos 函数返回 `needle` 在 `haystack` 中从 `startOffset` 开始的第一次出现的位置，
// 区分大小写。如果未找到，则返回 -1。
// md5:5f4ec91dd8b1819b
func Pos(haystack, needle string, startOffset ...int) int {
	length := len(haystack)
	offset := 0
	if len(startOffset) > 0 {
		offset = startOffset[0]
	}
	if length == 0 || offset > length || -offset > length {
		return -1
	}
	if offset < 0 {
		offset += length
	}
	pos := strings.Index(haystack[offset:], needle)
	if pos == NotFoundIndex {
		return NotFoundIndex
	}
	return pos + offset
}

// PosRune 行为类似于函数 Pos，但将 `haystack` 和 `needle` 视为Unicode字符串。 md5:8b42bd50a9b68293
func PosRune(haystack, needle string, startOffset ...int) int {
	pos := Pos(haystack, needle, startOffset...)
	if pos < 3 {
		return pos
	}
	return len([]rune(haystack[:pos]))
}

// PosI 函数从 `startOffset` 开始在 `haystack` 中按大小写不敏感的方式返回 `needle` 的第一个出现位置。
// 如果未找到，它将返回 -1。
// md5:9643b4a9d5243ed9
func PosI(haystack, needle string, startOffset ...int) int {
	length := len(haystack)
	offset := 0
	if len(startOffset) > 0 {
		offset = startOffset[0]
	}
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(strings.ToLower(haystack[offset:]), strings.ToLower(needle))
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// PosIRune函数的行为类似于PosI函数，但它将`haystack`和`needle`视为Unicode字符串。 md5:23449f37c0061789
func PosIRune(haystack, needle string, startOffset ...int) int {
	pos := PosI(haystack, needle, startOffset...)
	if pos < 3 {
		return pos
	}
	return len([]rune(haystack[:pos]))
}

// PosR 返回 `needle` 在 `haystack` 中从 `startOffset` 开始的最后一次出现的位置，
// 区分大小写。如果未找到，则返回 -1。
// md5:ceee68bb3f71cbed
func PosR(haystack, needle string, startOffset ...int) int {
	offset := 0
	if len(startOffset) > 0 {
		offset = startOffset[0]
	}
	pos, length := 0, len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		haystack = haystack[:offset+length+1]
	} else {
		haystack = haystack[offset:]
	}
	pos = strings.LastIndex(haystack, needle)
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

// PosRRune 行为类似于函数 PosR，但它将`haystack`和`needle`视为Unicode字符串。 md5:ef6eb08c3f0bdfdf
func PosRRune(haystack, needle string, startOffset ...int) int {
	pos := PosR(haystack, needle, startOffset...)
	if pos < 3 {
		return pos
	}
	return len([]rune(haystack[:pos]))
}

// PosRI 在从 `startOffset` 开始的 `haystack` 中，不区分大小写地返回 `needle` 的最后一次出现的位置。
// 如果未找到，它将返回 -1。
// md5:5ce310967675e981
func PosRI(haystack, needle string, startOffset ...int) int {
	offset := 0
	if len(startOffset) > 0 {
		offset = startOffset[0]
	}
	pos, length := 0, len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		haystack = haystack[:offset+length+1]
	} else {
		haystack = haystack[offset:]
	}
	pos = strings.LastIndex(strings.ToLower(haystack), strings.ToLower(needle))
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

// PosRIRune 行为类似于 PosRI 函数，但它将 `haystack` 和 `needle` 视为 Unicode 字符串。 md5:5dc327c84c5271a6
func PosRIRune(haystack, needle string, startOffset ...int) int {
	pos := PosRI(haystack, needle, startOffset...)
	if pos < 3 {
		return pos
	}
	return len([]rune(haystack[:pos]))
}
