// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类

import (
	"reflect"
	"strings"

	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/internal/utils"
)

// SliceUint是Uints的别名。 md5:75cea5e3f6b32ecb
func SliceUint(any interface{}) []uint {
	return Uints(any)
}

// SliceUint32是Uint32s的别名。 md5:128b14c8ffd2f906
func SliceUint32(any interface{}) []uint32 {
	return Uint32s(any)
}

// SliceUint64 是 Uint64s 的别名。 md5:8436f35c37880e8c
func SliceUint64(any interface{}) []uint64 {
	return Uint64s(any)
}

// Uints 将 `any` 转换为 []uint。 md5:5be2a9bb94384311
func Uints(any interface{}) []uint {
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
			return []uint{Uint(value)}
		}

	case []string:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = Uint(v)
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
			array[k] = Uint(v)
		}
	case []float64:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = Uint(v)
		}
	case []interface{}:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = Uint(v)
		}
	case [][]byte:
		array = make([]uint, len(value))
		for k, v := range value {
			array[k] = Uint(v)
		}
	}

	if array != nil {
		return array
	}

	// Default handler.
	if v, ok := any.(iUints); ok {
		return v.Uints()
	}
	if v, ok := any.(iInterfaces); ok {
		return Uints(v.Interfaces())
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
			slice  = make([]uint, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = Uint(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []uint{}
		}
		return []uint{Uint(any)}
	}
}

// Uint32s 将 `any` 转换为 []uint32。 md5:7a3545642e764e37
func Uint32s(any interface{}) []uint32 {
	if any == nil {
		return nil
	}
	var (
		array []uint32 = nil
	)
	switch value := any.(type) {
	case string:
		value = strings.TrimSpace(value)
		if value == "" {
			return []uint32{}
		}
		if utils.IsNumeric(value) {
			return []uint32{Uint32(value)}
		}
	case []string:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = Uint32(v)
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
			array[k] = Uint32(v)
		}
	case []float64:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = Uint32(v)
		}
	case []interface{}:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = Uint32(v)
		}
	case [][]byte:
		array = make([]uint32, len(value))
		for k, v := range value {
			array[k] = Uint32(v)
		}
	}
	if array != nil {
		return array
	}

	// Default handler.
	if v, ok := any.(iUints); ok {
		return Uint32s(v.Uints())
	}
	if v, ok := any.(iInterfaces); ok {
		return Uint32s(v.Interfaces())
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
			slice  = make([]uint32, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = Uint32(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []uint32{}
		}
		return []uint32{Uint32(any)}
	}
}

// Uint64s 将 `any` 转换为 []uint64。 md5:08b12c62032337ab
func Uint64s(any interface{}) []uint64 {
	if any == nil {
		return nil
	}
	var (
		array []uint64 = nil
	)
	switch value := any.(type) {
	case string:
		value = strings.TrimSpace(value)
		if value == "" {
			return []uint64{}
		}
		if utils.IsNumeric(value) {
			return []uint64{Uint64(value)}
		}

	case []string:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = Uint64(v)
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
			array[k] = Uint64(v)
		}
	case []float64:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = Uint64(v)
		}
	case []interface{}:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = Uint64(v)
		}
	case [][]byte:
		array = make([]uint64, len(value))
		for k, v := range value {
			array[k] = Uint64(v)
		}
	}
	if array != nil {
		return array
	}
	// Default handler.
	if v, ok := any.(iUints); ok {
		return Uint64s(v.Uints())
	}
	if v, ok := any.(iInterfaces); ok {
		return Uint64s(v.Interfaces())
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
			slice  = make([]uint64, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = Uint64(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []uint64{}
		}
		return []uint64{Uint64(any)}
	}
}
