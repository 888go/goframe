// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gvar

import "github.com/gogf/gf/v2/util/gconv"

// Ints 将 `v` 转换并返回为 []int 类型。. md5:063ba1cd966af441
func (v *Var) Ints() []int {
	return gconv.Ints(v.Val())
}

// Int64s 将 `v` 转换并返回为 []int64。. md5:def1c601eff8c141
func (v *Var) Int64s() []int64 {
	return gconv.Int64s(v.Val())
}

// Uints将`v`转换为[]uint并返回。. md5:fee8ec1b779988fe
func (v *Var) Uints() []uint {
	return gconv.Uints(v.Val())
}

// Uint64s 将 `v` 转换为 []uint64 并返回。. md5:b4b77de84a5454a9
func (v *Var) Uint64s() []uint64 {
	return gconv.Uint64s(v.Val())
}

// Floats 是 Float64s 的别名。. md5:df981a5dcf3a4356
func (v *Var) Floats() []float64 {
	return gconv.Floats(v.Val())
}

// Float32s 将 `v` 转换并返回为 []float32。. md5:c1c14441284dfcfd
func (v *Var) Float32s() []float32 {
	return gconv.Float32s(v.Val())
}

// Float64s 将 `v` 转换为 []float64 并返回。. md5:f515ce0f1efc3fca
func (v *Var) Float64s() []float64 {
	return gconv.Float64s(v.Val())
}

// Strings 将 `v` 转换为字符串切片并返回。. md5:7f2a45762d2cb9c8
func (v *Var) Strings() []string {
	return gconv.Strings(v.Val())
}

// Interfaces 将 `v` 转换并返回为 []interface{}。. md5:8b6ea8b55919ad7a
func (v *Var) Interfaces() []interface{} {
	return gconv.Interfaces(v.Val())
}

// Slice 是 Interfaces 的别名。. md5:94e5a660c49c5e34
func (v *Var) Slice() []interface{} {
	return v.Interfaces()
}

// Array是Interfaces的别名。. md5:137944f64a8b7b89
func (v *Var) Array() []interface{} {
	return v.Interfaces()
}

// Vars 将 `v` 转换为 []Var 并返回。. md5:fdbc7c7817f9852d
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
