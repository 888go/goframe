// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gfpool

import (
	"fmt"
	"os"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
)

// Open 函数根据给定的文件路径、标志和打开权限创建并返回一个文件项。当它首次被调用时，它会自动内部创建一个关联的文件指针池。然后，它从文件指针池中获取文件项。
// md5:94bbe2b7d15d2c1f
func Open(path string, flag int, perm os.FileMode, ttl ...time.Duration) (file *File, err error) {
	var fpTTL time.Duration
	if len(ttl) > 0 {
		fpTTL = ttl[0]
	}
	// 不要在這裡浪費性能搜索路径！
	// 保留以下代码只是为了警告你。
	//
	// path, err := gfile.Search(path)
	// 如果 err 不为 nil，则：
	// 返回 nil 和 err
	// md5:763fc7901895ec01
	pool := pools.GetOrSetFuncLock(
		fmt.Sprintf("%s&%d&%d&%d", path, flag, fpTTL, perm),
		func() interface{} {
			return New(path, flag, perm, fpTTL)
		},
	).(*Pool)

	return pool.File()
}

// Get 根据给定的文件路径、标志和打开权限返回一个文件项。
// 随后，它从文件指针池中检索一个文件项。
// md5:f56943d16a070df7
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

// Stat 返回描述文件的FileInfo结构。 md5:86e6f3f0a508aa53
func (f *File) Stat() (os.FileInfo, error) {
	if f.stat == nil {
		return nil, gerror.New("file stat is empty")
	}
	return f.stat, nil
}

// Close 将文件指针放回文件指针池。 md5:a47bacf277b7f774
func (f *File) Close(close ...bool) error {
	if len(close) > 0 && close[0] {
		f.File.Close()
	}

	if f.pid == f.pool.id.Val() {
		return f.pool.pool.Put(f)
	}
	return nil
}
