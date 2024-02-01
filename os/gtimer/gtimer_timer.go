// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtimer
import (
	"context"
	"time"
	
	"github.com/888go/goframe/container/gtype"
	)
// New 创建并返回一个 Timer。
func New(options ...TimerOptions) *Timer {
	t := &Timer{
		queue:  newPriorityQueue(),
		status: gtype.NewInt(StatusRunning),
		ticks:  gtype.NewInt64(),
	}
	if len(options) > 0 {
		t.options = options[0]
		if t.options.Interval == 0 {
			t.options.Interval = defaultInterval
		}
	} else {
		t.options = DefaultOptions()
	}
	go t.loop()
	return t
}

// Add 向定时器添加一个计时任务，该任务以 `interval` 为间隔运行。
func (t *Timer) Add(ctx context.Context, interval time.Duration, job JobFunc) *Entry {
	return t.createEntry(createEntryInput{
		Ctx:         ctx,
		Interval:    interval,
		Job:         job,
		IsSingleton: false,
		Times:       -1,
		Status:      StatusReady,
	})
}

// AddEntry 向定时器添加一个具有详细参数的定时任务。
//
// 参数 `interval` 指定了该任务的运行间隔时间。
//
// 参数 `singleton` 指定了任务是否以单例模式运行。如果是单例模式的任务，同一任务在运行时只允许存在一个实例。
//
// 参数 `times` 指定了任务运行次数的限制，意味着当任务运行次数超过 `times` 时，该任务将结束运行。
//
// 参数 `status` 指定了任务首次添加到定时器时的状态。
func (t *Timer) AddEntry(ctx context.Context, interval time.Duration, job JobFunc, isSingleton bool, times int, status int) *Entry {
	return t.createEntry(createEntryInput{
		Ctx:         ctx,
		Interval:    interval,
		Job:         job,
		IsSingleton: isSingleton,
		Times:       times,
		Status:      status,
	})
}

// AddSingleton 是一个用于添加单例模式任务的便捷函数。
func (t *Timer) AddSingleton(ctx context.Context, interval time.Duration, job JobFunc) *Entry {
	return t.createEntry(createEntryInput{
		Ctx:         ctx,
		Interval:    interval,
		Job:         job,
		IsSingleton: true,
		Times:       -1,
		Status:      StatusReady,
	})
}

// AddOnce 是一个便捷函数，用于添加一个仅运行一次然后退出的任务。
func (t *Timer) AddOnce(ctx context.Context, interval time.Duration, job JobFunc) *Entry {
	return t.createEntry(createEntryInput{
		Ctx:         ctx,
		Interval:    interval,
		Job:         job,
		IsSingleton: true,
		Times:       1,
		Status:      StatusReady,
	})
}

// AddTimes 是一个便捷函数，用于添加有一定运行次数限制的任务。
func (t *Timer) AddTimes(ctx context.Context, interval time.Duration, times int, job JobFunc) *Entry {
	return t.createEntry(createEntryInput{
		Ctx:         ctx,
		Interval:    interval,
		Job:         job,
		IsSingleton: true,
		Times:       times,
		Status:      StatusReady,
	})
}

// DelayAdd 在 `delay` 延迟时间后添加一个定时任务。
// 也可参考 Add。
func (t *Timer) DelayAdd(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc) {
	t.AddOnce(ctx, delay, func(ctx context.Context) {
		t.Add(ctx, interval, job)
	})
}

// DelayAddEntry 在`delay`延迟时间后添加一个定时任务。
// 也可参考 AddEntry。
func (t *Timer) DelayAddEntry(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc, isSingleton bool, times int, status int) {
	t.AddOnce(ctx, delay, func(ctx context.Context) {
		t.AddEntry(ctx, interval, job, isSingleton, times, status)
	})
}

// DelayAddSingleton在`delay`延迟时间后添加一个定时任务。
// 也请参阅AddSingleton。
func (t *Timer) DelayAddSingleton(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc) {
	t.AddOnce(ctx, delay, func(ctx context.Context) {
		t.AddSingleton(ctx, interval, job)
	})
}

// DelayAddOnce在`delay`延迟时间之后添加一个定时任务。
// 另请参阅AddOnce。
func (t *Timer) DelayAddOnce(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc) {
	t.AddOnce(ctx, delay, func(ctx context.Context) {
		t.AddOnce(ctx, interval, job)
	})
}

// DelayAddTimes 在`delay`延迟时间后添加一个定时任务。
// 另请参阅 AddTimes。
func (t *Timer) DelayAddTimes(ctx context.Context, delay time.Duration, interval time.Duration, times int, job JobFunc) {
	t.AddOnce(ctx, delay, func(ctx context.Context) {
		t.AddTimes(ctx, interval, times, job)
	})
}

// Start 开始计时器。
func (t *Timer) Start() {
	t.status.Set(StatusRunning)
}

// Stop 停止定时器。
func (t *Timer) Stop() {
	t.status.Set(StatusStopped)
}

// Close 关闭计时器。
func (t *Timer) Close() {
	t.status.Set(StatusClosed)
}

type createEntryInput struct {
	Ctx         context.Context
	Interval    time.Duration
	Job         JobFunc
	IsSingleton bool
	Times       int
	Status      int
}

// createEntry 创建并添加一个定时任务到计时器中。
func (t *Timer) createEntry(in createEntryInput) *Entry {
	var (
		infinite  = false
		nextTicks int64
	)
	if in.Times <= 0 {
		infinite = true
	}
	var (
		intervalTicksOfJob = int64(in.Interval / t.options.Interval)
	)
	if intervalTicksOfJob == 0 {
// 如果给定的间隔小于轮子的间隔，
// 则将其设置为一个刻度，这意味着它将在一个间隔后运行。
		intervalTicksOfJob = 1
	}
	if t.options.Quick {
// 如果启用了快速模式，意味着它将立即运行。
// 不需要等待第一个间隔时间。
		nextTicks = t.ticks.Val()
	} else {
		nextTicks = t.ticks.Val() + intervalTicksOfJob
	}
	var (
		entry = &Entry{
			job:         in.Job,
			ctx:         in.Ctx,
			timer:       t,
			ticks:       intervalTicksOfJob,
			times:       gtype.NewInt(in.Times),
			status:      gtype.NewInt(in.Status),
			isSingleton: gtype.NewBool(in.IsSingleton),
			nextTicks:   gtype.NewInt64(nextTicks),
			infinite:    gtype.NewBool(infinite),
		}
	)
	t.queue.Push(entry, nextTicks)
	return entry
}
