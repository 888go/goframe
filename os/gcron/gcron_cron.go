// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcron

import (
	"context"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/os/gtimer"
)

type Cron struct {
	idGen   *gtype.Int64    // 用于生成唯一名称。
	status  *gtype.Int      // 定时任务状态(0: 未开始; 1: 运行中; 2: 已停止; -1: 关闭)
	entries *gmap.StrAnyMap // 所有定时任务条目。
	logger  glog.ILogger    // Logger，默认情况下为nil。
}

// New 函数返回一个使用默认设置的新 Cron 对象。
func New() *Cron {
	return &Cron{
		idGen:   gtype.NewInt64(),
		status:  gtype.NewInt(StatusRunning),
		entries: gmap.NewStrAnyMap(true),
	}
}

// SetLogger 用于设置cron的日志记录器。
func (c *Cron) SetLogger(logger glog.ILogger) {
	c.logger = logger
}

// GetLogger 返回cron中的日志器。
func (c *Cron) GetLogger() glog.ILogger {
	return c.logger
}

// AddEntry 创建并返回一个新的 Entry 对象。
func (c *Cron) AddEntry(ctx context.Context, pattern string, job JobFunc, times int, isSingleton bool, name ...string) (*Entry, error) {
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
// 可以用唯一的 `name` 与定时任务进行绑定。
// 如果 `name` 已经被使用，则返回错误。
func (c *Cron) Add(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error) {
	return c.AddEntry(ctx, pattern, job, -1, false, name...)
}

// AddSingleton 添加一个单例定时任务。
// 单例定时任务是指在同一时刻只能运行一个实例的任务。
// 可以使用唯一的 `name` 与定时任务绑定。
// 如果 `name` 已被使用，则返回错误。
func (c *Cron) AddSingleton(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error) {
	return c.AddEntry(ctx, pattern, job, -1, true, name...)
}

// AddTimes 添加一个可运行指定次数的定时任务。
// 可以使用唯一的 `name` 与定时任务关联绑定。
// 如果 `name` 已被使用，则返回错误。
func (c *Cron) AddTimes(ctx context.Context, pattern string, times int, job JobFunc, name ...string) (*Entry, error) {
	return c.AddEntry(ctx, pattern, job, times, false, name...)
}

// AddOnce 添加一个仅能运行一次的定时任务。
// 可以为定时任务绑定一个唯一的 `name`。
// 如果 `name` 已被使用，则返回错误。
func (c *Cron) AddOnce(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error) {
	return c.AddEntry(ctx, pattern, job, 1, false, name...)
}

// DelayAddEntry 在 `delay` 时间后添加一个定时任务。
func (c *Cron) DelayAddEntry(ctx context.Context, delay time.Duration, pattern string, job JobFunc, times int, isSingleton bool, name ...string) {
	gtimer.AddOnce(ctx, delay, func(ctx context.Context) {
		if _, err := c.AddEntry(ctx, pattern, job, times, isSingleton, name...); err != nil {
			panic(err)
		}
	})
}

// DelayAdd 在 `delay` 时间后添加一个定时任务。
func (c *Cron) DelayAdd(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string) {
	gtimer.AddOnce(ctx, delay, func(ctx context.Context) {
		if _, err := c.Add(ctx, pattern, job, name...); err != nil {
			panic(err)
		}
	})
}

// DelayAddSingleton 在`delay`时间后添加一个单例定时任务。
func (c *Cron) DelayAddSingleton(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string) {
	gtimer.AddOnce(ctx, delay, func(ctx context.Context) {
		if _, err := c.AddSingleton(ctx, pattern, job, name...); err != nil {
			panic(err)
		}
	})
}

// DelayAddOnce在`delay`时间后添加一个定时任务。
// 这个定时任务只能运行一次。
func (c *Cron) DelayAddOnce(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string) {
	gtimer.AddOnce(ctx, delay, func(ctx context.Context) {
		if _, err := c.AddOnce(ctx, pattern, job, name...); err != nil {
			panic(err)
		}
	})
}

// DelayAddTimes 在 `delay` 时间后添加一个定时任务。
// 此定时任务可以运行指定次数。
func (c *Cron) DelayAddTimes(ctx context.Context, delay time.Duration, pattern string, times int, job JobFunc, name ...string) {
	gtimer.AddOnce(ctx, delay, func(ctx context.Context) {
		if _, err := c.AddTimes(ctx, pattern, times, job, name...); err != nil {
			panic(err)
		}
	})
}

// Search 返回具有指定`name`的已计划任务。
// 如果未找到，则返回nil。
func (c *Cron) Search(name string) *Entry {
	if v := c.entries.Get(name); v != nil {
		return v.(*Entry)
	}
	return nil
}

// Start 开始运行指定的定时任务，名为 `name`。
// 如果未指定 `name`，则启动整个 cron。
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

// Stop 停止运行指定的定时任务，该任务名为 `name`。
// 如果未指定 `name`，则停止整个 cron。
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

// Remove 删除名为 `name` 的已计划任务。
func (c *Cron) Remove(name string) {
	if v := c.entries.Get(name); v != nil {
		v.(*Entry).Close()
	}
}

// Close 停止并关闭当前的cron任务。
func (c *Cron) Close() {
	c.status.Set(StatusClosed)
}

// Size 返回定时任务的数量。
func (c *Cron) Size() int {
	return c.entries.Size()
}

// Entries 返回所有按注册时间升序排列的定时任务切片。
func (c *Cron) Entries() []*Entry {
	array := garray.NewSortedArraySize(c.entries.Size(), func(v1, v2 interface{}) int {
		entry1 := v1.(*Entry)
		entry2 := v2.(*Entry)
		if entry1.Time.Nanosecond() > entry2.Time.Nanosecond() {
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
