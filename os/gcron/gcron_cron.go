// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gcron

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtimer"
)

// Cron 存储所有的cron作业项。 md5:3a72b04261d69d0e
type Cron struct {
	idGen   *gtype.Int64    // 用于生成唯一名称。 md5:6a2c8b5e748394aa
	status  *gtype.Int      // 定时任务状态 (0: 未开始; 1: 运行中; 2: 已停止; -1: 已关闭). md5:c07ab3c74e0378b5
	entries *gmap.StrAnyMap // 所有定时任务条目。 md5:a74f7440115592e5
	logger  glog.ILogger    // Logger，默认情况下为nil。 md5:49f083354c677555
}

// New 返回一个具有默认设置的新 Cron 对象。 md5:d0ca0545e6ef9b46
// ff:
func New() *Cron {
	return &Cron{
		idGen:   gtype.NewInt64(),
		status:  gtype.NewInt(StatusRunning),
		entries: gmap.NewStrAnyMap(true),
	}
}

// SetLogger 为 cron 设置日志记录器。 md5:87e9222eac80e2a1
// ff:
// c:
// logger:
func (c *Cron) SetLogger(logger glog.ILogger) {
	c.logger = logger
}

// GetLogger 在cron中返回日志记录器。 md5:014719534223048e
// ff:
// c:
func (c *Cron) GetLogger() glog.ILogger {
	return c.logger
}

// AddEntry 创建并返回一个新的 Entry 对象。 md5:1a7d5709c2867437
// ff:
// c:
func (c *Cron) AddEntry(
	ctx context.Context,
	pattern string,
	job JobFunc,
	times int,
	isSingleton bool,
	name ...string,
) (*Entry, error) {
	var (
		entryName = ""
		infinite  = false
	)
	if len(name) > 0 {
		entryName = name[0]
	}
	if times <= 0 {
		infinite = true
	}
	return c.doAddEntry(doAddEntryInput{
		Name:        entryName,
		Job:         job,
		Ctx:         ctx,
		Times:       times,
		Pattern:     pattern,
		IsSingleton: isSingleton,
		Infinite:    infinite,
	})
}

// Add 添加一个定时任务。
// 一个唯一的`name`可以与定时任务绑定。
// 如果`name`已经被使用，它将返回一个错误。
// md5:850ebd654a2e3695
// ff:
// c:
// ctx:
// pattern:
// job:
// name:
func (c *Cron) Add(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error) {
	return c.AddEntry(ctx, pattern, job, -1, false, name...)
}

// AddSingleton 添加一个单例定时任务。
// 单例定时任务是指在同一时间只能运行一个实例的任务。
// 可以通过一个唯一的`name`与定时任务绑定。
// 如果`name`已经被使用，它将返回错误。
// md5:9e0e86c2aee09877
// ff:
// c:
// ctx:
// pattern:
// job:
// name:
func (c *Cron) AddSingleton(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error) {
	return c.AddEntry(ctx, pattern, job, -1, true, name...)
}

// AddTimes 添加一个定时任务，可以指定执行次数。
// 可以为定时任务绑定一个唯一的`name`。
// 如果`name`已存在，将返回一个错误。
// md5:b01e5695f9cc54d5
// ff:
// c:
// ctx:
// pattern:
// times:
// job:
// name:
func (c *Cron) AddTimes(ctx context.Context, pattern string, times int, job JobFunc, name ...string) (*Entry, error) {
	return c.AddEntry(ctx, pattern, job, times, false, name...)
}

// AddOnce 添加一个定时任务，该任务只能运行一次。
// 可以为定时任务绑定一个唯一的`name`。
// 如果`name`已使用，将返回一个错误。
// md5:fd5fb4f131e1f937
// ff:
// c:
// ctx:
// pattern:
// job:
// name:
func (c *Cron) AddOnce(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error) {
	return c.AddEntry(ctx, pattern, job, 1, false, name...)
}

// DelayAddEntry 在延迟`delay`时间后添加一个定时任务。 md5:a45391b1d2daacd5
// ff:
// c:
// ctx:
// delay:
// pattern:
// job:
// times:
// isSingleton:
// name:
func (c *Cron) DelayAddEntry(ctx context.Context, delay time.Duration, pattern string, job JobFunc, times int, isSingleton bool, name ...string) {
	gtimer.AddOnce(ctx, delay, func(ctx context.Context) {
		if _, err := c.AddEntry(ctx, pattern, job, times, isSingleton, name...); err != nil {
			panic(err)
		}
	})
}

