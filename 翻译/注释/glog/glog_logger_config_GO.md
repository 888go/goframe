
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Config is the configuration object for logger.
<原文结束>

# <翻译开始>
// Config 是 logger 的配置对象。
# <翻译结束>


<原文开始>
// Logger handlers which implement feature similar as middleware.
<原文结束>

# <翻译开始>
// 日志处理器实现类似中间件的功能。
# <翻译结束>







<原文开始>
// Extra flags for logging output features.
<原文结束>

# <翻译开始>
// 附加日志输出功能的标志。
# <翻译结束>












<原文开始>
// Format pattern for logging file.
<原文结束>

# <翻译开始>
// 日志文件的格式化模式。
# <翻译结束>


<原文开始>
// Prefix string for every logging content.
<原文结束>

# <翻译开始>
// Prefix 字符串，用于作为每条日志内容的前缀。
# <翻译结束>







<原文开始>
// Stack status(1: enabled - default; 0: disabled)
<原文结束>

# <翻译开始>
// 栈状态(1: 启用 - 默认值; 0: 禁用)
# <翻译结束>







<原文开始>
// Context keys for logging, which is used for value retrieving from context.
<原文结束>

# <翻译开始>
// Context keys 用于日志记录，它们被用于从context中检索值。
# <翻译结束>


<原文开始>
// Print header or not(true in default).
<原文结束>

# <翻译开始>
// 是否打印头部信息（默认为true）
# <翻译结束>


<原文开始>
// Output to stdout or not(true in default).
<原文结束>

# <翻译开始>
// 是否输出到标准输出（默认为true）
# <翻译结束>


<原文开始>
// Print level format string or not(true in default).
<原文结束>

# <翻译开始>
// 是否打印级别格式化字符串（默认为true）
# <翻译结束>


<原文开始>
// Logging level to its prefix string mapping.
<原文结束>

# <翻译开始>
// 日志级别到其前缀字符串的映射。
# <翻译结束>


<原文开始>
// Rotate the logging file if its size > 0 in bytes.
<原文结束>

# <翻译开始>
// 如果日志文件大小 > 0字节，则旋转日志文件。
# <翻译结束>


<原文开始>
// Rotate the logging file if its mtime exceeds this duration.
<原文结束>

# <翻译开始>
// 如果日志文件的修改时间超过这个持续时间，则旋转日志文件。
# <翻译结束>


<原文开始>
// Max backup for rotated files, default is 0, means no backups.
<原文结束>

# <翻译开始>
// Max 备份文件数量，默认为0，表示不进行备份。
# <翻译结束>


<原文开始>
// Max expires for rotated files, which is 0 in default, means no expiration.
<原文结束>

# <翻译开始>
// Max 为轮转文件设置的过期时间，默认为0，表示永不过期。
# <翻译结束>


<原文开始>
// Logging level prefix with color to writer or not (false in default).
<原文结束>

# <翻译开始>
// 是否（默认为false）在向writer输出时，以颜色前缀形式记录日志级别
# <翻译结束>


<原文开始>
// Whether the rotation feature initialized.
<原文结束>

# <翻译开始>
// 是否已初始化旋转功能
# <翻译结束>


<原文开始>
// DefaultConfig returns the default configuration for logger.
<原文结束>

# <翻译开始>
// DefaultConfig 返回日志器的默认配置。
# <翻译结束>


<原文开始>
// GetConfig returns the configuration of current Logger.
<原文结束>

# <翻译开始>
// GetConfig 返回当前 Logger 的配置。
# <翻译结束>


<原文开始>
// SetConfig set configurations for the logger.
<原文结束>

# <翻译开始>
// SetConfig 为日志器设置配置。
# <翻译结束>







<原文开始>
// SetConfigWithMap set configurations with map for the logger.
<原文结束>

# <翻译开始>
// SetConfigWithMap 通过map设置日志器的配置。
# <翻译结束>


