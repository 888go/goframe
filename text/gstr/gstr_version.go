// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gstr

import (
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
)

// IsGNUVersion 检查并返回给定的 `version` 是否为有效的 GNU 版本字符串。 md5:6400dc6a399e4aa3
func IsGNUVersion(version string) bool {
	if version != "" && (version[0] == 'v' || version[0] == 'V') {
		version = version[1:]
	}
	if version == "" {
		return false
	}
	var array = strings.Split(version, ".")
	if len(array) > 3 {
		return false
	}
	for _, v := range array {
		if v == "" {
			return false
		}
		if !IsNumeric(v) {
			return false
		}
		if v[0] == '-' || v[0] == '+' {
			return false
		}
	}
	return true
}

// CompareVersion 按照GNU版本标准比较 `a` 和 `b`。
//
// 如果 `a` 大于 `b`，返回 1。
//
// 如果 `a` 小于 `b`，返回 -1。
//
// 如果 `a` 等于 `b`，返回 0。
//
// GNU版本标准格式例如：
// v1.0
// 1
// 1.0.0
// v1.0.1
// v2.10.8
// 10.2.0
// 等等。 md5:2716e579b3f9ba4d
func CompareVersion(a, b string) int {
	if a != "" && a[0] == 'v' {
		a = a[1:]
	}
	if b != "" && b[0] == 'v' {
		b = b[1:]
	}
	var (
		array1 = strings.Split(a, ".")
		array2 = strings.Split(b, ".")
		diff   int
	)
	diff = len(array2) - len(array1)
	for i := 0; i < diff; i++ {
		array1 = append(array1, "0")
	}
	diff = len(array1) - len(array2)
	for i := 0; i < diff; i++ {
		array2 = append(array2, "0")
	}
	v1 := 0
	v2 := 0
	for i := 0; i < len(array1); i++ {
		v1 = gconv.Int(array1[i])
		v2 = gconv.Int(array2[i])
		if v1 > v2 {
			return 1
		}
		if v1 < v2 {
			return -1
		}
	}
	return 0
}

// CompareVersionGo 将 `a` 和 `b` 当作标准的 Go 语言版本进行比较。
//
// 如果 `a` 大于 `b`，返回 1。
//
// 如果 `a` 小于 `b`，返回 -1。
//
// 如果 `a` 等于 `b`，返回 0。
//
// 标准的 Go 语言版本格式如下：
// 1.0.0
// v1.0.1
// v2.10.8
// 10.2.0
// v0.0.0-20190626092158-b2ccc519800e
// v1.12.2-0.20200413154443-b17e3a6804fa
// v4.20.0+incompatible
// 等等。
//
// 文档：https://go.dev/doc/modules/version-numbers md5:27f202ad306995b3
func CompareVersionGo(a, b string) int {
	a = Trim(a)
	b = Trim(b)
	if a != "" && a[0] == 'v' {
		a = a[1:]
	}
	if b != "" && b[0] == 'v' {
		b = b[1:]
	}
	var (
		rawA = a
		rawB = b
	)
	if Count(a, "-") > 1 {
		if i := PosR(a, "-"); i > 0 {
			a = a[:i]
		}
	}
	if Count(b, "-") > 1 {
		if i := PosR(b, "-"); i > 0 {
			b = b[:i]
		}
	}
	if i := Pos(a, "+"); i > 0 {
		a = a[:i]
	}
	if i := Pos(b, "+"); i > 0 {
		b = b[:i]
	}
	a = Replace(a, "-", ".")
	b = Replace(b, "-", ".")
	var (
		array1 = strings.Split(a, ".")
		array2 = strings.Split(b, ".")
		diff   = len(array1) - len(array2)
	)

	for i := diff; i < 0; i++ {
		array1 = append(array1, "0")
	}
	for i := 0; i < diff; i++ {
		array2 = append(array2, "0")
	}

	// 首先检查Major.Minor.Patch. md5:098a0c10a68fabae
	v1, v2 := 0, 0
	for i := 0; i < len(array1); i++ {
		v1, v2 = gconv.Int(array1[i]), gconv.Int(array2[i])
		// 特别是在Go语言中：
		// 特别是：
		// "v1.12.2-0.20200413154443-b17e3a6804fa" < "v1.12.2" 		// 表示 v1.12.2-0.20200413154443-b17e3a6804fa 版本早于 v1.12.2
		// "v1.12.3-0.20200413154443-b17e3a6804fa" > "v1.12.2" 		// 表示 v1.12.3-0.20200413154443-b17e3a6804fa 版本晚于 v1.12.2 md5:685fe05f97473463
		if i == 4 && v1 != v2 && (v1 == 0 || v2 == 0) {
			if v1 > v2 {
				return -1
			} else {
				return 1
			}
		}

		if v1 > v2 {
			return 1
		}
		if v1 < v2 {
			return -1
		}
	}

	// 特别是在 Golang 中：
	// "v4.20.1+incompatible" 小于 "v4.20.1" md5:a292bd03375fd35c
	inA, inB := Contains(rawA, "+incompatible"), Contains(rawB, "+incompatible")
	if inA && !inB {
		return -1
	}
	if !inA && inB {
		return 1
	}

	return 0
}
