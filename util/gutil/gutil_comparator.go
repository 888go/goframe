// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gutil//bm:工具类

import (
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
)

// Comparator is a function that compare a and b, and returns the result as int.
//
// Should return a number:
//
//	negative , if a < b
//	zero     , if a == b
//	positive , if a > b
type Comparator func(a, b interface{}) int

// ComparatorString provides a fast comparison on strings.

// ff:比较文本
// b:
// a:
func ComparatorString(a, b interface{}) int {
	return strings.Compare(gconv.String(a), gconv.String(b))
}

// ComparatorInt provides a basic comparison on int.

// ff:比较整数
// b:
// a:
func ComparatorInt(a, b interface{}) int {
	return gconv.Int(a) - gconv.Int(b)
}

// ComparatorInt8 provides a basic comparison on int8.

// ff:比较整数8位
// b:
// a:
func ComparatorInt8(a, b interface{}) int {
	return int(gconv.Int8(a) - gconv.Int8(b))
}

// ComparatorInt16 provides a basic comparison on int16.

// ff:比较整数16位
// b:
// a:
func ComparatorInt16(a, b interface{}) int {
	return int(gconv.Int16(a) - gconv.Int16(b))
}

// ComparatorInt32 provides a basic comparison on int32.

// ff:比较整数32位
// b:
// a:
func ComparatorInt32(a, b interface{}) int {
	return int(gconv.Int32(a) - gconv.Int32(b))
}

// ComparatorInt64 provides a basic comparison on int64.

// ff:比较整数64位
// b:
// a:
func ComparatorInt64(a, b interface{}) int {
	return int(gconv.Int64(a) - gconv.Int64(b))
}

// ComparatorUint provides a basic comparison on uint.

// ff:比较正整数
// b:
// a:
func ComparatorUint(a, b interface{}) int {
	return int(gconv.Uint(a) - gconv.Uint(b))
}

// ComparatorUint8 provides a basic comparison on uint8.

// ff:比较正整数8位
// b:
// a:
func ComparatorUint8(a, b interface{}) int {
	return int(gconv.Uint8(a) - gconv.Uint8(b))
}

// ComparatorUint16 provides a basic comparison on uint16.

// ff:比较正整数16位
// b:
// a:
func ComparatorUint16(a, b interface{}) int {
	return int(gconv.Uint16(a) - gconv.Uint16(b))
}

// ComparatorUint32 provides a basic comparison on uint32.

// ff:比较正整数32位
// b:
// a:
func ComparatorUint32(a, b interface{}) int {
	return int(gconv.Uint32(a) - gconv.Uint32(b))
}

// ComparatorUint64 provides a basic comparison on uint64.

// ff:比较正整数64位
// b:
// a:
func ComparatorUint64(a, b interface{}) int {
	return int(gconv.Uint64(a) - gconv.Uint64(b))
}

// ComparatorFloat32 provides a basic comparison on float32.

// ff:比较小数32位
// b:
// a:
func ComparatorFloat32(a, b interface{}) int {
	aFloat := gconv.Float32(a)
	bFloat := gconv.Float32(b)
	if aFloat == bFloat {
		return 0
	}
	if aFloat > bFloat {
		return 1
	}
	return -1
}

// ComparatorFloat64 provides a basic comparison on float64.

// ff:比较小数64位
// b:
// a:
func ComparatorFloat64(a, b interface{}) int {
	aFloat := gconv.Float64(a)
	bFloat := gconv.Float64(b)
	if aFloat == bFloat {
		return 0
	}
	if aFloat > bFloat {
		return 1
	}
	return -1
}

// ComparatorByte provides a basic comparison on byte.

// ff:比较字节
// b:
// a:
func ComparatorByte(a, b interface{}) int {
	return int(gconv.Byte(a) - gconv.Byte(b))
}

// ComparatorRune provides a basic comparison on rune.

// ff:比较字符
// b:
// a:
func ComparatorRune(a, b interface{}) int {
	return int(gconv.Rune(a) - gconv.Rune(b))
}

// ComparatorTime provides a basic comparison on time.Time.

// ff:比较时间
// b:
// a:
func ComparatorTime(a, b interface{}) int {
	aTime := gconv.Time(a)
	bTime := gconv.Time(b)
	switch {
	case aTime.After(bTime):
		return 1
	case aTime.Before(bTime):
		return -1
	default:
		return 0
	}
}
