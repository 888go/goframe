// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 泛型类

import (
	"github.com/gogf/gf/v2/util/gconv"
)

// Vars 是一个 *Var 类型的切片。
type Vars []*Var

// Strings 将 `vs` 转换并返回为 []string 类型的切片。
func (vs Vars) X取文本切片() (切片 []string) {
	for _, v := range vs {
		切片 = append(切片, v.String())
	}
	return 切片
}

// Interfaces将`vs`转换并返回为[]interface{}类型。
func (vs Vars) X取any切片() (切片 []interface{}) {
	for _, v := range vs {
		切片 = append(切片, v.X取值())
	}
	return 切片
}

// Float32s将`vs`转换并返回为[]float32类型。
func (vs Vars) X取小数32位切片() (切片 []float32) {
	for _, v := range vs {
		切片 = append(切片, v.X取小数32位())
	}
	return 切片
}

// Float64s 将 `vs` 转换并返回为 []float64 类型的切片。
func (vs Vars) X取小数64位切片() (切片 []float64) {
	for _, v := range vs {
		切片 = append(切片, v.X取小数64位())
	}
	return 切片
}

// Ints将`vs`转换并返回为[]Int类型。
func (vs Vars) X取整数切片() (切片 []int) {
	for _, v := range vs {
		切片 = append(切片, v.X取整数())
	}
	return 切片
}

// Int8s将`vs`转换并返回为[]int8类型。
func (vs Vars) X取整数8位切片() (切片 []int8) {
	for _, v := range vs {
		切片 = append(切片, v.X取整数8位())
	}
	return 切片
}

// Int16s将`vs`转换并返回为[]int16类型。
func (vs Vars) X取整数16位切片() (切片 []int16) {
	for _, v := range vs {
		切片 = append(切片, v.X取整数16位())
	}
	return 切片
}

// Int32s 将 `vs` 转换并返回为 []int32 类型的切片。
func (vs Vars) X取整数32位切片() (切片 []int32) {
	for _, v := range vs {
		切片 = append(切片, v.X取整数32位())
	}
	return 切片
}

// Int64s将`vs`转换并返回为[]int64类型。
func (vs Vars) X取整数64位切片() (切片 []int64) {
	for _, v := range vs {
		切片 = append(切片, v.X取整数64位())
	}
	return 切片
}

// Uints 将 `vs` 转换并返回为 []uint 类型的切片。
func (vs Vars) X取正整数切片() (切片 []uint) {
	for _, v := range vs {
		切片 = append(切片, v.X取正整数())
	}
	return 切片
}

// Uint8s将`vs`转换并返回为[]uint8类型。
func (vs Vars) X取正整数8位切片() (切片 []uint8) {
	for _, v := range vs {
		切片 = append(切片, v.X取正整数8位())
	}
	return 切片
}

// Uint16s 将 `vs` 转换并返回为 []uint16 类型的切片。
func (vs Vars) X取正整数16位切片() (切片 []uint16) {
	for _, v := range vs {
		切片 = append(切片, v.X取正整数16位())
	}
	return 切片
}

// Uint32s 将 `vs` 转换并返回为 []uint32 类型的切片。
func (vs Vars) X取正整数32位切片() (切片 []uint32) {
	for _, v := range vs {
		切片 = append(切片, v.X取正整数32位())
	}
	return 切片
}

// Uint64s 将 `vs` 转换并返回为 []uint64 类型的切片。
func (vs Vars) X取正整数64位切片() (切片 []uint64) {
	for _, v := range vs {
		切片 = append(切片, v.X取正整数64位())
	}
	return 切片
}

// Scan 将 `vs` 转换为 []struct 或 []*struct 类型的切片。
func (vs Vars) X取结构体指针(结构体指针 interface{}, 名称映射 ...map[string]string) error {
	return gconv.Structs(vs.X取any切片(), 结构体指针, 名称映射...)
}
