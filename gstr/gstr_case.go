// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
//
//   | 函数名称                          | 结果示例             |
//   |-----------------------------------|--------------------|
//   | CaseSnake(s)                      | any_kind_of_string |
//   | CaseSnakeScreaming(s)             | ANY_KIND_OF_STRING |
//   | CaseSnakeFirstUpper("RGBCodeMd5") | rgb_code_md5       |
//   | CaseKebab(s)                      | any-kind-of-string |
//   | CaseKebabScreaming(s)             | ANY-KIND-OF-STRING |
//   | CaseDelimited(s, '.')             | any.kind.of.string |
//   | CaseDelimitedScreaming(s, '.')    | ANY.KIND.OF.STRING |
//   | CaseCamel(s)                      | AnyKindOfString    |
//   | CaseCamelLower(s)                 | anyKindOfString    |
// 注：这些函数提供了对字符串进行不同格式转换的功能，如蛇形命名、尖叫蛇形命名、驼峰命名（首字母大写和小写）以及使用特定分隔符的命名方式。
//
// 2024-01-23 
// 这些方法用于英文变量与方法命名, 基本都用不上. 

package 文本类

import (
	"regexp"
	"strings"
)

// CaseType 是 Case 的类型。
type CaseType string

// 情况类型常量。
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

// CaseTypeMatch 从字符串中匹配案例类型。
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

// CaseConvert 将字符串转换为指定的命名约定。
// 使用 CaseTypeMatch 从字符串中匹配案例类型。
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

// CaseCamel将字符串转换为大驼峰形式(首字母大写)。
// 如: hello world-->HelloWorld
func X命名转换到首字母大写驼峰(待转换文本 string) string {
	return toCamelInitCase(待转换文本, true)
}

// CaseCamelLowe将字符串转换为小驼峰形式(首字母小写)。
// 如: hello world-->helloWorld
func X命名转换到首字母小写驼峰(待转换文本 string) string {
	if 待转换文本 == "" {
		return 待转换文本
	}
	if r := rune(待转换文本[0]); r >= 'A' && r <= 'Z' {
		待转换文本 = strings.ToLower(string(r)) + 待转换文本[1:]
	}
	return toCamelInitCase(待转换文本, false)
}

// CaseSnake将字符串转换中的符号(下划线,空格,点,中横线)用下划线( _ )替换,并全部转换为小写字母。
// 如: hello world-->hello_world
func X命名转换到全小写蛇形(待转换文本 string) string {
	return X命名转换按符号(待转换文本, '_')
}

// CaseSnakeScreaming把字符串中的符号(下划线,空格,点,中横线),全部替换为下划线'_',并将所有英文字母转为大写。
// 如: hello world--> HELLO_WORLD
func X命名转换到大写蛇形(待转换文本 string) string {
	return X命名转换按符号与大小写(待转换文本, '_', true)
}

// CaseSnakeFirstUpper将字符串中的字母为大写时,将大写字母转换为小写字母并在其前面增加一个下划线'_',首字母大写时,只转换为小写,前面不增加下划线'_'
// 如:  RGBCodeMd5-->rgb_code_md5
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

// CaseKebab将字符串转换中的符号(下划线,空格,点,)用中横线'-'替换,并全部转换为小写字母。
// 如:  hello world-->hello-world
func X命名转换到小写短横线(待转换文本 string) string {
	return X命名转换按符号(待转换文本, '-')
}

// CaseKebabScreaming将字符串转换中的符号(下划线,空格,点,中横线)用中横线'-'替换,并全部转换为大写字母。
// 如:  hello world-->HELLO-WORLD
func X命名转换到大写驼峰短横线(待转换文本 string) string {
	return X命名转换按符号与大小写(待转换文本, '-', true)
}

// CaseDelimited将字符串转换中的符号进行替换。
// 如:  
// var (
// str    = `hello world`
// del    = byte('-')
// result = gstr.CaseDelimited(str, del)
// )
// fmt.Println(result) // hello-world
func X命名转换按符号(待转换文本 string, 连接符号 byte) string {
	return X命名转换按符号与大小写(待转换文本, 连接符号, false)
}

// CaseDelimitedScreaming将字符串中的符号(空格,下划线,点,中横线)用第二个参数进行替换,
// 该函数第二个参数为替换的字符,第三个参数为大小写转换,
// true为全部转换大写字母,false为全部转为小写字母。
// 如:
// func ExampleCaseDelimitedScreaming() {
// 	{
// 		var (
// 			str    = `hello world`
// 			del    = byte('-')
// 			result = gstr.CaseDelimitedScreaming(str, del, true)
// 		)
// 		fmt.Println(result)	//  HELLO-WORLD
// 	}
// 	{
// 		var (
// 			str    = `hello world`
// 			del    = byte('-')
// 			result = gstr.CaseDelimitedScreaming(str, del, false)
// 		)
// 		fmt.Println(result)	//  hello-world
// 	}
// }

func X命名转换按符号与大小写(待转换文本 string, 连接符号 uint8, 是否全大写 bool) string {
	待转换文本 = addWordBoundariesToNumbers(待转换文本)
	待转换文本 = strings.Trim(待转换文本, " ")
	n := ""
	for i, v := range 待转换文本 {
		// 将缩写视为单词处理，例如对于 JSONData，JSON 视为一个完整的单词
		nextCaseIsChanged := false
		if i+1 < len(待转换文本) {
			next := 待转换文本[i+1]
			if (v >= 'A' && v <= 'Z' && next >= 'a' && next <= 'z') || (v >= 'a' && v <= 'z' && next >= 'A' && next <= 'Z') {
				nextCaseIsChanged = true
			}
		}

		if i > 0 && n[len(n)-1] != 连接符号 && nextCaseIsChanged {
			// 如果下一个字母的大小写类型发生变化，则添加下划线
			if v >= 'A' && v <= 'Z' {
				n += string(连接符号) + string(v)
			} else if v >= 'a' && v <= 'z' {
				n += string(v) + string(连接符号)
			}
		} else if v == ' ' || v == '_' || v == '-' || v == '.' {
			// 将空格和下划线替换为分隔符
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

// 将字符串转换为驼峰式命名
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
