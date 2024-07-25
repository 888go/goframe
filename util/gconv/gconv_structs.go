// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gconv

import (
	"reflect"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// Structs 将任何切片转换为给定结构体类型的切片。
// 另请参见 Scan, Struct。
// md5:9f4251433eeb586c
func Structs(params interface{}, pointer interface{}, paramKeyToAttrMap ...map[string]string) (err error) {
	return Scan(params, pointer, paramKeyToAttrMap...)
}

// SliceStruct 是 Structs 的别名。 md5:844cd0606fb4edf0
func SliceStruct(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	return Structs(params, pointer, mapping...)
}

// StructsTag 作为 Structs 的功能增强版本，还支持优先级标签特性。它根据 `params` 键值对获取指定的标签，并将其映射到结构体属性名称上。
// 参数 `priorityTag` 支持多个标签，这些标签可以使用逗号分隔。
// md5:ddc344beca5956a8
func StructsTag(params interface{}, pointer interface{}, priorityTag string) (err error) {
	return doStructs(params, pointer, nil, priorityTag)
}

// doStructs 将任何切片转换为给定的结构体切片。
//
// 如果 `params` 是字符串或[]byte，它会自动检查并将其转换为[]map。
//
// 参数 `pointer` 应该是指向结构体切片的指针。注意，如果 `pointer` 是指向结构体切片的另一个指针的指针，
// 它会在内部创建结构体/指针。
// md5:0bddadd6970c6b0b
func doStructs(
	params interface{}, pointer interface{}, paramKeyToAttrMap map[string]string, priorityTag string,
) (err error) {
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

	// Pointer type check.
	pointerRv, ok := pointer.(reflect.Value)
	if !ok {
		pointerRv = reflect.ValueOf(pointer)
		if kind := pointerRv.Kind(); kind != reflect.Ptr {
			return gerror.NewCodef(
				gcode.CodeInvalidParameter,
				"pointer should be type of pointer, but got: %v", kind,
			)
		}
	}
		// 将`params`转换为映射切片。 md5:d0685c4290b475fe
	var (
		paramsList []interface{}
		paramsRv   = reflect.ValueOf(params)
		paramsKind = paramsRv.Kind()
	)
	for paramsKind == reflect.Ptr {
		paramsRv = paramsRv.Elem()
		paramsKind = paramsRv.Kind()
	}
	switch paramsKind {
	case reflect.Slice, reflect.Array:
		paramsList = make([]interface{}, paramsRv.Len())
		for i := 0; i < paramsRv.Len(); i++ {
			paramsList[i] = paramsRv.Index(i).Interface()
		}
	default:
		var paramsMaps = Maps(params)
		paramsList = make([]interface{}, len(paramsMaps))
		for i := 0; i < len(paramsMaps); i++ {
			paramsList[i] = paramsMaps[i]
		}
	}
		// 如果`params`是一个空切片，则不进行转换。 md5:c2aa546ea7052f08
	if len(paramsList) == 0 {
		return nil
	}
	var (
		reflectElemArray = reflect.MakeSlice(pointerRv.Type().Elem(), len(paramsList), len(paramsList))
		itemType         = reflectElemArray.Index(0).Type()
		itemTypeKind     = itemType.Kind()
		pointerRvElem    = pointerRv.Elem()
		pointerRvLength  = pointerRvElem.Len()
	)
	if itemTypeKind == reflect.Ptr {
		// Pointer element.
		for i := 0; i < len(paramsList); i++ {
			var tempReflectValue reflect.Value
			if i < pointerRvLength {
				// Might be nil.
				tempReflectValue = pointerRvElem.Index(i).Elem()
			}
			if !tempReflectValue.IsValid() {
				tempReflectValue = reflect.New(itemType.Elem()).Elem()
			}
			if err = doStruct(paramsList[i], tempReflectValue, paramKeyToAttrMap, priorityTag); err != nil {
				return err
			}
			reflectElemArray.Index(i).Set(tempReflectValue.Addr())
		}
	} else {
		// Struct element.
		for i := 0; i < len(paramsList); i++ {
			var tempReflectValue reflect.Value
			if i < pointerRvLength {
				tempReflectValue = pointerRvElem.Index(i)
			} else {
				tempReflectValue = reflect.New(itemType).Elem()
			}
			if err = doStruct(paramsList[i], tempReflectValue, paramKeyToAttrMap, priorityTag); err != nil {
				return err
			}
			reflectElemArray.Index(i).Set(tempReflectValue)
		}
	}
	pointerRv.Elem().Set(reflectElemArray)
	return nil
}
