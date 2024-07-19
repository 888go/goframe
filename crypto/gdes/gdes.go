// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gdes provides useful API for DES encryption/decryption algorithms.
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

// EncryptECB encrypts `plainText` using ECB mode.
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

// DecryptECB decrypts `cipherText` using ECB mode.
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

// EncryptECBTriple encrypts `plainText` using TripleDES and ECB mode.
// The length of the `key` should be either 16 or 24 bytes.
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

// DecryptECBTriple decrypts `cipherText` using TripleDES and ECB mode.
// The length of the `key` should be either 16 or 24 bytes.
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

// EncryptCBC encrypts `plainText` using CBC mode.
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

// DecryptCBC decrypts `cipherText` using CBC mode.
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

// EncryptCBCTriple encrypts `plainText` using TripleDES and CBC mode.
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

// DecryptCBCTriple decrypts `cipherText` using TripleDES and CBC mode.
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
