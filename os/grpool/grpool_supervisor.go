// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 协程类

import (
	"context"
	
	"github.com/888go/goframe/os/gtimer"
)

// supervisor 负责检查任务列表，如果存在待处理任务且池中无工作进程时，会创建新的 worker 协程来处理该任务。
func (p *Pool) supervisor(_ context.Context) {
	if p.IsClosed() {
		定时类.X退出()
	}
	if p.list.Size() > 0 && p.count.X取值() == 0 {
		var number = p.list.Size()
		if p.limit > 0 {
			number = p.limit
		}
		for i := 0; i < number; i++ {
			p.checkAndForkNewGoroutineWorker()
		}
	}
}
