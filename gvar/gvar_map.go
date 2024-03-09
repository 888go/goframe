// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 泛型类

import (
	"github.com/gogf/gf/v2/util/gconv"
)

// MapOption 定义了映射转换的选项。
type MapOption = gconv.MapOption

// Map 将 `v` 转换并以 map[string]interface{} 类型返回。
func (v *Var) X取Map(选项 ...MapOption) map[string]interface{} {
	return gconv.Map(v.X取值(), 选项...)
}

// MapStrAny 类似于函数 Map，但实现了 MapStrAny 接口。
func (v *Var) X取MapStrAny(选项 ...MapOption) map[string]interface{} {
	return v.X取Map(选项...)
}

// MapStrStr将`v`转换并返回为map[string]string类型。
func (v *Var) X取文本Map(选项 ...MapOption) map[string]string {
	return gconv.MapStrStr(v.X取值(), 选项...)
}

// MapStrVar 将 `v` 转换并返回为 map[string]Var 类型。
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

// MapDeep 递归地将 `v` 转换并返回为 map[string]interface{} 类型的值。
// 注意：已弃用，请改用 Map。
func (v *Var) MapDeep弃用(值类型标签 ...string) map[string]interface{} {
	return gconv.MapDeep(v.X取值(), 值类型标签...)
}

// MapStrStrDeep 递归地将 `v` 转换并返回为 map[string]string 类型。
// 已弃用：请改用 MapStrStr。
func (v *Var) MapStrStrDeep弃用(值类型标签 ...string) map[string]string {
	return gconv.MapStrStrDeep(v.X取值(), 值类型标签...)
}

// MapStrVarDeep 递归地将 `v` 转换并返回为 map[string]*Var 类型的值。
// 已弃用：请改用 MapStrVar。
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

// Maps 将 `v` 转换并以 map[string]string 类型返回。
// 参见 gconv.Maps。
func (v *Var) X取Map数组(选项 ...MapOption) []map[string]interface{} {
	return gconv.Maps(v.X取值(), 选项...)
}

// MapsDeep 递归地将 `value` 转换为 []map[string]interface{} 类型的切片。
// 注：已弃用，请改用 Maps。
func (v *Var) MapsDeep弃用(值类型标签 ...string) []map[string]interface{} {
	return gconv.MapsDeep(v.X取值(), 值类型标签...)
}

// MapToMap 将任意类型map变量 `params` 转换为另一种map类型变量 `pointer`。
// 参见 gconv.MapToMap。
func (v *Var) MapToMap(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMap(v.X取值(), pointer, mapping...)
}

// MapToMaps 将任意类型map变量 `params` 转换为另一种map类型变量 `pointer`。
// 参见 gconv.MapToMaps。
func (v *Var) MapToMaps(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMaps(v.X取值(), pointer, mapping...)
}

// MapToMapsDeep 将任意类型的 map 变量 `params` 递归地转换为另一种 map 类型变量
// `pointer`。
// 参见 gconv.MapToMapsDeep。
func (v *Var) MapToMapsDeep(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMaps(v.X取值(), pointer, mapping...)
}
