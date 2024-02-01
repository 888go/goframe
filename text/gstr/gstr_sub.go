// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstr


import (
	"strings"
	)
// Str 函数返回从 `haystack` 字符串中第一个（包含）出现的 `needle` 字符串开始，
// 直至 `haystack` 结尾的部分。参考：http://php.net/manual/en/function.strstr.php。
// 示例：Str("12345", "3") => "345"
func Str(haystack string, needle string) string {
	if needle == "" {
		return ""
	}
	pos := strings.Index(haystack, needle)
	if pos == NotFoundIndex {
		return ""
	}
	return haystack[pos+len([]byte(needle))-1:]
}

// StrEx 函数从 `haystack` 字符串中第一个不包含 `needle` 的位置开始，截取到 `haystack` 末尾的部分并返回。
// 示例：StrEx("12345", "3") => "45"
func StrEx(haystack string, needle string) string {
	if s := Str(haystack, needle); s != "" {
		return s[1:]
	}
	return ""
}

// StrTill 函数返回从 `haystack` 字符串起始位置到（并包含）
// 第一个出现的 `needle` 子串为止的部分字符串。
// 示例：StrTill("12345", "3") => 返回结果为 "123"
func StrTill(haystack string, needle string) string {
	pos := strings.Index(haystack, needle)
	if pos == NotFoundIndex || pos == 0 {
		return ""
	}
	return haystack[:pos+1]
}

// StrTillEx 从`haystack`字符串的起始位置截取，返回直到（但不包括）第一次出现`needle`子串的部分。
// 示例：StrTillEx("12345", "3") => "12"
func StrTillEx(haystack string, needle string) string {
	pos := strings.Index(haystack, needle)
	if pos == NotFoundIndex || pos == 0 {
		return ""
	}
	return haystack[:pos]
}

// SubStr 函数返回字符串 `str` 中由 `start` 和 `length` 参数指定的部分子串。
// 参数 `length` 是可选的，默认情况下会使用 `str` 的长度。
// 示例：SubStr("12345", 1, 2) => "23"
func SubStr(str string, start int, length ...int) (substr string) {
	strLength := len(str)
	if start < 0 {
		if -start > strLength {
			start = 0
		} else {
			start = strLength + start
		}
	} else if start > strLength {
		return ""
	}
	realLength := 0
	if len(length) > 0 {
		realLength = length[0]
		if realLength < 0 {
			if -realLength > strLength-start {
				realLength = 0
			} else {
				realLength = strLength - start + realLength
			}
		} else if realLength > strLength-start {
			realLength = strLength - start
		}
	} else {
		realLength = strLength - start
	}

	if realLength == strLength {
		return str
	} else {
		end := start + realLength
		return str[start:end]
	}
}

// SubStrRune 返回字符串 `str` 中由 `start` 和 `length` 参数指定的部分。
// SubStrRune 将参数 `str` 视为unicode字符串处理。
// 参数 `length` 是可选的，默认情况下它使用 `str` 的长度。
// 更详细的翻译：
// ```go
// SubStrRune 函数返回给定字符串 `str` 从 `start` 位置开始的一个子串。
// 在此函数中，我们把输入的字符串 `str` 当作Unicode字符序列进行处理。
// 参数 `length` 是可选的，如果不提供，则默认截取从 `start` 到字符串结尾的所有字符。
func SubStrRune(str string, start int, length ...int) (substr string) {
	// 转换为 []rune 以支持Unicode。
	var (
		runes       = []rune(str)
		runesLength = len(runes)
	)

	strLength := runesLength
	if start < 0 {
		if -start > strLength {
			start = 0
		} else {
			start = strLength + start
		}
	} else if start > strLength {
		return ""
	}
	realLength := 0
	if len(length) > 0 {
		realLength = length[0]
		if realLength < 0 {
			if -realLength > strLength-start {
				realLength = 0
			} else {
				realLength = strLength - start + realLength
			}
		} else if realLength > strLength-start {
			realLength = strLength - start
		}
	} else {
		realLength = strLength - start
	}
	end := start + realLength
	if end > runesLength {
		end = runesLength
	}
	return string(runes[start:end])
}

// StrLimit 返回字符串 `str` 指定长度为 `length` 的部分，如果 `str` 的长度大于 `length`，
// 则结果字符串后会追加 `suffix`。
func StrLimit(str string, length int, suffix ...string) string {
	if len(str) < length {
		return str
	}
	suffixStr := defaultSuffixForStrLimit
	if len(suffix) > 0 {
		suffixStr = suffix[0]
	}
	return str[0:length] + suffixStr
}

// StrLimitRune 返回字符串 `str` 指定长度的子串，若 `str` 的长度大于 `length`，
// 则结果字符串末尾会追加 `suffix`。此函数将参数 `str` 视为 unicode 字符串处理。
func StrLimitRune(str string, length int, suffix ...string) string {
	runes := []rune(str)
	if len(runes) < length {
		return str
	}
	suffixStr := defaultSuffixForStrLimit
	if len(suffix) > 0 {
		suffixStr = suffix[0]
	}
	return string(runes[0:length]) + suffixStr
}

// SubStrFrom 返回字符串 `str` 中从第一个出现并包含 `need` 的子串开始，直到 `str` 末尾的部分。
func SubStrFrom(str string, need string) (substr string) {
	pos := Pos(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos:]
}

// SubStrFromEx 从字符串 `str` 中返回从第一个出现且不包括 `need` 的子串到 `str` 结尾的部分。
func SubStrFromEx(str string, need string) (substr string) {
	pos := Pos(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos+len(need):]
}

// SubStrFromR 返回字符串 `str` 从最后一个出现且包含 `need` 的子串开始，直到 `str` 结尾的部分。
func SubStrFromR(str string, need string) (substr string) {
	pos := PosR(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos:]
}

// SubStrFromREx 函数从字符串 `str` 中返回从最后一次出现且不包含 `need` 的子串开始，直到 `str` 末尾的部分。
func SubStrFromREx(str string, need string) (substr string) {
	pos := PosR(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos+len(need):]
}
