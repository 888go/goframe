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

// X生成重复文本 函数返回一个新字符串，该字符串由输入字符串乘数次复制组成。
//
// 示例：
// X生成重复文本("a", 3) -> "aaa"
// md5:e4200cea299476dc
func X生成重复文本(文本 string, 重复次数 int) string {
	return strings.Repeat(文本, 重复次数)
}
