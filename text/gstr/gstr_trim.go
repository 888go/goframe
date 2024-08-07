// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文本类

import (
	"strings"

	"github.com/888go/goframe/internal/utils"
)

// X过滤首尾符并含空白 函数从字符串的开头和结尾删除空白符（或其他字符）。
// 可选参数 `characterMask` 指定了额外需要删除的字符。
// md5:a1d794d7a10c9c7a
func X过滤首尾符并含空白(文本 string, 可选过滤符号 ...string) string {
	return utils.Trim(文本, 可选过滤符号...)
}

// X过滤首尾 从字符串的开始和结束处删除给定的 `cut` 字符串。
// 请注意，它不会删除其开始或结束处的空格。
// md5:6180dbb1bb9965c4
func X过滤首尾(文本 string, 过滤符号 string, 可选过滤次数 ...int) string {
	return X过滤首字符(X过滤尾字符(文本, 过滤符号, 可选过滤次数...), 过滤符号, 可选过滤次数...)
}

// X过滤首字符并含空白 从字符串的开头移除空格（或其他字符）。 md5:648a4eb609597439
func X过滤首字符并含空白(文本 string, 可选过滤符号 ...string) string {
	trimChars := utils.DefaultTrimChars
	if len(可选过滤符号) > 0 {
		trimChars += 可选过滤符号[0]
	}
	return strings.TrimLeft(文本, trimChars)
}

// X过滤首字符 从字符串的开头移除所有给定的 `cut` 字符串。注意，它不会移除字符串开头的空格。
// md5:765cffbaed7e3cce
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

// X过滤尾字符并含空白 从字符串的末尾移除空白符（或其他字符）。 md5:c252aae10c16cd5a
func X过滤尾字符并含空白(文本 string, 可选过滤符号 ...string) string {
	trimChars := utils.DefaultTrimChars
	if len(可选过滤符号) > 0 {
		trimChars += 可选过滤符号[0]
	}
	return strings.TrimRight(文本, trimChars)
}

// X过滤尾字符 从字符串的末尾删除所有给定的 `cut` 字符串。
// 请注意，它不会删除字符串末尾的空格。
// md5:0ca4b90c9d600b39
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

// X过滤所有字符并含空白 函数会删除字符串 `str` 中的所有字符。 md5:2d770fefafe6bda4
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

// X开头判断 测试字符串 s 是否以 prefix 开头。 md5:8b4cd90bff10b28a
func X开头判断(文本, 开头文本 string) bool {
	return strings.HasPrefix(文本, 开头文本)
}

// X末尾判断 检测字符串 s 是否以 suffix 结尾。 md5:964d208ab1e843e1
func X末尾判断(文本, 末尾文本 string) bool {
	return strings.HasSuffix(文本, 末尾文本)
}
