// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gvar

import (
	"github.com/gogf/gf/v2/util/gconv"
)

// Vars是一个*Var类型的切片。 md5:e8d9e194fa744a27
type Vars []*Var

// Strings 将 `vs` 转换为字符串切片并返回。 md5:5a2b561ba123eb40
// yx:true
// ff:取文本切片
// vs:
// s:
func (vs Vars) Strings() (s []string) {
	for _, v := range vs {
		s = append(s, v.String())
	}
	return s
}

// Interfaces 将 `vs` 转换并返回为 []interface{} 类型。 md5:6fe626226748e2c9
// yx:true
// ff:取any切片
// vs:
// s:
func (vs Vars) Interfaces() (s []interface{}) {
	for _, v := range vs {
		s = append(s, v.Val())
	}
	return s
}

// Float32s 将 `vs` 转换并返回为 []float32。 md5:12c5ec8d220918de
// ff:取小数32位切片
// vs:
// s:切片
func (vs Vars) Float32s() (s []float32) {
	for _, v := range vs {
		s = append(s, v.Float32())
	}
	return s
}

// Float64s 将 `vs` 转换为并返回一个 []float64 类型。 md5:5cab9209bbdd927c
// ff:取小数64位切片
// vs:
// s:切片
func (vs Vars) Float64s() (s []float64) {
	for _, v := range vs {
		s = append(s, v.Float64())
	}
	return s
}

// Ints 将 `vs` 转换为 []Int 并返回。 md5:cb5b6b2458c0609e
// yx:true
// ff:取整数切片
// vs:
// s:
func (vs Vars) Ints() (s []int) {
	for _, v := range vs {
		s = append(s, v.Int())
	}
	return s
}

// Int8s 将 `vs` 转换并返回为 []int8 类型的切片。 md5:16fc28e3f01c7aaf
// ff:取整数8位切片
// vs:
// s:切片
func (vs Vars) Int8s() (s []int8) {
	for _, v := range vs {
		s = append(s, v.Int8())
	}
	return s
}

// Int16s 将 `vs` 转换并返回为 []int16。 md5:cc768f8dc87ad4bc
// ff:取整数16位切片
// vs:
// s:切片
func (vs Vars) Int16s() (s []int16) {
	for _, v := range vs {
		s = append(s, v.Int16())
	}
	return s
}

// Int32s 将 `vs` 转换为 []int32 并返回。 md5:7215f5b253a087b8
// ff:取整数32位切片
// vs:
// s:切片
func (vs Vars) Int32s() (s []int32) {
	for _, v := range vs {
		s = append(s, v.Int32())
	}
	return s
}

// Int64s 将 `vs` 转换为 []int64 并返回。 md5:94ba8c59fe9fa849
// ff:取整数64位切片
// vs:
// s:切片
func (vs Vars) Int64s() (s []int64) {
	for _, v := range vs {
		s = append(s, v.Int64())
	}
	return s
}

// Uints 将 `vs` 转换并返回为 []uint 类型的切片。 md5:14417d6f5180237c
// yx:true
// ff:取正整数切片
// vs:
// s:
func (vs Vars) Uints() (s []uint) {
	for _, v := range vs {
		s = append(s, v.Uint())
	}
	return s
}

// Uint8s将`vs`转换并返回为[]uint8。 md5:7270fa354e186ac3
// ff:取正整数8位切片
// vs:
// s:切片
func (vs Vars) Uint8s() (s []uint8) {
	for _, v := range vs {
		s = append(s, v.Uint8())
	}
	return s
}

// Uint16s 将 `vs` 转换为 []uint16 并返回。 md5:14a5805ab12b7fbe
// ff:取正整数16位切片
// vs:
// s:切片
func (vs Vars) Uint16s() (s []uint16) {
	for _, v := range vs {
		s = append(s, v.Uint16())
	}
	return s
}

// Uint32s 将 `vs` 转换为 []uint32 并返回。 md5:432ace9bd7f4941e
// ff:取正整数32位切片
// vs:
// s:切片
func (vs Vars) Uint32s() (s []uint32) {
	for _, v := range vs {
		s = append(s, v.Uint32())
	}
	return s
}

// Uint64s 将 `vs` 转换并返回为 []uint64 类型的切片。 md5:4c19a53bd828ba63
// ff:取正整数64位切片
// vs:
// s:切片
func (vs Vars) Uint64s() (s []uint64) {
	for _, v := range vs {
		s = append(s, v.Uint64())
	}
	return s
}

// Scan 将 `vs` 转换为 []struct/[]*struct。 md5:e0e26cf0f4eb31e6
// ff:取结构体指针
// vs:
// pointer:结构体指针
// mapping:名称映射
func (vs Vars) Scan(pointer interface{}, mapping ...map[string]string) error {
	return gconv.Structs(vs.Interfaces(), pointer, mapping...)
}
