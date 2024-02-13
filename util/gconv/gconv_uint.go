// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类

import (
	"math"
	"strconv"
	
	"github.com/888go/goframe/encoding/gbinary"
)

// Uint 将 `any` 类型转换为 uint 类型。
func X取正整数(值 interface{}) uint {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(uint); ok {
		return v
	}
	return uint(X取正整数64位(值))
}

// Uint8将`any`转换为uint8类型。
func X取正整数8位(值 interface{}) uint8 {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(uint8); ok {
		return v
	}
	return uint8(X取正整数64位(值))
}

// Uint16将`any`转换为uint16。
func X取正整数16位(值 interface{}) uint16 {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(uint16); ok {
		return v
	}
	return uint16(X取正整数64位(值))
}

// Uint32将`any`转换为uint32。
func X取正整数32位(值 interface{}) uint32 {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(uint32); ok {
		return v
	}
	return uint32(X取正整数64位(值))
}

// Uint64将`any`转换为uint64类型。
func X取正整数64位(值 interface{}) uint64 {
	if 值 == nil {
		return 0
	}
	switch value := 值.(type) {
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
		return 字节集类.DecodeToUint64(value)
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
