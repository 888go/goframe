
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
// Str returns part of `haystack` string starting from and including
// the first occurrence of `needle` to the end of `haystack`.
//
// This function performs exactly as function SubStr, but to implement the same function
// as PHP: http://php.net/manual/en/function.strstr.php.
//
// Example:
// Str("av.mp4", ".") -> ".mp4"
<原文结束>

# <翻译开始>
// Str 函数返回从 `haystack` 字符串开始，包括第一个出现的 `needle` 到 `haystack` 结尾的部分。
//
// 这个函数的行为与 SubStr 函数完全相同，但是为了实现与 PHP 相同的功能：http://php.net/manual/en/function.strstr.php。
//
// 示例：
// Str("av.mp4", ".") -> ".mp4"
// md5:35b0375920cdf357
# <翻译结束>


<原文开始>
// StrEx returns part of `haystack` string starting from and excluding
// the first occurrence of `needle` to the end of `haystack`.
//
// This function performs exactly as function SubStrEx, but to implement the same function
// as PHP: http://php.net/manual/en/function.strstr.php.
//
// Example:
// StrEx("av.mp4", ".") -> "mp4"
<原文结束>

# <翻译开始>
// StrEx 从`haystack`字符串中返回从第一个出现的`needle`开始到`haystack`末尾的部分。
//
// 这个函数的行为与SubStrEx函数完全相同，但实现了与PHP相同的函数：http://php.net/manual/en/function.strstr.php。
//
// 示例：
// StrEx("av.mp4", ".") -> "mp4"
// md5:1f6c13b098caa33a
# <翻译结束>


<原文开始>
// StrTill returns part of `haystack` string ending to and including
// the first occurrence of `needle` from the start of `haystack`.
//
// Example:
// StrTill("av.mp4", ".") -> "av."
<原文结束>

# <翻译开始>
// StrTill 返回 `haystack` 字符串中从开头到（包括）第一个 `needle` 出现的部分。
//
// 示例：
// StrTill("av.mp4", ".") -> "av."
// md5:f914c156abb95437
# <翻译结束>


<原文开始>
// StrTillEx returns part of `haystack` string ending to and excluding
// the first occurrence of `needle` from the start of `haystack`.
//
// Example:
// StrTillEx("av.mp4", ".") -> "av"
<原文结束>

# <翻译开始>
// StrTillEx 函数返回 `haystack` 字符串中从开始到（但不包括）第一次出现 `needle` 的部分。
//
// 示例：
// StrTillEx("av.mp4", ".") -> "av"
// md5:c0848291d8036d82
# <翻译结束>


<原文开始>
// SubStr returns a portion of string `str` specified by the `start` and `length` parameters.
// The parameter `length` is optional, it uses the length of `str` in default.
//
// Example:
// SubStr("123456", 1, 2) -> "23"
<原文结束>

# <翻译开始>
// SubStr 函数返回字符串 `str` 中由 `start` 和 `length` 参数指定的部分。参数 `length` 是可选的，如果未提供，则默认使用 `str` 的长度。
// 
// 示例：
// SubStr("123456", 1, 2) -> "23"
// md5:b6da71b3534fdbbc
# <翻译结束>


<原文开始>
// Converting to []rune to support unicode.
<原文结束>

# <翻译开始>
// 转换为[]rune以支持Unicode。 md5:459540c13f4e5603
# <翻译结束>


<原文开始>
// StrLimit returns a portion of string `str` specified by `length` parameters, if the length
// of `str` is greater than `length`, then the `suffix` will be appended to the result string.
//
// Example:
// StrLimit("123456", 3)      -> "123..."
// StrLimit("123456", 3, "~") -> "123~"
<原文结束>

# <翻译开始>
// StrLimit 函数返回字符串 `str` 中由 `length` 参数指定长度的部分。如果 `str` 的长度大于 `length`，则 `suffix` 将被添加到结果字符串的末尾。
//
// 示例：
// StrLimit("123456", 3)      -> "123..."
// StrLimit("123456", 3, "~") -> "123~"
// md5:bd8f96405a5594b5
# <翻译结束>


<原文开始>
// SubStrFrom returns a portion of string `str` starting from first occurrence of and including `need`
// to the end of `str`.
//
// Example:
// SubStrFrom("av.mp4", ".") -> ".mp4"
<原文结束>

# <翻译开始>
// SubStrFrom 从字符串 `str` 中从第一次出现 `need` 的位置开始，包括 `need` 到字符串末尾的部分。
//
// 示例：
// SubStrFrom("av.mp4", ".") -> ".mp4"
// md5:f4bff02c473abeff
# <翻译结束>


<原文开始>
// SubStrFromEx returns a portion of string `str` starting from first occurrence of and excluding `need`
// to the end of `str`.
//
// Example:
// SubStrFromEx("av.mp4", ".") -> "mp4"
<原文结束>

# <翻译开始>
// SubStrFromEx 从字符串 `str` 中返回从首次出现 `need` 到 `str` 结尾的部分（不包括 `need`）。
//
// 示例：
// SubStrFromEx("av.mp4", ".") -> "mp4"
// md5:88a817f03fc77455
# <翻译结束>


<原文开始>
// SubStrFromR returns a portion of string `str` starting from last occurrence of and including `need`
// to the end of `str`.
//
// Example:
// SubStrFromR("/dev/vda", "/") -> "/vda"
<原文结束>

# <翻译开始>
// SubStrFromR 从字符串 `str` 的最后一个出现的 `need` 开始并包括在内，返回一个子串。
// 示例：
// SubStrFromR("/dev/vda", "/") -> "/vda"
// md5:8f70ecc84d0338f8
# <翻译结束>


<原文开始>
// SubStrFromREx returns a portion of string `str` starting from last occurrence of and excluding `need`
// to the end of `str`.
//
// Example:
// SubStrFromREx("/dev/vda", "/") -> "vda"
<原文结束>

# <翻译开始>
// SubStrFromREx 从字符串 `str` 中最后一个出现的 `need` 子串之后的字符开始，直到 `str` 的末尾，返回这一部分子串。
//
// 示例：
// SubStrFromREx("/dev/vda", "/") -> "vda"
// md5:3de495ad97b12196
# <翻译结束>

