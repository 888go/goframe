// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类

import (
	"strings"
	
	"github.com/888go/goframe/gstr/internal/utils"
	"github.com/gogf/gf/v2/util/gconv"
)

// Split 函数通过一个分隔符字符串 `delimiter` 将字符串 `str` 分割成一个数组。
func X分割(文本, 用作分割的文本 string) []string {
	return strings.Split(文本, 用作分割的文本)
}

// SplitAndTrim通过字符串`delimiter`将字符串`str`分割成一个数组，
// 然后对数组中的每个元素调用Trim方法。它会忽略经过Trim处理后为空的元素。
func X分割并忽略空值(文本, 用作分割的文本 string, characterMask ...string) []string {
	return utils.SplitAndTrim(文本, 用作分割的文本, characterMask...)
}

// Join 函数将 `array` 中的元素连接起来生成一个单一的字符串。在生成的字符串中，分隔符字符串 `sep` 会被放置在各个元素之间。
func X连接(数组 []string, 连接符 string) string {
	return strings.Join(数组, 连接符)
}

// JoinAny 通过连接 `array` 中的元素来创建一个单一字符串。分隔符字符串
// `sep` 将会被放置在结果字符串中各元素之间。
//
// 参数 `array` 可以是任意类型的切片，只要它可以转换为字符串数组。
func X连接Any(数组 interface{}, 连接符 string) string {
	return strings.Join(gconv.Strings(数组), 连接符)
}

// Explode 将字符串 `str` 通过指定的分隔符 `delimiter` 进行拆分，结果存入数组中。
// 参考：http://php.net/manual/en/function.explode.php.
func Explode别名(delimiter, str string) []string {
	return X分割(str, delimiter)
}

// Implode 函数通过字符串 `glue` 连接数组元素 `pieces`。
// 参考：http://php.net/manual/en/function.implode.php
func Implode别名(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}

// ChunkSplit 将一个字符串分割成更小的块。
// 可用于将字符串分割成较小的块，这对于例如将 BASE64 字符串输出转换为匹配 RFC 2045 语义非常有用。
// 它会在每 chunkLen 个字符后插入 end。
// 它将参数 `body` 和 `end` 视为 unicode 字符串。
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

// Fields返回字符串中使用的单词作为一个切片。
func X单词分割(文本 string) []string {
	return strings.Fields(文本)
}
