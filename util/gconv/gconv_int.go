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

// X取整数 将 `any` 转换为整数。 md5:7d14e0d4caedf28f
func X取整数(值 interface{}) int {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(int); ok {
		return v
	}
	return int(X取整数64位(值))
}

// X取整数8位 将 `any` 类型转换为 int8。 md5:3a0a5e4022ab71df
func X取整数8位(值 interface{}) int8 {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(int8); ok {
		return v
	}
	return int8(X取整数64位(值))
}

// X取整数16位 将 `any` 转换为 int16。 md5:5acb21515934f62d
func X取整数16位(值 interface{}) int16 {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(int16); ok {
		return v
	}
	return int16(X取整数64位(值))
}

// X取整数32位 将 `any` 转换为 int32 类型。 md5:87b6a70499ffaf8a
func X取整数32位(值 interface{}) int32 {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(int32); ok {
		return v
	}
	return int32(X取整数64位(值))
}

// X取整数64位 将 `any` 转换为 int64 类型。 md5:e95a0233b73ea496
func X取整数64位(any interface{}) int64 {
	if any == nil {
		return 0
	}
	switch value := any.(type) {
	case int:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case uint:
		return int64(value)
	case uint8:
		return int64(value)
	case uint16:
		return int64(value)
	case uint32:
		return int64(value)
	case uint64:
		return int64(value)
	case float32:
		return int64(value)
	case float64:
		return int64(value)
	case bool:
		if value {
			return 1
		}
		return 0
	case []byte:
		return gbinary.DecodeToInt64(value)
	default:
		if f, ok := value.(iInt64); ok {
			return f.X取整数64位()
		}
		var (
			s       = String(value)
			isMinus = false
		)
		if len(s) > 0 {
			if s[0] == '-' {
				isMinus = true
				s = s[1:]
			} else if s[0] == '+' {
				s = s[1:]
			}
		}
		// Hexadecimal
		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			if v, e := strconv.ParseInt(s[2:], 16, 64); e == nil {
				if isMinus {
					return -v
				}
				return v
			}
		}
		// Decimal
		if v, e := strconv.ParseInt(s, 10, 64); e == nil {
			if isMinus {
				return -v
			}
			return v
		}
		// Float64
		if valueInt64 := X取小数64位(value); math.IsNaN(valueInt64) {
			return 0
		} else {
			return int64(valueInt64)
		}
	}
}
