// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

import "unsafe"

// UnsafeStrToBytes converts string to []byte without memory copy.
// Note that, if you completely sure you will never use `s` variable in the feature,
// you can use this unsafe function to implement type conversion in high performance.

// ff:文本到字节集_非安全
// s:文本
func UnsafeStrToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// UnsafeBytesToStr converts []byte to string without memory copy.
// Note that, if you completely sure you will never use `b` variable in the feature,
// you can use this unsafe function to implement type conversion in high performance.

// ff:字节集到文本_非安全
// b:字节集
func UnsafeBytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
