// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package json 提供了围绕标准库或第三方库的 JSON 操作，实现了对 JSON 的封装并忽略它们。
package json

import (
	"bytes"
	"encoding/json"
	"io"
	
	"github.com/888go/goframe/errors/gerror"
)

// RawMessage 是一个原始编码的 JSON 值。
// 它实现了 Marshaler 和 Unmarshaler 接口，可用于延迟 JSON 解码或预计算 JSON 编码。
type RawMessage = json.RawMessage

// Marshal 适应 json/encoding 库的 Marshal API.
//
// Marshal 返回参数 v 的 JSON 编码结果，此方法适应于 json/encoding 库中的 Marshal API，
// 更多信息请参考 https://godoc.org/encoding/json#Marshal 。
func Marshal(v interface{}) (marshaledBytes []byte, err error) {
	marshaledBytes, err = json.Marshal(v)
	if err != nil {
		err = 错误类.X多层错误(err, `json.Marshal failed`)
	}
	return
}

// MarshalIndent 与 json.MarshalIndent 功能相同。
func MarshalIndent(v interface{}, prefix, indent string) (marshaledBytes []byte, err error) {
	marshaledBytes, err = json.MarshalIndent(v, prefix, indent)
	if err != nil {
		err = 错误类.X多层错误(err, `json.MarshalIndent failed`)
	}
	return
}

// Unmarshal 适应 json/encoding 的 Unmarshal API
//
// Unmarshal 解析 JSON 编码的数据，并将结果存储在 v 指向的值中。
// 有关更多信息，请参考 https://godoc.org/encoding/json#Unmarshal 。
func Unmarshal(data []byte, v interface{}) (err error) {
	err = json.Unmarshal(data, v)
	if err != nil {
		err = 错误类.X多层错误(err, `json.Unmarshal failed`)
	}
	return
}

// UnmarshalUseNumber 使用数字选项将json数据字节解码到目标接口。
func UnmarshalUseNumber(data []byte, v interface{}) (err error) {
	decoder := NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	err = decoder.Decode(v)
	if err != nil {
		err = 错误类.X多层错误(err, `json.UnmarshalUseNumber failed`)
	}
	return
}

// NewEncoder 与 json.NewEncoder 功能相同
func NewEncoder(writer io.Writer) *json.Encoder {
	return json.NewEncoder(writer)
}

// NewDecoder 适应于 json/stream 的 NewDecoder API。
//
// NewDecoder 函数返回一个从 r 读取数据的新解码器。
//
// 返回的不是一个 json/encoding 包中的 Decoder，而是一个自定义的 Decoder。
// 有关更多信息，请参考 https://godoc.org/encoding/json#NewDecoder 。
func NewDecoder(reader io.Reader) *json.Decoder {
	return json.NewDecoder(reader)
}

// Valid 报告 data 是否为有效的 JSON 编码。
func Valid(data []byte) bool {
	return json.Valid(data)
}
