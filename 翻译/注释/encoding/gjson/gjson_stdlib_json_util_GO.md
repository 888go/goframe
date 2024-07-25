
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
// Valid checks whether `data` is a valid JSON data type.
// The parameter `data` specifies the json format data, which can be either
// bytes or string type.
<原文结束>

# <翻译开始>
// Valid 检查 `data` 是否是有效的 JSON 数据类型。参数 `data` 指定 JSON 格式的数据，可以是字节或字符串类型。
// md5:a1bbf790f78e4608
# <翻译结束>


<原文开始>
// Marshal is alias of Encode in order to fit the habit of json.Marshal/Unmarshal functions.
<原文结束>

# <翻译开始>
// Marshal 是 Encode 函数的别名，目的是为了适应 json.Marshal/Unmarshal 函数的习惯用法。 md5:ff4e462ef9c849f2
# <翻译结束>


<原文开始>
// MarshalIndent is alias of json.MarshalIndent in order to fit the habit of json.MarshalIndent function.
<原文结束>

# <翻译开始>
// MarshalIndent 是 json.MarshalIndent 的别名，以适应使用 json.MarshalIndent 函数的习惯。 md5:285efc00996caf06
# <翻译结束>


<原文开始>
// Unmarshal is alias of DecodeTo in order to fit the habit of json.Marshal/Unmarshal functions.
<原文结束>

# <翻译开始>
// Unmarshal是DecodeTo的别名，以适应json.Marshal/Unmarshal函数的习惯。 md5:dcb4ee0dfccdb10a
# <翻译结束>


<原文开始>
// Encode encodes any golang variable `value` to JSON bytes.
<原文结束>

# <翻译开始>
// Encode 将任何 Go 语言变量 `value` 编码为 JSON 字节。 md5:25418d619ec52d3a
# <翻译结束>


<原文开始>
// MustEncode performs as Encode, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustEncode 的行为与 Encode 相同，但如果发生任何错误，它会直接 panic。 md5:baf6676afd45559a
# <翻译结束>


<原文开始>
// EncodeString encodes any golang variable `value` to JSON string.
<原文结束>

# <翻译开始>
// EncodeString 将任何 Go 语言变量 `value` 编码为 JSON 字符串。 md5:b54e604cb403b55a
# <翻译结束>


<原文开始>
// MustEncodeString encodes any golang variable `value` to JSON string.
// It panics if any error occurs.
<原文结束>

# <翻译开始>
// MustEncodeString 将任何 Go 语言变量 `value` 编码为 JSON 字符串。如果发生任何错误，它将引发 panic。
// md5:05f6a19afa24c836
# <翻译结束>


<原文开始>
// Decode decodes json format `data` to golang variable.
// The parameter `data` can be either bytes or string type.
<原文结束>

# <翻译开始>
// Decode 将 JSON 格式的 `data` 解码为 Go 语言变量。
// 参数 `data` 可以是字节切片或字符串类型。
// md5:8c3a611dab2c0896
# <翻译结束>


<原文开始>
// DecodeTo decodes json format `data` to specified golang variable `v`.
// The parameter `data` can be either bytes or string type.
// The parameter `v` should be a pointer type.
<原文结束>

# <翻译开始>
// DecodeTo 将json格式的 `data` 解码到指定的golang变量 `v`。
// 参数 `data` 可以是字节切片或字符串类型。
// 参数 `v` 应该是一个指针类型。
// md5:bc0dc16b58d95bda
# <翻译结束>


<原文开始>
		// The StrNumber option is for certain situations, not for all.
		// For example, it causes converting issue for other data formats, for example: yaml.
<原文结束>

# <翻译开始>
		// StrNumber 选项适用于某些特定情况，而不是所有情况。
		// 例如，它会导致其他数据格式（如 YAML）的转换问题。
		// md5:304760f002a3649d
# <翻译结束>


<原文开始>
// DecodeToJson codes json format `data` to a Json object.
// The parameter `data` can be either bytes or string type.
<原文结束>

# <翻译开始>
// DecodeToJson 将JSON格式的`data`编码为一个Json对象。
// 参数`data`可以是字节或字符串类型。
// md5:f1745bf8c9553699
# <翻译结束>

