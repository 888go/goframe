// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 日志类

import (
	"context"
	"fmt"
	"os"
)

// X输出 使用 fmt.Sprintln 函数打印 `v` 并添加换行符。
// 参数 `v` 可以是多个变量。
// md5:6c0b3b96234f77ce
func (l *Logger) X输出(上下文 context.Context, 值 ...interface{}) {
	l.printStd(上下文, LEVEL_NONE, 值...)
}

// X输出并格式化 使用 fmt.Sprintf 根据格式 `format` 打印 `v`。
// 参数 `v` 可以是多个变量。
// md5:e3b6ab3d8750ad4c
func (l *Logger) X输出并格式化(上下文 context.Context, 格式 string, 值 ...interface{}) {
	l.printStd(上下文, LEVEL_NONE, l.format(格式, 值...))
}

// X输出FATA 以[FATA]标题和换行符打印日志内容，然后退出当前进程。 md5:4061b6551d7137a1
func (l *Logger) X输出FATA(上下文 context.Context, 值 ...interface{}) {
	l.printErr(上下文, LEVEL_FATA, 值...)
	os.Exit(1)
}

// X输出并格式化FATA 打印带有 [FATA] 头部、自定义格式和换行符的日志内容，然后退出当前进程。 md5:cbaf3fb7e2b92df9
func (l *Logger) X输出并格式化FATA(上下文 context.Context, 格式 string, 值 ...interface{}) {
	l.printErr(上下文, LEVEL_FATA, l.format(格式, 值...))
	os.Exit(1)
}

// X输出PANI 在输出带有 [PANI] 头部和换行符的日志内容后，引发 panic。 md5:bcde7bf5ff45073a
func (l *Logger) X输出PANI(上下文 context.Context, 值 ...interface{}) {
	l.printErr(上下文, LEVEL_PANI, 值...)
	panic(fmt.Sprint(值...))
}

// X输出并格式化PANI 函数打印带有 [PANI] 标头、自定义格式和换行的日志内容，然后触发恐慌（panic）。 md5:bf1df3ad28ded71f
func (l *Logger) X输出并格式化PANI(上下文 context.Context, 格式 string, 值 ...interface{}) {
	l.printErr(上下文, LEVEL_PANI, l.format(格式, 值...))
	panic(l.format(格式, 值...))
}

// X输出INFO 打印带有 "[INFO]" 标头和换行符的日志内容。 md5:1ed8e917ecca5ef4
func (l *Logger) X输出INFO(上下文 context.Context, 值 ...interface{}) {
	if l.checkLevel(LEVEL_INFO) {
		l.printStd(上下文, LEVEL_INFO, 值...)
	}
}

// X输出并格式化INFO 打印带有 [INFO] 标头、自定义格式和换行符的日志内容。 md5:fda1e57b2e2089d7
func (l *Logger) X输出并格式化INFO(上下文 context.Context, 格式 string, 值 ...interface{}) {
	if l.checkLevel(LEVEL_INFO) {
		l.printStd(上下文, LEVEL_INFO, l.format(格式, 值...))
	}
}

// X输出DEBU 打印带有 [DEBU] 标头和换行符的日志内容。 md5:7a64f5ebf48d4f92
func (l *Logger) X输出DEBU(上下文 context.Context, 值 ...interface{}) {
	if l.checkLevel(LEVEL_DEBU) {
		l.printStd(上下文, LEVEL_DEBU, 值...)
	}
}

// X输出并格式化DEBU 函数打印带有 [DEBU] 标头、自定义格式化内容和换行符的日志信息。 md5:0c0164b88b59a40c
func (l *Logger) X输出并格式化DEBU(上下文 context.Context, 格式 string, 值 ...interface{}) {
	if l.checkLevel(LEVEL_DEBU) {
		l.printStd(上下文, LEVEL_DEBU, l.format(格式, 值...))
	}
}

