// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 加密DES类_test

import (
	"encoding/hex"
	"testing"

	gdes "github.com/888go/goframe/crypto/gdes"
	gtest "github.com/888go/goframe/test/gtest"
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
		cipherText, err := gdes.X加密ECB(text, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.X解密ECB(cipherText, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "12345678")

						// 加密错误测试。当抛出异常时，err不等于nil，而字符串为nil. md5:e1fa7ef89e2e43d2
		errEncrypt, err := gdes.X加密ECB(text, key, errPadding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		errEncrypt, err = gdes.X加密ECB(text, errKey, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		// err decrypt test.
		errDecrypt, err := gdes.X解密ECB(cipherText, errKey, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errDecrypt, nil)
		errDecrypt, err = gdes.X解密ECB(cipherText, key, errPadding)
		t.AssertNE(err, nil)
		t.AssertEQ(errDecrypt, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		key := []byte("11111111")
		text := []byte("12345678")
		padding := gdes.PKCS5PADDING
		errPadding := 5
		result := "858b176da8b12503ad6a88b4fa37833d"
		cipherText, err := gdes.X加密ECB(text, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.X解密ECB(cipherText, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "12345678")

		// err test
		errEncrypt, err := gdes.X加密ECB(text, key, errPadding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		errDecrypt, err := gdes.X解密ECB(cipherText, errKey, padding)
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
		cipherText, err := gdes.X加密三重ECB(text, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.X解密三重ECB(cipherText, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "1234567812345678")
		// err test
		errEncrypt, err := gdes.X加密ECB(text, key, errPadding)
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
		cipherText, err := gdes.X加密三重ECB(text, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.X解密三重ECB(cipherText, key, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "123456789")
						// 错误测试，当键为err时，但文本和填充是正确的. md5:aa9db0c26aff396e
		errEncrypt, err := gdes.X加密三重ECB(text, errKey, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
				// 当填充出错，但密钥和文本是正确的时. md5:810554d84cfce555
		errEncrypt, err = gdes.X加密三重ECB(text, key, errPadding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
						// 解密错误测试，当密钥为错误时. md5:cd0ff4262b029c3f
		errEncrypt, err = gdes.X解密三重ECB(text, errKey, padding)
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
		cipherText, err := gdes.X加密CBC(text, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.X解密CBC(cipherText, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "1234567812345678")
		// encrypt err test.
		errEncrypt, err := gdes.X加密CBC(text, errKey, iv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		// the iv is err
		errEncrypt, err = gdes.X加密CBC(text, key, errIv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		// the padding is err
		errEncrypt, err = gdes.X加密CBC(text, key, iv, errPadding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
						// 解密错误测试。密钥是 err. md5:d37b14dae71a9447
		errDecrypt, err := gdes.X解密CBC(cipherText, errKey, iv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errDecrypt, nil)
		// the iv is err
		errDecrypt, err = gdes.X解密CBC(cipherText, key, errIv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errDecrypt, nil)
		// the padding is err
		errDecrypt, err = gdes.X解密CBC(cipherText, key, iv, errPadding)
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
		cipherText, err := gdes.X加密CBC(text, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.X解密CBC(cipherText, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "12345678")
		// err test
		errEncrypt, err := gdes.X加密CBC(text, key, errIv, padding)
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
		cipherText, err := gdes.X加密三重CBC(text, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.X解密三重CBC(cipherText, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "1234567812345678")
		// encrypt err test
		errEncrypt, err := gdes.X加密三重CBC(text, errKey, iv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		// the iv is err
		errEncrypt, err = gdes.X加密三重CBC(text, key, errIv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		// the padding is err
		errEncrypt, err = gdes.X加密三重CBC(text, key, iv, errPadding)
		t.AssertNE(err, nil)
		t.AssertEQ(errEncrypt, nil)
		// decrypt err test
		errDecrypt, err := gdes.X解密三重CBC(cipherText, errKey, iv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errDecrypt, nil)
		// the iv is err
		errDecrypt, err = gdes.X解密三重CBC(cipherText, key, errIv, padding)
		t.AssertNE(err, nil)
		t.AssertEQ(errDecrypt, nil)
		// the padding is err
		errDecrypt, err = gdes.X解密三重CBC(cipherText, key, iv, errPadding)
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
		cipherText, err := gdes.X加密三重CBC(text, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.X解密三重CBC(cipherText, key, iv, padding)
		t.AssertEQ(err, nil)
		t.AssertEQ(string(clearText), "12345678")
	})

}
