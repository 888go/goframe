// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// gaes 包提供了对AES加密/解密算法有用的API。 md5:a8dd4c4d404f7193
package 加密aes类

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
)

const (
		// IVDefaultValue 是初始向量(IV)的默认值。 md5:4a6e2194de451335
	IVDefaultValue = "I Love Go Frame!"
)

// Encrypt别名 是 EncryptCBC 的别名。 md5:d1191baf4cd313b4
func Encrypt别名(plainText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	return X加密CBC(plainText, key, iv...)
}

// Decrypt别名 是DecryptCBC的别名。 md5:8cf4ecbfea3f2867
func Decrypt别名(cipherText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	return X解密CBC(cipherText, key, iv...)
}

// 使用CBC模式加密`plainText`。
// 注意密钥长度必须为16/24/32位。
// 参数`iv`初始化向量是不必要的。
// md5:1628ebc8e55608ea
func X加密CBC(待加密 []byte, 秘钥 []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(秘钥)
	if err != nil {
		err = gerror.X多层错误码并格式化(gcode.CodeInvalidParameter, err, `aes.NewCipher failed for key "%s"`, 秘钥)
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

// X解密CBC 使用CBC模式解密`cipherText`。
// 请注意，密钥必须为16/24/32位长度。
// 参数`iv`初始化向量是可选的。
// md5:ffdc2bd43249f656
func X解密CBC(待解密 []byte, 秘钥 []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(秘钥)
	if err != nil {
		err = gerror.X多层错误码并格式化(gcode.CodeInvalidParameter, err, `aes.NewCipher failed for key "%s"`, 秘钥)
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(待解密) < blockSize {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "cipherText too short")
	}
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = []byte(IVDefaultValue)
	}
	if len(待解密)%blockSize != 0 {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "cipherText is not a multiple of the block size")
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

// PKCS5Padding 对源字节切片应用 PKCS#5 填充，以匹配给定的块大小。
//
// 如果未提供块大小，则默认为 8。
// md5:709e406aa572f106
func PKCS5Padding(src []byte, blockSize ...int) []byte {
	blockSizeTemp := 8
	if len(blockSize) > 0 {
		blockSizeTemp = blockSize[0]
	}
	return PKCS7Padding(src, blockSizeTemp)
}

// PKCS5UnPadding 根据给定的块大小，从源字节切片中移除PKCS#5填充。
// 
// 如果未提供块大小，则默认为8。
// md5:d61fd48ac346da68
func PKCS5UnPadding(src []byte, blockSize ...int) ([]byte, error) {
	blockSizeTemp := 8
	if len(blockSize) > 0 {
		blockSizeTemp = blockSize[0]
	}
	return PKCS7UnPadding(src, blockSizeTemp)
}

// PKCS7Padding 对源字节切片应用PKCS#7填充，以匹配给定的块大小。 md5:c93d69357ddcf364
func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// PKCS7UnPadding 根据给定的块大小从源字节切片中移除PKCS#7填充。 md5:405becc32a6915c2
func PKCS7UnPadding(src []byte, blockSize int) ([]byte, error) {
	length := len(src)
	if blockSize <= 0 {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, fmt.Sprintf("invalid blockSize: %d", blockSize))
	}

	if length%blockSize != 0 || length == 0 {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "invalid data len")
	}

	unpadding := int(src[length-1])
	if unpadding > blockSize || unpadding == 0 {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "invalid unpadding")
	}

	padding := src[length-unpadding:]
	for i := 0; i < unpadding; i++ {
		if padding[i] != byte(unpadding) {
			return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "invalid padding")
		}
	}

	return src[:(length - unpadding)], nil
}

// X加密CFB 使用CFB模式对`plainText`进行加密。
// 注意，密钥必须是16/24/32位长度。
// 参数`iv`（初始化向量）是不必要的。
// md5:cdcc74633b342790
func X加密CFB(待加密 []byte, 秘钥 []byte, padding *int, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(秘钥)
	if err != nil {
		err = gerror.X多层错误码并格式化(gcode.CodeInvalidParameter, err, `aes.NewCipher failed for key "%s"`, 秘钥)
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

// X解密CFB 使用CFB模式解密`plainText`。
// 注意，密钥必须是16/24/32位长度。
// 参数`iv`（初始化向量）是不必要的。
// md5:f6a0b1655dd052b7
func X解密CFB(待解密 []byte, 秘钥 []byte, unPadding int, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(秘钥)
	if err != nil {
		err = gerror.X多层错误码并格式化(gcode.CodeInvalidParameter, err, `aes.NewCipher failed for key "%s"`, 秘钥)
		return nil, err
	}
	if len(待解密) < aes.BlockSize {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "cipherText too short")
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
