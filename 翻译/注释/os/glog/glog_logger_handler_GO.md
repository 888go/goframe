
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Handler is function handler for custom logging content outputs.
<原文结束>

# <翻译开始>
// Handler 是用于自定义日志内容输出的函数处理器。 md5:486a8db7f7dd8188
# <翻译结束>


<原文开始>
// HandlerInput is the input parameter struct for logging Handler.
//
// The logging content is consisted in:
// TimeFormat [LevelFormat] {TraceId} {CtxStr} Prefix CallerFunc CallerPath Content Values Stack
//
// The header in the logging content is:
// TimeFormat [LevelFormat] {TraceId} {CtxStr} Prefix CallerFunc CallerPath
<原文结束>

# <翻译开始>
// HandlerInput是日志处理器的输入参数结构体。
//
// 日志内容由以下部分组成：
// 时间格式 [级别格式] {跟踪ID} {上下文字符串} 前缀 调用函数 调用路径 内容 值 堆栈
//
// 日志内容的头部是：
// 时间格式 [级别格式] {跟踪ID} {上下文字符串} 前缀 调用函数 调用路径 md5:6213dd0ebb4e9188
# <翻译结束>


<原文开始>
// Buffer for logging content outputs.
<原文结束>

# <翻译开始>
	// 用于存储日志输出内容的缓冲区。 md5:33224816c1505400
# <翻译结束>


<原文开始>
// (ReadOnly) Logging time, which is the time that logging triggers.
<原文结束>

# <翻译开始>
	// （只读）记录时间，即触发日志记录的时间。 md5:5ce6aaa482dcea28
# <翻译结束>


<原文开始>
// Formatted time string for output, like "2016-01-09 12:00:00".
<原文结束>

# <翻译开始>
	// 格式化的输出时间字符串，如 "2016-01-09 12:00:00"。 md5:530cb544b3906631
# <翻译结束>


<原文开始>
	// (ReadOnly) Using color constant value, like COLOR_RED, COLOR_BLUE, etc.
	// Example: 34
<原文结束>

# <翻译开始>
	//（只读）使用颜色常量值，如COLOR_RED，COLOR_BLUE等。
	// 示例：34 md5:e377684c0eb82b75
# <翻译结束>


<原文开始>
	// (ReadOnly) Using level, like LEVEL_INFO, LEVEL_ERRO, etc.
	// Example: 256
<原文结束>

# <翻译开始>
	// （只读）使用级别，如LEVEL_INFO、LEVEL_ERROR等。
	// 例子：256 md5:20e8f648c34222c9
# <翻译结束>


<原文开始>
	// Formatted level string for output, like "DEBU", "ERRO", etc.
	// Example: ERRO
<原文结束>

# <翻译开始>
	// 用于输出的格式化日志级别字符串，如 "DEBU"、"ERRO" 等。
	// 示例：ERRO md5:0bad424894695e93
# <翻译结束>


<原文开始>
// The source function name that calls logging, only available if F_CALLER_FN set.
<原文结束>

# <翻译开始>
	// 调用日志的源函数名称，仅在设置F_CALLER_FN时可用。 md5:2bfd8148853e8e4c
# <翻译结束>


<原文开始>
	// The source file path and its line number that calls logging,
	// only available if F_FILE_SHORT or F_FILE_LONG set.
<原文结束>

# <翻译开始>
	// 调用日志的源文件路径及其行号，只有在设置F_FILE_SHORT或F_FILE_LONG时可用。 md5:8e31a0cc592be662
# <翻译结束>


<原文开始>
	// The retrieved context value string from context, only available if Config.CtxKeys configured.
	// It's empty if no Config.CtxKeys configured.
<原文结束>

# <翻译开始>
	// 从上下文中获取的已配置的 context 值字符串。如果没有配置 Config.CtxKeys，它将为空。 md5:b854bd1bcad06fda
# <翻译结束>


<原文开始>
// Trace id, only available if OpenTelemetry is enabled, or else it's an empty string.
<原文结束>

