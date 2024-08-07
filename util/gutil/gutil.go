// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gutil提供了实用函数。 md5:94ecbd62b7aa8e76
package 工具类

import (
	"reflect"

	gconv "github.com/888go/goframe/util/gconv"
)

const (
	dumpIndent = `    `
)

// X取所有名称 从给定的映射或结构体中检索并返回键。 md5:90c2f70b36eaef9e
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

// X取所有值 从给定的映射或结构体中检索并返回值。 md5:e03f3d848bc2ef9a
func X取所有值(map或Struct interface{}) (值切片 []interface{}) {
	值切片 = make([]interface{}, 0)
	if m, ok := map或Struct.(map[string]interface{}); ok {
		for _, v := range m {
			值切片 = append(值切片, v)
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
			值切片 = append(值切片, reflectValue.MapIndex(k).Interface())
		}
	case reflect.Struct:
		var (
			fieldType   reflect.StructField
			reflectType = reflectValue.Type()
		)
		for i := 0; i < reflectValue.NumField(); i++ {
			fieldType = reflectType.Field(i)
			if fieldType.Anonymous {
				值切片 = append(值切片, X取所有值(reflectValue.Field(i))...)
			} else {
				值切片 = append(值切片, reflectValue.Field(i).Interface())
			}
		}
	}
	return
}