// DelayAdd 在`delay`时间后添加一个定时任务。 md5:027e39a4b8e3b167
// ff:
// c:
// ctx:
// delay:
// pattern:
// job:
// name:
func (c *Cron) DelayAdd(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string) {
	gtimer.AddOnce(ctx, delay, func(ctx context.Context) {
		if _, err := c.Add(ctx, pattern, job, name...); err != nil {
			panic(err)
		}
	})
}

// DelayAddSingleton 在`delay`时间后添加一个单例计时任务。 md5:c56847cf6733a3e4
// ff:
// c:
// ctx:
// delay:
// pattern:
// job:
// name:
func (c *Cron) DelayAddSingleton(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string) {
	gtimer.AddOnce(ctx, delay, func(ctx context.Context) {
		if _, err := c.AddSingleton(ctx, pattern, job, name...); err != nil {
			panic(err)
		}
	})
}

// DelayAddOnce 在 `delay` 时间后添加一个定时任务。
// 这个定时任务只能运行一次。
// md5:34aa0df8fb8e5477
// ff:
// c:
// ctx:
// delay:
// pattern:
// job:
// name:
func (c *Cron) DelayAddOnce(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string) {
	gtimer.AddOnce(ctx, delay, func(ctx context.Context) {
		if _, err := c.AddOnce(ctx, pattern, job, name...); err != nil {
			panic(err)
		}
	})
}

// DelayAddTimes 在延迟`delay`时间后添加一个定时任务。
// 该定时任务可以指定运行次数。
// md5:5ed58fb7650ed0bb
// ff:
// c:
// ctx:
// delay:
// pattern:
// times:
// job:
// name:
func (c *Cron) DelayAddTimes(ctx context.Context, delay time.Duration, pattern string, times int, job JobFunc, name ...string) {
	gtimer.AddOnce(ctx, delay, func(ctx context.Context) {
		if _, err := c.AddTimes(ctx, pattern, times, job, name...); err != nil {
			panic(err)
		}
	})
}

// Search 返回指定名称的计划任务。
// 如果未找到，则返回 nil。
// md5:b0da4b1e1203c6c7
// ff:
// c:
// name:
func (c *Cron) Search(name string) *Entry {
	if v := c.entries.Get(name); v != nil {
		return v.(*Entry)
	}
	return nil
}

// Start 启动指定名为 `name` 的定时任务。如果没有指定 `name`，则启动整个 cron。
// md5:d573745c6d8edaac
// ff:
// c:
// name:
func (c *Cron) Start(name ...string) {
	if len(name) > 0 {
		for _, v := range name {
			if entry := c.Search(v); entry != nil {
				entry.Start()
			}
		}
	} else {
		c.status.Set(StatusReady)
	}
}

// Stop 停止运行指定的定时任务，任务名为 `name`。如果未指定 `name`，则停止整个cron（cron job）。
// md5:68ed27359d633f5e
// ff:
// c:
// name:
func (c *Cron) Stop(name ...string) {
	if len(name) > 0 {
		for _, v := range name {
			if entry := c.Search(v); entry != nil {
				entry.Stop()
			}
		}
	} else {
		c.status.Set(StatusStopped)
	}
}

// Remove 删除名为`name`的计划任务。 md5:bc96b6bdb0bac57b
// ff:
// c:
// name:
func (c *Cron) Remove(name string) {
	if v := c.entries.Get(name); v != nil {
		v.(*Entry).Close()
	}
}

// Close 停止并关闭当前的cron。 md5:95a4276ef94fb50c
// ff:
// c:
func (c *Cron) Close() {
	c.status.Set(StatusClosed)
}

// Size返回定时任务的大小。 md5:a282381f7ca9bf53
// ff:
// c:
func (c *Cron) Size() int {
	return c.entries.Size()
}

// Entries 返回所有定时任务作为切片（按注册时间升序排列）。 md5:67b4f559a25d411e
// ff:
// c:
func (c *Cron) Entries() []*Entry {
	array := garray.NewSortedArraySize(c.entries.Size(), func(v1, v2 interface{}) int {
		entry1 := v1.(*Entry)
		entry2 := v2.(*Entry)
		if entry1.RegisterTime.Nanosecond() > entry2.RegisterTime.Nanosecond() {
			return 1
		}
		return -1
	}, true)
	c.entries.RLockFunc(func(m map[string]interface{}) {
		for _, v := range m {
			array.Add(v.(*Entry))
		}
	})
	entries := make([]*Entry, array.Len())
	array.RLockFunc(func(array []interface{}) {
		for k, v := range array {
			entries[k] = v.(*Entry)
		}
	})
	return entries
}
