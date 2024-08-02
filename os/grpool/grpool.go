// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包grpool实现了可重用的goroutine池。 md5:8908f4659795e87e
package 协程类

import (
	"context"
	"time"

	glist "github.com/888go/goframe/container/glist"
	gtype "github.com/888go/goframe/container/gtype"
	gtimer "github.com/888go/goframe/os/gtimer"
	grand "github.com/888go/goframe/util/grand"
)

// Func 是包含上下文参数的池函数。 md5:6974cc4f941bf840
type Func func(ctx context.Context)

// RecoverFunc 是一个带有上下文参数的池运行时panic恢复函数。 md5:745651d2d0aad841
type RecoverFunc func(ctx context.Context, exception error)

// Pool 使用池管理goroutine。 md5:b07df489dc176e1c
type Pool struct {
	limit  int         // 最大goroutine计数限制。 md5:d8f9a54c9665e042
	count  *gtype.Int  // 当前正在运行的goroutine数量。 md5:ab841a9a1dadf1a0
	list   *glist.List // 用于异步添加任务的列表。 md5:65d48b67f1f6833d
	closed *gtype.Bool // Is pool closed or not.
}

// localPoolItem是存储在作业列表中的作业项。 md5:2b97e4fa0813a9cc
type localPoolItem struct {
	Ctx  context.Context // Context.
	Func Func            // Job function.
}

const (
	minSupervisorTimerDuration = 500 * time.Millisecond
	maxSupervisorTimerDuration = 1500 * time.Millisecond
)

// 默认的goroutine池。 md5:a834f6aca53d4225
var (
	defaultPool = New()
)

// New 创建并返回一个新的 goroutine 池对象。
// 参数 `limit` 用于限制最大 goroutine 数量，
// 默认情况下不限制。
// md5:523f5833750663c7
func New(limit ...int) *Pool {
	var (
		pool = &Pool{
			limit:  -1,
			count:  gtype.NewInt(),
			list:   glist.New(true),
			closed: gtype.NewBool(),
		}
		timerDuration = grand.D(
			minSupervisorTimerDuration,
			maxSupervisorTimerDuration,
		)
	)
	if len(limit) > 0 && limit[0] > 0 {
		pool.limit = limit[0]
	}
	gtimer.Add(context.Background(), timerDuration, pool.supervisor)
	return pool
}

// Add 将一个新任务添加到默认的 Goroutine 池中。
// 该任务将会异步执行。
// md5:2edd63bb852da48c
func Add(ctx context.Context, f Func) error {
	return defaultPool.Add(ctx, f)
}

// AddWithRecover 将一个新的任务推送到默认池中，指定一个恢复函数。
// 
// 可选的 `recoverFunc` 在执行 `userFunc` 时遇到任何 panic 时被调用。如果未传递或给定 `nil`，则忽略来自 `userFunc` 的 panic。任务将异步执行。
// md5:4b448b4fd7caa604
func AddWithRecover(ctx context.Context, userFunc Func, recoverFunc RecoverFunc) error {
	return defaultPool.AddWithRecover(ctx, userFunc, recoverFunc)
}

// Size 返回默认goroutine池中的当前goroutine数量。 md5:f66351deb380810c
func Size() int {
	return defaultPool.Size()
}

// Jobs 返回默认goroutine池当前的任务数量。 md5:d9b300d28d86f6c3
func Jobs() int {
	return defaultPool.Jobs()
}
