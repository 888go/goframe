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

type (
	converterInType  = reflect.Type
	converterOutType = reflect.Type
	converterFunc    = reflect.Value
)

// customConverters 用于内部转换器的存储。 md5:eb816f1844daac79
var customConverters = make(map[converterInType]map[converterOutType]converterFunc)

// RegisterConverter 用于注册自定义转换器。
// 必须在使用此自定义转换功能之前进行注册。
// 建议在进程的启动程序中执行此操作。
//
// 注意：
//  1. 参数 `fn` 必须定义为模式 `func(T1) (T2, error)`。
//     它将类型 `T1` 转换为类型 `T2`。
//  2. `T1` 不应为指针类型，但 `T2` 应为指针类型。
// md5:8fbaa372837e6d8c
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
			"parameter must be type of converter function and defined as pattern `func(T1) (T2, error)`, but defined as `%s`",
			fnReflectType.String(),
		)
		return
	}

	// 转换映射的键和值不应该是指针。 md5:79bb068f1985b81a
	var (
		inType  = fnReflectType.In(0)
		outType = fnReflectType.Out(0)
	)
	if inType.Kind() == reflect.Pointer {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			"invalid converter function `%s`: invalid input parameter type `%s`, should not be type of pointer",
			fnReflectType.String(), inType.String(),
		)
		return
	}
	if outType.Kind() != reflect.Pointer {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			"invalid converter function `%s`: invalid output parameter type `%s` should be type of pointer",
			fnReflectType.String(), outType.String(),
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

func getRegisteredConverterFuncAndSrcType(
	srcReflectValue, dstReflectValueForRefer reflect.Value,
) (f converterFunc, srcType reflect.Type, ok bool) {
	if len(customConverters) == 0 {
		return reflect.Value{}, nil, false
	}
	srcType = srcReflectValue.Type()
	for srcType.Kind() == reflect.Pointer {
		srcType = srcType.Elem()
	}
	var registeredOutTypeMap map[converterOutType]converterFunc
	// 首先，它通过输入参数类型在映射中搜索。 md5:019f9d8418285668
	registeredOutTypeMap, ok = customConverters[srcType]
	if !ok {
		return reflect.Value{}, nil, false
	}
	var dstType = dstReflectValueForRefer.Type()
	if dstType.Kind() == reflect.Pointer {
		// 可能是**struct，这是设计上支持的。 md5:cb1f21754e39c3a1
		if dstType.Elem().Kind() == reflect.Pointer {
			dstType = dstType.Elem()
		}
	} else if dstReflectValueForRefer.IsValid() && dstReflectValueForRefer.CanAddr() {
		dstType = dstReflectValueForRefer.Addr().Type()
	} else {
		dstType = reflect.PointerTo(dstType)
	}
	// 其次，它会在输入参数类型映射中搜索
	// 并通过输出参数类型找到结果转换函数。
	// md5:3781290987232f09
	f, ok = registeredOutTypeMap[dstType]
	if !ok {
		return reflect.Value{}, nil, false
	}
	return
}

func callCustomConverterWithRefer(
	srcReflectValue, referReflectValue reflect.Value,
) (dstReflectValue reflect.Value, converted bool, err error) {
	registeredConverterFunc, srcType, ok := getRegisteredConverterFuncAndSrcType(srcReflectValue, referReflectValue)
	if !ok {
		return reflect.Value{}, false, nil
	}
	dstReflectValue = reflect.New(referReflectValue.Type()).Elem()
	converted, err = doCallCustomConverter(srcReflectValue, dstReflectValue, registeredConverterFunc, srcType)
	return
}

// callCustomConverter 调用自定义转换器。它会尝试一些可能的类型。 md5:44d83ddc5510baed
func callCustomConverter(srcReflectValue, dstReflectValue reflect.Value) (converted bool, err error) {
	registeredConverterFunc, srcType, ok := getRegisteredConverterFuncAndSrcType(srcReflectValue, dstReflectValue)
	if !ok {
		return false, nil
	}
	return doCallCustomConverter(srcReflectValue, dstReflectValue, registeredConverterFunc, srcType)
}

func doCallCustomConverter(
	srcReflectValue reflect.Value,
	dstReflectValue reflect.Value,
	registeredConverterFunc converterFunc,
	srcType reflect.Type,
) (converted bool, err error) {
	// 转换函数调用。 md5:1780fb4f627f751d
	for srcReflectValue.Type() != srcType {
		srcReflectValue = srcReflectValue.Elem()
	}
	result := registeredConverterFunc.Call([]reflect.Value{srcReflectValue})
	if !result[1].IsNil() {
		return false, result[1].Interface().(error)
	}
	// `result[0]`是一个指针。 md5:6505f86b6cd1e865
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
