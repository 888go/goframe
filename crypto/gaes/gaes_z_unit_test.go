// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package gaes_test

import (
	"testing"
	
	"github.com/888go/goframe/crypto/gaes"
	"github.com/888go/goframe/encoding/gbase64"
	"github.com/888go/goframe/test/gtest"
)

var (
	content          = []byte("pibigstar")
	content_16, _    = gbase64.DecodeString("v1jqsGHId/H8onlVHR8Vaw==")
	content_24, _    = gbase64.DecodeString("0TXOaj5KMoLhNWmJ3lxY1A==")
	content_32, _    = gbase64.DecodeString("qM/Waw1kkWhrwzek24rCSA==")
	content_16_iv, _ = gbase64.DecodeString("DqQUXiHgW/XFb6Qs98+hrA==")
	content_32_iv, _ = gbase64.DecodeString("ZuLgAOii+lrD5KJoQ7yQ8Q==")
	// iv 长度必须等于blockSize，只能为16
	iv         = []byte("Hello My GoFrame")
	key_16     = []byte("1234567891234567")
	key_17     = []byte("12345678912345670")
	key_24     = []byte("123456789123456789123456")
	key_32     = []byte("12345678912345678912345678912345")
	keys       = []byte("12345678912345678912345678912346")
	key_err    = []byte("1234")
	key_32_err = []byte("1234567891234567891234567891234 ")

	// cfb模式blockSize补位长度, add by zseeker
	padding_size      = 16 - len(content)
	content_16_cfb, _ = gbase64.DecodeString("oSmget3aBDT1nJnBp8u6kA==")
)

func TestEncrypt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		data, err := gaes.Encrypt(content, key_16)
		t.AssertNil(err)
		t.Assert(data, []byte(content_16))
		data, err = gaes.Encrypt(content, key_24)
		t.AssertNil(err)
		t.Assert(data, []byte(content_24))
		data, err = gaes.Encrypt(content, key_32)
		t.AssertNil(err)
		t.Assert(data, []byte(content_32))
		data, err = gaes.Encrypt(content, key_16, iv)
		t.AssertNil(err)
		t.Assert(data, []byte(content_16_iv))
		data, err = gaes.Encrypt(content, key_32, iv)
		t.AssertNil(err)
		t.Assert(data, []byte(content_32_iv))
	})
}

func TestDecrypt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		decrypt, err := gaes.Decrypt([]byte(content_16), key_16)
		t.AssertNil(err)
		t.Assert(decrypt, content)

		decrypt, err = gaes.Decrypt([]byte(content_24), key_24)
		t.AssertNil(err)
		t.Assert(decrypt, content)

		decrypt, err = gaes.Decrypt([]byte(content_32), key_32)
		t.AssertNil(err)
		t.Assert(decrypt, content)

		decrypt, err = gaes.Decrypt([]byte(content_16_iv), key_16, iv)
		t.AssertNil(err)
		t.Assert(decrypt, content)

		decrypt, err = gaes.Decrypt([]byte(content_32_iv), key_32, iv)
		t.AssertNil(err)
		t.Assert(decrypt, content)

		decrypt, err = gaes.Decrypt([]byte(content_32_iv), keys, iv)
		t.Assert(err, "invalid unpadding")
	})
}

func TestEncryptErr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// encrypt key error
		_, err := gaes.Encrypt(content, key_err)
		t.AssertNE(err, nil)
	})
}

func TestDecryptErr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// decrypt key error
		encrypt, err := gaes.Encrypt(content, key_16)
		_, err = gaes.Decrypt(encrypt, key_err)
		t.AssertNE(err, nil)

		// 解密内容过短错误
		_, err = gaes.Decrypt([]byte("test"), key_16)
		t.AssertNE(err, nil)

		// 解密内容大小错误
		_, err = gaes.Decrypt(key_17, key_16)
		t.AssertNE(err, nil)
	})
}

func TestPKCS5UnPaddingErr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// PKCS5UnPadding：对于给定的 blockSize，执行PKCS#5填充解码（即去除填充部分）
// 此函数通常用于对使用PKCS#5填充方式加密的数据进行解密前的填充剥离
		_, err := gaes.PKCS5UnPadding(content, 0)
		t.AssertNE(err, nil)

		// PKCS5UnPadding 解密时去除PKCS5填充, src为待解密的字节切片, len为src的长度
