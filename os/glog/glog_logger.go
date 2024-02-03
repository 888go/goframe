// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package glog

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
	
	"github.com/fatih/color"
	"go.opentelemetry.io/otel/trace"
	
	"github.com/888go/goframe/debug/gdebug"
	"github.com/888go/goframe/internal/consts"
	"github.com/888go/goframe/internal/errors"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gfpool"
	"github.com/888go/goframe/os/gmlock"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/util/gconv"
)

// Logger 是用于日志管理的结构体。
type Logger struct {
	parent *Logger // 父级日志器，如果非空，则表示该日志器在链式函数中使用。
	config Config  // 日志器配置。
}

const (
	defaultFileFormat                 = `{Y-m-d}.log`
	defaultFileFlags                  = os.O_CREATE | os.O_WRONLY | os.O_APPEND
	defaultFilePerm                   = os.FileMode(0666)
	defaultFileExpire                 = time.Minute
	pathFilterKey                     = "/os/glog/glog"
	memoryLockPrefixForPrintingToFile = "glog.printToFile:"
)

const (
	F_ASYNC      = 1 << iota // Print logging content asynchronously。
	F_FILE_LONG              // 打印完整的文件名和行号：/a/b/c/d.go:23.
	F_FILE_SHORT             // 打印文件名的最后一个元素和行号：d.go:23。该选项覆盖了 F_FILE_LONG。
	F_TIME_DATE              // 在本地时区打印日期：2009-01-23。
	F_TIME_TIME              // 在本地时区打印时间：01:23:23
// ```go
// 下面是示例代码（假设）
// package main
// import (
//     "fmt"
//     "time"
// )
// func main() {
    // 获取当前时间
//     now := time.Now()
    // 格式化为 HH:mm:ss 的格式
//     formattedTime := now.Format("15:04:05")
    // 打印在本地时区的时间：01:23:23
//     fmt.Println(formattedTime)
// }
// 这段注释描述的是要在本地时区打印时间的意图，但实际代码需要获取当前时间并进行格式化以匹配指定格式。
	F_TIME_MILLI             // 在本地时区打印带有毫秒的时间：01:23:23.675。
	F_CALLER_FN              // 打印调用函数名称和包名：main.main
	F_TIME_STD   = F_TIME_DATE | F_TIME_MILLI
)

// New 创建并返回一个自定义日志器。
func New() *Logger {
	return &Logger{
		config: DefaultConfig(),
	}
}

// NewWithWriter 通过 io.Writer 创建并返回一个自定义的日志记录器。
func NewWithWriter(writer io.Writer) *Logger {
	l := New()
	l.SetWriter(writer)
	return l
}

// Clone 返回一个新的日志器，它是当前日志器的“浅复制”。
// 注意，克隆后的日志器其 `config` 属性是对当前日志器该属性的浅复制。
func (l *Logger) Clone() *Logger {
	return &Logger{
		config: l.config,
		parent: l,
	}
}

// getFilePath 返回日志文件的路径。
// 日志文件名必须包含 ".log" 扩展名。
func (l *Logger) getFilePath(now time.Time) string {
	// 文件名中包含 "{}" 的内容将使用 gtime 进行格式化。
	file, _ := gregex.ReplaceStringFunc(`{.+?}`, l.config.File, func(s string) string {
		return gtime.New(now).Format(strings.Trim(s, "{}"))
	})
	file = gfile.Join(l.config.Path, file)
	return file
}

