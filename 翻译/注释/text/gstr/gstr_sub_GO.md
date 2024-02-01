
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
// Str returns part of `haystack` string starting from and including
// the first occurrence of `needle` to the end of `haystack`.
// See http://php.net/manual/en/function.strstr.php.
// Eg: Str("12345", "3") => "345"
<原文结束>

# <翻译开始>
// Str 函数返回从 `haystack` 字符串中第一个（包含）出现的 `needle` 字符串开始，
// 直至 `haystack` 结尾的部分。参考：http://php.net/manual/en/function.strstr.php。
// 示例：Str("12345", "3") => "345"
# <翻译结束>


<原文开始>
// StrEx returns part of `haystack` string starting from and excluding
// the first occurrence of `needle` to the end of `haystack`.
// Eg: StrEx("12345", "3") => "45"
<原文结束>

# <翻译开始>
// StrEx 函数从 `haystack` 字符串中第一个不包含 `needle` 的位置开始，截取到 `haystack` 末尾的部分并返回。
// 示例：StrEx("12345", "3") => "45"
# <翻译结束>


<原文开始>
// StrTill returns part of `haystack` string ending to and including
// the first occurrence of `needle` from the start of `haystack`.
// Eg: StrTill("12345", "3") => "123"
<原文结束>

# <翻译开始>
// StrTill 函数返回从 `haystack` 字符串起始位置到（并包含）
// 第一个出现的 `needle` 子串为止的部分字符串。
// 示例：StrTill("12345", "3") => 返回结果为 "123"
# <翻译结束>


<原文开始>
// StrTillEx returns part of `haystack` string ending to and excluding
// the first occurrence of `needle` from the start of `haystack`.
// Eg: StrTillEx("12345", "3") => "12"
<原文结束>

# <翻译开始>
// StrTillEx 从`haystack`字符串的起始位置截取，返回直到（但不包括）第一次出现`needle`子串的部分。
// 示例：StrTillEx("12345", "3") => "12"
# <翻译结束>


<原文开始>
// SubStr returns a portion of string `str` specified by the `start` and `length` parameters.
// The parameter `length` is optional, it uses the length of `str` in default.
// Eg: SubStr("12345", 1, 2) => "23"
<原文结束>

# <翻译开始>
// SubStr 函数返回字符串 `str` 中由 `start` 和 `length` 参数指定的部分子串。
// 参数 `length` 是可选的，默认情况下会使用 `str` 的长度。
// 示例：SubStr("12345", 1, 2) => "23"
# <翻译结束>


<原文开始>
// SubStrRune returns a portion of string `str` specified by the `start` and `length` parameters.
// SubStrRune considers parameter `str` as unicode string.
// The parameter `length` is optional, it uses the length of `str` in default.
<原文结束>

# <翻译开始>
// SubStrRune 返回字符串 `str` 中由 `start` 和 `length` 参数指定的部分。
// SubStrRune 将参数 `str` 视为unicode字符串处理。
// 参数 `length` 是可选的，默认情况下它使用 `str` 的长度。
// 更详细的翻译：
// ```go
// SubStrRune 函数返回给定字符串 `str` 从 `start` 位置开始的一个子串。
// 在此函数中，我们把输入的字符串 `str` 当作Unicode字符序列进行处理。
// 参数 `length` 是可选的，如果不提供，则默认截取从 `start` 到字符串结尾的所有字符。
# <翻译结束>


<原文开始>
// Converting to []rune to support unicode.
<原文结束>

# <翻译开始>
// 转换为 []rune 以支持Unicode。
# <翻译结束>


<原文开始>
// StrLimit returns a portion of string `str` specified by `length` parameters, if the length
// of `str` is greater than `length`, then the `suffix` will be appended to the result string.
<原文结束>

# <翻译开始>
// StrLimit 返回字符串 `str` 指定长度为 `length` 的部分，如果 `str` 的长度大于 `length`，
// 则结果字符串后会追加 `suffix`。
# <翻译结束>


<原文开始>
// StrLimitRune returns a portion of string `str` specified by `length` parameters, if the length
// of `str` is greater than `length`, then the `suffix` will be appended to the result string.
// StrLimitRune considers parameter `str` as unicode string.
<原文结束>

# <翻译开始>
// StrLimitRune 返回字符串 `str` 指定长度的子串，若 `str` 的长度大于 `length`，
// 则结果字符串末尾会追加 `suffix`。此函数将参数 `str` 视为 unicode 字符串处理。
# <翻译结束>


<原文开始>
// SubStrFrom returns a portion of string `str` starting from first occurrence of and including `need`
// to the end of `str`.
<原文结束>

# <翻译开始>
// SubStrFrom 返回字符串 `str` 中从第一个出现并包含 `need` 的子串开始，直到 `str` 末尾的部分。
# <翻译结束>


<原文开始>
// SubStrFromEx returns a portion of string `str` starting from first occurrence of and excluding `need`
// to the end of `str`.
<原文结束>

# <翻译开始>
// SubStrFromEx 从字符串 `str` 中返回从第一个出现且不包括 `need` 的子串到 `str` 结尾的部分。
# <翻译结束>


<原文开始>
// SubStrFromR returns a portion of string `str` starting from last occurrence of and including `need`
// to the end of `str`.
<原文结束>

# <翻译开始>
// SubStrFromR 返回字符串 `str` 从最后一个出现且包含 `need` 的子串开始，直到 `str` 结尾的部分。
# <翻译结束>


<原文开始>
// SubStrFromREx returns a portion of string `str` starting from last occurrence of and excluding `need`
// to the end of `str`.
<原文结束>

# <翻译开始>
// SubStrFromREx 函数从字符串 `str` 中返回从最后一次出现且不包含 `need` 的子串开始，直到 `str` 末尾的部分。
# <翻译结束>

