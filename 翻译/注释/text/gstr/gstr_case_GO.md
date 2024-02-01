
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
//   | Function                          | Result             |
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
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// CaseType is the type for Case.
<原文结束>

# <翻译开始>
// CaseType 是 Case 的类型。
# <翻译结束>







<原文开始>
// CaseTypeMatch matches the case type from string.
<原文结束>

# <翻译开始>
// CaseTypeMatch 从字符串中匹配案例类型。
# <翻译结束>


<原文开始>
// CaseConvert converts a string to the specified naming convention.
// Use CaseTypeMatch to match the case type from string.
<原文结束>

# <翻译开始>
// CaseConvert 将字符串转换为指定的命名约定。
// 使用 CaseTypeMatch 从字符串中匹配案例类型。
# <翻译结束>


<原文开始>
// CaseCamel converts a string to CamelCase.
<原文结束>

# <翻译开始>
// CaseCamel将字符串转换为大驼峰形式(首字母大写)。
// 如: hello world-->HelloWorld
# <翻译结束>


<原文开始>
// CaseCamelLower converts a string to lowerCamelCase.
<原文结束>

# <翻译开始>
// CaseCamelLowe将字符串转换为小驼峰形式(首字母小写)。
// 如: hello world-->helloWorld
# <翻译结束>


<原文开始>
// CaseSnake converts a string to snake_case.
<原文结束>

# <翻译开始>
// CaseSnake将字符串转换中的符号(下划线,空格,点,中横线)用下划线( _ )替换,并全部转换为小写字母。
// 如: hello world-->hello_world
# <翻译结束>


<原文开始>
// CaseSnakeScreaming converts a string to SNAKE_CASE_SCREAMING.
<原文结束>

# <翻译开始>
// CaseSnakeScreaming把字符串中的符号(下划线,空格,点,中横线),全部替换为下划线'_',并将所有英文字母转为大写。
// 如: hello world--> HELLO_WORLD
# <翻译结束>


<原文开始>
// CaseSnakeFirstUpper converts a string like "RGBCodeMd5" to "rgb_code_md5".
// TODO for efficiency should change regexp to traversing string in future.
<原文结束>

# <翻译开始>
// CaseSnakeFirstUpper将字符串中的字母为大写时,将大写字母转换为小写字母并在其前面增加一个下划线'_',首字母大写时,只转换为小写,前面不增加下划线'_'
// 如:  RGBCodeMd5-->rgb_code_md5
# <翻译结束>


<原文开始>
// CaseKebab converts a string to kebab-case
<原文结束>

# <翻译开始>
// CaseKebab将字符串转换中的符号(下划线,空格,点,)用中横线'-'替换,并全部转换为小写字母。
// 如:  hello world-->hello-world
# <翻译结束>


<原文开始>
// CaseKebabScreaming converts a string to KEBAB-CASE-SCREAMING.
<原文结束>

# <翻译开始>
// CaseKebabScreaming将字符串转换中的符号(下划线,空格,点,中横线)用中横线'-'替换,并全部转换为大写字母。
// 如:  hello world-->HELLO-WORLD
# <翻译结束>


<原文开始>
// CaseDelimited converts a string to snake.case.delimited.
<原文结束>

# <翻译开始>
// CaseDelimited将字符串转换中的符号进行替换。
// 如:  
// var (
// str    = `hello world`
// del    = byte('-')
// result = gstr.CaseDelimited(str, del)
// )
// fmt.Println(result) // hello-world
# <翻译结束>


<原文开始>
// CaseDelimitedScreaming converts a string to DELIMITED.SCREAMING.CASE or delimited.screaming.case.
<原文结束>

# <翻译开始>
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

# <翻译结束>


<原文开始>
// treat acronyms as words, eg for JSONData -> JSON is a whole word
<原文结束>

# <翻译开始>
// 将缩写视为单词处理，例如对于 JSONData，JSON 视为一个完整的单词
# <翻译结束>


<原文开始>
// add underscore if next letter case type is changed
<原文结束>

# <翻译开始>
// 如果下一个字母的大小写类型发生变化，则添加下划线
# <翻译结束>


<原文开始>
// replace spaces/underscores with delimiters
<原文结束>

# <翻译开始>
// 将空格和下划线替换为分隔符
# <翻译结束>


<原文开始>
// Converts a string to CamelCase
<原文结束>

# <翻译开始>
// 将字符串转换为驼峰式命名
# <翻译结束>


<原文开始>
// The case type constants.
<原文结束>

# <翻译开始>
// 情况类型常量。
# <翻译结束>

