// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类

import (
	"strings"
)

// Repeat 返回一个新的字符串，该字符串由 input 字符串重复 multiplier 次组成。
func X生成重复文本(文本 string, 重复次数 int) string {
	return strings.Repeat(文本, 重复次数)
}
