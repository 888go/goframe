// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 日志类

import (
	"io"

	gfile "github.com/888go/goframe/os/gfile"
)

// X重定向输出是一个链式函数，
// 它将当前日志内容输出重定向到指定的`writer`。
// md5:ebcc479882059a2e
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

// X文件路径 是一个链接函数，
// 用于设置当前日志内容输出的目录路径为 `path`。
//
// 注意，参数 `path` 是一个目录路径，而不是文件路径。
// md5:32049f7ff7fb26c7
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

// X文件分类是一个链式函数，
// 它将当前日志输出的内容类别设置为`category`。
// 参数`category`可以是分层的，例如：module/user。
// md5:a2af01bf0ef7b61a
func (l *Logger) X文件分类(类名称 string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	if logger.config.Path != "" {
		if err := logger.X设置文件路径(gfile.X路径生成(logger.config.Path, 类名称)); err != nil {
			panic(err)
		}
	}
	return logger
}

// X文件名格式是一个链式函数，
// 用于设置当前日志内容输出的文件名`pattern`。
// md5:2daa0ddd7ccf2b81
func (l *Logger) X文件名格式(文件名格式 string) *Logger {
	if 文件名格式 == "" {
		return l
	}
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	logger.X设置文件名格式(文件名格式)
	return logger
}

// X级别是一个链式函数，它为当前日志输出内容设置日志级别。
// md5:47d5e7fde254fd72
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

// X文本级别 是一个链式函数，
// 它使用级别字符串设置当前日志输出内容的记录级别。
// md5:1edd3ebd98ec4de9
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

// X堆栈偏移量 是一个链式函数，
// 用于设置当前日志内容输出的堆栈跳过级别。
// 当启用行号打印时，它也会影响调用者文件路径的检查。
// md5:f009e835989b7d56
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

// X堆栈选项是一个链式函数，
// 用于设置当前日志内容输出的堆栈选项。
// md5:aa980bb8d1f29075
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

// X堆栈过滤是一个链式函数，它为当前日志内容输出设置堆栈过滤器。
// md5:07c5327660880bce
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

// X是否同时输出到终端 是一个链式函数，
// 它允许/禁用当前日志内容输出到标准输出。默认启用。
// md5:47b9f2393a7d5d1f
func (l *Logger) X是否同时输出到终端(开启 ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
		// 如果未传入`enabled`，则启用标准输出打印。 md5:dd98fbe92ec5263e
	if len(开启) > 0 && !开启[0] {
		logger.config.StdoutPrint = false
	} else {
		logger.config.StdoutPrint = true
	}
	return logger
}

// X是否输出头信息 是一个链接函数，
// 用于启用或禁用当前日志内容输出的头部信息。
// 默认情况下，它是启用的。
// md5:0d1df22042835f38
func (l *Logger) X是否输出头信息(开启 ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
		// 如果没有传递`enabled`，则启用了header。 md5:a42dc44d8fc72606
	if len(开启) > 0 && !开启[0] {
		logger.X设置是否输出头信息(false)
	} else {
		logger.X设置是否输出头信息(true)
	}
	return logger
}

// X是否输出源文件路径与行号 is a chaining function,
// which enables/disables printing its caller file path along with its line number.
// The parameter `long` specified whether print the long absolute file path, eg: /a/b/c/d.go:23,
// or else short one: d.go:23.
func (l *Logger) X是否输出源文件路径与行号(开启 ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
	if len(开启) > 0 && 开启[0] {
		logger.config.Flags |= F_FILE_LONG
	} else {
		logger.config.Flags |= F_FILE_SHORT
	}
	return logger
}

// X是否异步输出 是一个链式函数，
// 用于启用或禁用异步日志输出功能。
// md5:8b08233b7c37c09f
func (l *Logger) X是否异步输出(开启 ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.X取副本()
	} else {
		logger = l
	}
		// 如果未传递`enabled`参数，则异步功能启用。 md5:93404b12c04ed6a0
	if len(开启) > 0 && !开启[0] {
		logger.X设置异步输出(false)
	} else {
		logger.X设置异步输出(true)
	}
	return logger
}
