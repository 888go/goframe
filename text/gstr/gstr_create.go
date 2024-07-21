// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gstr

import "strings"

// Repeat returns a new string consisting of multiplier copies of the string input.
//
// Repeat("a", 3) -> "aaa"
// ff:生成重复文本
// input:文本
// multiplier:重复次数
func Repeat(input string, multiplier int) string {
	return strings.Repeat(input, multiplier)
}
