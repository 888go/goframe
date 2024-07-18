// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gstr

import (
	"strings"

	"github.com/gogf/gf/v2/internal/utils"
)

// ToLower 返回一个字符串 s 的副本，其中所有Unicode字母都被转换为小写。 md5:d70b7de319e04fa7
// ff:到小写
// s:文本
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper 返回一个将所有 Unicode 字母映射为其大写形式的 s 的副本。 md5:b816796c284fd785
// ff:到大写
// s:文本
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// UcFirst 返回一个字符串s的副本，其中第一个字母映射为其大写形式。 md5:bc090531eef4b3e6
// ff:到首字母大写
// s:文本
func UcFirst(s string) string {
	return utils.UcFirst(s)
}

// LcFirst返回一个字符串s的副本，其中第一个字母映射为其小写形式。 md5:1597253ba084ce1e
// ff:到首字母小写
// s:文本
func LcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	if IsLetterUpper(s[0]) {
		return string(s[0]+32) + s[1:]
	}
	return s
}

// UcWords 将字符串中每个单词的首字母转换为大写。 md5:b471982d531c9077
// ff:到单词首字母大写
// str:文本
func UcWords(str string) string {
	return strings.Title(str)
}

// IsLetterLower 检查给定的字节 b 是否为小写字母。 md5:f298f88a463e6078
// ff:是否小写字符
// b:字符
func IsLetterLower(b byte) bool {
	return utils.IsLetterLower(b)
}

// IsLetterUpper 检查给定的字节 b 是否为大写字母。 md5:dfb8879b42135673
// ff:是否大写字符
// b:字符
func IsLetterUpper(b byte) bool {
	return utils.IsLetterUpper(b)
}
