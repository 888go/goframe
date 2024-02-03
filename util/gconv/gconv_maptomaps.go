// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gconv

import (
	"reflect"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
)

// MapToMaps 将任意切片类型变量 `params` 转换为另一种映射切片类型变量 `pointer`。
// 请参考 doMapToMaps 函数。
func MapToMaps(params interface{}, pointer interface{}, mapping ...map[string]string) error {
	return doMapToMaps(params, pointer, mapping...)
}

// doMapToMaps 将任意类型的 map 变量 `params` 转换为另一个 map 切片变量 `pointer`。
//
// 参数 `params` 可以为 []map、[]*map、[]struct 或 []*struct 类型。
//
// 参数 `pointer` 应为 []map 或 []*map 类型。
//
// 可选参数 `mapping` 用于 struct 属性到 map 键的映射，仅当 `params` 的元素类型为 struct 时才有意义。
// 这段代码注释翻译成中文后如下：
// ```go
// doMapToMaps 函数将任何类型的 map 变量 `params` 转换为另一种 map 切片变量 `pointer`。
//
// 参数 `params` 支持以下类型：[]map、[]*map、[]struct、([]*struct。
//
// 参数 `pointer` 必须是 []map 或 []*map 类型。
//
// 可选参数 `mapping` 用于进行 struct 属性到 map 键的映射，这个参数只有在 `params` 中的项为 struct 类型时才起作用。
func doMapToMaps(params interface{}, pointer interface{}, paramKeyToAttrMap ...map[string]string) (err error) {
	// 如果给定的`params`是JSON格式，那么它将使用json.Unmarshal进行转换。
	switch r := params.(type) {
	case []byte:
		if json.Valid(r) {
			if rv, ok := pointer.(reflect.Value); ok {
				if rv.Kind() == reflect.Ptr {
					return json.UnmarshalUseNumber(r, rv.Interface())
				}
			} else {
				return json.UnmarshalUseNumber(r, pointer)
			}
		}
	case string:
		if paramsBytes := []byte(r); json.Valid(paramsBytes) {
			if rv, ok := pointer.(reflect.Value); ok {
				if rv.Kind() == reflect.Ptr {
					return json.UnmarshalUseNumber(paramsBytes, rv.Interface())
				}
			} else {
				return json.UnmarshalUseNumber(paramsBytes, pointer)
			}
		}
	}
	// 参数及其元素类型检查
	var (
		paramsRv   reflect.Value
		paramsKind reflect.Kind
	)
	if v, ok := params.(reflect.Value); ok {
		paramsRv = v
	} else {
		paramsRv = reflect.ValueOf(params)
	}
	paramsKind = paramsRv.Kind()
	if paramsKind == reflect.Ptr {
		paramsRv = paramsRv.Elem()
		paramsKind = paramsRv.Kind()
	}
	if paramsKind != reflect.Array && paramsKind != reflect.Slice {
		return gerror.NewCode(gcode.CodeInvalidParameter, "params should be type of slice, eg: []map/[]*map/[]struct/[]*struct")
	}
	var (
		paramsElem     = paramsRv.Type().Elem()
		paramsElemKind = paramsElem.Kind()
	)
	if paramsElemKind == reflect.Ptr {
		paramsElem = paramsElem.Elem()
		paramsElemKind = paramsElem.Kind()
	}
	if paramsElemKind != reflect.Map && paramsElemKind != reflect.Struct && paramsElemKind != reflect.Interface {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "params element should be type of map/*map/struct/*struct, but got: %s", paramsElemKind)
	}
	// 空切片，无需继续。
	if paramsRv.Len() == 0 {
		return nil
	}
	// 指针及其元素类型检查。
	var (
		pointerRv   = reflect.ValueOf(pointer)
		pointerKind = pointerRv.Kind()
	)
	for pointerKind == reflect.Ptr {
		pointerRv = pointerRv.Elem()
		pointerKind = pointerRv.Kind()
	}
	if pointerKind != reflect.Array && pointerKind != reflect.Slice {
		return gerror.NewCode(gcode.CodeInvalidParameter, "pointer should be type of *[]map/*[]*map")
	}
	var (
		pointerElemType = pointerRv.Type().Elem()
		pointerElemKind = pointerElemType.Kind()
	)
	if pointerElemKind == reflect.Ptr {
		pointerElemKind = pointerElemType.Elem().Kind()
	}
	if pointerElemKind != reflect.Map {
		return gerror.NewCode(gcode.CodeInvalidParameter, "pointer element should be type of map/*map")
	}
	defer func() {
		// 捕获 panic，特别是反射操作引发的 panic。
		if exception := recover(); exception != nil {
			if v, ok := exception.(error); ok && gerror.HasStack(v) {
				err = v
			} else {
				err = gerror.NewCodeSkipf(gcode.CodeInternalPanic, 1, "%+v", exception)
			}
		}
	}()
	var (
		pointerSlice = reflect.MakeSlice(pointerRv.Type(), paramsRv.Len(), paramsRv.Len())
	)
	for i := 0; i < paramsRv.Len(); i++ {
		var item reflect.Value
		if pointerElemType.Kind() == reflect.Ptr {
			item = reflect.New(pointerElemType.Elem())
			if err = MapToMap(paramsRv.Index(i).Interface(), item, paramKeyToAttrMap...); err != nil {
				return err
			}
			pointerSlice.Index(i).Set(item)
		} else {
			item = reflect.New(pointerElemType)
			if err = MapToMap(paramsRv.Index(i).Interface(), item, paramKeyToAttrMap...); err != nil {
				return err
			}
			pointerSlice.Index(i).Set(item.Elem())
		}
	}
	pointerRv.Set(pointerSlice)
	return
}
