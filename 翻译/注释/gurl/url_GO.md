
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
// Package gurl provides useful API for URL handling.
<原文结束>

# <翻译开始>
// 包gurl提供了用于URL处理的有用API。
# <翻译结束>


<原文开始>
// Encode escapes the string so it can be safely placed
// inside an URL query.
<原文结束>

# <翻译开始>
// Encode 对字符串进行转义，以便它可以安全地放置在
// URL 查询参数内。
# <翻译结束>


<原文开始>
// Decode does the inverse transformation of Encode,
// converting each 3-byte encoded substring of the form "%AB" into the
// hex-decoded byte 0xAB.
// It returns an error if any % is not followed by two hexadecimal
// digits.
<原文结束>

# <翻译开始>
// Decode完成与Encode相反的转换操作，
// 将形式为"%AB"的每个3字节编码子串转化为十六进制解码后的字节0xAB。
// 如果存在百分号（%）后跟随的不是两个十六进制数字，则返回错误。
# <翻译结束>


<原文开始>
// RawEncode does encode the given string according
// URL-encode according to RFC 3986.
// See http://php.net/manual/en/function.rawurlencode.php.
<原文结束>

# <翻译开始>
// RawEncode 对给定的字符串进行编码，
// 根据 RFC 3986 进行 URL 编码。
// 参见 http://php.net/manual/en/function.rawurlencode.php.
# <翻译结束>


<原文开始>
// RawDecode does decode the given string
// Decode URL-encoded strings.
// See http://php.net/manual/en/function.rawurldecode.php.
<原文结束>

# <翻译开始>
// RawDecode用于解码给定的字符串
// 解码URL编码过的字符串。
// 参考：http://php.net/manual/en/function.rawurldecode.php.
# <翻译结束>


<原文开始>
// BuildQuery Generate URL-encoded query string.
// See http://php.net/manual/en/function.http-build-query.php.
<原文结束>

# <翻译开始>
// BuildQuery 生成 URL 编码的查询字符串。
// 参考：http://php.net/manual/en/function.http-build-query.php.
# <翻译结束>


<原文开始>
// ParseURL Parse an URL and return its components.
// -1: all; 1: scheme; 2: host; 4: port; 8: user; 16: pass; 32: path; 64: query; 128: fragment.
// See http://php.net/manual/en/function.parse-url.php.
<原文结束>

# <翻译开始>
// ParseURL 解析一个URL并返回其组成部分。
// 参数：
// -1：所有部分；1：方案（如http）；2：主机名；4：端口；8：用户名；16：密码；32：路径；64：查询字符串；128：片段（锚点）。
// 参考：http://php.net/manual/en/function.parse-url.php.
# <翻译结束>

