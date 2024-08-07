// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gdes提供了DES加密/解密算法的有用API。 md5:c8b6785595a2b6ed
package 加密DES类

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
)

const (
	NOPADDING = iota
	PKCS5PADDING
)

// ECB加密使用ECB模式对`plainText`进行加密。 md5:393855a628599a31
func X加密ECB(待加密 []byte, 秘钥 []byte, 填充 int) ([]byte, error) {
	text, err := Padding(待加密, 填充)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, len(text))
	block, err := des.NewCipher(秘钥)
	if err != nil {
		err = gerror.X多层错误码并格式化(gcode.CodeInvalidParameter, err, `des.NewCipher failed for key "%s"`, 秘钥)
		return nil, err
	}

	blockSize := block.BlockSize()
	for i, count := 0, len(text)/blockSize; i < count; i++ {
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Encrypt(cipherText[begin:end], text[begin:end])
	}
	return cipherText, nil
}

// 使用ECB模式解密`cipherText`。 md5:de6e332e6c59244e
func X解密ECB(待解密 []byte, 秘钥 []byte, 填充 int) ([]byte, error) {
	text := make([]byte, len(待解密))
	block, err := des.NewCipher(秘钥)
	if err != nil {
		err = gerror.X多层错误码并格式化(gcode.CodeInvalidParameter, err, `des.NewCipher failed for key "%s"`, 秘钥)
		return nil, err
	}

	blockSize := block.BlockSize()
	for i, count := 0, len(text)/blockSize; i < count; i++ {
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Decrypt(text[begin:end], 待解密[begin:end])
	}

	plainText, err := UnPadding(text, 填充)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

// X加密三重ECB 使用TripleDES和ECB模式加密`plainText`。
// `key`的长度应为16或24字节。
// md5:925f2b0f96f4bc56
func X加密三重ECB(待加密 []byte, 秘钥 []byte, 填充 int) ([]byte, error) {
	if len(秘钥) != 16 && len(秘钥) != 24 {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "key length error")
	}

	text, err := Padding(待加密, 填充)
	if err != nil {
		return nil, err
	}

	var newKey []byte
	if len(秘钥) == 16 {
		newKey = append([]byte{}, 秘钥...)
		newKey = append(newKey, 秘钥[:8]...)
	} else {
		newKey = append([]byte{}, 秘钥...)
	}

	block, err := des.NewTripleDESCipher(newKey)
	if err != nil {
		err = gerror.X多层错误码并格式化(gcode.CodeInvalidParameter, err, `des.NewTripleDESCipher failed for key "%s"`, newKey)
		return nil, err
	}

	blockSize := block.BlockSize()
	cipherText := make([]byte, len(text))
	for i, count := 0, len(text)/blockSize; i < count; i++ {
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Encrypt(cipherText[begin:end], text[begin:end])
	}
	return cipherText, nil
}

// X解密三重ECB 使用三重DES和ECB模式解密`cipherText`。
// `key`的长度应该是16或24字节。
// md5:c1bf1a1ac477aa6e
func X解密三重ECB(待解密 []byte, 秘钥 []byte, 填充 int) ([]byte, error) {
	if len(秘钥) != 16 && len(秘钥) != 24 {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "key length error")
	}

	var newKey []byte
	if len(秘钥) == 16 {
		newKey = append([]byte{}, 秘钥...)
		newKey = append(newKey, 秘钥[:8]...)
	} else {
		newKey = append([]byte{}, 秘钥...)
	}

	block, err := des.NewTripleDESCipher(newKey)
	if err != nil {
		err = gerror.X多层错误码并格式化(gcode.CodeInvalidParameter, err, `des.NewTripleDESCipher failed for key "%s"`, newKey)
		return nil, err
	}

	blockSize := block.BlockSize()
	text := make([]byte, len(待解密))
	for i, count := 0, len(text)/blockSize; i < count; i++ {
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Decrypt(text[begin:end], 待解密[begin:end])
	}

	plainText, err := UnPadding(text, 填充)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

