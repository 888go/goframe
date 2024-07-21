// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package grpool

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// Add 将一个新任务添加到池中。
// 该任务将会异步执行。
// md5:69389d53e280086b
// ff:
// p:
// ctx:
// f:
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
	// 检查并 fork 新的 worker。 md5:d3acb042c3373fa4
	p.checkAndForkNewGoroutineWorker()
	return nil
}

// AddWithRecover 将指定的恢复函数推送到队列中执行新任务。
// 
// 可选的 `recoverFunc` 在执行 `userFunc` 时发生任何 panic 时被调用。如果未传递或给定 nil，它将忽略来自 `userFunc` 的 panic。任务将异步执行。
// md5:764d1260466b9a5d
// ff:
// p:
// ctx:
// userFunc:
// recoverFunc:
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
// 如果没有限制，则返回-1。
// md5:1c6cae16429df1b2
// ff:
// p:
func (p *Pool) Cap() int {
	return p.limit
}

// Size 返回当前池中的goroutine数量。 md5:247eb1685633ccc3
// ff:
// p:
func (p *Pool) Size() int {
	return p.count.Val()
}

// Jobs 返回池中的当前任务数。
// 注意，它返回的不是工作器/goroutine的数量，而是任务的数量。
// md5:c82d92b33047974c
// ff:
// p:
func (p *Pool) Jobs() int {
	return p.list.Size()
}

// IsClosed 返回池是否已关闭。 md5:85755176347bcfea
// ff:
// p:
func (p *Pool) IsClosed() bool {
	return p.closed.Val()
}

// Close 关闭goroutine池，导致所有goroutines退出。 md5:3d9c73ed9b0f4643
// ff:
// p:
func (p *Pool) Close() {
	p.closed.Set(true)
}

// checkAndForkNewGoroutineWorker 检查并创建一个新的goroutine工作进程。
// 请注意，如果工作函数出现恐慌且该工作没有恢复处理，那么该工作进程将会死亡。
// md5:242a912451066181
func (p *Pool) checkAndForkNewGoroutineWorker() {
	// 检查是否需要在新的goroutine中 fork。 md5:20ef20b082ef0b86
	var n int
	for {
		n = p.count.Val()
		if p.limit != -1 && n >= p.limit {
			// 不需要启动新的goroutine。 md5:a4d7257aa086311e
			return
		}
		if p.count.Cas(n, n+1) {
			// 使用CAS（比较并交换）来保证操作的原子性。 md5:2337a31243acf132
			break
		}
	}

	// 在goroutine中创建任务函数。 md5:e0e70df051fd0b1a
	go func() {
		defer p.count.Add(-1)

		var (
			listItem interface{}
			poolItem *localPoolItem
		)
		// 哈丁工作，一个接一个，任务永无止境，工人永不消亡。 md5:625670ae6a926602
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
