// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文本类

// X切片查找 在字符串切片 `a` 中大小写敏感地搜索字符串 `s`,
// 并返回 `s` 在 `a` 中的索引。
// 如果 `s` 在 `a` 中未找到，它将返回 -1。
// md5:1f90173b093235c0
func X切片查找(切片 []string, 待查找值 string) int {
	for i, v := range 切片 {
		if 待查找值 == v {
			return i
		}
	}
	return NotFoundIndex
}

// X切片是否存在 检查字符串 `s` 是否存在于切片 `a` 中。 md5:a155970bbf0f5679
func X切片是否存在(切片 []string, 待查找值 string) bool {
	return X切片查找(切片, 待查找值) != NotFoundIndex
}

// X切片加前缀 为数组 `array` 中的每个元素添加前缀字符串 `prefix`。
//
// 示例：
// X切片加前缀(["a", "b"], "gf_") -> ["gf_a", "gf_b"]
// md5:b1320521881ddb0c
func X切片加前缀(切片 []string, 前缀 string) {
	for k, v := range 切片 {
		切片[k] = 前缀 + v
	}
}
