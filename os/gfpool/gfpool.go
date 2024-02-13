// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gfpool 提供了用于文件指针的可重用 io 资源池。
package 文件指针池类

import (
	"os"
	"time"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/container/gpool"
	"github.com/888go/goframe/container/gtype"
)

// Pool 指针池。
type Pool struct {
	id   *安全变量类.Int    // Pool id，用于标记此pool是否已重建。
	pool *对象复用类.Pool   // Underlying pool.
	init *安全变量类.Bool   // 是否已初始化，用于标记该文件是否已添加到fsnotify，且只能添加一次。
	ttl  time.Duration // 文件指针项的生存时间（TTL）
}

// File 是池中的一个项目。
type File struct {
	*os.File             // 底层文件指针。
	stat     os.FileInfo // 当前文件指针的状态。
	pid      int         // Belonging pool id, which is set when file pointer created. It's used to check whether the pool is recreated.
	pool     *Pool       // Belonging ool.
	flag     int         // 用于打开文件的闪存（快速暂存）
	perm     os.FileMode // 打开文件所需的权限。
	path     string      // 文件的绝对路径。
}

var (
	// 全局文件指针池。
	pools = map类.X创建StrAny(true)
)
