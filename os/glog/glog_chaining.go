// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类

import (
	"io"
)

// Expose 返回 glog 包的默认日志器。
// DefaultLogger()  方法的别名
func Expose别名() *Logger {
	return defaultLogger
}

// To 是一个链式函数，
// 它将当前日志内容输出重定向到指定的 `writer`。
func X重定向输出(writer io.Writer) *Logger {
	return defaultLogger.X重定向输出(writer)
}

// Path 是一个链式函数，
// 它用于将当前日志内容输出的目标目录路径设置为 `path`。
func X文件路径(文件路径 string) *Logger {
	return defaultLogger.X文件路径(文件路径)
}

// Cat 是一个链式函数，
// 它用于将当前日志内容输出的类别设置为 `category`。
func X文件分类(类名称 string) *Logger {
	return defaultLogger.X文件分类(类名称)
}

// File 是一个链式函数，
// 用于设置当前日志内容输出的文件名模式。
func X文件名格式(文件名格式 string) *Logger {
	return defaultLogger.X文件名格式(文件名格式)
}

// Level 是一个链式函数，
// 用于设置当前日志内容输出的记录级别。
func X级别(级别 int) *Logger {
	return defaultLogger.X级别(级别)
}

// LevelStr 是一个链式函数，
// 通过级别字符串设置当前日志内容输出的记录级别。
func X文本级别(文本级别 string) *Logger {
	return defaultLogger.X文本级别(文本级别)
}

// Skip 是一个链式函数，
// 用于设置当前日志内容输出时的堆栈跳过级别。
// 当启用行号打印时，它也会影响调用文件路径的检查。
func X堆栈偏移量(偏移量 int) *Logger {
	return defaultLogger.X堆栈偏移量(偏移量)
}

// Stack 是一个链式函数，
// 用于为当前日志内容输出设置堆栈选项。
func X堆栈选项(开启 bool, 偏移量 ...int) *Logger {
	return defaultLogger.X堆栈选项(开启, 偏移量...)
}

// StackWithFilter 是一个链式函数，
// 用于为当前日志内容输出设置堆栈过滤器。
func X堆栈过滤(过滤器 string) *Logger {
	return defaultLogger.X堆栈过滤(过滤器)
}

// Stdout 是一个链式函数，
// 用于启用/禁用当前日志内容输出到标准输出（stdout）。
// 默认情况下它是启用的。
func X是否同时输出到终端(开启 ...bool) *Logger {
	return defaultLogger.X是否同时输出到终端(开启...)
}

// Header 是一个链式函数，
// 用于启用/禁用当前日志内容输出的头部信息。
// 默认情况下它是启用的。
func X是否输出头信息(开启 ...bool) *Logger {
	return defaultLogger.X是否输出头信息(开启...)
}

// Line 是一个链式函数，
// 它用于启用/禁用在输出时附带调用者文件及其行号。
// 参数 `long` 指定是否打印完整的绝对文件路径，例如：/a/b/c/d.go:23。
func X是否输出源文件路径与行号(开启 ...bool) *Logger {
	return defaultLogger.X是否输出源文件路径与行号(开启...)
}

// Async 是一个链式函数，
// 用于启用/禁用异步日志输出功能。
func X是否异步输出(开启 ...bool) *Logger {
	return defaultLogger.X是否异步输出(开启...)
}
