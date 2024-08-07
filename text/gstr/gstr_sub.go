// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文本类

import (
	"strings"
)

// X取右边并含关键字 函数返回从 `haystack` 字符串开始，包括第一个出现的 `needle` 到 `haystack` 结尾的部分。
//
// 这个函数的行为与 SubStr 函数完全相同，但是为了实现与 PHP 相同的功能：http://php.net/manual/en/function.strstr.php。
//
// 示例：
// X取右边并含关键字("av.mp4", ".") -> ".mp4"
// md5:35b0375920cdf357
func X取右边并含关键字(文本 string, 欲寻找的文本 string) string {
	if 欲寻找的文本 == "" {
		return ""
	}
	pos := strings.Index(文本, 欲寻找的文本)
	if pos == NotFoundIndex {
		return ""
	}
	return 文本[pos+len([]byte(欲寻找的文本))-1:]
}

// X取右边 从`haystack`字符串中返回从第一个出现的`needle`开始到`haystack`末尾的部分。
//
// 这个函数的行为与SubStrEx函数完全相同，但实现了与PHP相同的函数：http://php.net/manual/en/function.strstr.php。
//
// 示例：
// X取右边("av.mp4", ".") -> "mp4"
// md5:1f6c13b098caa33a
func X取右边(文本 string, 欲寻找的文本 string) string {
	if s := X取右边并含关键字(文本, 欲寻找的文本); s != "" {
		return s[1:]
	}
	return ""
}

// X取左边并含关键字 返回 `haystack` 字符串中从开头到（包括）第一个 `needle` 出现的部分。
//
// 示例：
// X取左边并含关键字("av.mp4", ".") -> "av."
// md5:f914c156abb95437
func X取左边并含关键字(文本 string, 欲寻找的文本 string) string {
	pos := strings.Index(文本, 欲寻找的文本)
	if pos == NotFoundIndex || pos == 0 {
		return ""
	}
	return 文本[:pos+1]
}

// X取左边 函数返回 `haystack` 字符串中从开始到（但不包括）第一次出现 `needle` 的部分。
//
// 示例：
// X取左边("av.mp4", ".") -> "av"
// md5:c0848291d8036d82
func X取左边(文本 string, 欲寻找的文本 string) string {
	pos := strings.Index(文本, 欲寻找的文本)
	if pos == NotFoundIndex || pos == 0 {
		return ""
	}
	return 文本[:pos]
}

// X按长度取文本 函数返回字符串 `str` 中由 `start` 和 `length` 参数指定的部分。参数 `length` 是可选的，如果未提供，则默认使用 `str` 的长度。
// 
// 示例：
// X按长度取文本("123456", 1, 2) -> "23"
// md5:b6da71b3534fdbbc
func X按长度取文本(文本 string, 起始位置 int, 长度 ...int) (返回 string) {
	strLength := len(文本)
	if 起始位置 < 0 {
		if -起始位置 > strLength {
			起始位置 = 0
		} else {
			起始位置 = strLength + 起始位置
		}
	} else if 起始位置 > strLength {
		return ""
	}
	realLength := 0
	if len(长度) > 0 {
		realLength = 长度[0]
		if realLength < 0 {
			if -realLength > strLength-起始位置 {
				realLength = 0
			} else {
				realLength = strLength - 起始位置 + realLength
			}
		} else if realLength > strLength-起始位置 {
			realLength = strLength - 起始位置
		}
	} else {
		realLength = strLength - 起始位置
	}

	if realLength == strLength {
		return 文本
	} else {
		end := 起始位置 + realLength
		return 文本[起始位置:end]
	}
}

