// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gvar

import "github.com/gogf/gf/v2/util/gconv"

// MapOption specifies the option for map converting.
type MapOption = gconv.MapOption

// Map converts and returns `v` as map[string]interface{}.
// ff:取Map
// v:
// option:选项
func (v *Var) Map(option ...MapOption) map[string]interface{} {
	return gconv.Map(v.Val(), option...)
}

// MapStrAny is like function Map, but implements the interface of MapStrAny.
// yx:true
// ff:取MapStrAny
// v:
// option:
func (v *Var) MapStrAny(option ...MapOption) map[string]interface{} {
	return v.Map(option...)
}

// MapStrStr converts and returns `v` as map[string]string.
// ff:取文本Map
// v:
// option:选项
func (v *Var) MapStrStr(option ...MapOption) map[string]string {
	return gconv.MapStrStr(v.Val(), option...)
}

// MapStrVar converts and returns `v` as map[string]Var.
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

// Maps converts and returns `v` as map[string]string.
// See gconv.Maps.
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

// MapToMap converts any map type variable `params` to another map type variable `pointer`.
// See gconv.MapToMap.
// ff:
// v:
// pointer:
// mapping:
// err:
func (v *Var) MapToMap(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMap(v.Val(), pointer, mapping...)
}

// MapToMaps converts any map type variable `params` to another map type variable `pointer`.
// See gconv.MapToMaps.
// ff:
// v:
// pointer:
// mapping:
// err:
func (v *Var) MapToMaps(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMaps(v.Val(), pointer, mapping...)
}

// MapToMapsDeep converts any map type variable `params` to another map type variable
// `pointer` recursively.
// See gconv.MapToMapsDeep.
// ff:
// v:
// pointer:
// mapping:
// err:
func (v *Var) MapToMapsDeep(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMaps(v.Val(), pointer, mapping...)
}
