// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gstr

import "strings"

// Str 函数返回从 `haystack` 字符串开始，包括第一个出现的 `needle` 到 `haystack` 结尾的部分。
//
// 这个函数的行为与 SubStr 函数完全相同，但是为了实现与 PHP 相同的功能：http://php.net/manual/en/function.strstr.php。
//
// 示例：
// Str("av.mp4", ".") -> ".mp4"
// md5:35b0375920cdf357
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

// StrEx 从`haystack`字符串中返回从第一个出现的`needle`开始到`haystack`末尾的部分。
//
// 这个函数的行为与SubStrEx函数完全相同，但实现了与PHP相同的函数：http://php.net/manual/en/function.strstr.php。
//
// 示例：
// StrEx("av.mp4", ".") -> "mp4"
// md5:1f6c13b098caa33a
func StrEx(haystack string, needle string) string {
	if s := Str(haystack, needle); s != "" {
		return s[1:]
	}
	return ""
}

// StrTill 返回 `haystack` 字符串中从开头到（包括）第一个 `needle` 出现的部分。
//
// 示例：
// StrTill("av.mp4", ".") -> "av."
// md5:f914c156abb95437
func StrTill(haystack string, needle string) string {
	pos := strings.Index(haystack, needle)
	if pos == NotFoundIndex || pos == 0 {
		return ""
	}
	return haystack[:pos+1]
}

// StrTillEx 函数返回 `haystack` 字符串中从开始到（但不包括）第一次出现 `needle` 的部分。
//
// 示例：
// StrTillEx("av.mp4", ".") -> "av"
// md5:c0848291d8036d82
func StrTillEx(haystack string, needle string) string {
	pos := strings.Index(haystack, needle)
	if pos == NotFoundIndex || pos == 0 {
		return ""
	}
	return haystack[:pos]
}

// SubStr 函数返回字符串 `str` 中由 `start` 和 `length` 参数指定的部分。参数 `length` 是可选的，如果未提供，则默认使用 `str` 的长度。
// 
// 示例：
// SubStr("123456", 1, 2) -> "23"
// md5:b6da71b3534fdbbc
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

// SubStrRune returns a portion of string `str` specified by the `start` and `length` parameters.
// SubStrRune considers parameter `str` as unicode string.
// The parameter `length` is optional, it uses the length of `str` in default.
//
// Example:
// SubStrRune("一起学习吧！", 2, 2) -> "学习"
func SubStrRune(str string, start int, length ...int) (substr string) {
	// 转换为[]rune以支持Unicode。 md5:459540c13f4e5603
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

// StrLimit 函数返回字符串 `str` 中由 `length` 参数指定长度的部分。如果 `str` 的长度大于 `length`，则 `suffix` 将被添加到结果字符串的末尾。
//
// 示例：
// StrLimit("123456", 3)      -> "123..."
// StrLimit("123456", 3, "~") -> "123~"
// md5:bd8f96405a5594b5
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

// StrLimitRune returns a portion of string `str` specified by `length` parameters, if the length
// of `str` is greater than `length`, then the `suffix` will be appended to the result string.
// StrLimitRune considers parameter `str` as unicode string.
//
// Example:
// StrLimitRune("一起学习吧！", 2)      -> "一起..."
// StrLimitRune("一起学习吧！", 2, "~") -> "一起~"
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

// SubStrFrom 从字符串 `str` 中从第一次出现 `need` 的位置开始，包括 `need` 到字符串末尾的部分。
//
// 示例：
// SubStrFrom("av.mp4", ".") -> ".mp4"
// md5:f4bff02c473abeff
func SubStrFrom(str string, need string) (substr string) {
	pos := Pos(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos:]
}

// SubStrFromEx 从字符串 `str` 中返回从首次出现 `need` 到 `str` 结尾的部分（不包括 `need`）。
//
// 示例：
// SubStrFromEx("av.mp4", ".") -> "mp4"
// md5:88a817f03fc77455
func SubStrFromEx(str string, need string) (substr string) {
	pos := Pos(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos+len(need):]
}

// SubStrFromR 从字符串 `str` 的最后一个出现的 `need` 开始并包括在内，返回一个子串。
// 示例：
// SubStrFromR("/dev/vda", "/") -> "/vda"
// md5:8f70ecc84d0338f8
func SubStrFromR(str string, need string) (substr string) {
	pos := PosR(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos:]
}

// SubStrFromREx 从字符串 `str` 中最后一个出现的 `need` 子串之后的字符开始，直到 `str` 的末尾，返回这一部分子串。
//
// 示例：
// SubStrFromREx("/dev/vda", "/") -> "vda"
// md5:3de495ad97b12196
func SubStrFromREx(str string, need string) (substr string) {
	pos := PosR(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos+len(need):]
}
