
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Package gregex provides high performance API for regular expression functionality.
<原文结束>

# <翻译开始>
// Package gregex 提供了用于正则表达式功能的高性能 API。
# <翻译结束>


<原文开始>
// Quote quotes `s` by replacing special chars in `s`
// to match the rules of regular expression pattern.
// And returns the copy.
//
// Eg: Quote(`[foo]`) returns `\[foo\]`.
<原文结束>

# <翻译开始>
// Quote通过将`s`中的特殊字符替换为符合正则表达式模式的规则，
// 并返回处理后的副本。
//
// 例如：Quote(`[foo]`) 将返回 `\[foo\]`。
# <翻译结束>


<原文开始>
// Validate checks whether given regular expression pattern `pattern` valid.
<原文结束>

# <翻译开始>
// Validate 检查给定的正则表达式模式 `pattern` 是否有效。
# <翻译结束>


<原文开始>
// IsMatch checks whether given bytes `src` matches `pattern`.
<原文结束>

# <翻译开始>
// IsMatch 检查给定的字节序列 `src` 是否匹配模式 `pattern`。
# <翻译结束>


<原文开始>
// IsMatchString checks whether given string `src` matches `pattern`.
<原文结束>

# <翻译开始>
// IsMatchString 检查给定的字符串 `src` 是否与 `pattern` 匹配。
# <翻译结束>


<原文开始>
// Match return bytes slice that matched `pattern`.
<原文结束>

# <翻译开始>
// Match 返回匹配 `pattern` 的字节切片。
# <翻译结束>


<原文开始>
// MatchString return strings that matched `pattern`.
<原文结束>

# <翻译开始>
// MatchString 返回匹配`pattern`的字符串。
# <翻译结束>


<原文开始>
// MatchAll return all bytes slices that matched `pattern`.
<原文结束>

# <翻译开始>
// MatchAll 返回所有匹配 `pattern` 的字节切片。
# <翻译结束>


<原文开始>
// MatchAllString return all strings that matched `pattern`.
<原文结束>

# <翻译开始>
// MatchAllString 返回所有匹配 `pattern` 的字符串。
# <翻译结束>


<原文开始>
// Replace replaces all matched `pattern` in bytes `src` with bytes `replace`.
<原文结束>

# <翻译开始>
// Replace 将字节 `src` 中所有匹配到的 `pattern` 替换为字节 `replace`。
# <翻译结束>


<原文开始>
// ReplaceString replace all matched `pattern` in string `src` with string `replace`.
<原文结束>

# <翻译开始>
// ReplaceString 将字符串 `src` 中所有匹配到的 `pattern` 替换为字符串 `replace`。
# <翻译结束>


<原文开始>
// ReplaceFunc replace all matched `pattern` in bytes `src`
// with custom replacement function `replaceFunc`.
<原文结束>

# <翻译开始>
// ReplaceFunc 将字节切片 `src` 中所有匹配到的 `pattern` 用自定义替换函数 `replaceFunc` 进行替换。
# <翻译结束>


<原文开始>
// ReplaceFuncMatch replace all matched `pattern` in bytes `src`
// with custom replacement function `replaceFunc`.
// The parameter `match` type for `replaceFunc` is [][]byte,
// which is the result contains all sub-patterns of `pattern` using Match function.
<原文结束>

# <翻译开始>
// ReplaceFuncMatch 在字节切片 `src` 中替换所有匹配的 `pattern`，
// 使用自定义替换函数 `replaceFunc` 进行替换。
// `replaceFunc` 参数中的 `match` 类型为 [][]byte，
// 它是使用 Match 函数得到的所有 `pattern` 子模式的结果。
# <翻译结束>


<原文开始>
// ReplaceStringFunc replace all matched `pattern` in string `src`
// with custom replacement function `replaceFunc`.
<原文结束>

# <翻译开始>
// ReplaceStringFunc 在字符串 `src` 中使用自定义替换函数 `replaceFunc` 替换所有匹配到的 `pattern`。
# <翻译结束>


<原文开始>
// ReplaceStringFuncMatch replace all matched `pattern` in string `src`
// with custom replacement function `replaceFunc`.
// The parameter `match` type for `replaceFunc` is []string,
// which is the result contains all sub-patterns of `pattern` using MatchString function.
<原文结束>

# <翻译开始>
// ReplaceStringFuncMatch 在字符串 `src` 中使用自定义替换函数 `replaceFunc` 替换所有匹配到的 `pattern`。
// 参数 `replaceFunc` 的形参类型为 `[]string`，该结果包含通过 MatchString 函数得到的所有 `pattern` 的子模式。
# <翻译结束>


<原文开始>
// Split slices `src` into substrings separated by the expression and returns a slice of
// the substrings between those expression matches.
<原文结束>

# <翻译开始>
// Split 函数将 `src` 切片按照表达式进行分割，并返回由这些表达式匹配之间的子字符串构成的切片。
# <翻译结束>

