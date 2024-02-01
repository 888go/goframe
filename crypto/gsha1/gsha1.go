// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gsha1 提供了用于SHA1加密算法的有用API。
package gsha1
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
func Encrypt(v interface{}) string {
	r := sha1.Sum(gconv.Bytes(v))
	return hex.EncodeToString(r[:])
}

// EncryptFile 使用SHA1算法加密`path`指定文件的内容。
func EncryptFile(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		err = gerror.Wrapf(err, `os.Open failed for name "%s"`, path)
		return "", err
	}
	defer f.Close()
	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		err = gerror.Wrap(err, `io.Copy failed`)
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// MustEncryptFile 使用SHA1算法加密`path`指定文件的内容。
// 如果出现任何错误，将会触发panic。
func MustEncryptFile(path string) string {
	result, err := EncryptFile(path)
	if err != nil {
		panic(err)
	}
	return result
}
