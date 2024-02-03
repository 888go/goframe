// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtime

import (
	"time"
)

// wrapper 是对标准库 struct time.Time 的一个包装器。
// 它用于重写 time.Time 中的一些函数，例如：String。
type wrapper struct {
	time.Time
}

// String 重写 time.Time 的 String 函数。
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
