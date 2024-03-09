// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gutil提供了实用工具函数。
package 工具类

import (
	"reflect"
	
	"github.com/888go/goframe/gutil/internal/empty"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	dumpIndent = `    `
)

// IsEmpty 检查给定的 `value` 是否为空。
// 当 `value` 为：整数(0)、布尔值(false)、切片/映射(len=0)或 nil 时，返回 false；
// 否则返回 true。
func X是否为空(值 interface{}) bool {
	return empty.IsEmpty(值)
}

// Keys 从给定的 map 或 struct 中获取并返回其键（key）。
func X取所有名称(map或Struct interface{}) (键或属性 []string) {
	键或属性 = make([]string, 0)
	if m, ok := map或Struct.(map[string]interface{}); ok {
		for k := range m {
			键或属性 = append(键或属性, k)
		}
		return
	}
	var (
		reflectValue reflect.Value
		reflectKind  reflect.Kind
	)
	if v, ok := map或Struct.(reflect.Value); ok {
		reflectValue = v
	} else {
		reflectValue = reflect.ValueOf(map或Struct)
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
			键或属性 = append(键或属性, gconv.String(k.Interface()))
		}
	case reflect.Struct:
		var (
			fieldType   reflect.StructField
			reflectType = reflectValue.Type()
		)
		for i := 0; i < reflectValue.NumField(); i++ {
			fieldType = reflectType.Field(i)
			if fieldType.Anonymous {
				键或属性 = append(键或属性, X取所有名称(reflectValue.Field(i))...)
			} else {
				键或属性 = append(键或属性, fieldType.Name)
			}
		}
	}
	return
}

// Values 从给定的 map 或 struct 中获取并返回其值。
func X取所有值(map或Struct interface{}) (值数组 []interface{}) {
	值数组 = make([]interface{}, 0)
	if m, ok := map或Struct.(map[string]interface{}); ok {
		for _, v := range m {
			值数组 = append(值数组, v)
		}
		return
	}
	var (
		reflectValue reflect.Value
		reflectKind  reflect.Kind
	)
	if v, ok := map或Struct.(reflect.Value); ok {
		reflectValue = v
	} else {
		reflectValue = reflect.ValueOf(map或Struct)
	}
	reflectKind = reflectValue.Kind()
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	switch reflectKind {
	case reflect.Map:
		for _, k := range reflectValue.MapKeys() {
			值数组 = append(值数组, reflectValue.MapIndex(k).Interface())
		}
	case reflect.Struct:
		var (
			fieldType   reflect.StructField
			reflectType = reflectValue.Type()
		)
		for i := 0; i < reflectValue.NumField(); i++ {
			fieldType = reflectType.Field(i)
			if fieldType.Anonymous {
				值数组 = append(值数组, X取所有值(reflectValue.Field(i))...)
			} else {
				值数组 = append(值数组, reflectValue.Field(i).Interface())
			}
		}
	}
	return
}
