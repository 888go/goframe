// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gdes 提供了用于DES加密/解密算法的有用API。
package gdes
import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	)
const (
	NOPADDING = iota
	PKCS5PADDING
)

// EncryptECB 使用ECB模式加密`plainText`。
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

// DecryptECB 使用ECB模式解密`cipherText`。
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

// EncryptECBTriple 使用TripleDES加密算法及ECB模式加密`plainText`。
// `key`的长度应为16字节或24字节。
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

// DecryptECBTriple 使用TripleDES和ECB模式解密`cipherText`。
// `key`的长度应为16字节或24字节。
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

// EncryptCBC 使用CBC模式加密`plainText`。
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

// DecryptCBC 使用CBC模式解密`cipherText`。
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

// EncryptCBCTriple 使用TripleDES算法和CBC模式加密`plainText`。
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

// DecryptCBCTriple 使用3DES加密算法和CBC模式解密`cipherText`。
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
			return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid text length")
		}

	case PKCS5PADDING:
		return PaddingPKCS5(text, 8), nil

	default:
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `unsupported padding type "%d"`, padding)
	}

	return text, nil
}

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
