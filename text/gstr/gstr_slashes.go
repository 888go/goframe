// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gstr

import (
	"bytes"

	"github.com/gogf/gf/v2/internal/utils"
)

// AddSlashes 为字符 '"'、'\' 和 '.' 添加反斜杠 `\` 进行转义。. md5:7e939499c0cd825f
func AddSlashes(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// StripSlashes通过AddSlashes方法去引号化一个带引号的字符串。. md5:fdba1646284db234
func StripSlashes(str string) string {
	return utils.StripSlashes(str)
}

// QuoteMeta 返回一个带有反斜杠字符 (`\`) 的 `str` 的版本。
// 如果没有提供自定义字符 `chars`，则使用默认字符：.\+*?[^]($)
// md5:e6bfeeabc9838069
func QuoteMeta(str string, chars ...string) string {
	var buf bytes.Buffer
	for _, char := range str {
		if len(chars) > 0 {
			for _, c := range chars[0] {
				if c == char {
					buf.WriteRune('\\')
					break
				}
			}
		} else {
			switch char {
			case '.', '+', '\\', '(', '$', ')', '[', '^', ']', '*', '?':
				buf.WriteRune('\\')
			}
		}
		buf.WriteRune(char)
	}
	return buf.String()
}
