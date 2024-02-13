// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类

import (
	"strings"
	
	"github.com/888go/goframe/util/gconv"
)

// IsGNUVersion 检查并返回给定的 `version` 是否为有效的 GNU 版本字符串。
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

// CompareVersion 按照标准 GNU 版本格式比较 `a` 和 `b`。
//
// 如果 `a` > `b`，则返回 1。
//
// 如果 `a` < `b`，则返回 -1。
//
// 如果 `a` = `b`，则返回 0。
//
// 标准 GNU 版本格式例如：
// v1.0
// 1
// 1.0.0
// v1.0.1
// v2.10.8
// 10.2.0
// 等等。
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
		v1 = 转换类.X取整数(array1[i])
		v2 = 转换类.X取整数(array2[i])
		if v1 > v2 {
			return 1
		}
		if v1 < v2 {
			return -1
		}
	}
	return 0
}

// CompareVersionGo 按照 Golang 标准版本格式比较 `a` 和 `b`。
//
// 如果 `a` > `b`，则返回 1。
//
// 如果 `a` < `b`，则返回 -1。
//
// 如果 `a` = `b`，则返回 0。
//
// Golang 标准版本格式例如：
// 1.0.0
// v1.0.1
// v2.10.8
// 10.2.0
// v0.0.0-20190626092158-b2ccc519800e
// v1.12.2-0.20200413154443-b17e3a6804fa
// v4.20.0+incompatible
// 等等。
//
// 文档参考：https://go.dev/doc/modules/version-numbers
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

	// 首先检查 Major.Minor.Patch
	v1, v2 := 0, 0
	for i := 0; i < len(array1); i++ {
		v1, v2 = 转换类.X取整数(array1[i]), 转换类.X取整数(array2[i])
// 特别地在 Golang 中：
// "v1.12.2-0.20200413154443-b17e3a6804fa" 小于 "v1.12.2"
// "v1.12.3-0.20200413154443-b17e3a6804fa" 大于 "v1.12.2"
// 这段代码注释是关于 Golang 中版本字符串比较的特殊规则：
// 在 Golang 中，对于包含预发布版本号（如 "-0.20200413154443-b17e3a6804fa"）的版本字符串，在进行字符串比较时，主版本号、次版本号和补丁版本号部分会被优先比较。当这部分相同时，带有预发布版本号的版本会认为小于不带预发布版本号的版本。
// 因此，尽管 "v1.12.2-0.20200413154443-b17e3a6804fa" 的主要部分与 "v1.12.2" 相同，但由于其附加了预发布标识，所以在比较中它被认为小于 "v1.12.2"。
// 同样，"v1.12.3-0.20200413154443-b17e3a6804fa" 由于其主版本号部分高于 "v1.12.2"，所以即使它也有预发布版本号，依然会在比较中大于 "v1.12.2"。
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

// 特别在 Golang 中：
// "v4.20.1+incompatible" < "v4.20.1"
// 这表示包含 "+incompatible" 后缀的版本号在比较时，会被视为小于不包含此后缀的相同主版本号和次版本号的版本。这是 Go 语言处理依赖版本的一种方式，"+incompatible" 后缀通常用于标记非模块化的包或不符合语义化版本控制规范的包。
	inA, inB := X是否包含(rawA, "+incompatible"), X是否包含(rawB, "+incompatible")
	if inA && !inB {
		return -1
	}
	if !inA && inB {
		return 1
	}

	return 0
}
