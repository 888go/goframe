// 版权所有GoFrame作者(https://goframe.org)，保留所有权利。
//
// 本源代码形式受MIT许可条款约束。若未随此文件分发MIT许可的副本，
// 您可以从https://github.com/gogf/gf获取。
//
//   | 功能                             | 结果               |
//   |-----------------------------------|---------------------|
//   | CaseSnake(s)                       | 任意类型的字符串   |
//   | CaseSnakeScreaming(s)              | 任意类型的大写字符串 |
//   | CaseSnakeFirstUpper("RGBCodeMd5")   | rgb_code_md5        |
//   | CaseKebab(s)                       | 任意-类型-的字符串  |
//   | CaseKebabScreaming(s)              | 任意-类型-的大写字符串 |
//   | CaseDelimited(s, '.')              | 任何.类型.的字符串  |
//   | CaseDelimitedScreaming(s, '.')     | 任何.类型.的大写字符串 |
//   | CaseCamel(s)                       | AnyKindOfString     |
//   | CaseCamelLower(s)                  | anyKindOfString     |
// md5:66511e05f3151030

package 文本类

import (
	"regexp"
	"strings"
)

// CaseType 是 Case 的类型。 md5:3bd1a46ccb6a5474
type CaseType string

// 案例类型常量。 md5:09b430778f19740d
const (
	Camel           CaseType = "Camel"
	CamelLower      CaseType = "CamelLower"
	Snake           CaseType = "Snake"
	SnakeFirstUpper CaseType = "SnakeFirstUpper"
	SnakeScreaming  CaseType = "SnakeScreaming"
	Kebab           CaseType = "Kebab"
	KebabScreaming  CaseType = "KebabScreaming"
	Lower           CaseType = "Lower"
)

var (
	numberSequence      = regexp.MustCompile(`([a-zA-Z]{0,1})(\d+)([a-zA-Z]{0,1})`)
	firstCamelCaseStart = regexp.MustCompile(`([A-Z]+)([A-Z]?[_a-z\d]+)|$`)
	firstCamelCaseEnd   = regexp.MustCompile(`([\w\W]*?)([_]?[A-Z]+)$`)
)

// X命名方式判断 从字符串匹配案件类型。 md5:e9d0c49161bc12ae
func X命名方式判断(待判断名称 string) CaseType {
	caseTypes := []CaseType{
		Camel,
		CamelLower,
		Snake,
		SnakeFirstUpper,
		SnakeScreaming,
		Kebab,
		KebabScreaming,
		Lower,
	}

	for _, caseType := range caseTypes {
		if X相等比较并忽略大小写(待判断名称, string(caseType)) {
			return caseType
		}
	}

	return CaseType(待判断名称)
}

// X命名转换 将字符串转换为指定的命名约定。
// 使用 CaseTypeMatch 从字符串中匹配 case 类型。
// md5:3c58b688150ee2a3
func X命名转换(待转换文本 string, 类型 CaseType) string {
	if 待转换文本 == "" || 类型 == "" {
		return 待转换文本
	}

	switch 类型 {
	case Camel:
		return X命名转换到首字母大写驼峰(待转换文本)

	case CamelLower:
		return X命名转换到首字母小写驼峰(待转换文本)

	case Kebab:
		return X命名转换到小写短横线(待转换文本)

	case KebabScreaming:
		return X命名转换到大写驼峰短横线(待转换文本)

	case Snake:
		return X命名转换到全小写蛇形(待转换文本)

	case SnakeFirstUpper:
		return X命名转换到全小写蛇形2(待转换文本)

	case SnakeScreaming:
		return X命名转换到大写蛇形(待转换文本)

	case Lower:
		return X到小写(待转换文本)

	default:
		return 待转换文本
	}
}

// X命名转换到首字母大写驼峰 将一个字符串转换为驼峰式写法。
//
// 示例：
// X命名转换到首字母大写驼峰("any_kind_of_string") -> AnyKindOfString
// md5:189cc8dcd6a04d2c
func X命名转换到首字母大写驼峰(待转换文本 string) string {
	return toCamelInitCase(待转换文本, true)
}

// X命名转换到首字母小写驼峰 将一个字符串转换为下划线驼峰式（lowerCamelCase）。
//
// 例子：
// X命名转换到首字母小写驼峰("any_kind_of_string") -> anyKindOfString
// md5:dc604c858a2452d4
func X命名转换到首字母小写驼峰(待转换文本 string) string {
	if 待转换文本 == "" {
		return 待转换文本
	}
	if r := rune(待转换文本[0]); r >= 'A' && r <= 'Z' {
		待转换文本 = strings.ToLower(string(r)) + 待转换文本[1:]
	}
	return toCamelInitCase(待转换文本, false)
}

// X命名转换到全小写蛇形将一个字符串转换为蛇形命名（snake_case）。
//
// 示例：
// X命名转换到全小写蛇形("AnyKindOfString") -> any_kind_of_string
// md5:348ee5cd8cb1cd34
func X命名转换到全小写蛇形(待转换文本 string) string {
	return X命名转换按符号(待转换文本, '_')
}

