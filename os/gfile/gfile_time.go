// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gfile

import (
	"os"
	"time"
)

// MTime 返回由`path`指定的文件的修改时间（以秒为单位）。. md5:66dbc182c71f7ffb
func MTime(path string) time.Time {
	s, e := os.Stat(path)
	if e != nil {
		return time.Time{}
	}
	return s.ModTime()
}

// MTimestamp 返回给定路径`path`的文件修改时间，以秒为单位。. md5:bb848f3c89f3cb71
func MTimestamp(path string) int64 {
	mtime := MTime(path)
	if mtime.IsZero() {
		return -1
	}
	return mtime.Unix()
}

// MTimestampMilli 返回由 `path` 指定的文件的修改时间（以毫秒为单位）。. md5:2def39248c3bde9b
func MTimestampMilli(path string) int64 {
	mtime := MTime(path)
	if mtime.IsZero() {
		return -1
	}
	return mtime.UnixNano() / 1000000
}
