// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package glog

import (
	"io"

	"github.com/gogf/gf/v2/os/gfile"
)

// To是一个链式函数，
// 它将当前日志内容输出重定向到指定的`writer`。
// md5:ebcc479882059a2e
// ff:重定向输出
// l:
// writer:
func (l *Logger) To(writer io.Writer) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetWriter(writer)
	return logger
}

// Path 是一个链接函数，
// 用于设置当前日志内容输出的目录路径为 `path`。
//
// 注意，参数 `path` 是一个目录路径，而不是文件路径。
// md5:32049f7ff7fb26c7
// ff:文件路径
// l:
// path:文件路径
func (l *Logger) Path(path string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	if path != "" {
		if err := logger.SetPath(path); err != nil {
			panic(err)
		}
	}
	return logger
}

// Cat是一个链式函数，
// 它将当前日志输出的内容类别设置为`category`。
// 参数`category`可以是分层的，例如：module/user。
// md5:a2af01bf0ef7b61a
// ff:文件分类
// l:
// category:类名称
func (l *Logger) Cat(category string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	if logger.config.Path != "" {
		if err := logger.SetPath(gfile.Join(logger.config.Path, category)); err != nil {
			panic(err)
		}
	}
	return logger
}

// File是一个链式函数，
// 用于设置当前日志内容输出的文件名`pattern`。
// md5:2daa0ddd7ccf2b81
// ff:文件名格式
// l:
// file:文件名格式
func (l *Logger) File(file string) *Logger {
	if file == "" {
		return l
	}
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetFile(file)
	return logger
}

// Level是一个链式函数，它为当前日志输出内容设置日志级别。
// md5:47d5e7fde254fd72
// ff:级别
// l:
// level:级别
func (l *Logger) Level(level int) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetLevel(level)
	return logger
}

// LevelStr 是一个链式函数，
// 它使用级别字符串设置当前日志输出内容的记录级别。
// md5:1edd3ebd98ec4de9
// ff:文本级别
// l:
// levelStr:文本级别
func (l *Logger) LevelStr(levelStr string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	if err := logger.SetLevelStr(levelStr); err != nil {
		panic(err)
	}
	return logger
}

// Skip 是一个链式函数，
// 用于设置当前日志内容输出的堆栈跳过级别。
// 当启用行号打印时，它也会影响调用者文件路径的检查。
// md5:f009e835989b7d56
// ff:堆栈偏移量
// l:
// skip:偏移量
func (l *Logger) Skip(skip int) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetStackSkip(skip)
	return logger
}

// Stack是一个链式函数，
// 用于设置当前日志内容输出的堆栈选项。
// md5:aa980bb8d1f29075
// ff:堆栈选项
// l:
// enabled:开启
// skip:偏移量
func (l *Logger) Stack(enabled bool, skip ...int) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetStack(enabled)
	if len(skip) > 0 {
		logger.SetStackSkip(skip[0])
	}
	return logger
}

// StackWithFilter是一个链式函数，它为当前日志内容输出设置堆栈过滤器。
// md5:07c5327660880bce
// ff:堆栈过滤
// l:
// filter:过滤器
func (l *Logger) StackWithFilter(filter string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetStack(true)
	logger.SetStackFilter(filter)
	return logger
}

// Stdout 是一个链式函数，
// 它允许/禁用当前日志内容输出到标准输出。默认启用。
// md5:47b9f2393a7d5d1f
// ff:是否同时输出到终端
// l:
// enabled:开启
func (l *Logger) Stdout(enabled ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	// 如果未传入`enabled`，则启用标准输出打印。 md5:dd98fbe92ec5263e
	if len(enabled) > 0 && !enabled[0] {
		logger.config.StdoutPrint = false
	} else {
		logger.config.StdoutPrint = true
	}
	return logger
}

// Header 是一个链接函数，
// 用于启用或禁用当前日志内容输出的头部信息。
// 默认情况下，它是启用的。
// md5:0d1df22042835f38
// ff:是否输出头信息
// l:
// enabled:开启
func (l *Logger) Header(enabled ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	// 如果没有传递`enabled`，则启用了header。 md5:a42dc44d8fc72606
	if len(enabled) > 0 && !enabled[0] {
		logger.SetHeaderPrint(false)
	} else {
		logger.SetHeaderPrint(true)
	}
	return logger
}

// Line is a chaining function,
// which enables/disables printing its caller file path along with its line number.
// The parameter `long` specified whether print the long absolute file path, eg: /a/b/c/d.go:23,
// or else short one: d.go:23.
// ff:是否输出源文件路径与行号
// l:
// long:开启
func (l *Logger) Line(long ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	if len(long) > 0 && long[0] {
		logger.config.Flags |= F_FILE_LONG
	} else {
		logger.config.Flags |= F_FILE_SHORT
	}
	return logger
}

// Async 是一个链式函数，
// 用于启用或禁用异步日志输出功能。
// md5:8b08233b7c37c09f
// ff:是否异步输出
// l:
// enabled:开启
func (l *Logger) Async(enabled ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	// 如果未传递`enabled`参数，则异步功能启用。 md5:93404b12c04ed6a0
	if len(enabled) > 0 && !enabled[0] {
		logger.SetAsync(false)
	} else {
		logger.SetAsync(true)
	}
	return logger
}
