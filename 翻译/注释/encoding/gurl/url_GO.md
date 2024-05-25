
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
// Package gurl provides useful API for URL handling.
<原文结束>

# <翻译开始>
// 包gurl提供了处理URL的有用API。 md5:3954efb697af4a41
# <翻译结束>


<原文开始>
// Encode escapes the string so it can be safely placed
// inside an URL query.
<原文结束>

# <翻译开始>
// Encode将字符串进行转义，以便安全地放置在URL查询中。
// md5:2e139b94de8d8e81
# <翻译结束>


<原文开始>
// Decode does the inverse transformation of Encode,
// converting each 3-byte encoded substring of the form "%AB" into the
// hex-decoded byte 0xAB.
// It returns an error if any % is not followed by two hexadecimal
// digits.
<原文结束>

# <翻译开始>
// Decode 执行与 Encode 相反的转换，
// 将形如 "%AB" 的每3字节编码子串转换为其十六进制解码字节 0xAB。
// 如果任何百分号（%）后面没有跟随两个十六进制数字，它将返回一个错误。
// md5:c8ff43c799b800c0
# <翻译结束>


<原文开始>
// RawEncode does encode the given string according
// URL-encode according to RFC 3986.
// See http://php.net/manual/en/function.rawurlencode.php.
<原文结束>

# <翻译开始>
// RawEncode 按照 RFC 3986 标准对给定的字符串进行原始URL编码。
// 参考 http://php.net/manual/en/function.rawurlencode.php。
// md5:b116dd32b351afc8
# <翻译结束>


<原文开始>
// RawDecode does decode the given string
// Decode URL-encoded strings.
// See http://php.net/manual/en/function.rawurldecode.php.
<原文结束>

# <翻译开始>
// RawDecode 用于解码给定的字符串
// 解码 URL 编码的字符串。
// 参考：http://php.net/manual/zh/function.rawurldecode.php
// md5:ffbb20457d038fe3
# <翻译结束>


<原文开始>
// BuildQuery Generate URL-encoded query string.
// See http://php.net/manual/en/function.http-build-query.php.
<原文结束>

# <翻译开始>
// BuildQuery 生成 URL 编码的查询字符串。
// 参考：http://php.net/manual/zh/function.http-build-query.php。
// md5:f0e4222e29189a30
# <翻译结束>


<原文开始>
// ParseURL Parse an URL and return its components.
// -1: all; 1: scheme; 2: host; 4: port; 8: user; 16: pass; 32: path; 64: query; 128: fragment.
// See http://php.net/manual/en/function.parse-url.php.
<原文结束>

# <翻译开始>
// ParseURL 解析一个URL并返回其组成部分。
// 参数可选值：-1表示全部；1表示方案（scheme）；2表示主机（host）；4表示端口（port）；
// 8表示用户名（user）；16表示密码（pass）；32表示路径（path）；
// 64表示查询字符串（query）；128表示片段（fragment）。
// 参考 PHP 手册中的函数 parse-url：http://php.net/manual/en/function.parse-url.php。
// md5:ab33f23dd1fe61ca
# <翻译结束>

