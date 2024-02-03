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

// Structs 将任何切片转换为给定的结构体切片。
// 也可以参考 Scan, Struct。
func Structs(params interface{}, pointer interface{}, paramKeyToAttrMap ...map[string]string) (err error) {
	return Scan(params, pointer, paramKeyToAttrMap...)
}

// SliceStruct 是 Structs 的别名。
func SliceStruct(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	return Structs(params, pointer, mapping...)
}

// StructsTag的行为类似于Structs，但增加了对优先级标签功能的支持，该功能用于获取指定的标签，以便将`params`键值对映射到结构体属性名称。
// 参数`priorityTag`支持多个标签，这些标签可以使用字符','连接。
func StructsTag(params interface{}, pointer interface{}, priorityTag string) (err error) {
	return doStructs(params, pointer, nil, priorityTag)
}

// doStructs 将任何切片转换为给定的结构体切片。
//
// 如果`params`是字符串或[]byte，它会自动检查并转换为json字符串到[]map。
//
// 参数`pointer`应为指向结构体切片的指针类型。
// 注意，如果`pointer`是指向另一个结构体切片类型的指针的指针，
// 它将内部创建结构体/指针。
// 以下是更详细的中文注释：
// ```go
// doStructs 函数用于将任意类型的切片转换为目标结构体切片。
//
// 当传入参数 `params` 为字符串或字节切片（[]byte）时，函数会自动检测并将其转换成 JSON 字符串，进一步解析为 []map 类型的数据。
//
// 参数 `pointer` 需要是指向结构体切片的指针类型。特别地，如果 `pointer` 是一个指向结构体切片指针的指针，该函数将在内部自行创建所需结构体及相应的指针。
func doStructs(
	params interface{}, pointer interface{}, paramKeyToAttrMap map[string]string, priorityTag string,
) (err error) {
	if params == nil {
		// 如果`params`为nil，则不进行转换。
		return nil
	}
	if pointer == nil {
		return gerror.NewCode(gcode.CodeInvalidParameter, "object pointer cannot be nil")
	}

	if doStructsByDirectReflectSet(params, pointer) {
		return nil
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
	// 指针类型检查。
	pointerRv, ok := pointer.(reflect.Value)
	if !ok {
		pointerRv = reflect.ValueOf(pointer)
		if kind := pointerRv.Kind(); kind != reflect.Ptr {
			return gerror.NewCodef(gcode.CodeInvalidParameter, "pointer should be type of pointer, but got: %v", kind)
		}
	}
	// 将`params`转换为映射切片。
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
	// 如果`params`是一个空切片，则不进行转换。
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

// doStructsByDirectReflectSet 直接使用 reflect.Set 进行转换操作。
// 如果转换成功，返回 true，否则返回 false。
func doStructsByDirectReflectSet(params interface{}, pointer interface{}) (ok bool) {
	v1 := reflect.ValueOf(pointer)
	v2 := reflect.ValueOf(params)
	if v1.Kind() == reflect.Ptr {
		if elem := v1.Elem(); elem.IsValid() && elem.Type() == v2.Type() {
			elem.Set(v2)
			ok = true
		}
	}
	return ok
}
