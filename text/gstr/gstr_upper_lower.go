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

// X到小写 返回一个字符串 s 的副本，其中所有Unicode字母都被转换为小写。 md5:d70b7de319e04fa7
func X到小写(文本 string) string {
	return strings.ToLower(文本)
}

// X到大写 返回一个将所有 Unicode 字母映射为其大写形式的 s 的副本。 md5:b816796c284fd785
func X到大写(文本 string) string {
	return strings.ToUpper(文本)
}

// X到首字母大写 返回一个字符串s的副本，其中第一个字母映射为其大写形式。 md5:bc090531eef4b3e6
func X到首字母大写(文本 string) string {
	return utils.UcFirst(文本)
}

// X到首字母小写返回一个字符串s的副本，其中第一个字母映射为其小写形式。 md5:1597253ba084ce1e
func X到首字母小写(文本 string) string {
	if len(文本) == 0 {
		return 文本
	}
	if X是否大写字符(文本[0]) {
		return string(文本[0]+32) + 文本[1:]
	}
	return 文本
}

// X到单词首字母大写 将字符串中每个单词的首字母转换为大写。 md5:b471982d531c9077
func X到单词首字母大写(文本 string) string {
	return strings.Title(文本)
}

// X是否小写字符 检查给定的字节 b 是否为小写字母。 md5:f298f88a463e6078
func X是否小写字符(字符 byte) bool {
	return utils.IsLetterLower(字符)
}

// X是否大写字符 检查给定的字节 b 是否为大写字母。 md5:dfb8879b42135673
func X是否大写字符(字符 byte) bool {
	return utils.IsLetterUpper(字符)
}
