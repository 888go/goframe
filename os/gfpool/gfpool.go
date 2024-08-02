// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Package gfpool 提供可重用的文件指针io池。 md5:994211336c178d55
package 文件指针池类

import (
	"os"
	"time"

	gmap "github.com/888go/goframe/container/gmap"
	gpool "github.com/888go/goframe/container/gpool"
	gtype "github.com/888go/goframe/container/gtype"
)

// Pool pointer pool.
type Pool struct {
	id   *gtype.Int    // 池ID，用于标记此池是否已重新创建。 md5:c0eec2e36167c140
	pool *gpool.Pool   // Underlying pool.
	init *gtype.Bool   // 是否已初始化，用于标记该文件是否已添加到fsnotify，且只能添加一次。 md5:d40a9cbe1c0bc4a8
	ttl  time.Duration // 文件指针项的生存时间。 md5:d52a45f685325e0d
}

// File是池中的一个项目。 md5:eea9020d954d8acb
type File struct {
	*os.File             // 底层文件指针。 md5:c0f18bf647334bf6
	stat     os.FileInfo // 当前文件指针的状态。 md5:60200c5416353351
	pid      int         // Belonging pool id, which is set when file pointer created. It's used to check whether the pool is recreated.
	pool     *Pool       // Belonging ool.
	flag     int         // 打开文件时使用的闪存。 md5:5051cbf609bc85ed
	perm     os.FileMode // 打开文件的权限。 md5:029a045bad083703
	path     string      // 文件的绝对路径。 md5:6e81d00aa7e09eca
}

var (
		// 全局文件指针池。 md5:4db4f1674b75c6e1
	pools = gmap.NewStrAnyMap(true)
)
