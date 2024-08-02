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
	"github.com/888go/goframe/internal/json"
)

// Scan 自动检查`pointer`的类型，并将`params`转换为`pointer`。
// 它支持以下类型的`pointer`进行转换：`*map/*[]map/*[]*map/*struct/**struct/*[]struct/*[]*struct`。
//
// 待办：将`paramKeyToAttrMap`改为`ScanOption`以提高可扩展性；为`ScanOption`添加`DeepCopy`选项。
// md5:6b1a82a906dd8ec9
func Scan(srcValue interface{}, dstPointer interface{}, paramKeyToAttrMap ...map[string]string) (err error) {
	if srcValue == nil {
				// 如果`srcValue`为nil，不进行转换。 md5:72ad5bbbd6b824ce
		return nil
	}
	if dstPointer == nil {
		return gerror.NewCode(
			gcode.CodeInvalidParameter,
			`destination pointer should not be nil`,
		)
	}

	// json converting check.
	ok, err := doConvertWithJsonCheck(srcValue, dstPointer)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}

	var (
		dstPointerReflectType  reflect.Type
		dstPointerReflectValue reflect.Value
	)
	if v, ok := dstPointer.(reflect.Value); ok {
		dstPointerReflectValue = v
		dstPointerReflectType = v.Type()
	} else {
		dstPointerReflectValue = reflect.ValueOf(dstPointer)
				// 不要使用dstPointerReflectValue.Type()，因为dstPointerReflectValue可能为零。 md5:8930f50c4dd8479c
		dstPointerReflectType = reflect.TypeOf(dstPointer)
	}

		// 指针类型验证。 md5:4a9fffd5103820ed
	var dstPointerReflectKind = dstPointerReflectType.Kind()
	if dstPointerReflectKind != reflect.Ptr {
		if dstPointerReflectValue.CanAddr() {
			dstPointerReflectValue = dstPointerReflectValue.Addr()
			dstPointerReflectType = dstPointerReflectValue.Type()
			dstPointerReflectKind = dstPointerReflectType.Kind()
		} else {
			return gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`destination pointer should be type of pointer, but got type: %v`,
				dstPointerReflectType,
			)
		}
	}
			// 直接赋值检查！. md5:fd96cdd5962c2f14
	var srcValueReflectValue reflect.Value
	if v, ok := srcValue.(reflect.Value); ok {
		srcValueReflectValue = v
	} else {
		srcValueReflectValue = reflect.ValueOf(srcValue)
	}
	// 如果`srcValue`和`dstPointer`是相同的类型，直接进行赋值。
	// 为了提高性能。
	// md5:5d0efd73f7a58b6f
	var dstPointerReflectValueElem = dstPointerReflectValue.Elem()
	// 如果`srcValue`和`dstPointer`是相同的类型，直接进行赋值。
	// 为了提高性能。
	// md5:5d0efd73f7a58b6f
	if ok = doConvertWithTypeCheck(srcValueReflectValue, dstPointerReflectValueElem); ok {
		return nil
	}

	// do the converting.
	var (
		dstPointerReflectTypeElem     = dstPointerReflectType.Elem()
		dstPointerReflectTypeElemKind = dstPointerReflectTypeElem.Kind()
		keyToAttributeNameMapping     map[string]string
	)
	if len(paramKeyToAttrMap) > 0 {
		keyToAttributeNameMapping = paramKeyToAttrMap[0]
	}
	switch dstPointerReflectTypeElemKind {
	case reflect.Map:
		return doMapToMap(srcValue, dstPointer, paramKeyToAttrMap...)

	case reflect.Array, reflect.Slice:
		var (
			sliceElem     = dstPointerReflectTypeElem.Elem()
			sliceElemKind = sliceElem.Kind()
		)
		for sliceElemKind == reflect.Ptr {
			sliceElem = sliceElem.Elem()
			sliceElemKind = sliceElem.Kind()
		}
		if sliceElemKind == reflect.Map {
			return doMapToMaps(srcValue, dstPointer, paramKeyToAttrMap...)
		}
		return doStructs(srcValue, dstPointer, keyToAttributeNameMapping, "")

	default:
		return doStruct(srcValue, dstPointer, keyToAttributeNameMapping, "")
	}
}

