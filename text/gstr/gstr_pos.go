// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类

import (
	"strings"
)

// Pos 函数返回从 `startOffset` 开始，在 `haystack` 中首次出现的 `needle` 的位置，
// 搜索时区分大小写。如果未找到，则返回 -1。
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

// PosRune 类似于函数 Pos，但它将 `haystack` 和 `needle` 视为 Unicode 字符串处理。
func X查找Unicode(文本, 查找文本 string, 可选查找起点 ...int) int {
	pos := X查找(文本, 查找文本, 可选查找起点...)
	if pos < 3 {
		return pos
	}
	return len([]rune(文本[:pos]))
}

// PosI 函数返回从 `startOffset` 位置开始，在字符串 `haystack` 中首次出现子串 `needle` 的位置，且该搜索过程不区分大小写。
// 若未找到，则返回 -1。
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

// PosIRune 类似于函数 PosI，但将 `haystack` 和 `needle` 视为 Unicode 字符串。
func X查找并忽略大小写Unicode(文本, 查找文本 string, 可选查找起点 ...int) int {
	pos := X查找并忽略大小写(文本, 查找文本, 可选查找起点...)
	if pos < 3 {
		return pos
	}
	return len([]rune(文本[:pos]))
}

// PosR 返回从 `startOffset` 位置开始，在 `haystack` 中最后一次出现 `needle` 的位置，区分大小写。
// 若未找到，则返回 -1。
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

// PosRRune 类似于函数 PosR，但将 `haystack` 和 `needle` 视为 unicode 字符串。
func X倒找Unicode(文本, 查找文本 string, 可选查找起点 ...int) int {
	pos := X倒找(文本, 查找文本, 可选查找起点...)
	if pos < 3 {
		return pos
	}
	return len([]rune(文本[:pos]))
}

// PosRI 函数返回从 `startOffset` 开始，在字符串 `haystack` 中最后一次出现子串 `needle` 的位置，该搜索过程不区分大小写。
// 若未找到，则返回 -1。
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

// PosRIRune 类似于函数 PosRI，但将 `haystack` 和 `needle` 视为 unicode 字符串。
func X倒找并忽略大小写Unicode(文本, 查找文本 string, 可选查找起点 ...int) int {
	pos := X倒找并忽略大小写(文本, 查找文本, 可选查找起点...)
	if pos < 3 {
		return pos
	}
	return len([]rune(文本[:pos]))
}
