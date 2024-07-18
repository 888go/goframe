// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gutil

import (
	"reflect"

	"github.com/gogf/gf/v2/internal/utils"
)

// ListItemValues 从所有元素为映射或结构体的切片（list）中，根据给定的键（key）获取并返回对应的值。
// 注意，参数 list 应该是包含 map 或结构体元素的切片，否则将返回一个空切片。
// 
// 参数 list 支持的类型包括：
// []map[string]interface{}
// []map[string]子映射
// []struct
// []struct:子结构体
// 如果提供了可选参数 `subKey`，子映射/子结构体才有意义。
// md5:9523dac525318de2
// ff:
// list:
// key:
// subKey:
// values:
func ListItemValues(list interface{}, key interface{}, subKey ...interface{}) (values []interface{}) {
	var reflectValue reflect.Value
	if v, ok := list.(reflect.Value); ok {
		reflectValue = v
	} else {
		reflectValue = reflect.ValueOf(list)
	}
	reflectKind := reflectValue.Kind()
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	switch reflectKind {
	case reflect.Slice, reflect.Array:
		if reflectValue.Len() == 0 {
			return
		}
		values = []interface{}{}
		for i := 0; i < reflectValue.Len(); i++ {
			if value, ok := ItemValue(reflectValue.Index(i), key); ok {
				if len(subKey) > 0 && subKey[0] != nil {
					if subValue, ok := ItemValue(value, subKey[0]); ok {
						value = subValue
					} else {
						continue
					}
				}
				if array, ok := value.([]interface{}); ok {
					values = append(values, array...)
				} else {
					values = append(values, value)
				}
			}
		}
	}
	return
}

// ItemValue 获取并返回由`key`指定的名称/属性的值。
// 参数`item`可以是地图(map)、指针地图(*map)、结构体(struct)或指针结构体(*struct)类型。
// md5:ca5bcda09a11157b
// ff:
// item:
// key:
// value:
// found:
func ItemValue(item interface{}, key interface{}) (value interface{}, found bool) {
	var reflectValue reflect.Value
	if v, ok := item.(reflect.Value); ok {
		reflectValue = v
	} else {
		reflectValue = reflect.ValueOf(item)
	}
	reflectKind := reflectValue.Kind()
	if reflectKind == reflect.Interface {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	var keyValue reflect.Value
	if v, ok := key.(reflect.Value); ok {
		keyValue = v
	} else {
		keyValue = reflect.ValueOf(key)
	}
	switch reflectKind {
	case reflect.Array, reflect.Slice:
		// `key` 必须是字符串类型。 md5:6ffd36d1a5fc0de1
		values := ListItemValues(reflectValue, keyValue.String())
		if values == nil {
			return nil, false
		}
		return values, true

	case reflect.Map:
		v := reflectValue.MapIndex(keyValue)
		if v.IsValid() {
			found = true
			value = v.Interface()
		}

	case reflect.Struct:
		// `mapKey`必须是字符串类型。 md5:d2b6db36f99feed4
		v := reflectValue.FieldByName(keyValue.String())
		if v.IsValid() {
			found = true
			value = v.Interface()
		}
	}
	return
}

// ListItemValuesUnique 获取并返回具有键为`key`的所有结构体/映射的独特元素。
// 请注意，参数`list`应为包含映射或结构体元素的切片类型，否则将返回一个空切片。
// md5:0f361d3ff901d0a1
// ff:
// list:
// key:
// subKey:
func ListItemValuesUnique(list interface{}, key string, subKey ...interface{}) []interface{} {
	values := ListItemValues(list, key, subKey...)
	if len(values) > 0 {
		var (
			ok bool
			m  = make(map[interface{}]struct{}, len(values))
		)
		for i := 0; i < len(values); {
			value := values[i]
			if t, ok := value.([]byte); ok {
				// 将字节切片设置为可比较. md5:1ca9d6fe5290f517
				value = string(t)
			}
			if _, ok = m[value]; ok {
				values = SliceDelete(values, i)
			} else {
				m[value] = struct{}{}
				i++
			}
		}
	}
	return values
}

// ListToMapByKey 将 `list` 转换为一个键为 `key` 的 map[string]interface{}。注意，项的值可能为切片类型。
// md5:6509753e629d5dc6
// ff:
// list:
// key:
func ListToMapByKey(list []map[string]interface{}, key string) map[string]interface{} {
	return utils.ListToMapByKey(list, key)
}
