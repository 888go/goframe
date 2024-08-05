// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package gmd5_test

import (
	"os"
	"testing"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/test/gtest"
)

var (
	s = "pibigstar"
		// 在线生成的MD5值. md5:2b9ec6bcf95f7091
	result = "d175a1ff66aedde64344785f7f7a3df8"
)

type user struct {
	name     string
	password string
	age      int
}

func TestEncrypt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		encryptString, _ := gmd5.Encrypt(s)
		t.Assert(encryptString, result)

		result := "1427562bb29f88a1161590b76398ab72"
		encrypt, _ := gmd5.Encrypt(123456)
		t.AssertEQ(encrypt, result)
	})

	gtest.C(t, func(t *gtest.T) {
		user := &user{
			name:     "派大星",
			password: "123456",
			age:      23,
		}
		result := "70917ebce8bd2f78c736cda63870fb39"
		encrypt, _ := gmd5.Encrypt(user)
		t.AssertEQ(encrypt, result)
	})
}

func TestEncryptString(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		encryptString, _ := gmd5.EncryptString(s)
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
		encryptFile, _ := gmd5.EncryptFile(path)
		t.AssertEQ(encryptFile, result)
				// 当文件不存在时，encrypt会返回空字符串. md5:2282711167b98cb7
		errEncrypt, _ := gmd5.EncryptFile(errorPath)
		t.AssertEQ(errEncrypt, "")
	})

}