// X加密CBC 使用CBC模式对`plainText`进行加密。 md5:f3a1f35b734799c4
func X加密CBC(待加密 []byte, 密钥 []byte, iv []byte, 填充 int) ([]byte, error) {
	block, err := des.NewCipher(密钥)
	if err != nil {
		err = gerror.X多层错误码并格式化(gcode.CodeInvalidParameter, err, `des.NewCipher failed for key "%s"`, 密钥)
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "invalid iv length")
	}

	text, err := Padding(待加密, 填充)
	if err != nil {
		return nil, err
	}
	cipherText := make([]byte, len(text))

	encryptor := cipher.NewCBCEncrypter(block, iv)
	encryptor.CryptBlocks(cipherText, text)

	return cipherText, nil
}

// 使用CBC模式解密`cipherText`。 md5:232eecbdd9458374
func X解密CBC(待解密 []byte, 密钥 []byte, iv []byte, 填充 int) ([]byte, error) {
	block, err := des.NewCipher(密钥)
	if err != nil {
		err = gerror.X多层错误码并格式化(gcode.CodeInvalidParameter, err, `des.NewCipher failed for key "%s"`, 密钥)
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "iv length invalid")
	}

	text := make([]byte, len(待解密))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(text, 待解密)

	plainText, err := UnPadding(text, 填充)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

// X加密三重CBC 使用TripleDES和CBC模式加密`plainText`。 md5:8fb88673fa27c4a0
func X加密三重CBC(待加密 []byte, 密钥 []byte, iv []byte, 填充 int) ([]byte, error) {
	if len(密钥) != 16 && len(密钥) != 24 {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "key length invalid")
	}

	var newKey []byte
	if len(密钥) == 16 {
		newKey = append([]byte{}, 密钥...)
		newKey = append(newKey, 密钥[:8]...)
	} else {
		newKey = append([]byte{}, 密钥...)
	}

	block, err := des.NewTripleDESCipher(newKey)
	if err != nil {
		err = gerror.X多层错误码并格式化(gcode.CodeInvalidParameter, err, `des.NewTripleDESCipher failed for key "%s"`, newKey)
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "invalid iv length")
	}

	text, err := Padding(待加密, 填充)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, len(text))
	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(cipherText, text)

	return cipherText, nil
}

// X解密三重CBC 使用三重DES和CBC模式解密`cipherText`。 md5:53d3127b0b4fcbb0
func X解密三重CBC(待解密 []byte, 密钥 []byte, iv []byte, 填充 int) ([]byte, error) {
	if len(密钥) != 16 && len(密钥) != 24 {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "key length invalid")
	}

	var newKey []byte
	if len(密钥) == 16 {
		newKey = append([]byte{}, 密钥...)
		newKey = append(newKey, 密钥[:8]...)
	} else {
		newKey = append([]byte{}, 密钥...)
	}

	block, err := des.NewTripleDESCipher(newKey)
	if err != nil {
		err = gerror.X多层错误码并格式化(gcode.CodeInvalidParameter, err, `des.NewTripleDESCipher failed for key "%s"`, newKey)
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "invalid iv length")
	}

	text := make([]byte, len(待解密))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(text, 待解密)

	plainText, err := UnPadding(text, 填充)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

func PaddingPKCS5(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padText...)
}

func UnPaddingPKCS5(text []byte) []byte {
	length := len(text)
	padText := int(text[length-1])
	return text[:(length - padText)]
}

func Padding(text []byte, padding int) ([]byte, error) {
	switch padding {
	case NOPADDING:
		if len(text)%8 != 0 {
			return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "invalid text length")
		}

	case PKCS5PADDING:
		return PaddingPKCS5(text, 8), nil

	default:
		return nil, gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `unsupported padding type "%d"`, padding)
	}

	return text, nil
}

func UnPadding(text []byte, padding int) ([]byte, error) {
	switch padding {
	case NOPADDING:
		if len(text)%8 != 0 {
			return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "invalid text length")
		}

	case PKCS5PADDING:
		return UnPaddingPKCS5(text), nil

	default:
		return nil, gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `unsupported padding type "%d"`, padding)
	}
	return text, nil
}
