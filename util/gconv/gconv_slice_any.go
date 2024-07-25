// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gconv

import (
	"reflect"

	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/reflection"
)

// SliceAny 是 Interfaces 的别名。 md5:4a67bb7703f4770e
func SliceAny(any interface{}) []interface{} {
	return Interfaces(any)
}

// Interfaces 将 `any` 转换为 []interface{}。 md5:74035d098fdf06f6
func Interfaces(any interface{}) []interface{} {
	if any == nil {
		return nil
	}
	var array []interface{}
	switch value := any.(type) {
	case []interface{}:
		array = value
	case []string:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []int:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []int8:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []int16:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []int32:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []int64:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []uint:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &array)
		} else {
			array = make([]interface{}, len(value))
			for k, v := range value {
				array[k] = v
			}
		}
	case []uint16:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []uint32:
		for _, v := range value {
			array = append(array, v)
		}
	case []uint64:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []bool:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []float32:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []float64:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	}
	if array != nil {
		return array
	}
	if v, ok := any.(iInterfaces); ok {
		return v.Interfaces()
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
			slice  = make([]interface{}, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = originValueAndKind.OriginValue.Index(i).Interface()
		}
		return slice

	default:
		return []interface{}{any}
	}
}
