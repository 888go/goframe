// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gcrc32 provides useful API for CRC32 encryption algorithms.
package 加密crc32类

import (
	"hash/crc32"

	gconv "github.com/888go/goframe/util/gconv"
)

// Encrypt encrypts any type of variable using CRC32 algorithms.
// It uses gconv package to convert `v` to its bytes type.
func Encrypt(v interface{}) uint32 {
	return crc32.ChecksumIEEE(gconv.Bytes(v))
}
