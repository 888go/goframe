// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package utils
import (
	"bytes"
	"strings"
	)
var (
	// DefaultTrimChars 是默认被 Trim* 函数去除的字符集合。
	DefaultTrimChars = string([]byte{
		'\t', // Tab.
		'\v', // Vertical tab.
		'\n', // 新行（换行符）。
		'\r', // Carriage return.
		'\f', // New page.
		' ',  // Ordinary space.
		0x00, // NUL-byte.
		0x85, // Delete.
		0xA0, // 非换行空格。
	})
)

// IsLetterUpper 检查给定的字节 b 是否为大写字母。
func IsLetterUpper(b byte) bool {
	if b >= byte('A') && b <= byte('Z') {
		return true
	}
	return false
}

// IsLetterLower 检查给定的字节 b 是否为小写字母。
func IsLetterLower(b byte) bool {
	if b >= byte('a') && b <= byte('z') {
		return true
	}
	return false
}

// IsLetter 检查给定的字节 b 是否为字母。
func IsLetter(b byte) bool {
	return IsLetterUpper(b) || IsLetterLower(b)
}

// IsNumeric 检查给定的字符串 s 是否为数值类型。
// 注意，小数形式的字符串如 "123.456" 也被视为数值类型。
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

// UcFirst 返回一个字符串 s 的副本，其中首字母已转换为大写。
func UcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	if IsLetterLower(s[0]) {
		return string(s[0]-32) + s[1:]
	}
	return s
}

// ReplaceByMap 函数返回 `origin` 的一个副本，
// 并使用一个无序的映射进行替换，且替换操作区分大小写。
func ReplaceByMap(origin string, replaces map[string]string) string {
	for k, v := range replaces {
		origin = strings.ReplaceAll(origin, k, v)
	}
	return origin
}

// RemoveSymbols 从字符串中移除所有符号，仅保留数字和字母。
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

// EqualFoldWithoutChars 检查字符串 `s1` 和 `s2` 是否不区分大小写相等，
// 同时在有/无 '-'/'_'/'.'/' ' 这些字符的情况下进行比较。
func EqualFoldWithoutChars(s1, s2 string) bool {
	return strings.EqualFold(RemoveSymbols(s1), RemoveSymbols(s2))
}

// SplitAndTrim通过字符串`delimiter`将字符串`str`分割成一个数组，
// 然后对数组中的每个元素调用Trim方法。它会忽略经过Trim处理后为空的元素。
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

// Trim 函数用于从字符串的开头和结尾去除空格（或其他字符）。
// 可选参数 `characterMask` 指定了需要额外去除的字符。
func Trim(str string, characterMask ...string) string {
	trimChars := DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	return strings.Trim(str, trimChars)
}

// FormatCmdKey 将字符串`s`格式化为统一格式的命令键。
func FormatCmdKey(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, "_", "."))
}

// FormatEnvKey 将字符串`s`格式化为统一格式的环境变量键。
func FormatEnvKey(s string) string {
	return strings.ToUpper(strings.ReplaceAll(s, ".", "_"))
}

// StripSlashes 将通过 AddSlashes 方法添加了反斜杠的字符串进行去除引用操作。
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
