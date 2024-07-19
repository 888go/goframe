// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gstr

import (
	"bytes"
	"strings"
	"unicode"
)

// Count counts the number of `substr` appears in `s`.
// It returns 0 if no `substr` found in `s`.
// ff:统计次数
// s:文本
// substr:子文本
func Count(s, substr string) int {
	return strings.Count(s, substr)
}

// CountI counts the number of `substr` appears in `s`, case-insensitively.
// It returns 0 if no `substr` found in `s`.
// ff:统计次数并忽略大小写
// s:文本
// substr:子文本
func CountI(s, substr string) int {
	return strings.Count(ToLower(s), ToLower(substr))
}

// CountWords returns information about words' count used in a string.
// It considers parameter `str` as unicode string.
// ff:统计单词数
// str:文本
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

// CountChars returns information about chars' count used in a string.
// It considers parameter `str` as unicode string.
// ff:统计unicode字符数
// str:文本
// noSpace:是否计算空格
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
