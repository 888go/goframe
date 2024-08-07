// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文本类

import (
	"github.com/888go/goframe/internal/utils"
)

// X是否为数字 检查给定的字符串 s 是否为数字。 md5:c82137abf4b164a4
func X是否为数字(值 string) bool {
	return utils.IsNumeric(值)
}
