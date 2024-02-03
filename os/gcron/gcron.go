// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gcron实现了cron模式解析器和任务执行器。
package gcron

import (
	"context"
	"time"
	
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/os/gtimer"
)

const (
	StatusReady   = gtimer.StatusReady
	StatusRunning = gtimer.StatusRunning
	StatusStopped = gtimer.StatusStopped
	StatusClosed  = gtimer.StatusClosed
)

var (
	// 默认的cron对象。
	defaultCron = New()
)

// SetLogger 用于设置cron的日志记录器。
func SetLogger(logger glog.ILogger) {
	defaultCron.SetLogger(logger)
}

// GetLogger 返回cron中的日志器。
func GetLogger() glog.ILogger {
	return defaultCron.GetLogger()
}

// Add 向默认的 cron 对象添加一个定时任务。
// 可以用唯一的 `name` 与定时任务关联。
// 如果 `name` 已经被使用，则返回错误。
func Add(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error) {
	return defaultCron.Add(ctx, pattern, job, name...)
}

// AddSingleton 添加单例定时任务到默认的cron对象。
// 单例定时任务是指在同一时间只能运行一个实例的任务。
// 可以通过唯一的`name`与定时任务关联绑定。
// 如果`name`已经被使用，则返回错误。
func AddSingleton(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error) {
	return defaultCron.AddSingleton(ctx, pattern, job, name...)
}

// AddOnce 向默认的cron对象添加一个仅能执行一次的定时任务。
// 可以为定时任务绑定一个唯一的`name`标识。
// 如果`name`已被使用，则返回错误。
func AddOnce(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error) {
	return defaultCron.AddOnce(ctx, pattern, job, name...)
}

// AddTimes 向默认的 cron 对象添加一个可运行指定次数的定时任务。
// 可以为定时任务绑定一个唯一的 `name`。
// 如果 `name` 已经被使用，将会返回错误。
func AddTimes(ctx context.Context, pattern string, times int, job JobFunc, name ...string) (*Entry, error) {
	return defaultCron.AddTimes(ctx, pattern, times, job, name...)
}

// DelayAdd 在 `delay` 时间后向默认的 cron 对象添加一个定时任务。
func DelayAdd(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string) {
	defaultCron.DelayAdd(ctx, delay, pattern, job, name...)
}

// DelayAddSingleton 在`delay`时间后向默认的cron对象添加一个单例定时任务。
func DelayAddSingleton(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string) {
	defaultCron.DelayAddSingleton(ctx, delay, pattern, job, name...)
}

// DelayAddOnce在`delay`时间后向默认的cron对象添加一个定时任务。
// 这个定时任务只能运行一次。
func DelayAddOnce(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string) {
	defaultCron.DelayAddOnce(ctx, delay, pattern, job, name...)
}

// DelayAddTimes 在`delay`时间后向默认的cron对象添加一个定时任务。
// 这个定时任务可以运行指定次数。
func DelayAddTimes(ctx context.Context, delay time.Duration, pattern string, times int, job JobFunc, name ...string) {
	defaultCron.DelayAddTimes(ctx, delay, pattern, times, job, name...)
}

// Search 返回具有指定名称 `name` 的已计划任务。
// 如果未找到，则返回 nil。
func Search(name string) *Entry {
	return defaultCron.Search(name)
}

// Remove 删除名为 `name` 的已计划任务。
func Remove(name string) {
	defaultCron.Remove(name)
}

// Size 返回默认Cron定时任务的数量。
func Size() int {
	return defaultCron.Size()
}

// Entries 返回所有定时任务作为一个切片。
func Entries() []*Entry {
	return defaultCron.Entries()
}

// Start 开始运行指定的定时任务，名为 `name`。
// 如果未指定 `name`，则启动整个 cron。
func Start(name ...string) {
	defaultCron.Start(name...)
}

// Stop 停止运行指定的定时任务，该任务名为 `name`。
// 如果未指定 `name`，则停止整个 cron。
func Stop(name ...string) {
	defaultCron.Stop(name...)
}
