// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文本类

import (
	"strings"

	"github.com/888go/goframe/internal/utils"
)

// Replace 函数返回一个副本字符串 `origin`
// 在这个副本中，所有字符串 `search` 都会被 `replace` 替换，此替换区分大小写。
// md5:452df01ae43b07d9
func Replace(origin, search, replace string, count ...int) string {
	n := -1
	if len(count) > 0 {
		n = count[0]
	}
	return strings.Replace(origin, search, replace, n)
}

// ReplaceI 返回一个字符串 `origin` 的副本，
// 在该副本中不区分大小写地将字符串 `search` 替换为 `replace`。
// md5:f667575fd12d3732
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

// ReplaceByArray返回一个`origin`的副本，它被按顺序、区分大小写的方式替换为一个切片。
// md5:3b7b1a35fd597e47
func ReplaceByArray(origin string, array []string) string {
	for i := 0; i < len(array); i += 2 {
		if i+1 >= len(array) {
			break
		}
		origin = Replace(origin, array[i], array[i+1])
	}
	return origin
}

// ReplaceIByArray 返回一个副本 `origin`，它被按顺序、不区分大小写地替换为一个切片。
// md5:45d1fbd66515d9dd
func ReplaceIByArray(origin string, array []string) string {
	for i := 0; i < len(array); i += 2 {
		if i+1 >= len(array) {
			break
		}
		origin = ReplaceI(origin, array[i], array[i+1])
	}
	return origin
}

// ReplaceByMap 返回一个`origin`的副本，
// 使用映射无序地替换其中的内容，且区分大小写。
// md5:c047c08d8be640ad
func ReplaceByMap(origin string, replaces map[string]string) string {
	return utils.ReplaceByMap(origin, replaces)
}

// ReplaceIByMap 返回 `origin` 的一个副本，
// 其中内容被一个映射无序地替换，且替换时不区分大小写。
// md5:5b002ab4f7bd0cd8
func ReplaceIByMap(origin string, replaces map[string]string) string {
	for k, v := range replaces {
		origin = ReplaceI(origin, k, v)
	}
	return origin
}
