// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

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

	"github.com/gogf/gf/v2/debug/gdebug"
	"github.com/gogf/gf/v2/internal/consts"
	"github.com/gogf/gf/v2/internal/errors"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gfpool"
	"github.com/gogf/gf/v2/os/gmlock"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gconv"
)

// Logger 是用于日志管理的结构体。 md5:c338807dca943ae3
type Logger struct {
	parent *Logger // 父日志记录器，如果非空，表示该日志记录器用于链式调用功能。 md5:9efa0349702d1c2c
	config Config  // Logger configuration.
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
	F_ASYNC      = 1 << iota // 异步打印日志内容。 md5:aae4973631499f41
	F_FILE_LONG              // 打印完整的文件名和行号：/a/b/c/d.go:23。 md5:b0d552af751b5a59
	F_FILE_SHORT             // 打印最终的文件名元素和行号：d.go:23。会覆盖F_FILE_LONG设置。 md5:b2f804c31e821aba
	F_TIME_DATE              // 在本地时区打印日期：2009-01-23。 md5:d7310c166de59388
	F_TIME_TIME              // 在本地时区打印时间：01:23:23。 md5:547ab7e84bd2a146
	F_TIME_MILLI             // 在本地时区打印带有毫秒的时间：01:23:23.675。 md5:063a060ef145f39d
	F_CALLER_FN              // 打印调用者函数名称和包：main.main. md5:1d729cf0b4dce445
	F_TIME_STD   = F_TIME_DATE | F_TIME_MILLI
)

// New 创建并返回一个自定义的日志器。 md5:0205650422cdd95e
func New() *Logger {
	return &Logger{
		config: DefaultConfig(),
	}
}

// NewWithWriter 创建并返回一个具有 io.Writer 的自定义日志器。 md5:51edfcbd62ded572
func NewWithWriter(writer io.Writer) *Logger {
	l := New()
	l.SetWriter(writer)
	return l
}

// Clone 返回一个新的记录器，它是当前记录器的`浅拷贝`。
// 注意，克隆体的`config`属性是对当前记录器配置的浅拷贝。
// md5:c70ded0c6903f4be
func (l *Logger) Clone() *Logger {
	return &Logger{
		config: l.config,
		parent: l,
	}
}

// getFilePath 返回日志文件的路径。
// 日志文件的名称必须带有".log"扩展名。
// md5:f3fb565d6de6de8e
func (l *Logger) getFilePath(now time.Time) string {
		// 文件名中包含“{}”的内容将使用gtime进行格式化。 md5:9502dc758bde7fca
	file, _ := gregex.ReplaceStringFunc(`{.+?}`, l.config.File, func(s string) string {
		return gtime.New(now).Format(strings.Trim(s, "{}"))
	})
	file = gfile.Join(l.config.Path, file)
	return file
}

