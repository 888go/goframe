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

// Vars是一个*Var类型的切片。 md5:e8d9e194fa744a27
type Vars []*Var

// X取文本切片 将 `vs` 转换为字符串切片并返回。 md5:5a2b561ba123eb40
func (vs Vars) X取文本切片() (s []string) {
	for _, v := range vs {
		s = append(s, v.String())
	}
	return s
}

// X取any切片 将 `vs` 转换并返回为 []interface{} 类型。 md5:6fe626226748e2c9
func (vs Vars) X取any切片() (s []interface{}) {
	for _, v := range vs {
		s = append(s, v.X取值())
	}
	return s
}

// X取小数32位切片 将 `vs` 转换并返回为 []float32。 md5:12c5ec8d220918de
func (vs Vars) X取小数32位切片() (切片 []float32) {
	for _, v := range vs {
		切片 = append(切片, v.X取小数32位())
	}
	return 切片
}

// X取小数64位切片 将 `vs` 转换为并返回一个 []float64 类型。 md5:5cab9209bbdd927c
func (vs Vars) X取小数64位切片() (切片 []float64) {
	for _, v := range vs {
		切片 = append(切片, v.X取小数64位())
	}
	return 切片
}

// X取整数切片 将 `vs` 转换为 []Int 并返回。 md5:cb5b6b2458c0609e
func (vs Vars) X取整数切片() (s []int) {
	for _, v := range vs {
		s = append(s, v.X取整数())
	}
	return s
}

// X取整数8位切片 将 `vs` 转换并返回为 []int8 类型的切片。 md5:16fc28e3f01c7aaf
func (vs Vars) X取整数8位切片() (切片 []int8) {
	for _, v := range vs {
		切片 = append(切片, v.X取整数8位())
	}
	return 切片
}

// X取整数16位切片 将 `vs` 转换并返回为 []int16。 md5:cc768f8dc87ad4bc
func (vs Vars) X取整数16位切片() (切片 []int16) {
	for _, v := range vs {
		切片 = append(切片, v.X取整数16位())
	}
	return 切片
}

// X取整数32位切片 将 `vs` 转换为 []int32 并返回。 md5:7215f5b253a087b8
func (vs Vars) X取整数32位切片() (切片 []int32) {
	for _, v := range vs {
		切片 = append(切片, v.X取整数32位())
	}
	return 切片
}

// X取整数64位切片 将 `vs` 转换为 []int64 并返回。 md5:94ba8c59fe9fa849
func (vs Vars) X取整数64位切片() (切片 []int64) {
	for _, v := range vs {
		切片 = append(切片, v.X取整数64位())
	}
	return 切片
}

// X取正整数切片 将 `vs` 转换并返回为 []uint 类型的切片。 md5:14417d6f5180237c
func (vs Vars) X取正整数切片() (s []uint) {
	for _, v := range vs {
		s = append(s, v.X取正整数())
	}
	return s
}

// X取正整数8位切片将`vs`转换并返回为[]uint8。 md5:7270fa354e186ac3
func (vs Vars) X取正整数8位切片() (切片 []uint8) {
	for _, v := range vs {
		切片 = append(切片, v.X取正整数8位())
	}
	return 切片
}

// X取正整数16位切片 将 `vs` 转换为 []uint16 并返回。 md5:14a5805ab12b7fbe
func (vs Vars) X取正整数16位切片() (切片 []uint16) {
	for _, v := range vs {
		切片 = append(切片, v.X取正整数16位())
	}
	return 切片
}

// X取正整数32位切片 将 `vs` 转换为 []uint32 并返回。 md5:432ace9bd7f4941e
func (vs Vars) X取正整数32位切片() (切片 []uint32) {
	for _, v := range vs {
		切片 = append(切片, v.X取正整数32位())
	}
	return 切片
}

// X取正整数64位切片 将 `vs` 转换并返回为 []uint64 类型的切片。 md5:4c19a53bd828ba63
func (vs Vars) X取正整数64位切片() (切片 []uint64) {
	for _, v := range vs {
		切片 = append(切片, v.X取正整数64位())
	}
	return 切片
}

// X取结构体指针 将 `vs` 转换为 []struct/[]*struct。 md5:e0e26cf0f4eb31e6
func (vs Vars) X取结构体指针(结构体指针 interface{}, 名称映射 ...map[string]string) error {
	return gconv.Structs(vs.X取any切片(), 结构体指针, 名称映射...)
}
