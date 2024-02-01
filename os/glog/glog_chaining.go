// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package glog
import (
	"io"
	)
// Expose 返回 glog 包的默认日志器。
// DefaultLogger()  方法的别名
func Expose() *Logger {
	return defaultLogger
}

// To 是一个链式函数，
// 它将当前日志内容输出重定向到指定的 `writer`。
func To(writer io.Writer) *Logger {
	return defaultLogger.To(writer)
}

// Path 是一个链式函数，
// 它用于将当前日志内容输出的目标目录路径设置为 `path`。
func Path(path string) *Logger {
	return defaultLogger.Path(path)
}

// Cat 是一个链式函数，
// 它用于将当前日志内容输出的类别设置为 `category`。
func Cat(category string) *Logger {
	return defaultLogger.Cat(category)
}

// File 是一个链式函数，
// 用于设置当前日志内容输出的文件名模式。
func File(pattern string) *Logger {
	return defaultLogger.File(pattern)
}

// Level 是一个链式函数，
// 用于设置当前日志内容输出的记录级别。
func Level(level int) *Logger {
	return defaultLogger.Level(level)
}

// LevelStr 是一个链式函数，
// 通过级别字符串设置当前日志内容输出的记录级别。
func LevelStr(levelStr string) *Logger {
	return defaultLogger.LevelStr(levelStr)
}

// Skip 是一个链式函数，
// 用于设置当前日志内容输出时的堆栈跳过级别。
// 当启用行号打印时，它也会影响调用文件路径的检查。
func Skip(skip int) *Logger {
	return defaultLogger.Skip(skip)
}

// Stack 是一个链式函数，
// 用于为当前日志内容输出设置堆栈选项。
func Stack(enabled bool, skip ...int) *Logger {
	return defaultLogger.Stack(enabled, skip...)
}

// StackWithFilter 是一个链式函数，
// 用于为当前日志内容输出设置堆栈过滤器。
func StackWithFilter(filter string) *Logger {
	return defaultLogger.StackWithFilter(filter)
}

// Stdout 是一个链式函数，
// 用于启用/禁用当前日志内容输出到标准输出（stdout）。
// 默认情况下它是启用的。
func Stdout(enabled ...bool) *Logger {
	return defaultLogger.Stdout(enabled...)
}

// Header 是一个链式函数，
// 用于启用/禁用当前日志内容输出的头部信息。
// 默认情况下它是启用的。
func Header(enabled ...bool) *Logger {
	return defaultLogger.Header(enabled...)
}

// Line 是一个链式函数，
// 它用于启用/禁用在输出时附带调用者文件及其行号。
// 参数 `long` 指定是否打印完整的绝对文件路径，例如：/a/b/c/d.go:23。
func Line(long ...bool) *Logger {
	return defaultLogger.Line(long...)
}

// Async 是一个链式函数，
// 用于启用/禁用异步日志输出功能。
func Async(enabled ...bool) *Logger {
	return defaultLogger.Async(enabled...)
}
