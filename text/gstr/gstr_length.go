// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstr


import (
	"unicode/utf8"
	)
// LenRune 返回字符串的Unicode长度。
func LenRune(str string) int {
	return utf8.RuneCountInString(str)
}
