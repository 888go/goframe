// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包 glog 实现了强大且易于使用的分级日志功能。
package glog

import (
	"context"
	
	"github.com/888go/goframe/internal/command"
	"github.com/888go/goframe/os/grpool"
	"github.com/888go/goframe/util/gconv"
)

// ILogger 是 logger 的 API 接口。
type ILogger interface {
	Print(ctx context.Context, v ...interface{})
	Printf(ctx context.Context, format string, v ...interface{})
	Debug(ctx context.Context, v ...interface{})
	Debugf(ctx context.Context, format string, v ...interface{})
	Info(ctx context.Context, v ...interface{})
	Infof(ctx context.Context, format string, v ...interface{})
	Notice(ctx context.Context, v ...interface{})
	Noticef(ctx context.Context, format string, v ...interface{})
	Warning(ctx context.Context, v ...interface{})
	Warningf(ctx context.Context, format string, v ...interface{})
	Error(ctx context.Context, v ...interface{})
	Errorf(ctx context.Context, format string, v ...interface{})
	Critical(ctx context.Context, v ...interface{})
	Criticalf(ctx context.Context, format string, v ...interface{})
	Panic(ctx context.Context, v ...interface{})
	Panicf(ctx context.Context, format string, v ...interface{})
	Fatal(ctx context.Context, v ...interface{})
	Fatalf(ctx context.Context, format string, v ...interface{})
}

const (
	commandEnvKeyForDebug = "gf.glog.debug"
)

var (
	// 确保 Logger 实现了 ILogger 接口。
	_ ILogger = &Logger{}

	// 默认日志器对象，用于包内方法的使用。
	defaultLogger = New()

// Goroutine 池用于异步日志输出。
// 它仅使用一个异步工作者以确保日志按序输出。
	asyncPool = grpool.New(1)

// defaultDebug 默认是否开启调试级别，可以通过命令行选项或系统环境进行配置。
	defaultDebug = true
)

func init() {
	defaultDebug = gconv.Bool(command.GetOptWithEnv(commandEnvKeyForDebug, "true"))
	SetDebug(defaultDebug)
}

// DefaultLogger 返回默认日志器。
func DefaultLogger() *Logger {
	return defaultLogger
}

// SetDefaultLogger 为 glog 包设置默认日志器。
// 注意，如果在不同 goroutine 中调用此函数可能存在并发安全问题。
func SetDefaultLogger(l *Logger) {
	defaultLogger = l
}
