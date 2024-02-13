// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包 glog 实现了强大且易于使用的分级日志功能。
package 日志类

import (
	"context"
	
	"github.com/888go/goframe/internal/command"
	"github.com/888go/goframe/os/grpool"
	"github.com/888go/goframe/util/gconv"
)

// ILogger 是 logger 的 API 接口。
type ILogger interface {
	X输出(上下文 context.Context, 值 ...interface{})
	X输出并格式化(上下文 context.Context, 格式 string, 值 ...interface{})
	X输出DEBU(上下文 context.Context, 值 ...interface{})
	X输出并格式化DEBU(上下文 context.Context, 格式 string, 值 ...interface{})
	X输出INFO(上下文 context.Context, 值 ...interface{})
	X输出并格式化INFO(上下文 context.Context, 格式 string, 值 ...interface{})
	X输出NOTI(上下文 context.Context, 值 ...interface{})
	X输出并格式化NOTI(上下文 context.Context, 格式 string, 值 ...interface{})
	X输出WARN(上下文 context.Context, 值 ...interface{})
	X输出并格式化WARN(上下文 context.Context, 格式 string, 值 ...interface{})
	Error(上下文 context.Context, 值 ...interface{})
	X输出并格式化ERR(上下文 context.Context, 格式 string, 值 ...interface{})
	X输出CRIT(上下文 context.Context, 值 ...interface{})
	X输出并格式化CRIT(上下文 context.Context, 格式 string, 值 ...interface{})
	X输出PANI(上下文 context.Context, 值 ...interface{})
	X输出并格式化PANI(上下文 context.Context, 格式 string, 值 ...interface{})
	X输出FATA(上下文 context.Context, 值 ...interface{})
	X输出并格式化FATA(上下文 context.Context, 格式 string, 值 ...interface{})
}

const (
	commandEnvKeyForDebug = "gf.glog.debug"
)

var (
	// 确保 Logger 实现了 ILogger 接口。
	_ ILogger = &Logger{}

	// 默认日志器对象，用于包内方法的使用。
	defaultLogger = X创建()

// Goroutine 池用于异步日志输出。
// 它仅使用一个异步工作者以确保日志按序输出。
	asyncPool = 协程类.New(1)

// defaultDebug 默认是否开启调试级别，可以通过命令行选项或系统环境进行配置。
	defaultDebug = true
)

func init() {
	defaultDebug = 转换类.X取布尔(command.GetOptWithEnv(commandEnvKeyForDebug, "true"))
	X设置debug(defaultDebug)
}

// DefaultLogger 返回默认日志器。
func X取默认日志类() *Logger {
	return defaultLogger
}

// SetDefaultLogger 为 glog 包设置默认日志器。
// 注意，如果在不同 goroutine 中调用此函数可能存在并发安全问题。
func X设置默认日志类(l *Logger) {
	defaultLogger = l
}
