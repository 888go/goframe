
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Config is the configuration object for logger.
<原文结束>

# <翻译开始>
// Config 是记录器的配置对象。 md5:df2a8ab047bea305
# <翻译结束>


<原文开始>
// Logger handlers which implement feature similar as middleware.
<原文结束>

# <翻译开始>
// Logger handlers 实现了类似于中间件的功能。 md5:dba4d3d0c7f592b9
# <翻译结束>


<原文开始>
// Extra flags for logging output features.
<原文结束>

# <翻译开始>
// 用于日志输出功能的额外标志。 md5:6b323bf0cac304e0
# <翻译结束>


<原文开始>
// Logging directory path.
<原文结束>

# <翻译开始>
// 日志目录路径。 md5:a9b74f93138f8158
# <翻译结束>


<原文开始>
// Format pattern for logging file.
<原文结束>

# <翻译开始>
// 日志文件的格式化模式。 md5:88ed9324a5afc4c8
# <翻译结束>


<原文开始>
// Prefix string for every logging content.
<原文结束>

# <翻译开始>
// 每个日志内容的前缀字符串。 md5:9318d5ac0b1e3e0e
# <翻译结束>


<原文开始>
// Skipping count for stack.
<原文结束>

# <翻译开始>
// 跳过栈的计数。 md5:dd0842336cee717c
# <翻译结束>


<原文开始>
// Stack status(1: enabled - default; 0: disabled)
<原文结束>

# <翻译开始>
// 栈的状态（1：启用 - 默认；0：禁用）. md5:3a07964ef563d1f6
# <翻译结束>


<原文开始>
// Context keys for logging, which is used for value retrieving from context.
<原文结束>

# <翻译开始>
// 用于日志记录的上下文键，用于从上下文中检索值。 md5:d4a4f8b7c1027e23
# <翻译结束>


<原文开始>
// Print header or not(true in default).
<原文结束>

# <翻译开始>
// 是否打印头部信息，默认为true。 md5:8b587e739278ffe3
# <翻译结束>


<原文开始>
// Output to stdout or not(true in default).
<原文结束>

# <翻译开始>
// 是否将输出写入stdout（默认为true）。 md5:4f790cec19c3aa5a
# <翻译结束>


<原文开始>
// Print level format string or not(true in default).
<原文结束>

# <翻译开始>
// 是否打印级别格式字符串（默认为 true）。 md5:2261e6b6d1cccb2d
# <翻译结束>


<原文开始>
// Logging level to its prefix string mapping.
<原文结束>

# <翻译开始>
// 日志级别与其前缀字符串的映射。 md5:c4a5a8857bc37946
# <翻译结束>


<原文开始>
// Rotate the logging file if its size > 0 in bytes.
<原文结束>

# <翻译开始>
// 如果日志文件的大小大于0字节，则进行日志文件轮换。 md5:9fb4614dcea49823
# <翻译结束>


<原文开始>
// Rotate the logging file if its mtime exceeds this duration.
<原文结束>

# <翻译开始>
// 如果日志文件的修改时间超过这个持续时间，就旋转日志文件。 md5:0832ad6a5113efe9
# <翻译结束>


<原文开始>
// Max backup for rotated files, default is 0, means no backups.
<原文结束>

# <翻译开始>
// 旋转文件的最大备份数量，默认为0，表示不备份。 md5:67b4908c6a850b93
# <翻译结束>


<原文开始>
// Max expires for rotated files, which is 0 in default, means no expiration.
<原文结束>

# <翻译开始>
// Max 为旋转文件的过期时间，默认为0，表示永不过期。 md5:eedec2e3ee56fc5d
# <翻译结束>


<原文开始>
// Logging level prefix with color to writer or not (false in default).
<原文结束>

# <翻译开始>
// 是否向写入器输出带有颜色的日志级别前缀（默认为false）。 md5:cd5684396601fdfc
# <翻译结束>


<原文开始>
// Whether the rotation feature initialized.
<原文结束>

# <翻译开始>
// 是否启用了旋转功能。 md5:32f779f6bf9c7aee
# <翻译结束>


<原文开始>
// DefaultConfig returns the default configuration for logger.
<原文结束>

