// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstr
import (
	"strings"
	
	"github.com/888go/goframe/internal/utils"
	)
// ToLower 返回一个字符串 s 的副本，其中所有 Unicode 字母都转换为小写。
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper 返回字符串 s 的副本，其中所有 Unicode 字母都转换为它们的大写形式。
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// UcFirst 返回一个字符串 s 的副本，其中首字母已转换为大写。
func UcFirst(s string) string {
	return utils.UcFirst(s)
}

// LcFirst 返回一个字符串s的副本，其中首字母被转换为小写。
func LcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	if IsLetterUpper(s[0]) {
		return string(s[0]+32) + s[1:]
	}
	return s
}

// UcWords 将字符串中每个单词的首字母转换为大写。
func UcWords(str string) string {
	return strings.Title(str)
}

// IsLetterLower 测试给定的字节 b 是否为小写字母。
func IsLetterLower(b byte) bool {
	return utils.IsLetterLower(b)
}

// IsLetterUpper测试给定的字节b是否为大写字母。
func IsLetterUpper(b byte) bool {
	return utils.IsLetterUpper(b)
}
