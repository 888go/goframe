// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类

import (
	"reflect"

	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/reflection"
)

// SliceFloat别名是Floats的别名。 md5:d770be8332715271
func SliceFloat别名(any interface{}) []float64 {
	return X取小数切片(any)
}

// SliceFloat32别名 是 Float32s 的别名。 md5:686d576236624be0
func SliceFloat32别名(any interface{}) []float32 {
	return X取小数32位切片(any)
}

// SliceFloat64别名 是 Float64s 的别名。 md5:8fe51e4f2e2255df
func SliceFloat64别名(any interface{}) []float64 {
	return X取小数切片(any)
}

// X取小数切片 将 `any` 转换为 []float64。 md5:888b7822b64d033c
func X取小数切片(any interface{}) []float64 {
	return X取小数64位切片(any)
}

// X取小数32位切片 将 `any` 转换为 []float32。 md5:e1b7a8e84a68d538
func X取小数32位切片(值 interface{}) []float32 {
	if 值 == nil {
		return nil
	}
	var (
		array []float32 = nil
	)
	switch value := 值.(type) {
	case string:
		if value == "" {
			return []float32{}
		}
		return []float32{X取小数32位(value)}
	case []string:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = X取小数32位(v)
		}
	case []int:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = X取小数32位(v)
		}
	case []int8:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = X取小数32位(v)
		}
	case []int16:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = X取小数32位(v)
		}
	case []int32:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = X取小数32位(v)
		}
	case []int64:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = X取小数32位(v)
		}
	case []uint:
		for _, v := range value {
			array = append(array, X取小数32位(v))
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &array)
		} else {
			array = make([]float32, len(value))
			for k, v := range value {
				array[k] = X取小数32位(v)
			}
		}
	case []uint16:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = X取小数32位(v)
		}
	case []uint32:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = X取小数32位(v)
		}
	case []uint64:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = X取小数32位(v)
		}
	case []bool:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = X取小数32位(v)
		}
	case []float32:
		array = value
	case []float64:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = X取小数32位(v)
		}
	case []interface{}:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = X取小数32位(v)
		}
	}
	if array != nil {
		return array
	}
	if v, ok := 值.(iFloats); ok {
		return X取小数32位切片(v.X取小数切片())
	}
	if v, ok := 值.(iInterfaces); ok {
		return X取小数32位切片(v.X取any切片())
	}
		// 将JSON格式的字符串值转换。 md5:60b4567e3f65e08a
	if checkJsonAndUnmarshalUseNumber(值, &array) {
		return array
	}
		// 并非常见类型，因此它使用反射来进行转换。 md5:a4126e9dfe7a56bd
	originValueAndKind := reflection.OriginValueAndKind(值)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		var (
			length = originValueAndKind.OriginValue.Len()
			slice  = make([]float32, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = X取小数32位(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []float32{}
		}
		return []float32{X取小数32位(值)}
	}
}

// X取小数64位切片 将 `any` 转换为 []float64。 md5:f66a670e62a3b46e
func X取小数64位切片(值 interface{}) []float64 {
	if 值 == nil {
		return nil
	}
	var (
		array []float64 = nil
	)
	switch value := 值.(type) {
	case string:
		if value == "" {
			return []float64{}
		}
		return []float64{X取小数64位(value)}
	case []string:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = X取小数64位(v)
		}
	case []int:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = X取小数64位(v)
		}
	case []int8:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = X取小数64位(v)
		}
	case []int16:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = X取小数64位(v)
		}
	case []int32:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = X取小数64位(v)
		}
	case []int64:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = X取小数64位(v)
		}
	case []uint:
		for _, v := range value {
			array = append(array, X取小数64位(v))
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &array)
		} else {
			array = make([]float64, len(value))
			for k, v := range value {
				array[k] = X取小数64位(v)
			}
		}
	case []uint16:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = X取小数64位(v)
		}
	case []uint32:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = X取小数64位(v)
		}
	case []uint64:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = X取小数64位(v)
		}
	case []bool:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = X取小数64位(v)
		}
	case []float32:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = X取小数64位(v)
		}
	case []float64:
		array = value
	case []interface{}:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = X取小数64位(v)
		}
	}
	if array != nil {
		return array
	}
	if v, ok := 值.(iFloats); ok {
		return v.X取小数切片()
	}
	if v, ok := 值.(iInterfaces); ok {
		return X取小数切片(v.X取any切片())
	}
		// 将JSON格式的字符串值转换。 md5:60b4567e3f65e08a
	if checkJsonAndUnmarshalUseNumber(值, &array) {
		return array
	}
		// 并非常见类型，因此它使用反射来进行转换。 md5:a4126e9dfe7a56bd
	originValueAndKind := reflection.OriginValueAndKind(值)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		var (
			length = originValueAndKind.OriginValue.Len()
			slice  = make([]float64, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = X取小数64位(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []float64{}
		}
		return []float64{X取小数64位(值)}
	}
}
