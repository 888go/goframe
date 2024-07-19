// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gjson

import (
	"bytes"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/util/gconv"
)

// Valid checks whether `data` is a valid JSON data type.
// The parameter `data` specifies the json format data, which can be either
// bytes or string type.
// ff:是否为有效json
// data:值
func Valid(data interface{}) bool {
	return json.Valid(gconv.Bytes(data))
}

// Marshal is alias of Encode in order to fit the habit of json.Marshal/Unmarshal functions.
// ff:Marshal别名
// v:
// marshaledBytes:
// err:
func Marshal(v interface{}) (marshaledBytes []byte, err error) {
	return Encode(v)
}

// MarshalIndent is alias of json.MarshalIndent in order to fit the habit of json.MarshalIndent function.
// ff:MarshalIndent别名
// v:
// prefix:
// indent:
// marshaledBytes:
// err:
func MarshalIndent(v interface{}, prefix, indent string) (marshaledBytes []byte, err error) {
	return json.MarshalIndent(v, prefix, indent)
}

// Unmarshal is alias of DecodeTo in order to fit the habit of json.Marshal/Unmarshal functions.
// ff:Unmarshal别名
// data:
// v:
// err:
func Unmarshal(data []byte, v interface{}) (err error) {
	return DecodeTo(data, v)
}

// Encode encodes any golang variable `value` to JSON bytes.
// ff:变量到json字节集
// value:值
func Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

// MustEncode performs as Encode, but it panics if any error occurs.
// ff:变量到json字节集PANI
// value:值
func MustEncode(value interface{}) []byte {
	b, err := Encode(value)
	if err != nil {
		panic(err)
	}
	return b
}

// EncodeString encodes any golang variable `value` to JSON string.
// ff:变量到json文本
// value:值
func EncodeString(value interface{}) (string, error) {
	b, err := json.Marshal(value)
	return string(b), err
}

// MustEncodeString encodes any golang variable `value` to JSON string.
// It panics if any error occurs.
// ff:变量到json文本PANI
// value:值
func MustEncodeString(value interface{}) string {
	return string(MustEncode(value))
}

// Decode decodes json format `data` to golang variable.
// The parameter `data` can be either bytes or string type.
// ff:Json格式到变量
// data:值
// options:选项
func Decode(data interface{}, options ...Options) (interface{}, error) {
	var value interface{}
	if err := DecodeTo(gconv.Bytes(data), &value, options...); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

// DecodeTo decodes json format `data` to specified golang variable `v`.
// The parameter `data` can be either bytes or string type.
// The parameter `v` should be a pointer type.
// ff:Json格式到变量指针
// data:值
// v:变量指针
// options:选项
// err:错误
func DecodeTo(data interface{}, v interface{}, options ...Options) (err error) {
	decoder := json.NewDecoder(bytes.NewReader(gconv.Bytes(data)))
	if len(options) > 0 {
		// The StrNumber option is for certain situations, not for all.
		// For example, it causes converting issue for other data formats, for example: yaml.
		if options[0].StrNumber {
			decoder.UseNumber()
		}
	}
	if err = decoder.Decode(v); err != nil {
		err = gerror.Wrap(err, `json Decode failed`)
	}
	return
}

// DecodeToJson codes json format `data` to a Json object.
// The parameter `data` can be either bytes or string type.
// ff:解码到json
// data:值
// options:选项
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
