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

// Int将`any`转换为int。
func X取整数(值 interface{}) int {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(int); ok {
		return v
	}
	return int(X取整数64位(值))
}

// Int8将`any`转换为int8类型。
func X取整数8位(值 interface{}) int8 {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(int8); ok {
		return v
	}
	return int8(X取整数64位(值))
}

// Int16将`any`转换为int16类型。
func X取整数16位(值 interface{}) int16 {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(int16); ok {
		return v
	}
	return int16(X取整数64位(值))
}

// Int32将`any`转换为int32类型。
func X取整数32位(值 interface{}) int32 {
	if 值 == nil {
		return 0
	}
	if v, ok := 值.(int32); ok {
		return v
	}
	return int32(X取整数64位(值))
}

// Int64将`any`转换为int64类型。
func X取整数64位(值 interface{}) int64 {
	if 值 == nil {
		return 0
	}
	switch value := 值.(type) {
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
		return 字节集类.DecodeToInt64(value)
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
