// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类

import (
	"strings"
)

// Compare 函数返回两个字符串的字典序比较结果，是一个整数。
// 如果 a==b，则结果为 0；如果 a < b，则结果为 -1；如果 a > b，则结果为 +1。
func X顺序比较(a, b string) int {
	return strings.Compare(a, b)
}

// Equal 判断当 `a` 和 `b` 被解释为 UTF-8 字符串时，
// 在Unicode大小写折叠规则下，它们是否大小写不敏感地相等。
func X相等比较并忽略大小写(a, b string) bool {
	return strings.EqualFold(a, b)
}
