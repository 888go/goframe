// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类

// SearchArray 在字符串切片 `a` 中按大小写敏感方式搜索字符串 `s`，
// 并返回其在 `a` 中的索引。
// 若 `s` 未在 `a` 中找到，则返回 -1。
func X数组查找(数组 []string, 待查找值 string) int {
	for i, v := range 数组 {
		if 待查找值 == v {
			return i
		}
	}
	return NotFoundIndex
}

// InArray 检查字符串 `s` 是否在切片 `a` 中。
func X数组是否存在(数组 []string, 待查找值 string) bool {
	return X数组查找(数组, 待查找值) != NotFoundIndex
}

// PrefixArray 为 `array` 中的每个元素添加 `prefix` 字符串。
func X数组加前缀(数组 []string, 前缀 string) {
	for k, v := range 数组 {
		数组[k] = 前缀 + v
	}
}