# <翻译开始>
// DefaultConfig 返回日志记录器的默认配置。 md5:307781636b8ca142
# <翻译结束>


<原文开始>
// GetConfig returns the configuration of current Logger.
<原文结束>

# <翻译开始>
// GetConfig 返回当前Logger的配置。 md5:34aac9175b86a456
# <翻译结束>


<原文开始>
// SetConfig set configurations for the logger.
<原文结束>

# <翻译开始>
// SetConfig 为logger设置配置。 md5:d219673b9a3ec8b0
# <翻译结束>


<原文开始>
// SetConfigWithMap set configurations with map for the logger.
<原文结束>

# <翻译开始>
// SetConfigWithMap 使用映射为日志器设置配置。 md5:a4d4197c666898a3
# <翻译结束>


<原文开始>
	// The m now is a shallow copy of m.
	// A little tricky, isn't it?
<原文结束>

# <翻译开始>
// 现在m是一个浅拷贝 of m。
// 有点巧妙，不是吗？
// md5:644970336da24c9d
# <翻译结束>


<原文开始>
// Change string configuration to int value for level.
<原文结束>

# <翻译开始>
// 将字符串配置更改为级别的整数值。 md5:e990c1dc64df6943
# <翻译结束>


<原文开始>
// Change string configuration to int value for file rotation size.
<原文结束>

# <翻译开始>
// 将字符串配置转换为文件轮转大小的整数值。 md5:b9efebe5c9999270
# <翻译结束>


<原文开始>
// SetDebug enables/disables the debug level for logger.
// The debug level is enabled in default.
<原文结束>

# <翻译开始>
// SetDebug 用于启用/禁用日志记录器的调试级别。默认情况下，调试级别是启用的。
// md5:72f0f67e25416b8e
# <翻译结束>


<原文开始>
// SetAsync enables/disables async logging output feature.
<原文结束>

# <翻译开始>
// SetAsync 启用/禁用异步日志输出功能。 md5:10096a3a0860346e
# <翻译结束>


<原文开始>
// SetFlags sets extra flags for logging output features.
<原文结束>

# <翻译开始>
// SetFlags 为日志输出功能设置额外的标志。 md5:40253d4ed662de77
# <翻译结束>


<原文开始>
// GetFlags returns the flags of logger.
<原文结束>

# <翻译开始>
// GetFlags 返回记录器的标志。 md5:b9a17daa74081d07
# <翻译结束>


<原文开始>
// SetStack enables/disables the stack feature in failure logging outputs.
<原文结束>

# <翻译开始>
// SetStack 启用/禁用失败日志输出中的堆栈功能。 md5:3c80a664fff650de
# <翻译结束>


<原文开始>
// SetStackSkip sets the stack offset from the end point.
<原文结束>

# <翻译开始>
// SetStackSkip 设置从终点开始的堆栈偏移量。 md5:98a83cd0e38dc56c
# <翻译结束>


<原文开始>
// SetStackFilter sets the stack filter from the end point.
<原文结束>

# <翻译开始>
// SetStackFilter 从端点设置堆栈过滤器。 md5:7eabd577c24907f2
# <翻译结束>


<原文开始>
// SetCtxKeys sets the context keys for logger. The keys is used for retrieving values
// from context and printing them to logging content.
//
// Note that multiple calls of this function will overwrite the previous set context keys.
<原文结束>

# <翻译开始>
// SetCtxKeys 为日志器设置上下文键。这些键用于从上下文中检索值并将其打印到日志内容中。
// 
// 注意，多次调用此函数会覆盖之前设置的上下文键。
// md5:f7244f6c7fa79db2
# <翻译结束>


<原文开始>
// AppendCtxKeys appends extra keys to logger.
// It ignores the key if it is already appended to the logger previously.
<原文结束>

# <翻译开始>
// AppendCtxKeys 向记录器添加额外的键。
// 如果该键已先前被添加到记录器中，则此操作会忽略该键。
// md5:f989e696d285ffc1
# <翻译结束>


<原文开始>
// GetCtxKeys retrieves and returns the context keys for logging.
<原文结束>

# <翻译开始>
// GetCtxKeys检索并返回日志记录的上下文键。 md5:1e780871ada8c59c
# <翻译结束>


