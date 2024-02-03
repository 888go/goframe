// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdes_test

import (
	"encoding/hex"
	"testing"
	
	"github.com/888go/goframe/crypto/gdes"
	"github.com/888go/goframe/test/gtest"
)

var (
	errKey     = []byte("1111111111111234123456789")
	errIv      = []byte("123456789")
	errPadding = 5
)

func TestDesECB(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		key := []byte("11111111")
		text := []byte("12345678")
		padding := gdes.NOPADDING
		result := "858b176da8b12503"
		// encrypt test
		cipherText, err := gdes.EncryptECB(text, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptECB(cipherText, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "12345678")

		// 加密错误测试。当抛出异常时，err不等于nil，并且字符串为nil
		errEncrypt, err := gdes.EncryptECB(text, key, errPadding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		errEncrypt, err = gdes.EncryptECB(text, errKey, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		// err decrypt test.
		errDecrypt, err := gdes.DecryptECB(cipherText, errKey, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errDecrypt, nil)
		errDecrypt, err = gdes.DecryptECB(cipherText, key, errPadding)
		t.AssertNE(err, nil)
		t.AssertEQ(errDecrypt, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		key := []byte("11111111")
		text := []byte("12345678")
		padding := gdes.PKCS5PADDING
		errPadding := 5
		result := "858b176da8b12503ad6a88b4fa37833d"
		cipherText, err := gdes.EncryptECB(text, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptECB(cipherText, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "12345678")

		// err test
		errEncrypt, err := gdes.EncryptECB(text, key, errPadding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		errDecrypt, err := gdes.DecryptECB(cipherText, errKey, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errDecrypt, nil)
	})
}

func Test3DesECB(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		key := []byte("1111111111111234")
		text := []byte("1234567812345678")
		padding := gdes.NOPADDING
		result := "a23ee24b98c26263a23ee24b98c26263"
		// encrypt test
		cipherText, err := gdes.EncryptECBTriple(text, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptECBTriple(cipherText, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "1234567812345678")
		// err test
		errEncrypt, err := gdes.EncryptECB(text, key, errPadding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		key := []byte("111111111111123412345678")
		text := []byte("123456789")
		padding := gdes.PKCS5PADDING
		errPadding := 5
		result := "37989b1effc07a6d00ff89a7d052e79f"
		// encrypt test
		cipherText, err := gdes.EncryptECBTriple(text, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptECBTriple(cipherText, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "123456789")
		// err 测试，当 key 为错误值时，但 text 和 padding 是正确的
		errEncrypt, err := gdes.EncryptECBTriple(text, errKey, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		// 当填充错误时，但密钥和文本是正确的
// 此句英文注释翻译成中文注释为：
// ```go
// 当填充格式错误时，但密钥和原文内容是正确的
		errEncrypt, err = gdes.EncryptECBTriple(text, key, errPadding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		// 解密错误测试，当密钥错误时
		errEncrypt, err = gdes.DecryptECBTriple(text, errKey, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
	})
}

func TestDesCBC(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		key := []byte("11111111")
		text := []byte("1234567812345678")
		padding := gdes.NOPADDING
		iv := []byte("12345678")
		result := "40826a5800608c87585ca7c9efabee47"
		// encrypt test
		cipherText, err := gdes.EncryptCBC(text, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptCBC(cipherText, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "1234567812345678")
		// encrypt err test.
		errEncrypt, err := gdes.EncryptCBC(text, errKey, iv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		// the iv is err
		errEncrypt, err = gdes.EncryptCBC(text, key, errIv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		// 这里的填充有误
		errEncrypt, err = gdes.EncryptCBC(text, key, iv, errPadding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		// 解密错误测试，关键字为 err
		errDecrypt, err := gdes.DecryptCBC(cipherText, errKey, iv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errDecrypt, nil)
		// the iv is err
		errDecrypt, err = gdes.DecryptCBC(cipherText, key, errIv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errDecrypt, nil)
		// 这里的填充有误
		errDecrypt, err = gdes.DecryptCBC(cipherText, key, iv, errPadding)
		t.AssertNE(err, nil)
		t.AssertEQ(errDecrypt, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		key := []byte("11111111")
		text := []byte("12345678")
		padding := gdes.PKCS5PADDING
		iv := []byte("12345678")
		result := "40826a5800608c87100a25d86ac7c52c"
		// encrypt test
		cipherText, err := gdes.EncryptCBC(text, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptCBC(cipherText, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "12345678")
		// err test
		errEncrypt, err := gdes.EncryptCBC(text, key, errIv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
	})
}

func Test3DesCBC(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		key := []byte("1111111112345678")
		text := []byte("1234567812345678")
		padding := gdes.NOPADDING
		iv := []byte("12345678")
		result := "bfde1394e265d5f738d5cab170c77c88"
		// encrypt test
		cipherText, err := gdes.EncryptCBCTriple(text, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptCBCTriple(cipherText, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "1234567812345678")
		// encrypt err test
		errEncrypt, err := gdes.EncryptCBCTriple(text, errKey, iv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		// the iv is err
		errEncrypt, err = gdes.EncryptCBCTriple(text, key, errIv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		// 这里的填充有误
		errEncrypt, err = gdes.EncryptCBCTriple(text, key, iv, errPadding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		// decrypt err test
		errDecrypt, err := gdes.DecryptCBCTriple(cipherText, errKey, iv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errDecrypt, nil)
		// the iv is err
		errDecrypt, err = gdes.DecryptCBCTriple(cipherText, key, errIv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errDecrypt, nil)
		// 这里的填充有误
		errDecrypt, err = gdes.DecryptCBCTriple(cipherText, key, iv, errPadding)
		t.AssertNE(err, nil)
		t.AssertEQ(errDecrypt, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		key := []byte("111111111234567812345678")
		text := []byte("12345678")
		padding := gdes.PKCS5PADDING
		iv := []byte("12345678")
		result := "40826a5800608c87100a25d86ac7c52c"
		// encrypt test
		cipherText, err := gdes.EncryptCBCTriple(text, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptCBCTriple(cipherText, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "12345678")
	})

}
