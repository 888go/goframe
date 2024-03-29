// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 定时类

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/errors/gerror"
)

// Entry 是定时任务。
type Entry struct {
	job         JobFunc         // The job function.
	ctx         context.Context // 该上下文为作业的上下文，仅用于 READ ONLY（只读）。
	timer       *Timer          // Belonged timer.
	ticks       int64           // 该任务在每个时间间隔运行一次。
	times       *gtype.Int      // 限制运行次数。
	status      *gtype.Int      // Job status.
	isSingleton *gtype.Bool     // Singleton mode.
	nextTicks   *gtype.Int64    // 下一次运行该任务的时间刻度
	infinite    *gtype.Bool     // No times limit.
}

// JobFunc 是在定时器中被定时调用的任务函数。
type JobFunc = func(ctx context.Context)

// Status 返回作业的状态。
func (entry *Entry) X取任务状态() int {
	return entry.status.Val()
}

// Run 启动计时器任务并异步执行。
func (entry *Entry) X异步运行() {
	if !entry.infinite.Val() {
		leftRunningTimes := entry.times.Add(-1)
		// 它检查运行时间是否超过限制。
		if leftRunningTimes < 0 {
			entry.status.Set(StatusClosed)
			return
		}
	}
	go func() {
		defer func() {
			if exception := recover(); exception != nil {
				if exception != panicExit {
					if v, ok := exception.(error); ok && gerror.HasStack(v) {
						panic(v)
					} else {
						panic(gerror.NewCodef(gcode.CodeInternalPanic, "exception recovered: %+v", exception))
					}
				} else {
					entry.X关闭任务()
					return
				}
			}
			if entry.X取任务状态() == StatusRunning {
				entry.X设置任务状态(StatusReady)
			}
		}()
		entry.job(entry.ctx)
	}()
}

// doCheckAndRunByTicks 根据给定的定时器刻度检查任务是否可以运行，
// 如果给定的 `currentTimerTicks` 满足条件，则异步执行该任务；
// 否则，增加其刻度并等待下次运行检查。
func (entry *Entry) doCheckAndRunByTicks(currentTimerTicks int64) {
	// Ticks check.
	if currentTimerTicks < entry.nextTicks.Val() {
		return
	}
	entry.nextTicks.Set(currentTimerTicks + entry.ticks)
	// 执行任务检查。
	switch entry.status.Val() {
	case StatusRunning:
		if entry.X是否单例模式() {
			return
		}
	case StatusReady:
		if !entry.status.Cas(StatusReady, StatusRunning) {
			return
		}
	case StatusStopped:
		return
	case StatusClosed:
		return
	}
	// 执行任务运行。
	entry.X异步运行()
}

// SetStatus 自定义设置任务的状态。
func (entry *Entry) X设置任务状态(状态 int) int {
	return entry.status.Set(状态)
}

// Start 启动任务。
func (entry *Entry) X开始工作() {
	entry.status.Set(StatusReady)
}

// Stop 停止作业。
func (entry *Entry) X暂停工作() {
	entry.status.Set(StatusStopped)
}

// Close 关闭作业，随后该作业将从计时器中移除。
func (entry *Entry) X关闭任务() {
	entry.status.Set(StatusClosed)
}

// Reset 重置作业，这将重置其下一次运行的滴答次数。
func (entry *Entry) X重置任务() {
	entry.nextTicks.Set(entry.timer.ticks.Val() + entry.ticks)
}

// IsSingleton 检查并返回该任务是否处于单例模式。
func (entry *Entry) X是否单例模式() bool {
	return entry.isSingleton.Val()
}

// SetSingleton 设置作业单例模式。
func (entry *Entry) X设置单例模式(单例模式 bool) {
	entry.isSingleton.Set(单例模式)
}

// Job 返回此任务的工作函数。
func (entry *Entry) X取任务函数() JobFunc {
	return entry.job
}

// Ctx 返回此任务初始化后的上下文。
func (entry *Entry) X取任务上下文() context.Context {
	return entry.ctx
}

// SetTimes 设置作业的最大运行次数。
func (entry *Entry) X设置任务次数(次数 int) {
	entry.times.Set(次数)
	entry.infinite.Set(false)
}
