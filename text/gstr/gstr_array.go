// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gstr

// SearchArray searches string `s` in string slice `a` case-sensitively,
// returns its index in `a`.
// If `s` is not found in `a`, it returns -1.
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

// InArray checks whether string `s` in slice `a`.
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
