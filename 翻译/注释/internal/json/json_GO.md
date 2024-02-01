
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
// Package json provides json operations wrapping ignoring stdlib or third-party lib json.
<原文结束>

# <翻译开始>
// Package json 提供了围绕标准库或第三方库的 JSON 操作，实现了对 JSON 的封装并忽略它们。
# <翻译结束>


<原文开始>
// RawMessage is a raw encoded JSON value.
// It implements Marshaler and Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.
<原文结束>

# <翻译开始>
// RawMessage 是一个原始编码的 JSON 值。
// 它实现了 Marshaler 和 Unmarshaler 接口，可用于延迟 JSON 解码或预计算 JSON 编码。
# <翻译结束>


<原文开始>
// Marshal adapts to json/encoding Marshal API.
//
// Marshal returns the JSON encoding of v, adapts to json/encoding Marshal API
// Refer to https://godoc.org/encoding/json#Marshal for more information.
<原文结束>

# <翻译开始>
// Marshal 适应 json/encoding 库的 Marshal API.
//
// Marshal 返回参数 v 的 JSON 编码结果，此方法适应于 json/encoding 库中的 Marshal API，
// 更多信息请参考 https://godoc.org/encoding/json#Marshal 。
# <翻译结束>


<原文开始>
// MarshalIndent same as json.MarshalIndent.
<原文结束>

# <翻译开始>
// MarshalIndent 与 json.MarshalIndent 功能相同。
# <翻译结束>


<原文开始>
// Unmarshal adapts to json/encoding Unmarshal API
//
// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
// Refer to https://godoc.org/encoding/json#Unmarshal for more information.
<原文结束>

# <翻译开始>
// Unmarshal 适应 json/encoding 的 Unmarshal API
//
// Unmarshal 解析 JSON 编码的数据，并将结果存储在 v 指向的值中。
// 有关更多信息，请参考 https://godoc.org/encoding/json#Unmarshal 。
# <翻译结束>


<原文开始>
// UnmarshalUseNumber decodes the json data bytes to target interface using number option.
<原文结束>

# <翻译开始>
// UnmarshalUseNumber 使用数字选项将json数据字节解码到目标接口。
# <翻译结束>


<原文开始>
// NewEncoder same as json.NewEncoder
<原文结束>

# <翻译开始>
// NewEncoder 与 json.NewEncoder 功能相同
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
// NewDecoder 适应于 json/stream 的 NewDecoder API。
//
// NewDecoder 函数返回一个从 r 读取数据的新解码器。
//
// 返回的不是一个 json/encoding 包中的 Decoder，而是一个自定义的 Decoder。
// 有关更多信息，请参考 https://godoc.org/encoding/json#NewDecoder 。
# <翻译结束>


<原文开始>
// Valid reports whether data is a valid JSON encoding.
<原文结束>

# <翻译开始>
// Valid 报告 data 是否为有效的 JSON 编码。
# <翻译结束>

