// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类

import (
	"reflect"
	"strings"
	
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/internal/utils"
)

// SliceUint 是 Uints 的别名。
func SliceUint别名(any interface{}) []uint {
	return X取正整数数组(any)
}

// SliceUint32 是 Uint32s 的别名。
func SliceUint32别名(any interface{}) []uint32 {
	return X取正整数32位数组(any)
}

// SliceUint64 是 Uint64s 的别名。
func SliceUint64别名(any interface{}) []uint64 {
	return X取正整数64位数组(any)
}

// Uints 将 `any` 转换为 []uint 类型的切片。
func X取正整数数组(any interface{}) []uint {
	if any == nil {
		return nil
	}

	var (
		array []uint = nil
	)
	switch value := any.(type) {
	case string:
		value = strings.TrimSpace(value)
		if value == "" {
			return []uint{}
		}
		if utils.IsNumeric(value) {
			return []uint{X取正整数(value)}
		}

	case []string:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = X取正整数(v)
		}
	case []int8:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = uint(v)
		}
	case []int16:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = uint(v)
		}
	case []int32:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = uint(v)
		}
	case []int64:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = uint(v)
		}
	case []uint:
		array = value
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &array)
		} else {
			array = make([]uint, len(value))
			for k, v := range value {
				array[k] = uint(v)
			}
		}
	case []uint16:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = uint(v)
		}
	case []uint32:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = uint(v)
		}
	case []uint64:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = uint(v)
		}
	case []bool:
		array = make([]uint, len(value))
		for k, v := range value {
			if v {
				array[k] = 1
			} else {
				array[k] = 0
			}
		}
	case []float32:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = X取正整数(v)
		}
	case []float64:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = X取正整数(v)
		}
	case []interface{}:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = X取正整数(v)
		}
	case [][]byte:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = X取正整数(v)
		}
	}

	if array != nil {
		return array
	}

	// Default handler.
	if v, ok := any.(iUints); ok {
		return v.X取正整数数组()
	}
	if v, ok := any.(iInterfaces); ok {
		return X取正整数数组(v.X取any数组())
	}
	// JSON格式字符串值转换
	if checkJsonAndUnmarshalUseNumber(any, &array) {
		return array
	}
	// 如果不是常见类型，它将使用反射进行转换。
	originValueAndKind := reflection.OriginValueAndKind(any)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		var (
			length = originValueAndKind.OriginValue.Len()
			slice  = make([]uint, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = X取正整数(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []uint{}
		}
		return []uint{X取正整数(any)}
	}
}

// Uint32s将`any`转换为[]uint32。
func X取正整数32位数组(值 interface{}) []uint32 {
	if 值 == nil {
		return nil
	}
	var (
		array []uint32 = nil
	)
	switch value := 值.(type) {
	case string:
		value = strings.TrimSpace(value)
		if value == "" {
			return []uint32{}
		}
		if utils.IsNumeric(value) {
			return []uint32{X取正整数32位(value)}
		}
	case []string:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = X取正整数32位(v)
		}
	case []int8:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = uint32(v)
		}
	case []int16:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = uint32(v)
		}
	case []int32:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = uint32(v)
		}
	case []int64:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = uint32(v)
		}
	case []uint:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = uint32(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &array)
		} else {
			array = make([]uint32, len(value))
			for k, v := range value {
				array[k] = uint32(v)
			}
		}
	case []uint16:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = uint32(v)
		}
	case []uint32:
		array = value
	case []uint64:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = uint32(v)
		}
	case []bool:
		array = make([]uint32, len(value))
		for k, v := range value {
			if v {
				array[k] = 1
			} else {
				array[k] = 0
			}
		}
	case []float32:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = X取正整数32位(v)
		}
	case []float64:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = X取正整数32位(v)
		}
	case []interface{}:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = X取正整数32位(v)
		}
	case [][]byte:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = X取正整数32位(v)
		}
	}
	if array != nil {
		return array
	}

	// Default handler.
	if v, ok := 值.(iUints); ok {
		return X取正整数32位数组(v.X取正整数数组())
	}
	if v, ok := 值.(iInterfaces); ok {
		return X取正整数32位数组(v.X取any数组())
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
			slice  = make([]uint32, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = X取正整数32位(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []uint32{}
		}
		return []uint32{X取正整数32位(值)}
	}
}

// Uint64s 将 `any` 类型转换为 []uint64 类型。
func X取正整数64位数组(值 interface{}) []uint64 {
	if 值 == nil {
		return nil
	}
	var (
		array []uint64 = nil
	)
	switch value := 值.(type) {
	case string:
		value = strings.TrimSpace(value)
		if value == "" {
			return []uint64{}
		}
		if utils.IsNumeric(value) {
			return []uint64{X取正整数64位(value)}
		}

	case []string:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = X取正整数64位(v)
		}
	case []int8:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = uint64(v)
		}
	case []int16:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = uint64(v)
		}
	case []int32:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = uint64(v)
		}
	case []int64:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = uint64(v)
		}
	case []uint:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = uint64(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &array)
		} else {
			array = make([]uint64, len(value))
			for k, v := range value {
				array[k] = uint64(v)
			}
		}
	case []uint16:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = uint64(v)
		}
	case []uint32:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = uint64(v)
		}
	case []uint64:
		array = value
	case []bool:
		array = make([]uint64, len(value))
		for k, v := range value {
			if v {
				array[k] = 1
			} else {
				array[k] = 0
			}
		}
	case []float32:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = X取正整数64位(v)
		}
	case []float64:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = X取正整数64位(v)
		}
	case []interface{}:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = X取正整数64位(v)
		}
	case [][]byte:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = X取正整数64位(v)
		}
	}
	if array != nil {
		return array
	}
	// Default handler.
	if v, ok := 值.(iUints); ok {
		return X取正整数64位数组(v.X取正整数数组())
	}
	if v, ok := 值.(iInterfaces); ok {
		return X取正整数64位数组(v.X取any数组())
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
			slice  = make([]uint64, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = X取正整数64位(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []uint64{}
		}
		return []uint64{X取正整数64位(值)}
	}
}
