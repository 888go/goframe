// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包glog实现了强大且易于使用的等级化日志记录功能。 md5:b8685dc39c4dd154
package 日志类

import (
	"context"

	"github.com/888go/goframe/internal/command"
	grpool "github.com/888go/goframe/os/grpool"
	gconv "github.com/888go/goframe/util/gconv"
)

// ILogger是日志记录器的API接口。 md5:762449020563f6b9
type ILogger interface {
	X输出(ctx context.Context, v ...interface{})
	X输出并格式化(ctx context.Context, format string, v ...interface{})
	X输出DEBU(ctx context.Context, v ...interface{})
	X输出并格式化DEBU(ctx context.Context, format string, v ...interface{})
	X输出INFO(ctx context.Context, v ...interface{})
	X输出并格式化INFO(ctx context.Context, format string, v ...interface{})
	X输出NOTI(ctx context.Context, v ...interface{})
	X输出并格式化NOTI(ctx context.Context, format string, v ...interface{})
	X输出WARN(ctx context.Context, v ...interface{})
	X输出并格式化WARN(ctx context.Context, format string, v ...interface{})
	Error(ctx context.Context, v ...interface{})
	X输出并格式化ERR(ctx context.Context, format string, v ...interface{})
	X输出CRIT(ctx context.Context, v ...interface{})
	X输出并格式化CRIT(ctx context.Context, format string, v ...interface{})
	X输出PANI(ctx context.Context, v ...interface{})
	X输出并格式化PANI(ctx context.Context, format string, v ...interface{})
	X输出FATA(ctx context.Context, v ...interface{})
	X输出并格式化FATA(ctx context.Context, format string, v ...interface{})
}

const (
	commandEnvKeyForDebug = "gf.glog.debug"
)

var (
		// 确保 Logger 实现 ILogger 接口。 md5:451e2e42ba395d3a
	_ ILogger = &Logger{}

		// 默认的日志记录器对象，供包内方法使用。 md5:f3aa5266bd3b033f
	defaultLogger = X创建()

	// 用于异步日志输出的Goroutine池。
	// 它仅使用一个异步工作者来确保日志的顺序。
	// md5:b8cbf70a6cb430e0
	asyncPool = grpool.New(1)

	// defaultDebug 表示是否默认启用调试级别，这可以通过命令选项或系统环境进行配置。
	// md5:db02f93ae09ddc6a
	defaultDebug = true
)

func init() {
	defaultDebug = gconv.X取布尔(command.GetOptWithEnv(commandEnvKeyForDebug, "true"))
	X设置debug(defaultDebug)
}

// X取默认日志类 返回默认的logger。 md5:375e904736d75955
func X取默认日志类() *Logger {
	return defaultLogger
}

// X设置默认日志类 为 glog 包设置默认的日志记录器。
// 注意，如果在不同的 goroutine 中调用此函数，可能会存在并发安全问题。
// md5:acb1633d3882d5ab
func X设置默认日志类(l *Logger) {
	defaultLogger = l
}
