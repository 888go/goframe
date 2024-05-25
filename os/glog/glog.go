// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包glog实现了强大且易于使用的等级化日志记录功能。. md5:b8685dc39c4dd154
package glog

import (
	"context"

	"github.com/gogf/gf/v2/internal/command"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/util/gconv"
)

// ILogger是日志记录器的API接口。. md5:762449020563f6b9
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
	// 确保 Logger 实现 ILogger 接口。. md5:451e2e42ba395d3a
	_ ILogger = &Logger{}

	// 默认的日志记录器对象，供包内方法使用。. md5:f3aa5266bd3b033f
	defaultLogger = New()

// 用于异步日志输出的Goroutine池。
// 它仅使用一个异步工作者来确保日志的顺序。
// md5:b8cbf70a6cb430e0
	asyncPool = grpool.New(1)

// defaultDebug 表示是否默认启用调试级别，这可以通过命令选项或系统环境进行配置。
// md5:db02f93ae09ddc6a
	defaultDebug = true
)

func init() {
	defaultDebug = gconv.Bool(command.GetOptWithEnv(commandEnvKeyForDebug, "true"))
	SetDebug(defaultDebug)
}

// DefaultLogger 返回默认的logger。. md5:375e904736d75955
func DefaultLogger() *Logger {
	return defaultLogger
}

// SetDefaultLogger 为 glog 包设置默认的日志记录器。
// 注意，如果在不同的 goroutine 中调用此函数，可能会存在并发安全问题。
// md5:acb1633d3882d5ab
func SetDefaultLogger(l *Logger) {
	defaultLogger = l
}
