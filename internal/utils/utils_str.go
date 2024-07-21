// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package utils

import (
	"bytes"
	"strings"
)

var (
	// DefaultTrimChars 是Trim*函数默认删除的字符。 md5:4841fe294d977402
	DefaultTrimChars = string([]byte{
		'\t', // Tab.
		'\v', // Vertical tab.
		'\n', // New line (line feed).
		'\r', // Carriage return.
		'\f', // New page.
		' ',  // Ordinary space.
		0x00, // NUL-byte.
		0x85, // Delete.
		0xA0, // Non-breaking space.
	})
)

// IsLetterUpper 检查给定字节 b 是否为大写字母。 md5:92d2fefb65d65bd7
func IsLetterUpper(b byte) bool {
	if b >= byte('A') && b <= byte('Z') {
		return true
	}
	return false
}

// IsLetterLower 检查给定的字节 b 是否为小写。 md5:6f3845d2ae021561
func IsLetterLower(b byte) bool {
	if b >= byte('a') && b <= byte('z') {
		return true
	}
	return false
}

// IsLetter 检查给定的字节b是否为字母。 md5:f483df3b4a456648
func IsLetter(b byte) bool {
	return IsLetterUpper(b) || IsLetterLower(b)
}

// IsNumeric 检查给定的字符串 s 是否为数字。
// 注意，像 "123.456" 这样的浮点数字符串也被视为数字。
// md5:630bec1763c4d982
func IsNumeric(s string) bool {
	var (
		dotCount = 0
		length   = len(s)
	)
	if length == 0 {
		return false
	}
	for i := 0; i < length; i++ {
		if s[i] == '-' && i == 0 {
			continue
		}
		if s[i] == '.' {
			dotCount++
			if i > 0 && i < length-1 {
				continue
			} else {
				return false
			}
		}
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return dotCount <= 1
}

// UcFirst 返回一个字符串s的副本，其中第一个字母映射为其大写形式。 md5:bc090531eef4b3e6
func UcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	if IsLetterLower(s[0]) {
		return string(s[0]-32) + s[1:]
	}
	return s
}

// ReplaceByMap 返回一个`origin`的副本，
// 使用映射无序地替换其中的内容，且区分大小写。
// md5:c047c08d8be640ad
func ReplaceByMap(origin string, replaces map[string]string) string {
	for k, v := range replaces {
		origin = strings.ReplaceAll(origin, k, v)
	}
	return origin
}

// RemoveSymbols 从字符串中移除所有符号，只保留数字和字母。 md5:cdaeb32f53496617
func RemoveSymbols(s string) string {
	var b = make([]rune, 0, len(s))
	for _, c := range s {
		if c > 127 {
			b = append(b, c)
		} else if (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			b = append(b, c)
		}
	}
	return string(b)
}

// EqualFoldWithoutChars 检查字符串 `s1` 和 `s2` 是否在大小写不敏感的情况下相等，包括/不包括字符 '-'、'_'、'.' 和 ' '。
// md5:e519885b86d35bd5
func EqualFoldWithoutChars(s1, s2 string) bool {
	return strings.EqualFold(RemoveSymbols(s1), RemoveSymbols(s2))
}

// SplitAndTrim 通过字符串 `delimiter` 将字符串 `str` 分割成一个数组，然后对数组中的每个元素调用 Trim 函数。它会忽略 Trim 后为空的元素。
// md5:20d7e1d120928c19
func SplitAndTrim(str, delimiter string, characterMask ...string) []string {
	array := make([]string, 0)
	for _, v := range strings.Split(str, delimiter) {
		v = Trim(v, characterMask...)
		if v != "" {
			array = append(array, v)
		}
	}
	return array
}

// Trim 函数从字符串的开头和结尾删除空白符（或其他字符）。
// 可选参数 `characterMask` 指定了额外需要删除的字符。
// md5:a1d794d7a10c9c7a
func Trim(str string, characterMask ...string) string {
	trimChars := DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	return strings.Trim(str, trimChars)
}

// FormatCmdKey 使用统一的格式将字符串 `s` 格式化为命令键。 md5:c946276756bc1c34
func FormatCmdKey(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, "_", "."))
}

// FormatEnvKey 使用统一格式对字符串`s`进行格式化，以便作为环境键。 md5:1ded8906c5ce105d
func FormatEnvKey(s string) string {
	return strings.ToUpper(strings.ReplaceAll(s, ".", "_"))
}

// StripSlashes通过AddSlashes方法去引号化一个带引号的字符串。 md5:fdba1646284db234
func StripSlashes(str string) string {
	var buf bytes.Buffer
	l, skip := len(str), false
	for i, char := range str {
		if skip {
			skip = false
		} else if char == '\\' {
			if i+1 < l && str[i+1] == '\\' {
				skip = true
			}
			continue
		}
		buf.WriteRune(char)
	}
	return buf.String()
}
