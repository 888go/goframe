// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类

import (
	"io"
	
	"github.com/888go/goframe/os/gfile"
)

// To 是一个链式函数，
// 它将当前日志内容输出重定向到指定的 `writer`。
func (l *Logger) X重定向输出(writer io.Writer) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	logger.X设置Writer(writer)
	return logger
}

// Path 是一个链式函数，
// 用于将当前日志内容输出的目标目录路径设置为 `path`。
//
// 注意，参数 `path` 应为一个目录路径，而非文件路径。
func (l *Logger) X文件路径(文件路径 string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	if 文件路径 != "" {
		if err := logger.X设置文件路径(文件路径); err != nil {
			panic(err)
		}
	}
	return logger
}

// Cat 是一个链式函数，
// 它用于为当前日志内容输出设置类别为 `category`。
// 参数 `category` 可以是层级式的，例如：module/user。
func (l *Logger) X文件分类(类名称 string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	if logger.config.X文件路径 != "" {
		if err := logger.X设置文件路径(文件类.X路径生成(logger.config.X文件路径, 类名称)); err != nil {
			panic(err)
		}
	}
	return logger
}

// File 是一个链式函数，
// 用于设置当前日志内容输出的文件名模式。
func (l *Logger) X文件名格式(文件名格式 string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	logger.X设置文件名格式(文件名格式)
	return logger
}

// Level 是一个链式函数，
// 用于设置当前日志内容输出的记录级别。
func (l *Logger) X级别(级别 int) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	logger.X设置级别(级别)
	return logger
}

// LevelStr 是一个链式函数，
// 通过级别字符串设置当前日志内容输出的记录级别。
func (l *Logger) X文本级别(文本级别 string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	if err := logger.X设置文本级别(文本级别); err != nil {
		panic(err)
	}
	return logger
}

// Skip 是一个链式函数，
// 用于设置当前日志内容输出时的堆栈跳过级别。
// 当启用行号打印时，它也会影响调用文件路径的检查。
func (l *Logger) X堆栈偏移量(偏移量 int) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	logger.X设置堆栈偏移量(偏移量)
	return logger
}

// Stack 是一个链式函数，
// 用于为当前日志内容输出设置堆栈选项。
func (l *Logger) X堆栈选项(开启 bool, 偏移量 ...int) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	logger.X设置堆栈跟踪(开启)
	if len(偏移量) > 0 {
		logger.X设置堆栈偏移量(偏移量[0])
	}
	return logger
}

// StackWithFilter 是一个链式函数，
// 用于为当前日志内容输出设置堆栈过滤器。
func (l *Logger) X堆栈过滤(过滤器 string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	logger.X设置堆栈跟踪(true)
	logger.X设置堆栈过滤(过滤器)
	return logger
}

// Stdout 是一个链式函数，
// 用于启用/禁用当前日志内容输出到标准输出（stdout）。
// 默认情况下它是启用的。
func (l *Logger) X是否同时输出到终端(开启 ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	// 如果未传递 `enabled`，则启用标准输出打印。
	if len(开启) > 0 && !开启[0] {
		logger.config.X是否同时输出到终端 = false
	} else {
		logger.config.X是否同时输出到终端 = true
	}
	return logger
}

// Header 是一个链式函数，
// 用于启用/禁用当前日志内容输出的头部信息。
// 默认情况下它是启用的。
func (l *Logger) X是否输出头信息(开启 ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	// 如果未传递`enabled`参数，则启用header。
	if len(开启) > 0 && !开启[0] {
		logger.X设置是否输出头信息(false)
	} else {
		logger.X设置是否输出头信息(true)
	}
	return logger
}

// Line 是一个链式函数，
// 它用于启用/禁用在输出时附带调用者所在的文件路径及行号。
// 参数 `long` 指定了是否打印完整的绝对文件路径，例如：/a/b/c/d.go:23，
// 否则只打印简短形式：d.go:23。
func (l *Logger) X是否输出源文件路径与行号(开启 ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	if len(开启) > 0 && 开启[0] {
		logger.config.X日志标识 |= F_FILE_LONG
	} else {
		logger.config.X日志标识 |= F_FILE_SHORT
	}
	return logger
}

// Async 是一个链式函数，
// 用于启用/禁用异步日志输出功能。
func (l *Logger) X是否异步输出(开启 ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	// 如果未传递`enabled`参数，则启用异步功能。
	if len(开启) > 0 && !开启[0] {
		logger.X设置异步输出(false)
	} else {
		logger.X设置异步输出(true)
	}
	return logger
}
