// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gconv

import (
	"reflect"

	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/reflection"
)

// SliceFloat是Floats的别名。 md5:d770be8332715271
func SliceFloat(any interface{}) []float64 {
	return Floats(any)
}

// SliceFloat32 是 Float32s 的别名。 md5:686d576236624be0
func SliceFloat32(any interface{}) []float32 {
	return Float32s(any)
}

// SliceFloat64 是 Float64s 的别名。 md5:8fe51e4f2e2255df
func SliceFloat64(any interface{}) []float64 {
	return Floats(any)
}

// Floats 将 `any` 转换为 []float64。 md5:888b7822b64d033c
func Floats(any interface{}) []float64 {
	return Float64s(any)
}

// Float32s 将 `any` 转换为 []float32。 md5:e1b7a8e84a68d538
func Float32s(any interface{}) []float32 {
	if any == nil {
		return nil
	}
	var (
		array []float32 = nil
	)
	switch value := any.(type) {
	case string:
		if value == "" {
			return []float32{}
		}
		return []float32{Float32(value)}
	case []string:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []int:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []int8:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []int16:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []int32:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []int64:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []uint:
		for _, v := range value {
			array = append(array, Float32(v))
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &array)
		} else {
			array = make([]float32, len(value))
			for k, v := range value {
				array[k] = Float32(v)
			}
		}
	case []uint16:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []uint32:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []uint64:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []bool:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []float32:
		array = value
	case []float64:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []interface{}:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	}
	if array != nil {
		return array
	}
	if v, ok := any.(iFloats); ok {
		return Float32s(v.Floats())
	}
	if v, ok := any.(iInterfaces); ok {
		return Float32s(v.Interfaces())
	}
	// 将JSON格式的字符串值转换。 md5:60b4567e3f65e08a
	if checkJsonAndUnmarshalUseNumber(any, &array) {
		return array
	}
	// 并非常见类型，因此它使用反射来进行转换。 md5:a4126e9dfe7a56bd
	originValueAndKind := reflection.OriginValueAndKind(any)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		var (
			length = originValueAndKind.OriginValue.Len()
			slice  = make([]float32, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = Float32(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []float32{}
		}
		return []float32{Float32(any)}
	}
}

// Float64s 将 `any` 转换为 []float64。 md5:f66a670e62a3b46e
func Float64s(any interface{}) []float64 {
	if any == nil {
		return nil
	}
	var (
		array []float64 = nil
	)
	switch value := any.(type) {
	case string:
		if value == "" {
			return []float64{}
		}
		return []float64{Float64(value)}
	case []string:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []int:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []int8:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []int16:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []int32:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []int64:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []uint:
		for _, v := range value {
			array = append(array, Float64(v))
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &array)
		} else {
			array = make([]float64, len(value))
			for k, v := range value {
				array[k] = Float64(v)
			}
		}
	case []uint16:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []uint32:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []uint64:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []bool:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []float32:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []float64:
		array = value
	case []interface{}:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	}
	if array != nil {
		return array
	}
	if v, ok := any.(iFloats); ok {
		return v.Floats()
	}
	if v, ok := any.(iInterfaces); ok {
		return Floats(v.Interfaces())
	}
	// 将JSON格式的字符串值转换。 md5:60b4567e3f65e08a
	if checkJsonAndUnmarshalUseNumber(any, &array) {
		return array
	}
	// 并非常见类型，因此它使用反射来进行转换。 md5:a4126e9dfe7a56bd
	originValueAndKind := reflection.OriginValueAndKind(any)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		var (
			length = originValueAndKind.OriginValue.Len()
			slice  = make([]float64, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = Float64(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []float64{}
		}
		return []float64{Float64(any)}
	}
}
