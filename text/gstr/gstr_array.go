// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gstr

// SearchArray 在字符串切片 `a` 中大小写敏感地搜索字符串 `s`,
// 并返回 `s` 在 `a` 中的索引。
// 如果 `s` 在 `a` 中未找到，它将返回 -1。
// md5:1f90173b093235c0
// ff:切片查找
// a:切片
// s:待查找值
func SearchArray(a []string, s string) int {
	for i, v := range a {
		if s == v {
			return i
		}
	}
	return NotFoundIndex
}

// InArray 检查字符串 `s` 是否存在于切片 `a` 中。 md5:a155970bbf0f5679
// ff:切片是否存在
// a:切片
// s:待查找值
func InArray(a []string, s string) bool {
	return SearchArray(a, s) != NotFoundIndex
}

// PrefixArray adds `prefix` string for each item of `array`.
//
// PrefixArray(["a","b"], "gf_") -> ["gf_a", "gf_b"]
// ff:切片加前缀
// array:切片
// prefix:前缀
func PrefixArray(array []string, prefix string) {
	for k, v := range array {
		array[k] = prefix + v
	}
}
