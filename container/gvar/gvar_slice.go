// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 泛型类

import (
	gconv "github.com/888go/goframe/util/gconv"
)

// X取整数切片 将 `v` 转换并返回为 []int 类型。 md5:063ba1cd966af441
func (v *Var) X取整数切片() []int {
	return gconv.X取整数切片(v.X取值())
}

// X取整数64位切片 将 `v` 转换并返回为 []int64。 md5:def1c601eff8c141
func (v *Var) X取整数64位切片() []int64 {
	return gconv.X取整数64位切片(v.X取值())
}

// X取正整数切片将`v`转换为[]uint并返回。 md5:fee8ec1b779988fe
func (v *Var) X取正整数切片() []uint {
	return gconv.X取正整数切片(v.X取值())
}

// X取正整数64位切片 将 `v` 转换为 []uint64 并返回。 md5:b4b77de84a5454a9
func (v *Var) X取正整数64位切片() []uint64 {
	return gconv.X取正整数64位切片(v.X取值())
}

// X取小数切片 是 Float64s 的别名。 md5:df981a5dcf3a4356
func (v *Var) X取小数切片() []float64 {
	return gconv.X取小数切片(v.X取值())
}

// X取小数32位切片 将 `v` 转换并返回为 []float32。 md5:c1c14441284dfcfd
func (v *Var) X取小数32位切片() []float32 {
	return gconv.X取小数32位切片(v.X取值())
}

// X取小数64位切片 将 `v` 转换为 []float64 并返回。 md5:f515ce0f1efc3fca
func (v *Var) X取小数64位切片() []float64 {
	return gconv.X取小数64位切片(v.X取值())
}

// X取文本切片 将 `v` 转换为字符串切片并返回。 md5:7f2a45762d2cb9c8
func (v *Var) X取文本切片() []string {
	return gconv.X取文本切片(v.X取值())
}

// X取any切片 将 `v` 转换并返回为 []interface{}。 md5:8b6ea8b55919ad7a
func (v *Var) X取any切片() []interface{} {
	return gconv.X取any切片(v.X取值())
}

// Slice别名 是 Interfaces 的别名。 md5:94e5a660c49c5e34
func (v *Var) Slice别名() []interface{} {
	return v.X取any切片()
}

// Array别名是Interfaces的别名。 md5:137944f64a8b7b89
func (v *Var) Array别名() []interface{} {
	return v.X取any切片()
}

// X取泛型类切片 将 `v` 转换为 []Var 并返回。 md5:fdbc7c7817f9852d
func (v *Var) X取泛型类切片() []*Var {
	array := gconv.X取any切片(v.X取值())
	if len(array) == 0 {
		return nil
	}
	vars := make([]*Var, len(array))
	for k, v := range array {
		vars[k] = X创建(v)
	}
	return vars
}
