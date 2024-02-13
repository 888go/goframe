// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gsha1 提供了用于SHA1加密算法的有用API。
package 加密sha1类

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/util/gconv"
)

// Encrypt 使用SHA1算法对任意类型的变量进行加密。
// 它使用gconv包将`v`转换为字节类型。
func X加密(值 interface{}) string {
	r := sha1.Sum(转换类.X取字节集(值))
	return hex.EncodeToString(r[:])
}

// EncryptFile 使用SHA1算法加密`path`指定文件的内容。
func X加密文件(路径 string) (sha1值 string, 错误 error) {
	f, 错误 := os.Open(路径)
	if 错误 != nil {
		错误 = 错误类.X多层错误并格式化(错误, `os.Open failed for name "%s"`, 路径)
		return "", 错误
	}
	defer f.Close()
	h := sha1.New()
	_, 错误 = io.Copy(h, f)
	if 错误 != nil {
		错误 = 错误类.X多层错误(错误, `io.Copy failed`)
		return "", 错误
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// MustEncryptFile 使用SHA1算法加密`path`指定文件的内容。
// 如果出现任何错误，将会触发panic。
func X加密文件PANI(路径 string) string {
	result, err := X加密文件(路径)
	if err != nil {
		panic(err)
	}
	return result
}
