// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类

import (
	"strings"
)

// Contains 函数用于报告 `substr` 是否在 `str` 中，区分大小写。
func X是否包含(文本, 欲寻找的文本 string) bool {
	return strings.Contains(文本, 欲寻找的文本)
}

// ContainsI 判断 substr 是否在 str 中，忽略大小写。
func X是否包含并忽略大小写(文本, 欲寻找的文本 string) bool {
	return X查找并忽略大小写(文本, 欲寻找的文本) != -1
}

// ContainsAny 报告字符串 `s` 中是否包含任意来自 `chars` 中的Unicode字符。
func X是否包含Any(文本, 欲寻找的文本 string) bool {
	return strings.ContainsAny(文本, 欲寻找的文本)
}
