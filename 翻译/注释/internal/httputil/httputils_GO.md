
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
// Package httputil provides HTTP functions for internal usage only.
<原文结束>

# <翻译开始>
// 包 httputil 提供仅供内部使用的 HTTP 函数。 md5:68a87514ccfd9190
# <翻译结束>


<原文开始>
// BuildParams builds the request string for the http client. The `params` can be type of:
// string/[]byte/map/struct/*struct.
//
// The optional parameter `noUrlEncode` specifies whether ignore the url encoding for the data.
<原文结束>

# <翻译开始>
// BuildParams 为http客户端构建请求字符串。`params`可以是以下类型：
// 字符串/字节切片/映射/结构体/结构体指针。
//
// 可选参数`noUrlEncode`指定是否忽略数据的URL编码。 md5:664ad104f4b3f610
# <翻译结束>


<原文开始>
// If given string/[]byte, converts and returns it directly as string.
<原文结束>

# <翻译开始>
	// 如果给定字符串/字节切片，会直接转换并返回它作为字符串。 md5:80d9827515b7e847
# <翻译结束>


<原文开始>
// Else converts it to map and does the url encoding.
<原文结束>

# <翻译开始>
	// 否则，它会将它转换为map并进行URL编码。 md5:932b0b216ae84f60
# <翻译结束>


<原文开始>
// If there's file uploading, it ignores the url encoding.
<原文结束>

# <翻译开始>
	// 如果有文件上传，它将忽略URL编码。 md5:e349803af0cef3a3
# <翻译结束>


<原文开始>
// No url encoding if uploading file.
<原文结束>

# <翻译开始>
				// 如果正在上传文件，则不进行URL编码。 md5:1d89b2d337a7a0e9
# <翻译结束>


<原文开始>
// HeaderToMap coverts request headers to map.
<原文结束>

# <翻译开始>
// HeaderToMap 将请求头转换为映射。 md5:d7b057a672ffda30
# <翻译结束>

