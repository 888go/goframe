// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 定时类

import (
	"context"
	gcode "github.com/888go/goframe/errors/gcode"

	gtype "github.com/888go/goframe/container/gtype"
	gerror "github.com/888go/goframe/errors/gerror"
)

// Entry是定时任务。 md5:50a66c0cedad73c2
type Entry struct {
	job         JobFunc         // The job function.
	ctx         context.Context // 作业的上下文，只读。 md5:f8c9c52b497bf322
	timer       *Timer          // Belonged timer.
	ticks       int64           // 任务每滴答一声就执行一次。 md5:ee9f167eaedbe210
	times       *gtype.Int      // Limit running times.
	status      *gtype.Int      // Job status.
	isSingleton *gtype.Bool     // Singleton mode.
	nextTicks   *gtype.Int64    // 下一次运行任务的ticks。 md5:41046b30b7471bf3
	infinite    *gtype.Bool     // No times limit.
}

// JobFunc 是定时器中调用的定时任务函数。 md5:8958cfb2ccc06eff
type JobFunc = func(ctx context.Context)

// X取任务状态 返回作业的状态。 md5:2147922a20ade271
func (entry *Entry) X取任务状态() int {
	return entry.status.X取值()
}

// X异步运行 异步运行计时器任务。 md5:11fb4e6232736ab7
func (entry *Entry) X异步运行() {
	if !entry.infinite.X取值() {
		leftRunningTimes := entry.times.Add(-1)
				// 检查其运行时间是否超时。 md5:040ad43c2af11b3d
		if leftRunningTimes < 0 {
			entry.status.X设置值(StatusClosed)
			return
		}
	}
	go func() {
		defer func() {
			if exception := recover(); exception != nil {
				if exception != panicExit {
					if v, ok := exception.(error); ok && gerror.X判断是否带堆栈(v) {
						panic(v)
					} else {
						panic(gerror.X创建错误码并格式化(gcode.CodeInternalPanic, "exception recovered: %+v", exception))
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

// doCheckAndRunByTicks 检查在给定的计时器周期内，任务是否可以运行。如果当前的`currentTimerTicks`满足条件，它会异步运行；否则，它会增加其周期并等待下一次运行检查。
// md5:44d5223afb2e4f9d
func (entry *Entry) doCheckAndRunByTicks(currentTimerTicks int64) {
	// Ticks check.
	if currentTimerTicks < entry.nextTicks.X取值() {
		return
	}
	entry.nextTicks.X设置值(currentTimerTicks + entry.ticks)
	// Perform job checking.
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
	// Perform job running.
	entry.X异步运行()
}

// X设置任务状态 自定义设置作业的状态。 md5:c143d90d99990f2c
func (entry *Entry) X设置任务状态(状态 int) int {
	return entry.status.X设置值(状态)
}

// X开始工作 starts the job.
func (entry *Entry) X开始工作() {
	entry.status.X设置值(StatusReady)
}

// X暂停工作 stops the job.
func (entry *Entry) X暂停工作() {
	entry.status.X设置值(StatusStopped)
}

// X关闭任务 方法关闭任务，随后该任务将从计时器中移除。 md5:f499b51290bff676
func (entry *Entry) X关闭任务() {
	entry.status.X设置值(StatusClosed)
}

// X重置任务 重置作业，这将为下次运行重置其计数器。 md5:5a5ab5e4b73a76fe
func (entry *Entry) X重置任务() {
	entry.nextTicks.X设置值(entry.timer.ticks.X取值() + entry.ticks)
}

// X是否单例模式 检查并返回当前任务是否处于单例模式。 md5:a380e519564eb9da
func (entry *Entry) X是否单例模式() bool {
	return entry.isSingleton.X取值()
}

// X设置单例模式 设置单例模式。 md5:3fd379a01f57d11e
func (entry *Entry) X设置单例模式(单例模式 bool) {
	entry.isSingleton.X设置值(单例模式)
}

// X取任务函数 返回此任务的工作函数。 md5:38a44e496baf9d51
func (entry *Entry) X取任务函数() JobFunc {
	return entry.job
}

// X取任务上下文 返回此任务的初始化上下文。 md5:c2a3cb7932f8bf8a
func (entry *Entry) X取任务上下文() context.Context {
	return entry.ctx
}

// X设置任务次数 设置作业的运行次数限制。 md5:812717e2b2bcce7c
func (entry *Entry) X设置任务次数(次数 int) {
	entry.times.X设置值(次数)
	entry.infinite.X设置值(false)
}
