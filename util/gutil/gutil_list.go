// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gutil
import (
	"reflect"
	
	"github.com/888go/goframe/internal/utils"
	)
// ListItemValues 通过键 `key` 获取并返回所有项（item）结构体或映射中的元素。
// 注意，参数 `list` 应为包含映射或结构体元素的切片类型，否则将返回一个空切片。
//
// 参数 `list` 支持以下类型：
// []map[string]interface{}
// []map[string]子映射
// []struct
// []struct:子结构体
// 注意，只有当提供可选参数 `subKey` 时，子映射/子结构体才有意义。
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

// ItemValue 通过 `key` 参数指定的名称/属性获取并返回其对应的值。
// 参数 `item` 可以是 map/*map/struct/*struct 类型。
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
		// `key`必须为字符串类型。
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
		// `mapKey`必须为字符串类型。
		v := reflectValue.FieldByName(keyValue.String())
		if v.IsValid() {
			found = true
			value = v.Interface()
		}
	}
	return
}

// ListItemValuesUnique 通过键 `key` 获取并返回所有结构体或映射中的唯一元素。
// 注意，参数 `list` 应为包含映射或结构体元素的切片类型，
// 否则将返回一个空切片。
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
				// 使字节切片可比较
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

// ListToMapByKey 将 `list` 转换为一个 map[string]interface{}，其中键由 `key` 指定。
// 注意，项的值可能为 slice 类型。
func ListToMapByKey(list []map[string]interface{}, key string) map[string]interface{} {
	return utils.ListToMapByKey(list, key)
}
