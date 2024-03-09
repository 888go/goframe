// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类

import (
	"reflect"
	
	"github.com/888go/goframe/gutil/internal/utils"
)

// MapCopy 对于最常用的 map 类型 map[string]interface{}，执行从 `data` 到 `copy` 的浅复制。
func MapCopy(data map[string]interface{}) (copy map[string]interface{}) {
	copy = make(map[string]interface{}, len(data))
	for k, v := range data {
		copy[k] = v
	}
	return
}

// MapContains 检查字典 `data` 是否包含键 `key`。
func MapContains(data map[string]interface{}, key string) (ok bool) {
	if len(data) == 0 {
		return
	}
	_, ok = data[key]
	return
}

// MapDelete 从 map `data` 中删除所有 `keys`。
func MapDelete(data map[string]interface{}, keys ...string) {
	if len(data) == 0 {
		return
	}
	for _, key := range keys {
		delete(data, key)
	}
}

// MapMerge 将从 `src` 中的所有映射合并到映射 `dst`。
func MapMerge(dst map[string]interface{}, src ...map[string]interface{}) {
	if dst == nil {
		return
	}
	for _, m := range src {
		for k, v := range m {
			dst[k] = v
		}
	}
}

// MapMergeCopy 创建并返回一个新的映射，该映射将合并来自 `src` 中的所有映射。
func MapMergeCopy(src ...map[string]interface{}) (copy map[string]interface{}) {
	copy = make(map[string]interface{})
	for _, m := range src {
		for k, v := range m {
			copy[k] = v
		}
	}
	return
}

// MapPossibleItemByKey尝试通过忽略大小写和符号的方式，为给定的键查找可能的键值对。
//
// 注意：此函数可能性能较低。
func MapPossibleItemByKey(data map[string]interface{}, key string) (foundKey string, foundValue interface{}) {
	return utils.MapPossibleItemByKey(data, key)
}

// MapContainsPossibleKey 检查给定的 `key` 是否包含在给定的映射 `data` 中。
// 它在忽略大小写和符号的情况下检查键值。
//
// 注意，此函数可能性能较低。
func MapContainsPossibleKey(data map[string]interface{}, key string) bool {
	return utils.MapContainsPossibleKey(data, key)
}

// MapOmitEmpty 从给定的映射中删除所有空值。
func MapOmitEmpty(data map[string]interface{}) {
	if len(data) == 0 {
		return
	}
	for k, v := range data {
		if X是否为空(v) {
			delete(data, k)
		}
	}
}

// MapToSlice 将映射转换为切片，其中所有键和值都是其元素。
// 例如：{"K1": "v1", "K2": "v2"} => ["K1", "v1", "K2", "v2"]
func MapToSlice(data interface{}) []interface{} {
	var (
		reflectValue = reflect.ValueOf(data)
		reflectKind  = reflectValue.Kind()
	)
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	switch reflectKind {
	case reflect.Map:
		array := make([]interface{}, 0)
		for _, key := range reflectValue.MapKeys() {
			array = append(array, key.Interface())
			array = append(array, reflectValue.MapIndex(key).Interface())
		}
		return array
	}
	return nil
}