// print 将 `s` 打印输出到预定义的 writer（写入器）、日志文件或传入的 `std`。
func (l *Logger) print(ctx context.Context, level int, stack string, values ...any) {
// 为日志旋转功能进行惰性初始化。
// 使用原子读取操作以提升性能检查的效率。
// 这里使用了CAP以保证性能和并发安全性。
// 对于每个日志器，仅初始化一次。
	if l.config.RotateSize > 0 || l.config.RotateExpire > 0 {
		if !l.config.rotatedHandlerInitialized.Val() && l.config.rotatedHandlerInitialized.Cas(false, true) {
			l.rotateChecksTimely(ctx)
			intlog.Printf(ctx, "logger rotation initialized: every %s", l.config.RotateCheckInterval.String())
		}
	}

	var (
		now   = time.Now()
		input = &HandlerInput{
			internalHandlerInfo: internalHandlerInfo{
				index: -1,
			},
			Logger: l,
			Buffer: bytes.NewBuffer(nil),
			Time:   now,
			Color:  defaultLevelColor[level],
			Level:  level,
			Stack:  stack,
			Values: values,
		}
	)

	// Logging handlers.
	if len(l.config.Handlers) > 0 {
		input.handlers = append(input.handlers, l.config.Handlers...)
	} else if defaultHandler != nil {
		input.handlers = []Handler{defaultHandler}
	}
	input.handlers = append(input.handlers, doFinalPrint)

	// Time.
	timeFormat := ""
	if l.config.TimeFormat != "" {
		timeFormat = l.config.TimeFormat
	} else {
		if l.config.Flags&F_TIME_DATE > 0 {
			timeFormat += "2006-01-02"
		}
		if l.config.Flags&F_TIME_TIME > 0 {
			if timeFormat != "" {
				timeFormat += " "
			}
			timeFormat += "15:04:05"
		}
		if l.config.Flags&F_TIME_MILLI > 0 {
			if timeFormat != "" {
				timeFormat += " "
			}
			timeFormat += "15:04:05.000"
		}
	}

	if len(timeFormat) > 0 {
		input.TimeFormat = now.Format(timeFormat)
	}

	// Level string.
	input.LevelFormat = l.GetLevelPrefix(level)

	// 调用路径和函数名称。
	if l.config.Flags&(F_FILE_LONG|F_FILE_SHORT|F_CALLER_FN) > 0 {
		callerFnName, path, line := gdebug.CallerWithFilter(
			[]string{consts.StackFilterKeyForGoFrame},
			l.config.StSkip,
		)
		if l.config.Flags&F_CALLER_FN > 0 {
			if len(callerFnName) > 2 {
				input.CallerFunc = fmt.Sprintf(`[%s]`, callerFnName)
			}
		}
		if line >= 0 && len(path) > 1 {
			if l.config.Flags&F_FILE_LONG > 0 {
				input.CallerPath = fmt.Sprintf(`%s:%d:`, path, line)
			}
			if l.config.Flags&F_FILE_SHORT > 0 {
				input.CallerPath = fmt.Sprintf(`%s:%d:`, gfile.Basename(path), line)
			}
		}
	}
	// Prefix.
	if len(l.config.Prefix) > 0 {
		input.Prefix = l.config.Prefix
	}

	// 将值转换为字符串。
	if ctx != nil {
		// Tracing values.
		spanCtx := trace.SpanContextFromContext(ctx)
		if traceId := spanCtx.TraceID(); traceId.IsValid() {
			input.TraceId = traceId.String()
		}
		// Context values.
		if len(l.config.CtxKeys) > 0 {
			for _, ctxKey := range l.config.CtxKeys {
				var ctxValue interface{}
				if ctxValue = ctx.Value(ctxKey); ctxValue == nil {
					ctxValue = ctx.Value(gctx.StrKey(gconv.String(ctxKey)))
				}
				if ctxValue != nil {
					if input.CtxStr != "" {
						input.CtxStr += ", "
					}
					input.CtxStr += gconv.String(ctxValue)
				}
			}
		}
	}
	if l.config.Flags&F_ASYNC > 0 {
		input.IsAsync = true
		err := asyncPool.Add(ctx, func(ctx context.Context) {
			input.Next(ctx)
		})
		if err != nil {
			intlog.Errorf(ctx, `%+v`, err)
		}
	} else {
		input.Next(ctx)
	}
}

// doFinalPrint 根据配置输出日志内容。
func (l *Logger) doFinalPrint(ctx context.Context, input *HandlerInput) *bytes.Buffer {
	var buffer *bytes.Buffer
	// 是否允许输出到标准输出（stdout）？
	if l.config.StdoutPrint {
		if buf := l.printToStdout(ctx, input); buf != nil {
			buffer = buf
		}
	}

	// 将内容输出到磁盘文件。
	if l.config.Path != "" {
		if buf := l.printToFile(ctx, input.Time, input); buf != nil {
			buffer = buf
		}
	}

	// 使用了自定义的写入器。
	if l.config.Writer != nil {
		// 将输出发送到自定义写入器。
		if buf := l.printToWriter(ctx, input); buf != nil {
			buffer = buf
		}
	}
	return buffer
}

// printToWriter 将缓冲区内容写入到writer中。
func (l *Logger) printToWriter(ctx context.Context, input *HandlerInput) *bytes.Buffer {
	if l.config.Writer != nil {
		var buffer = input.getRealBuffer(l.config.WriterColorEnable)
		if _, err := l.config.Writer.Write(buffer.Bytes()); err != nil {
			intlog.Errorf(ctx, `%+v`, err)
		}
		return buffer
	}
	return nil
}

