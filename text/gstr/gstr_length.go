// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文本类

import (
	"unicode/utf8"
)

// LenRune返回Unicode字符串的长度。 md5:153854e98ec30abe
func LenRune(str string) int {
	return utf8.RuneCountInString(str)
}
