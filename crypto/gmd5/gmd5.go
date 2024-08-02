// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// gmd5包提供了对MD5加密算法的实用API。 md5:637f00f8697c325b
package 加密md5类

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"

	gerror "github.com/888go/goframe/errors/gerror"
	gconv "github.com/888go/goframe/util/gconv"
)

// Encrypt 使用MD5算法加密任何类型的变量。
// 它使用gconv包将`v`转换为字节类型。
// md5:0ab9d73eb0da5581
func Encrypt(data interface{}) (encrypt string, err error) {
	return EncryptBytes(gconv.Bytes(data))
}

// MustEncrypt 使用MD5算法对任何类型的变量进行加密。
// 它使用gconv包将`v`转换为其字节类型。
// 如果发生任何错误，它将引发恐慌。
// md5:759531471ae8fb5a
func MustEncrypt(data interface{}) string {
	result, err := Encrypt(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptBytes 使用MD5算法对`data`进行加密。 md5:28a4ffde44149352
func EncryptBytes(data []byte) (encrypt string, err error) {
	h := md5.New()
	if _, err = h.Write(data); err != nil {
		err = gerror.Wrap(err, `hash.Write failed`)
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// MustEncryptBytes 使用MD5算法对`data`进行加密。
// 如果发生任何错误，它将直接 panic。
// md5:24200c4b1dd0cbd3
func MustEncryptBytes(data []byte) string {
	result, err := EncryptBytes(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptString 使用MD5算法对字符串`data`进行加密。 md5:c2472c71dca7f578
func EncryptString(data string) (encrypt string, err error) {
	return EncryptBytes([]byte(data))
}

// MustEncryptString 使用MD5算法对字符串`data`进行加密。如果发生任何错误，它将引发恐慌。
// md5:54e2ed7e76b2c713
func MustEncryptString(data string) string {
	result, err := EncryptString(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptFile 使用MD5算法对`path`指定文件的内容进行加密。 md5:169266aa4496a2b4
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

// MustEncryptFile 使用MD5算法加密`path`文件的内容。
// 如果发生任何错误，它将直接 panic。
// md5:71f4a9ffa26ff10c
func MustEncryptFile(path string) string {
	result, err := EncryptFile(path)
	if err != nil {
		panic(err)
	}
	return result
}
