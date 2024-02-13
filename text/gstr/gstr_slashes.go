// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类

import (
	"bytes"
	
	"github.com/888go/goframe/internal/utils"
)

// AddSlashes 为特定字符添加反斜杠 `\` 进行转义，这些字符包括：`"`、`\` 和 `.`。
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

// StripSlashes 将通过 AddSlashes 方法添加了反斜杠的字符串进行去除引用操作。
func X转义还原(文本 string) string {
	return utils.StripSlashes(文本)
}

// QuoteMeta 返回一个带有反斜杠字符 (`\`) 的 `str` 版本。
// 若未提供自定义字符集 `chars`，则使用默认字符集：.+\*?[^]($)
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
