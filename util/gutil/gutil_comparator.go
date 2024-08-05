// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gutil

import (
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
)

// Comparator 是一个函数，用于比较 a 和 b，并返回一个整数结果。
//
// 应该返回以下数值：
//
// 如果 a < b，则返回负数
// 如果 a == b，则返回零
// 如果 a > b，则返回正数
// md5:c1e00d03f187b96f
type Comparator func(a, b interface{}) int

// ComparatorString 提供了一个快速的字符串比较功能。 md5:d0e6023eaf1cf6f5
func ComparatorString(a, b interface{}) int {
	return strings.Compare(gconv.String(a), gconv.String(b))
}

// ComparatorInt 在int类型上提供基本的比较功能。 md5:c8660a6b9a98a8f5
func ComparatorInt(a, b interface{}) int {
	return gconv.Int(a) - gconv.Int(b)
}

// ComparatorInt8 提供了对int8的基本比较操作。 md5:b727bd550fef2616
func ComparatorInt8(a, b interface{}) int {
	return int(gconv.Int8(a) - gconv.Int8(b))
}

// ComparatorInt16 提供对 int16 的基本比较。 md5:9cf435a64f01ae86
func ComparatorInt16(a, b interface{}) int {
	return int(gconv.Int16(a) - gconv.Int16(b))
}

// ComparatorInt32 提供对 int32 的基本比较。 md5:b491ec4c9c46cfde
func ComparatorInt32(a, b interface{}) int {
	return int(gconv.Int32(a) - gconv.Int32(b))
}

// ComparatorInt64 提供了对int64类型的基本比较功能。 md5:0a827d5d4781dc3b
func ComparatorInt64(a, b interface{}) int {
	return int(gconv.Int64(a) - gconv.Int64(b))
}

// ComparatorUint 提供了一个基础的无符号整数比较功能。 md5:a11699413547d752
func ComparatorUint(a, b interface{}) int {
	return int(gconv.Uint(a) - gconv.Uint(b))
}

// ComparatorUint8 提供uint8的基本比较。 md5:5f83ec701ceeb5c2
func ComparatorUint8(a, b interface{}) int {
	return int(gconv.Uint8(a) - gconv.Uint8(b))
}

// ComparatorUint16 提供对 uint16 的基本比较。 md5:74670078ce1de876
func ComparatorUint16(a, b interface{}) int {
	return int(gconv.Uint16(a) - gconv.Uint16(b))
}

// ComparatorUint32 在uint32类型上提供基本的比较功能。 md5:f08e9a1b63d90f06
func ComparatorUint32(a, b interface{}) int {
	return int(gconv.Uint32(a) - gconv.Uint32(b))
}

// ComparatorUint64 提供了对 uint64 类型的基本比较功能。 md5:e08126378b4b83b9
func ComparatorUint64(a, b interface{}) int {
	return int(gconv.Uint64(a) - gconv.Uint64(b))
}

// ComparatorFloat32 提供对float32的基本比较。 md5:5f3a1a46408ed33d
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

// ComparatorFloat64 提供对float64的基本比较。 md5:15afd5d88c222774
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

// ComparatorByte 在字节上提供基本的比较功能。 md5:74dad260972241f7
func ComparatorByte(a, b interface{}) int {
	return int(gconv.Byte(a) - gconv.Byte(b))
}

// ComparatorRune 提供了一个基本的 rune 对比功能。 md5:e30baf0412d73c95
func ComparatorRune(a, b interface{}) int {
	return int(gconv.Rune(a) - gconv.Rune(b))
}

// ComparatorTime 提供了对time.Time的基本比较。 md5:58d0e4c6e29cff00
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
