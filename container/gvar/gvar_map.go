// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gvar

import "github.com/gogf/gf/v2/util/gconv"

// MapOption 定义了映射转换的选项。 md5:8dc53d6fdc486bf8
type MapOption = gconv.MapOption

// Map 将 `v` 转换为 map[string]interface{} 并返回。 md5:88b2687bddef8ca6
// ff:取Map
// v:
// option:选项
func (v *Var) Map(option ...MapOption) map[string]interface{} {
	return gconv.Map(v.Val(), option...)
}

// MapStrAny 类似于 Map 函数，但它实现了 MapStrAny 接口。 md5:f7cf0af70c8cbee9
// yx:true
// ff:取MapStrAny
// v:
// option:
func (v *Var) MapStrAny(option ...MapOption) map[string]interface{} {
	return v.Map(option...)
}

// MapStrStr 将 `v` 转换并返回为 map[string]string 类型。 md5:b5af3d144d89aad9
// ff:取文本Map
// v:
// option:选项
func (v *Var) MapStrStr(option ...MapOption) map[string]string {
	return gconv.MapStrStr(v.Val(), option...)
}

// MapStrVar将`v`转换为map[string]Var并返回。 md5:18642fac7292b37a
// ff:取泛型类Map
// v:
// option:选项
func (v *Var) MapStrVar(option ...MapOption) map[string]*Var {
	m := v.Map(option...)
	if len(m) > 0 {
		vMap := make(map[string]*Var, len(m))
		for k, v := range m {
			vMap[k] = New(v)
		}
		return vMap
	}
	return nil
}

// MapDeep converts and returns `v` as map[string]interface{} recursively.
// ff:MapDeep弃用
// v:
// tags:值类型标签
func (v *Var) MapDeep(tags ...string) map[string]interface{} {
	return gconv.MapDeep(v.Val(), tags...)
}

// MapStrStrDeep converts and returns `v` as map[string]string recursively.
// ff:MapStrStrDeep弃用
// v:
// tags:值类型标签
func (v *Var) MapStrStrDeep(tags ...string) map[string]string {
	return gconv.MapStrStrDeep(v.Val(), tags...)
}

// MapStrVarDeep converts and returns `v` as map[string]*Var recursively.
// ff:取泛型类Map_递归
// v:
// tags:值类型标签
func (v *Var) MapStrVarDeep(tags ...string) map[string]*Var {
	m := v.MapDeep(tags...)
	if len(m) > 0 {
		vMap := make(map[string]*Var, len(m))
		for k, v := range m {
			vMap[k] = New(v)
		}
		return vMap
	}
	return nil
}

// Maps 将 `v` 转换为 map[string]string 类型并返回。参考 gconv.Maps。
// md5:0240a1ad5bd80743
// ff:取Map切片
// v:
// option:选项
func (v *Var) Maps(option ...MapOption) []map[string]interface{} {
	return gconv.Maps(v.Val(), option...)
}

// MapsDeep converts `value` to []map[string]interface{} recursively.
// ff:MapsDeep弃用
// v:
// tags:值类型标签
func (v *Var) MapsDeep(tags ...string) []map[string]interface{} {
	return gconv.MapsDeep(v.Val(), tags...)
}

// MapToMap 将任意映射类型变量 `params` 转换为另一个映射类型变量 `pointer`。
// 参见 gconv.MapToMap。
// md5:bc3ef5f9ee0920e7
// ff:
// v:
// pointer:
// mapping:
// err:
func (v *Var) MapToMap(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMap(v.Val(), pointer, mapping...)
}

// MapToMaps 将任何类型的映射变量 `params` 转换为另一个映射类型变量 `pointer`。
// 参考 gconv.MapToMaps。
// md5:1bd0da08c8937a10
// ff:
// v:
// pointer:
// mapping:
// err:
func (v *Var) MapToMaps(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMaps(v.Val(), pointer, mapping...)
}

// MapToMapsDeep 将任何映射类型变量 `params` 递归地转换为另一个映射类型变量 `pointer`。
// 参考 gconv.MapToMapsDeep。
// md5:77546446f168a41a
// ff:
// v:
// pointer:
// mapping:
// err:
func (v *Var) MapToMapsDeep(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMaps(v.Val(), pointer, mapping...)
}
