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

package gstr

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

// CaseTypeMatch 从字符串匹配案件类型。 md5:e9d0c49161bc12ae
// ff:命名方式判断
// caseStr:待判断名称
func CaseTypeMatch(caseStr string) CaseType {
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
		if Equal(caseStr, string(caseType)) {
			return caseType
		}
	}

	return CaseType(caseStr)
}

// CaseConvert 将字符串转换为指定的命名约定。
// 使用 CaseTypeMatch 从字符串中匹配 case 类型。
// md5:3c58b688150ee2a3
// ff:命名转换
// s:待转换文本
// caseType:类型
func CaseConvert(s string, caseType CaseType) string {
	if s == "" || caseType == "" {
		return s
	}

	switch caseType {
	case Camel:
		return CaseCamel(s)

	case CamelLower:
		return CaseCamelLower(s)

	case Kebab:
		return CaseKebab(s)

	case KebabScreaming:
		return CaseKebabScreaming(s)

	case Snake:
		return CaseSnake(s)

	case SnakeFirstUpper:
		return CaseSnakeFirstUpper(s)

	case SnakeScreaming:
		return CaseSnakeScreaming(s)

	case Lower:
		return ToLower(s)

	default:
		return s
	}
}

// CaseCamel converts a string to CamelCase.
//
// CaseCamel("any_kind_of_string") -> AnyKindOfString
// ff:命名转换到首字母大写驼峰
// s:待转换文本
func CaseCamel(s string) string {
	return toCamelInitCase(s, true)
}

// CaseCamelLower converts a string to lowerCamelCase.
//
// CaseCamelLower("any_kind_of_string") -> anyKindOfString
// ff:命名转换到首字母小写驼峰
// s:待转换文本
func CaseCamelLower(s string) string {
	if s == "" {
		return s
	}
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}
	return toCamelInitCase(s, false)
}

// CaseSnake converts a string to snake_case.
//
// CaseSnake("AnyKindOfString") -> any_kind_of_string
// ff:命名转换到全小写蛇形
// s:待转换文本
func CaseSnake(s string) string {
	return CaseDelimited(s, '_')
}

// CaseSnakeScreaming converts a string to SNAKE_CASE_SCREAMING.
//
// CaseSnakeScreaming("AnyKindOfString") -> ANY_KIND_OF_STRING
// ff:命名转换到大写蛇形
// s:待转换文本
func CaseSnakeScreaming(s string) string {
	return CaseDelimitedScreaming(s, '_', true)
}

// CaseSnakeFirstUpper converts a string like "RGBCodeMd5" to "rgb_code_md5".
// TODO for efficiency should change regexp to traversing string in future.
//
// CaseSnakeFirstUpper("RGBCodeMd5") -> rgb_code_md5
// ff:命名转换到全小写蛇形2
// word:待转换文本
// underscore:可选连接符
func CaseSnakeFirstUpper(word string, underscore ...string) string {
	replace := "_"
	if len(underscore) > 0 {
		replace = underscore[0]
	}

	m := firstCamelCaseEnd.FindAllStringSubmatch(word, 1)
	if len(m) > 0 {
		word = m[0][1] + replace + TrimLeft(ToLower(m[0][2]), replace)
	}

	for {
		m = firstCamelCaseStart.FindAllStringSubmatch(word, 1)
		if len(m) > 0 && m[0][1] != "" {
			w := strings.ToLower(m[0][1])
			w = w[:len(w)-1] + replace + string(w[len(w)-1])

			word = strings.Replace(word, m[0][1], w, 1)
		} else {
			break
		}
	}

	return TrimLeft(word, replace)
}

// CaseKebab converts a string to kebab-case.
//
// CaseKebab("AnyKindOfString") -> any-kind-of-string
// ff:命名转换到小写短横线
// s:待转换文本
func CaseKebab(s string) string {
	return CaseDelimited(s, '-')
}

// CaseKebabScreaming converts a string to KEBAB-CASE-SCREAMING.
//
// CaseKebab("AnyKindOfString") -> ANY-KIND-OF-STRING
// ff:命名转换到大写驼峰短横线
// s:待转换文本
func CaseKebabScreaming(s string) string {
	return CaseDelimitedScreaming(s, '-', true)
}

// CaseDelimited converts a string to snake.case.delimited.
//
// CaseDelimited("AnyKindOfString", '.') -> any.kind.of.string
// ff:命名转换按符号
// s:待转换文本
// del:连接符号
func CaseDelimited(s string, del byte) string {
	return CaseDelimitedScreaming(s, del, false)
}

// CaseDelimitedScreaming converts a string to DELIMITED.SCREAMING.CASE or delimited.screaming.case.
//
// CaseDelimitedScreaming("AnyKindOfString", '.') -> ANY.KIND.OF.STRING
// ff:命名转换按符号与大小写
// s:待转换文本
// del:连接符号
// screaming:是否全大写
func CaseDelimitedScreaming(s string, del uint8, screaming bool) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	for i, v := range s {
		// 将首字母缩写视为单词，例如 JSONData -> JSON 是一个完整的单词. md5:48305e8e01011121
		nextCaseIsChanged := false
		if i+1 < len(s) {
			next := s[i+1]
			if (v >= 'A' && v <= 'Z' && next >= 'a' && next <= 'z') || (v >= 'a' && v <= 'z' && next >= 'A' && next <= 'Z') {
				nextCaseIsChanged = true
			}
		}

		if i > 0 && n[len(n)-1] != del && nextCaseIsChanged {
			// 如果下一个字母的大小写类型改变，请添加下划线. md5:6409a4f72dd6f8df
			if v >= 'A' && v <= 'Z' {
				n += string(del) + string(v)
			} else if v >= 'a' && v <= 'z' {
				n += string(v) + string(del)
			}
		} else if v == ' ' || v == '_' || v == '-' || v == '.' {
			// 将空格/下划线替换为分隔符. md5:f2c26bc8f3ea056f
			n += string(del)
		} else {
			n = n + string(v)
		}
	}

	if screaming {
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