// X按长度取文本Unicode returns a portion of string `str` specified by the `start` and `length` parameters.
// X按长度取文本Unicode considers parameter `str` as unicode string.
// The parameter `length` is optional, it uses the length of `str` in default.
//
// Example:
// X按长度取文本Unicode("一起学习吧！", 2, 2) -> "学习"
func X按长度取文本Unicode(文本 string, 起始位置 int, 长度 ...int) (返回 string) {
		// 转换为[]rune以支持Unicode。 md5:459540c13f4e5603
	var (
		runes       = []rune(文本)
		runesLength = len(runes)
	)

	strLength := runesLength
	if 起始位置 < 0 {
		if -起始位置 > strLength {
			起始位置 = 0
		} else {
			起始位置 = strLength + 起始位置
		}
	} else if 起始位置 > strLength {
		return ""
	}
	realLength := 0
	if len(长度) > 0 {
		realLength = 长度[0]
		if realLength < 0 {
			if -realLength > strLength-起始位置 {
				realLength = 0
			} else {
				realLength = strLength - 起始位置 + realLength
			}
		} else if realLength > strLength-起始位置 {
			realLength = strLength - 起始位置
		}
	} else {
		realLength = strLength - 起始位置
	}
	end := 起始位置 + realLength
	if end > runesLength {
		end = runesLength
	}
	return string(runes[起始位置:end])
}

// X按长度取左边并带前缀 函数返回字符串 `str` 中由 `length` 参数指定长度的部分。如果 `str` 的长度大于 `length`，则 `suffix` 将被添加到结果字符串的末尾。
//
// 示例：
// X按长度取左边并带前缀("123456", 3)      -> "123..."
// X按长度取左边并带前缀("123456", 3, "~") -> "123~"
// md5:bd8f96405a5594b5
func X按长度取左边并带前缀(文本 string, 长度 int, 后缀 ...string) string {
	if len(文本) < 长度 {
		return 文本
	}
	suffixStr := defaultSuffixForStrLimit
	if len(后缀) > 0 {
		suffixStr = 后缀[0]
	}
	return 文本[0:长度] + suffixStr
}

// X按长度取左边并带前缀Unicode returns a portion of string `str` specified by `length` parameters, if the length
// of `str` is greater than `length`, then the `suffix` will be appended to the result string.
// X按长度取左边并带前缀Unicode considers parameter `str` as unicode string.
//
// Example:
// X按长度取左边并带前缀Unicode("一起学习吧！", 2)      -> "一起..."
// X按长度取左边并带前缀Unicode("一起学习吧！", 2, "~") -> "一起~"
func X按长度取左边并带前缀Unicode(文本 string, 长度 int, 后缀 ...string) string {
	runes := []rune(文本)
	if len(runes) < 长度 {
		return 文本
	}
	suffixStr := defaultSuffixForStrLimit
	if len(后缀) > 0 {
		suffixStr = 后缀[0]
	}
	return string(runes[0:长度]) + suffixStr
}

// SubStrFrom别名 从字符串 `str` 中从第一次出现 `need` 的位置开始，包括 `need` 到字符串末尾的部分。
//
// 示例：
// SubStrFrom别名("av.mp4", ".") -> ".mp4"
// md5:f4bff02c473abeff
func SubStrFrom别名(str string, need string) (substr string) {
	pos := X查找(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos:]
}

// SubStrFromEx别名 从字符串 `str` 中返回从首次出现 `need` 到 `str` 结尾的部分（不包括 `need`）。
//
// 示例：
// SubStrFromEx别名("av.mp4", ".") -> "mp4"
// md5:88a817f03fc77455
func SubStrFromEx别名(str string, need string) (substr string) {
	pos := X查找(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos+len(need):]
}

// X取右边并倒找与含关键字 从字符串 `str` 的最后一个出现的 `need` 开始并包括在内，返回一个子串。
// 示例：
// X取右边并倒找与含关键字("/dev/vda", "/") -> "/vda"
// md5:8f70ecc84d0338f8
func X取右边并倒找与含关键字(文本 string, 欲寻找的文本 string) (文本结果 string) {
	pos := X倒找(文本, 欲寻找的文本)
	if pos < 0 {
		return ""
	}
	return 文本[pos:]
}

// X取右边并倒找 从字符串 `str` 中最后一个出现的 `need` 子串之后的字符开始，直到 `str` 的末尾，返回这一部分子串。
//
// 示例：
// X取右边并倒找("/dev/vda", "/") -> "vda"
// md5:3de495ad97b12196
func X取右边并倒找(文本 string, 欲寻找的文本 string) (文本结果 string) {
	pos := X倒找(文本, 欲寻找的文本)
	if pos < 0 {
		return ""
	}
	return 文本[pos+len(欲寻找的文本):]
}
