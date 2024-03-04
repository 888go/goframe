// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gregex 提供了用于正则表达式功能的高性能 API。
package gregex

import (
	"regexp"
)

// Quote通过将`s`中的特殊字符替换为符合正则表达式模式的规则，
// 并返回处理后的副本。
//
// 例如：Quote(`[foo]`) 将返回 `\[foo\]`。
func Quote(s string) string {
	return regexp.QuoteMeta(s)
}

// Validate 检查给定的正则表达式模式 `pattern` 是否有效。
func Validate(pattern string) error {
	_, err := getRegexp(pattern)
	return err
}

// IsMatch 检查给定的字节序列 `src` 是否匹配模式 `pattern`。
func IsMatch(pattern string, src []byte) bool {
	if r, err := getRegexp(pattern); err == nil {
		return r.Match(src)
	}
	return false
}

// IsMatchString 检查给定的字符串 `src` 是否与 `pattern` 匹配。
func IsMatchString(pattern string, src string) bool {
	return IsMatch(pattern, []byte(src))
}

// Match 返回匹配 `pattern` 的字节切片。
func Match(pattern string, src []byte) ([][]byte, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindSubmatch(src), nil
	} else {
		return nil, err
	}
}

// MatchString 返回匹配`pattern`的字符串。
func MatchString(pattern string, src string) ([]string, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindStringSubmatch(src), nil
	} else {
		return nil, err
	}
}

// MatchAll 返回所有匹配 `pattern` 的字节切片。
func MatchAll(pattern string, src []byte) ([][][]byte, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindAllSubmatch(src, -1), nil
	} else {
		return nil, err
	}
}

// MatchAllString 返回所有匹配 `pattern` 的字符串。
func MatchAllString(pattern string, src string) ([][]string, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindAllStringSubmatch(src, -1), nil
	} else {
		return nil, err
	}
}

// Replace 将字节 `src` 中所有匹配到的 `pattern` 替换为字节 `replace`。
func Replace(pattern string, replace, src []byte) ([]byte, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.ReplaceAll(src, replace), nil
	} else {
		return nil, err
	}
}

// ReplaceString 将字符串 `src` 中所有匹配到的 `pattern` 替换为字符串 `replace`。
func ReplaceString(pattern, replace, src string) (string, error) {
	r, e := Replace(pattern, []byte(replace), []byte(src))
	return string(r), e
}

// ReplaceFunc 将字节切片 `src` 中所有匹配到的 `pattern` 用自定义替换函数 `replaceFunc` 进行替换。
func ReplaceFunc(pattern string, src []byte, replaceFunc func(b []byte) []byte) ([]byte, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.ReplaceAllFunc(src, replaceFunc), nil
	} else {
		return nil, err
	}
}

// ReplaceFuncMatch 在字节切片 `src` 中替换所有匹配的 `pattern`，
// 使用自定义替换函数 `replaceFunc` 进行替换。
// `replaceFunc` 参数中的 `match` 类型为 [][]byte，
// 它是使用 Match 函数得到的所有 `pattern` 子模式的结果。
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

// ReplaceStringFunc 在字符串 `src` 中使用自定义替换函数 `replaceFunc` 替换所有匹配到的 `pattern`。
func ReplaceStringFunc(pattern string, src string, replaceFunc func(s string) string) (string, error) {
	bytes, err := ReplaceFunc(pattern, []byte(src), func(bytes []byte) []byte {
		return []byte(replaceFunc(string(bytes)))
	})
	return string(bytes), err
}

// ReplaceStringFuncMatch 在字符串 `src` 中使用自定义替换函数 `replaceFunc` 替换所有匹配到的 `pattern`。
// 参数 `replaceFunc` 的形参类型为 `[]string`，该结果包含通过 MatchString 函数得到的所有 `pattern` 的子模式。
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

// Split 函数将 `src` 切片按照表达式进行分割，并返回由这些表达式匹配之间的子字符串构成的切片。
func Split(pattern string, src string) []string {
	if r, err := getRegexp(pattern); err == nil {
		return r.Split(src, -1)
	}
	return nil
}
