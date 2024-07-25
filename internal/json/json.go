// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// json包提供了对json操作的封装，忽略了标准库或第三方库的json。 md5:8c700638d650aacd
package json

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gogf/gf/v2/errors/gerror"
)

// RawMessage 是一个原始的编码JSON值。
// 它实现了Marshaler和Unmarshaler接口，可以用于延迟JSON解码或预先计算JSON编码。 md5:2ea51fc4bfe3af87
type RawMessage = json.RawMessage

// Marshal 适应了json/encoding.Marshal API。
//
// Marshal 返回 v 的 JSON 编码，适应了 json/encoding.Marshal API。更多信息请参考 https://godoc.org/encoding/json#Marshal。 md5:e67e9d2efbcb1d3c
func Marshal(v interface{}) (marshaledBytes []byte, err error) {
	marshaledBytes, err = json.Marshal(v)
	if err != nil {
		err = gerror.Wrap(err, `json.Marshal failed`)
	}
	return
}

// MarshalIndent 和 json.MarshalIndent 功能相同。 md5:1c3ee8bca0354fac
func MarshalIndent(v interface{}, prefix, indent string) (marshaledBytes []byte, err error) {
	marshaledBytes, err = json.MarshalIndent(v, prefix, indent)
	if err != nil {
		err = gerror.Wrap(err, `json.MarshalIndent failed`)
	}
	return
}

// Unmarshal 适应于 json/encoding 的 Unmarshal API
//
// Unmarshal 解析 JSON 编码的数据，并将结果存储到由 v 指向的值中。
// 更多信息，请参考 https://godoc.org/encoding/json#Unmarshal。 md5:7272cfbd647a7f0f
func Unmarshal(data []byte, v interface{}) (err error) {
	err = json.Unmarshal(data, v)
	if err != nil {
		err = gerror.Wrap(err, `json.Unmarshal failed`)
	}
	return
}

// UnmarshalUseNumber 使用数字选项将json数据字节解码到目标接口。 md5:b04f82165b9c9933
func UnmarshalUseNumber(data []byte, v interface{}) (err error) {
	decoder := NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	err = decoder.Decode(v)
	if err != nil {
		err = gerror.Wrap(err, `json.UnmarshalUseNumber failed`)
	}
	return
}

// NewEncoder 等同于 json.NewEncoder. md5:4f4b9a0b5e712ba5
func NewEncoder(writer io.Writer) *json.Encoder {
	return json.NewEncoder(writer)
}

// NewDecoder 封装了 json/stream 中的 NewDecoder API。
//
// NewDecoder 从 r 读取数据，返回一个新的解码器。
//
// 返回的不是一个 json/encoding 中的 Decoder，而是 Decoder。更多详细信息请参考：https://godoc.org/encoding/json#NewDecoder。 md5:eea702294e95e71d
func NewDecoder(reader io.Reader) *json.Decoder {
	return json.NewDecoder(reader)
}

// Valid 报告数据是否为有效的JSON编码。 md5:db76d6317d6c0b76
func Valid(data []byte) bool {
	return json.Valid(data)
}
