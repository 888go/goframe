// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gvar

import (
	"github.com/gogf/gf/v2/util/gconv"
)

// Ints 将 `v` 转换并返回为 []int 类型的切片。
func (v *Var) Ints() []int {
	return gconv.Ints(v.Val())
}

// Int64s 将 `v` 转换并返回为 []int64 类型的切片。
func (v *Var) Int64s() []int64 {
	return gconv.Int64s(v.Val())
}

// Uints 将 `v` 转换并返回为 []uint 类型。
func (v *Var) Uints() []uint {
	return gconv.Uints(v.Val())
}

// Uint64s将`v`转换并返回为[]uint64类型。
func (v *Var) Uint64s() []uint64 {
	return gconv.Uint64s(v.Val())
}

// Floats 是 Float64s 的别名。
func (v *Var) Floats() []float64 {
	return gconv.Floats(v.Val())
}

// Float32s 将 `v` 转换并返回为 []float32 类型的切片。
func (v *Var) Float32s() []float32 {
	return gconv.Float32s(v.Val())
}

// Float64s将`v`转换并返回为[]float64类型。
func (v *Var) Float64s() []float64 {
	return gconv.Float64s(v.Val())
}

// Strings 将 `v` 转换并返回为 []string 类型。
func (v *Var) Strings() []string {
	return gconv.Strings(v.Val())
}

// Interfaces 将 `v` 转换并返回为 []interfaces{} 类型的切片。
func (v *Var) Interfaces() []interface{} {
	return gconv.Interfaces(v.Val())
}

// Slice 是 Interfaces 的别名。
func (v *Var) Slice() []interface{} {
	return v.Interfaces()
}

// Array 是 Interfaces 的别名。
func (v *Var) Array() []interface{} {
	return v.Interfaces()
}

// Vars 将 `v` 转换并返回为 []Var 类型的切片。
func (v *Var) Vars() []*Var {
	array := gconv.Interfaces(v.Val())
	if len(array) == 0 {
		return nil
	}
	vars := make([]*Var, len(array))
	for k, v := range array {
		vars[k] = New(v)
	}
	return vars
}
