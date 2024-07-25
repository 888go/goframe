// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gutil

import (
	"reflect"

	"github.com/gogf/gf/v2/util/gconv"
)

// SliceCopy 对于最常用的切片类型（[]interface{}）执行浅拷贝操作。
// md5:d119afb140e8324a
func SliceCopy(slice []interface{}) []interface{} {
	newSlice := make([]interface{}, len(slice))
	copy(newSlice, slice)
	return newSlice
}

// SliceInsertBefore 将 `values` 插入到 `index` 位置的前面，并返回一个新的切片。 md5:d1cbe1d61df82fb6
func SliceInsertBefore(slice []interface{}, index int, values ...interface{}) (newSlice []interface{}) {
	if index < 0 || index >= len(slice) {
		return slice
	}
	newSlice = make([]interface{}, len(slice)+len(values))
	copy(newSlice, slice[0:index])
	copy(newSlice[index:], values)
	copy(newSlice[index+len(values):], slice[index:])
	return
}

// SliceInsertAfter 在 `index` 后方插入 `values`，并返回一个新的切片。 md5:bb561266d0b2b921
func SliceInsertAfter(slice []interface{}, index int, values ...interface{}) (newSlice []interface{}) {
	if index < 0 || index >= len(slice) {
		return slice
	}
	newSlice = make([]interface{}, len(slice)+len(values))
	copy(newSlice, slice[0:index+1])
	copy(newSlice[index+1:], values)
	copy(newSlice[index+1+len(values):], slice[index+1:])
	return
}

// SliceDelete 从索引 `index` 处删除一个元素，并返回新的切片。如果给定的 `index` 不有效，它不会做任何操作。
// md5:f57a8afe207e8169
func SliceDelete(slice []interface{}, index int) (newSlice []interface{}) {
	if index < 0 || index >= len(slice) {
		return slice
	}
		// 在删除时确定数组边界，以提高删除效率。 md5:bc969ee880edf699
	if index == 0 {
		return slice[1:]
	} else if index == len(slice)-1 {
		return slice[:index]
	}
	// 如果是一个非边界删除，
	// 它将涉及创建一个数组，
	// 那么删除操作效率较低。
	// md5:6a664196d66bc968
	return append(slice[:index], slice[index+1:]...)
}

// SliceToMap 将切片类型变量 `slice` 转换为 `map[string]interface{}` 类型。注意，如果 `slice` 的长度不是偶数，它将返回nil。
// 例如：
// ["K1", "v1", "K2", "v2"] => {"K1": "v1", "K2": "v2"}
// ["K1", "v1", "K2"]       => nil
// md5:f7b4384607d6bfbe
func SliceToMap(slice interface{}) map[string]interface{} {
	var (
		reflectValue = reflect.ValueOf(slice)
		reflectKind  = reflectValue.Kind()
	)
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	switch reflectKind {
	case reflect.Slice, reflect.Array:
		length := reflectValue.Len()
		if length%2 != 0 {
			return nil
		}
		data := make(map[string]interface{})
		for i := 0; i < reflectValue.Len(); i += 2 {
			data[gconv.String(reflectValue.Index(i).Interface())] = reflectValue.Index(i + 1).Interface()
		}
		return data
	}
	return nil
}

// SliceToMapWithColumnAsKey 将切片类型变量 `slice` 转换为 `map[interface{}]interface{}`
// 切片中指定列的值将作为返回映射的键。
// 例如：
// SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K1") => {"v1": {"K1": "v1", "K2": 1}, "v2": {"K1": "v2", "K2": 2}}
// SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K2") => {1: {"K1": "v1", "K2": 1}, 2: {"K1": "v2", "K2": 2}}
// md5:bf178f79b93bd1b2
func SliceToMapWithColumnAsKey(slice interface{}, key interface{}) map[interface{}]interface{} {
	var (
		reflectValue = reflect.ValueOf(slice)
		reflectKind  = reflectValue.Kind()
	)
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	data := make(map[interface{}]interface{})
	switch reflectKind {
	case reflect.Slice, reflect.Array:
		for i := 0; i < reflectValue.Len(); i++ {
			if k, ok := ItemValue(reflectValue.Index(i), key); ok {
				data[k] = reflectValue.Index(i).Interface()
			}
		}
	}
	return data
}
