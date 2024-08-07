// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文本类

import (
	"bytes"
	"strings"
	"unicode"
)

// X统计次数 计算字符串 `s` 中子串 `substr` 出现的次数。
// 如果 `s` 中不存在 `substr`，则返回 0。
// md5:eae2450dece3dd59
func X统计次数(文本, 子文本 string) int {
	return strings.Count(文本, 子文本)
}

// X统计次数并忽略大小写 计算字符串 `s` 中不区分大小写地出现 `substr` 的次数。
// 如果在 `s` 中没有找到 `substr`，则返回 0。
// md5:b8cd78125984641d
func X统计次数并忽略大小写(文本, 子文本 string) int {
	return strings.Count(X到小写(文本), X到小写(子文本))
}

// X统计单词数 返回字符串中单词的数量信息。
// 将参数 `str` 视为Unicode字符串。
// md5:825c2248f4df3a7d
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

// X统计unicode字符数 函数返回一个字符串中字符的计数信息。
// 它将参数 `str` 视为 Unicode 字符串。
// md5:365441f7aeb7819f
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
