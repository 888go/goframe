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

// MapOption 定义了映射转换的选项。 md5:8dc53d6fdc486bf8
type MapOption = gconv.MapOption

// X取Map 将 `v` 转换为 map[string]interface{} 并返回。 md5:88b2687bddef8ca6
func (v *Var) X取Map(选项 ...MapOption) map[string]interface{} {
	return gconv.X取Map(v.X取值(), 选项...)
}

// X取MapStrAny 类似于 Map 函数，但它实现了 X取MapStrAny 接口。 md5:f7cf0af70c8cbee9
func (v *Var) X取MapStrAny(option ...MapOption) map[string]interface{} {
	return v.X取Map(option...)
}

// X取文本Map 将 `v` 转换并返回为 map[string]string 类型。 md5:b5af3d144d89aad9
func (v *Var) X取文本Map(选项 ...MapOption) map[string]string {
	return gconv.X取文本Map(v.X取值(), 选项...)
}

// X取泛型类Map将`v`转换为map[string]Var并返回。 md5:18642fac7292b37a
func (v *Var) X取泛型类Map(选项 ...MapOption) map[string]*Var {
	m := v.X取Map(选项...)
	if len(m) > 0 {
		vMap := make(map[string]*Var, len(m))
		for k, v := range m {
			vMap[k] = X创建(v)
		}
		return vMap
	}
	return nil
}

// MapDeep弃用 递归地将 `v` 转换为 map[string]interface{} 并返回。
// 提示：已使用 Map 函数代替。
// md5:1a45b51f1c84bb44
func (v *Var) MapDeep弃用(值类型标签 ...string) map[string]interface{} {
	return gconv.X取Map_递归(v.X取值(), 值类型标签...)
}

// MapStrStrDeep弃用 递归地将 `v` 转换并返回为 map[string]string 类型。
// 已废弃：请使用 MapStrStr 函数代替。
// md5:9f5885f5e2b8a6e4
func (v *Var) MapStrStrDeep弃用(值类型标签 ...string) map[string]string {
	return gconv.X取文本Map_递归(v.X取值(), 值类型标签...)
}

// X取泛型类Map_递归 递归地将`v`转换并返回为map[string]*Var。
// 注意：请使用MapStrVar代替。
// md5:b37116aff42f6b15
func (v *Var) X取泛型类Map_递归(值类型标签 ...string) map[string]*Var {
	m := v.MapDeep弃用(值类型标签...)
	if len(m) > 0 {
		vMap := make(map[string]*Var, len(m))
		for k, v := range m {
			vMap[k] = X创建(v)
		}
		return vMap
	}
	return nil
}

// X取Map切片 将 `v` 转换为 map[string]string 类型并返回。参考 gconv.X取Map切片。
// md5:0240a1ad5bd80743
func (v *Var) X取Map切片(选项 ...MapOption) []map[string]interface{} {
	return gconv.X取Map切片(v.X取值(), 选项...)
}

// MapsDeep弃用 将 `value` 递归地转换为 []map[string]interface{} 类型。
// 警告：已使用 Maps 替代。
// md5:10c733fd844f279f
func (v *Var) MapsDeep弃用(值类型标签 ...string) []map[string]interface{} {
	return gconv.X取Map切片_递归(v.X取值(), 值类型标签...)
}

// MapToMap 将任意映射类型变量 `params` 转换为另一个映射类型变量 `pointer`。
// 参见 gconv.MapToMap。
// md5:bc3ef5f9ee0920e7
func (v *Var) MapToMap(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMap(v.X取值(), pointer, mapping...)
}

// MapToMaps 将任何类型的映射变量 `params` 转换为另一个映射类型变量 `pointer`。
// 参考 gconv.MapToMaps。
// md5:1bd0da08c8937a10
func (v *Var) MapToMaps(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMaps(v.X取值(), pointer, mapping...)
}

// MapToMapsDeep 将任何映射类型变量 `params` 递归地转换为另一个映射类型变量 `pointer`。
// 参考 gconv.MapToMapsDeep。
// md5:77546446f168a41a
func (v *Var) MapToMapsDeep(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMaps(v.X取值(), pointer, mapping...)
}
