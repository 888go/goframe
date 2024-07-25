
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
# <翻译结束>


<原文开始>
// CaseType is the type for Case.
<原文结束>

# <翻译开始>
// CaseType 是 Case 的类型。 md5:3bd1a46ccb6a5474
# <翻译结束>


<原文开始>
// The case type constants.
<原文结束>

# <翻译开始>
// 案例类型常量。 md5:09b430778f19740d
# <翻译结束>


<原文开始>
// CaseTypeMatch matches the case type from string.
<原文结束>

# <翻译开始>
// CaseTypeMatch 从字符串匹配案件类型。 md5:e9d0c49161bc12ae
# <翻译结束>


<原文开始>
// CaseConvert converts a string to the specified naming convention.
// Use CaseTypeMatch to match the case type from string.
<原文结束>

# <翻译开始>
// CaseConvert 将字符串转换为指定的命名约定。
// 使用 CaseTypeMatch 从字符串中匹配 case 类型。
// md5:3c58b688150ee2a3
# <翻译结束>


<原文开始>
// CaseCamel converts a string to CamelCase.
//
// Example:
// CaseCamel("any_kind_of_string") -> AnyKindOfString
<原文结束>

# <翻译开始>
// CaseCamel 将一个字符串转换为驼峰式写法。
//
// 示例：
// CaseCamel("any_kind_of_string") -> AnyKindOfString
// md5:189cc8dcd6a04d2c
# <翻译结束>


<原文开始>
// CaseCamelLower converts a string to lowerCamelCase.
//
// Example:
// CaseCamelLower("any_kind_of_string") -> anyKindOfString
<原文结束>

# <翻译开始>
// CaseCamelLower 将一个字符串转换为下划线驼峰式（lowerCamelCase）。
//
// 例子：
// CaseCamelLower("any_kind_of_string") -> anyKindOfString
// md5:dc604c858a2452d4
# <翻译结束>


<原文开始>
// CaseSnake converts a string to snake_case.
//
// Example:
// CaseSnake("AnyKindOfString") -> any_kind_of_string
<原文结束>

# <翻译开始>
// CaseSnake将一个字符串转换为蛇形命名（snake_case）。
//
// 示例：
// CaseSnake("AnyKindOfString") -> any_kind_of_string
// md5:348ee5cd8cb1cd34
# <翻译结束>


<原文开始>
// CaseSnakeScreaming converts a string to SNAKE_CASE_SCREAMING.
//
// Example:
// CaseSnakeScreaming("AnyKindOfString") -> ANY_KIND_OF_STRING
<原文结束>

# <翻译开始>
// CaseSnakeScreaming 将一个字符串转换为 SNAKE_CASE_SCREAMING 格式。
//
// 示例：
// CaseSnakeScreaming("AnyKindOfString") -> "ANY_KIND_OF_STRING"
// md5:9f2e1f082921e42e
# <翻译结束>


<原文开始>
// CaseSnakeFirstUpper converts a string like "RGBCodeMd5" to "rgb_code_md5".
// TODO for efficiency should change regexp to traversing string in future.
//
// Example:
// CaseSnakeFirstUpper("RGBCodeMd5") -> rgb_code_md5
<原文结束>

# <翻译开始>
// CaseSnakeFirstUpper 将类似 "RGBCodeMd5" 的字符串转换为 "rgb_code_md5"。
// TODO 为了提高效率，未来应将正则表达式改为遍历字符串的方式。
//
// 示例：
// CaseSnakeFirstUpper("RGBCodeMd5") -> rgb_code_md5
// md5:aff36f9f5f3a68d7
# <翻译结束>


<原文开始>
// CaseKebab converts a string to kebab-case.
//
// Example:
// CaseKebab("AnyKindOfString") -> any-kind-of-string
<原文结束>

# <翻译开始>
// CaseKebab 将字符串转换为kebab-case形式。
//
// 例子：
// CaseKebab("AnyKindOfString") -> any-kind-of-string
// md5:885475f21356c510
# <翻译结束>


<原文开始>
// CaseKebabScreaming converts a string to KEBAB-CASE-SCREAMING.
//
// Example:
// CaseKebab("AnyKindOfString") -> ANY-KIND-OF-STRING
<原文结束>

# <翻译开始>
// CaseKebabScreaming 将一个字符串转换为KEBAB-CASE-SCREAMING格式。
//
// 示例：
// CaseKebab("AnyKindOfString") -> "ANY-KIND-OF-STRING"
// md5:64e3399ff1b60dad
# <翻译结束>


<原文开始>
// CaseDelimited converts a string to snake.case.delimited.
//
// Example:
// CaseDelimited("AnyKindOfString", '.') -> any.kind.of.string
<原文结束>

# <翻译开始>
// CaseDelimited 将字符串转换为 snake_case_delimited 形式。
//
// 示例：
// CaseDelimited("AnyKindOfString", '.') -> any.kind.of.string
// md5:8edd65912cb80360
# <翻译结束>


<原文开始>
// CaseDelimitedScreaming converts a string to DELIMITED.SCREAMING.CASE or delimited.screaming.case.
//
// Example:
// CaseDelimitedScreaming("AnyKindOfString", '.') -> ANY.KIND.OF.STRING
<原文结束>

# <翻译开始>
// CaseDelimitedScreaming 将字符串转换为 DELIMITED.SCREAMING.CASE 或 delimited.screaming.case 格式。
//
// 示例：
// CaseDelimitedScreaming("AnyKindOfString", '.') -> ANY.KIND.OF.STRING
// md5:e81c17d2e4a95231
# <翻译结束>


<原文开始>
// treat acronyms as words, eg for JSONData -> JSON is a whole word
<原文结束>

# <翻译开始>
		// 将首字母缩写视为单词，例如 JSONData -> JSON 是一个完整的单词. md5:48305e8e01011121
# <翻译结束>


<原文开始>
// add underscore if next letter case type is changed
<原文结束>

# <翻译开始>
			// 如果下一个字母的大小写类型改变，请添加下划线. md5:6409a4f72dd6f8df
# <翻译结束>


<原文开始>
// replace spaces/underscores with delimiters
<原文结束>

# <翻译开始>
			// 将空格/下划线替换为分隔符. md5:f2c26bc8f3ea056f
# <翻译结束>


<原文开始>
// Converts a string to CamelCase
<原文结束>

# <翻译开始>
// 将字符串转换为驼峰式命名. md5:e7c9de8ba3801cd9
# <翻译结束>

