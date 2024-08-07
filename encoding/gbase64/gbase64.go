// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// gbase64 包提供了对 BASE64 编码/解码算法的有用API。 md5:0524d2c59dedacdf
package 编码base64类

import (
	"encoding/base64"
	"os"

	gerror "github.com/888go/goframe/errors/gerror"
)

// X字节集编码 使用BASE64算法对字节进行编码。 md5:9148f0cc83085b94
func X字节集编码(字节集 []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(字节集)))
	base64.StdEncoding.Encode(dst, 字节集)
	return dst
}

// X文本编码 使用BASE64算法对字符串进行编码。 md5:3fd8a2a5fa2a8981
func X文本编码(文本 string) string {
	return X字节集编码到文本([]byte(文本))
}

// X字节集编码到文本 使用BASE64算法将字节编码为字符串。 md5:53ccd7caaa813781
func X字节集编码到文本(字节集 []byte) string {
	return string(X字节集编码(字节集))
}

// X文件编码到字节集 使用BASE64算法对路径`path`指向的文件内容进行编码。 md5:66b77bf9fafafe1f
func X文件编码到字节集(文件路径 string) ([]byte, error) {
	content, err := os.ReadFile(文件路径)
	if err != nil {
		err = gerror.X多层错误并格式化(err, `os.ReadFile failed for filename "%s"`, 文件路径)
		return nil, err
	}
	return X字节集编码(content), nil
}

// X文件编码到字节集PANI 使用BASE64算法对`path`路径的文件内容进行编码。
// 如果出现任何错误，它将引发 panic。
// md5:d31d543637e7eb6a
func X文件编码到字节集PANI(文件路径 string) []byte {
	result, err := X文件编码到字节集(文件路径)
	if err != nil {
		panic(err)
	}
	return result
}

// X文件编码到文本 使用BASE64算法将`path`路径下的文件内容编码为字符串。 md5:06aa514a3aa28d74
func X文件编码到文本(文件路径 string) (string, error) {
	content, err := X文件编码到字节集(文件路径)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// X文件编码到文本PANI 使用BASE64算法将路径`path`中的文件内容编码为字符串。如果发生任何错误，它会引发恐慌。
// md5:f3091087e9edde81
func X文件编码到文本PANI(文件路径 string) string {
	result, err := X文件编码到文本(文件路径)
	if err != nil {
		panic(err)
	}
	return result
}

// X字节集解码 使用BASE64算法对字节进行解码。 md5:dc5d8eec3a1bfb59
func X字节集解码(字节集 []byte) ([]byte, error) {
	var (
		src    = make([]byte, base64.StdEncoding.DecodedLen(len(字节集)))
		n, err = base64.StdEncoding.Decode(src, 字节集)
	)
	if err != nil {
		err = gerror.X多层错误(err, `base64.StdEncoding.Decode failed`)
	}
	return src[:n], err
}

// X字节集解码PANI 使用BASE64算法解码字节。
// 如果出现任何错误，它将引发恐慌。
// md5:92f5fbea62981c7a
func X字节集解码PANI(字节集 []byte) []byte {
	result, err := X字节集解码(字节集)
	if err != nil {
		panic(err)
	}
	return result
}

// X文本解码到字节集 使用BASE64算法解码字符串。 md5:8d6d8625cd51a554
func X文本解码到字节集(文本 string) ([]byte, error) {
	return X字节集解码([]byte(文本))
}

// X文本解码到字节集PANI 使用BASE64算法解码字符串。
// 如果发生任何错误，它将引发恐慌。
// md5:215040f1e41de5af
func X文本解码到字节集PANI(文本 string) []byte {
	result, err := X文本解码到字节集(文本)
	if err != nil {
		panic(err)
	}
	return result
}

// X文本解码 使用BASE64算法对字符串进行解码。 md5:3cff6b7395c38704
func X文本解码(文本 string) (string, error) {
	b, err := X文本解码到字节集(文本)
	return string(b), err
}

// X文本解码PANI 使用BASE64算法解码字符串。
// 如果发生错误，该函数将引发恐慌。
// md5:e2e017d04a30d409
func X文本解码PANI(文本 string) string {
	result, err := X文本解码(文本)
	if err != nil {
		panic(err)
	}
	return result
}
