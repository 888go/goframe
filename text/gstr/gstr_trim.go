// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gstr

import (
	"strings"

	"github.com/gogf/gf/v2/internal/utils"
)

// Trim 函数从字符串的开头和结尾删除空白符（或其他字符）。
// 可选参数 `characterMask` 指定了额外需要删除的字符。
// md5:a1d794d7a10c9c7a
func Trim(str string, characterMask ...string) string {
	return utils.Trim(str, characterMask...)
}

// TrimStr 从字符串的开始和结束处删除给定的 `cut` 字符串。
// 请注意，它不会删除其开始或结束处的空格。
// md5:6180dbb1bb9965c4
func TrimStr(str string, cut string, count ...int) string {
	return TrimLeftStr(TrimRightStr(str, cut, count...), cut, count...)
}

// TrimLeft 从字符串的开头移除空格（或其他字符）。 md5:648a4eb609597439
func TrimLeft(str string, characterMask ...string) string {
	trimChars := utils.DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	return strings.TrimLeft(str, trimChars)
}

// TrimLeftStr 从字符串的开头移除所有给定的 `cut` 字符串。注意，它不会移除字符串开头的空格。
// md5:765cffbaed7e3cce
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

// TrimRight 从字符串的末尾移除空白符（或其他字符）。 md5:c252aae10c16cd5a
func TrimRight(str string, characterMask ...string) string {
	trimChars := utils.DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	return strings.TrimRight(str, trimChars)
}

// TrimRightStr 从字符串的末尾删除所有给定的 `cut` 字符串。
// 请注意，它不会删除字符串末尾的空格。
// md5:0ca4b90c9d600b39
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

// TrimAll 函数会删除字符串 `str` 中的所有字符。 md5:2d770fefafe6bda4
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

// HasPrefix 测试字符串 s 是否以 prefix 开头。 md5:8b4cd90bff10b28a
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix 检测字符串 s 是否以 suffix 结尾。 md5:964d208ab1e843e1
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}
