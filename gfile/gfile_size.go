// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Size 返回指定路径 `path` 文件的大小，单位为字节。
func X取大小(路径 string) int64 {
	s, e := os.Stat(路径)
	if e != nil {
		return 0
	}
	return s.Size()
}

// SizeFormat 返回指定路径 `path` 下文件的大小，格式化为字符串。
func X取大小并易读格式(路径 string) string {
	return X字节长度转易读格式(X取大小(路径))
}

// ReadableSize 格式化给定路径 `path` 的文件大小，使其更易于人类阅读。
func ReadableSize别名(路径 string) string {
	return X字节长度转易读格式(X取大小(路径))
}

// StrToSize 将格式化的大小字符串转换为其字节表示的大小。
func X易读格式转字节长度(大小文本 string) int64 {
	i := 0
	for ; i < len(大小文本); i++ {
		if 大小文本[i] == '.' || (大小文本[i] >= '0' && 大小文本[i] <= '9') {
			continue
		} else {
			break
		}
	}
	var (
		unit      = 大小文本[i:]
		number, _ = strconv.ParseFloat(大小文本[:i], 64)
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

// FormatSize 将大小 `raw` 格式化为更便于人工阅读的形式。
func X字节长度转易读格式(文件大小 int64) string {
	var r float64 = float64(文件大小)
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
