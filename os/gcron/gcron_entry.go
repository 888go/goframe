// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcron
import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"time"
	
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/os/gtimer"
	"github.com/888go/goframe/util/gconv"
	)
// JobFunc 是 cron 中被定时调用的任务函数。
type JobFunc = gtimer.JobFunc

// Entry 是定时任务的入口。
type Entry struct {
	cron       *Cron         // Cron对象所属的。
	timerEntry *gtimer.Entry // 关联的定时器条目。
	schedule   *cronSchedule // 定时调度对象
	jobName    string        // 回调函数名称(地址信息)。
	times      *gtype.Int    // 运行次数限制。
	infinite   *gtype.Bool   // No times limit.
	Name       string        // Entry name.
	Job        JobFunc       `json:"-"` // 回调函数。
	Time       time.Time     // Registered time.
}

type doAddEntryInput struct {
	Name        string          // Name 为该条目设置名称以便进行手动控制。
	Job         JobFunc         // Job 是用于定时任务执行的回调函数。
	Ctx         context.Context // 该作业的上下文。
	Times       int             // Times 指定该条目运行的限制次数。
	Pattern     string          // Pattern 是用于调度器的 crontab 风格字符串。
	IsSingleton bool            // Singleton 指定定时任务是否以单例模式执行。
	Infinite    bool            // Infinite 指定此条目是否在无时间限制的情况下运行。
}

// doAddEntry 创建并返回一个新的 Entry 对象。
func (c *Cron) doAddEntry(in doAddEntryInput) (*Entry, error) {
	if in.Name != "" {
		if c.Search(in.Name) != nil {
			return nil, gerror.NewCodef(gcode.CodeInvalidOperation, `cron job "%s" already exists`, in.Name)
		}
	}
	schedule, err := newSchedule(in.Pattern)
	if err != nil {
		return nil, err
	}
	// 对于 `times` 没有限制，表示该定时器用于每秒检查调度。
	entry := &Entry{
		cron:     c,
		schedule: schedule,
		jobName:  runtime.FuncForPC(reflect.ValueOf(in.Job).Pointer()).Name(),
		times:    gtype.NewInt(in.Times),
		infinite: gtype.NewBool(in.Infinite),
		Job:      in.Job,
		Time:     time.Now(),
	}
	if in.Name != "" {
		entry.Name = in.Name
	} else {
		entry.Name = "cron-" + gconv.String(c.idGen.Add(1))
	}
// 当你添加一个计划任务时，不能允许它立即运行。
// 它不能在添加到timer时就开始运行。
// 应该在该任务条目被添加到Cron的entries映射中之后才开始运行，以避免在添加过程中由于entries尚未拥有该条目的信息而运行任务，这可能会导致panic。
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

// IsSingleton 返回该条目是否为单例定时任务。
func (entry *Entry) IsSingleton() bool {
	return entry.timerEntry.IsSingleton()
}

// SetSingleton 设置入口以单例模式运行。
func (entry *Entry) SetSingleton(enabled bool) {
	entry.timerEntry.SetSingleton(enabled)
}

// SetTimes 设置条目可以运行的时间。
func (entry *Entry) SetTimes(times int) {
	entry.times.Set(times)
	entry.infinite.Set(false)
}

// Status 返回 entry 的状态。
func (entry *Entry) Status() int {
	return entry.timerEntry.Status()
}

// SetStatus 设置条目的状态。
func (entry *Entry) SetStatus(status int) int {
	return entry.timerEntry.SetStatus(status)
}

// Start 启动运行入口。
func (entry *Entry) Start() {
	entry.timerEntry.Start()
}

// Stop 停止运行 entry。
func (entry *Entry) Stop() {
	entry.timerEntry.Stop()
}

// Close 停止该任务并从 cron 中移除该条目。
func (entry *Entry) Close() {
	entry.cron.entries.Remove(entry.Name)
	entry.timerEntry.Close()
}

// checkAndRun 是核心的定时任务检查逻辑。
func (entry *Entry) checkAndRun(ctx context.Context) {
	currentTime := time.Now()
	if !entry.schedule.checkMeetAndUpdateLastSeconds(ctx, currentTime) {
		return
	}
	switch entry.cron.status.Val() {
	case StatusStopped:
		return

	case StatusClosed:
		entry.logDebugf(ctx, `cron job "%s" is removed`, entry.getJobNameWithPattern())
		entry.Close()

	case StatusReady, StatusRunning:
		defer func() {
			if exception := recover(); exception != nil {
				// 捕获到异常，按照默认行为将错误内容记录到日志器中。
				entry.logErrorf(ctx,
					`cron job "%s(%s)" end with error: %+v`,
					entry.jobName, entry.schedule.pattern, exception,
				)
			} else {
				entry.logDebugf(ctx, `cron job "%s" ends`, entry.getJobNameWithPattern())
			}
			if entry.timerEntry.Status() == StatusClosed {
				entry.Close()
			}
		}()

		// 运行时间检查。
		if !entry.infinite.Val() {
			times := entry.times.Add(-1)
			if times <= 0 {
				if entry.timerEntry.SetStatus(StatusClosed) == StatusClosed || times < 0 {
					return
				}
			}
		}
		entry.logDebugf(ctx, `cron job "%s" starts`, entry.getJobNameWithPattern())
		entry.Job(ctx)
	}
}

func (entry *Entry) getJobNameWithPattern() string {
	return fmt.Sprintf(`%s(%s)`, entry.jobName, entry.schedule.pattern)
}

func (entry *Entry) logDebugf(ctx context.Context, format string, v ...interface{}) {
	if logger := entry.cron.GetLogger(); logger != nil {
		logger.Debugf(ctx, format, v...)
	}
}

func (entry *Entry) logErrorf(ctx context.Context, format string, v ...interface{}) {
	logger := entry.cron.GetLogger()
	if logger == nil {
		logger = glog.DefaultLogger()
	}
	logger.Errorf(ctx, format, v...)
}
