// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gutil提供了实用工具函数。
package gutil

import (
	"reflect"
	
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/util/gconv"
)

const (
	dumpIndent = `    `
)

// IsEmpty 检查给定的 `value` 是否为空。
// 当 `value` 为：整数(0)、布尔值(false)、切片/映射(len=0)或 nil 时，返回 false；
// 否则返回 true。
func IsEmpty(value interface{}) bool {
	return empty.IsEmpty(value)
}

// Keys 从给定的 map 或 struct 中获取并返回其键（key）。
func Keys(mapOrStruct interface{}) (keysOrAttrs []string) {
	keysOrAttrs = make([]string, 0)
	if m, ok := mapOrStruct.(map[string]interface{}); ok {
		for k := range m {
			keysOrAttrs = append(keysOrAttrs, k)
		}
		return
	}
	var (
		reflectValue reflect.Value
		reflectKind  reflect.Kind
	)
	if v, ok := mapOrStruct.(reflect.Value); ok {
		reflectValue = v
	} else {
		reflectValue = reflect.ValueOf(mapOrStruct)
	}
	reflectKind = reflectValue.Kind()
	for reflectKind == reflect.Ptr {
		if !reflectValue.IsValid() || reflectValue.IsNil() {
			reflectValue = reflect.New(reflectValue.Type().Elem()).Elem()
			reflectKind = reflectValue.Kind()
		} else {
			reflectValue = reflectValue.Elem()
			reflectKind = reflectValue.Kind()
		}
	}
	switch reflectKind {
	case reflect.Map:
		for _, k := range reflectValue.MapKeys() {
			keysOrAttrs = append(keysOrAttrs, gconv.String(k.Interface()))
		}
	case reflect.Struct:
		var (
			fieldType   reflect.StructField
			reflectType = reflectValue.Type()
		)
		for i := 0; i < reflectValue.NumField(); i++ {
			fieldType = reflectType.Field(i)
			if fieldType.Anonymous {
				keysOrAttrs = append(keysOrAttrs, Keys(reflectValue.Field(i))...)
			} else {
				keysOrAttrs = append(keysOrAttrs, fieldType.Name)
			}
		}
	}
	return
}

// Values 从给定的 map 或 struct 中获取并返回其值。
func Values(mapOrStruct interface{}) (values []interface{}) {
	values = make([]interface{}, 0)
	if m, ok := mapOrStruct.(map[string]interface{}); ok {
		for _, v := range m {
			values = append(values, v)
		}
		return
	}
	var (
		reflectValue reflect.Value
		reflectKind  reflect.Kind
	)
	if v, ok := mapOrStruct.(reflect.Value); ok {
		reflectValue = v
	} else {
		reflectValue = reflect.ValueOf(mapOrStruct)
	}
	reflectKind = reflectValue.Kind()
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	switch reflectKind {
	case reflect.Map:
		for _, k := range reflectValue.MapKeys() {
			values = append(values, reflectValue.MapIndex(k).Interface())
		}
	case reflect.Struct:
		var (
			fieldType   reflect.StructField
			reflectType = reflectValue.Type()
		)
		for i := 0; i < reflectValue.NumField(); i++ {
			fieldType = reflectType.Field(i)
			if fieldType.Anonymous {
				values = append(values, Values(reflectValue.Field(i))...)
			} else {
				values = append(values, reflectValue.Field(i).Interface())
			}
		}
	}
	return
}
