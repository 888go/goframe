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

// X加密 使用MD5算法加密任何类型的变量。
// 它使用gconv包将`v`转换为字节类型。
// md5:0ab9d73eb0da5581
func X加密(值 interface{}) (md5值 string, 错误 error) {
	return X加密字节集(gconv.X取字节集(值))
}

// X加密PANI 使用MD5算法对任何类型的变量进行加密。
// 它使用gconv包将`v`转换为其字节类型。
// 如果发生任何错误，它将引发恐慌。
// md5:759531471ae8fb5a
func X加密PANI(值 interface{}) string {
	result, err := X加密(值)
	if err != nil {
		panic(err)
	}
	return result
}

// X加密字节集 使用MD5算法对`data`进行加密。 md5:28a4ffde44149352
func X加密字节集(字节集 []byte) (md5值 string, 错误 error) {
	h := md5.New()
	if _, 错误 = h.Write(字节集); 错误 != nil {
		错误 = gerror.X多层错误(错误, `hash.Write failed`)
		return "", 错误
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// X加密字节集PANI 使用MD5算法对`data`进行加密。
// 如果发生任何错误，它将直接 panic。
// md5:24200c4b1dd0cbd3
func X加密字节集PANI(字节集 []byte) string {
	result, err := X加密字节集(字节集)
	if err != nil {
		panic(err)
	}
	return result
}

// X加密文本 使用MD5算法对字符串`data`进行加密。 md5:c2472c71dca7f578
func X加密文本(值 string) (md5值 string, 错误 error) {
	return X加密字节集([]byte(值))
}

// X加密文本PANI 使用MD5算法对字符串`data`进行加密。如果发生任何错误，它将引发恐慌。
// md5:54e2ed7e76b2c713
func X加密文本PANI(值 string) string {
	result, err := X加密文本(值)
	if err != nil {
		panic(err)
	}
	return result
}

// X加密文件 使用MD5算法对`path`指定文件的内容进行加密。 md5:169266aa4496a2b4
func X加密文件(路径 string) (md5值 string, 错误 error) {
	f, 错误 := os.Open(路径)
	if 错误 != nil {
		错误 = gerror.X多层错误并格式化(错误, `os.Open failed for name "%s"`, 路径)
		return "", 错误
	}
	defer f.Close()
	h := md5.New()
	_, 错误 = io.Copy(h, f)
	if 错误 != nil {
		错误 = gerror.X多层错误(错误, `io.Copy failed`)
		return "", 错误
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// X加密文件PANI 使用MD5算法加密`path`文件的内容。
// 如果发生任何错误，它将直接 panic。
// md5:71f4a9ffa26ff10c
func X加密文件PANI(路径 string) string {
	result, err := X加密文件(路径)
	if err != nil {
		panic(err)
	}
	return result
}
