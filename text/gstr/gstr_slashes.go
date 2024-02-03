// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstr

import (
	"bytes"
	
	"github.com/888go/goframe/internal/utils"
)

// AddSlashes 为特定字符添加反斜杠 `\` 进行转义，这些字符包括：`"`、`\` 和 `.`。
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

// StripSlashes 将通过 AddSlashes 方法添加了反斜杠的字符串进行去除引用操作。
func StripSlashes(str string) string {
	return utils.StripSlashes(str)
}

// QuoteMeta 返回一个带有反斜杠字符 (`\`) 的 `str` 版本。
// 若未提供自定义字符集 `chars`，则使用默认字符集：.+\*?[^]($)
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
