// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// package fileinfo 提供给定信息的虚拟os.FileInfo。. md5:cbdfc1b24e190cd9
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
