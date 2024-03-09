// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 加密md5类_test

import (
	"os"
	"testing"
	
	"github.com/888go/goframe/gmd5"
	"github.com/gogf/gf/v2/test/gtest"
)

var (
	s = "pibigstar"
	// 在线生成的MD5值
	result = "d175a1ff66aedde64344785f7f7a3df8"
)

type user struct {
	name     string
	password string
	age      int
}

func TestEncrypt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		encryptString, _ := 加密md5类.X加密(s)
		t.Assert(encryptString, result)

		result := "1427562bb29f88a1161590b76398ab72"
		encrypt, _ := 加密md5类.X加密(123456)
		t.AssertEQ(encrypt, result)
	})

	gtest.C(t, func(t *gtest.T) {
		user := &user{
			name:     "派大星",
			password: "123456",
			age:      23,
		}
		result := "70917ebce8bd2f78c736cda63870fb39"
		encrypt, _ := 加密md5类.X加密(user)
		t.AssertEQ(encrypt, result)
	})
}

func TestEncryptString(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		encryptString, _ := 加密md5类.X加密文本(s)
		t.Assert(encryptString, result)
	})
}

func TestEncryptFile(t *testing.T) {
	path := "test.text"
	errorPath := "err.txt"
	result := "e6e6e1cd41895beebff16d5452dfce12"
	gtest.C(t, func(t *gtest.T) {
		file, err := os.Create(path)
		defer os.Remove(path)
		defer file.Close()
		t.AssertNil(err)
		_, _ = file.Write([]byte("Hello Go Frame"))
		encryptFile, _ := 加密md5类.X加密文件(path)
		t.AssertEQ(encryptFile, result)
		// 当文件不存在时，encrypt将返回空字符串
		errEncrypt, _ := 加密md5类.X加密文件(errorPath)
		t.AssertEQ(errEncrypt, "")
	})

}
