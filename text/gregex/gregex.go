// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gregex提供了正则表达式功能的高性能API。 md5:5d43833868579329
package gregex//bm:正则类

import (
	"regexp"
)

// Quote quotes `s` by replacing special chars in `s`
// to match the rules of regular expression pattern.
// And returns the copy.
//
// ff:转义特殊符号
// s:文本
func Quote(s string) string {
	return regexp.QuoteMeta(s)
}

// Validate 检查给定的正则表达式模式 `pattern` 是否有效。 md5:39fda51666585abe
// ff:表达式验证
// pattern:表达式
func Validate(pattern string) error {
	_, err := getRegexp(pattern)
	return err
}

// IsMatch 检查给定的字节`src`是否匹配`pattern`。 md5:7f26688fb33d288d
// ff:是否匹配字节集
// pattern:表达式
// src:字节集
func IsMatch(pattern string, src []byte) bool {
	if r, err := getRegexp(pattern); err == nil {
		return r.Match(src)
	}
	return false
}

// IsMatchString 检查给定的字符串 `src` 是否与 `pattern` 匹配。 md5:b4afd86a0688ae19
// ff:是否匹配文本
// pattern:表达式
// src:字节集
func IsMatchString(pattern string, src string) bool {
	return IsMatch(pattern, []byte(src))
}

// Match 返回匹配`pattern`的字节切片。 md5:7f3dd939e61e4db8
// ff:匹配字节集
// pattern:表达式
// src:字节集
func Match(pattern string, src []byte) ([][]byte, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindSubmatch(src), nil
	} else {
		return nil, err
	}
}

// MatchString 返回与`pattern`匹配的字符串。 md5:d4d623eec4e1f3ec
// ff:匹配文本
// pattern:表达式
// src:文本
func MatchString(pattern string, src string) ([]string, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindStringSubmatch(src), nil
	} else {
		return nil, err
	}
}

// MatchAll 返回所有匹配`pattern`的字节切片。 md5:64871f15e4916377
// ff:匹配全部字节集
// pattern:表达式
// src:字节集
func MatchAll(pattern string, src []byte) ([][][]byte, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindAllSubmatch(src, -1), nil
	} else {
		return nil, err
	}
}

// MatchAllString 返回所有与`pattern`匹配的字符串。 md5:ec34b1802db69c97
// ff:匹配全部文本
// pattern:表达式
// src:文本
func MatchAllString(pattern string, src string) ([][]string, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindAllStringSubmatch(src, -1), nil
	} else {
		return nil, err
	}
}

// Replace 将 `src` 字节中的所有匹配 `pattern` 的部分替换为 `replace` 字节。 md5:5d5c7ad162f72858
// ff:替换字节集
// pattern:表达式
// replace:替换字节集
// src:字节集
func Replace(pattern string, replace, src []byte) ([]byte, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.ReplaceAll(src, replace), nil
	} else {
		return nil, err
	}
}

// ReplaceString 将字符串 `src` 中所有匹配的 `pattern` 替换为字符串 `replace`。 md5:1e7fdbe12a62e763
// ff:替换文本
// pattern:表达式
// replace:替换文本
// src:文本
func ReplaceString(pattern, replace, src string) (string, error) {
	r, e := Replace(pattern, []byte(replace), []byte(src))
	return string(r), e
}

// ReplaceFunc 使用自定义的替换函数 `replaceFunc`，将字节切片 `src` 中所有匹配的 `pattern` 替换。
// md5:3b66619bd59d4056
// ff:替换字节集_函数
// pattern:表达式
// src:字节集
// replaceFunc:回调函数
// b:
func ReplaceFunc(pattern string, src []byte, replaceFunc func(b []byte) []byte) ([]byte, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.ReplaceAllFunc(src, replaceFunc), nil
	} else {
		return nil, err
	}
}

// ReplaceFuncMatch：在字节`src`中使用自定义替换函数`replaceFunc`替换所有匹配的`pattern`。`replaceFunc`的参数`match`类型为`[][]byte`，它包含了`pattern`使用Match函数的所有子模式的结果。
// md5:cdbed5cefac02741
// ff:
// pattern:
// src:
// replaceFunc:
// match:
func ReplaceFuncMatch(pattern string, src []byte, replaceFunc func(match [][]byte) []byte) ([]byte, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.ReplaceAllFunc(src, func(bytes []byte) []byte {
			match, _ := Match(pattern, bytes)
			return replaceFunc(match)
		}), nil
	} else {
		return nil, err
	}
}

// ReplaceStringFunc 函数会在字符串 `src` 中替换所有匹配的 `pattern`，使用自定义的替换函数 `replaceFunc`。
// md5:8575760795474682
// ff:替换文本_函数
// pattern:表达式
// src:文本
// replaceFunc:回调函数
// s:
func ReplaceStringFunc(pattern string, src string, replaceFunc func(s string) string) (string, error) {
	bytes, err := ReplaceFunc(pattern, []byte(src), func(bytes []byte) []byte {
		return []byte(replaceFunc(string(bytes)))
	})
	return string(bytes), err
}

// ReplaceStringFuncMatch 将字符串 `src` 中所有与 `pattern` 匹配的部分
// 使用自定义替换函数 `replaceFunc` 进行替换。
// 替换函数 `replaceFunc` 的参数 `match` 类型为 []string，
// 它包含使用 MatchString 函数得到的 `pattern` 中所有子模式匹配结果。
// md5:b24f208b16cfd56a
// ff:
// pattern:
// src:
// replaceFunc:
// match:
func ReplaceStringFuncMatch(pattern string, src string, replaceFunc func(match []string) string) (string, error) {
	if r, err := getRegexp(pattern); err == nil {
		return string(r.ReplaceAllFunc([]byte(src), func(bytes []byte) []byte {
			match, _ := MatchString(pattern, string(bytes))
			return []byte(replaceFunc(match))
		})), nil
	} else {
		return "", err
	}
}

	// Split 将切片 `src` 按照给定的表达式分割成多个子字符串，并返回这些匹配之间的子字符串切片。
	// md5:e0809df699cf82c1
// ff:分割
// pattern:表达式
// src:文本
func Split(pattern string, src string) []string {
	if r, err := getRegexp(pattern); err == nil {
		return r.Split(src, -1)
	}
	return nil
}
