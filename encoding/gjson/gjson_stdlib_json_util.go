// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package json类

import (
	"bytes"

	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	gconv "github.com/888go/goframe/util/gconv"
)

// X是否为有效json 检查 `data` 是否是有效的 JSON 数据类型。参数 `data` 指定 JSON 格式的数据，可以是字节或字符串类型。
// md5:a1bbf790f78e4608
func X是否为有效json(值 interface{}) bool {
	return json.Valid(gconv.X取字节集(值))
}

// Marshal别名 是 Encode 函数的别名，目的是为了适应 json.Marshal别名/Unmarshal 函数的习惯用法。 md5:ff4e462ef9c849f2
func Marshal别名(v interface{}) (marshaledBytes []byte, err error) {
	return X变量到json字节集(v)
}

// MarshalIndent别名 是 json.MarshalIndent别名 的别名，以适应使用 json.MarshalIndent别名 函数的习惯。 md5:285efc00996caf06
func MarshalIndent别名(v interface{}, prefix, indent string) (marshaledBytes []byte, err error) {
	return json.MarshalIndent(v, prefix, indent)
}

// Unmarshal别名是DecodeTo的别名，以适应json.Marshal/Unmarshal别名函数的习惯。 md5:dcb4ee0dfccdb10a
func Unmarshal别名(data []byte, v interface{}) (err error) {
	return Json格式到变量指针(data, v)
}

// X变量到json字节集 将任何 Go 语言变量 `value` 编码为 JSON 字节。 md5:25418d619ec52d3a
func X变量到json字节集(值 interface{}) ([]byte, error) {
	return json.Marshal(值)
}

// X变量到json字节集PANI 的行为与 Encode 相同，但如果发生任何错误，它会直接 panic。 md5:baf6676afd45559a
func X变量到json字节集PANI(值 interface{}) []byte {
	b, err := X变量到json字节集(值)
	if err != nil {
		panic(err)
	}
	return b
}

// X变量到json文本 将任何 Go 语言变量 `value` 编码为 JSON 字符串。 md5:b54e604cb403b55a
func X变量到json文本(值 interface{}) (string, error) {
	b, err := json.Marshal(值)
	return string(b), err
}

// X变量到json文本PANI 将任何 Go 语言变量 `value` 编码为 JSON 字符串。如果发生任何错误，它将引发 panic。
// md5:05f6a19afa24c836
func X变量到json文本PANI(值 interface{}) string {
	return string(X变量到json字节集PANI(值))
}

// Json格式到变量 将 JSON 格式的 `data` 解码为 Go 语言变量。
// 参数 `data` 可以是字节切片或字符串类型。
// md5:8c3a611dab2c0896
func Json格式到变量(值 interface{}, 选项 ...Options) (interface{}, error) {
	var value interface{}
	if err := Json格式到变量指针(gconv.X取字节集(值), &value, 选项...); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

// Json格式到变量指针 将json格式的 `data` 解码到指定的golang变量 `v`。
// 参数 `data` 可以是字节切片或字符串类型。
// 参数 `v` 应该是一个指针类型。
// md5:bc0dc16b58d95bda
func Json格式到变量指针(值 interface{}, 变量指针 interface{}, 选项 ...Options) (错误 error) {
	decoder := json.NewDecoder(bytes.NewReader(gconv.X取字节集(值)))
	if len(选项) > 0 {
		// StrNumber 选项适用于某些特定情况，而不是所有情况。
		// 例如，它会导致其他数据格式（如 YAML）的转换问题。
		// md5:304760f002a3649d
		if 选项[0].StrNumber {
			decoder.UseNumber()
		}
	}
	if 错误 = decoder.Decode(变量指针); 错误 != nil {
		错误 = gerror.X多层错误(错误, `json Decode failed`)
	}
	return
}

// X解码到json 将JSON格式的`data`编码为一个Json对象。
// 参数`data`可以是字节或字符串类型。
// md5:f1745bf8c9553699
func X解码到json(值 interface{}, 选项 ...Options) (*Json, error) {
	if v, err := Json格式到变量(gconv.X取字节集(值), 选项...); err != nil {
		return nil, err
	} else {
		if len(选项) > 0 {
			return X创建(v, 选项[0].Safe), nil
		}
		return X创建(v), nil
	}
}
