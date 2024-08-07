// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 定时类

import (
	"context"
	"time"

	gtype "github.com/888go/goframe/container/gtype"
)

// X创建 创建并返回一个计时器。 md5:3db3235abce3957a
func X创建(选项 ...TimerOptions) *Timer {
	t := &Timer{
		queue:  newPriorityQueue(),
		status: gtype.NewInt(StatusRunning),
		ticks:  gtype.NewInt64(),
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

// X加入循环任务 向定时器添加一个定时任务，该任务以 `interval` 为间隔运行。 md5:358decaca6313fa2
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

// X加入详细循环任务 向计时器添加一个具有详细参数的定时任务。
//
// 参数 `interval` 指定任务的运行间隔。
//
// 参数 `singleton` 指定任务是否以单例模式运行。当任务为单例模式时，只允许运行一个相同的任务。
//
// 参数 `times` 指定任务的最大运行次数，即如果任务的运行次数超过 `times`，则退出。
//
// 参数 `status` 指定任务初次添加到计时器时的状态。
// md5:22c21ed21d95479c
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

// X加入单例循环任务 是一个方便的函数，用于添加单例模式任务。 md5:8530c92e0f700eff
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

// X加入单次任务 是一个方便函数，用于添加一个只运行一次然后退出的工作。 md5:7674bfdda1236b76
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

// X加入指定次数任务 是一个方便的函数，用于添加一个有限运行次数的任务。 md5:98b9f39d64b7906c
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

// X延时加入循环任务 在延迟`delay`持续时间后添加一个定时任务。
// 参阅 Add。
// md5:5db03c917e923b89
func (t *Timer) X延时加入循环任务(上下文 context.Context, 延时加入 time.Duration, 间隔时长 time.Duration, 任务函数 JobFunc) {
	t.X加入单次任务(上下文, 延时加入, func(ctx context.Context) {
		t.X加入循环任务(ctx, 间隔时长, 任务函数)
	})
}

// X延时加入详细循环任务 在延迟`delay`持续时间后，添加一个定时任务。
// 参见AddEntry。
// md5:6f230211b81dca10
func (t *Timer) X延时加入详细循环任务(上下文 context.Context, 延时加入 time.Duration, 间隔时长 time.Duration, 任务函数 JobFunc, 是否单例模式 bool, 次数 int, 任务状态 int) {
	t.X加入单次任务(上下文, 延时加入, func(ctx context.Context) {
		t.X加入详细循环任务(ctx, 间隔时长, 任务函数, 是否单例模式, 次数, 任务状态)
	})
}

// X延时加入单例循环任务 在延迟`delay`时间后添加一个定时任务。
// 另请参阅AddSingleton。
// md5:4df66d0755ab0371
func (t *Timer) X延时加入单例循环任务(上下文 context.Context, 延时加入 time.Duration, 间隔时长 time.Duration, 任务函数 JobFunc) {
	t.X加入单次任务(上下文, 延时加入, func(ctx context.Context) {
		t.X加入单例循环任务(ctx, 间隔时长, 任务函数)
	})
}

// X延时加入单次任务 在延迟 `delay` 持续时间后添加一个定时任务。
// 参见 AddOnce。
// md5:613532ca791628bf
func (t *Timer) X延时加入单次任务(上下文 context.Context, 延时加入 time.Duration, 间隔时长 time.Duration, 任务函数 JobFunc) {
	t.X加入单次任务(上下文, 延时加入, func(ctx context.Context) {
		t.X加入单次任务(ctx, 间隔时长, 任务函数)
	})
}

// X延时加入指定次数任务 在`delay`持续时间后添加一个定时任务。
// 参阅 AddTimes。
// md5:8a53452ea0989047
func (t *Timer) X延时加入指定次数任务(上下文 context.Context, 延时加入 time.Duration, 间隔时长 time.Duration, 次数 int, 任务函数 JobFunc) {
	t.X加入单次任务(上下文, 延时加入, func(ctx context.Context) {
		t.X加入指定次数任务(ctx, 间隔时长, 次数, 任务函数)
	})
}

// X开始工作 开始计时器。 md5:3330d3e996e68d8f
func (t *Timer) X开始工作() {
	t.status.X设置值(StatusRunning)
}

// X暂停工作 stops the timer.
func (t *Timer) X暂停工作() {
	t.status.X设置值(StatusStopped)
}

// X关闭任务 关闭定时器。 md5:cc2ed98b62fc4904
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

// createEntry 创建并添加一个计时任务到计时器中。 md5:8d0d62888bb8b536
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
		// 如果给定的间隔小于车轮的间隔，那么将其设置为一个刻度，这意味着它将在一个间隔内运行。
		// md5:4f1ce7e56c107c6e
		intervalTicksOfJob = 1
	}
	if t.options.Quick {
		// 如果启用了快速模式，意味着它将立即执行。
		// 无需等待第一个时间间隔。
		// md5:6d9ecc987797b1ba
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
