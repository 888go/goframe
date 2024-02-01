// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package utils

// MapPossibleItemByKey尝试通过忽略大小写和符号的方式，为给定的键查找可能的键值对。
//
// 注意：此函数可能性能较低。
func MapPossibleItemByKey(data map[string]interface{}, key string) (foundKey string, foundValue interface{}) {
	if len(data) == 0 {
		return
	}
	if v, ok := data[key]; ok {
		return key, v
	}
	// Loop checking.
	for k, v := range data {
		if EqualFoldWithoutChars(k, key) {
			return k, v
		}
	}
	return "", nil
}

// MapContainsPossibleKey 检查给定的 `key` 是否包含在给定的映射 `data` 中。
// 它在忽略大小写和符号的情况下检查键值。
//
// 注意，此函数可能性能较低。
func MapContainsPossibleKey(data map[string]interface{}, key string) bool {
	if k, _ := MapPossibleItemByKey(data, key); k != "" {
		return true
	}
	return false
}
