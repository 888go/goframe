// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类

import (
	"reflect"
	
	"github.com/888go/goframe/util/gconv"
)

// SliceCopy 对于最常用的切片类型 []interface{}，执行浅拷贝操作，复制 `data` 切片。
func SliceCopy(slice []interface{}) []interface{} {
	newSlice := make([]interface{}, len(slice))
	copy(newSlice, slice)
	return newSlice
}

// SliceInsertBefore 将 `values` 插入到 `index` 位置之前，并返回一个新的切片。
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

// SliceInsertAfter 在 `index` 后面插入 `values`，并返回一个新的切片。
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

// SliceDelete 在`index`处删除一个元素并返回新的切片。
// 如果给定的`index`无效，则不做任何操作。
func SliceDelete(slice []interface{}, index int) (newSlice []interface{}) {
	if index < 0 || index >= len(slice) {
		return slice
	}
	// 确定删除时的数组边界以提高删除效率
	if index == 0 {
		return slice[1:]
	} else if index == len(slice)-1 {
		return slice[:index]
	}
// 如果这是一个非边界删除，
// 那么它将涉及创建一个数组，
// 因此，删除操作效率较低。
	return append(slice[:index], slice[index+1:]...)
}

// SliceToMap 将切片类型变量 `slice` 转换为 `map[string]interface{}` 类型。
// 注意，如果 `slice` 的长度不是偶数，则返回 nil。
// 示例：
// ["K1", "v1", "K2", "v2"] => {"K1": "v1", "K2": "v2"}
// ["K1", "v1", "K2"]       => nil
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
			data[转换类.String(reflectValue.Index(i).Interface())] = reflectValue.Index(i + 1).Interface()
		}
		return data
	}
	return nil
}

// SliceToMapWithColumnAsKey 将切片类型变量 `slice` 转换为 `map[interface{}]interface{}`
// 指定列的值作为返回映射中的键。
// 示例：
// SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K1") => {"v1": {"K1": "v1", "K2": 1}, "v2": {"K1": "v2", "K2": 2}}
// SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K2") => {1: {"K1": "v1", "K2": 1}, 2: {"K1": "v2", "K2": 2}}
// 这段Go语言代码注释翻译成中文后，其内容如下：
// SliceToMapWithColumnAsKey 函数将 slice 类型变量转换为 map 类型变量，其中 map 的键是指定列的值。
// 例如：
// 当调用 SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K1") 时，返回结果为 {"v1": {"K1": "v1", "K2": 1}, "v2": {"K1": "v2", "K2": 2}}
// 当调用 SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K2") 时，返回结果为 {1: {"K1": "v1", "K2": 1}, 2: {"K1": "v2", "K2": 2}}
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
