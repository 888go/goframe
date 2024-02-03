// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstr

import (
	"bytes"
	"strings"
	"unicode"
)

// Count 计算 `substr` 在 `s` 中出现的次数。
// 如果在 `s` 中未找到 `substr`，则返回 0。
func Count(s, substr string) int {
	return strings.Count(s, substr)
}

// CountI 计算字符串 `s` 中不区分大小写出现的子串 `substr` 的次数。
// 如果在 `s` 中未找到 `substr`，则返回 0。
func CountI(s, substr string) int {
	return strings.Count(ToLower(s), ToLower(substr))
}

// CountWords 返回关于字符串中单词数量的信息。
// 它将参数`str`视为unicode字符串。
func CountWords(str string) map[string]int {
	m := make(map[string]int)
	buffer := bytes.NewBuffer(nil)
	for _, r := range []rune(str) {
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
func CountChars(str string, noSpace ...bool) map[string]int {
	m := make(map[string]int)
	countSpace := true
	if len(noSpace) > 0 && noSpace[0] {
		countSpace = false
	}
	for _, r := range []rune(str) {
		if !countSpace && unicode.IsSpace(r) {
			continue
		}
		m[string(r)]++
	}
	return m
}
