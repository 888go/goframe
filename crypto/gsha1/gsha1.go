// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gsha1 provides useful API for SHA1 encryption algorithms.
package gsha1//bm:加密sha1类

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

// Encrypt encrypts any type of variable using SHA1 algorithms.
// It uses package gconv to convert `v` to its bytes type.

// ff:加密
// v:值
func Encrypt(v interface{}) string {
	r := sha1.Sum(gconv.Bytes(v))
	return hex.EncodeToString(r[:])
}

// EncryptFile encrypts file content of `path` using SHA1 algorithms.

// ff:加密文件
// err:错误
// encrypt:sha1值
// path:路径
func EncryptFile(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		err = gerror.Wrapf(err, `os.Open failed for name "%s"`, path)
		return "", err
	}
	defer f.Close()
	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		err = gerror.Wrap(err, `io.Copy failed`)
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// MustEncryptFile encrypts file content of `path` using SHA1 algorithms.
// It panics if any error occurs.

// ff:加密文件PANI
// path:路径
func MustEncryptFile(path string) string {
	result, err := EncryptFile(path)
	if err != nil {
		panic(err)
	}
	return result
}
