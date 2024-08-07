// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文本类

import (
	"strings"
)

// X是否包含报告`substr`是否在`str`中（区分大小写）。 md5:23e258f54d1c33e6
func X是否包含(文本, 欲寻找的文本 string) bool {
	return strings.Contains(文本, 欲寻找的文本)
}

// X是否包含并忽略大小写 报告 substr 是否在 str 中，不区分大小写。 md5:1dee4dc6bff3de90
func X是否包含并忽略大小写(文本, 欲寻找的文本 string) bool {
	return X查找并忽略大小写(文本, 欲寻找的文本) != -1
}

// X是否包含Any 报告 `s` 中是否包含 `chars` 里的任何Unicode字符。 md5:68982aa112f62c02
func X是否包含Any(文本, 欲寻找的文本 string) bool {
	return strings.ContainsAny(文本, 欲寻找的文本)
}
