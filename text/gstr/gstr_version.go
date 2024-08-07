// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文本类

import (
	"strings"

	gconv "github.com/888go/goframe/util/gconv"
)

// X版本号是否有效 检查并返回给定的 `version` 是否为有效的 GNU 版本字符串。 md5:6400dc6a399e4aa3
func X版本号是否有效(版本号 string) bool {
	if 版本号 != "" && (版本号[0] == 'v' || 版本号[0] == 'V') {
		版本号 = 版本号[1:]
	}
	if 版本号 == "" {
		return false
	}
	var array = strings.Split(版本号, ".")
	if len(array) > 3 {
		return false
	}
	for _, v := range array {
		if v == "" {
			return false
		}
		if !X是否为数字(v) {
			return false
		}
		if v[0] == '-' || v[0] == '+' {
			return false
		}
	}
	return true
}

// X版本号比较GNU格式 按照GNU版本标准比较 `a` 和 `b`。
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
// 等等。
// md5:2716e579b3f9ba4d
func X版本号比较GNU格式(a, b string) int {
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
		v1 = gconv.X取整数(array1[i])
		v2 = gconv.X取整数(array2[i])
		if v1 > v2 {
			return 1
		}
		if v1 < v2 {
			return -1
		}
	}
	return 0
}

// X版本号比较GO格式 将 `a` 和 `b` 当作标准的 Go 语言版本进行比较。
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
// 文档：https://go.dev/doc/modules/version-numbers
// md5:27f202ad306995b3
func X版本号比较GO格式(a, b string) int {
	a = X过滤首尾符并含空白(a)
	b = X过滤首尾符并含空白(b)
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
	if X统计次数(a, "-") > 1 {
		if i := X倒找(a, "-"); i > 0 {
			a = a[:i]
		}
	}
	if X统计次数(b, "-") > 1 {
		if i := X倒找(b, "-"); i > 0 {
			b = b[:i]
		}
	}
	if i := X查找(a, "+"); i > 0 {
		a = a[:i]
	}
	if i := X查找(b, "+"); i > 0 {
		b = b[:i]
	}
	a = X替换(a, "-", ".")
	b = X替换(b, "-", ".")
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
		v1, v2 = gconv.X取整数(array1[i]), gconv.X取整数(array2[i])
		// 特别是在Go语言中：
		// 特别是：
		// "v1.12.2-0.20200413154443-b17e3a6804fa" < "v1.12.2" 		// 表示 v1.12.2-0.20200413154443-b17e3a6804fa 版本早于 v1.12.2
		// "v1.12.3-0.20200413154443-b17e3a6804fa" > "v1.12.2" 		// 表示 v1.12.3-0.20200413154443-b17e3a6804fa 版本晚于 v1.12.2
		// md5:685fe05f97473463
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
	// "v4.20.1+incompatible" 小于 "v4.20.1"
	// md5:a292bd03375fd35c
	inA, inB := X是否包含(rawA, "+incompatible"), X是否包含(rawB, "+incompatible")
	if inA && !inB {
		return -1
	}
	if !inA && inB {
		return 1
	}

	return 0
}
