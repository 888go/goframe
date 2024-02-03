// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfpool

import (
	"fmt"
	"os"
	"time"
	
	"github.com/888go/goframe/errors/gerror"
)

// Open 函数通过给定的文件路径、标志和打开权限创建并返回一个文件项。
// 当首次调用时，它会自动内部创建一个相关的文件指针池。
// 此后，它将从该文件指针池中获取文件项。
func Open(path string, flag int, perm os.FileMode, ttl ...time.Duration) (file *File, err error) {
	var fpTTL time.Duration
	if len(ttl) > 0 {
		fpTTL = ttl[0]
	}
// **不要**在这里搜索路径以免浪费性能！
// 保留以下代码只是为了给您一个警告。
//
// path, err = gfile.Search(path)
// if err != nil {
//     return nil, err
// }
	pool := pools.GetOrSetFuncLock(
		fmt.Sprintf("%s&%d&%d&%d", path, flag, fpTTL, perm),
		func() interface{} {
			return New(path, flag, perm, fpTTL)
		},
	).(*Pool)

	return pool.File()
}

// Get 函数通过给定的文件路径、标志和打开权限获取一个文件项。
// 然后，它从文件指针池中检索一个文件项。
func Get(path string, flag int, perm os.FileMode, ttl ...time.Duration) (file *File) {
	var fpTTL time.Duration
	if len(ttl) > 0 {
		fpTTL = ttl[0]
	}

	f, found := pools.Search(fmt.Sprintf("%s&%d&%d&%d", path, flag, fpTTL, perm))
	if !found {
		return nil
	}

	fp, _ := f.(*Pool).pool.Get()
	return fp.(*File)
}

// Stat返回描述文件的FileInfo结构体。
func (f *File) Stat() (os.FileInfo, error) {
	if f.stat == nil {
		return nil, gerror.New("file stat is empty")
	}
	return f.stat, nil
}

// Close将文件指针放回文件指针池。
func (f *File) Close(close ...bool) error {
	if len(close) > 0 && close[0] {
		f.File.Close()
	}

	if f.pid == f.pool.id.Val() {
		return f.pool.pool.Put(f)
	}
	return nil
}
