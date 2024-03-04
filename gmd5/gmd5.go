// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gmd5提供了MD5加密算法的实用API。
package gmd5

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

// Encrypt 使用MD5算法加密任意类型的变量。
// 它使用gconv包将`v`转换为字节类型。
func Encrypt(data interface{}) (encrypt string, err error) {
	return EncryptBytes(gconv.Bytes(data))
}

// MustEncrypt 使用MD5算法对任意类型的变量进行加密。
// 它使用gconv包将`v`转换为字节类型。
// 如果发生任何错误，它会引发panic。
func MustEncrypt(data interface{}) string {
	result, err := Encrypt(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptBytes 使用MD5算法加密`data`。
func EncryptBytes(data []byte) (encrypt string, err error) {
	h := md5.New()
	if _, err = h.Write(data); err != nil {
		err = gerror.Wrap(err, `hash.Write failed`)
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// MustEncryptBytes 使用MD5算法加密`data`。
// 如果出现任何错误，将会导致panic。
func MustEncryptBytes(data []byte) string {
	result, err := EncryptBytes(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptString 使用MD5算法加密字符串`data`。
func EncryptString(data string) (encrypt string, err error) {
	return EncryptBytes([]byte(data))
}

// MustEncryptString 使用MD5算法加密字符串`data`。
// 如果发生任何错误，它将引发panic。
func MustEncryptString(data string) string {
	result, err := EncryptString(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptFile 使用MD5算法加密`path`指定文件的内容。
func EncryptFile(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		err = gerror.Wrapf(err, `os.Open failed for name "%s"`, path)
		return "", err
	}
	defer f.Close()
	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		err = gerror.Wrap(err, `io.Copy failed`)
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// MustEncryptFile 使用MD5算法加密`path`指定文件的内容。
// 如果发生任何错误，将会导致程序panic。
func MustEncryptFile(path string) string {
	result, err := EncryptFile(path)
	if err != nil {
		panic(err)
	}
	return result
}
