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

// SliceInt 是 Ints 的别名。
func SliceIne别名(any interface{}) []int {
	return X取整数数组(any)
}

// SliceInt32 是 Int32s 的别名。
func SliceIet32别名(any interface{}) []int32 {
	return X取整数32位数组(any)
}

// SliceInt64 是 Int64s 的别名。
func SliceInt64别名(any interface{}) []int64 {
	return X取整数64位数组(any)
}

// Ints 将 `any` 类型转换为 []int 类型。
func X取整数数组(值 interface{}) []int {
	if 值 == nil {
		return nil
	}
	var (
		array []int = nil
	)
	switch value := 值.(type) {
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
	if v, ok := 值.(iInts); ok {
		return v.X取整数数组()
	}
	if v, ok := 值.(iInterfaces); ok {
		return X取整数数组(v.X取any数组())
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
		return []int{X取整数(值)}
	}
}

// Int32s将`any`转换为[]int32。
func X取整数32位数组(值 interface{}) []int32 {
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
		return X取整数32位数组(v.X取整数数组())
	}
	if v, ok := 值.(iInterfaces); ok {
		return X取整数32位数组(v.X取any数组())
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

// Int64s 将 `any` 转换为 []int64 类型的切片。
func X取整数64位数组(值 interface{}) []int64 {
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
		return X取整数64位数组(v.X取整数数组())
	}
	if v, ok := 值.(iInterfaces); ok {
		return X取整数64位数组(v.X取any数组())
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
