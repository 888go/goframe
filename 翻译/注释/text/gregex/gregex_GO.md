
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Package gregex provides high performance API for regular expression functionality.
<原文结束>

# <翻译开始>
// 包gregex提供了正则表达式功能的高性能API。. md5:5d43833868579329
# <翻译结束>


<原文开始>
// Quote quotes `s` by replacing special chars in `s`
// to match the rules of regular expression pattern.
// And returns the copy.
//
// Eg: Quote(`[foo]`) returns `\[foo\]`.
<原文结束>

# <翻译开始>
// Quote 通过替换`s`中的特殊字符，使其符合正则表达式模式的规则，对`s`进行引号包裹。然后返回修改后的字符串。
// 
// 例如：Quote(`[foo]`) 返回 `\[foo\]`。
// md5:baab9e0870efe45f
# <翻译结束>


<原文开始>
// Validate checks whether given regular expression pattern `pattern` valid.
<原文结束>

# <翻译开始>
// Validate 检查给定的正则表达式模式 `pattern` 是否有效。. md5:39fda51666585abe
# <翻译结束>


<原文开始>
// IsMatch checks whether given bytes `src` matches `pattern`.
<原文结束>

# <翻译开始>
// IsMatch 检查给定的字节`src`是否匹配`pattern`。. md5:7f26688fb33d288d
# <翻译结束>


<原文开始>
// IsMatchString checks whether given string `src` matches `pattern`.
<原文结束>

# <翻译开始>
// IsMatchString 检查给定的字符串 `src` 是否与 `pattern` 匹配。. md5:b4afd86a0688ae19
# <翻译结束>


<原文开始>
// Match return bytes slice that matched `pattern`.
<原文结束>

# <翻译开始>
// Match 返回匹配`pattern`的字节切片。. md5:7f3dd939e61e4db8
# <翻译结束>


<原文开始>
// MatchString return strings that matched `pattern`.
<原文结束>

# <翻译开始>
// MatchString 返回与`pattern`匹配的字符串。. md5:d4d623eec4e1f3ec
# <翻译结束>


<原文开始>
// MatchAll return all bytes slices that matched `pattern`.
<原文结束>

# <翻译开始>
// MatchAll 返回所有匹配`pattern`的字节切片。. md5:64871f15e4916377
# <翻译结束>


<原文开始>
// MatchAllString return all strings that matched `pattern`.
<原文结束>

# <翻译开始>
// MatchAllString 返回所有与`pattern`匹配的字符串。. md5:ec34b1802db69c97
# <翻译结束>


<原文开始>
// Replace replaces all matched `pattern` in bytes `src` with bytes `replace`.
<原文结束>

# <翻译开始>
// Replace 将 `src` 字节中的所有匹配 `pattern` 的部分替换为 `replace` 字节。. md5:5d5c7ad162f72858
# <翻译结束>


<原文开始>
// ReplaceString replace all matched `pattern` in string `src` with string `replace`.
<原文结束>

# <翻译开始>
// ReplaceString 将字符串 `src` 中所有匹配的 `pattern` 替换为字符串 `replace`。. md5:1e7fdbe12a62e763
# <翻译结束>


<原文开始>
// ReplaceFunc replace all matched `pattern` in bytes `src`
// with custom replacement function `replaceFunc`.
<原文结束>

# <翻译开始>
// ReplaceFunc 使用自定义的替换函数 `replaceFunc`，将字节切片 `src` 中所有匹配的 `pattern` 替换。
// md5:3b66619bd59d4056
# <翻译结束>


<原文开始>
// ReplaceFuncMatch replace all matched `pattern` in bytes `src`
// with custom replacement function `replaceFunc`.
// The parameter `match` type for `replaceFunc` is [][]byte,
// which is the result contains all sub-patterns of `pattern` using Match function.
<原文结束>

# <翻译开始>
// ReplaceFuncMatch：在字节`src`中使用自定义替换函数`replaceFunc`替换所有匹配的`pattern`。`replaceFunc`的参数`match`类型为`[][]byte`，它包含了`pattern`使用Match函数的所有子模式的结果。
// md5:cdbed5cefac02741
# <翻译结束>


<原文开始>
// ReplaceStringFunc replace all matched `pattern` in string `src`
// with custom replacement function `replaceFunc`.
<原文结束>

# <翻译开始>
// ReplaceStringFunc 函数会在字符串 `src` 中替换所有匹配的 `pattern`，使用自定义的替换函数 `replaceFunc`。
// md5:8575760795474682
# <翻译结束>


<原文开始>
// ReplaceStringFuncMatch replace all matched `pattern` in string `src`
// with custom replacement function `replaceFunc`.
// The parameter `match` type for `replaceFunc` is []string,
// which is the result contains all sub-patterns of `pattern` using MatchString function.
<原文结束>

# <翻译开始>
// ReplaceStringFuncMatch 将字符串 `src` 中所有与 `pattern` 匹配的部分
// 使用自定义替换函数 `replaceFunc` 进行替换。
// 替换函数 `replaceFunc` 的参数 `match` 类型为 []string，
// 它包含使用 MatchString 函数得到的 `pattern` 中所有子模式匹配结果。
// md5:b24f208b16cfd56a
# <翻译结束>


<原文开始>
// Split slices `src` into substrings separated by the expression and returns a slice of
// the substrings between those expression matches.
<原文结束>

# <翻译开始>
// Split 将切片 `src` 按照给定的表达式分割成多个子字符串，并返回这些匹配之间的子字符串切片。
// md5:e0809df699cf82c1
# <翻译结束>

