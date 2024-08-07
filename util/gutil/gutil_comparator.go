// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 工具类

import (
	"strings"

	gconv "github.com/888go/goframe/util/gconv"
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

// X比较文本 提供了一个快速的字符串比较功能。 md5:d0e6023eaf1cf6f5
func X比较文本(a, b interface{}) int {
	return strings.Compare(gconv.String(a), gconv.String(b))
}

// X比较整数 在int类型上提供基本的比较功能。 md5:c8660a6b9a98a8f5
func X比较整数(a, b interface{}) int {
	return gconv.X取整数(a) - gconv.X取整数(b)
}

// X比较整数8位 提供了对int8的基本比较操作。 md5:b727bd550fef2616
func X比较整数8位(a, b interface{}) int {
	return int(gconv.X取整数8位(a) - gconv.X取整数8位(b))
}

// X比较整数16位 提供对 int16 的基本比较。 md5:9cf435a64f01ae86
func X比较整数16位(a, b interface{}) int {
	return int(gconv.X取整数16位(a) - gconv.X取整数16位(b))
}

// X比较整数32位 提供对 int32 的基本比较。 md5:b491ec4c9c46cfde
func X比较整数32位(a, b interface{}) int {
	return int(gconv.X取整数32位(a) - gconv.X取整数32位(b))
}

// X比较整数64位 提供了对int64类型的基本比较功能。 md5:0a827d5d4781dc3b
func X比较整数64位(a, b interface{}) int {
	return int(gconv.X取整数64位(a) - gconv.X取整数64位(b))
}

// X比较正整数 提供了一个基础的无符号整数比较功能。 md5:a11699413547d752
func X比较正整数(a, b interface{}) int {
	return int(gconv.X取正整数(a) - gconv.X取正整数(b))
}

// X比较正整数8位 提供uint8的基本比较。 md5:5f83ec701ceeb5c2
func X比较正整数8位(a, b interface{}) int {
	return int(gconv.X取正整数8位(a) - gconv.X取正整数8位(b))
}

// X比较正整数16位 提供对 uint16 的基本比较。 md5:74670078ce1de876
func X比较正整数16位(a, b interface{}) int {
	return int(gconv.X取正整数16位(a) - gconv.X取正整数16位(b))
}

// X比较正整数32位 在uint32类型上提供基本的比较功能。 md5:f08e9a1b63d90f06
func X比较正整数32位(a, b interface{}) int {
	return int(gconv.X取正整数32位(a) - gconv.X取正整数32位(b))
}

// X比较正整数64位 提供了对 uint64 类型的基本比较功能。 md5:e08126378b4b83b9
func X比较正整数64位(a, b interface{}) int {
	return int(gconv.X取正整数64位(a) - gconv.X取正整数64位(b))
}

// X比较小数32位 提供对float32的基本比较。 md5:5f3a1a46408ed33d
func X比较小数32位(a, b interface{}) int {
	aFloat := gconv.X取小数32位(a)
	bFloat := gconv.X取小数32位(b)
	if aFloat == bFloat {
		return 0
	}
	if aFloat > bFloat {
		return 1
	}
	return -1
}

// X比较小数64位 提供对float64的基本比较。 md5:15afd5d88c222774
func X比较小数64位(a, b interface{}) int {
	aFloat := gconv.X取小数64位(a)
	bFloat := gconv.X取小数64位(b)
	if aFloat == bFloat {
		return 0
	}
	if aFloat > bFloat {
		return 1
	}
	return -1
}

// X比较字节 在字节上提供基本的比较功能。 md5:74dad260972241f7
func X比较字节(a, b interface{}) int {
	return int(gconv.X取字节(a) - gconv.X取字节(b))
}

// X比较字符 提供了一个基本的 rune 对比功能。 md5:e30baf0412d73c95
func X比较字符(a, b interface{}) int {
	return int(gconv.X取字符(a) - gconv.X取字符(b))
}

// X比较时间 提供了对time.Time的基本比较。 md5:58d0e4c6e29cff00
func X比较时间(a, b interface{}) int {
	aTime := gconv.X取时间(a)
	bTime := gconv.X取时间(b)
	switch {
	case aTime.After(bTime):
		return 1
	case aTime.Before(bTime):
		return -1
	default:
		return 0
	}
}
