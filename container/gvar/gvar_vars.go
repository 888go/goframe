// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gvar
import (
	"github.com/888go/goframe/util/gconv"
	)
// Vars 是一个 *Var 类型的切片。
type Vars []*Var

// Strings 将 `vs` 转换并返回为 []string 类型的切片。
func (vs Vars) Strings() (s []string) {
	for _, v := range vs {
		s = append(s, v.String())
	}
	return s
}

// Interfaces将`vs`转换并返回为[]interface{}类型。
func (vs Vars) Interfaces() (s []interface{}) {
	for _, v := range vs {
		s = append(s, v.Val())
	}
	return s
}

// Float32s将`vs`转换并返回为[]float32类型。
func (vs Vars) Float32s() (s []float32) {
	for _, v := range vs {
		s = append(s, v.Float32())
	}
	return s
}

// Float64s 将 `vs` 转换并返回为 []float64 类型的切片。
func (vs Vars) Float64s() (s []float64) {
	for _, v := range vs {
		s = append(s, v.Float64())
	}
	return s
}

// Ints将`vs`转换并返回为[]Int类型。
func (vs Vars) Ints() (s []int) {
	for _, v := range vs {
		s = append(s, v.Int())
	}
	return s
}

// Int8s将`vs`转换并返回为[]int8类型。
func (vs Vars) Int8s() (s []int8) {
	for _, v := range vs {
		s = append(s, v.Int8())
	}
	return s
}

// Int16s将`vs`转换并返回为[]int16类型。
func (vs Vars) Int16s() (s []int16) {
	for _, v := range vs {
		s = append(s, v.Int16())
	}
	return s
}

// Int32s 将 `vs` 转换并返回为 []int32 类型的切片。
func (vs Vars) Int32s() (s []int32) {
	for _, v := range vs {
		s = append(s, v.Int32())
	}
	return s
}

// Int64s将`vs`转换并返回为[]int64类型。
func (vs Vars) Int64s() (s []int64) {
	for _, v := range vs {
		s = append(s, v.Int64())
	}
	return s
}

// Uints 将 `vs` 转换并返回为 []uint 类型的切片。
func (vs Vars) Uints() (s []uint) {
	for _, v := range vs {
		s = append(s, v.Uint())
	}
	return s
}

// Uint8s将`vs`转换并返回为[]uint8类型。
func (vs Vars) Uint8s() (s []uint8) {
	for _, v := range vs {
		s = append(s, v.Uint8())
	}
	return s
}

// Uint16s 将 `vs` 转换并返回为 []uint16 类型的切片。
func (vs Vars) Uint16s() (s []uint16) {
	for _, v := range vs {
		s = append(s, v.Uint16())
	}
	return s
}

// Uint32s 将 `vs` 转换并返回为 []uint32 类型的切片。
func (vs Vars) Uint32s() (s []uint32) {
	for _, v := range vs {
		s = append(s, v.Uint32())
	}
	return s
}

// Uint64s 将 `vs` 转换并返回为 []uint64 类型的切片。
func (vs Vars) Uint64s() (s []uint64) {
	for _, v := range vs {
		s = append(s, v.Uint64())
	}
	return s
}

// Scan 将 `vs` 转换为 []struct 或 []*struct 类型的切片。
func (vs Vars) Scan(pointer interface{}, mapping ...map[string]string) error {
	return gconv.Structs(vs.Interfaces(), pointer, mapping...)
}
