// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类

import (
	"strings"
	
	"github.com/888go/goframe/gstr/internal/utils"
)

// Replace 返回 `origin` 字符串的一个副本，
// 其中字符串 `search` 被 `replace` 按照大小写敏感的方式替换。
func X替换(文本, 替换文本, 用作替换文本 string, 替换次数 ...int) string {
	n := -1
	if len(替换次数) > 0 {
		n = 替换次数[0]
	}
	return strings.Replace(文本, 替换文本, 用作替换文本, n)
}

// ReplaceI返回一个`origin`字符串的副本，
// 其中不区分大小写地将字符串`search`替换为`replace`。
func X替换并忽略大小写(文本, 替换文本, 用作替换文本 string, 替换次数 ...int) string {
	n := -1
	if len(替换次数) > 0 {
		n = 替换次数[0]
	}
	if n == 0 {
		return 文本
	}
	var (
		searchLength  = len(替换文本)
		replaceLength = len(用作替换文本)
		searchLower   = strings.ToLower(替换文本)
		originLower   string
		pos           int
	)
	for {
		originLower = strings.ToLower(文本)
		if pos = X查找(originLower, searchLower, pos); pos != -1 {
			文本 = 文本[:pos] + 用作替换文本 + 文本[pos+searchLength:]
			pos += replaceLength
			if n--; n == 0 {
				break
			}
		} else {
			break
		}
	}
	return 文本
}

// ReplaceByArray 返回 `origin` 的副本，
// 其中内容会按照顺序被切片中的元素替换，并且是区分大小写的。
func X数组替换(文本 string, 数组 []string) string {
	for i := 0; i < len(数组); i += 2 {
		if i+1 >= len(数组) {
			break
		}
		文本 = X替换(文本, 数组[i], 数组[i+1])
	}
	return 文本
}

// ReplaceIByArray 返回 `origin` 的副本，
// 其中内容按照顺序被切片替换，并且不区分大小写。
func X数组替换并忽略大小写(文本 string, 数组 []string) string {
	for i := 0; i < len(数组); i += 2 {
		if i+1 >= len(数组) {
			break
		}
		文本 = X替换并忽略大小写(文本, 数组[i], 数组[i+1])
	}
	return 文本
}

// ReplaceByMap 函数返回 `origin` 的一个副本，
// 并使用一个无序的映射进行替换，且替换操作区分大小写。
func Map替换(文本 string, 用作替换的Map map[string]string) string {
	return utils.ReplaceByMap(文本, 用作替换的Map)
}

// ReplaceIByMap 返回 `origin` 的副本，
// 其中内容将以无序方式、不区分大小写地通过一个映射表进行替换。
func Map替换并忽略大小写(文本 string, map数组 map[string]string) string {
	for k, v := range map数组 {
		文本 = X替换并忽略大小写(文本, k, v)
	}
	return 文本
}