// print 将`s`打印到已定义的writer（日志文件）或传递的`std`。 md5:2368d31e4b600609
func (l *Logger) print(ctx context.Context, level int, stack string, values ...any) {
	// 延迟初始化旋转特性。
	// 它采用原子读取操作来增强性能检查。
	// 此处使用CAP以确保性能和并发安全性。
	// 每个日志器仅初始化一次。
	// md5:1562dbed8f576bc2
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

		// 调用者路径和函数名。 md5:92543c6732cddd2d
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

		// 将值转换为字符串。 md5:c172ad8437ce5ebf
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

// doFinalPrint 根据配置输出日志内容。 md5:d442b45c82ee936e
func (l *Logger) doFinalPrint(ctx context.Context, input *HandlerInput) *bytes.Buffer {
	var buffer *bytes.Buffer
		// 允许输出到标准输出（stdout）吗？. md5:0f9801ce26bfb625
	if l.config.StdoutPrint {
		if buf := l.printToStdout(ctx, input); buf != nil {
			buffer = buf
		}
	}

		// 将内容输出到磁盘文件。 md5:eac9c01dcbb64a4f
	if l.config.Path != "" {
		if buf := l.printToFile(ctx, input.Time, input); buf != nil {
			buffer = buf
		}
	}

	// Used custom writer.
	if l.config.Writer != nil {
				// 向自定义写入器输出。 md5:ef7c09b52f5d355b
		if buf := l.printToWriter(ctx, input); buf != nil {
			buffer = buf
		}
	}
	return buffer
}

// printToWriter 将缓冲区写入writer。 md5:5b62d7e007bf0275
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

// printToStdout 将日志内容输出到标准输出（stdout）。 md5:2fb96a5229ec7af7
func (l *Logger) printToStdout(ctx context.Context, input *HandlerInput) *bytes.Buffer {
	if l.config.StdoutPrint {
		var (
			err    error
			buffer = input.getRealBuffer(!l.config.StdoutColorDisabled)
		)
		// 这将在Windows操作系统中丢失颜色信息。请勿使用。
		// if _, err := os.Stdout.Write(input.getRealBuffer(true).Bytes()); err != nil {
		// md5:29dd90df2339a223

				// 这将在Windows操作系统中打印颜色。 md5:c2abebac838c5747
		if _, err = fmt.Fprint(color.Output, buffer.String()); err != nil {
			intlog.Errorf(ctx, `%+v`, err)
		}
		return buffer
	}
	return nil
}

// printToFile 将日志内容输出到磁盘文件中。 md5:20d4379608c45b69
func (l *Logger) printToFile(ctx context.Context, t time.Time, in *HandlerInput) *bytes.Buffer {
	var (
		buffer        = in.getRealBuffer(l.config.WriterColorEnable)
		logFilePath   = l.getFilePath(t)
		memoryLockKey = memoryLockPrefixForPrintingToFile + logFilePath
	)
	gmlock.Lock(memoryLockKey)
	defer gmlock.Unlock(memoryLockKey)

		// 旋转文件大小检查。 md5:82f7b948ac1657a5
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
		// 将日志内容输出到磁盘文件。 md5:c3b494b8895acc38
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

// createFpInPool 从文件池中检索并返回一个文件指针。 md5:4acfc3ca923fca99
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

// getFpFromPool 从文件池中检索并返回一个文件指针。 md5:e3cccf00ba4439d4
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

// printStd 在不打印调用栈的情况下输出内容 `s`。 md5:f9e76daf3a16b514
func (l *Logger) printStd(ctx context.Context, level int, values ...interface{}) {
	l.print(ctx, level, "", values...)
}

// printErr 打印内容 `s`，并检查堆栈信息。 md5:d9883612e4e85bcb
func (l *Logger) printErr(ctx context.Context, level int, values ...interface{}) {
	var stack string
	if l.config.StStatus == 1 {
		stack = l.GetStack()
	}
		// 从顺序上来说，这里不要使用stderr，而是要使用相同的stdout。 md5:1f1258ae1ca0856e
	l.print(ctx, level, stack, values...)
}

// format 使用fmt.Sprintf格式化`values`。 md5:bd4bb80582842100
func (l *Logger) format(format string, values ...interface{}) string {
	return fmt.Sprintf(format, values...)
}

// PrintStack 打印调用堆栈，
// 可选参数 `skip` 用于指定从堆栈终点开始忽略的偏移量。
// md5:ef6cd40820765783
func (l *Logger) PrintStack(ctx context.Context, skip ...int) {
	if s := l.GetStack(skip...); s != "" {
		l.Print(ctx, "Stack:\n"+s)
	} else {
		l.Print(ctx)
	}
}

// GetStack 返回调用者栈的内容，
// 可选参数 `skip` 指定从终点开始要跳过的栈偏移量。
// md5:13592be3061e779d
func (l *Logger) GetStack(skip ...int) string {
	stackSkip := l.config.StSkip
	if len(skip) > 0 {
		stackSkip += skip[0]
	}
	filters := []string{pathFilterKey}
	if l.config.StFilter != "" {
		filters = append(filters, l.config.StFilter)
	}
		// 是否过滤框架错误堆栈。 md5:7cf82960065281e9
	if errors.IsStackModeBrief() {
		filters = append(filters, consts.StackFilterKeyForGoFrame)
	}
	return gdebug.StackWithFilters(filters, stackSkip)
}
