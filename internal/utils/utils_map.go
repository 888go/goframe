// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package utils

// MapPossibleItemByKey 尝试根据给定键（忽略大小写和符号）找到可能的 key-value 对。
//
// 请注意，此函数的性能可能较低。 md5:4dd7c7511eb401cd
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

// MapContainsPossibleKey 检查给定的`key`是否包含在给定的映射`data`中。
// 它在检查键时忽略大小写和符号。
//
// 注意：此函数可能性能较低。 md5:f1b183430304dc85
func MapContainsPossibleKey(data map[string]interface{}, key string) bool {
	if k, _ := MapPossibleItemByKey(data, key); k != "" {
		return true
	}
	return false
}
