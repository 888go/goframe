// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 泛型类

import (
	"github.com/888go/goframe/util/gconv"
)

// Ints 将 `v` 转换并返回为 []int 类型的切片。
func (v *Var) X取整数数组() []int {
	return 转换类.X取整数数组(v.X取值())
}

// Int64s 将 `v` 转换并返回为 []int64 类型的切片。
func (v *Var) X取整数64位数组() []int64 {
	return 转换类.X取整数64位数组(v.X取值())
}

// Uints 将 `v` 转换并返回为 []uint 类型。
func (v *Var) X取正整数数组() []uint {
	return 转换类.X取正整数数组(v.X取值())
}

// Uint64s将`v`转换并返回为[]uint64类型。
func (v *Var) X取正整数64位数组() []uint64 {
	return 转换类.X取正整数64位数组(v.X取值())
}

// Floats 是 Float64s 的别名。
func (v *Var) X取小数数组() []float64 {
	return 转换类.X取小数数组(v.X取值())
}

// Float32s 将 `v` 转换并返回为 []float32 类型的切片。
func (v *Var) X取小数32位数组() []float32 {
	return 转换类.X取小数32位数组(v.X取值())
}

// Float64s将`v`转换并返回为[]float64类型。
func (v *Var) X取小数64位数组() []float64 {
	return 转换类.X取小数64位数组(v.X取值())
}

// Strings 将 `v` 转换并返回为 []string 类型。
func (v *Var) X取文本数组() []string {
	return 转换类.X取文本数组(v.X取值())
}

// Interfaces 将 `v` 转换并返回为 []interfaces{} 类型的切片。
func (v *Var) X取any数组() []interface{} {
	return 转换类.X取any数组(v.X取值())
}

// Slice 是 Interfaces 的别名。
func (v *Var) Slice别名() []interface{} {
	return v.X取any数组()
}

// Array 是 Interfaces 的别名。
func (v *Var) Array别名() []interface{} {
	return v.X取any数组()
}

// Vars 将 `v` 转换并返回为 []Var 类型的切片。
func (v *Var) X取泛型类数组() []*Var {
	array := 转换类.X取any数组(v.X取值())
	if len(array) == 0 {
		return nil
	}
	vars := make([]*Var, len(array))
	for k, v := range array {
		vars[k] = X创建(v)
	}
	return vars
}
