// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类

import (
	"context"
	"fmt"
	"os"
)

// Print 使用 fmt.Sprintln 函数打印变量 `v` 及其换行。参数 `v` 可以是多个变量。
func (l *Logger) X输出(上下文 context.Context, 值 ...interface{}) {
	l.printStd(上下文, LEVEL_NONE, 值...)
}

// Printf 通过使用 fmt.Sprintf 格式化方法打印变量 `v`。
// 参数 `v` 可以是多个变量。
func (l *Logger) X输出并格式化(上下文 context.Context, 格式 string, 值 ...interface{}) {
	l.printStd(上下文, LEVEL_NONE, l.format(格式, 值...))
}

// Fatal 打印带有[FATA]头部和换行符的日志内容，然后退出当前进程。
func (l *Logger) X输出FATA(上下文 context.Context, 值 ...interface{}) {
	l.printErr(上下文, LEVEL_FATA, 值...)
	os.Exit(1)
}

// Fatalf 打印日志内容，其头部为[FATA]，采用自定义格式，并在末尾添加换行符，然后退出当前进程。
func (l *Logger) X输出并格式化FATA(上下文 context.Context, 格式 string, 值 ...interface{}) {
	l.printErr(上下文, LEVEL_FATA, l.format(格式, 值...))
	os.Exit(1)
}

// Panic 会打印带有[PANI]头部和换行符的日志内容，然后触发 panic。
func (l *Logger) X输出PANI(上下文 context.Context, 值 ...interface{}) {
	l.printErr(上下文, LEVEL_PANI, 值...)
	panic(fmt.Sprint(值...))
}

// Panicf 函数会打印带有[PANI]头部、自定义格式以及换行符的日志内容，然后触发 panic。
func (l *Logger) X输出并格式化PANI(上下文 context.Context, 格式 string, 值 ...interface{}) {
	l.printErr(上下文, LEVEL_PANI, l.format(格式, 值...))
	panic(l.format(格式, 值...))
}

// Info打印日志内容，前面带有[INFO]头部和换行符。
func (l *Logger) X输出INFO(上下文 context.Context, 值 ...interface{}) {
	if l.checkLevel(LEVEL_INFO) {
		l.printStd(上下文, LEVEL_INFO, 值...)
	}
}

// Infof 格式化并打印日志内容，带有 [INFO] 标头、自定义格式及换行。
// ```go
// Infof 根据提供的格式和参数，以 [INFO] 头部形式输出格式化信息，并在末尾添加换行。
// 示例用法：
// log.Infof("用户 %s 登录成功", username)
func (l *Logger) X输出并格式化INFO(上下文 context.Context, 格式 string, 值 ...interface{}) {
	if l.checkLevel(LEVEL_INFO) {
		l.printStd(上下文, LEVEL_INFO, l.format(格式, 值...))
	}
}

// Debug 以 [DEBU] 标头和换行符打印日志内容。
func (l *Logger) X输出DEBU(上下文 context.Context, 值 ...interface{}) {
	if l.checkLevel(LEVEL_DEBU) {
		l.printStd(上下文, LEVEL_DEBU, 值...)
	}
}

// Debugf按照[DEBU]头部，自定义格式及换行打印日志内容。
func (l *Logger) X输出并格式化DEBU(上下文 context.Context, 格式 string, 值 ...interface{}) {
	if l.checkLevel(LEVEL_DEBU) {
		l.printStd(上下文, LEVEL_DEBU, l.format(格式, 值...))
	}
}

// Notice 以 [NOTI] 标头和换行符打印日志内容。
// 如果启用了堆栈功能，还会打印调用堆栈信息。
func (l *Logger) X输出NOTI(上下文 context.Context, 值 ...interface{}) {
	if l.checkLevel(LEVEL_NOTI) {
		l.printStd(上下文, LEVEL_NOTI, 值...)
	}
}

// Noticef函数以[NOTI]头、自定义格式及换行符打印日志内容。
// 如果启用了堆栈追踪功能，它还会打印调用堆栈信息。
func (l *Logger) X输出并格式化NOTI(上下文 context.Context, 格式 string, 值 ...interface{}) {
	if l.checkLevel(LEVEL_NOTI) {
		l.printStd(上下文, LEVEL_NOTI, l.format(格式, 值...))
	}
}

// Warning 以 [WARN] 标头和换行符打印日志内容。
// 如果启用了堆栈功能，它还会打印调用者堆栈信息。
func (l *Logger) X输出WARN(上下文 context.Context, 值 ...interface{}) {
	if l.checkLevel(LEVEL_WARN) {
		l.printStd(上下文, LEVEL_WARN, 值...)
	}
}

// Warningf函数会打印带有[WARN]头的记录内容，使用自定义格式并添加换行符。
// 如果启用了堆栈功能，它还会打印调用堆栈信息。
func (l *Logger) X输出并格式化WARN(上下文 context.Context, 格式 string, 值 ...interface{}) {
	if l.checkLevel(LEVEL_WARN) {
		l.printStd(上下文, LEVEL_WARN, l.format(格式, 值...))
	}
}

// Error 以 [ERRO] 标头和换行符打印日志内容。
// 如果启用了堆栈追踪功能，它还会打印调用者堆栈信息。
func (l *Logger) Error(上下文 context.Context, 值 ...interface{}) {
	if l.checkLevel(LEVEL_ERRO) {
		l.printErr(上下文, LEVEL_ERRO, 值...)
	}
}

// Errorf 函数以 [ERRO] 标头、自定义格式及换行符打印日志内容。
// 若启用了堆栈追踪功能，它还会打印调用者堆栈信息。
func (l *Logger) X输出并格式化ERR(上下文 context.Context, 格式 string, 值 ...interface{}) {
	if l.checkLevel(LEVEL_ERRO) {
		l.printErr(上下文, LEVEL_ERRO, l.format(格式, 值...))
	}
}

// Critical 以 [CRIT] 标头和换行符打印日志内容。
// 如果启用了堆栈追踪功能，还会打印调用堆栈信息。
func (l *Logger) X输出CRIT(上下文 context.Context, 值 ...interface{}) {
	if l.checkLevel(LEVEL_CRIT) {
		l.printErr(上下文, LEVEL_CRIT, 值...)
	}
}

// Criticalf以[CRIT]头部、自定义格式和换行符打印日志内容。
// 如果启用了堆栈追踪功能，它还会打印调用者堆栈信息。
func (l *Logger) X输出并格式化CRIT(上下文 context.Context, 格式 string, 值 ...interface{}) {
	if l.checkLevel(LEVEL_CRIT) {
		l.printErr(上下文, LEVEL_CRIT, l.format(格式, 值...))
	}
}

// checkLevel 检查给定的 `level` 是否可以输出。
func (l *Logger) checkLevel(level int) bool {
	return l.config.X级别&level > 0
}
