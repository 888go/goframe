// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package utils

import (
	"reflect"

	"github.com/gogf/gf/v2/internal/empty"
)

// IsNil 检查 `value` 是否为 nil，特别是对于 interface{} 类型的值。. md5:7789b39820de4e0c
func IsNil(value interface{}) bool {
	return empty.IsNil(value)
}

// IsEmpty 检查 `value` 是否为空。. md5:b2015a35f4930a22
func IsEmpty(value interface{}) bool {
	return empty.IsEmpty(value)
}

// IsInt 检查 `value` 是否为 int 类型。. md5:f996dc2da0e38f8d
func IsInt(value interface{}) bool {
	switch value.(type) {
	case int, *int, int8, *int8, int16, *int16, int32, *int32, int64, *int64:
		return true
	}
	return false
}

// IsUint 检查 `value` 是否为 uint 类型。. md5:cc17b6c670e71ccd
func IsUint(value interface{}) bool {
	switch value.(type) {
	case uint, *uint, uint8, *uint8, uint16, *uint16, uint32, *uint32, uint64, *uint64:
		return true
	}
	return false
}

// IsFloat检查`value`是否为浮点类型。. md5:11f3d9e95720de15
func IsFloat(value interface{}) bool {
	switch value.(type) {
	case float32, *float32, float64, *float64:
		return true
	}
	return false
}

// IsSlice 检查 `value` 是否为切片类型。. md5:7ea6cb09ee3f9841
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

// IsMap 检查 `value` 是否为映射类型。. md5:a3bcf5df99373e2a
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

// IsStruct 检查 `value` 是否为结构体类型。. md5:3be94afb9f9b40ea
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
