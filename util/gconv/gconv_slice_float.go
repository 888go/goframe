// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类

import (
	"reflect"
	
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/reflection"
)

// SliceFloat 是 Floats 的别名。
func SliceFloat别名(any interface{}) []float64 {
	return X取小数数组(any)
}

// SliceFloat32 是 Float32s 的别名。
func SliceFloat32别名(any interface{}) []float32 {
	return X取小数32位数组(any)
}

// SliceFloat64 是 Float64s 的别名。
func SliceFloat64别名(any interface{}) []float64 {
	return X取小数数组(any)
}

// Floats 将 `any` 转换为 []float64 类型的切片。
func X取小数数组(值 interface{}) []float64 {
	return X取小数64位数组(值)
}

// Float32s 将 `any` 类型转换为 []float32 类型。
func X取小数32位数组(值 interface{}) []float32 {
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
		return X取小数32位数组(v.X取小数数组())
	}
	if v, ok := 值.(iInterfaces); ok {
		return X取小数32位数组(v.X取any数组())
	}
	// JSON格式字符串值转换
	if checkJsonAndUnmarshalUseNumber(值, &array) {
		return array
	}
	// 如果不是常见类型，它将使用反射进行转换。
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

// Float64s将`any`转换为[]float64。
func X取小数64位数组(值 interface{}) []float64 {
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
		return v.X取小数数组()
	}
	if v, ok := 值.(iInterfaces); ok {
		return X取小数数组(v.X取any数组())
	}
	// JSON格式字符串值转换
	if checkJsonAndUnmarshalUseNumber(值, &array) {
		return array
	}
	// 如果不是常见类型，它将使用反射进行转换。
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
