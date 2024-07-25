// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// gcron 包实现了cron模式解析器和任务运行器。 md5:56d461ee2c75e1f7
package gcron

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtimer"
)

const (
	StatusReady   = gtimer.StatusReady
	StatusRunning = gtimer.StatusRunning
	StatusStopped = gtimer.StatusStopped
	StatusClosed  = gtimer.StatusClosed
)

var (
	// Default cron object.
	defaultCron = New()
)

// SetLogger 设置 cron 的全局日志记录器。 md5:b0a9c7514a8c8b29
func SetLogger(logger glog.ILogger) {
	defaultCron.SetLogger(logger)
}

// GetLogger 返回 cron 中的全局日志记录器。 md5:908af8c408f25d3d
func GetLogger() glog.ILogger {
	return defaultCron.GetLogger()
}

// Add 向默认的 cron 对象添加一个定时任务。
// 一个唯一的 `name` 可以与定时任务绑定。
// 如果 `name` 已经被使用，它将返回一个错误。
// md5:0f46b08a5a96144a
func Add(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error) {
	return defaultCron.Add(ctx, pattern, job, name...)
}

// AddSingleton 向默认的 cron 对象添加一个单例定时任务。
// 单例定时任务是指在同一时间只能运行一个实例的任务。
// 可以将一个唯一的 `name` 与定时任务绑定。
// 如果 `name` 已经被使用，它将返回错误。
// md5:96268d3e8373359e
func AddSingleton(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error) {
	return defaultCron.AddSingleton(ctx, pattern, job, name...)
}

// AddOnce 向默认cron对象添加一个仅能运行一次的定时任务。可以为定时任务绑定一个唯一的`name`。
// 如果`name`已存在，它将返回一个错误。
// md5:9701998ce952852d
func AddOnce(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error) {
	return defaultCron.AddOnce(ctx, pattern, job, name...)
}

// AddTimes 将一个定时任务添加到默认cron对象中，该任务可以执行指定次数。可以为定时任务绑定一个唯一的`name`。
// 如果`name`已经存在，它将返回一个错误。
// md5:505a2bbf10529705
func AddTimes(ctx context.Context, pattern string, times int, job JobFunc, name ...string) (*Entry, error) {
	return defaultCron.AddTimes(ctx, pattern, times, job, name...)
}

// DelayAdd 在延迟'delay'时间后，向默认的cron对象添加一个定时任务。 md5:7c28aadbf64d1362
func DelayAdd(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string) {
	defaultCron.DelayAdd(ctx, delay, pattern, job, name...)
}

// DelayAddSingleton 在默认的cron对象中延迟`delay`时间后添加一个单例定时任务。 md5:d6c26c1edc16c19a
func DelayAddSingleton(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string) {
	defaultCron.DelayAddSingleton(ctx, delay, pattern, job, name...)
}

// DelayAddOnce 在 `delay` 时间后向默认cron对象添加一个定时任务。
// 这个定时任务只能运行一次。
// md5:56e7e748e4d4c63d
func DelayAddOnce(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string) {
	defaultCron.DelayAddOnce(ctx, delay, pattern, job, name...)
}

// DelayAddTimes 在`delay`时间后向默认cron对象添加一个定时任务。这个定时任务可以运行指定次数。
// md5:b4ecf68ee4e86408
func DelayAddTimes(ctx context.Context, delay time.Duration, pattern string, times int, job JobFunc, name ...string) {
	defaultCron.DelayAddTimes(ctx, delay, pattern, times, job, name...)
}

// Search 返回具有指定`name`的计划任务。
// 如果未找到，返回nil。
// md5:06da84fc0766d888
func Search(name string) *Entry {
	return defaultCron.Search(name)
}

// Remove 删除名为`name`的计划任务。 md5:bc96b6bdb0bac57b
func Remove(name string) {
	defaultCron.Remove(name)
}

// Size 返回默认cron中定时任务的数量。 md5:e87d680e31031739
func Size() int {
	return defaultCron.Size()
}

// Entries 返回所有定时任务作为切片。 md5:16823f1ebafbd9cc
func Entries() []*Entry {
	return defaultCron.Entries()
}

// Start 启动指定名为 `name` 的定时任务。如果没有指定 `name`，则启动整个 cron。
// md5:d573745c6d8edaac
func Start(name ...string) {
	defaultCron.Start(name...)
}

// Stop 停止运行指定的定时任务，任务名为 `name`。如果未指定 `name`，则停止整个cron（cron job）。
// md5:68ed27359d633f5e
func Stop(name ...string) {
	defaultCron.Stop(name...)
}
