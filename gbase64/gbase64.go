// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gbase64 提供了用于 BASE64 编码/解码算法的有用 API。
package 编码base64类

import (
	"encoding/base64"
	"os"
	
	"github.com/gogf/gf/v2/errors/gerror"
)

// Encode 使用BASE64算法对字节进行编码。
func X字节集编码(字节集 []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(字节集)))
	base64.StdEncoding.Encode(dst, 字节集)
	return dst
}

// EncodeString 使用BASE64算法对字符串进行编码。
func X文本编码(文本 string) string {
	return X字节集编码到文本([]byte(文本))
}

// EncodeToString 使用BASE64算法将字节编码为字符串。
func X字节集编码到文本(字节集 []byte) string {
	return string(X字节集编码(字节集))
}

// EncodeFile 使用BASE64算法对`path`指定的文件内容进行编码。
func X文件编码到字节集(文件路径 string) ([]byte, error) {
	content, err := os.ReadFile(文件路径)
	if err != nil {
		err = gerror.Wrapf(err, `os.ReadFile failed for filename "%s"`, 文件路径)
		return nil, err
	}
	return X字节集编码(content), nil
}

// MustEncodeFile 使用BASE64算法对`path`指定文件的内容进行编码。
// 如果出现任何错误，将会导致panic。
func X文件编码到字节集PANI(文件路径 string) []byte {
	result, err := X文件编码到字节集(文件路径)
	if err != nil {
		panic(err)
	}
	return result
}

// EncodeFileToString 使用BASE64算法将`path`指定文件的内容编码为字符串。
func X文件编码到文本(文件路径 string) (string, error) {
	content, err := X文件编码到字节集(文件路径)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// MustEncodeFileToString 使用BASE64算法将`path`指定文件的内容编码为字符串。
// 如果发生任何错误，它将会引发panic。
func X文件编码到文本PANI(文件路径 string) string {
	result, err := X文件编码到文本(文件路径)
	if err != nil {
		panic(err)
	}
	return result
}

// Decode 使用BASE64算法解码字节。
func X字节集解码(字节集 []byte) ([]byte, error) {
	var (
		src    = make([]byte, base64.StdEncoding.DecodedLen(len(字节集)))
		n, err = base64.StdEncoding.Decode(src, 字节集)
	)
	if err != nil {
		err = gerror.Wrap(err, `base64.StdEncoding.Decode failed`)
	}
	return src[:n], err
}

// MustDecode 使用BASE64算法解码字节。
// 如果发生任何错误，它会触发panic。
func X字节集解码PANI(字节集 []byte) []byte {
	result, err := X字节集解码(字节集)
	if err != nil {
		panic(err)
	}
	return result
}

// DecodeString 使用BASE64算法解码字符串。
func X文本解码到字节集(文本 string) ([]byte, error) {
	return X字节集解码([]byte(文本))
}

// MustDecodeString 使用BASE64算法解码字符串。
// 如果发生任何错误，它会引发panic。
func X文本解码到字节集PANI(文本 string) []byte {
	result, err := X文本解码到字节集(文本)
	if err != nil {
		panic(err)
	}
	return result
}

// DecodeToString 使用BASE64算法解码字符串。
func X文本解码(文本 string) (string, error) {
	b, err := X文本解码到字节集(文本)
	return string(b), err
}

// MustDecodeToString 使用BASE64算法解码字符串。
// 如果出现任何错误，将会导致panic。
func X文本解码PANI(文本 string) string {
	result, err := X文本解码(文本)
	if err != nil {
		panic(err)
	}
	return result
}
