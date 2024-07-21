// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gstr

import (
	"bytes"
	"strings"
	"unicode"
)

// Count 计算字符串 `s` 中子串 `substr` 出现的次数。
// 如果 `s` 中不存在 `substr`，则返回 0。
// md5:eae2450dece3dd59
// ff:统计次数
// s:文本
// substr:子文本
func Count(s, substr string) int {
	return strings.Count(s, substr)
}

// CountI 计算字符串 `s` 中不区分大小写地出现 `substr` 的次数。
// 如果在 `s` 中没有找到 `substr`，则返回 0。
// md5:b8cd78125984641d
// ff:统计次数并忽略大小写
// s:文本
// substr:子文本
func CountI(s, substr string) int {
	return strings.Count(ToLower(s), ToLower(substr))
}

// CountWords 返回字符串中单词的数量信息。
// 将参数 `str` 视为Unicode字符串。
// md5:825c2248f4df3a7d
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

	// CountChars 函数返回一个字符串中字符的计数信息。
	// 它将参数 `str` 视为 Unicode 字符串。
	// md5:365441f7aeb7819f
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
