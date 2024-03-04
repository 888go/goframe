// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gconv

import (
	"reflect"
	
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/888go/goframe/gconv/internal/json"
)

// MapToMap 使用 reflect 将任意类型的 map 变量 `params` 转换为另一种 map 类型变量 `pointer`
// 详细实现请参考 doMapToMap 函数。
func MapToMap(params interface{}, pointer interface{}, mapping ...map[string]string) error {
	return doMapToMap(params, pointer, mapping...)
}

// doMapToMap 将任意类型的 map 变量 `params` 转换为另一种 map 类型变量 `pointer`。
//
// 参数 `params` 可以是任意类型的 map，例如：
// map[string]string, map[string]struct, map[string]*struct, reflect.Value 等等。
//
// 参数 `pointer` 应该是指向 map 的类型，例如：
// *map[int]string, *map[string]struct, *map[string]*struct, *reflect.Value 等等。
//
// 可选参数 `mapping` 用于 struct 属性到 map 键的映射，只有当原 map `params` 中的项为 struct 类型时才有意义。
func doMapToMap(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
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
	var (
		paramsRv                  reflect.Value
		paramsKind                reflect.Kind
		keyToAttributeNameMapping map[string]string
	)
	if len(mapping) > 0 {
		keyToAttributeNameMapping = mapping[0]
	}
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
	if paramsKind != reflect.Map {
		return doMapToMap(Map(params), pointer, mapping...)
	}
	// 空参数映射，无需继续。
	if paramsRv.Len() == 0 {
		return nil
	}
	var pointerRv reflect.Value
	if v, ok := pointer.(reflect.Value); ok {
		pointerRv = v
	} else {
		pointerRv = reflect.ValueOf(pointer)
	}
	pointerKind := pointerRv.Kind()
	for pointerKind == reflect.Ptr {
		pointerRv = pointerRv.Elem()
		pointerKind = pointerRv.Kind()
	}
	if pointerKind != reflect.Map {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "pointer should be type of *map, but got:%s", pointerKind)
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
		paramsKeys       = paramsRv.MapKeys()
		pointerKeyType   = pointerRv.Type().Key()
		pointerValueType = pointerRv.Type().Elem()
		pointerValueKind = pointerValueType.Kind()
		dataMap          = reflect.MakeMapWithSize(pointerRv.Type(), len(paramsKeys))
	)
	// 获取目标映射的真正元素类型。
	if pointerValueKind == reflect.Ptr {
		pointerValueKind = pointerValueType.Elem().Kind()
	}
	for _, key := range paramsKeys {
		mapValue := reflect.New(pointerValueType).Elem()
		switch pointerValueKind {
		case reflect.Map, reflect.Struct:
			if err = doStruct(paramsRv.MapIndex(key).Interface(), mapValue, keyToAttributeNameMapping, ""); err != nil {
				return err
			}
		default:
			mapValue.Set(
				reflect.ValueOf(
					doConvert(doConvertInput{
						FromValue:  paramsRv.MapIndex(key).Interface(),
						ToTypeName: pointerValueType.String(),
						ReferValue: mapValue,
						Extra:      nil,
					}),
				),
			)
		}
		var mapKey = reflect.ValueOf(
			doConvert(doConvertInput{
				FromValue:  key.Interface(),
				ToTypeName: pointerKeyType.Name(),
				ReferValue: reflect.New(pointerKeyType).Elem().Interface(),
				Extra:      nil,
			}),
		)
		dataMap.SetMapIndex(mapKey, mapValue)
	}
	pointerRv.Set(dataMap)
	return nil
}