<原文开始>
	// The m now is a shallow copy of m.
	// A little tricky, isn't it?
<原文结束>

# <翻译开始>
// 现在的m是m的一个浅拷贝。
// 这有点巧妙，不是吗？
# <翻译结束>


<原文开始>
// Change string configuration to int value for level.
<原文结束>

# <翻译开始>
// 将字符串配置转换为级别对应的整数值。
# <翻译结束>


<原文开始>
// Change string configuration to int value for file rotation size.
<原文结束>

# <翻译开始>
// 将字符串配置转换为文件旋转大小的整数值。
# <翻译结束>


<原文开始>
// SetDebug enables/disables the debug level for logger.
// The debug level is enabled in default.
<原文结束>

# <翻译开始>
// SetDebug 用于开启或关闭日志器的调试级别。
// 默认情况下，调试级别是启用的。
# <翻译结束>


<原文开始>
// SetAsync enables/disables async logging output feature.
<原文结束>

# <翻译开始>
// SetAsync 启用/禁用异步日志输出功能。
# <翻译结束>


<原文开始>
// SetFlags sets extra flags for logging output features.
<原文结束>

# <翻译开始>
// SetFlags 设置日志输出功能的额外标志。
# <翻译结束>


<原文开始>
// GetFlags returns the flags of logger.
<原文结束>

# <翻译开始>
// GetFlags 返回日志器的标志。
# <翻译结束>


<原文开始>
// SetStack enables/disables the stack feature in failure logging outputs.
<原文结束>

# <翻译开始>
// SetStack 启用/禁用失败日志输出中的堆栈跟踪功能。
# <翻译结束>


<原文开始>
// SetStackSkip sets the stack offset from the end point.
<原文结束>

# <翻译开始>
// SetStackSkip 设置从终点开始的堆栈偏移量。
# <翻译结束>


<原文开始>
// SetStackFilter sets the stack filter from the end point.
<原文结束>

# <翻译开始>
// SetStackFilter 从终点设置堆栈过滤器。
# <翻译结束>


<原文开始>
// SetCtxKeys sets the context keys for logger. The keys is used for retrieving values
// from context and printing them to logging content.
//
// Note that multiple calls of this function will overwrite the previous set context keys.
<原文结束>

# <翻译开始>
// SetCtxKeys 设置日志器的上下文键。这些键用于从上下文中检索值并将其打印到日志内容中。
//
// 注意，多次调用此函数将覆盖之前设置的上下文键。
# <翻译结束>


<原文开始>
// AppendCtxKeys appends extra keys to logger.
// It ignores the key if it is already appended to the logger previously.
<原文结束>

# <翻译开始>
// AppendCtxKeys 向日志器追加额外键。
// 如果该键之前已向日志器追加过，则忽略此次操作。
# <翻译结束>


<原文开始>
// GetCtxKeys retrieves and returns the context keys for logging.
<原文结束>

# <翻译开始>
// GetCtxKeys 获取并返回用于日志记录的上下文键。
# <翻译结束>


<原文开始>
// SetWriter sets the customized logging `writer` for logging.
// The `writer` object should implement the io.Writer interface.
// Developer can use customized logging `writer` to redirect logging output to another service,
// eg: kafka, mysql, mongodb, etc.
<原文结束>

# <翻译开始>
// SetWriter 设置自定义的日志 `writer` 用于日志记录。
// `writer` 对象应实现 io.Writer 接口。
// 开发者可以使用自定义的日志 `writer` 将日志输出重定向到其他服务，
// 例如：kafka、mysql、mongodb 等。
# <翻译结束>


<原文开始>
// GetWriter returns the customized writer object, which implements the io.Writer interface.
// It returns nil if no writer previously set.
<原文结束>

# <翻译开始>
// GetWriter 返回自定义的writer对象，该对象实现了io.Writer接口。
// 如果之前未设置过writer，则返回nil。
# <翻译结束>


<原文开始>
// SetPath sets the directory path for file logging.
<原文结束>

