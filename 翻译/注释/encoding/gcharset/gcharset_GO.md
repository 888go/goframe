
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
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
// Package gcharset 实现字符集转换功能。
//
// 支持的字符集：
//
// 中文：GBK/GB18030/GB2312/Big5
//
// 日本语：EUCJP/ISO2022JP/ShiftJIS
//
// 韩语：EUCKR
//
// Unicode：UTF-8/UTF-16/UTF-16BE/UTF-16LE
//
// 其他：macintosh/IBM*/Windows*/ISO-*（*表示各种变种） md5:5f95b30c9186a77b
# <翻译结束>


<原文开始>
// Supported returns whether charset `charset` is supported.
<原文结束>

# <翻译开始>
// Supported 返回字符集 `charset` 是否受支持。 md5:ecb209536b99e114
# <翻译结束>


<原文开始>
// Convert converts `src` charset encoding from `srcCharset` to `dstCharset`,
// and returns the converted string.
// It returns `src` as `dst` if it fails converting.
<原文结束>

# <翻译开始>
// Convert 将 `src` 字符串的编码从 `srcCharset` 转换为 `dstCharset`，并返回转换后的字符串。
// 如果转换失败，则返回原 `src` 作为 `dst`。 md5:d579c6167a34081f
# <翻译结束>


<原文开始>
// Converting `src` to UTF-8.
<原文结束>

# <翻译开始>
	// 将 `src` 转换为 UTF-8 编码。 md5:345cd013199770a3
# <翻译结束>


<原文开始>
// Do the converting from UTF-8 to `dstCharset`.
<原文结束>

# <翻译开始>
	// 将UTF-8转换为`dstCharset`。 md5:4caf3880c33fb49d
# <翻译结束>


<原文开始>
// ToUTF8 converts `src` charset encoding from `srcCharset` to UTF-8 ,
// and returns the converted string.
<原文结束>

# <翻译开始>
// ToUTF8 将 `src` 字符串的字符集编码从 `srcCharset` 转换为 UTF-8，
// 并返回转换后的字符串。 md5:ed113e096f11dcee
# <翻译结束>


<原文开始>
// UTF8To converts `src` charset encoding from UTF-8 to `dstCharset`,
// and returns the converted string.
<原文结束>

# <翻译开始>
// UTF8To 将 `src` 字符集编码从 UTF-8 转换为 `dstCharset`，
// 并返回转换后的字符串。 md5:6d376918eb2969a6
# <翻译结束>


<原文开始>
// getEncoding returns the encoding.Encoding interface object for `charset`.
// It returns nil if `charset` is not supported.
<原文结束>

# <翻译开始>
// getEncoding 返回与 `charset` 对应的 encoding.Encoding 接口对象。如果 `charset` 不被支持，它将返回 nil。 md5:8770abf28a404b1b
# <翻译结束>

