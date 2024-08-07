// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类

import (
	"strconv"

	gbinary "github.com/888go/goframe/encoding/gbinary"
)

// X取小数32位 将 `any` 转换为 float32 类型。 md5:ae355a911909e343
func X取小数32位(any interface{}) float32 {
	if any == nil {
		return 0
	}
	switch value := any.(type) {
	case float32:
		return value
	case float64:
		return float32(value)
	case []byte:
		return gbinary.DecodeToFloat32(value)
	default:
		if f, ok := value.(iFloat32); ok {
			return f.X取小数32位()
		}
		v, _ := strconv.ParseFloat(String(any), 64)
		return float32(v)
	}
}

// X取小数64位 将 `any` 转换为 float64 类型。 md5:c0bd7cb237571bff
func X取小数64位(any interface{}) float64 {
	if any == nil {
		return 0
	}
	switch value := any.(type) {
	case float32:
		return float64(value)
	case float64:
		return value
	case []byte:
		return gbinary.DecodeToFloat64(value)
	default:
		if f, ok := value.(iFloat64); ok {
			return f.X取小数64位()
		}
		v, _ := strconv.ParseFloat(String(any), 64)
		return v
	}
}
