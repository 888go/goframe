// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// gmd5包提供了对MD5加密算法的实用API。 md5:637f00f8697c325b
package gmd5//bm:加密md5类

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

// Encrypt 使用MD5算法加密任何类型的变量。
// 它使用gconv包将`v`转换为字节类型。
// md5:0ab9d73eb0da5581
// ff:加密
// data:值
// encrypt:md5值
// err:错误
func Encrypt(data interface{}) (encrypt string, err error) {
	return EncryptBytes(gconv.Bytes(data))
}

// MustEncrypt 使用MD5算法对任何类型的变量进行加密。
// 它使用gconv包将`v`转换为其字节类型。
// 如果发生任何错误，它将引发恐慌。
// md5:759531471ae8fb5a
// ff:加密PANI
// data:值
func MustEncrypt(data interface{}) string {
	result, err := Encrypt(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptBytes 使用MD5算法对`data`进行加密。 md5:28a4ffde44149352
// ff:加密字节集
// data:字节集
// encrypt:md5值
// err:错误
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
// ff:加密字节集PANI
// data:字节集
func MustEncryptBytes(data []byte) string {
	result, err := EncryptBytes(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptString 使用MD5算法对字符串`data`进行加密。 md5:c2472c71dca7f578
// ff:加密文本
// data:值
// encrypt:md5值
// err:错误
func EncryptString(data string) (encrypt string, err error) {
	return EncryptBytes([]byte(data))
}

// MustEncryptString 使用MD5算法对字符串`data`进行加密。如果发生任何错误，它将引发恐慌。
// md5:54e2ed7e76b2c713
// ff:加密文本PANI
// data:值
func MustEncryptString(data string) string {
	result, err := EncryptString(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptFile 使用MD5算法对`path`指定文件的内容进行加密。 md5:169266aa4496a2b4
// ff:加密文件
// path:路径
// encrypt:md5值
// err:错误
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
// ff:加密文件PANI
// path:路径
func MustEncryptFile(path string) string {
	result, err := EncryptFile(path)
	if err != nil {
		panic(err)
	}
	return result
}
