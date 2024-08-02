// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 定时cron类

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"time"

	gtype "github.com/888go/goframe/container/gtype"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	glog "github.com/888go/goframe/os/glog"
	gtimer "github.com/888go/goframe/os/gtimer"
	gconv "github.com/888go/goframe/util/gconv"
)

// JobFunc 是在cron中定时被调用的作业函数。 md5:476072dbc1ca96ff
type JobFunc = gtimer.JobFunc

// Entry 是定时任务的入口。 md5:ebc0ffec0c819fa5
type Entry struct {
	cron         *Cron         // 属于的Cron对象。 md5:b0d31cd3bc86924f
	timerEntry   *gtimer.Entry // 相关的定时器条目。 md5:cc0c92f42e2f28c0
	schedule     *cronSchedule // Timed schedule object.
	jobName      string        // 回调函数名称（地址信息）。 md5:7f130ad66e787420
	times        *gtype.Int    // Running times limit.
	infinite     *gtype.Bool   // No times limit.
	Name         string        // Entry name.
	RegisterTime time.Time     // Registered time.
	Job          JobFunc       `json:"-"` // Callback function.
}

type doAddEntryInput struct {
	Name        string          // Name 为手动控制此条目提供名称。 md5:ab6add69f0d6cead
	Job         JobFunc         // Job是定时任务执行的回调函数。 md5:7faad71757692791
	Ctx         context.Context // 作业的上下文。 md5:505a832c1cca7c0e
	Times       int             // Times 指定了条目的运行限制次数。 md5:bb45eebc85ac5fcc
	Pattern     string          // Pattern 是用于调度的 crontab 风格的字符串。 md5:2076c5de0b0a0288
	IsSingleton bool            // Singleton 指定是否以单例模式执行定时任务。 md5:f98d6fb605fc538f
	Infinite    bool            // Infinite表示此条目是否无限运行，不受时间限制。 md5:14267dbbffb84b38
}

// doAddEntry 创建并返回一个新的 Entry 对象。 md5:c29537f522377c29
func (c *Cron) doAddEntry(in doAddEntryInput) (*Entry, error) {
	if in.Name != "" {
		if c.Search(in.Name) != nil {
			return nil, gerror.NewCodef(
				gcode.CodeInvalidOperation,
				`duplicated cron job name "%s", already exists`,
				in.Name,
			)
		}
	}
	schedule, err := newSchedule(in.Pattern)
	if err != nil {
		return nil, err
	}
		// 对于`times`没有限制，用于每秒检查调度的计时器。 md5:dce371ba28b68c21
	entry := &Entry{
		cron:         c,
		schedule:     schedule,
		jobName:      runtime.FuncForPC(reflect.ValueOf(in.Job).Pointer()).Name(),
		times:        gtype.NewInt(in.Times),
		infinite:     gtype.NewBool(in.Infinite),
		RegisterTime: time.Now(),
		Job:          in.Job,
	}
	if in.Name != "" {
		entry.Name = in.Name
	} else {
		entry.Name = "cron-" + gconv.String(c.idGen.Add(1))
	}
	// 当你添加一个定时任务时，不能让它立即运行。
	// 它在添加到计时器时不能开始执行。
	// 应该在将条目添加到Cron条目映射后开始运行，以防止在添加过程中（此时条目信息可能不完整）任务运行，从而可能导致恐慌。
	// md5:e2b503aef8166c84
	entry.timerEntry = gtimer.AddEntry(
		in.Ctx,
		time.Second,
		entry.checkAndRun,
		in.IsSingleton,
		-1,
		gtimer.StatusStopped,
	)
	c.entries.Set(entry.Name, entry)
	entry.timerEntry.Start()
	return entry, nil
}

// IsSingleton 判断这个条目是否是定时单例任务。 md5:171967c731b60f88
func (e *Entry) IsSingleton() bool {
	return e.timerEntry.IsSingleton()
}

// SetSingleton 设置以单例模式运行的条目。 md5:6c81a3a09d0ef0b1
func (e *Entry) SetSingleton(enabled bool) {
	e.timerEntry.SetSingleton(enabled)
}

// SetTimes 设置条目可以运行的时间。 md5:663af054d5aab5e9
func (e *Entry) SetTimes(times int) {
	e.times.Set(times)
	e.infinite.Set(false)
}

// Status 返回条目的状态。 md5:6a9d3438dc575881
func (e *Entry) Status() int {
	return e.timerEntry.Status()
}

// SetStatus 设置条目的状态。 md5:ea0ecb4171f3f017
func (e *Entry) SetStatus(status int) int {
	return e.timerEntry.SetStatus(status)
}

// Start 开始运行入口函数。 md5:aa729d73eb626ca1
func (e *Entry) Start() {
	e.timerEntry.Start()
}

// Stop 停止运行条目。 md5:06d53148d6536ce9
func (e *Entry) Stop() {
	e.timerEntry.Stop()
}

// Close 停止并从 cron 中移除条目。 md5:a2a5eee9228cd918
func (e *Entry) Close() {
	e.cron.entries.Remove(e.Name)
	e.timerEntry.Close()
}

// checkAndRun是核心定时任务检查逻辑。
// 这个函数每秒被调用一次。
// md5:e4a94345d01fd7df
func (e *Entry) checkAndRun(ctx context.Context) {
	currentTime := time.Now()
	if !e.schedule.checkMeetAndUpdateLastSeconds(ctx, currentTime) {
		return
	}
	switch e.cron.status.Val() {
	case StatusStopped:
		return

	case StatusClosed:
		e.logDebugf(ctx, `cron job "%s" is removed`, e.getJobNameWithPattern())
		e.Close()

	case StatusReady, StatusRunning:
		defer func() {
			if exception := recover(); exception != nil {
								// 捕获到异常，默认行为是将错误内容记录到日志中。 md5:9af43ee64f5795bf
				e.logErrorf(ctx,
					`cron job "%s(%s)" end with error: %+v`,
					e.jobName, e.schedule.pattern, exception,
				)
			} else {
				e.logDebugf(ctx, `cron job "%s" ends`, e.getJobNameWithPattern())
			}
			if e.timerEntry.Status() == StatusClosed {
				e.Close()
			}
		}()

		// Running times check.
		if !e.infinite.Val() {
			times := e.times.Add(-1)
			if times <= 0 {
				if e.timerEntry.SetStatus(StatusClosed) == StatusClosed || times < 0 {
					return
				}
			}
		}
		e.logDebugf(ctx, `cron job "%s" starts`, e.getJobNameWithPattern())
		e.Job(ctx)
	}
}

func (e *Entry) getJobNameWithPattern() string {
	return fmt.Sprintf(`%s(%s)`, e.jobName, e.schedule.pattern)
}

func (e *Entry) logDebugf(ctx context.Context, format string, v ...interface{}) {
	if logger := e.cron.GetLogger(); logger != nil {
		logger.Debugf(ctx, format, v...)
	}
}

func (e *Entry) logErrorf(ctx context.Context, format string, v ...interface{}) {
	logger := e.cron.GetLogger()
	if logger == nil {
		logger = glog.DefaultLogger()
	}
	logger.Errorf(ctx, format, v...)
}