# <翻译开始>
	// 跟踪ID，仅在启用OpenTelemetry时可用，否则为空字符串。 md5:0cd8e77f80286121
# <翻译结束>


<原文开始>
	// Custom prefix string in logging content header part.
	// Note that, it takes no effect if HeaderPrint is disabled.
<原文结束>

# <翻译开始>
	// 在日志内容头部自定义的前缀字符串。
	// 请注意，如果已禁用HeaderPrint，此设置将不会生效。 md5:004eed7afe3ca2dd
# <翻译结束>


<原文开始>
// Custom logging content for logging.
<原文结束>

# <翻译开始>
	// 用于日志记录的自定义日志内容。 md5:9749c3bafd8e33d5
# <翻译结束>


<原文开始>
// The passed un-formatted values array to logger.
<原文结束>

# <翻译开始>
	// 传递给日志记录器的未经格式化的值数组。 md5:854ab8e84e01371d
# <翻译结束>


<原文开始>
	// Stack string produced by logger, only available if Config.StStatus configured.
	// Note that there are usually multiple lines in stack content.
<原文结束>

# <翻译开始>
	// 由记录器生成的堆栈字符串，仅在配置了Config.StStatus时可用。
	// 注意，堆栈内容中通常包含多行。 md5:c36e69fdfae3ac16
# <翻译结束>


<原文开始>
// IsAsync marks it is in asynchronous logging.
<原文结束>

# <翻译开始>
	// IsAsync 标记为异步日志记录。 md5:e138a9a968506347
# <翻译结束>


<原文开始>
// Middleware handling index for internal usage.
<原文结束>

# <翻译开始>
// 处理内部使用的索引的中间件。 md5:61d366e59aee7159
# <翻译结束>


<原文开始>
// Handler array calling bu index.
<原文结束>

# <翻译开始>
// 通过索引调用处理器数组。 md5:7cb772c2e129fd27
# <翻译结束>


<原文开始>
// defaultHandler is the default handler for package.
<原文结束>

# <翻译开始>
// defaultHandler 是包的默认处理程序。 md5:0f4cafed00a48af2
# <翻译结束>


<原文开始>
// doFinalPrint is a handler for logging content printing.
// This handler outputs logging content to file/stdout/write if any of them configured.
<原文结束>

# <翻译开始>
// doFinalPrint 是用于记录内容打印的处理器。
// 此处理器将日志内容输出到文件/stdout/write，如果它们中有任何被配置的话。 md5:794b81b9fa0a2bd6
# <翻译结束>


<原文开始>
// SetDefaultHandler sets default handler for package.
<原文结束>

# <翻译开始>
// SetDefaultHandler 设置包的默认处理器。 md5:33a213aebe83e5ed
# <翻译结束>


<原文开始>
// GetDefaultHandler returns the default handler of package.
<原文结束>

# <翻译开始>
// GetDefaultHandler 返回包的默认处理器。 md5:8812c42db1189f3b
# <翻译结束>


<原文开始>
// Next calls the next logging handler in middleware way.
<原文结束>

# <翻译开始>
// Next 以中间件的方式调用下一个日志处理程序。 md5:ab91f9dfe65c4322
# <翻译结束>


<原文开始>
// String returns the logging content formatted by default logging handler.
<原文结束>

# <翻译开始>
// String 返回默认日志处理器格式化的日志内容。 md5:e78613962fe54276
# <翻译结束>


<原文开始>
// ValuesContent converts and returns values as string content.
<原文结束>

# <翻译开始>
// ValuesContent 将值转换为字符串内容并返回。 md5:da3a0fd9093d35c9
# <翻译结束>


<原文开始>
// Remove one blank line(\n\n).
<原文结束>

# <翻译开始>
		// 移除一个空行（\n\n）。 md5:777d73ee86014d2c
# <翻译结束>


<原文开始>
// avoid a single space at the end of a line.
<原文结束>

# <翻译开始>
	// 避免在行尾留下单个空格。 md5:f107ec37b9775773
# <翻译结束>