// 当len为0时，表示处理特殊情况，可能由于数据异常或者已无填充
// 在密码学中，PKCS5Padding是对明文进行填充以满足块加密算法对输入长度要求的一种方法。在解密时需要将这些填充的部分去除。上述注释表明这是一个用于在解密后执行PKCS5填充去除操作的函数，并特别指出了当输入长度为0时的情况，这可能意味着原始数据不包含有效的填充或已经不是完整的块。
		_, err = gaes.PKCS5UnPadding([]byte(""), 16)
		t.AssertNE(err, nil)

		// PKCS5UnPadding 当原始数据长度大于块大小时
// 这段注释表明这是一个Go语言函数，用于执行PKCS#5填充的解填充操作，前提是源数据的长度大于块大小。在密码学中，PKCS#5填充是一种用于确保明文数据长度适配分组密码（如AES）所需块大小的方法，在解密后需要进行解填充以恢复原始数据。
		_, err = gaes.PKCS5UnPadding(key_17, 16)
		t.AssertNE(err, nil)

		// PKCS5UnPadding 当原始数据长度大于块大小时
// 这段注释表明这是一个Go语言函数，用于执行PKCS#5填充的解填充操作，前提是源数据的长度大于块大小。在密码学中，PKCS#5填充是一种用于确保明文数据长度适配分组密码（如AES）所需块大小的方法，在解密后需要进行解填充以恢复原始数据。
		_, err = gaes.PKCS5UnPadding(key_32_err, 32)
		t.AssertNE(err, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		// PKCS7UnPadding 函数，用于移除PKCS7填充
// blockSize 指定块大小（以字节为单位），在解密时需要提供该参数以便正确地移除填充
// ```go
// PKCS7UnPadding func
// 参数:
//   - paddedData: 已经过PKCS7填充的原始数据
//   - blockSize: 块大小，以字节为单位
// func PKCS7UnPadding(paddedData []byte, blockSize int) ([]byte, error) {
    // ...
// }
		_, err := gaes.PKCS7UnPadding(content, 0)
		t.AssertNE(err, nil)

		// PKCS7UnPadding 对PKCS7填充方式进行解码，src为待解码的源数据，len表示源数据长度，需要注意的是这里的源数据长度可能为零。
		_, err = gaes.PKCS7UnPadding([]byte(""), 16)
		t.AssertNE(err, nil)

		// PKCS7UnPadding 函数，用于处理src（源数据）的PKCS#7解填充，要求src的长度大于块大小（blockSize）
// 在密码学中，PKCS7Padding是一种常用的对称加密数据填充方式，确保加密后数据长度是块大小的整数倍。当解密时，需要通过PKCS7UnPadding函数去除末尾的填充数据。这段注释说明了该函数的作用和调用前的约束条件——待解填充的数据长度必须大于块大小。
		_, err = gaes.PKCS7UnPadding(key_17, 16)
		t.AssertNE(err, nil)

		// PKCS7UnPadding 函数，用于处理src（源数据）的PKCS#7解填充，要求src的长度大于块大小（blockSize）
// 在密码学中，PKCS7Padding是一种常用的对称加密数据填充方式，确保加密后数据长度是块大小的整数倍。当解密时，需要通过PKCS7UnPadding函数去除末尾的填充数据。这段注释说明了该函数的作用和调用前的约束条件——待解填充的数据长度必须大于块大小。
		_, err = gaes.PKCS7UnPadding(key_32_err, 32)
		t.AssertNE(err, nil)
	})
}

func TestEncryptCFB(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var padding int = 0
		data, err := gaes.EncryptCFB(content, key_16, &padding, iv)
		t.AssertNil(err)
		t.Assert(padding, padding_size)
		t.Assert(data, []byte(content_16_cfb))
	})
}

func TestDecryptCFB(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		decrypt, err := gaes.DecryptCFB([]byte(content_16_cfb), key_16, padding_size, iv)
		t.AssertNil(err)
		t.Assert(decrypt, content)
	})
}
