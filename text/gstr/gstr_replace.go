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

// X替换 函数返回一个副本字符串 `origin`
// 在这个副本中，所有字符串 `search` 都会被 `replace` 替换，此替换区分大小写。
// md5:452df01ae43b07d9
func X替换(文本, 替换文本, 用作替换文本 string, 替换次数 ...int) string {
	n := -1
	if len(替换次数) > 0 {
		n = 替换次数[0]
	}
	return strings.Replace(文本, 替换文本, 用作替换文本, n)
}

// X替换并忽略大小写 返回一个字符串 `origin` 的副本，
// 在该副本中不区分大小写地将字符串 `search` 替换为 `replace`。
// md5:f667575fd12d3732
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

// X切片替换返回一个`origin`的副本，它被按顺序、区分大小写的方式替换为一个切片。
// md5:3b7b1a35fd597e47
func X切片替换(文本 string, 切片 []string) string {
	for i := 0; i < len(切片); i += 2 {
		if i+1 >= len(切片) {
			break
		}
		文本 = X替换(文本, 切片[i], 切片[i+1])
	}
	return 文本
}

// X切片替换并忽略大小写 返回一个副本 `origin`，它被按顺序、不区分大小写地替换为一个切片。
// md5:45d1fbd66515d9dd
func X切片替换并忽略大小写(文本 string, 切片 []string) string {
	for i := 0; i < len(切片); i += 2 {
		if i+1 >= len(切片) {
			break
		}
		文本 = X替换并忽略大小写(文本, 切片[i], 切片[i+1])
	}
	return 文本
}

// Map替换 返回一个`origin`的副本，
// 使用映射无序地替换其中的内容，且区分大小写。
// md5:c047c08d8be640ad
func Map替换(文本 string, 用作替换的Map map[string]string) string {
	return utils.ReplaceByMap(文本, 用作替换的Map)
}

// Map替换并忽略大小写 返回 `origin` 的一个副本，
// 其中内容被一个映射无序地替换，且替换时不区分大小写。
// md5:5b002ab4f7bd0cd8
func Map替换并忽略大小写(文本 string, map切片 map[string]string) string {
	for k, v := range map切片 {
		文本 = X替换并忽略大小写(文本, k, v)
	}
	return 文本
}
