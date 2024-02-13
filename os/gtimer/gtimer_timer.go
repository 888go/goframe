// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 定时类

import (
	"context"
	"time"
	
	"github.com/888go/goframe/container/gtype"
)

// New 创建并返回一个 Timer。
func X创建(选项 ...TimerOptions) *Timer {
	t := &Timer{
		queue:  newPriorityQueue(),
		status: 安全变量类.NewInt(StatusRunning),
		ticks:  安全变量类.NewInt64(),
	}
	if len(选项) > 0 {
		t.options = 选项[0]
		if t.options.Interval == 0 {
			t.options.Interval = defaultInterval
		}
	} else {
		t.options = X取单例对象()
	}
	go t.loop()
	return t
}

// Add 向定时器添加一个计时任务，该任务以 `interval` 为间隔运行。
func (t *Timer) X加入循环任务(上下文 context.Context, 间隔时长 time.Duration, 任务函数 JobFunc) *Entry {
	return t.createEntry(createEntryInput{
		Ctx:         上下文,
		Interval:    间隔时长,
		Job:         任务函数,
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
func (t *Timer) X加入详细循环任务(上下文 context.Context, 间隔时长 time.Duration, 任务函数 JobFunc, 是否单例模式 bool, 次数 int, 任务状态 int) *Entry {
	return t.createEntry(createEntryInput{
		Ctx:         上下文,
		Interval:    间隔时长,
		Job:         任务函数,
		IsSingleton: 是否单例模式,
		Times:       次数,
		Status:      任务状态,
	})
}

// AddSingleton 是一个用于添加单例模式任务的便捷函数。
func (t *Timer) X加入单例循环任务(上下文 context.Context, 间隔时长 time.Duration, 任务函数 JobFunc) *Entry {
	return t.createEntry(createEntryInput{
		Ctx:         上下文,
		Interval:    间隔时长,
		Job:         任务函数,
		IsSingleton: true,
		Times:       -1,
		Status:      StatusReady,
	})
}

// AddOnce 是一个便捷函数，用于添加一个仅运行一次然后退出的任务。
func (t *Timer) X加入单次任务(上下文 context.Context, 间隔时长 time.Duration, 任务函数 JobFunc) *Entry {
	return t.createEntry(createEntryInput{
		Ctx:         上下文,
		Interval:    间隔时长,
		Job:         任务函数,
		IsSingleton: true,
		Times:       1,
		Status:      StatusReady,
	})
}

// AddTimes 是一个便捷函数，用于添加有一定运行次数限制的任务。
func (t *Timer) X加入指定次数任务(上下文 context.Context, 间隔时长 time.Duration, 次数 int, 任务函数 JobFunc) *Entry {
	return t.createEntry(createEntryInput{
		Ctx:         上下文,
		Interval:    间隔时长,
		Job:         任务函数,
		IsSingleton: true,
		Times:       次数,
		Status:      StatusReady,
	})
}

// DelayAdd 在 `delay` 延迟时间后添加一个定时任务。
// 也可参考 Add。
func (t *Timer) X延时加入循环任务(上下文 context.Context, 延时加入 time.Duration, 间隔时长 time.Duration, 任务函数 JobFunc) {
	t.X加入单次任务(上下文, 延时加入, func(ctx context.Context) {
		t.X加入循环任务(ctx, 间隔时长, 任务函数)
	})
}

// DelayAddEntry 在`delay`延迟时间后添加一个定时任务。
// 也可参考 AddEntry。
func (t *Timer) X延时加入详细循环任务(上下文 context.Context, 延时加入 time.Duration, 间隔时长 time.Duration, 任务函数 JobFunc, 是否单例模式 bool, 次数 int, 任务状态 int) {
	t.X加入单次任务(上下文, 延时加入, func(ctx context.Context) {
		t.X加入详细循环任务(ctx, 间隔时长, 任务函数, 是否单例模式, 次数, 任务状态)
	})
}

// DelayAddSingleton在`delay`延迟时间后添加一个定时任务。
// 也请参阅AddSingleton。
func (t *Timer) X延时加入单例循环任务(上下文 context.Context, 延时加入 time.Duration, 间隔时长 time.Duration, 任务函数 JobFunc) {
	t.X加入单次任务(上下文, 延时加入, func(ctx context.Context) {
		t.X加入单例循环任务(ctx, 间隔时长, 任务函数)
	})
}

// DelayAddOnce在`delay`延迟时间之后添加一个定时任务。
// 另请参阅AddOnce。
func (t *Timer) X延时加入单次任务(上下文 context.Context, 延时加入 time.Duration, 间隔时长 time.Duration, 任务函数 JobFunc) {
	t.X加入单次任务(上下文, 延时加入, func(ctx context.Context) {
		t.X加入单次任务(ctx, 间隔时长, 任务函数)
	})
}

// DelayAddTimes 在`delay`延迟时间后添加一个定时任务。
// 另请参阅 AddTimes。
func (t *Timer) X延时加入指定次数任务(上下文 context.Context, 延时加入 time.Duration, 间隔时长 time.Duration, 次数 int, 任务函数 JobFunc) {
	t.X加入单次任务(上下文, 延时加入, func(ctx context.Context) {
		t.X加入指定次数任务(ctx, 间隔时长, 次数, 任务函数)
	})
}

// Start 开始计时器。
func (t *Timer) X开始工作() {
	t.status.X设置值(StatusRunning)
}

// Stop 停止定时器。
func (t *Timer) X暂停工作() {
	t.status.X设置值(StatusStopped)
}

// Close 关闭计时器。
func (t *Timer) X关闭任务() {
	t.status.X设置值(StatusClosed)
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
		nextTicks = t.ticks.X取值()
	} else {
		nextTicks = t.ticks.X取值() + intervalTicksOfJob
	}
	var (
		entry = &Entry{
			job:         in.Job,
			ctx:         in.Ctx,
			timer:       t,
			ticks:       intervalTicksOfJob,
			times:       安全变量类.NewInt(in.Times),
			status:      安全变量类.NewInt(in.Status),
			isSingleton: 安全变量类.NewBool(in.IsSingleton),
			nextTicks:   安全变量类.NewInt64(nextTicks),
			infinite:    安全变量类.NewBool(infinite),
		}
	)
	t.queue.Push(entry, nextTicks)
	return entry
}
