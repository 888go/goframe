// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类

import (
	"bytes"
	"strings"
	"unicode"
)

// Count 计算 `substr` 在 `s` 中出现的次数。
// 如果在 `s` 中未找到 `substr`，则返回 0。
func X统计次数(文本, 子文本 string) int {
	return strings.Count(文本, 子文本)
}

// CountI 计算字符串 `s` 中不区分大小写出现的子串 `substr` 的次数。
// 如果在 `s` 中未找到 `substr`，则返回 0。
func X统计次数并忽略大小写(文本, 子文本 string) int {
	return strings.Count(X到小写(文本), X到小写(子文本))
}

// CountWords 返回关于字符串中单词数量的信息。
// 它将参数`str`视为unicode字符串。
func X统计单词数(文本 string) map[string]int {
	m := make(map[string]int)
	buffer := bytes.NewBuffer(nil)
	for _, r := range []rune(文本) {
		if unicode.IsSpace(r) {
			if buffer.Len() > 0 {
				m[buffer.String()]++
				buffer.Reset()
			}
		} else {
			buffer.WriteRune(r)
		}
	}
	if buffer.Len() > 0 {
		m[buffer.String()]++
	}
	return m
}

// CountChars 返回关于字符串中使用字符数量的信息。
// 它将参数 `str` 视为unicode字符串。
func X统计unicode字符数(文本 string, 是否计算空格 ...bool) map[string]int {
	m := make(map[string]int)
	countSpace := true
	if len(是否计算空格) > 0 && 是否计算空格[0] {
		countSpace = false
	}
	for _, r := range []rune(文本) {
		if !countSpace && unicode.IsSpace(r) {
			continue
		}
		m[string(r)]++
	}
	return m
}
