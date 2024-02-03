// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gconv

import (
	"reflect"
	
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/reflection"
)

// SliceAny 是 Interfaces 的别名。
func SliceAny(any interface{}) []interface{} {
	return Interfaces(any)
}

// Interfaces 将 `any` 类型转换为 []interface{} 类型。
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
