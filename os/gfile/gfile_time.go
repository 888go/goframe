// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件类

import (
	"os"
	"time"
)

// X取修改时间秒 返回由`path`指定的文件的修改时间（以秒为单位）。 md5:66dbc182c71f7ffb
func X取修改时间秒(路径 string) time.Time {
	s, e := os.Stat(路径)
	if e != nil {
		return time.Time{}
	}
	return s.ModTime()
}

// X取修改时间戳秒 返回给定路径`path`的文件修改时间，以秒为单位。 md5:bb848f3c89f3cb71
func X取修改时间戳秒(路径 string) int64 {
	mtime := X取修改时间秒(路径)
	if mtime.IsZero() {
		return -1
	}
	return mtime.Unix()
}

// X取修改时间戳毫秒 返回由 `path` 指定的文件的修改时间（以毫秒为单位）。 md5:2def39248c3bde9b
func X取修改时间戳毫秒(路径 string) int64 {
	mtime := X取修改时间秒(路径)
	if mtime.IsZero() {
		return -1
	}
	return mtime.UnixNano() / 1000000
}
