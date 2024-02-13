// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包grpool实现了goroutine可复用池。
package 协程类

import (
	"context"
	"time"
	
	"github.com/888go/goframe/container/glist"
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/os/gtimer"
	"github.com/888go/goframe/util/grand"
)

// Func 是一个包含上下文参数的池函数。
type Func func(ctx context.Context)

// RecoverFunc 是包含上下文参数的运行时 panic 恢复函数，用于池化组件。
type RecoverFunc func(ctx context.Context, exception error)

// Pool 通过使用 goroutine 池来管理 goroutines。
type Pool struct {
	limit  int         // 最大goroutine数量限制。
	count  *安全变量类.Int  // 当前运行的 goroutine 数量
	list   *链表类.List // 用于异步任务添加目的的列表。
	closed *安全变量类.Bool // 是否已关闭池。
}

// localPoolItem 是存储在任务列表中的工作项。
type localPoolItem struct {
	Ctx  context.Context // Context.
	Func Func            // Job function.
}

const (
	minSupervisorTimerDuration = 500 * time.Millisecond
	maxSupervisorTimerDuration = 1500 * time.Millisecond
)

// 默认的 goroutine 池。
var (
	defaultPool = New()
)

// New 创建并返回一个新的 goroutine 池对象。
// 参数 `limit` 用于限制最大goroutine数量，默认情况下不限制。
func New(limit ...int) *Pool {
	var (
		pool = &Pool{
			limit:  -1,
			count:  安全变量类.NewInt(),
			list:   链表类.New(true),
			closed: 安全变量类.NewBool(),
		}
		timerDuration = 随机类.X时长(
			minSupervisorTimerDuration,
			maxSupervisorTimerDuration,
		)
	)
	if len(limit) > 0 && limit[0] > 0 {
		pool.limit = limit[0]
	}
	定时类.X加入循环任务(context.Background(), timerDuration, pool.supervisor)
	return pool
}

// Add 向默认的goroutine池中添加一个新的任务。
// 该任务将会被异步执行。
func Add(ctx context.Context, f Func) error {
	return defaultPool.Add(ctx, f)
}

// AddWithRecover 将带有指定恢复函数的新任务推送到默认工作池。
//
// 可选的 `recoverFunc` 参数在执行 `userFunc` 函数期间发生任何 panic 时被调用。
// 如果未传递或给定 nil，则忽略来自 `userFunc` 的 panic。
// 该任务将被异步执行。
func AddWithRecover(ctx context.Context, userFunc Func, recoverFunc RecoverFunc) error {
	return defaultPool.AddWithRecover(ctx, userFunc, recoverFunc)
}

// Size 返回默认协程池中当前的协程数量。
func Size() int {
	return defaultPool.Size()
}

// Jobs 返回默认goroutine池当前的任务数量。
func Jobs() int {
	return defaultPool.Jobs()
}