# <翻译开始>
// SetPath 设置文件日志的目录路径。
# <翻译结束>


<原文开始>
// GetPath returns the logging directory path for file logging.
// It returns empty string if no directory path set.
<原文结束>

# <翻译开始>
// GetPath 返回用于文件日志记录的日志目录路径。
// 如果未设置目录路径，则返回空字符串。
# <翻译结束>


<原文开始>
// SetFile sets the file name `pattern` for file logging.
// Datetime pattern can be used in `pattern`, eg: access-{Ymd}.log.
// The default file name pattern is: Y-m-d.log, eg: 2018-01-01.log
<原文结束>

# <翻译开始>
// SetFile 设置文件日志的文件名`pattern`。
// 在`pattern`中可以使用日期时间模式，例如：access-{Ymd}.log。
// 默认的文件名模式是：Y-m-d.log，例如：2018-01-01.log
// 这段Go语言代码注释翻译成中文为：
// 设置文件日志的文件名为 `pattern`。
// 可以在 `pattern` 中使用日期时间格式化字符串，例如：access-{Ymd}.log（表示按年月日生成不同文件）。
// 默认的文件名格式是：Y-m-d.log，例如：2018-01-01.log
# <翻译结束>


<原文开始>
// SetTimeFormat sets the time format for the logging time.
<原文结束>

# <翻译开始>
// SetTimeFormat 设置日志时间的时间格式。
# <翻译结束>


<原文开始>
// SetStdoutPrint sets whether output the logging contents to stdout, which is true in default.
<原文结束>

# <翻译开始>
// SetStdoutPrint 设置是否将日志内容输出到标准输出(stdout)，默认为true。
# <翻译结束>


<原文开始>
// SetHeaderPrint sets whether output header of the logging contents, which is true in default.
<原文结束>

# <翻译开始>
// SetHeaderPrint 设置是否输出日志内容的头部，默认为true。
# <翻译结束>


<原文开始>
// SetLevelPrint sets whether output level string of the logging contents, which is true in default.
<原文结束>

# <翻译开始>
// SetLevelPrint 设置是否输出日志内容的级别字符串，默认为true。
# <翻译结束>


<原文开始>
// SetPrefix sets prefix string for every logging content.
// Prefix is part of header, which means if header output is shut, no prefix will be output.
<原文结束>

# <翻译开始>
// SetPrefix 设置每个日志内容的前缀字符串。
// 前缀是头部的一部分，这意味着如果关闭了头部输出，则不会输出任何前缀。
# <翻译结束>


<原文开始>
// SetHandlers sets the logging handlers for current logger.
<原文结束>

# <翻译开始>
// SetHandlers 设置当前日志器的处理程序。
# <翻译结束>


<原文开始>
// SetWriterColorEnable enables file/writer logging with color.
<原文结束>

# <翻译开始>
// SetWriterColorEnable 开启文件/写入器日志的彩色输出功能。
# <翻译结束>


<原文开始>
// SetStdoutColorDisabled disables stdout logging with color.
<原文结束>

# <翻译开始>
// SetStdoutColorDisabled 禁用 stdout 日志颜色输出。
# <翻译结束>


<原文开始>
// Customized io.Writer.
<原文结束>

# <翻译开始>
// 自定义 io.Writer。
# <翻译结束>


<原文开始>
// Logging time format
<原文结束>

# <翻译开始>
// 日志时间格式
# <翻译结束>


<原文开始>
// Logging directory path.
<原文结束>

# <翻译开始>
// 日志目录路径。
# <翻译结束>


<原文开始>
// Skipping count for stack.
<原文结束>

# <翻译开始>
// 跳过堆栈的计数。
# <翻译结束>


<原文开始>
// Stack string filter.
<原文结束>

# <翻译开始>
// 字符串过滤栈
# <翻译结束>


<原文开始>
// Necessary validation.
<原文结束>

# <翻译开始>
// 必要的验证
# <翻译结束>

