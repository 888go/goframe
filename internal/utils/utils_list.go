// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package utils

import (
	"fmt"
)

// ListToMapByKey 将 `list` 转换为一个 map[string]interface{}，其中键由 `key` 指定。
// 注意，项的值可能为 slice 类型。
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
