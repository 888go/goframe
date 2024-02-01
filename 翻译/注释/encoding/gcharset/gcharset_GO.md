
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
// Package gcharset implements character-set conversion functionality.
//
// Supported Character Set:
//
// Chinese : GBK/GB18030/GB2312/Big5
//
// Japanese: EUCJP/ISO2022JP/ShiftJIS
//
// Korean  : EUCKR
//
// Unicode : UTF-8/UTF-16/UTF-16BE/UTF-16LE
//
// Other   : macintosh/IBM*/Windows*/ISO-*
<原文结束>

# <翻译开始>
// Package gcharset 提供字符集转换功能的实现。
//
// 支持的字符集：
//
// 中文 : GBK/GB18030/GB2312/Big5
//
// 日文: EUCJP/ISO2022JP/ShiftJIS
//
// 韩文  : EUCKR
//
// Unicode : UTF-8/UTF-16/UTF-16BE/UTF-16LE
//
// 其他   : macintosh/IBM*/Windows*/ISO-*
// 这段注释描述了一个名为 `gcharset` 的 Go 语言包，它实现了字符集转换功能，并列出了该包支持的多种字符集，包括中文、日文、韩文对应的常见编码以及 Unicode 和其他一些通用或特定平台的字符集。
# <翻译结束>







<原文开始>
// Supported returns whether charset `charset` is supported.
<原文结束>

# <翻译开始>
// Supported 返回 `charset` 字符集是否被支持。
# <翻译结束>


<原文开始>
// Convert converts `src` charset encoding from `srcCharset` to `dstCharset`,
// and returns the converted string.
// It returns `src` as `dst` if it fails converting.
<原文结束>

# <翻译开始>
// Convert 将 `src` 字符串从 `srcCharset` 编码转换为 `dstCharset` 编码，
// 并返回转换后的字符串。
// 如果转换失败，则直接返回原字符串 `src` 作为结果。
# <翻译结束>







<原文开始>
// Do the converting from UTF-8 to `dstCharset`.
<原文结束>

# <翻译开始>
// 将UTF-8转换为`dstCharset`。
# <翻译结束>


<原文开始>
// ToUTF8 converts `src` charset encoding from `srcCharset` to UTF-8 ,
// and returns the converted string.
<原文结束>

# <翻译开始>
// ToUTF8 将 `src` 字符集从 `srcCharset` 转换为 UTF-8，
// 并返回转换后的字符串。
# <翻译结束>


<原文开始>
// UTF8To converts `src` charset encoding from UTF-8 to `dstCharset`,
// and returns the converted string.
<原文结束>

# <翻译开始>
// UTF8To将`src`的字符集编码从UTF-8转换为`dstCharset`，
// 并返回转换后的字符串。
# <翻译结束>


<原文开始>
// getEncoding returns the encoding.Encoding interface object for `charset`.
// It returns nil if `charset` is not supported.
<原文结束>

# <翻译开始>
// getEncoding 函数根据 `charset` 参数返回对应的 encoding.Encoding 接口对象。
// 如果 `charset` 不被支持，则返回 nil。
# <翻译结束>


<原文开始>
// Package_gcharset implements character-set conversion functionality.
//
// Supported Character Set:
//
// Chinese : GBK/GB18030/GB2312/Big5
//
// Japanese: EUCJP/ISO2022JP/ShiftJIS
//
// Korean  : EUCKR
//
// Unicode : UTF-8/UTF-16/UTF-16BE/UTF-16LE
//
// Other   : macintosh/IBM*/Windows*/ISO-*
<原文结束>

# <翻译开始>
// Package_gcharset实现了字符集转换功能。
//
// 支持的字符集：
//
// 中文 : GBK/GB18030/GB2312/Big5
//
// 日文: EUCJP/ISO2022JP/ShiftJIS
//
// 韩文  : EUCKR
//
// Unicode : UTF-8/UTF-16/UTF-16BE/UTF-16LE
//
// 其他   : macintosh/IBM系列/Windows系列/ISO系列
# <翻译结束>


<原文开始>
// Alias for charsets.
<原文结束>

# <翻译开始>
// 对字符集的别名。
# <翻译结束>


<原文开始>
// Converting `src` to UTF-8.
<原文结束>

# <翻译开始>
// 将 `src` 转换为 UTF-8。
# <翻译结束>