// X命名转换到大写蛇形 将一个字符串转换为 SNAKE_CASE_SCREAMING 格式。
//
// 示例：
// X命名转换到大写蛇形("AnyKindOfString") -> "ANY_KIND_OF_STRING"
// md5:9f2e1f082921e42e
func X命名转换到大写蛇形(待转换文本 string) string {
	return X命名转换按符号与大小写(待转换文本, '_', true)
}

// X命名转换到全小写蛇形2 将类似 "RGBCodeMd5" 的字符串转换为 "rgb_code_md5"。
// TODO 为了提高效率，未来应将正则表达式改为遍历字符串的方式。
//
// 示例：
// X命名转换到全小写蛇形2("RGBCodeMd5") -> rgb_code_md5
// md5:aff36f9f5f3a68d7
func X命名转换到全小写蛇形2(待转换文本 string, 可选连接符 ...string) string {
	replace := "_"
	if len(可选连接符) > 0 {
		replace = 可选连接符[0]
	}

	m := firstCamelCaseEnd.FindAllStringSubmatch(待转换文本, 1)
	if len(m) > 0 {
		待转换文本 = m[0][1] + replace + X过滤首字符并含空白(X到小写(m[0][2]), replace)
	}

	for {
		m = firstCamelCaseStart.FindAllStringSubmatch(待转换文本, 1)
		if len(m) > 0 && m[0][1] != "" {
			w := strings.ToLower(m[0][1])
			w = w[:len(w)-1] + replace + string(w[len(w)-1])

			待转换文本 = strings.Replace(待转换文本, m[0][1], w, 1)
		} else {
			break
		}
	}

	return X过滤首字符并含空白(待转换文本, replace)
}

// X命名转换到小写短横线 将字符串转换为kebab-case形式。
//
// 例子：
// X命名转换到小写短横线("AnyKindOfString") -> any-kind-of-string
// md5:885475f21356c510
func X命名转换到小写短横线(待转换文本 string) string {
	return X命名转换按符号(待转换文本, '-')
}

// X命名转换到大写驼峰短横线 将一个字符串转换为KEBAB-CASE-SCREAMING格式。
//
// 示例：
// CaseKebab("AnyKindOfString") -> "ANY-KIND-OF-STRING"
// md5:64e3399ff1b60dad
func X命名转换到大写驼峰短横线(待转换文本 string) string {
	return X命名转换按符号与大小写(待转换文本, '-', true)
}

// X命名转换按符号 将字符串转换为 snake_case_delimited 形式。
//
// 示例：
// X命名转换按符号("AnyKindOfString", '.') -> any.kind.of.string
// md5:8edd65912cb80360
func X命名转换按符号(待转换文本 string, 连接符号 byte) string {
	return X命名转换按符号与大小写(待转换文本, 连接符号, false)
}

// X命名转换按符号与大小写 将字符串转换为 DELIMITED.SCREAMING.CASE 或 delimited.screaming.case 格式。
//
// 示例：
// X命名转换按符号与大小写("AnyKindOfString", '.') -> ANY.KIND.OF.STRING
// md5:e81c17d2e4a95231
func X命名转换按符号与大小写(待转换文本 string, 连接符号 uint8, 是否全大写 bool) string {
	待转换文本 = addWordBoundariesToNumbers(待转换文本)
	待转换文本 = strings.Trim(待转换文本, " ")
	n := ""
	for i, v := range 待转换文本 {
				// 将首字母缩写视为单词，例如 JSONData -> JSON 是一个完整的单词. md5:48305e8e01011121
		nextCaseIsChanged := false
		if i+1 < len(待转换文本) {
			next := 待转换文本[i+1]
			if (v >= 'A' && v <= 'Z' && next >= 'a' && next <= 'z') || (v >= 'a' && v <= 'z' && next >= 'A' && next <= 'Z') {
				nextCaseIsChanged = true
			}
		}

		if i > 0 && n[len(n)-1] != 连接符号 && nextCaseIsChanged {
									// 如果下一个字母的大小写类型改变，请添加下划线. md5:6409a4f72dd6f8df
			if v >= 'A' && v <= 'Z' {
				n += string(连接符号) + string(v)
			} else if v >= 'a' && v <= 'z' {
				n += string(v) + string(连接符号)
			}
		} else if v == ' ' || v == '_' || v == '-' || v == '.' {
						// 将空格/下划线替换为分隔符. md5:f2c26bc8f3ea056f
			n += string(连接符号)
		} else {
			n = n + string(v)
		}
	}

	if 是否全大写 {
		n = strings.ToUpper(n)
	} else {
		n = strings.ToLower(n)
	}
	return n
}

func addWordBoundariesToNumbers(s string) string {
	r := numberSequence.ReplaceAllFunc([]byte(s), func(bytes []byte) []byte {
		var result []byte
		match := numberSequence.FindSubmatch(bytes)
		if len(match[1]) > 0 {
			result = append(result, match[1]...)
			result = append(result, []byte(" ")...)
		}
		result = append(result, match[2]...)
		if len(match[3]) > 0 {
			result = append(result, []byte(" ")...)
			result = append(result, match[3]...)
		}
		return result
	})
	return string(r)
}

// 将字符串转换为驼峰式命名. md5:e7c9de8ba3801cd9
func toCamelInitCase(s string, initCase bool) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	capNext := initCase
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' || v == '.' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}
