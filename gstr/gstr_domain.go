// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类

import (
	"strings"
)

// IsSubDomain 检查 `subDomain` 是否为 `mainDomain` 的子域名。
// 它支持在 `mainDomain` 中使用 '*'。
func X是否为子域名(子域名 string, 主域名 string) bool {
	if p := strings.IndexByte(子域名, ':'); p != -1 {
		子域名 = 子域名[0:p]
	}
	if p := strings.IndexByte(主域名, ':'); p != -1 {
		主域名 = 主域名[0:p]
	}
	var (
		subArray   = strings.Split(子域名, ".")
		mainArray  = strings.Split(主域名, ".")
		subLength  = len(subArray)
		mainLength = len(mainArray)
	)
// 示例：
// "goframe.org" 不是 "s.goframe.org" 的子域名。
	if mainLength > subLength {
		for i := range mainArray[0 : mainLength-subLength] {
			if mainArray[i] != "*" {
				return false
			}
		}
	}

// 示例：
// "s.s.goframe.org" 不是 "*.goframe.org" 的子域名，
// 但是
// "s.s.goframe.org" 是 "goframe.org" 的子域名。
	if mainLength > 2 && subLength > mainLength {
		return false
	}
	minLength := subLength
	if mainLength < minLength {
		minLength = mainLength
	}
	for i := minLength; i > 0; i-- {
		if mainArray[mainLength-i] == "*" {
			continue
		}
		if mainArray[mainLength-i] != subArray[subLength-i] {
			return false
		}
	}
	return true
}
