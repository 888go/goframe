// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// fileinfo 包为给定的信息提供虚拟的 os.FileInfo。
package fileinfo

import (
	"os"
	"time"
)

type Info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func New(name string, size int64, mode os.FileMode, modTime time.Time) *Info {
	return &Info{
		name:    name,
		size:    size,
		mode:    mode,
		modTime: modTime,
	}
}

func (i *Info) Name() string {
	return i.name
}

func (i *Info) Size() int64 {
	return i.size
}

func (i *Info) IsDir() bool {
	return i.mode.IsDir()
}

func (i *Info) Mode() os.FileMode {
	return i.mode
}

func (i *Info) ModTime() time.Time {
	return i.modTime
}

func (i *Info) Sys() interface{} {
	return nil
}
