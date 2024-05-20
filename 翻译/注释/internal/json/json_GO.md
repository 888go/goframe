
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
// Package json provides json operations wrapping ignoring stdlib or third-party lib json.
<原文结束>

# <翻译开始>
// json包提供了对json操作的封装，忽略了标准库或第三方库的json。. md5:8c700638d650aacd
# <翻译结束>


<原文开始>
// RawMessage is a raw encoded JSON value.
// It implements Marshaler and Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.
<原文结束>

# <翻译开始>
// RawMessage 是一个原始的编码JSON值。
// 它实现了Marshaler和Unmarshaler接口，可以用于延迟JSON解码或预先计算JSON编码。
// md5:2ea51fc4bfe3af87
# <翻译结束>


<原文开始>
// Marshal adapts to json/encoding Marshal API.
//
// Marshal returns the JSON encoding of v, adapts to json/encoding Marshal API
// Refer to https://godoc.org/encoding/json#Marshal for more information.
<原文结束>

# <翻译开始>
// Marshal 适应了json/encoding.Marshal API。
//
// Marshal 返回 v 的 JSON 编码，适应了 json/encoding.Marshal API。更多信息请参考 https://godoc.org/encoding/json#Marshal。
// md5:e67e9d2efbcb1d3c
# <翻译结束>


<原文开始>
// MarshalIndent same as json.MarshalIndent.
<原文结束>

# <翻译开始>
// MarshalIndent 和 json.MarshalIndent 功能相同。. md5:1c3ee8bca0354fac
# <翻译结束>


<原文开始>
// Unmarshal adapts to json/encoding Unmarshal API
//
// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
// Refer to https://godoc.org/encoding/json#Unmarshal for more information.
<原文结束>

# <翻译开始>
// Unmarshal 适应于 json/encoding 的 Unmarshal API
//
// Unmarshal 解析 JSON 编码的数据，并将结果存储到由 v 指向的值中。
// 更多信息，请参考 https://godoc.org/encoding/json#Unmarshal。
// md5:7272cfbd647a7f0f
# <翻译结束>


<原文开始>
// UnmarshalUseNumber decodes the json data bytes to target interface using number option.
<原文结束>

# <翻译开始>
// UnmarshalUseNumber 使用数字选项将json数据字节解码到目标接口。. md5:b04f82165b9c9933
# <翻译结束>


<原文开始>
// NewEncoder same as json.NewEncoder
<原文结束>

# <翻译开始>
// NewEncoder 等同于 json.NewEncoder. md5:4f4b9a0b5e712ba5
# <翻译结束>


<原文开始>
// NewDecoder adapts to json/stream NewDecoder API.
//
// NewDecoder returns a new decoder that reads from r.
//
// Instead of a json/encoding Decoder, a Decoder is returned
// Refer to https://godoc.org/encoding/json#NewDecoder for more information.
<原文结束>

# <翻译开始>
// NewDecoder 封装了 json/stream 中的 NewDecoder API。
//
// NewDecoder 从 r 读取数据，返回一个新的解码器。
//
// 返回的不是一个 json/encoding 中的 Decoder，而是 Decoder。更多详细信息请参考：https://godoc.org/encoding/json#NewDecoder。
// md5:eea702294e95e71d
# <翻译结束>


<原文开始>
// Valid reports whether data is a valid JSON encoding.
<原文结束>

# <翻译开始>
// Valid 报告数据是否为有效的JSON编码。. md5:db76d6317d6c0b76
# <翻译结束>

