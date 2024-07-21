// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtimer

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"

	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/errors/gerror"
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

// Status 返回作业的状态。 md5:2147922a20ade271
func (entry *Entry) Status() int {
	return entry.status.Val()
}

// Run 异步运行计时器任务。 md5:11fb4e6232736ab7
func (entry *Entry) Run() {
	if !entry.infinite.Val() {
		leftRunningTimes := entry.times.Add(-1)
		// 检查其运行时间是否超时。 md5:040ad43c2af11b3d
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
					entry.Close()
					return
				}
			}
			if entry.Status() == StatusRunning {
				entry.SetStatus(StatusReady)
			}
		}()
		entry.job(entry.ctx)
	}()
}

// doCheckAndRunByTicks 检查在给定的计时器周期内，任务是否可以运行。如果当前的`currentTimerTicks`满足条件，它会异步运行；否则，它会增加其周期并等待下一次运行检查。
// md5:44d5223afb2e4f9d
func (entry *Entry) doCheckAndRunByTicks(currentTimerTicks int64) {
	// Ticks check.
	if currentTimerTicks < entry.nextTicks.Val() {
		return
	}
	entry.nextTicks.Set(currentTimerTicks + entry.ticks)
	// Perform job checking.
	switch entry.status.Val() {
	case StatusRunning:
		if entry.IsSingleton() {
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
	entry.Run()
}

// SetStatus 自定义设置作业的状态。 md5:c143d90d99990f2c
func (entry *Entry) SetStatus(status int) int {
	return entry.status.Set(status)
}

// Start starts the job.
func (entry *Entry) Start() {
	entry.status.Set(StatusReady)
}

// Stop stops the job.
func (entry *Entry) Stop() {
	entry.status.Set(StatusStopped)
}

// Close 方法关闭任务，随后该任务将从计时器中移除。 md5:f499b51290bff676
func (entry *Entry) Close() {
	entry.status.Set(StatusClosed)
}

// Reset 重置作业，这将为下次运行重置其计数器。 md5:5a5ab5e4b73a76fe
func (entry *Entry) Reset() {
	entry.nextTicks.Set(entry.timer.ticks.Val() + entry.ticks)
}

// IsSingleton 检查并返回当前任务是否处于单例模式。 md5:a380e519564eb9da
func (entry *Entry) IsSingleton() bool {
	return entry.isSingleton.Val()
}

// SetSingleton 设置单例模式。 md5:3fd379a01f57d11e
func (entry *Entry) SetSingleton(enabled bool) {
	entry.isSingleton.Set(enabled)
}

// Job 返回此任务的工作函数。 md5:38a44e496baf9d51
func (entry *Entry) Job() JobFunc {
	return entry.job
}

// Ctx 返回此任务的初始化上下文。 md5:c2a3cb7932f8bf8a
func (entry *Entry) Ctx() context.Context {
	return entry.ctx
}

// SetTimes 设置作业的运行次数限制。 md5:812717e2b2bcce7c
func (entry *Entry) SetTimes(times int) {
	entry.times.Set(times)
	entry.infinite.Set(false)
}
