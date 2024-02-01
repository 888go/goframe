// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile
import (
	"os"
	"time"
	)
// MTime 返回给定路径 `path` 下文件的修改时间（以秒为单位）。
func MTime(path string) time.Time {
	s, e := os.Stat(path)
	if e != nil {
		return time.Time{}
	}
	return s.ModTime()
}

// MTimestamp 返回给定路径 `path` 文件的修改时间（以秒为单位）。
func MTimestamp(path string) int64 {
	mtime := MTime(path)
	if mtime.IsZero() {
		return -1
	}
	return mtime.Unix()
}

// MTimestampMilli 返回给定路径 `path` 下文件的修改时间，单位为毫秒。
func MTimestampMilli(path string) int64 {
	mtime := MTime(path)
	if mtime.IsZero() {
		return -1
	}
	return mtime.UnixNano() / 1000000
}
