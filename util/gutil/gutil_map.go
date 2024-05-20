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

// MapCopy 对于最常用的映射类型map[string]interface{}，执行从数据映射`data`到`copy`的浅复制。
// md5:f29e782d6173a097
func MapCopy(data map[string]interface{}) (copy map[string]interface{}) {
	copy = make(map[string]interface{}, len(data))
	for k, v := range data {
		copy[k] = v
	}
	return
}

// MapContains 检查映射数据 `data` 是否包含键 `key`。. md5:be6a3dd5c1d28037
func MapContains(data map[string]interface{}, key string) (ok bool) {
	if len(data) == 0 {
		return
	}
	_, ok = data[key]
	return
}

// MapDelete 从 map `data` 中删除所有 `keys`。. md5:3e89d2fe52825284
func MapDelete(data map[string]interface{}, keys ...string) {
	if len(data) == 0 {
		return
	}
	for _, key := range keys {
		delete(data, key)
	}
}

// MapMerge 将源（src）映射中的所有映射合并到目标（dst）映射中。. md5:aa4647e55af49733
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

// MapMergeCopy 创建并返回一个新地图，该地图合并了来自 `src` 的所有地图。. md5:b880aa85a1899404
func MapMergeCopy(src ...map[string]interface{}) (copy map[string]interface{}) {
	copy = make(map[string]interface{})
	for _, m := range src {
		for k, v := range m {
			copy[k] = v
		}
	}
	return
}

// MapPossibleItemByKey 尝试根据给定键（忽略大小写和符号）找到可能的 key-value 对。
//
// 请注意，此函数的性能可能较低。
// md5:4dd7c7511eb401cd
func MapPossibleItemByKey(data map[string]interface{}, key string) (foundKey string, foundValue interface{}) {
	return utils.MapPossibleItemByKey(data, key)
}

// MapContainsPossibleKey 检查给定的`key`是否包含在给定的映射`data`中。
// 它在检查键时忽略大小写和符号。
//
// 注意：此函数可能性能较低。
// md5:f1b183430304dc85
func MapContainsPossibleKey(data map[string]interface{}, key string) bool {
	return utils.MapContainsPossibleKey(data, key)
}

// MapOmitEmpty 从给定的映射中删除所有空值。. md5:9e670086c7ddbc0e
func MapOmitEmpty(data map[string]interface{}) {
	if len(data) == 0 {
		return
	}
	for k, v := range data {
		if IsEmpty(v) {
			delete(data, k)
		}
	}
}

// MapToSlice 将映射转换为包含所有键和值的切片。
// 例如：{"K1": "v1", "K2": "v2"} => ["K1", "v1", "K2", "v2"]
// md5:bf557c548edfb4a6
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
