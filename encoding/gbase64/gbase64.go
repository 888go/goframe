// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// gbase64 包提供了对 BASE64 编码/解码算法的有用API。. md5:0524d2c59dedacdf
package gbase64

import (
	"encoding/base64"
	"os"

	"github.com/gogf/gf/v2/errors/gerror"
)

// Encode 使用BASE64算法对字节进行编码。. md5:9148f0cc83085b94
func Encode(src []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

// EncodeString 使用BASE64算法对字符串进行编码。. md5:3fd8a2a5fa2a8981
func EncodeString(src string) string {
	return EncodeToString([]byte(src))
}

// EncodeToString 使用BASE64算法将字节编码为字符串。. md5:53ccd7caaa813781
func EncodeToString(src []byte) string {
	return string(Encode(src))
}

// EncodeFile 使用BASE64算法对路径`path`指向的文件内容进行编码。. md5:66b77bf9fafafe1f
func EncodeFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		err = gerror.Wrapf(err, `os.ReadFile failed for filename "%s"`, path)
		return nil, err
	}
	return Encode(content), nil
}

// MustEncodeFile 使用BASE64算法对`path`路径的文件内容进行编码。
// 如果出现任何错误，它将引发 panic。
// md5:d31d543637e7eb6a
func MustEncodeFile(path string) []byte {
	result, err := EncodeFile(path)
	if err != nil {
		panic(err)
	}
	return result
}

// EncodeFileToString 使用BASE64算法将`path`路径下的文件内容编码为字符串。. md5:06aa514a3aa28d74
func EncodeFileToString(path string) (string, error) {
	content, err := EncodeFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// MustEncodeFileToString 使用BASE64算法将路径`path`中的文件内容编码为字符串。如果发生任何错误，它会引发恐慌。
// md5:f3091087e9edde81
func MustEncodeFileToString(path string) string {
	result, err := EncodeFileToString(path)
	if err != nil {
		panic(err)
	}
	return result
}

// Decode 使用BASE64算法对字节进行解码。. md5:dc5d8eec3a1bfb59
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
// 如果出现任何错误，它将引发恐慌。
// md5:92f5fbea62981c7a
func MustDecode(data []byte) []byte {
	result, err := Decode(data)
	if err != nil {
		panic(err)
	}
	return result
}

// DecodeString 使用BASE64算法解码字符串。. md5:8d6d8625cd51a554
func DecodeString(data string) ([]byte, error) {
	return Decode([]byte(data))
}

// MustDecodeString 使用BASE64算法解码字符串。
// 如果发生任何错误，它将引发恐慌。
// md5:215040f1e41de5af
func MustDecodeString(data string) []byte {
	result, err := DecodeString(data)
	if err != nil {
		panic(err)
	}
	return result
}

// DecodeToString 使用BASE64算法对字符串进行解码。. md5:3cff6b7395c38704
func DecodeToString(data string) (string, error) {
	b, err := DecodeString(data)
	return string(b), err
}

// MustDecodeToString 使用BASE64算法解码字符串。
// 如果发生错误，该函数将引发恐慌。
// md5:e2e017d04a30d409
func MustDecodeToString(data string) string {
	result, err := DecodeToString(data)
	if err != nil {
		panic(err)
	}
	return result
}
