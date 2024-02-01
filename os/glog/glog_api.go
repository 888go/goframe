// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package glog
import (
	"context"
	)
// Print 使用 fmt.Sprintln 函数打印变量 `v` 及其换行。参数 `v` 可以是多个变量。
func Print(ctx context.Context, v ...interface{}) {
	defaultLogger.Print(ctx, v...)
}

// Printf 通过使用 fmt.Sprintf 格式化方法打印变量 `v`。
// 参数 `v` 可以是多个变量。
func Printf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Printf(ctx, format, v...)
}

// Fatal 打印带有[FATA]头部和换行符的日志内容，然后退出当前进程。
func Fatal(ctx context.Context, v ...interface{}) {
	defaultLogger.Fatal(ctx, v...)
}

// Fatalf 打印日志内容，其头部为[FATA]，采用自定义格式，并在末尾添加换行符，然后退出当前进程。
func Fatalf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Fatalf(ctx, format, v...)
}

// Panic 会打印带有[PANI]头部和换行符的日志内容，然后触发 panic。
func Panic(ctx context.Context, v ...interface{}) {
	defaultLogger.Panic(ctx, v...)
}

// Panicf 函数会打印带有[PANI]头部、自定义格式以及换行符的日志内容，然后触发 panic。
func Panicf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Panicf(ctx, format, v...)
}

// Info打印日志内容，前面带有[INFO]头部和换行符。
func Info(ctx context.Context, v ...interface{}) {
	defaultLogger.Info(ctx, v...)
}

// Infof 格式化并打印日志内容，带有 [INFO] 标头、自定义格式及换行。
// ```go
// Infof 根据提供的格式和参数，以 [INFO] 头部形式输出格式化信息，并在末尾添加换行。
// 示例用法：
// log.Infof("用户 %s 登录成功", username)
func Infof(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Infof(ctx, format, v...)
}

// Debug 以 [DEBU] 标头和换行符打印日志内容。
func Debug(ctx context.Context, v ...interface{}) {
	defaultLogger.Debug(ctx, v...)
}

// Debugf按照[DEBU]头部，自定义格式及换行打印日志内容。
func Debugf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Debugf(ctx, format, v...)
}

// Notice 以 [NOTI] 标头和换行符打印日志内容。
// 如果启用了堆栈功能，还会打印调用堆栈信息。
func Notice(ctx context.Context, v ...interface{}) {
	defaultLogger.Notice(ctx, v...)
}

// Noticef函数以[NOTI]头、自定义格式及换行符打印日志内容。
// 如果启用了堆栈追踪功能，它还会打印调用堆栈信息。
func Noticef(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Noticef(ctx, format, v...)
}

// Warning 以 [WARN] 标头和换行符打印日志内容。
// 如果启用了堆栈功能，它还会打印调用者堆栈信息。
func Warning(ctx context.Context, v ...interface{}) {
	defaultLogger.Warning(ctx, v...)
}

// Warningf函数会打印带有[WARN]头的记录内容，使用自定义格式并添加换行符。
// 如果启用了堆栈功能，它还会打印调用堆栈信息。
func Warningf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Warningf(ctx, format, v...)
}

// Error 以 [ERRO] 标头和换行符打印日志内容。
// 如果启用了堆栈追踪功能，它还会打印调用者堆栈信息。
func Error(ctx context.Context, v ...interface{}) {
	defaultLogger.Error(ctx, v...)
}

// Errorf 函数以 [ERRO] 标头、自定义格式及换行符打印日志内容。
// 若启用了堆栈追踪功能，它还会打印调用者堆栈信息。
func Errorf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Errorf(ctx, format, v...)
}

// Critical 以 [CRIT] 标头和换行符打印日志内容。
// 如果启用了堆栈追踪功能，还会打印调用堆栈信息。
func Critical(ctx context.Context, v ...interface{}) {
	defaultLogger.Critical(ctx, v...)
}

// Criticalf以[CRIT]头部、自定义格式和换行符打印日志内容。
// 如果启用了堆栈追踪功能，它还会打印调用者堆栈信息。
func Criticalf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Criticalf(ctx, format, v...)
}
