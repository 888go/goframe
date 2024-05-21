// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package glog

import (
	"io"
)

// Expose returns the default logger of package glog.

// ff:Expose别名
func Expose() *Logger {
	return defaultLogger
}

// To is a chaining function,
// which redirects current logging content output to the sepecified `writer`.

// ff:重定向输出
// writer:
func To(writer io.Writer) *Logger {
	return defaultLogger.To(writer)
}

// Path is a chaining function,
// which sets the directory path to `path` for current logging content output.

// ff:文件路径
// path:文件路径
func Path(path string) *Logger {
	return defaultLogger.Path(path)
}

// Cat is a chaining function,
// which sets the category to `category` for current logging content output.

// ff:文件分类
// category:类名称
func Cat(category string) *Logger {
	return defaultLogger.Cat(category)
}

// File is a chaining function,
// which sets file name `pattern` for the current logging content output.

// ff:文件名格式
// pattern:文件名格式
func File(pattern string) *Logger {
	return defaultLogger.File(pattern)
}

// Level is a chaining function,
// which sets logging level for the current logging content output.

// ff:级别
// level:级别
func Level(level int) *Logger {
	return defaultLogger.Level(level)
}

// LevelStr is a chaining function,
// which sets logging level for the current logging content output using level string.

// ff:文本级别
// levelStr:文本级别
func LevelStr(levelStr string) *Logger {
	return defaultLogger.LevelStr(levelStr)
}

// Skip is a chaining function,
// which sets stack skip for the current logging content output.
// It also affects the caller file path checks when line number printing enabled.

// ff:堆栈偏移量
// skip:偏移量
func Skip(skip int) *Logger {
	return defaultLogger.Skip(skip)
}

// Stack is a chaining function,
// which sets stack options for the current logging content output .

// ff:堆栈选项
// skip:偏移量
// enabled:开启
func Stack(enabled bool, skip ...int) *Logger {
	return defaultLogger.Stack(enabled, skip...)
}

// StackWithFilter is a chaining function,
// which sets stack filter for the current logging content output .

// ff:堆栈过滤
// filter:过滤器
func StackWithFilter(filter string) *Logger {
	return defaultLogger.StackWithFilter(filter)
}

// Stdout is a chaining function,
// which enables/disables stdout for the current logging content output.
// It's enabled in default.

// ff:是否同时输出到终端
// enabled:开启
func Stdout(enabled ...bool) *Logger {
	return defaultLogger.Stdout(enabled...)
}

// Header is a chaining function,
// which enables/disables log header for the current logging content output.
// It's enabled in default.

// ff:是否输出头信息
// enabled:开启
func Header(enabled ...bool) *Logger {
	return defaultLogger.Header(enabled...)
}

// Line is a chaining function,
// which enables/disables printing its caller file along with its line number.
// The parameter `long` specified whether print the long absolute file path, eg: /a/b/c/d.go:23.

// ff:是否输出源文件路径与行号
// long:开启
func Line(long ...bool) *Logger {
	return defaultLogger.Line(long...)
}

// Async is a chaining function,
// which enables/disables async logging output feature.

// ff:是否异步输出
// enabled:开启
func Async(enabled ...bool) *Logger {
	return defaultLogger.Async(enabled...)
}
