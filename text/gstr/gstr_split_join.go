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
	gconv "github.com/888go/goframe/util/gconv"
)

// X分割将字符串`str`按照字符串`delimiter`进行分割，返回一个数组。 md5:905c146c396933a0
func X分割(文本, 用作分割的文本 string) []string {
	return strings.Split(文本, 用作分割的文本)
}

// X分割并忽略空值 通过字符串 `delimiter` 将字符串 `str` 分割成一个数组，然后对数组中的每个元素调用 Trim 函数。它会忽略 Trim 后为空的元素。
// md5:20d7e1d120928c19
func X分割并忽略空值(文本, 用作分割的文本 string, characterMask ...string) []string {
	return utils.SplitAndTrim(文本, 用作分割的文本, characterMask...)
}

// X连接 函数将 `array` 中的元素连接起来以创建一个单一的字符串。分隔符字符串 `sep` 会被放置在结果字符串中各元素之间。
// md5:c796a29e291a2864
func X连接(切片 []string, 连接符 string) string {
	return strings.Join(切片, 连接符)
}

// X连接Any 将 `array` 中的元素连接成一个单一的字符串。在结果字符串中，元素之间由分隔符字符串 `sep` 分隔。
//
// 参数 `array` 可以是任何类型的切片，它将被转换为字符串数组。
// md5:fc531415278a603b
func X连接Any(切片 interface{}, 连接符 string) string {
	return strings.Join(gconv.X取文本切片(切片), 连接符)
}

// Explode别名 函数将字符串 `str` 按照字符串 `delimiter` 进行分割，返回一个数组。
// 参考：http://php.net/manual/zh/function.explode.php。
// md5:28fb7a55d9ec56dc
func Explode别名(delimiter, str string) []string {
	return X分割(str, delimiter)
}

// Implode别名 将数组元素 `pieces` 用字符串 `glue` 进行连接。// 参考：http://php.net/manual/zh/function.implode.php
// md5:e20b4f4a627b156b
func Implode别名(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}

// X长度分割 将字符串分割成更小的块。
// 可用于将字符串分割成更小的块，这对于
// 例如，将BASE64字符串输出转换为符合RFC 2045语义的情况非常有用。
// 它会在每chunkLen个字符后插入结束标记`end`。
// 它将参数`body`和`end`视为Unicode字符串处理。
// md5:94b7c0d7df7ca2e3
func X长度分割(文本 string, 分割长度 int, 分割符 string) string {
	if 分割符 == "" {
		分割符 = "\r\n"
	}
	runes, endRunes := []rune(文本), []rune(分割符)
	l := len(runes)
	if l <= 1 || l < 分割长度 {
		return 文本 + 分割符
	}
	ns := make([]rune, 0, len(runes)+len(endRunes))
	for i := 0; i < l; i += 分割长度 {
		if i+分割长度 > l {
			ns = append(ns, runes[i:]...)
		} else {
			ns = append(ns, runes[i:i+分割长度]...)
		}
		ns = append(ns, endRunes...)
	}
	return string(ns)
}

// X单词分割 将字符串中的单词返回为一个切片。 md5:b66b97aa739d583c
func X单词分割(文本 string) []string {
	return strings.Fields(文本)
}
