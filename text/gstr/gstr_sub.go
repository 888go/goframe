// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gstr

import "strings"

// Str returns part of `haystack` string starting from and including
// the first occurrence of `needle` to the end of `haystack`.
//
// This function performs exactly as function SubStr, but to implement the same function
// as PHP: http://php.net/manual/en/function.strstr.php.
//
// Str("av.mp4", ".") -> ".mp4"
// ff:取右边并含关键字
// haystack:文本
// needle:欲寻找的文本
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

// StrEx returns part of `haystack` string starting from and excluding
// the first occurrence of `needle` to the end of `haystack`.
//
// This function performs exactly as function SubStrEx, but to implement the same function
// as PHP: http://php.net/manual/en/function.strstr.php.
//
// StrEx("av.mp4", ".") -> "mp4"
// ff:取右边
// haystack:文本
// needle:欲寻找的文本
func StrEx(haystack string, needle string) string {
	if s := Str(haystack, needle); s != "" {
		return s[1:]
	}
	return ""
}

// StrTill returns part of `haystack` string ending to and including
// the first occurrence of `needle` from the start of `haystack`.
//
// StrTill("av.mp4", ".") -> "av."
// ff:取左边并含关键字
// haystack:文本
// needle:欲寻找的文本
func StrTill(haystack string, needle string) string {
	pos := strings.Index(haystack, needle)
	if pos == NotFoundIndex || pos == 0 {
		return ""
	}
	return haystack[:pos+1]
}

// StrTillEx returns part of `haystack` string ending to and excluding
// the first occurrence of `needle` from the start of `haystack`.
//
// StrTillEx("av.mp4", ".") -> "av"
// ff:取左边
// haystack:文本
// needle:欲寻找的文本
func StrTillEx(haystack string, needle string) string {
	pos := strings.Index(haystack, needle)
	if pos == NotFoundIndex || pos == 0 {
		return ""
	}
	return haystack[:pos]
}

// SubStr returns a portion of string `str` specified by the `start` and `length` parameters.
// The parameter `length` is optional, it uses the length of `str` in default.
//
// SubStr("123456", 1, 2) -> "23"
// ff:按长度取文本
// str:文本
// start:起始位置
// length:长度
// substr:返回
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
// SubStrRune("一起学习吧！", 2, 2) -> "学习"
// ff:按长度取文本Unicode
// str:文本
// start:起始位置
// length:长度
// substr:返回
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

// StrLimit returns a portion of string `str` specified by `length` parameters, if the length
// of `str` is greater than `length`, then the `suffix` will be appended to the result string.
//
// StrLimit("123456", 3)      -> "123..."
// StrLimit("123456", 3, "~") -> "123~"
// ff:按长度取左边并带前缀
// str:文本
// length:长度
// suffix:后缀
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
// StrLimitRune("一起学习吧！", 2)      -> "一起..."
// StrLimitRune("一起学习吧！", 2, "~") -> "一起~"
// ff:按长度取左边并带前缀Unicode
// str:文本
// length:长度
// suffix:后缀
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

// SubStrFrom returns a portion of string `str` starting from first occurrence of and including `need`
// to the end of `str`.
//
// SubStrFrom("av.mp4", ".") -> ".mp4"
// ff:SubStrFrom别名
// str:
// need:
// substr:
func SubStrFrom(str string, need string) (substr string) {
	pos := Pos(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos:]
}

// SubStrFromEx returns a portion of string `str` starting from first occurrence of and excluding `need`
// to the end of `str`.
//
// SubStrFromEx("av.mp4", ".") -> "mp4"
// ff:SubStrFromEx别名
// str:
// need:
// substr:
func SubStrFromEx(str string, need string) (substr string) {
	pos := Pos(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos+len(need):]
}

// SubStrFromR returns a portion of string `str` starting from last occurrence of and including `need`
// to the end of `str`.
//
// SubStrFromR("/dev/vda", "/") -> "/vda"
// ff:取右边并倒找与含关键字
// str:文本
// need:欲寻找的文本
// substr:文本结果
func SubStrFromR(str string, need string) (substr string) {
	pos := PosR(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos:]
}

// SubStrFromREx returns a portion of string `str` starting from last occurrence of and excluding `need`
// to the end of `str`.
//
// SubStrFromREx("/dev/vda", "/") -> "vda"
// ff:取右边并倒找
// str:文本
// need:欲寻找的文本
// substr:文本结果
func SubStrFromREx(str string, need string) (substr string) {
	pos := PosR(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos+len(need):]
}
