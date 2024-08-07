// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文本类

import (
	"bytes"

	"github.com/888go/goframe/internal/utils"
)

// X转义 为字符 '"'、'\' 和 '.' 添加反斜杠 `\` 进行转义。 md5:7e939499c0cd825f
func X转义(文本 string) string {
	var buf bytes.Buffer
	for _, char := range 文本 {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// X转义还原通过AddSlashes方法去引号化一个带引号的字符串。 md5:fdba1646284db234
func X转义还原(文本 string) string {
	return utils.StripSlashes(文本)
}

// X转义并按字符 返回一个带有反斜杠字符 (`\`) 的 `str` 的版本。
// 如果没有提供自定义字符 `chars`，则使用默认字符：.\+*?[^]($)
// md5:e6bfeeabc9838069
func X转义并按字符(文本 string, 可选转义字符 ...string) string {
	var buf bytes.Buffer
	for _, char := range 文本 {
		if len(可选转义字符) > 0 {
			for _, c := range 可选转义字符[0] {
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
