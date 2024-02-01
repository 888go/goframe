// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gbase64 提供了用于 BASE64 编码/解码算法的有用 API。
package gbase64
import (
	"encoding/base64"
	"os"
	
	"github.com/888go/goframe/errors/gerror"
	)
// Encode 使用BASE64算法对字节进行编码。
func Encode(src []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

// EncodeString 使用BASE64算法对字符串进行编码。
func EncodeString(src string) string {
	return EncodeToString([]byte(src))
}

// EncodeToString 使用BASE64算法将字节编码为字符串。
func EncodeToString(src []byte) string {
	return string(Encode(src))
}

// EncodeFile 使用BASE64算法对`path`指定的文件内容进行编码。
func EncodeFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		err = gerror.Wrapf(err, `os.ReadFile failed for filename "%s"`, path)
		return nil, err
	}
	return Encode(content), nil
}

// MustEncodeFile 使用BASE64算法对`path`指定文件的内容进行编码。
// 如果出现任何错误，将会导致panic。
func MustEncodeFile(path string) []byte {
	result, err := EncodeFile(path)
	if err != nil {
		panic(err)
	}
	return result
}

// EncodeFileToString 使用BASE64算法将`path`指定文件的内容编码为字符串。
func EncodeFileToString(path string) (string, error) {
	content, err := EncodeFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// MustEncodeFileToString 使用BASE64算法将`path`指定文件的内容编码为字符串。
// 如果发生任何错误，它将会引发panic。
func MustEncodeFileToString(path string) string {
	result, err := EncodeFileToString(path)
	if err != nil {
		panic(err)
	}
	return result
}

// Decode 使用BASE64算法解码字节。
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

// MustDecode 使用BASE64算法解码字节。
// 如果发生任何错误，它会触发panic。
func MustDecode(data []byte) []byte {
	result, err := Decode(data)
	if err != nil {
		panic(err)
	}
	return result
}

// DecodeString 使用BASE64算法解码字符串。
func DecodeString(data string) ([]byte, error) {
	return Decode([]byte(data))
}

// MustDecodeString 使用BASE64算法解码字符串。
// 如果发生任何错误，它会引发panic。
func MustDecodeString(data string) []byte {
	result, err := DecodeString(data)
	if err != nil {
		panic(err)
	}
	return result
}

// DecodeToString 使用BASE64算法解码字符串。
func DecodeToString(data string) (string, error) {
	b, err := DecodeString(data)
	return string(b), err
}

// MustDecodeToString 使用BASE64算法解码字符串。
// 如果出现任何错误，将会导致panic。
func MustDecodeToString(data string) string {
	result, err := DecodeToString(data)
	if err != nil {
		panic(err)
	}
	return result
}
