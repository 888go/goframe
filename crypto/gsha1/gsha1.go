// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gsha1提供了SHA1加密算法的有用API。 md5:4ebe688b6095e4db
package 加密sha1类

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	gerror "github.com/888go/goframe/errors/gerror"
	gconv "github.com/888go/goframe/util/gconv"
)

// 使用SHA1算法对任何类型的变量进行加密。
// 它使用gconv包将`v`转换为其字节类型。
// md5:3bcfc7ac2d70d9e3
func Encrypt(v interface{}) string {
	r := sha1.Sum(gconv.Bytes(v))
	return hex.EncodeToString(r[:])
}

// EncryptFile 使用SHA1算法对`path`路径下的文件内容进行加密。 md5:25246a5477d29491
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

// MustEncryptFile 使用SHA1算法对`path`指定文件的内容进行加密。如果发生任何错误，它将引发恐慌。
// md5:ee1a2c634d668ad2
func MustEncryptFile(path string) string {
	result, err := EncryptFile(path)
	if err != nil {
		panic(err)
	}
	return result
}
