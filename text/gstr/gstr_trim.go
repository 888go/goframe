// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类

import (
	"strings"
	
	"github.com/888go/goframe/internal/utils"
)

// Trim 函数用于从字符串的开头和结尾去除空格（或其他字符）。
// 可选参数 `characterMask` 指定了需要额外去除的字符。
func X过滤首尾符并含空白(文本 string, 可选过滤符号 ...string) string {
	return utils.Trim(文本, 可选过滤符号...)
}

// TrimStr 从字符串的开头和结尾去除所有给定的`cut`子串。
// 注意，它不会去除其开头或结尾的空格。
func X过滤首尾(文本 string, 过滤符号 string, 可选过滤次数 ...int) string {
	return X过滤首字符(X过滤尾字符(文本, 过滤符号, 可选过滤次数...), 过滤符号, 可选过滤次数...)
}

// TrimLeft 从字符串的开头剥离空白（或其他字符）。
func X过滤首字符并含空白(文本 string, 可选过滤符号 ...string) string {
	trimChars := utils.DefaultTrimChars
	if len(可选过滤符号) > 0 {
		trimChars += 可选过滤符号[0]
	}
	return strings.TrimLeft(文本, trimChars)
}

// TrimLeftStr 从字符串的开头移除所有给定的 `cut` 字符串。
// 注意，它不会移除其开头的空白字符。
func X过滤首字符(文本 string, 过滤符号 string, 可选过滤次数 ...int) string {
	var (
		lenCut   = len(过滤符号)
		cutCount = 0
	)
	for len(文本) >= lenCut && 文本[0:lenCut] == 过滤符号 {
		文本 = 文本[lenCut:]
		cutCount++
		if len(可选过滤次数) > 0 && 可选过滤次数[0] != -1 && cutCount >= 可选过滤次数[0] {
			break
		}
	}
	return 文本
}

// TrimRight 从字符串的末尾删除空白字符（或其他字符）。
func X过滤尾字符并含空白(文本 string, 可选过滤符号 ...string) string {
	trimChars := utils.DefaultTrimChars
	if len(可选过滤符号) > 0 {
		trimChars += 可选过滤符号[0]
	}
	return strings.TrimRight(文本, trimChars)
}

// TrimRightStr 从字符串的末尾剥离所有给定的 `cut` 字符串。
// 注意，它不会剥离其末尾的空白字符。
func X过滤尾字符(文本 string, 过滤符号 string, 可选过滤次数 ...int) string {
	var (
		lenStr   = len(文本)
		lenCut   = len(过滤符号)
		cutCount = 0
	)
	for lenStr >= lenCut && 文本[lenStr-lenCut:lenStr] == 过滤符号 {
		lenStr = lenStr - lenCut
		文本 = 文本[:lenStr]
		cutCount++
		if len(可选过滤次数) > 0 && 可选过滤次数[0] != -1 && cutCount >= 可选过滤次数[0] {
			break
		}
	}
	return 文本
}

// TrimAll 删除字符串 `str` 中的所有字符。
func X过滤所有字符并含空白(文本 string, 可选过滤符号 ...string) string {
	trimChars := utils.DefaultTrimChars
	if len(可选过滤符号) > 0 {
		trimChars += 可选过滤符号[0]
	}
	var (
		filtered bool
		slice    = make([]rune, 0, len(文本))
	)
	for _, char := range 文本 {
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
func X开头判断(文本, 开头文本 string) bool {
	return strings.HasPrefix(文本, 开头文本)
}

// HasSuffix 测试字符串 s 是否以 suffix 结尾。
func X末尾判断(文本, 末尾文本 string) bool {
	return strings.HasSuffix(文本, 末尾文本)
}
