// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gbase64 provides useful API for BASE64 encoding/decoding algorithm.
package gbase64//bm:编码base64类

import (
	"encoding/base64"
	"os"

	"github.com/gogf/gf/v2/errors/gerror"
)

// Encode encodes bytes with BASE64 algorithm.

// ff:字节集编码
// src:字节集
func Encode(src []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

// EncodeString encodes string with BASE64 algorithm.

// ff:文本编码
// src:文本
func EncodeString(src string) string {
	return EncodeToString([]byte(src))
}

// EncodeToString encodes bytes to string with BASE64 algorithm.

// ff:字节集编码到文本
// src:字节集
func EncodeToString(src []byte) string {
	return string(Encode(src))
}

// EncodeFile encodes file content of `path` using BASE64 algorithms.

// ff:文件编码到字节集
// path:文件路径
func EncodeFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		err = gerror.Wrapf(err, `os.ReadFile failed for filename "%s"`, path)
		return nil, err
	}
	return Encode(content), nil
}

// MustEncodeFile encodes file content of `path` using BASE64 algorithms.
// It panics if any error occurs.

// ff:文件编码到字节集PANI
// path:文件路径
func MustEncodeFile(path string) []byte {
	result, err := EncodeFile(path)
	if err != nil {
		panic(err)
	}
	return result
}

// EncodeFileToString encodes file content of `path` to string using BASE64 algorithms.

// ff:文件编码到文本
// path:文件路径
func EncodeFileToString(path string) (string, error) {
	content, err := EncodeFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// MustEncodeFileToString encodes file content of `path` to string using BASE64 algorithms.
// It panics if any error occurs.

// ff:文件编码到文本PANI
// path:文件路径
func MustEncodeFileToString(path string) string {
	result, err := EncodeFileToString(path)
	if err != nil {
		panic(err)
	}
	return result
}

// Decode decodes bytes with BASE64 algorithm.

// ff:字节集解码
// data:字节集
func Decode(data []byte) ([]byte, error) {
	var (
		src    = make([]byte, base64.StdEncoding.DecodedLen(len(data)))
		n, err = base64.StdEncoding.Decode(src, data)
	)
	if err != nil {
		err = gerror.Wrap(err, `base64.StdEncoding.Decode failed`)
	}
	return src[:n], err
}

// MustDecode decodes bytes with BASE64 algorithm.
// It panics if any error occurs.

// ff:字节集解码PANI
// data:字节集
func MustDecode(data []byte) []byte {
	result, err := Decode(data)
	if err != nil {
		panic(err)
	}
	return result
}

// DecodeString decodes string with BASE64 algorithm.

// ff:文本解码到字节集
// data:文本
func DecodeString(data string) ([]byte, error) {
	return Decode([]byte(data))
}

// MustDecodeString decodes string with BASE64 algorithm.
// It panics if any error occurs.

// ff:文本解码到字节集PANI
// data:文本
func MustDecodeString(data string) []byte {
	result, err := DecodeString(data)
	if err != nil {
		panic(err)
	}
	return result
}

// DecodeToString decodes string with BASE64 algorithm.

// ff:文本解码
// data:文本
func DecodeToString(data string) (string, error) {
	b, err := DecodeString(data)
	return string(b), err
}

// MustDecodeToString decodes string with BASE64 algorithm.
// It panics if any error occurs.

// ff:文本解码PANI
// data:文本
func MustDecodeToString(data string) string {
	result, err := DecodeToString(data)
	if err != nil {
		panic(err)
	}
	return result
}
