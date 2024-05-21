// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gstr

import (
	"bytes"

	"github.com/gogf/gf/v2/internal/utils"
)

// AddSlashes quotes with slashes `\` for chars: '"\.

// ff:转义
// str:文本
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

// StripSlashes un-quotes a quoted string by AddSlashes.

// ff:转义还原
// str:文本
func StripSlashes(str string) string {
	return utils.StripSlashes(str)
}

// QuoteMeta returns a version of `str` with a backslash character (`\`).
// If custom chars `chars` not given, it uses default chars: .\+*?[^]($)

// ff:转义并按字符
// chars:可选转义字符
// str:文本
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
