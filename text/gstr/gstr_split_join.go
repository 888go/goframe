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
	"github.com/gogf/gf/v2/util/gconv"
)

// Split将字符串`str`按照字符串`delimiter`进行分割，返回一个数组。 md5:905c146c396933a0
func Split(str, delimiter string) []string {
	return strings.Split(str, delimiter)
}

// SplitAndTrim 通过字符串 `delimiter` 将字符串 `str` 分割成一个数组，然后对数组中的每个元素调用 Trim 函数。它会忽略 Trim 后为空的元素。
// md5:20d7e1d120928c19
func SplitAndTrim(str, delimiter string, characterMask ...string) []string {
	return utils.SplitAndTrim(str, delimiter, characterMask...)
}

// Join 函数将 `array` 中的元素连接起来以创建一个单一的字符串。分隔符字符串 `sep` 会被放置在结果字符串中各元素之间。
// md5:c796a29e291a2864
func Join(array []string, sep string) string {
	return strings.Join(array, sep)
}

// JoinAny 将 `array` 中的元素连接成一个单一的字符串。在结果字符串中，元素之间由分隔符字符串 `sep` 分隔。
//
// 参数 `array` 可以是任何类型的切片，它将被转换为字符串数组。
// md5:fc531415278a603b
func JoinAny(array interface{}, sep string) string {
	return strings.Join(gconv.Strings(array), sep)
}

// Explode 函数将字符串 `str` 按照字符串 `delimiter` 进行分割，返回一个数组。
// 参考：http://php.net/manual/zh/function.explode.php。
// md5:28fb7a55d9ec56dc
func Explode(delimiter, str string) []string {
	return Split(str, delimiter)
}

// Implode 将数组元素 `pieces` 用字符串 `glue` 进行连接。// 参考：http://php.net/manual/zh/function.implode.php
// md5:e20b4f4a627b156b
func Implode(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}

// ChunkSplit 将字符串分割成更小的块。
// 可用于将字符串分割成更小的块，这对于
// 例如，将BASE64字符串输出转换为符合RFC 2045语义的情况非常有用。
// 它会在每chunkLen个字符后插入结束标记`end`。
// 它将参数`body`和`end`视为Unicode字符串处理。
// md5:94b7c0d7df7ca2e3
func ChunkSplit(body string, chunkLen int, end string) string {
	if end == "" {
		end = "\r\n"
	}
	runes, endRunes := []rune(body), []rune(end)
	l := len(runes)
	if l <= 1 || l < chunkLen {
		return body + end
	}
	ns := make([]rune, 0, len(runes)+len(endRunes))
	for i := 0; i < l; i += chunkLen {
		if i+chunkLen > l {
			ns = append(ns, runes[i:]...)
		} else {
			ns = append(ns, runes[i:i+chunkLen]...)
		}
		ns = append(ns, endRunes...)
	}
	return string(ns)
}

// Fields 将字符串中的单词返回为一个切片。 md5:b66b97aa739d583c
func Fields(str string) []string {
	return strings.Fields(str)
}
