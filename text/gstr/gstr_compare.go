// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gstr

import "strings"

// Compare函数返回一个整数，用于按字典顺序比较两个字符串。
// 如果a==b，结果为0；如果a < b，结果为-1；如果a > b，结果为+1。
// md5:a9e868f5edeb584f
func Compare(a, b string) int {
	return strings.Compare(a, b)
}

// Equal报告是否`a`和`b`，解释为UTF-8字符串，在Unicode大小写转换和不区分大小写的情况下相等。
// md5:4d0c0e7a06fc1578
func Equal(a, b string) bool {
	return strings.EqualFold(a, b)
}
