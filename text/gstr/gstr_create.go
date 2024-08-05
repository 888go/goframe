// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gstr

import "strings"

// Repeat 函数返回一个新字符串，该字符串由输入字符串乘数次复制组成。
//
// 示例：
// Repeat("a", 3) -> "aaa"
// md5:e4200cea299476dc
func Repeat(input string, multiplier int) string {
	return strings.Repeat(input, multiplier)
}
