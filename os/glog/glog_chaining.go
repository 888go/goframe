// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 日志类

import (
	"io"
)

// Expose别名 返回 glog 包的默认日志记录器。 md5:20ba79507675e0e8
func Expose别名() *Logger {
	return defaultLogger
}

// X重定向输出是一个链式函数，
// 它将当前日志内容输出重定向到指定的`writer`。
// md5:198a22b4aa1923ef
func X重定向输出(writer io.Writer) *Logger {
	return defaultLogger.X重定向输出(writer)
}

// X文件路径是一个链式函数，它为当前日志内容输出设置目录路径为`path`。
// md5:9f023b954ca40c7e
func X文件路径(文件路径 string) *Logger {
	return defaultLogger.X文件路径(文件路径)
}

// X文件分类 是一个链式函数，
// 用于将当前日志内容输出的类别设置为 `category`。
// md5:9233c1ec32751e89
func X文件分类(类名称 string) *Logger {
	return defaultLogger.X文件分类(类名称)
}

// X文件名格式是一个链式函数，
// 用于设置当前日志内容输出的文件名`pattern`。
// md5:2daa0ddd7ccf2b81
func X文件名格式(文件名格式 string) *Logger {
	return defaultLogger.X文件名格式(文件名格式)
}

// X级别是一个链式函数，它为当前日志输出内容设置日志级别。
// md5:47d5e7fde254fd72
func X级别(级别 int) *Logger {
	return defaultLogger.X级别(级别)
}

// X文本级别 是一个链式函数，
// 它使用级别字符串设置当前日志输出内容的记录级别。
// md5:1edd3ebd98ec4de9
func X文本级别(文本级别 string) *Logger {
	return defaultLogger.X文本级别(文本级别)
}

// X堆栈偏移量 是一个链式函数，
// 用于设置当前日志内容输出的堆栈跳过级别。
// 当启用行号打印时，它也会影响调用者文件路径的检查。
// md5:f009e835989b7d56
func X堆栈偏移量(偏移量 int) *Logger {
	return defaultLogger.X堆栈偏移量(偏移量)
}

// X堆栈选项是一个链式函数，
// 用于设置当前日志内容输出的堆栈选项。
// md5:aa980bb8d1f29075
func X堆栈选项(开启 bool, 偏移量 ...int) *Logger {
	return defaultLogger.X堆栈选项(开启, 偏移量...)
}

// X堆栈过滤是一个链式函数，它为当前日志内容输出设置堆栈过滤器。
// md5:07c5327660880bce
func X堆栈过滤(过滤器 string) *Logger {
	return defaultLogger.X堆栈过滤(过滤器)
}

// X是否同时输出到终端 是一个链式函数，
// 它允许/禁用当前日志内容输出到标准输出。默认启用。
// md5:47b9f2393a7d5d1f
func X是否同时输出到终端(开启 ...bool) *Logger {
	return defaultLogger.X是否同时输出到终端(开启...)
}

// X是否输出头信息 是一个链接函数，
// 用于启用或禁用当前日志内容输出的头部信息。
// 默认情况下，它是启用的。
// md5:0d1df22042835f38
func X是否输出头信息(开启 ...bool) *Logger {
	return defaultLogger.X是否输出头信息(开启...)
}

// X是否输出源文件路径与行号 is a chaining function,
// which enables/disables printing its caller file along with its line number.
// The parameter `long` specified whether print the long absolute file path, eg: /a/b/c/d.go:23.
func X是否输出源文件路径与行号(开启 ...bool) *Logger {
	return defaultLogger.X是否输出源文件路径与行号(开启...)
}

// X是否异步输出 是一个链式函数，
// 用于启用或禁用异步日志输出功能。
// md5:8b08233b7c37c09f
func X是否异步输出(开启 ...bool) *Logger {
	return defaultLogger.X是否异步输出(开启...)
}
