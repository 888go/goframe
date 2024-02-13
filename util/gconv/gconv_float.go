// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类

import (
	"strconv"
	
	"github.com/888go/goframe/encoding/gbinary"
)

// Float32将`any`转换为float32类型。
func X取小数32位(值 interface{}) float32 {
	if 值 == nil {
		return 0
	}
	switch value := 值.(type) {
	case float32:
		return value
	case float64:
		return float32(value)
	case []byte:
		return 字节集类.DecodeToFloat32(value)
	default:
		if f, ok := value.(iFloat32); ok {
			return f.X取小数32位()
		}
		v, _ := strconv.ParseFloat(String(值), 64)
		return float32(v)
	}
}

// Float64将`any`转换为float64类型。
func X取小数64位(值 interface{}) float64 {
	if 值 == nil {
		return 0
	}
	switch value := 值.(type) {
	case float32:
		return float64(value)
	case float64:
		return value
	case []byte:
		return 字节集类.DecodeToFloat64(value)
	default:
		if f, ok := value.(iFloat64); ok {
			return f.X取小数64位()
		}
		v, _ := strconv.ParseFloat(String(值), 64)
		return v
	}
}
