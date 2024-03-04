// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstr

import (
	"strings"
	
	"github.com/888go/goframe/gstr/internal/utils"
)

// Trim 函数用于从字符串的开头和结尾去除空格（或其他字符）。
// 可选参数 `characterMask` 指定了需要额外去除的字符。
func Trim(str string, characterMask ...string) string {
	return utils.Trim(str, characterMask...)
}

// TrimStr 从字符串的开头和结尾去除所有给定的`cut`子串。
// 注意，它不会去除其开头或结尾的空格。
func TrimStr(str string, cut string, count ...int) string {
	return TrimLeftStr(TrimRightStr(str, cut, count...), cut, count...)
}

// TrimLeft 从字符串的开头剥离空白（或其他字符）。
func TrimLeft(str string, characterMask ...string) string {
	trimChars := utils.DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	return strings.TrimLeft(str, trimChars)
}

// TrimLeftStr 从字符串的开头移除所有给定的 `cut` 字符串。
// 注意，它不会移除其开头的空白字符。
func TrimLeftStr(str string, cut string, count ...int) string {
	var (
		lenCut   = len(cut)
		cutCount = 0
	)
	for len(str) >= lenCut && str[0:lenCut] == cut {
		str = str[lenCut:]
		cutCount++
		if len(count) > 0 && count[0] != -1 && cutCount >= count[0] {
			break
		}
	}
	return str
}

// TrimRight 从字符串的末尾删除空白字符（或其他字符）。
func TrimRight(str string, characterMask ...string) string {
	trimChars := utils.DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	return strings.TrimRight(str, trimChars)
}

// TrimRightStr 从字符串的末尾剥离所有给定的 `cut` 字符串。
// 注意，它不会剥离其末尾的空白字符。
func TrimRightStr(str string, cut string, count ...int) string {
	var (
		lenStr   = len(str)
		lenCut   = len(cut)
		cutCount = 0
	)
	for lenStr >= lenCut && str[lenStr-lenCut:lenStr] == cut {
		lenStr = lenStr - lenCut
		str = str[:lenStr]
		cutCount++
		if len(count) > 0 && count[0] != -1 && cutCount >= count[0] {
			break
		}
	}
	return str
}

// TrimAll 删除字符串 `str` 中的所有字符。
func TrimAll(str string, characterMask ...string) string {
	trimChars := utils.DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	var (
		filtered bool
		slice    = make([]rune, 0, len(str))
	)
	for _, char := range str {
		filtered = false
		for _, trimChar := range trimChars {
			if char == trimChar {
				filtered = true
				break
			}
		}
		if !filtered {
			slice = append(slice, char)
		}
	}
	return string(slice)
}

// HasPrefix 测试字符串 s 是否以 prefix 开头。
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix 测试字符串 s 是否以 suffix 结尾。
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}
