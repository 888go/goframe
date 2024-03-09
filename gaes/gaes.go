// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gaes提供了AES加密/解密算法的有用API。
package 加密aes类

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

const (
	// IVDefaultValue 是 IV 的默认值。
	IVDefaultValue = "I Love Go Frame!"
)

// Encrypt 是 EncryptCBC 的别名。
func Encrypt别名(plainText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	return X加密CBC(plainText, key, iv...)
}

// Decrypt 是 DecryptCBC 的别名。
func Decrypt别名(cipherText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	return X解密CBC(cipherText, key, iv...)
}

// EncryptCBC 使用CBC模式加密`plainText`。
// 注意，密钥必须为16/24/32位长度。
// 参数`iv`（初始化向量）是不必要的。
func X加密CBC(待加密 []byte, 秘钥 []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(秘钥)
	if err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `aes.NewCipher failed for key "%s"`, 秘钥)
		return nil, err
	}
	blockSize := block.BlockSize()
	待加密 = PKCS7Padding(待加密, blockSize)
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = []byte(IVDefaultValue)
	}
	blockMode := cipher.NewCBCEncrypter(block, ivValue)
	cipherText := make([]byte, len(待加密))
	blockMode.CryptBlocks(cipherText, 待加密)

	return cipherText, nil
}

// DecryptCBC 使用CBC模式解密`cipherText`。
// 注意，密钥必须为16/24/32比特长度。
// 参数`iv`初始化向量是不必要的。
func X解密CBC(待解密 []byte, 秘钥 []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(秘钥)
	if err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `aes.NewCipher failed for key "%s"`, 秘钥)
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(待解密) < blockSize {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "cipherText too short")
	}
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = []byte(IVDefaultValue)
	}
	if len(待解密)%blockSize != 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "cipherText is not a multiple of the block size")
	}
	blockModel := cipher.NewCBCDecrypter(block, ivValue)
	plainText := make([]byte, len(待解密))
	blockModel.CryptBlocks(plainText, 待解密)
	plainText, e := PKCS7UnPadding(plainText, blockSize)
	if e != nil {
		return nil, e
	}
	return plainText, nil
}

// PKCS5Padding 对源字节切片应用PKCS#5填充，以匹配给定的块大小。
//
// 若未提供块大小，则默认为8。
func PKCS5Padding(src []byte, blockSize ...int) []byte {
	blockSizeTemp := 8
	if len(blockSize) > 0 {
		blockSizeTemp = blockSize[0]
	}
	return PKCS7Padding(src, blockSizeTemp)
}

// PKCS5UnPadding 根据给定的块大小，从源字节切片中移除PKCS#5填充。
//
// 若未提供块大小，则默认为8。
func PKCS5UnPadding(src []byte, blockSize ...int) ([]byte, error) {
	blockSizeTemp := 8
	if len(blockSize) > 0 {
		blockSizeTemp = blockSize[0]
	}
	return PKCS7UnPadding(src, blockSizeTemp)
}

// PKCS7Padding 对源字节切片应用PKCS#7填充，以匹配给定的块大小。
func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// PKCS7UnPadding 根据给定的块大小，从源字节切片中移除PKCS#7填充。
func PKCS7UnPadding(src []byte, blockSize int) ([]byte, error) {
	length := len(src)
	if blockSize <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, fmt.Sprintf("invalid blockSize: %d", blockSize))
	}

	if length%blockSize != 0 || length == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid data len")
	}

	unpadding := int(src[length-1])
	if unpadding > blockSize || unpadding == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid unpadding")
	}

	padding := src[length-unpadding:]
	for i := 0; i < unpadding; i++ {
		if padding[i] != byte(unpadding) {
			return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid padding")
		}
	}

	return src[:(length - unpadding)], nil
}

// EncryptCFB 使用CFB模式加密`plainText`。
// 注意，密钥必须为16/24/32比特长度。
// 参数`iv`（初始化向量）不是必需的。
func X加密CFB(待加密 []byte, 秘钥 []byte, padding *int, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(秘钥)
	if err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `aes.NewCipher failed for key "%s"`, 秘钥)
		return nil, err
	}
	blockSize := block.BlockSize()
	待加密, *padding = ZeroPadding(待加密, blockSize)
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = []byte(IVDefaultValue)
	}
	stream := cipher.NewCFBEncrypter(block, ivValue)
	cipherText := make([]byte, len(待加密))
	stream.XORKeyStream(cipherText, 待加密)
	return cipherText, nil
}

// DecryptCFB 使用CFB模式解密`plainText`。
// 注意，密钥必须为16/24/32位长度。
// 参数`iv`初始化向量是不必要的。
func X解密CFB(待解密 []byte, 秘钥 []byte, unPadding int, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(秘钥)
	if err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `aes.NewCipher failed for key "%s"`, 秘钥)
		return nil, err
	}
	if len(待解密) < aes.BlockSize {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "cipherText too short")
	}
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = []byte(IVDefaultValue)
	}
	stream := cipher.NewCFBDecrypter(block, ivValue)
	plainText := make([]byte, len(待解密))
	stream.XORKeyStream(plainText, 待解密)
	plainText = ZeroUnPadding(plainText, unPadding)
	return plainText, nil
}

func ZeroPadding(cipherText []byte, blockSize int) ([]byte, int) {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(0)}, padding)
	return append(cipherText, padText...), padding
}

func ZeroUnPadding(plaintext []byte, unPadding int) []byte {
	length := len(plaintext)
	return plaintext[:(length - unPadding)]
}
