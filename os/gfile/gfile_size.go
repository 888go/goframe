// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件类

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Size 返回指定路径`path`的文件大小，以字节为单位。 md5:bb6ab734af5c941f
func Size(path string) int64 {
	s, e := os.Stat(path)
	if e != nil {
		return 0
	}
	return s.Size()
}

// SizeFormat 返回由 `path` 指定的文件的大小，格式为字符串。 md5:17a4b324cd8ec9a7
func SizeFormat(path string) string {
	return FormatSize(Size(path))
}

// ReadableSize 格式化给定路径（`path`）的文件大小，使其更易于人类阅读。 md5:04f6ed7d21e25298
func ReadableSize(path string) string {
	return FormatSize(Size(path))
}

// StrToSize 将格式化后的尺寸字符串转换为字节大小。 md5:b40fff33ad10f088
func StrToSize(sizeStr string) int64 {
	i := 0
	for ; i < len(sizeStr); i++ {
		if sizeStr[i] == '.' || (sizeStr[i] >= '0' && sizeStr[i] <= '9') {
			continue
		} else {
			break
		}
	}
	var (
		unit      = sizeStr[i:]
		number, _ = strconv.ParseFloat(sizeStr[:i], 64)
	)
	if unit == "" {
		return int64(number)
	}
	switch strings.ToLower(unit) {
	case "b", "bytes":
		return int64(number)
	case "k", "kb", "ki", "kib", "kilobyte":
		return int64(number * 1024)
	case "m", "mb", "mi", "mib", "mebibyte":
		return int64(number * 1024 * 1024)
	case "g", "gb", "gi", "gib", "gigabyte":
		return int64(number * 1024 * 1024 * 1024)
	case "t", "tb", "ti", "tib", "terabyte":
		return int64(number * 1024 * 1024 * 1024 * 1024)
	case "p", "pb", "pi", "pib", "petabyte":
		return int64(number * 1024 * 1024 * 1024 * 1024 * 1024)
	case "e", "eb", "ei", "eib", "exabyte":
		return int64(number * 1024 * 1024 * 1024 * 1024 * 1024 * 1024)
	case "z", "zb", "zi", "zib", "zettabyte":
		return int64(number * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024)
	case "y", "yb", "yi", "yib", "yottabyte":
		return int64(number * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024)
	case "bb", "brontobyte":
		return int64(number * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024)
	}
	return -1
}

// FormatSize 将大小 `raw` 格式化为更易读的形式。 md5:7b0eb2b4570b5bb2
func FormatSize(raw int64) string {
	var r float64 = float64(raw)
	var t float64 = 1024
	var d float64 = 1
	if r < t {
		return fmt.Sprintf("%.2fB", r/d)
	}
	d *= 1024
	t *= 1024
	if r < t {
		return fmt.Sprintf("%.2fK", r/d)
	}
	d *= 1024
	t *= 1024
	if r < t {
		return fmt.Sprintf("%.2fM", r/d)
	}
	d *= 1024
	t *= 1024
	if r < t {
		return fmt.Sprintf("%.2fG", r/d)
	}
	d *= 1024
	t *= 1024
	if r < t {
		return fmt.Sprintf("%.2fT", r/d)
	}
	d *= 1024
	t *= 1024
	if r < t {
		return fmt.Sprintf("%.2fP", r/d)
	}
	d *= 1024
	t *= 1024
	if r < t {
		return fmt.Sprintf("%.2fE", r/d)
	}
	d *= 1024
	t *= 1024
	if r < t {
		return fmt.Sprintf("%.2fZ", r/d)
	}
	d *= 1024
	t *= 1024
	if r < t {
		return fmt.Sprintf("%.2fY", r/d)
	}
	d *= 1024
	t *= 1024
	if r < t {
		return fmt.Sprintf("%.2fBB", r/d)
	}
	return "TooLarge"
}