// ```go
// X输出NOTI 打印日志内容，前缀为 [NOTI] 并在末尾添加换行符。
// 如果启用了堆栈跟踪功能，它还会打印调用者堆栈信息。
// ```
// md5:c36d548c618d1251
func (l *Logger) X输出NOTI(上下文 context.Context, 值 ...interface{}) {
	if l.checkLevel(LEVEL_NOTI) {
		l.printStd(上下文, LEVEL_NOTI, 值...)
	}
}

// X输出并格式化NOTI 打印带有 [NOTI] 标头的日志内容，自定义格式和换行符。如果启用了堆栈功能，它还会打印调用者堆栈信息。
// md5:f9d4f5af91856cd9
func (l *Logger) X输出并格式化NOTI(上下文 context.Context, 格式 string, 值 ...interface{}) {
	if l.checkLevel(LEVEL_NOTI) {
		l.printStd(上下文, LEVEL_NOTI, l.format(格式, 值...))
	}
}

// X输出WARN 打印带有 [WARN] 头部和换行符的日志内容。如果启用了堆栈功能，它还会打印调用者堆栈信息。
// md5:8142c86f6be53ee0
func (l *Logger) X输出WARN(上下文 context.Context, 值 ...interface{}) {
	if l.checkLevel(LEVEL_WARN) {
		l.printStd(上下文, LEVEL_WARN, 值...)
	}
}

// X输出并格式化WARN 函数打印带有 [WARN] 标头的记录内容、自定义格式化字符串以及换行。
// 如果启用了堆栈追踪功能，它还会打印调用者堆栈信息。
// md5:025f0baa4a1f8600
func (l *Logger) X输出并格式化WARN(上下文 context.Context, 格式 string, 值 ...interface{}) {
	if l.checkLevel(LEVEL_WARN) {
		l.printStd(上下文, LEVEL_WARN, l.format(格式, 值...))
	}
}

// Error 打印带有 [ERRO] 标头和换行符的日志内容。
// 如果启用了堆栈跟踪功能，它还会打印调用者堆栈信息。
// md5:f2aa6f6c0e4d2061
func (l *Logger) Error(上下文 context.Context, 值 ...interface{}) {
	if l.checkLevel(LEVEL_ERRO) {
		l.printErr(上下文, LEVEL_ERRO, 值...)
	}
}

// X输出并格式化ERR 函数会打印带有 [ERRO] 标头的日志内容，使用自定义格式并添加换行符。如果启用了堆栈功能，它还会打印调用者堆栈信息。
// md5:4a90789d1de07943
func (l *Logger) X输出并格式化ERR(上下文 context.Context, 格式 string, 值 ...interface{}) {
	if l.checkLevel(LEVEL_ERRO) {
		l.printErr(上下文, LEVEL_ERRO, l.format(格式, 值...))
	}
}

// X输出CRIT 打印带有 [CRIT] 头部和换行符的日志内容。如果启用了堆栈功能，它还会打印调用者堆栈信息。
// md5:f9fb0900ff8f602f
func (l *Logger) X输出CRIT(上下文 context.Context, 值 ...interface{}) {
	if l.checkLevel(LEVEL_CRIT) {
		l.printErr(上下文, LEVEL_CRIT, 值...)
	}
}

// X输出并格式化CRIT 函数打印带有 [CRIT] 标头、自定义格式和换行的日志内容。
// 如果启用了堆栈追踪功能，它还会打印调用者堆栈信息。
// md5:fa381bbe7b0465d0
func (l *Logger) X输出并格式化CRIT(上下文 context.Context, 格式 string, 值 ...interface{}) {
	if l.checkLevel(LEVEL_CRIT) {
		l.printErr(上下文, LEVEL_CRIT, l.format(格式, 值...))
	}
}

// checkLevel 检查给定的 `level` 是否可以输出。 md5:59e82f73882a5ac4
func (l *Logger) checkLevel(level int) bool {
	return l.config.Level&level > 0
}
