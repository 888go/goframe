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
func (vs Vars) X取文本数组() (数组 []string) {
	for _, v := range vs {
		数组 = append(数组, v.String())
	}
	return 数组
}

// Interfaces将`vs`转换并返回为[]interface{}类型。
func (vs Vars) X取any数组() (数组 []interface{}) {
	for _, v := range vs {
		数组 = append(数组, v.X取值())
	}
	return 数组
}

// Float32s将`vs`转换并返回为[]float32类型。
func (vs Vars) X取小数32位数组() (数组 []float32) {
	for _, v := range vs {
		数组 = append(数组, v.X取小数32位())
	}
	return 数组
}

// Float64s 将 `vs` 转换并返回为 []float64 类型的切片。
func (vs Vars) X取小数64位数组() (数组 []float64) {
	for _, v := range vs {
		数组 = append(数组, v.X取小数64位())
	}
	return 数组
}

// Ints将`vs`转换并返回为[]Int类型。
func (vs Vars) X取整数数组() (数组 []int) {
	for _, v := range vs {
		数组 = append(数组, v.X取整数())
	}
	return 数组
}

// Int8s将`vs`转换并返回为[]int8类型。
func (vs Vars) X取整数8位数组() (数组 []int8) {
	for _, v := range vs {
		数组 = append(数组, v.X取整数8位())
	}
	return 数组
}

// Int16s将`vs`转换并返回为[]int16类型。
func (vs Vars) X取整数16位数组() (数组 []int16) {
	for _, v := range vs {
		数组 = append(数组, v.X取整数16位())
	}
	return 数组
}

// Int32s 将 `vs` 转换并返回为 []int32 类型的切片。
func (vs Vars) X取整数32位数组() (数组 []int32) {
	for _, v := range vs {
		数组 = append(数组, v.X取整数32位())
	}
	return 数组
}

// Int64s将`vs`转换并返回为[]int64类型。
func (vs Vars) X取整数64位数组() (数组 []int64) {
	for _, v := range vs {
		数组 = append(数组, v.X取整数64位())
	}
	return 数组
}

// Uints 将 `vs` 转换并返回为 []uint 类型的切片。
func (vs Vars) X取正整数数组() (数组 []uint) {
	for _, v := range vs {
		数组 = append(数组, v.X取正整数())
	}
	return 数组
}

// Uint8s将`vs`转换并返回为[]uint8类型。
func (vs Vars) X取正整数8位数组() (数组 []uint8) {
	for _, v := range vs {
		数组 = append(数组, v.X取正整数8位())
	}
	return 数组
}

// Uint16s 将 `vs` 转换并返回为 []uint16 类型的切片。
func (vs Vars) X取正整数16位数组() (数组 []uint16) {
	for _, v := range vs {
		数组 = append(数组, v.X取正整数16位())
	}
	return 数组
}

// Uint32s 将 `vs` 转换并返回为 []uint32 类型的切片。
func (vs Vars) X取正整数32位数组() (数组 []uint32) {
	for _, v := range vs {
		数组 = append(数组, v.X取正整数32位())
	}
	return 数组
}

// Uint64s 将 `vs` 转换并返回为 []uint64 类型的切片。
func (vs Vars) X取正整数64位数组() (数组 []uint64) {
	for _, v := range vs {
		数组 = append(数组, v.X取正整数64位())
	}
	return 数组
}

// Scan 将 `vs` 转换为 []struct 或 []*struct 类型的切片。
func (vs Vars) X取结构体指针(结构体指针 interface{}, 名称映射 ...map[string]string) error {
	return gconv.Structs(vs.X取any数组(), 结构体指针, 名称映射...)
}
