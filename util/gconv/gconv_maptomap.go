// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gconv

import (
	"reflect"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// MapToMap 通过反射将任何map类型变量`params`转换为另一个map类型变量`pointer`。
// 参考 doMapToMap。 md5:8fbdb048d4cad524
func MapToMap(params interface{}, pointer interface{}, mapping ...map[string]string) error {
	return Scan(params, pointer, mapping...)
}

// doMapToMap 将任何类型的映射变量 `params` 转换为另一个映射类型变量 `pointer`。
//
// 参数 `params` 可以是任何类型的映射，例如：map[string]string, map[string]struct, map[string]*struct, reflect.Value 等。
//
// 参数 `pointer` 应该是 *map 类型，例如：map[int]string, map[string]struct, map[string]*struct, reflect.Value 等。
//
// 可选参数 `mapping` 用于结构体属性到映射键的映射，只有当原始映射 `params` 的项是结构体类型时，这个参数才有意义。 md5:08b8fa82edaf8b08
func doMapToMap(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
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
	// 空参数映射，无需继续。 md5:7734e4bea4d21319
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
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`destination pointer should be type of *map, but got: %s`,
			pointerKind,
		)
	}
	defer func() {
		// 捕获panic，尤其是反射操作引发的panic。 md5:dd183bf8028f513a
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
	// 获取目标映射的真正元素类型。 md5:7e93cce5ee0c27e1
	if pointerValueKind == reflect.Ptr {
		pointerValueKind = pointerValueType.Elem().Kind()
	}
	for _, key := range paramsKeys {
		mapValue := reflect.New(pointerValueType).Elem()
		switch pointerValueKind {
		case reflect.Map, reflect.Struct:
			if err = doStruct(
				paramsRv.MapIndex(key).Interface(), mapValue, keyToAttributeNameMapping, "",
			); err != nil {
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
