// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package glog
import (
	"context"
	"fmt"
	"os"
	)
// Print 使用 fmt.Sprintln 函数打印变量 `v` 及其换行。参数 `v` 可以是多个变量。
func (l *Logger) Print(ctx context.Context, v ...interface{}) {
	l.printStd(ctx, LEVEL_NONE, v...)
}

// Printf 通过使用 fmt.Sprintf 格式化方法打印变量 `v`。
// 参数 `v` 可以是多个变量。
func (l *Logger) Printf(ctx context.Context, format string, v ...interface{}) {
	l.printStd(ctx, LEVEL_NONE, l.format(format, v...))
}

// Fatal 打印带有[FATA]头部和换行符的日志内容，然后退出当前进程。
func (l *Logger) Fatal(ctx context.Context, v ...interface{}) {
	l.printErr(ctx, LEVEL_FATA, v...)
	os.Exit(1)
}

// Fatalf 打印日志内容，其头部为[FATA]，采用自定义格式，并在末尾添加换行符，然后退出当前进程。
func (l *Logger) Fatalf(ctx context.Context, format string, v ...interface{}) {
	l.printErr(ctx, LEVEL_FATA, l.format(format, v...))
	os.Exit(1)
}

// Panic 会打印带有[PANI]头部和换行符的日志内容，然后触发 panic。
func (l *Logger) Panic(ctx context.Context, v ...interface{}) {
	l.printErr(ctx, LEVEL_PANI, v...)
	panic(fmt.Sprint(v...))
}

// Panicf 函数会打印带有[PANI]头部、自定义格式以及换行符的日志内容，然后触发 panic。
func (l *Logger) Panicf(ctx context.Context, format string, v ...interface{}) {
	l.printErr(ctx, LEVEL_PANI, l.format(format, v...))
	panic(l.format(format, v...))
}

// Info打印日志内容，前面带有[INFO]头部和换行符。
func (l *Logger) Info(ctx context.Context, v ...interface{}) {
	if l.checkLevel(LEVEL_INFO) {
		l.printStd(ctx, LEVEL_INFO, v...)
	}
}

// Infof 格式化并打印日志内容，带有 [INFO] 标头、自定义格式及换行。
// ```go
// Infof 根据提供的格式和参数，以 [INFO] 头部形式输出格式化信息，并在末尾添加换行。
// 示例用法：
// log.Infof("用户 %s 登录成功", username)
func (l *Logger) Infof(ctx context.Context, format string, v ...interface{}) {
	if l.checkLevel(LEVEL_INFO) {
		l.printStd(ctx, LEVEL_INFO, l.format(format, v...))
	}
}

// Debug 以 [DEBU] 标头和换行符打印日志内容。
func (l *Logger) Debug(ctx context.Context, v ...interface{}) {
	if l.checkLevel(LEVEL_DEBU) {
		l.printStd(ctx, LEVEL_DEBU, v...)
	}
}

// Debugf按照[DEBU]头部，自定义格式及换行打印日志内容。
func (l *Logger) Debugf(ctx context.Context, format string, v ...interface{}) {
	if l.checkLevel(LEVEL_DEBU) {
		l.printStd(ctx, LEVEL_DEBU, l.format(format, v...))
	}
}

// Notice 以 [NOTI] 标头和换行符打印日志内容。
// 如果启用了堆栈功能，还会打印调用堆栈信息。
func (l *Logger) Notice(ctx context.Context, v ...interface{}) {
	if l.checkLevel(LEVEL_NOTI) {
		l.printStd(ctx, LEVEL_NOTI, v...)
	}
}

// Noticef函数以[NOTI]头、自定义格式及换行符打印日志内容。
// 如果启用了堆栈追踪功能，它还会打印调用堆栈信息。
func (l *Logger) Noticef(ctx context.Context, format string, v ...interface{}) {
	if l.checkLevel(LEVEL_NOTI) {
		l.printStd(ctx, LEVEL_NOTI, l.format(format, v...))
	}
}

// Warning 以 [WARN] 标头和换行符打印日志内容。
// 如果启用了堆栈功能，它还会打印调用者堆栈信息。
func (l *Logger) Warning(ctx context.Context, v ...interface{}) {
	if l.checkLevel(LEVEL_WARN) {
		l.printStd(ctx, LEVEL_WARN, v...)
	}
}

// Warningf函数会打印带有[WARN]头的记录内容，使用自定义格式并添加换行符。
// 如果启用了堆栈功能，它还会打印调用堆栈信息。
func (l *Logger) Warningf(ctx context.Context, format string, v ...interface{}) {
	if l.checkLevel(LEVEL_WARN) {
		l.printStd(ctx, LEVEL_WARN, l.format(format, v...))
	}
}

// Error 以 [ERRO] 标头和换行符打印日志内容。
// 如果启用了堆栈追踪功能，它还会打印调用者堆栈信息。
func (l *Logger) Error(ctx context.Context, v ...interface{}) {
	if l.checkLevel(LEVEL_ERRO) {
		l.printErr(ctx, LEVEL_ERRO, v...)
	}
}

// Errorf 函数以 [ERRO] 标头、自定义格式及换行符打印日志内容。
// 若启用了堆栈追踪功能，它还会打印调用者堆栈信息。
func (l *Logger) Errorf(ctx context.Context, format string, v ...interface{}) {
	if l.checkLevel(LEVEL_ERRO) {
		l.printErr(ctx, LEVEL_ERRO, l.format(format, v...))
	}
}

// Critical 以 [CRIT] 标头和换行符打印日志内容。
// 如果启用了堆栈追踪功能，还会打印调用堆栈信息。
func (l *Logger) Critical(ctx context.Context, v ...interface{}) {
	if l.checkLevel(LEVEL_CRIT) {
		l.printErr(ctx, LEVEL_CRIT, v...)
	}
}

// Criticalf以[CRIT]头部、自定义格式和换行符打印日志内容。
// 如果启用了堆栈追踪功能，它还会打印调用者堆栈信息。
func (l *Logger) Criticalf(ctx context.Context, format string, v ...interface{}) {
	if l.checkLevel(LEVEL_CRIT) {
		l.printErr(ctx, LEVEL_CRIT, l.format(format, v...))
	}
}

// checkLevel 检查给定的 `level` 是否可以输出。
func (l *Logger) checkLevel(level int) bool {
	return l.config.Level&level > 0
}
