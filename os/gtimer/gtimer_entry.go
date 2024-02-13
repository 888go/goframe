// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 定时类

import (
	"context"
	"github.com/888go/goframe/errors/gcode"
	
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/errors/gerror"
)

// Entry 是定时任务。
type Entry struct {
	job         JobFunc         // The job function.
	ctx         context.Context // 该上下文为作业的上下文，仅用于 READ ONLY（只读）。
	timer       *Timer          // Belonged timer.
	ticks       int64           // 该任务在每个时间间隔运行一次。
	times       *安全变量类.Int      // 限制运行次数。
	status      *安全变量类.Int      // Job status.
	isSingleton *安全变量类.Bool     // Singleton mode.
	nextTicks   *安全变量类.Int64    // 下一次运行该任务的时间刻度
	infinite    *安全变量类.Bool     // No times limit.
}

// JobFunc 是在定时器中被定时调用的任务函数。
type JobFunc = func(ctx context.Context)

// Status 返回作业的状态。
func (entry *Entry) X取任务状态() int {
	return entry.status.X取值()
}

// Run 启动计时器任务并异步执行。
func (entry *Entry) X异步运行() {
	if !entry.infinite.X取值() {
		leftRunningTimes := entry.times.Add(-1)
		// 它检查运行时间是否超过限制。
		if leftRunningTimes < 0 {
			entry.status.X设置值(StatusClosed)
			return
		}
	}
	go func() {
		defer func() {
			if exception := recover(); exception != nil {
				if exception != panicExit {
					if v, ok := exception.(error); ok && 错误类.X判断是否带堆栈(v) {
						panic(v)
					} else {
						panic(错误类.X创建错误码并格式化(错误码类.CodeInternalPanic, "exception recovered: %+v", exception))
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
	if currentTimerTicks < entry.nextTicks.X取值() {
		return
	}
	entry.nextTicks.X设置值(currentTimerTicks + entry.ticks)
	// 执行任务检查。
	switch entry.status.X取值() {
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
	return entry.status.X设置值(状态)
}

// Start 启动任务。
func (entry *Entry) X开始工作() {
	entry.status.X设置值(StatusReady)
}

// Stop 停止作业。
func (entry *Entry) X暂停工作() {
	entry.status.X设置值(StatusStopped)
}

// Close 关闭作业，随后该作业将从计时器中移除。
func (entry *Entry) X关闭任务() {
	entry.status.X设置值(StatusClosed)
}

// Reset 重置作业，这将重置其下一次运行的滴答次数。
func (entry *Entry) X重置任务() {
	entry.nextTicks.X设置值(entry.timer.ticks.X取值() + entry.ticks)
}

// IsSingleton 检查并返回该任务是否处于单例模式。
func (entry *Entry) X是否单例模式() bool {
	return entry.isSingleton.X取值()
}

// SetSingleton 设置作业单例模式。
func (entry *Entry) X设置单例模式(单例模式 bool) {
	entry.isSingleton.X设置值(单例模式)
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
	entry.times.X设置值(次数)
	entry.infinite.X设置值(false)
}
