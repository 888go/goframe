// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gconv

import (
	"reflect"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
)

type (
	converterInType  = reflect.Type
	converterOutType = reflect.Type
	converterFunc    = reflect.Value
)

// customConverters 用于内部转换器存储。
var customConverters = make(map[converterInType]map[converterOutType]converterFunc)

// RegisterConverter 用于注册自定义转换器。
// 在使用此自定义转换功能之前，必须先进行注册。
// 建议在初始化阶段完成此操作。
// 注意：
// 1. 参数 `fn` 必须定义为模式 `func(T1) (T2, error)`。
//    它将把类型 `T1` 转换为类型 `T2`。
// 2. `T1` 不应为指针类型，但 `T2` 应为指针类型。
func RegisterConverter(fn interface{}) (err error) {
	var (
		fnReflectType = reflect.TypeOf(fn)
		errType       = reflect.TypeOf((*error)(nil)).Elem()
	)
	if fnReflectType.Kind() != reflect.Func ||
		fnReflectType.NumIn() != 1 || fnReflectType.NumOut() != 2 ||
		!fnReflectType.Out(1).Implements(errType) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			"parameter must be type of function and defined as pattern `func(T1) (T2, error)`, but defined as `%s`",
			fnReflectType.String(),
		)
		return
	}

	// 转换器映射中的键和值不应为指针类型。
	var (
		inType  = fnReflectType.In(0)
		outType = fnReflectType.Out(0)
	)
	if inType.Kind() == reflect.Pointer {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			"invalid input parameter type `%s`: should not be type of pointer",
			inType.String(),
		)
		return
	}
	if outType.Kind() != reflect.Pointer {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			"invalid output parameter type `%s`: should be type of pointer",
			outType.String(),
		)
		return
	}

	registeredOutTypeMap, ok := customConverters[inType]
	if !ok {
		registeredOutTypeMap = make(map[converterOutType]converterFunc)
		customConverters[inType] = registeredOutTypeMap
	}
	if _, ok = registeredOutTypeMap[outType]; ok {
		err = gerror.NewCodef(
			gcode.CodeInvalidOperation,
			"the converter parameter type `%s` to type `%s` has already been registered",
			inType.String(), outType.String(),
		)
		return
	}
	registeredOutTypeMap[outType] = reflect.ValueOf(fn)
	return
}

// callCustomConverter 调用自定义转换器。它会尝试一些可能的类型。
func callCustomConverter(srcReflectValue reflect.Value, dstReflectValue reflect.Value) (converted bool, err error) {
	if len(customConverters) == 0 {
		return false, nil
	}
	var (
		ok      bool
		srcType = srcReflectValue.Type()
	)
	for srcType.Kind() == reflect.Pointer {
		srcType = srcType.Elem()
	}
	var (
		registeredOutTypeMap    map[converterOutType]converterFunc
		registeredConverterFunc converterFunc
	)
	// 首先，通过输入参数类型搜索映射。
	registeredOutTypeMap, ok = customConverters[srcType]
	if !ok {
		return false, nil
	}
	var dstType = dstReflectValue.Type()
	if dstType.Kind() == reflect.Pointer && dstReflectValue.Elem().Kind() == reflect.Pointer {
		dstType = dstReflectValue.Elem().Type()
	} else if dstType.Kind() != reflect.Pointer && dstReflectValue.CanAddr() {
		dstType = dstReflectValue.Addr().Type()
	}
// 其次，它在输入参数类型映射中搜索
// 并通过输出参数类型找到结果转换函数。
	registeredConverterFunc, ok = registeredOutTypeMap[dstType]
	if !ok {
		return false, nil
	}
	// 转换器函数调用。
	for srcReflectValue.Type() != srcType {
		srcReflectValue = srcReflectValue.Elem()
	}
	result := registeredConverterFunc.Call([]reflect.Value{srcReflectValue})
	if !result[1].IsNil() {
		return false, result[1].Interface().(error)
	}
	// `result[0]` 是一个指针。
	if result[0].IsNil() {
		return false, nil
	}
	var resultValue = result[0]
	for {
		if resultValue.Type() == dstReflectValue.Type() && dstReflectValue.CanSet() {
			dstReflectValue.Set(resultValue)
			converted = true
		} else if dstReflectValue.Kind() == reflect.Pointer {
			if resultValue.Type() == dstReflectValue.Elem().Type() && dstReflectValue.Elem().CanSet() {
				dstReflectValue.Elem().Set(resultValue)
				converted = true
			}
		}
		if converted {
			break
		}
		if resultValue.Kind() == reflect.Pointer {
			resultValue = resultValue.Elem()
		} else {
			break
		}
	}

	return converted, nil
}
