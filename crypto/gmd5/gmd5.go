// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gmd5提供了MD5加密算法的实用API。
package 加密md5类

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/util/gconv"
)

// Encrypt 使用MD5算法加密任意类型的变量。
// 它使用gconv包将`v`转换为字节类型。
func X加密(值 interface{}) (md5值 string, 错误 error) {
	return X加密字节集(转换类.X取字节集(值))
}

// MustEncrypt 使用MD5算法对任意类型的变量进行加密。
// 它使用gconv包将`v`转换为字节类型。
// 如果发生任何错误，它会引发panic。
func X加密PANI(值 interface{}) string {
	result, err := X加密(值)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptBytes 使用MD5算法加密`data`。
func X加密字节集(字节集 []byte) (md5值 string, 错误 error) {
	h := md5.New()
	if _, 错误 = h.Write(字节集); 错误 != nil {
		错误 = 错误类.X多层错误(错误, `hash.Write failed`)
		return "", 错误
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// MustEncryptBytes 使用MD5算法加密`data`。
// 如果出现任何错误，将会导致panic。
func X加密字节集PANI(字节集 []byte) string {
	result, err := X加密字节集(字节集)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptString 使用MD5算法加密字符串`data`。
func X加密文本(值 string) (md5值 string, 错误 error) {
	return X加密字节集([]byte(值))
}

// MustEncryptString 使用MD5算法加密字符串`data`。
// 如果发生任何错误，它将引发panic。
func X加密文本PANI(值 string) string {
	result, err := X加密文本(值)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptFile 使用MD5算法加密`path`指定文件的内容。
func X加密文件(路径 string) (md5值 string, 错误 error) {
	f, 错误 := os.Open(路径)
	if 错误 != nil {
		错误 = 错误类.X多层错误并格式化(错误, `os.Open failed for name "%s"`, 路径)
		return "", 错误
	}
	defer f.Close()
	h := md5.New()
	_, 错误 = io.Copy(h, f)
	if 错误 != nil {
		错误 = 错误类.X多层错误(错误, `io.Copy failed`)
		return "", 错误
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// MustEncryptFile 使用MD5算法加密`path`指定文件的内容。
// 如果发生任何错误，将会导致程序panic。
func X加密文件PANI(路径 string) string {
	result, err := X加密文件(路径)
	if err != nil {
		panic(err)
	}
	return result
}
