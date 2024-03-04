// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gtimer 实现了用于处理和管理间隔/延迟任务的定时器。
//
// 该包设计用于高效管理数百万计的定时任务。gtimer 和 gcron 包之间的区别如下：
//  1. 包 gcron 是基于 gtimer 包实现的。
//  2. gtimer 针对高性能场景设计，适用于处理数百万级别的定时任务。
//  3. gcron 支持类似 Linux crontab 的配置模式语法，更便于人工阅读。
//  4. gtimer 的基准操作（OP）性能以纳秒为单位衡量，而 gcron 的基准操作性能则以微秒为单位衡量。
//
// 另外，请特别注意定时器常见的延迟问题：https://github.com/golang/go/issues/14410
package gtimer

import (
	"context"
	"strconv"
	"sync"
	"time"
	
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/888go/goframe/gtimer/internal/command"
)

// Timer 是计时器管理器，它利用滴答（ticks）来计算定时间隔。
type Timer struct {
	mu      sync.RWMutex
	queue   *priorityQueue // queue 是一个基于堆结构的优先队列。
	status  *gtype.Int     // status 是当前计时器的状态。
	ticks   *gtype.Int64   // ticks 是定时器已处理的间隔数。
	options TimerOptions   // timer options 用于配置定时器。
}

// TimerOptions 是 Timer 的配置对象。
type TimerOptions struct {
	Interval time.Duration // (可选) Interval 是该定时器底层的滚动间隔刻度。
	Quick    bool          // Quick 用于快速计时器，这意味着计时器不会等待第一个间隔时间过去。
}

// internalPanic 是用于内部使用的自定义恐慌函数。
type internalPanic string

const (
	StatusReady                        = 0      // Job 或 Timer 准备就绪，可以开始运行。
	StatusRunning                      = 1      // 任务或计时器已经在运行中。
	StatusStopped                      = 2      // Job 或 Timer 已停止。
	StatusClosed                       = -1     // 作业或计时器已关闭，正在等待被删除。
	panicExit            internalPanic = "exit" // panicExit 用于通过 panic 进行自定义作业退出。
	defaultTimerInterval               = "100"  // defaultTimerInterval 是默认计时器间隔，单位为毫秒。
	// commandEnvKeyForInterval 是用于配置定时器默认间隔持续时间的命令行参数或环境变量的键。
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

// DefaultOptions 创建并返回一个用于创建Timer的默认选项对象。
func DefaultOptions() TimerOptions {
	return TimerOptions{
		Interval: defaultInterval,
	}
}

// SetTimeout 在 `delay` 延迟时间过后执行一次任务。
// 它类似于 JavaScript 中的 SetTimeout。
func SetTimeout(ctx context.Context, delay time.Duration, job JobFunc) {
	AddOnce(ctx, delay, job)
}

// SetInterval 每隔 `delay` 时间间隔运行 job。
// 它类似于 JavaScript 中的 SetInterval。
func SetInterval(ctx context.Context, interval time.Duration, job JobFunc) {
	Add(ctx, interval, job)
}

// Add 将一个定时任务添加到默认计时器，该计时器以`interval`为间隔运行。
func Add(ctx context.Context, interval time.Duration, job JobFunc) *Entry {
	return defaultTimer.Add(ctx, interval, job)
}

// AddEntry 向默认计时器添加一个具有详细参数的定时任务。
//
// 参数 `interval` 指定了任务的运行间隔时间。
//
// 参数 `singleton` 指定了任务是否以单例模式运行。如果任务是单例模式，那么同一任务在运行时只允许有一个实例。
//
// 参数 `times` 指定了任务运行次数的限制，这意味着当任务运行次数超过 `times` 时，该任务将退出。
//
// 参数 `status` 指定了任务首次添加到计时器时的状态。
func AddEntry(ctx context.Context, interval time.Duration, job JobFunc, isSingleton bool, times int, status int) *Entry {
	return defaultTimer.AddEntry(ctx, interval, job, isSingleton, times, status)
}

// AddSingleton 是一个用于添加单例模式任务的便捷函数。
func AddSingleton(ctx context.Context, interval time.Duration, job JobFunc) *Entry {
	return defaultTimer.AddSingleton(ctx, interval, job)
}

// AddOnce 是一个便捷函数，用于添加一个仅运行一次然后退出的任务。
func AddOnce(ctx context.Context, interval time.Duration, job JobFunc) *Entry {
	return defaultTimer.AddOnce(ctx, interval, job)
}

// AddTimes 是一个便捷函数，用于添加有一定运行次数限制的任务。
func AddTimes(ctx context.Context, interval time.Duration, times int, job JobFunc) *Entry {
	return defaultTimer.AddTimes(ctx, interval, times, job)
}

// DelayAdd 在 `interval` 延迟时间后添加一个定时任务。
// 另请参阅 Add。
func DelayAdd(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc) {
	defaultTimer.DelayAdd(ctx, delay, interval, job)
}

// DelayAddEntry 在 `interval` 延迟后添加一个定时任务。
// 也可参考 AddEntry。
func DelayAddEntry(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc, isSingleton bool, times int, status int) {
	defaultTimer.DelayAddEntry(ctx, delay, interval, job, isSingleton, times, status)
}

// DelayAddSingleton 在`interval`延迟时间后添加一个定时任务。
// 另请参阅 AddSingleton。
func DelayAddSingleton(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc) {
	defaultTimer.DelayAddSingleton(ctx, delay, interval, job)
}

// DelayAddOnce在`interval`延迟时间后添加一个定时任务。
// 另请参阅AddOnce。
func DelayAddOnce(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc) {
	defaultTimer.DelayAddOnce(ctx, delay, interval, job)
}

// DelayAddTimes 在 `interval` 延迟后添加一个定时任务。
// 另请参阅 AddTimes。
func DelayAddTimes(ctx context.Context, delay time.Duration, interval time.Duration, times int, job JobFunc) {
	defaultTimer.DelayAddTimes(ctx, delay, interval, times, job)
}
