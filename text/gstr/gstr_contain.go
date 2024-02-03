// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstr

import (
	"strings"
)

// Contains 函数用于报告 `substr` 是否在 `str` 中，区分大小写。
func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// ContainsI 判断 substr 是否在 str 中，忽略大小写。
func ContainsI(str, substr string) bool {
	return PosI(str, substr) != -1
}

// ContainsAny 报告字符串 `s` 中是否包含任意来自 `chars` 中的Unicode字符。
func ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}
