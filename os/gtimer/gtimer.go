// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// gtimer 包实现了定时/延迟任务的执行与管理。
//
// 本包旨在管理数百万级别的定时任务。gtimer 与 gcron 的区别在于：
// 1. gcron 包是基于 gtimer 包实现的。
// 2. gtimer 侧重于高性能，适用于处理百万级的定时任务。
// 3. gcron 支持类似 Linux crontab 的配置模式语法，更加便于人工阅读。
// 4. gtimer 的基准操作时间单位为纳秒，而 gcron 的基准操作时间单位为微秒。
//
// 另外，特别注意定时器常见的延迟问题：https://github.com/golang/go/issues/14410 md5:c0dc9837a603eb26
package gtimer

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/command"
)

// Timer 是计时器管理器，它使用ticks来计算时间间隔。 md5:eab16bf4737f12a8
type Timer struct {
	mu      sync.RWMutex
	queue   *priorityQueue // queue 是基于堆结构的优先队列。 md5:6393b80eaafadb06
	status  *gtype.Int     // status是当前计时器的状态。 md5:86cd8448a559a7ec
	ticks   *gtype.Int64   // ticks 是计时器经过的间隔数量。 md5:6162dcb7888ba5fd
	options TimerOptions   // timer options 用于配置定时器的选项。 md5:a922957501fcab8b
}

// TimerOptions是Timer的配置对象。 md5:9d2942910cd72ea4
type TimerOptions struct {
	Interval time.Duration // (可选) Interval是计时器的基础滚动间隔刻度。 md5:4c933d7fc9ae2121
	Quick    bool          // Quick 用于快速计时器，这意味着计时器不会等待第一个间隔结束便立即触发。 md5:7b719aee1cbeb308
}

// internalPanic 是用于内部使用的自定义恐慌。 md5:287806e552654f1d
type internalPanic string

const (
	StatusReady                        = 0      // 作业或定时器已准备好运行。 md5:043261776a379433
	StatusRunning                      = 1      // 已经有任务或定时器在运行。 md5:148ed9b5d89215fe
	StatusStopped                      = 2      // 任务或定时器已停止。 md5:3cc8479a2061f7db
	StatusClosed                       = -1     // 作业或计时器已关闭，正在等待被删除。 md5:9b775a18fd7fe5b9
	panicExit            internalPanic = "exit" // panicExit 用于在发生恐慌时自定义作业退出。 md5:b22834505d9c74ec
	defaultTimerInterval               = "100"  // defaultTimerInterval 是默认的定时器间隔（以毫秒为单位）。 md5:8b2b5568f4526000
	// commandEnvKeyForInterval 是用于命令参数或环境配置中，定时器的默认间隔持续时间的键。 md5:baf94b4095117907
	commandEnvKeyForInterval = "gf.gtimer.interval"
)

var (
	defaultInterval = getDefaultInterval()
	defaultTimer    = New()
)

func getDefaultInterval() time.Duration {
	interval := command.GetOptWithEnv(commandEnvKeyForInterval, defaultTimerInterval)
	n, err := strconv.Atoi(interval)
	if err != nil {
		panic(gerror.WrapCodef(
			gcode.CodeInvalidConfiguration, err, `error converting string "%s" to int number`,
			interval,
		))
	}
	return time.Duration(n) * time.Millisecond
}

// DefaultOptions 创建并返回一个用于创建计时器的默认选项对象。 md5:67460fb8c6e56969
func DefaultOptions() TimerOptions {
	return TimerOptions{
		Interval: defaultInterval,
	}
}

// SetTimeout 在`delay`时间间隔后执行一次该任务。
// 它的作用类似于JavaScript中的同名函数。 md5:e0477460fecac4d8
func SetTimeout(ctx context.Context, delay time.Duration, job JobFunc) {
	AddOnce(ctx, delay, job)
}

// SetInterval 每隔 `delay` 的时间运行一次任务。
// 类似于 JavaScript 中的 SetInterval。 md5:1a47e33b3567d415
func SetInterval(ctx context.Context, interval time.Duration, job JobFunc) {
	Add(ctx, interval, job)
}

// Add 将一个计时任务添加到默认计时器中，该计时器以`interval`的间隔运行。 md5:222cab00f4afd790
func Add(ctx context.Context, interval time.Duration, job JobFunc) *Entry {
	return defaultTimer.Add(ctx, interval, job)
}

// AddEntry 向默认计时器添加一个带有详细参数的定时任务。
//
// 参数 `interval` 指定任务的运行间隔。
//
// 参数 `singleton` 指定任务是否以单例模式运行。
// 当任务为单例模式时，同一任务只允许有一个实例在运行。
//
// 参数 `times` 限制了任务的运行次数，意味着如果任务运行次数超过 `times` 就会退出。
//
// 参数 `status` 指定任务首次添加到计时器时的状态。 md5:0f65b7fd26b5f483
func AddEntry(ctx context.Context, interval time.Duration, job JobFunc, isSingleton bool, times int, status int) *Entry {
	return defaultTimer.AddEntry(ctx, interval, job, isSingleton, times, status)
}

// AddSingleton 是一个方便的函数，用于添加单例模式任务。 md5:8530c92e0f700eff
func AddSingleton(ctx context.Context, interval time.Duration, job JobFunc) *Entry {
	return defaultTimer.AddSingleton(ctx, interval, job)
}

// AddOnce 是一个方便函数，用于添加一个只运行一次然后退出的工作。 md5:7674bfdda1236b76
func AddOnce(ctx context.Context, interval time.Duration, job JobFunc) *Entry {
	return defaultTimer.AddOnce(ctx, interval, job)
}

// AddTimes 是一个方便的函数，用于添加一个有限运行次数的任务。 md5:98b9f39d64b7906c
func AddTimes(ctx context.Context, interval time.Duration, times int, job JobFunc) *Entry {
	return defaultTimer.AddTimes(ctx, interval, times, job)
}

// DelayAdd 在延迟`interval`持续时间后，添加一个定时任务。
// 参见Add。 md5:e1bb93aeff16693d
func DelayAdd(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc) {
	defaultTimer.DelayAdd(ctx, delay, interval, job)
}

// DelayAddEntry 在 `interval` 时长后添加一个定时任务。
// 另请参阅 AddEntry。 md5:e6e85b15472aaf98
func DelayAddEntry(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc, isSingleton bool, times int, status int) {
	defaultTimer.DelayAddEntry(ctx, delay, interval, job, isSingleton, times, status)
}

// DelayAddSingleton 在延迟`interval`时长后添加一个定时任务。
// 参阅 AddSingleton。 md5:80f70090fa17a370
func DelayAddSingleton(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc) {
	defaultTimer.DelayAddSingleton(ctx, delay, interval, job)
}

// DelayAddOnce 在延迟 `interval` 持续时间后添加一个定时任务。
// 参见 AddOnce。 md5:71c388c8096e5e48
func DelayAddOnce(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc) {
	defaultTimer.DelayAddOnce(ctx, delay, interval, job)
}

// DelayAddTimes 在延迟`interval`持续时间后，添加一个定时任务。
// 参见AddTimes。 md5:62463bf6c56709b2
func DelayAddTimes(ctx context.Context, delay time.Duration, interval time.Duration, times int, job JobFunc) {
	defaultTimer.DelayAddTimes(ctx, delay, interval, times, job)
}
