// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gdes提供了DES加密/解密算法的有用API。 md5:c8b6785595a2b6ed
package gdes//bm:加密DES类

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

const (
	NOPADDING = iota
	PKCS5PADDING
)

// ECB加密使用ECB模式对`plainText`进行加密。 md5:393855a628599a31
// ff:加密ECB
// plainText:待加密
// key:秘钥
// padding:填充
func EncryptECB(plainText []byte, key []byte, padding int) ([]byte, error) {
	text, err := Padding(plainText, padding)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, len(text))
	block, err := des.NewCipher(key)
	if err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `des.NewCipher failed for key "%s"`, key)
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
// ff:解密ECB
// cipherText:待解密
// key:秘钥
// padding:填充
func DecryptECB(cipherText []byte, key []byte, padding int) ([]byte, error) {
	text := make([]byte, len(cipherText))
	block, err := des.NewCipher(key)
	if err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `des.NewCipher failed for key "%s"`, key)
		return nil, err
	}

	blockSize := block.BlockSize()
	for i, count := 0, len(text)/blockSize; i < count; i++ {
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Decrypt(text[begin:end], cipherText[begin:end])
	}

	plainText, err := UnPadding(text, padding)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

// EncryptECBTriple 使用TripleDES和ECB模式加密`plainText`。
// `key`的长度应为16或24字节。
// md5:925f2b0f96f4bc56
// ff:加密三重ECB
// plainText:待加密
// key:秘钥
// padding:填充
func EncryptECBTriple(plainText []byte, key []byte, padding int) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "key length error")
	}

	text, err := Padding(plainText, padding)
	if err != nil {
		return nil, err
	}

	var newKey []byte
	if len(key) == 16 {
		newKey = append([]byte{}, key...)
		newKey = append(newKey, key[:8]...)
	} else {
		newKey = append([]byte{}, key...)
	}

	block, err := des.NewTripleDESCipher(newKey)
	if err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `des.NewTripleDESCipher failed for key "%s"`, newKey)
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

// DecryptECBTriple 使用三重DES和ECB模式解密`cipherText`。
// `key`的长度应该是16或24字节。
// md5:c1bf1a1ac477aa6e
// ff:解密三重ECB
// cipherText:待解密
// key:秘钥
// padding:填充
func DecryptECBTriple(cipherText []byte, key []byte, padding int) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "key length error")
	}

	var newKey []byte
	if len(key) == 16 {
		newKey = append([]byte{}, key...)
		newKey = append(newKey, key[:8]...)
	} else {
		newKey = append([]byte{}, key...)
	}

	block, err := des.NewTripleDESCipher(newKey)
	if err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `des.NewTripleDESCipher failed for key "%s"`, newKey)
		return nil, err
	}

	blockSize := block.BlockSize()
	text := make([]byte, len(cipherText))
	for i, count := 0, len(text)/blockSize; i < count; i++ {
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Decrypt(text[begin:end], cipherText[begin:end])
	}

	plainText, err := UnPadding(text, padding)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

// EncryptCBC 使用CBC模式对`plainText`进行加密。 md5:f3a1f35b734799c4
// ff:加密CBC
// plainText:待加密
// key:密钥
// iv:
// padding:填充
func EncryptCBC(plainText []byte, key []byte, iv []byte, padding int) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `des.NewCipher failed for key "%s"`, key)
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid iv length")
	}

	text, err := Padding(plainText, padding)
	if err != nil {
		return nil, err
	}
	cipherText := make([]byte, len(text))

	encryptor := cipher.NewCBCEncrypter(block, iv)
	encryptor.CryptBlocks(cipherText, text)

	return cipherText, nil
}

// 使用CBC模式解密`cipherText`。 md5:232eecbdd9458374
// ff:解密CBC
// cipherText:待解密
// key:密钥
// iv:
// padding:填充
func DecryptCBC(cipherText []byte, key []byte, iv []byte, padding int) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `des.NewCipher failed for key "%s"`, key)
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "iv length invalid")
	}

	text := make([]byte, len(cipherText))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(text, cipherText)

	plainText, err := UnPadding(text, padding)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

// EncryptCBCTriple 使用TripleDES和CBC模式加密`plainText`。 md5:8fb88673fa27c4a0
// ff:加密三重CBC
// plainText:待加密
// key:密钥
// iv:
// padding:填充
func EncryptCBCTriple(plainText []byte, key []byte, iv []byte, padding int) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "key length invalid")
	}

	var newKey []byte
	if len(key) == 16 {
		newKey = append([]byte{}, key...)
		newKey = append(newKey, key[:8]...)
	} else {
		newKey = append([]byte{}, key...)
	}

	block, err := des.NewTripleDESCipher(newKey)
	if err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `des.NewTripleDESCipher failed for key "%s"`, newKey)
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid iv length")
	}

	text, err := Padding(plainText, padding)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, len(text))
	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(cipherText, text)

	return cipherText, nil
}

// DecryptCBCTriple 使用三重DES和CBC模式解密`cipherText`。 md5:53d3127b0b4fcbb0
// ff:解密三重CBC
// cipherText:待解密
// key:密钥
// iv:
// padding:填充
func DecryptCBCTriple(cipherText []byte, key []byte, iv []byte, padding int) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "key length invalid")
	}

	var newKey []byte
	if len(key) == 16 {
		newKey = append([]byte{}, key...)
		newKey = append(newKey, key[:8]...)
	} else {
		newKey = append([]byte{}, key...)
	}

	block, err := des.NewTripleDESCipher(newKey)
	if err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `des.NewTripleDESCipher failed for key "%s"`, newKey)
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid iv length")
	}

	text := make([]byte, len(cipherText))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(text, cipherText)

	plainText, err := UnPadding(text, padding)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

// ff:
// text:
// blockSize:
func PaddingPKCS5(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padText...)
}

// ff:
// text:
func UnPaddingPKCS5(text []byte) []byte {
	length := len(text)
	padText := int(text[length-1])
	return text[:(length - padText)]
}

// ff:
// text:
// padding:
func Padding(text []byte, padding int) ([]byte, error) {
	switch padding {
	case NOPADDING:
		if len(text)%8 != 0 {
			return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid text length")
		}

	case PKCS5PADDING:
		return PaddingPKCS5(text, 8), nil

	default:
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `unsupported padding type "%d"`, padding)
	}

	return text, nil
}

// ff:
// text:
// padding:
func UnPadding(text []byte, padding int) ([]byte, error) {
	switch padding {
	case NOPADDING:
		if len(text)%8 != 0 {
			return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid text length")
		}

	case PKCS5PADDING:
		return UnPaddingPKCS5(text), nil

	default:
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `unsupported padding type "%d"`, padding)
	}
	return text, nil
}