func doConvertWithTypeCheck(srcValueReflectValue, dstPointerReflectValueElem reflect.Value) (ok bool) {
	if !dstPointerReflectValueElem.IsValid() || !srcValueReflectValue.IsValid() {
		return false
	}
	switch {
	// 示例：
	// 
	// UploadFile    => 上传文件
	// []UploadFile  => 上传文件切片
	// *UploadFile   => 指向上传文件的指针
	// *[]UploadFile => 指向上传文件切片的指针
	// map           => 映射
	// []map         => 映射切片
	// *[]map        => 指向映射切片的指针
	// md5:f6ba941ba3b0269f
	case dstPointerReflectValueElem.Type() == srcValueReflectValue.Type():
		dstPointerReflectValueElem.Set(srcValueReflectValue)
		return true

	// 示例：
	// UploadFile    => *UploadFile
	// []UploadFile  => *[]UploadFile
	// map           => *map
	// []map         => *[]map
	// 
	// 这些注释表示的是Go语言中的指针和数据结构的转换。在Go中，`*`符号用于表示指针类型。这里展示了如何将非指针类型转换为指针类型：
	// 
	// - `UploadFile` 是一个类型，`*UploadFile` 是它的指针类型。
	// - `[]UploadFile` 是 `UploadFile` 类型的切片（数组），`*[]UploadFile` 是这个切片类型的指针。
	// - `map` 表示一个未指定类型的映射，`*map` 则是这个映射类型的指针。
	// - `[]map` 表示一个包含映射的切片，`*[]map` 是这个切片类型的指针。
	// md5:3b743427a52ed67e
	case dstPointerReflectValueElem.Kind() == reflect.Ptr &&
		dstPointerReflectValueElem.Elem().IsValid() &&
		dstPointerReflectValueElem.Elem().Type() == srcValueReflectValue.Type():
		dstPointerReflectValueElem.Elem().Set(srcValueReflectValue)
		return true

	// 示例：
	// *UploadFile    => 上传文件
	// *[]UploadFile  => 上传文件的切片
	// *map           => 映射（字典）
	// *[]map         => 映射的切片
	// md5:a787c0f77f0eaa64
	case srcValueReflectValue.Kind() == reflect.Ptr &&
		srcValueReflectValue.Elem().IsValid() &&
		dstPointerReflectValueElem.Type() == srcValueReflectValue.Elem().Type():
		dstPointerReflectValueElem.Set(srcValueReflectValue.Elem())
		return true

	default:
		return false
	}
}

// doConvertWithJsonCheck 做 JSON 转换检查。
// 如果给定的 `params` 是 JSON，那么它会使用 json.Unmarshal 进行转换。
// md5:aa79d041fb48e2db
func doConvertWithJsonCheck(srcValue interface{}, dstPointer interface{}) (ok bool, err error) {
	switch valueResult := srcValue.(type) {
	case []byte:
		if json.Valid(valueResult) {
			if dstPointerReflectType, ok := dstPointer.(reflect.Value); ok {
				if dstPointerReflectType.Kind() == reflect.Ptr {
					if dstPointerReflectType.IsNil() {
						return false, nil
					}
					return true, json.UnmarshalUseNumber(valueResult, dstPointerReflectType.Interface())
				} else if dstPointerReflectType.CanAddr() {
					return true, json.UnmarshalUseNumber(valueResult, dstPointerReflectType.Addr().Interface())
				}
			} else {
				return true, json.UnmarshalUseNumber(valueResult, dstPointer)
			}
		}

	case string:
		if valueBytes := []byte(valueResult); json.Valid(valueBytes) {
			if dstPointerReflectType, ok := dstPointer.(reflect.Value); ok {
				if dstPointerReflectType.Kind() == reflect.Ptr {
					if dstPointerReflectType.IsNil() {
						return false, nil
					}
					return true, json.UnmarshalUseNumber(valueBytes, dstPointerReflectType.Interface())
				} else if dstPointerReflectType.CanAddr() {
					return true, json.UnmarshalUseNumber(valueBytes, dstPointerReflectType.Addr().Interface())
				}
			} else {
				return true, json.UnmarshalUseNumber(valueBytes, dstPointer)
			}
		}

	default:
				// `params`可能是一个实现了接口函数Interface的结构体，例如：gvar.Var。 md5:c02e870a76bafdaa
		if v, ok := srcValue.(iInterface); ok {
			return doConvertWithJsonCheck(v.Interface(), dstPointer)
		}
	}
	return false, nil
}
