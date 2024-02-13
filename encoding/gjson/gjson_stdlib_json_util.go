// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package json类

import (
	"bytes"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/util/gconv"
)

// Valid 检查 `data` 是否为有效的 JSON 数据类型。
// 参数 `data` 指定了 json 格式的数据，可以是字节切片类型或字符串类型。
func X是否为有效json(值 interface{}) bool {
	return json.Valid(转换类.X取字节集(值))
}

// Marshal 是 Encode 的别名，以便与 json.Marshal/Unmarshal 函数的习惯用法保持一致。
func Marshal别名(v interface{}) (marshaledBytes []byte, err error) {
	return X变量到json字节集(v)
}

// MarshalIndent 是 json.MarshalIndent 函数的别名，目的是为了符合使用 json.MarshalIndent 函数的习惯。
func MarshalIndent别名(v interface{}, prefix, indent string) (marshaledBytes []byte, err error) {
	return json.MarshalIndent(v, prefix, indent)
}

// Unmarshal 是 DecodeTo 的别名，目的是为了符合 json.Marshal/Unmarshal 函数的习惯用法。
func Unmarshal别名(data []byte, v interface{}) (err error) {
	return Json格式到变量指针(data, v)
}

// Encode 将任何 Go 语言变量 `value` 编码为 JSON 字节。
func X变量到json字节集(值 interface{}) ([]byte, error) {
	return json.Marshal(值)
}

// MustEncode 的行为与 Encode 相同，但当发生任何错误时，它会触发 panic（异常）。
func X变量到json字节集PANI(值 interface{}) []byte {
	b, err := X变量到json字节集(值)
	if err != nil {
		panic(err)
	}
	return b
}

// EncodeString 将任意的 Go 语言变量 `value` 编码为 JSON 字符串。
func X变量到json文本(值 interface{}) (string, error) {
	b, err := json.Marshal(值)
	return string(b), err
}

// MustEncodeString 将任何 Go 语言变量 `value` 编码为 JSON 字符串。
// 如果出现任何错误，它将引发恐慌（panic）。
func X变量到json文本PANI(值 interface{}) string {
	return string(X变量到json字节集PANI(值))
}

// Decode 解码 JSON 格式的 `data` 为 Go 语言变量。
// 参数 `data` 可以是字节切片（bytes）类型或字符串（string）类型。
func Json格式到变量(值 interface{}, 选项 ...Options) (interface{}, error) {
	var value interface{}
	if err := Json格式到变量指针(转换类.X取字节集(值), &value, 选项...); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

// DecodeTo 将 json 格式的 `data` 解码到指定的 golang 变量 `v`。
// 参数 `data` 可以是 bytes 类型或 string 类型。
// 参数 `v` 应该是指针类型。
func Json格式到变量指针(值 interface{}, 变量指针 interface{}, 选项 ...Options) (错误 error) {
	decoder := json.NewDecoder(bytes.NewReader(转换类.X取字节集(值)))
	if len(选项) > 0 {
// StrNumber 选项适用于某些特定场景，但并不适用于所有场景。
// 例如，它可能会导致与其他数据格式（如 yaml）转换时出现问题。
		if 选项[0].StrNumber {
			decoder.UseNumber()
		}
	}
	if 错误 = decoder.Decode(变量指针); 错误 != nil {
		错误 = 错误类.X多层错误(错误, `json Decode failed`)
	}
	return
}

// DecodeToJson将json格式的`data`解码成一个Json对象。
// 参数`data`可以是字节类型或字符串类型。
func X解码到json(值 interface{}, 选项 ...Options) (*Json, error) {
	if v, err := Json格式到变量(转换类.X取字节集(值), 选项...); err != nil {
		return nil, err
	} else {
		if len(选项) > 0 {
			return X创建(v, 选项[0].Safe), nil
		}
		return X创建(v), nil
	}
}
