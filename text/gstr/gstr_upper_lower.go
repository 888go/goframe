// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类

import (
	"strings"
	
	"github.com/888go/goframe/internal/utils"
)

// ToLower 返回一个字符串 s 的副本，其中所有 Unicode 字母都转换为小写。
func X到小写(文本 string) string {
	return strings.ToLower(文本)
}

// ToUpper 返回字符串 s 的副本，其中所有 Unicode 字母都转换为它们的大写形式。
func X到大写(文本 string) string {
	return strings.ToUpper(文本)
}

// UcFirst 返回一个字符串 s 的副本，其中首字母已转换为大写。
func X到首字母大写(文本 string) string {
	return utils.UcFirst(文本)
}

// LcFirst 返回一个字符串s的副本，其中首字母被转换为小写。
func X到首字母小写(文本 string) string {
	if len(文本) == 0 {
		return 文本
	}
	if X是否大写字符(文本[0]) {
		return string(文本[0]+32) + 文本[1:]
	}
	return 文本
}

// UcWords 将字符串中每个单词的首字母转换为大写。
func X到单词首字母大写(文本 string) string {
	return strings.Title(文本)
}

// IsLetterLower 测试给定的字节 b 是否为小写字母。
func X是否小写字符(字符 byte) bool {
	return utils.IsLetterLower(字符)
}

// IsLetterUpper测试给定的字节b是否为大写字母。
func X是否大写字符(字符 byte) bool {
	return utils.IsLetterUpper(字符)
}