// printToStdout 将日志内容输出到 stdout（标准输出）。
func (l *Logger) printToStdout(ctx context.Context, input *HandlerInput) *bytes.Buffer {
	if l.config.StdoutPrint {
		var (
			err    error
			buffer = input.getRealBuffer(!l.config.StdoutColorDisabled)
		)
// 这将在Windows操作系统中丢失颜色。请勿使用。
// 如果 _, err := os.Stdout.Write(input.getRealBuffer(true).Bytes()); 出现错误err，则不为nil {

		// 这将在Windows操作系统中打印颜色。
		if _, err = fmt.Fprint(color.Output, buffer.String()); err != nil {
			intlog.Errorf(ctx, `%+v`, err)
		}
		return buffer
	}
	return nil
}

// printToFile 将日志内容输出到磁盘文件。
func (l *Logger) printToFile(ctx context.Context, t time.Time, in *HandlerInput) *bytes.Buffer {
	var (
		buffer        = in.getRealBuffer(l.config.WriterColorEnable)
		logFilePath   = l.getFilePath(t)
		memoryLockKey = memoryLockPrefixForPrintingToFile + logFilePath
	)
	gmlock.Lock(memoryLockKey)
	defer gmlock.Unlock(memoryLockKey)

	// 旋转文件大小检查。
	if l.config.RotateSize > 0 && gfile.Size(logFilePath) > l.config.RotateSize {
		if runtime.GOOS == "windows" {
			file := l.createFpInPool(ctx, logFilePath)
			if file == nil {
				intlog.Errorf(ctx, `got nil file pointer for: %s`, logFilePath)
				return buffer
			}

			if _, err := file.Write(buffer.Bytes()); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}

			if err := file.Close(true); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
			l.rotateFileBySize(ctx, t)

			return buffer
		}

		l.rotateFileBySize(ctx, t)
	}
	// 将日志内容输出到磁盘文件。
	if file := l.createFpInPool(ctx, logFilePath); file == nil {
		intlog.Errorf(ctx, `got nil file pointer for: %s`, logFilePath)
	} else {
		if _, err := file.Write(buffer.Bytes()); err != nil {
			intlog.Errorf(ctx, `%+v`, err)
		}
		if err := file.Close(); err != nil {
			intlog.Errorf(ctx, `%+v`, err)
		}
	}
	return buffer
}

// createFpInPool 从文件池中获取并返回一个文件指针。
func (l *Logger) createFpInPool(ctx context.Context, path string) *gfpool.File {
	file, err := gfpool.Open(
		path,
		defaultFileFlags,
		defaultFilePerm,
		defaultFileExpire,
	)
	if err != nil {
		// panic(err)
		intlog.Errorf(ctx, `%+v`, err)
	}
	return file
}

// getFpFromPool 从文件池中获取并返回一个文件指针。
func (l *Logger) getFpFromPool(ctx context.Context, path string) *gfpool.File {
	file := gfpool.Get(
		path,
		defaultFileFlags,
		defaultFilePerm,
		defaultFileExpire,
	)
	if file == nil {
		intlog.Errorf(ctx, `can not find the file, path:%s`, path)
	}
	return file
}

// printStd 不带堆栈地打印内容`s`。
func (l *Logger) printStd(ctx context.Context, level int, values ...interface{}) {
	l.print(ctx, level, "", values...)
}

// printStd 在进行堆栈检查的情况下打印内容`s`。
func (l *Logger) printErr(ctx context.Context, level int, values ...interface{}) {
	var stack string
	if l.config.StStatus == 1 {
		stack = l.GetStack()
	}
	// 在顺序输出方面，此处不要使用 stderr，而应使用相同的 stdout。
	l.print(ctx, level, stack, values...)
}

// format 使用 fmt.Sprintf 对 `values` 进行格式化。
func (l *Logger) format(format string, values ...interface{}) string {
	return fmt.Sprintf(format, values...)
}

// PrintStack 打印调用栈，
// 可选参数 `skip` 指定了从终点开始需要跳过的堆栈偏移量。
func (l *Logger) PrintStack(ctx context.Context, skip ...int) {
	if s := l.GetStack(skip...); s != "" {
		l.Print(ctx, "Stack:\n"+s)
	} else {
		l.Print(ctx)
	}
}

// GetStack 返回调用堆栈的内容，
// 可选参数 `skip` 指定了从终点开始跳过的堆栈偏移量。
func (l *Logger) GetStack(skip ...int) string {
	stackSkip := l.config.StSkip
	if len(skip) > 0 {
		stackSkip += skip[0]
	}
	filters := []string{pathFilterKey}
	if l.config.StFilter != "" {
		filters = append(filters, l.config.StFilter)
	}
	// 是否过滤框架错误堆栈。
	if errors.IsStackModeBrief() {
		filters = append(filters, consts.StackFilterKeyForGoFrame)
	}
	return gdebug.StackWithFilters(filters, stackSkip)
}
