
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
// Package httputil provides HTTP functions for internal usage only.
<原文结束>

# <翻译开始>
// httputil 包提供内部使用的HTTP功能。
# <翻译结束>


<原文开始>
// BuildParams builds the request string for the http client. The `params` can be type of:
// string/[]byte/map/struct/*struct.
//
// The optional parameter `noUrlEncode` specifies whether ignore the url encoding for the data.
<原文结束>

# <翻译开始>
// BuildParams 为 http 客户端构建请求字符串。`params` 参数可以是以下类型：
// string/[]byte/map/struct/*struct。
//
// 可选参数 `noUrlEncode` 指定是否忽略数据的 URL 编码。
// 详细解释：
// 此 Go 语言注释描述了一个名为 `BuildParams` 的函数，该函数用于根据给定的参数构建一个 HTTP 客户端的请求字符串。`params` 参数支持多种数据类型，包括但不限于：字符串(string)、字节切片([]byte)、映射(map)、结构体(struct)以及指向结构体的指针(*struct)。
// 另外，该函数接受一个可选参数 `noUrlEncode`，这个布尔类型的参数决定了在构建请求字符串时是否跳过对数据进行 URL 编码处理。如果设为 `true`，则表示不进行 URL 编码；否则（默认情况或设为 `false`），将对数据进行 URL 编码。
# <翻译结束>


<原文开始>
// If given string/[]byte, converts and returns it directly as string.
<原文结束>

# <翻译开始>
// 如果给定的是字符串或[]byte，直接将其转换并返回为字符串。
# <翻译结束>


<原文开始>
// Else converts it to map and does the url encoding.
<原文结束>

# <翻译开始>
// Else 将其转换为 map，并执行 URL 编码。
# <翻译结束>


<原文开始>
// If there's file uploading, it ignores the url encoding.
<原文结束>

# <翻译开始>
// 如果存在文件上传，则忽略URL编码。
# <翻译结束>







<原文开始>
// No url encoding if uploading file.
<原文结束>

# <翻译开始>
// 如果上传文件，则不进行URL编码
# <翻译结束>


<原文开始>
// HeaderToMap coverts request headers to map.
<原文结束>

# <翻译开始>
// HeaderToMap 将请求头转换为映射（map）。
# <翻译结束>


<原文开始>
// Ignore nil attributes.
<原文结束>

# <翻译开始>
// 忽略nil属性。
# <翻译结束>

