// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gutil

import (
	"strings"
	
	"github.com/gogf/gf/v2/util/gconv"
)

// Comparator 是一个函数，用于比较 a 和 b，并返回一个整数形式的结果。
//
// 应该返回一个数字：
//
// 如果 a < b，则返回负数
// 如果 a == b，则返回零
// 如果 a > b，则返回正数
type Comparator func(a, b interface{}) int

// ComparatorString 提供了一个对字符串进行快速比较的功能。
func ComparatorString(a, b interface{}) int {
	return strings.Compare(gconv.String(a), gconv.String(b))
}

// ComparatorInt 提供了一个对 int 类型的基本比较功能。
func ComparatorInt(a, b interface{}) int {
	return gconv.Int(a) - gconv.Int(b)
}

// ComparatorInt8 提供了对 int8 类型的基本比较功能。
func ComparatorInt8(a, b interface{}) int {
	return int(gconv.Int8(a) - gconv.Int8(b))
}

// ComparatorInt16 提供了对 int16 类型的基本比较功能。
func ComparatorInt16(a, b interface{}) int {
	return int(gconv.Int16(a) - gconv.Int16(b))
}

// ComparatorInt32 提供了对 int32 类型的基本比较功能。
func ComparatorInt32(a, b interface{}) int {
	return int(gconv.Int32(a) - gconv.Int32(b))
}

// ComparatorInt64 提供了对 int64 类型的基本比较功能。
func ComparatorInt64(a, b interface{}) int {
	return int(gconv.Int64(a) - gconv.Int64(b))
}

// ComparatorUint 提供了一个在 uint 类型上的基本比较功能。
func ComparatorUint(a, b interface{}) int {
	return int(gconv.Uint(a) - gconv.Uint(b))
}

// ComparatorUint8 提供了一个基本的 uint8 类型比较功能。
func ComparatorUint8(a, b interface{}) int {
	return int(gconv.Uint8(a) - gconv.Uint8(b))
}

// ComparatorUint16 提供了对 uint16 类型的基本比较功能。
func ComparatorUint16(a, b interface{}) int {
	return int(gconv.Uint16(a) - gconv.Uint16(b))
}

// ComparatorUint32 提供了对 uint32 类型的基本比较功能。
func ComparatorUint32(a, b interface{}) int {
	return int(gconv.Uint32(a) - gconv.Uint32(b))
}

// ComparatorUint64 提供了一个在 uint64 类型上的基础比较功能。
func ComparatorUint64(a, b interface{}) int {
	return int(gconv.Uint64(a) - gconv.Uint64(b))
}

// ComparatorFloat32 提供了对 float32 类型的基本比较功能。
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

// ComparatorFloat64 提供了对 float64 类型的基本比较功能。
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

// ComparatorByte 提供了一个在 byte 类型上的基础比较功能。
func ComparatorByte(a, b interface{}) int {
	return int(gconv.Byte(a) - gconv.Byte(b))
}

// ComparatorRune 提供了对 rune 类型的基本比较功能。
func ComparatorRune(a, b interface{}) int {
	return int(gconv.Rune(a) - gconv.Rune(b))
}

// ComparatorTime 提供了对 time.Time 类型的基本比较功能。
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
