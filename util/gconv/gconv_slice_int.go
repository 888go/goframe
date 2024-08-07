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

// SliceIne别名 是 Ints 的别名。 md5:1918d6b770b537c7
func SliceIne别名(any interface{}) []int {
	return X取整数切片(any)
}

// SliceIet32别名 是 Int32s 的别名。 md5:ea50c4495697e1ad
func SliceIet32别名(any interface{}) []int32 {
	return X取整数32位切片(any)
}

// SliceInt64别名 是 Int64s 的别名。 md5:c8c6a11351f3c7dc
func SliceInt64别名(any interface{}) []int64 {
	return X取整数64位切片(any)
}

// X取整数切片 将 `any` 类型转换为 []int 类型。 md5:2e68bd53fc00e307
func X取整数切片(any interface{}) []int {
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
			array[k] = X取整数(v)
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
			array[k] = X取整数(v)
		}
	case []float64:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = X取整数(v)
		}
	case []interface{}:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = X取整数(v)
		}
	case [][]byte:
		array = make([]int, len(value))
		for k, v := range value {
			array[k] = X取整数(v)
		}
	}
	if array != nil {
		return array
	}
	if v, ok := any.(iInts); ok {
		return v.X取整数切片()
	}
	if v, ok := any.(iInterfaces); ok {
		return X取整数切片(v.X取any切片())
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
			slice[i] = X取整数(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []int{}
		}
		return []int{X取整数(any)}
	}
}

// X取整数32位切片 将 `any` 转换为 []int32。 md5:3b96167e68fc609e
func X取整数32位切片(值 interface{}) []int32 {
	if 值 == nil {
		return nil
	}
	var (
		array []int32 = nil
	)
	switch value := 值.(type) {
	case []string:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = X取整数32位(v)
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
			array[k] = X取整数32位(v)
		}
	case []float64:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = X取整数32位(v)
		}
	case []interface{}:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = X取整数32位(v)
		}
	case [][]byte:
		array = make([]int32, len(value))
		for k, v := range value {
			array[k] = X取整数32位(v)
		}
	}
	if array != nil {
		return array
	}
	if v, ok := 值.(iInts); ok {
		return X取整数32位切片(v.X取整数切片())
	}
	if v, ok := 值.(iInterfaces); ok {
		return X取整数32位切片(v.X取any切片())
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
			slice  = make([]int32, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = X取整数32位(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []int32{}
		}
		return []int32{X取整数32位(值)}
	}
}

// X取整数64位切片 将 `any` 转换为 []int64。 md5:c28e69b40a68187f
func X取整数64位切片(值 interface{}) []int64 {
	if 值 == nil {
		return nil
	}
	var (
		array []int64 = nil
	)
	switch value := 值.(type) {
	case []string:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = X取整数64位(v)
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
			array[k] = X取整数64位(v)
		}
	case []float64:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = X取整数64位(v)
		}
	case []interface{}:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = X取整数64位(v)
		}
	case [][]byte:
		array = make([]int64, len(value))
		for k, v := range value {
			array[k] = X取整数64位(v)
		}
	}
	if array != nil {
		return array
	}
	if v, ok := 值.(iInts); ok {
		return X取整数64位切片(v.X取整数切片())
	}
	if v, ok := 值.(iInterfaces); ok {
		return X取整数64位切片(v.X取any切片())
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
			slice  = make([]int64, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = X取整数64位(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []int64{}
		}
		return []int64{X取整数64位(值)}
	}
}