<原文开始>
// SetWriter sets the customized logging `writer` for logging.
// The `writer` object should implement the io.Writer interface.
// Developer can use customized logging `writer` to redirect logging output to another service,
// eg: kafka, mysql, mongodb, etc.
<原文结束>

# <翻译开始>
// SetWriter 设置自定义的日志记录器`writer`。
// `writer`对象应实现io.Writer接口。
// 开发者可以使用自定义的`writer`将日志输出重定向到其他服务，
// 例如：kafka，mysql，mongodb等。
// md5:8f650a69c1fe2b4b
# <翻译结束>


<原文开始>
// GetWriter returns the customized writer object, which implements the io.Writer interface.
// It returns nil if no writer previously set.
<原文结束>

# <翻译开始>
// GetWriter 返回自定义的 writer 对象，该对象实现了 io.Writer 接口。
// 如果之前未设置 writer，则返回 nil。
// md5:cce0a2679c717d75
# <翻译结束>


<原文开始>
// SetPath sets the directory path for file logging.
<原文结束>

# <翻译开始>
// SetPath 设置文件日志的目录路径。 md5:817e6d2802241584
# <翻译结束>


<原文开始>
// GetPath returns the logging directory path for file logging.
// It returns empty string if no directory path set.
<原文结束>

# <翻译开始>
// GetPath 返回用于文件日志记录的目录路径。
// 如果未设置目录路径，它将返回空字符串。
// md5:f69da996992ffd9e
# <翻译结束>


<原文开始>
// SetFile sets the file name `pattern` for file logging.
// Datetime pattern can be used in `pattern`, eg: access-{Ymd}.log.
// The default file name pattern is: Y-m-d.log, eg: 2018-01-01.log
<原文结束>

# <翻译开始>
// SetFile 设置文件日志的文件名模式为 `pattern`。
// 在 `pattern` 中可以使用日期时间模式，例如：access-YYYYMMDD.log。
// 默认的文件名模式为：Y-m-d.log，例如：2018-01-01.log
// md5:03b3a973ce783b24
# <翻译结束>


<原文开始>
// SetTimeFormat sets the time format for the logging time.
<原文结束>

# <翻译开始>
// SetTimeFormat 设置日志时间的时间格式。 md5:258a98926fba4588
# <翻译结束>


<原文开始>
// SetStdoutPrint sets whether output the logging contents to stdout, which is true in default.
<原文结束>

# <翻译开始>
// SetStdoutPrint 设置是否将日志内容输出到标准输出，默认为true。 md5:b212437cebfd423a
# <翻译结束>


<原文开始>
// SetHeaderPrint sets whether output header of the logging contents, which is true in default.
<原文结束>

# <翻译开始>
// SetHeaderPrint 设置日志输出的头部是否打印，默认为 true。 md5:3e71cb67564384cc
# <翻译结束>


<原文开始>
// SetLevelPrint sets whether output level string of the logging contents, which is true in default.
<原文结束>

# <翻译开始>
// SetLevelPrint 设置是否输出日志内容的级别字符串，默认为true。 md5:6ba8899e4d3d1c1b
# <翻译结束>


<原文开始>
// SetPrefix sets prefix string for every logging content.
// Prefix is part of header, which means if header output is shut, no prefix will be output.
<原文结束>

# <翻译开始>
// SetPrefix 设置日志内容的前缀字符串。
// 前缀是日志头的一部分，如果关闭了头部输出，就不会显示前缀。
// md5:31d8e3c101c1eea6
# <翻译结束>


<原文开始>
// SetHandlers sets the logging handlers for current logger.
<原文结束>

# <翻译开始>
// SetHandlers 设置当前日志记录器的处理程序。 md5:7b876afcd04a669e
# <翻译结束>


<原文开始>
// SetWriterColorEnable enables file/writer logging with color.
<原文结束>

# <翻译开始>
// SetWriterColorEnable 启用带有颜色的文件/写入器日志记录。 md5:deef19b9707bd4df
# <翻译结束>


<原文开始>
// SetStdoutColorDisabled disables stdout logging with color.
<原文结束>

# <翻译开始>
// SetStdoutColorDisabled 禁用带有颜色的stdout日志记录。 md5:aed9b0e4a2ba0f72
# <翻译结束>

