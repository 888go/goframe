// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gregex提供了正则表达式功能的高性能API。 md5:5d43833868579329
package 正则类

import (
	"regexp"
)

// X转义特殊符号 通过替换`s`中的特殊字符，使其符合正则表达式模式的规则，对`s`进行引号包裹。然后返回修改后的字符串。
// 
// 例如：X转义特殊符号(`[foo]`) 返回 `\[foo\]`。
// md5:baab9e0870efe45f
func X转义特殊符号(文本 string) string {
	return regexp.QuoteMeta(文本)
}

// X表达式验证 检查给定的正则表达式模式 `pattern` 是否有效。 md5:39fda51666585abe
func X表达式验证(表达式 string) error {
	_, err := getRegexp(表达式)
	return err
}

// X是否匹配字节集 检查给定的字节`src`是否匹配`pattern`。 md5:7f26688fb33d288d
func X是否匹配字节集(表达式 string, 字节集 []byte) bool {
	if r, err := getRegexp(表达式); err == nil {
		return r.Match(字节集)
	}
	return false
}

// X是否匹配文本 检查给定的字符串 `src` 是否与 `pattern` 匹配。 md5:b4afd86a0688ae19
func X是否匹配文本(表达式 string, 字节集 string) bool {
	return X是否匹配字节集(表达式, []byte(字节集))
}

// X匹配字节集 返回匹配`pattern`的字节切片。 md5:7f3dd939e61e4db8
func X匹配字节集(表达式 string, 字节集 []byte) ([][]byte, error) {
	if r, err := getRegexp(表达式); err == nil {
		return r.FindSubmatch(字节集), nil
	} else {
		return nil, err
	}
}

// X匹配文本 返回与`pattern`匹配的字符串。 md5:d4d623eec4e1f3ec
func X匹配文本(表达式 string, 文本 string) ([]string, error) {
	if r, err := getRegexp(表达式); err == nil {
		return r.FindStringSubmatch(文本), nil
	} else {
		return nil, err
	}
}

// X匹配全部字节集 返回所有匹配`pattern`的字节切片。 md5:64871f15e4916377
func X匹配全部字节集(表达式 string, 字节集 []byte) ([][][]byte, error) {
	if r, err := getRegexp(表达式); err == nil {
		return r.FindAllSubmatch(字节集, -1), nil
	} else {
		return nil, err
	}
}

// X匹配全部文本 返回所有与`pattern`匹配的字符串。 md5:ec34b1802db69c97
func X匹配全部文本(表达式 string, 文本 string) ([][]string, error) {
	if r, err := getRegexp(表达式); err == nil {
		return r.FindAllStringSubmatch(文本, -1), nil
	} else {
		return nil, err
	}
}

// X替换字节集 将 `src` 字节中的所有匹配 `pattern` 的部分替换为 `replace` 字节。 md5:5d5c7ad162f72858
func X替换字节集(表达式 string, 替换字节集, 字节集 []byte) ([]byte, error) {
	if r, err := getRegexp(表达式); err == nil {
		return r.ReplaceAll(字节集, 替换字节集), nil
	} else {
		return nil, err
	}
}

// X替换文本 将字符串 `src` 中所有匹配的 `pattern` 替换为字符串 `replace`。 md5:1e7fdbe12a62e763
func X替换文本(表达式, 替换文本, 文本 string) (string, error) {
	r, e := X替换字节集(表达式, []byte(替换文本), []byte(文本))
	return string(r), e
}

// X替换字节集_函数 使用自定义的替换函数 `replaceFunc`，将字节切片 `src` 中所有匹配的 `pattern` 替换。
// md5:3b66619bd59d4056
func X替换字节集_函数(表达式 string, 字节集 []byte, 回调函数 func(b []byte) []byte) ([]byte, error) {
	if r, err := getRegexp(表达式); err == nil {
		return r.ReplaceAllFunc(字节集, 回调函数), nil
	} else {
		return nil, err
	}
}

// ReplaceFuncMatch：在字节`src`中使用自定义替换函数`replaceFunc`替换所有匹配的`pattern`。`replaceFunc`的参数`match`类型为`[][]byte`，它包含了`pattern`使用Match函数的所有子模式的结果。
// md5:cdbed5cefac02741
func ReplaceFuncMatch(pattern string, src []byte, replaceFunc func(match [][]byte) []byte) ([]byte, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.ReplaceAllFunc(src, func(bytes []byte) []byte {
			match, _ := X匹配字节集(pattern, bytes)
			return replaceFunc(match)
		}), nil
	} else {
		return nil, err
	}
}

// X替换文本_函数 函数会在字符串 `src` 中替换所有匹配的 `pattern`，使用自定义的替换函数 `replaceFunc`。
// md5:8575760795474682
func X替换文本_函数(表达式 string, 文本 string, 回调函数 func(s string) string) (string, error) {
	bytes, err := X替换字节集_函数(表达式, []byte(文本), func(bytes []byte) []byte {
		return []byte(回调函数(string(bytes)))
	})
	return string(bytes), err
}

// ReplaceStringFuncMatch 将字符串 `src` 中所有与 `pattern` 匹配的部分
// 使用自定义替换函数 `replaceFunc` 进行替换。
// 替换函数 `replaceFunc` 的参数 `match` 类型为 []string，
// 它包含使用 MatchString 函数得到的 `pattern` 中所有子模式匹配结果。
// md5:b24f208b16cfd56a
func ReplaceStringFuncMatch(pattern string, src string, replaceFunc func(match []string) string) (string, error) {
	if r, err := getRegexp(pattern); err == nil {
		return string(r.ReplaceAllFunc([]byte(src), func(bytes []byte) []byte {
			match, _ := X匹配文本(pattern, string(bytes))
			return []byte(replaceFunc(match))
		})), nil
	} else {
		return "", err
	}
}

// X分割 将切片 `src` 按照给定的表达式分割成多个子字符串，并返回这些匹配之间的子字符串切片。
// md5:e0809df699cf82c1
func X分割(表达式 string, 文本 string) []string {
	if r, err := getRegexp(表达式); err == nil {
		return r.Split(文本, -1)
	}
	return nil
}
