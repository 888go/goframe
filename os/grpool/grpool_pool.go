// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package grpool
import (
	"context"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	)
// Add 向任务池中添加一个新的任务。
// 该任务将会被异步执行。
func (p *Pool) Add(ctx context.Context, f Func) error {
	for p.closed.Val() {
		return gerror.NewCode(
			gcode.CodeInvalidOperation,
			"goroutine defaultPool is already closed",
		)
	}
	p.list.PushFront(&localPoolItem{
		Ctx:  ctx,
		Func: f,
	})
	// 检查并创建新的工作进程。
	p.checkAndForkNewGoroutineWorker()
	return nil
}

// AddWithRecover 向池中添加一个带有指定恢复函数的新任务。
//
// 可选的 `recoverFunc` 参数会在执行 `userFunc` 函数期间发生 panic 时被调用。
// 如果未传递 `recoverFunc` 或其值为 nil，则会忽略来自 `userFunc` 的 panic。
// 该任务将被异步执行。
func (p *Pool) AddWithRecover(ctx context.Context, userFunc Func, recoverFunc RecoverFunc) error {
	return p.Add(ctx, func(ctx context.Context) {
		defer func() {
			if exception := recover(); exception != nil {
				if recoverFunc != nil {
					if v, ok := exception.(error); ok && gerror.HasStack(v) {
						recoverFunc(ctx, v)
					} else {
						recoverFunc(ctx, gerror.NewCodef(gcode.CodeInternalPanic, "%+v", exception))
					}
				}
			}
		}()
		userFunc(ctx)
	})
}

// Cap 返回池的容量。
// 这个容量在创建池时定义。
// 如果没有限制，则返回 -1。
func (p *Pool) Cap() int {
	return p.limit
}

// Size 返回当前 goroutine 池中的goroutine数量。
func (p *Pool) Size() int {
	return p.count.Val()
}

// Jobs 返回当前工作池中的任务数量。
// 注意，它返回的是任务/任务计数，而不是工作者/协程计数。
func (p *Pool) Jobs() int {
	return p.list.Size()
}

// IsClosed 返回 pool 是否已关闭。
func (p *Pool) IsClosed() bool {
	return p.closed.Val()
}

// Close 关闭 goroutine 池，这将使所有 goroutine 退出。
func (p *Pool) Close() {
	p.closed.Set(true)
}

// checkAndForkNewGoroutineWorker 检查并创建一个新的 goroutine 工作者。
// 注意，如果 job 函数出现 panic，并且该任务未进行 recover 处理，则该工作者会终止。
func (p *Pool) checkAndForkNewGoroutineWorker() {
	// 检查是否需要新建一个goroutine。
	var n int
	for {
		n = p.count.Val()
		if p.limit != -1 && n >= p.limit {
			// 无需创建新的goroutine。
			return
		}
		if p.count.Cas(n, n+1) {
			// 使用CAS（Compare And Swap）来保证原子性。
			break
		}
	}

	// 在goroutine中创建任务函数。
	go func() {
		defer p.count.Add(-1)

		var (
			listItem interface{}
			poolItem *localPoolItem
		)
		// 哈丁辛勤工作，逐个完成任务，工作永不停歇，工作者永不消亡。
		for !p.closed.Val() {
			listItem = p.list.PopBack()
			if listItem == nil {
				return
			}
			poolItem = listItem.(*localPoolItem)
			poolItem.Func(poolItem.Ctx)
		}
	}()
}
