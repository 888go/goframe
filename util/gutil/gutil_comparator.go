// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类

import (
	"strings"
	
	"github.com/888go/goframe/util/gconv"
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
func X比较文本(a, b interface{}) int {
	return strings.Compare(转换类.String(a), 转换类.String(b))
}

// ComparatorInt 提供了一个对 int 类型的基本比较功能。
func X比较整数(a, b interface{}) int {
	return 转换类.X取整数(a) - 转换类.X取整数(b)
}

// ComparatorInt8 提供了对 int8 类型的基本比较功能。
func X比较整数8位(a, b interface{}) int {
	return int(转换类.X取整数8位(a) - 转换类.X取整数8位(b))
}

// ComparatorInt16 提供了对 int16 类型的基本比较功能。
func X比较整数16位(a, b interface{}) int {
	return int(转换类.X取整数16位(a) - 转换类.X取整数16位(b))
}

// ComparatorInt32 提供了对 int32 类型的基本比较功能。
func X比较整数32位(a, b interface{}) int {
	return int(转换类.X取整数32位(a) - 转换类.X取整数32位(b))
}

// ComparatorInt64 提供了对 int64 类型的基本比较功能。
func X比较整数64位(a, b interface{}) int {
	return int(转换类.X取整数64位(a) - 转换类.X取整数64位(b))
}

// ComparatorUint 提供了一个在 uint 类型上的基本比较功能。
func X比较正整数(a, b interface{}) int {
	return int(转换类.X取正整数(a) - 转换类.X取正整数(b))
}

// ComparatorUint8 提供了一个基本的 uint8 类型比较功能。
func X比较正整数8位(a, b interface{}) int {
	return int(转换类.X取正整数8位(a) - 转换类.X取正整数8位(b))
}

// ComparatorUint16 提供了对 uint16 类型的基本比较功能。
func X比较正整数16位(a, b interface{}) int {
	return int(转换类.X取正整数16位(a) - 转换类.X取正整数16位(b))
}

// ComparatorUint32 提供了对 uint32 类型的基本比较功能。
func X比较正整数32位(a, b interface{}) int {
	return int(转换类.X取正整数32位(a) - 转换类.X取正整数32位(b))
}

// ComparatorUint64 提供了一个在 uint64 类型上的基础比较功能。
func X比较正整数64位(a, b interface{}) int {
	return int(转换类.X取正整数64位(a) - 转换类.X取正整数64位(b))
}

// ComparatorFloat32 提供了对 float32 类型的基本比较功能。
func X比较小数32位(a, b interface{}) int {
	aFloat := 转换类.X取小数32位(a)
	bFloat := 转换类.X取小数32位(b)
	if aFloat == bFloat {
		return 0
	}
	if aFloat > bFloat {
		return 1
	}
	return -1
}

// ComparatorFloat64 提供了对 float64 类型的基本比较功能。
func X比较小数64位(a, b interface{}) int {
	aFloat := 转换类.X取小数64位(a)
	bFloat := 转换类.X取小数64位(b)
	if aFloat == bFloat {
		return 0
	}
	if aFloat > bFloat {
		return 1
	}
	return -1
}

// ComparatorByte 提供了一个在 byte 类型上的基础比较功能。
func X比较字节(a, b interface{}) int {
	return int(转换类.X取字节(a) - 转换类.X取字节(b))
}

// ComparatorRune 提供了对 rune 类型的基本比较功能。
func X比较字符(a, b interface{}) int {
	return int(转换类.X取字符(a) - 转换类.X取字符(b))
}

// ComparatorTime 提供了对 time.Time 类型的基本比较功能。
func X比较时间(a, b interface{}) int {
	aTime := 转换类.X取时间(a)
	bTime := 转换类.X取时间(b)
	switch {
	case aTime.After(bTime):
		return 1
	case aTime.Before(bTime):
		return -1
	default:
		return 0
	}
}
