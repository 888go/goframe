// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package utils

import "fmt"

// ListToMapByKey 将 `list` 转换为一个键为 `key` 的 map[string]interface{}。注意，项的值可能为切片类型。
// md5:6509753e629d5dc6
func ListToMapByKey(list []map[string]interface{}, key string) map[string]interface{} {
	var (
		s              = ""
		m              = make(map[string]interface{})
		tempMap        = make(map[string][]interface{})
		hasMultiValues bool
	)
	for _, item := range list {
		if k, ok := item[key]; ok {
			s = fmt.Sprintf(`%v`, k)
			tempMap[s] = append(tempMap[s], item)
			if len(tempMap[s]) > 1 {
				hasMultiValues = true
			}
		}
	}
	for k, v := range tempMap {
		if hasMultiValues {
			m[k] = v
		} else {
			m[k] = v[0]
		}
	}
	return m
}
