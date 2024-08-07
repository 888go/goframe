// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类

import (
	"reflect"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
)

// MapToMaps 将任何切片类型变量 `params` 转换为另一个映射切片类型变量 `pointer`。
// 参见 doMapToMaps。
// md5:70b6d6cf0e63da31
func MapToMaps(params interface{}, pointer interface{}, mapping ...map[string]string) error {
	return Scan(params, pointer, mapping...)
}

// doMapToMaps 将任何类型的映射变量`params`转换为另一个映射切片变量`pointer`。
//
// 参数`params`可以是[]map, []*map, []struct, []*struct类型。
//
// 参数`pointer`应该是[]map, []*map类型。
//
// 可选参数`mapping`用于结构体属性到映射键的映射，只有当`params`的元素类型为struct时才有意义。
// md5:e5da204851e0f1b9
func doMapToMaps(params interface{}, pointer interface{}, paramKeyToAttrMap ...map[string]string) (err error) {
		// 检查参数及其元素类型。 md5:9678a18f11496e59
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
		return gerror.X创建错误码(
			gcode.CodeInvalidParameter,
			"params should be type of slice, example: []map/[]*map/[]struct/[]*struct",
		)
	}
	var (
		paramsElem     = paramsRv.Type().Elem()
		paramsElemKind = paramsElem.Kind()
	)
	if paramsElemKind == reflect.Ptr {
		paramsElem = paramsElem.Elem()
		paramsElemKind = paramsElem.Kind()
	}
	if paramsElemKind != reflect.Map &&
		paramsElemKind != reflect.Struct &&
		paramsElemKind != reflect.Interface {
		return gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			"params element should be type of map/*map/struct/*struct, but got: %s",
			paramsElemKind,
		)
	}
		// 空切片，无需继续。 md5:3e185b94ae24e0b3
	if paramsRv.Len() == 0 {
		return nil
	}
		// 指针及其元素类型的检查。 md5:b460debe108087f5
	var (
		pointerRv   = reflect.ValueOf(pointer)
		pointerKind = pointerRv.Kind()
	)
	for pointerKind == reflect.Ptr {
		pointerRv = pointerRv.Elem()
		pointerKind = pointerRv.Kind()
	}
	if pointerKind != reflect.Array && pointerKind != reflect.Slice {
		return gerror.X创建错误码(gcode.CodeInvalidParameter, "pointer should be type of *[]map/*[]*map")
	}
	var (
		pointerElemType = pointerRv.Type().Elem()
		pointerElemKind = pointerElemType.Kind()
	)
	if pointerElemKind == reflect.Ptr {
		pointerElemKind = pointerElemType.Elem().Kind()
	}
	if pointerElemKind != reflect.Map {
		return gerror.X创建错误码(gcode.CodeInvalidParameter, "pointer element should be type of map/*map")
	}
	defer func() {
				// 捕获panic，尤其是反射操作引发的panic。 md5:dd183bf8028f513a
		if exception := recover(); exception != nil {
			if v, ok := exception.(error); ok && gerror.X判断是否带堆栈(v) {
				err = v
			} else {
				err = gerror.X创建错误码并跳过堆栈与格式化(gcode.CodeInternalPanic, 1, "%+v", exception)
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
