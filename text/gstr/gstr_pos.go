// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gstr

import "strings"

// Pos returns the position of the first occurrence of `needle`
// in `haystack` from `startOffset`, case-sensitively.
// It returns -1, if not found.

// ff:查找
// startOffset:可选查找起点
// needle:查找文本
// haystack:文本
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

// PosRune acts like function Pos but considers `haystack` and `needle` as unicode string.

// ff:查找Unicode
// startOffset:可选查找起点
// needle:查找文本
// haystack:文本
func PosRune(haystack, needle string, startOffset ...int) int {
	pos := Pos(haystack, needle, startOffset...)
	if pos < 3 {
		return pos
	}
	return len([]rune(haystack[:pos]))
}

// PosI returns the position of the first occurrence of `needle`
// in `haystack` from `startOffset`, case-insensitively.
// It returns -1, if not found.

// ff:查找并忽略大小写
// startOffset:可选查找起点
// needle:查找文本
// haystack:文本
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

// PosIRune acts like function PosI but considers `haystack` and `needle` as unicode string.

// ff:查找并忽略大小写Unicode
// startOffset:可选查找起点
// needle:查找文本
// haystack:文本
func PosIRune(haystack, needle string, startOffset ...int) int {
	pos := PosI(haystack, needle, startOffset...)
	if pos < 3 {
		return pos
	}
	return len([]rune(haystack[:pos]))
}

// PosR returns the position of the last occurrence of `needle`
// in `haystack` from `startOffset`, case-sensitively.
// It returns -1, if not found.

// ff:倒找
// startOffset:可选查找起点
// needle:查找文本
// haystack:文本
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

// PosRRune acts like function PosR but considers `haystack` and `needle` as unicode string.

// ff:倒找Unicode
// startOffset:可选查找起点
// needle:查找文本
// haystack:文本
func PosRRune(haystack, needle string, startOffset ...int) int {
	pos := PosR(haystack, needle, startOffset...)
	if pos < 3 {
		return pos
	}
	return len([]rune(haystack[:pos]))
}

// PosRI returns the position of the last occurrence of `needle`
// in `haystack` from `startOffset`, case-insensitively.
// It returns -1, if not found.

// ff:倒找并忽略大小写
// startOffset:可选查找起点
// needle:查找文本
// haystack:文本
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

// PosRIRune acts like function PosRI but considers `haystack` and `needle` as unicode string.

// ff:倒找并忽略大小写Unicode
// startOffset:可选查找起点
// needle:查找文本
// haystack:文本
func PosRIRune(haystack, needle string, startOffset ...int) int {
	pos := PosRI(haystack, needle, startOffset...)
	if pos < 3 {
		return pos
	}
	return len([]rune(haystack[:pos]))
}
