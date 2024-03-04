// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gjson

import (
	"bytes"
	
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/888go/goframe/gjson/internal/json"
	"github.com/gogf/gf/v2/util/gconv"
)

// Valid 检查 `data` 是否为有效的 JSON 数据类型。
// 参数 `data` 指定了 json 格式的数据，可以是字节切片类型或字符串类型。
func Valid(data interface{}) bool {
	return json.Valid(gconv.Bytes(data))
}

// Marshal 是 Encode 的别名，以便与 json.Marshal/Unmarshal 函数的习惯用法保持一致。
func Marshal(v interface{}) (marshaledBytes []byte, err error) {
	return Encode(v)
}

// MarshalIndent 是 json.MarshalIndent 函数的别名，目的是为了符合使用 json.MarshalIndent 函数的习惯。
func MarshalIndent(v interface{}, prefix, indent string) (marshaledBytes []byte, err error) {
	return json.MarshalIndent(v, prefix, indent)
}

// Unmarshal 是 DecodeTo 的别名，目的是为了符合 json.Marshal/Unmarshal 函数的习惯用法。
func Unmarshal(data []byte, v interface{}) (err error) {
	return DecodeTo(data, v)
}

// Encode 将任何 Go 语言变量 `value` 编码为 JSON 字节。
func Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

// MustEncode 的行为与 Encode 相同，但当发生任何错误时，它会触发 panic（异常）。
func MustEncode(value interface{}) []byte {
	b, err := Encode(value)
	if err != nil {
		panic(err)
	}
	return b
}

// EncodeString 将任意的 Go 语言变量 `value` 编码为 JSON 字符串。
func EncodeString(value interface{}) (string, error) {
	b, err := json.Marshal(value)
	return string(b), err
}

// MustEncodeString 将任何 Go 语言变量 `value` 编码为 JSON 字符串。
// 如果出现任何错误，它将引发恐慌（panic）。
func MustEncodeString(value interface{}) string {
	return string(MustEncode(value))
}

// Decode 解码 JSON 格式的 `data` 为 Go 语言变量。
// 参数 `data` 可以是字节切片（bytes）类型或字符串（string）类型。
func Decode(data interface{}, options ...Options) (interface{}, error) {
	var value interface{}
	if err := DecodeTo(gconv.Bytes(data), &value, options...); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

// DecodeTo 将 json 格式的 `data` 解码到指定的 golang 变量 `v`。
// 参数 `data` 可以是 bytes 类型或 string 类型。
// 参数 `v` 应该是指针类型。
func DecodeTo(data interface{}, v interface{}, options ...Options) (err error) {
	decoder := json.NewDecoder(bytes.NewReader(gconv.Bytes(data)))
	if len(options) > 0 {
// StrNumber 选项适用于某些特定场景，但并不适用于所有场景。
// 例如，它可能会导致与其他数据格式（如 yaml）转换时出现问题。
		if options[0].StrNumber {
			decoder.UseNumber()
		}
	}
	if err = decoder.Decode(v); err != nil {
		err = gerror.Wrap(err, `json Decode failed`)
	}
	return
}

// DecodeToJson将json格式的`data`解码成一个Json对象。
// 参数`data`可以是字节类型或字符串类型。
func DecodeToJson(data interface{}, options ...Options) (*Json, error) {
	if v, err := Decode(gconv.Bytes(data), options...); err != nil {
		return nil, err
	} else {
		if len(options) > 0 {
			return New(v, options[0].Safe), nil
		}
		return New(v), nil
	}
}
