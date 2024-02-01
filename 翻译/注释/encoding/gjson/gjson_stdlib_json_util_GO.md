
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
// Valid checks whether `data` is a valid JSON data type.
// The parameter `data` specifies the json format data, which can be either
// bytes or string type.
<原文结束>

# <翻译开始>
// Valid 检查 `data` 是否为有效的 JSON 数据类型。
// 参数 `data` 指定了 json 格式的数据，可以是字节切片类型或字符串类型。
# <翻译结束>


<原文开始>
// Marshal is alias of Encode in order to fit the habit of json.Marshal/Unmarshal functions.
<原文结束>

# <翻译开始>
// Marshal 是 Encode 的别名，以便与 json.Marshal/Unmarshal 函数的习惯用法保持一致。
# <翻译结束>


<原文开始>
// MarshalIndent is alias of json.MarshalIndent in order to fit the habit of json.MarshalIndent function.
<原文结束>

# <翻译开始>
// MarshalIndent 是 json.MarshalIndent 函数的别名，目的是为了符合使用 json.MarshalIndent 函数的习惯。
# <翻译结束>


<原文开始>
// Unmarshal is alias of DecodeTo in order to fit the habit of json.Marshal/Unmarshal functions.
<原文结束>

# <翻译开始>
// Unmarshal 是 DecodeTo 的别名，目的是为了符合 json.Marshal/Unmarshal 函数的习惯用法。
# <翻译结束>


<原文开始>
// Encode encodes any golang variable `value` to JSON bytes.
<原文结束>

# <翻译开始>
// Encode 将任何 Go 语言变量 `value` 编码为 JSON 字节。
# <翻译结束>


<原文开始>
// MustEncode performs as Encode, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustEncode 的行为与 Encode 相同，但当发生任何错误时，它会触发 panic（异常）。
# <翻译结束>


<原文开始>
// EncodeString encodes any golang variable `value` to JSON string.
<原文结束>

# <翻译开始>
// EncodeString 将任意的 Go 语言变量 `value` 编码为 JSON 字符串。
# <翻译结束>


<原文开始>
// MustEncodeString encodes any golang variable `value` to JSON string.
// It panics if any error occurs.
<原文结束>

# <翻译开始>
// MustEncodeString 将任何 Go 语言变量 `value` 编码为 JSON 字符串。
// 如果出现任何错误，它将引发恐慌（panic）。
# <翻译结束>


<原文开始>
// Decode decodes json format `data` to golang variable.
// The parameter `data` can be either bytes or string type.
<原文结束>

# <翻译开始>
// Decode 解码 JSON 格式的 `data` 为 Go 语言变量。
// 参数 `data` 可以是字节切片（bytes）类型或字符串（string）类型。
# <翻译结束>


<原文开始>
// DecodeTo decodes json format `data` to specified golang variable `v`.
// The parameter `data` can be either bytes or string type.
// The parameter `v` should be a pointer type.
<原文结束>

# <翻译开始>
// DecodeTo 将 json 格式的 `data` 解码到指定的 golang 变量 `v`。
// 参数 `data` 可以是 bytes 类型或 string 类型。
// 参数 `v` 应该是指针类型。
# <翻译结束>


<原文开始>
		// The StrNumber option is for certain situations, not for all.
		// For example, it causes converting issue for other data formats, for example: yaml.
<原文结束>

# <翻译开始>
// StrNumber 选项适用于某些特定场景，但并不适用于所有场景。
// 例如，它可能会导致与其他数据格式（如 yaml）转换时出现问题。
# <翻译结束>


<原文开始>
// DecodeToJson codes json format `data` to a Json object.
// The parameter `data` can be either bytes or string type.
<原文结束>

# <翻译开始>
// DecodeToJson将json格式的`data`解码成一个Json对象。
// 参数`data`可以是字节类型或字符串类型。
# <翻译结束>

