// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类

import (
	"math"
	"strconv"

	gbinary "github.com/888go/goframe/encoding/gbinary"
)

// X取正整数 将 `any` 转换为 uint 类型。 md5:0a9e343966117c44
func X取正整数(值 interface{}) uint {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(uint); ok {
		return v
	}
	return uint(X取正整数64位(值))
}

// X取正整数8位 将 `any` 类型转换为 uint8 类型。 md5:330b11711227c6f3
func X取正整数8位(值 interface{}) uint8 {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(uint8); ok {
		return v
	}
	return uint8(X取正整数64位(值))
}

// X取正整数16位 将 `any` 转换为 uint16。 md5:7763a0b90bd179e4
func X取正整数16位(值 interface{}) uint16 {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(uint16); ok {
		return v
	}
	return uint16(X取正整数64位(值))
}

// X取正整数32位 将 `any` 转换为 uint32 类型。 md5:26c3b815ff56d271
func X取正整数32位(值 interface{}) uint32 {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(uint32); ok {
		return v
	}
	return uint32(X取正整数64位(值))
}

// X取正整数64位 将 `any` 转换为 uint64 类型。 md5:14c7f76fe909ea8e
func X取正整数64位(any interface{}) uint64 {
	if any == nil {
		return 0
	}
	switch value := any.(type) {
	case int:
		return uint64(value)
	case int8:
		return uint64(value)
	case int16:
		return uint64(value)
	case int32:
		return uint64(value)
	case int64:
		return uint64(value)
	case uint:
		return uint64(value)
	case uint8:
		return uint64(value)
	case uint16:
		return uint64(value)
	case uint32:
		return uint64(value)
	case uint64:
		return value
	case float32:
		return uint64(value)
	case float64:
		return uint64(value)
	case bool:
		if value {
			return 1
		}
		return 0
	case []byte:
		return gbinary.DecodeToUint64(value)
	default:
		if f, ok := value.(iUint64); ok {
			return f.X取正整数64位()
		}
		s := String(value)
		// Hexadecimal
		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			if v, e := strconv.ParseUint(s[2:], 16, 64); e == nil {
				return v
			}
		}
		// Decimal
		if v, e := strconv.ParseUint(s, 10, 64); e == nil {
			return v
		}
		// Float64
		if valueFloat64 := X取小数64位(value); math.IsNaN(valueFloat64) {
			return 0
		} else {
			return uint64(valueFloat64)
		}
	}
}
