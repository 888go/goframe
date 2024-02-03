// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gvar

import (
	"github.com/888go/goframe/util/gconv"
)

// MapOption 定义了映射转换的选项。
type MapOption = gconv.MapOption

// Map 将 `v` 转换并以 map[string]interface{} 类型返回。
func (v *Var) Map(option ...MapOption) map[string]interface{} {
	return gconv.Map(v.Val(), option...)
}

// MapStrAny 类似于函数 Map，但实现了 MapStrAny 接口。
func (v *Var) MapStrAny(option ...MapOption) map[string]interface{} {
	return v.Map(option...)
}

// MapStrStr将`v`转换并返回为map[string]string类型。
func (v *Var) MapStrStr(option ...MapOption) map[string]string {
	return gconv.MapStrStr(v.Val(), option...)
}

// MapStrVar 将 `v` 转换并返回为 map[string]Var 类型。
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

// MapDeep 递归地将 `v` 转换并返回为 map[string]interface{} 类型的值。
// 注意：已弃用，请改用 Map。
func (v *Var) MapDeep(tags ...string) map[string]interface{} {
	return gconv.MapDeep(v.Val(), tags...)
}

// MapStrStrDeep 递归地将 `v` 转换并返回为 map[string]string 类型。
// 已弃用：请改用 MapStrStr。
func (v *Var) MapStrStrDeep(tags ...string) map[string]string {
	return gconv.MapStrStrDeep(v.Val(), tags...)
}

// MapStrVarDeep 递归地将 `v` 转换并返回为 map[string]*Var 类型的值。
// 已弃用：请改用 MapStrVar。
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

// Maps 将 `v` 转换并以 map[string]string 类型返回。
// 参见 gconv.Maps。
func (v *Var) Maps(option ...MapOption) []map[string]interface{} {
	return gconv.Maps(v.Val(), option...)
}

// MapsDeep 递归地将 `value` 转换为 []map[string]interface{} 类型的切片。
// 注：已弃用，请改用 Maps。
func (v *Var) MapsDeep(tags ...string) []map[string]interface{} {
	return gconv.MapsDeep(v.Val(), tags...)
}

// MapToMap 将任意类型map变量 `params` 转换为另一种map类型变量 `pointer`。
// 参见 gconv.MapToMap。
func (v *Var) MapToMap(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMap(v.Val(), pointer, mapping...)
}

// MapToMaps 将任意类型map变量 `params` 转换为另一种map类型变量 `pointer`。
// 参见 gconv.MapToMaps。
func (v *Var) MapToMaps(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMaps(v.Val(), pointer, mapping...)
}

// MapToMapsDeep 将任意类型的 map 变量 `params` 递归地转换为另一种 map 类型变量
// `pointer`。
// 参见 gconv.MapToMapsDeep。
func (v *Var) MapToMapsDeep(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMaps(v.Val(), pointer, mapping...)
}
