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

// SliceStr别名是Strings的别名。 md5:dacb4ebc45c023cf
func SliceStr别名(any interface{}) []string {
	return X取文本切片(any)
}

// X取文本切片 将 `any` 转换为 []string。 md5:cbac28ee26158116
func X取文本切片(any interface{}) []string {
	if any == nil {
		return nil
	}
	var (
		array []string = nil
	)
	switch value := any.(type) {
	case []int:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = String(v)
		}
	case []int8:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = String(v)
		}
	case []int16:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = String(v)
		}
	case []int32:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = String(v)
		}
	case []int64:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = String(v)
		}
	case []uint:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = String(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &array)
		}
		if array == nil {
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
			return array
		}
	case string:
		byteValue := []byte(value)
		if json.Valid(byteValue) {
			_ = json.UnmarshalUseNumber(byteValue, &array)
		}
		if array == nil {
			if value == "" {
				return []string{}
			}
			// 防止字符串为null
			// 查看问题3465以获取详细信息
			// md5:7177702384700ffb
			return []string{value}
		}
	case []uint16:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = String(v)
		}
	case []uint32:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = String(v)
		}
	case []uint64:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = String(v)
		}
	case []bool:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = String(v)
		}
	case []float32:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = String(v)
		}
	case []float64:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = String(v)
		}
	case []interface{}:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = String(v)
		}
	case []string:
		array = value
	case [][]byte:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = String(v)
		}
	}
	if array != nil {
		return array
	}
	if v, ok := any.(iStrings); ok {
		return v.X取文本切片()
	}
	if v, ok := any.(iInterfaces); ok {
		return X取文本切片(v.X取any切片())
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
			slice  = make([]string, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = String(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []string{}
		}
		return []string{String(any)}
	}
}
