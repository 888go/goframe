// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package utils

import (
	"reflect"
	
	"github.com/888go/goframe/internal/empty"
)

// IsNil 检查 `value` 是否为 nil，特别是针对 interface{} 类型的值。
func X是否为Nil(value interface{}) bool {
	return empty.X是否为Nil(value)
}

// IsEmpty 检查 `value` 是否为空。
func IsEmpty(value interface{}) bool {
	return empty.IsEmpty(value)
}

// IsInt 检查 `value` 是否为 int 类型。
func IsInt(value interface{}) bool {
	switch value.(type) {
	case int, *int, int8, *int8, int16, *int16, int32, *int32, int64, *int64:
		return true
	}
	return false
}

// IsUint 检查 `value` 是否为 uint 类型。
func IsUint(value interface{}) bool {
	switch value.(type) {
	case uint, *uint, uint8, *uint8, uint16, *uint16, uint32, *uint32, uint64, *uint64:
		return true
	}
	return false
}

// IsFloat 检查 `value` 是否为浮点类型。
func IsFloat(value interface{}) bool {
	switch value.(type) {
	case float32, *float32, float64, *float64:
		return true
	}
	return false
}

// IsSlice 检查 `value` 是否为切片类型。
func IsSlice(value interface{}) bool {
	var (
		reflectValue = reflect.ValueOf(value)
		reflectKind  = reflectValue.Kind()
	)
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	switch reflectKind {
	case reflect.Slice, reflect.Array:
		return true
	}
	return false
}

// IsMap 检查 `value` 是否为 map 类型。
func IsMap(value interface{}) bool {
	var (
		reflectValue = reflect.ValueOf(value)
		reflectKind  = reflectValue.Kind()
	)
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	switch reflectKind {
	case reflect.Map:
		return true
	}
	return false
}

// IsStruct 检查 `value` 是否为结构体类型。
func IsStruct(value interface{}) bool {
	var reflectType = reflect.TypeOf(value)
	if reflectType == nil {
		return false
	}
	var reflectKind = reflectType.Kind()
	for reflectKind == reflect.Ptr {
		reflectType = reflectType.Elem()
		reflectKind = reflectType.Kind()
	}
	switch reflectKind {
	case reflect.Struct:
		return true
	}
	return false
}
