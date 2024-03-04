// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstr

import (
	"strings"
	
	"github.com/888go/goframe/gstr/internal/utils"
)

// Replace 返回 `origin` 字符串的一个副本，
// 其中字符串 `search` 被 `replace` 按照大小写敏感的方式替换。
func Replace(origin, search, replace string, count ...int) string {
	n := -1
	if len(count) > 0 {
		n = count[0]
	}
	return strings.Replace(origin, search, replace, n)
}

// ReplaceI返回一个`origin`字符串的副本，
// 其中不区分大小写地将字符串`search`替换为`replace`。
func ReplaceI(origin, search, replace string, count ...int) string {
	n := -1
	if len(count) > 0 {
		n = count[0]
	}
	if n == 0 {
		return origin
	}
	var (
		searchLength  = len(search)
		replaceLength = len(replace)
		searchLower   = strings.ToLower(search)
		originLower   string
		pos           int
	)
	for {
		originLower = strings.ToLower(origin)
		if pos = Pos(originLower, searchLower, pos); pos != -1 {
			origin = origin[:pos] + replace + origin[pos+searchLength:]
			pos += replaceLength
			if n--; n == 0 {
				break
			}
		} else {
			break
		}
	}
	return origin
}

// ReplaceByArray 返回 `origin` 的副本，
// 其中内容会按照顺序被切片中的元素替换，并且是区分大小写的。
func ReplaceByArray(origin string, array []string) string {
	for i := 0; i < len(array); i += 2 {
		if i+1 >= len(array) {
			break
		}
		origin = Replace(origin, array[i], array[i+1])
	}
	return origin
}

// ReplaceIByArray 返回 `origin` 的副本，
// 其中内容按照顺序被切片替换，并且不区分大小写。
func ReplaceIByArray(origin string, array []string) string {
	for i := 0; i < len(array); i += 2 {
		if i+1 >= len(array) {
			break
		}
		origin = ReplaceI(origin, array[i], array[i+1])
	}
	return origin
}

// ReplaceByMap 函数返回 `origin` 的一个副本，
// 并使用一个无序的映射进行替换，且替换操作区分大小写。
func ReplaceByMap(origin string, replaces map[string]string) string {
	return utils.ReplaceByMap(origin, replaces)
}

// ReplaceIByMap 返回 `origin` 的副本，
// 其中内容将以无序方式、不区分大小写地通过一个映射表进行替换。
func ReplaceIByMap(origin string, replaces map[string]string) string {
	for k, v := range replaces {
		origin = ReplaceI(origin, k, v)
	}
	return origin
}
