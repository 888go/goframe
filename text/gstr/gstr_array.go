// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstr

// SearchArray 在字符串切片 `a` 中按大小写敏感方式搜索字符串 `s`，
// 并返回其在 `a` 中的索引。
// 若 `s` 未在 `a` 中找到，则返回 -1。
func SearchArray(a []string, s string) int {
	for i, v := range a {
		if s == v {
			return i
		}
	}
	return NotFoundIndex
}

// InArray 检查字符串 `s` 是否在切片 `a` 中。
func InArray(a []string, s string) bool {
	return SearchArray(a, s) != NotFoundIndex
}

// PrefixArray 为 `array` 中的每个元素添加 `prefix` 字符串。
func PrefixArray(array []string, prefix string) {
	for k, v := range array {
		array[k] = prefix + v
	}
}
