// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtime

import (
	"time"
)

// wrapper 是标准库 time.Time 结构体的包装器。
// 它用于重写 time.Time 的某些函数，例如：String。
// md5:c8307623baa7274a
type wrapper struct {
	time.Time
}

// String 方法重写了 time.Time 类型的 String 函数。 md5:c519f17d90f0ddcb
func (t wrapper) String() string {
	if t.IsZero() {
		return ""
	}
	if t.Year() == 0 {
		// Only time.
		return t.Format("15:04:05")
	}
	return t.Format("2006-01-02 15:04:05")
}
