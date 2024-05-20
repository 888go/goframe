// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package grpool

import (
	"context"

	"github.com/gogf/gf/v2/os/gtimer"
)

// supervisor 检查工作列表，如果工作列表中有任务但工作池中没有工人goroutine，就 fork 新的工人goroutine来处理任务。
// md5:02b61e26a9994363
func (p *Pool) supervisor(_ context.Context) {
	if p.IsClosed() {
		gtimer.Exit()
	}
	if p.list.Size() > 0 && p.count.Val() == 0 {
		var number = p.list.Size()
		if p.limit > 0 {
			number = p.limit
		}
		for i := 0; i < number; i++ {
			p.checkAndForkNewGoroutineWorker()
		}
	}
}
