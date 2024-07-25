// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gstr

import "strings"

// IsSubDomain 检查 `subDomain` 是否为 `mainDomain` 的子域名。
// 它支持在 `mainDomain` 中使用通配符 '*'。 md5:a9d75f1129f8ee85
func IsSubDomain(subDomain string, mainDomain string) bool {
	if p := strings.IndexByte(subDomain, ':'); p != -1 {
		subDomain = subDomain[0:p]
	}
	if p := strings.IndexByte(mainDomain, ':'); p != -1 {
		mainDomain = mainDomain[0:p]
	}
	var (
		subArray   = strings.Split(subDomain, ".")
		mainArray  = strings.Split(mainDomain, ".")
		subLength  = len(subArray)
		mainLength = len(mainArray)
	)
	// 例如：
	// "goframe.org" 不是 "s.goframe.org" 的子域名。 md5:82e988c659c65b8f
	if mainLength > subLength {
		for i := range mainArray[0 : mainLength-subLength] {
			if mainArray[i] != "*" {
				return false
			}
		}
	}

	// 例如：
	// "s.s.goframe.org" 不是 "*.goframe.org" 的子域名
	// 但是
	// "s.s.goframe.org" 是 "goframe.org" 的子域名 md5:8cdcf05bb87c9aa4
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
