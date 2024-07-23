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

// SliceInt 是 Ints 的别名。 md5:1918d6b770b537c7
func SliceInt(any interface{}) []int {
	return Ints(any)
}

// SliceInt32 是 Int32s 的别名。 md5:ea50c4495697e1ad
func SliceInt32(any interface{}) []int32 {
	return Int32s(any)
}

// SliceInt64 是 Int64s 的别名。 md5:c8c6a11351f3c7dc
func SliceInt64(any interface{}) []int64 {
	return Int64s(any)
}

// Ints 将 `any` 类型转换为 []int 类型。 md5:2e68bd53fc00e307
func Ints(any interface{}) []int {
	if any == nil {
		return nil
	}
	var (
		array []int = nil
	)
	switch value := any.(type) {
	case []string:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = Int(v)
		}
	case []int:
		array = value
	case []int8:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = int(v)
		}
	case []int16:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = int(v)
		}
	case []int32:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = int(v)
		}
	case []int64:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = int(v)
		}
	case []uint:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = int(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &array)
		} else {
			array = make([]int, len(value))
			for k, v := range value {
				array[k] = int(v)
			}
		}
	case []uint16:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = int(v)
		}
	case []uint32:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = int(v)
		}
	case []uint64:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = int(v)
		}
	case []bool:
		array = make([]int, len(value))
		for k, v := range value {
			if v {
				array[k] = 1
			} else {
				array[k] = 0
			}
		}
	case []float32:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = Int(v)
		}
	case []float64:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = Int(v)
		}
	case []interface{}:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = Int(v)
		}
	case [][]byte:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = Int(v)
		}
	}
	if array != nil {
		return array
	}
	if v, ok := any.(iInts); ok {
		return v.Ints()
	}
	if v, ok := any.(iInterfaces); ok {
		return Ints(v.Interfaces())
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
			slice  = make([]int, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = Int(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []int{}
		}
		return []int{Int(any)}
	}
}

// Int32s 将 `any` 转换为 []int32。 md5:3b96167e68fc609e
func Int32s(any interface{}) []int32 {
	if any == nil {
		return nil
	}
	var (
		array []int32 = nil
	)
	switch value := any.(type) {
	case []string:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = Int32(v)
		}
	case []int:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = int32(v)
		}
	case []int8:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = int32(v)
		}
	case []int16:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = int32(v)
		}
	case []int32:
		array = value
	case []int64:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = int32(v)
		}
	case []uint:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = int32(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &array)
		} else {
			array = make([]int32, len(value))
			for k, v := range value {
				array[k] = int32(v)
			}
		}
	case []uint16:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = int32(v)
		}
	case []uint32:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = int32(v)
		}
	case []uint64:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = int32(v)
		}
	case []bool:
		array = make([]int32, len(value))
		for k, v := range value {
			if v {
				array[k] = 1
			} else {
				array[k] = 0
			}
		}
	case []float32:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = Int32(v)
		}
	case []float64:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = Int32(v)
		}
	case []interface{}:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = Int32(v)
		}
	case [][]byte:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = Int32(v)
		}
	}
	if array != nil {
		return array
	}
	if v, ok := any.(iInts); ok {
		return Int32s(v.Ints())
	}
	if v, ok := any.(iInterfaces); ok {
		return Int32s(v.Interfaces())
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
			slice  = make([]int32, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = Int32(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []int32{}
		}
		return []int32{Int32(any)}
	}
}

// Int64s 将 `any` 转换为 []int64。 md5:c28e69b40a68187f
func Int64s(any interface{}) []int64 {
	if any == nil {
		return nil
	}
	var (
		array []int64 = nil
	)
	switch value := any.(type) {
	case []string:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = Int64(v)
		}
	case []int:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = int64(v)
		}
	case []int8:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = int64(v)
		}
	case []int16:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = int64(v)
		}
	case []int32:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = int64(v)
		}
	case []int64:
		array = value
	case []uint:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = int64(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &array)
		} else {
			array = make([]int64, len(value))
			for k, v := range value {
				array[k] = int64(v)
			}
		}
	case []uint16:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = int64(v)
		}
	case []uint32:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = int64(v)
		}
	case []uint64:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = int64(v)
		}
	case []bool:
		array = make([]int64, len(value))
		for k, v := range value {
			if v {
				array[k] = 1
			} else {
				array[k] = 0
			}
		}
	case []float32:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = Int64(v)
		}
	case []float64:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = Int64(v)
		}
	case []interface{}:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = Int64(v)
		}
	case [][]byte:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = Int64(v)
		}
	}
	if array != nil {
		return array
	}
	if v, ok := any.(iInts); ok {
		return Int64s(v.Ints())
	}
	if v, ok := any.(iInterfaces); ok {
		return Int64s(v.Interfaces())
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
			slice  = make([]int64, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = Int64(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []int64{}
		}
		return []int64{Int64(any)}
	}
}
