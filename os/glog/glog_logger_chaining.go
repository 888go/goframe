// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package glog

import (
	"io"
	
	"github.com/888go/goframe/os/gfile"
)

// To 是一个链式函数，
// 它将当前日志内容输出重定向到指定的 `writer`。
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

// Path 是一个链式函数，
// 用于将当前日志内容输出的目标目录路径设置为 `path`。
//
// 注意，参数 `path` 应为一个目录路径，而非文件路径。
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

// Cat 是一个链式函数，
// 它用于为当前日志内容输出设置类别为 `category`。
// 参数 `category` 可以是层级式的，例如：module/user。
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

// File 是一个链式函数，
// 用于设置当前日志内容输出的文件名模式。
func (l *Logger) File(file string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetFile(file)
	return logger
}

// Level 是一个链式函数，
// 用于设置当前日志内容输出的记录级别。
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
// 通过级别字符串设置当前日志内容输出的记录级别。
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
// 用于设置当前日志内容输出时的堆栈跳过级别。
// 当启用行号打印时，它也会影响调用文件路径的检查。
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

// Stack 是一个链式函数，
// 用于为当前日志内容输出设置堆栈选项。
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

// StackWithFilter 是一个链式函数，
// 用于为当前日志内容输出设置堆栈过滤器。
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
// 用于启用/禁用当前日志内容输出到标准输出（stdout）。
// 默认情况下它是启用的。
func (l *Logger) Stdout(enabled ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	// 如果未传递 `enabled`，则启用标准输出打印。
	if len(enabled) > 0 && !enabled[0] {
		logger.config.StdoutPrint = false
	} else {
		logger.config.StdoutPrint = true
	}
	return logger
}

// Header 是一个链式函数，
// 用于启用/禁用当前日志内容输出的头部信息。
// 默认情况下它是启用的。
func (l *Logger) Header(enabled ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	// 如果未传递`enabled`参数，则启用header。
	if len(enabled) > 0 && !enabled[0] {
		logger.SetHeaderPrint(false)
	} else {
		logger.SetHeaderPrint(true)
	}
	return logger
}

// Line 是一个链式函数，
// 它用于启用/禁用在输出时附带调用者所在的文件路径及行号。
// 参数 `long` 指定了是否打印完整的绝对文件路径，例如：/a/b/c/d.go:23，
// 否则只打印简短形式：d.go:23。
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
// 用于启用/禁用异步日志输出功能。
func (l *Logger) Async(enabled ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	// 如果未传递`enabled`参数，则启用异步功能。
	if len(enabled) > 0 && !enabled[0] {
		logger.SetAsync(false)
	} else {
		logger.SetAsync(true)
	}
	return logger
}
