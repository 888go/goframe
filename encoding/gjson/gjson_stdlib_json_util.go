// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gjson

import (
	"bytes"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/util/gconv"
)

// Valid 检查 `data` 是否是有效的 JSON 数据类型。参数 `data` 指定 JSON 格式的数据，可以是字节或字符串类型。
// md5:a1bbf790f78e4608
func Valid(data interface{}) bool {
	return json.Valid(gconv.Bytes(data))
}

// Marshal 是 Encode 函数的别名，目的是为了适应 json.Marshal/Unmarshal 函数的习惯用法。 md5:ff4e462ef9c849f2
func Marshal(v interface{}) (marshaledBytes []byte, err error) {
	return Encode(v)
}

// MarshalIndent 是 json.MarshalIndent 的别名，以适应使用 json.MarshalIndent 函数的习惯。 md5:285efc00996caf06
func MarshalIndent(v interface{}, prefix, indent string) (marshaledBytes []byte, err error) {
	return json.MarshalIndent(v, prefix, indent)
}

// Unmarshal是DecodeTo的别名，以适应json.Marshal/Unmarshal函数的习惯。 md5:dcb4ee0dfccdb10a
func Unmarshal(data []byte, v interface{}) (err error) {
	return DecodeTo(data, v)
}

// Encode 将任何 Go 语言变量 `value` 编码为 JSON 字节。 md5:25418d619ec52d3a
func Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

// MustEncode 的行为与 Encode 相同，但如果发生任何错误，它会直接 panic。 md5:baf6676afd45559a
func MustEncode(value interface{}) []byte {
	b, err := Encode(value)
	if err != nil {
		panic(err)
	}
	return b
}

// EncodeString 将任何 Go 语言变量 `value` 编码为 JSON 字符串。 md5:b54e604cb403b55a
func EncodeString(value interface{}) (string, error) {
	b, err := json.Marshal(value)
	return string(b), err
}

// MustEncodeString 将任何 Go 语言变量 `value` 编码为 JSON 字符串。如果发生任何错误，它将引发 panic。
// md5:05f6a19afa24c836
func MustEncodeString(value interface{}) string {
	return string(MustEncode(value))
}

// Decode 将 JSON 格式的 `data` 解码为 Go 语言变量。
// 参数 `data` 可以是字节切片或字符串类型。
// md5:8c3a611dab2c0896
func Decode(data interface{}, options ...Options) (interface{}, error) {
	var value interface{}
	if err := DecodeTo(gconv.Bytes(data), &value, options...); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

// DecodeTo 将json格式的 `data` 解码到指定的golang变量 `v`。
// 参数 `data` 可以是字节切片或字符串类型。
// 参数 `v` 应该是一个指针类型。
// md5:bc0dc16b58d95bda
func DecodeTo(data interface{}, v interface{}, options ...Options) (err error) {
	decoder := json.NewDecoder(bytes.NewReader(gconv.Bytes(data)))
	if len(options) > 0 {
		// StrNumber 选项适用于某些特定情况，而不是所有情况。
		// 例如，它会导致其他数据格式（如 YAML）的转换问题。
		// md5:304760f002a3649d
		if options[0].StrNumber {
			decoder.UseNumber()
		}
	}
	if err = decoder.Decode(v); err != nil {
		err = gerror.Wrap(err, `json Decode failed`)
	}
	return
}

// DecodeToJson 将JSON格式的`data`编码为一个Json对象。
// 参数`data`可以是字节或字符串类型。
// md5:f1745bf8c9553699
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
