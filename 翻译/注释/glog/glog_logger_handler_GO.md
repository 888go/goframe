
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
// Handler is function handler for custom logging content outputs.
<原文结束>

# <翻译开始>
// Handler 是用于自定义日志内容输出的函数处理器。
# <翻译结束>


<原文开始>
// HandlerInput is the input parameter struct for logging Handler.
<原文结束>

# <翻译开始>
// HandlerInput 是 logging Handler 的输入参数结构体。
# <翻译结束>







<原文开始>
// Buffer for logging content outputs.
<原文结束>

# <翻译开始>
// Buffer，用于日志内容输出。
# <翻译结束>


<原文开始>
// Logging time, which is the time that logging triggers.
<原文结束>

# <翻译开始>
// 日志时间，即触发日志记录的时间。
# <翻译结束>


<原文开始>
// Using color, like COLOR_RED, COLOR_BLUE, etc. Eg: 34
<原文结束>

# <翻译开始>
// 使用颜色，如COLOR_RED、COLOR_BLUE等。例如：34
# <翻译结束>


<原文开始>
// Using level, like LEVEL_INFO, LEVEL_ERRO, etc. Eg: 256
<原文结束>

# <翻译开始>
// 使用级别，如 LEVEL_INFO, LEVEL_ERRO 等。例如：256
# <翻译结束>


<原文开始>
// The source function name that calls logging, only available if F_CALLER_FN set.
<原文结束>

# <翻译开始>
// 如果设置了F_CALLER_FN，该变量记录调用日志函数的源函数名。
# <翻译结束>


<原文开始>
// The source file path and its line number that calls logging, only available if F_FILE_SHORT or F_FILE_LONG set.
<原文结束>

# <翻译开始>
// 调用日志记录的源文件路径及其行号，仅在设置了 F_FILE_SHORT 或 F_FILE_LONG 时可用。
# <翻译结束>


<原文开始>
// The retrieved context value string from context, only available if Config.CtxKeys configured.
<原文结束>

# <翻译开始>
// 从context中获取的字符串类型的上下文值，但只有在配置了Config.CtxKeys时才可用。
# <翻译结束>


<原文开始>
// Trace id, only available if OpenTelemetry is enabled.
<原文结束>

# <翻译开始>
// 跟踪ID，仅在启用OpenTelemetry时可用。
# <翻译结束>


<原文开始>
// Custom prefix string for logging content.
<原文结束>

# <翻译开始>
// 自定义日志内容前缀字符串。
# <翻译结束>


<原文开始>
// Content is the main logging content without error stack string produced by logger.
<原文结束>

# <翻译开始>
// Content 是由 logger 生成的、不包含错误堆栈信息的主要日志内容。
# <翻译结束>


<原文开始>
// The passed un-formatted values array to logger.
<原文结束>

# <翻译开始>
// 传递给 logger 的未格式化的值数组。
# <翻译结束>


<原文开始>
// Stack string produced by logger, only available if Config.StStatus configured.
<原文结束>

# <翻译开始>
// Stack 字符串由 logger 生成，仅在配置了 Config.StStatus 时可用。
# <翻译结束>


<原文开始>
// IsAsync marks it is in asynchronous logging.
<原文结束>

# <翻译开始>
// IsAsync 标记它处于异步日志记录状态。
# <翻译结束>


<原文开始>
// Middleware handling index for internal usage.
<原文结束>

# <翻译开始>
// 此中间件用于内部使用，处理索引功能。
# <翻译结束>


<原文开始>
// Handler array calling bu index.
<原文结束>

# <翻译开始>
// 通过索引调用处理器数组
# <翻译结束>


<原文开始>
// defaultHandler is the default handler for package.
<原文结束>

# <翻译开始>
// defaultHandler 是该包的默认处理器。
# <翻译结束>


<原文开始>
// doFinalPrint is a handler for logging content printing.
// This handler outputs logging content to file/stdout/write if any of them configured.
<原文结束>

# <翻译开始>
// doFinalPrint 是一个用于记录内容打印的处理器。
// 如果其中任意一项被配置，此处理器会将日志内容输出到文件、标准输出(stdout)或写入指定位置。
# <翻译结束>


<原文开始>
// SetDefaultHandler sets default handler for package.
<原文结束>

# <翻译开始>
// SetDefaultHandler 设置包的默认处理器。
# <翻译结束>


<原文开始>
// GetDefaultHandler returns the default handler of package.
<原文结束>

# <翻译开始>
// GetDefaultHandler 返回该包的默认处理器。
# <翻译结束>


<原文开始>
// Next calls the next logging handler in middleware way.
<原文结束>

# <翻译开始>
// Next 以中间件方式调用下一个日志处理程序。
# <翻译结束>


<原文开始>
// String returns the logging content formatted by default logging handler.
<原文结束>

# <翻译开始>
// String 返回由默认日志处理程序格式化的日志内容。
# <翻译结束>


<原文开始>
// Convert values string content.
<原文结束>

# <翻译开始>
// 将values字符串内容进行转换
# <翻译结束>


<原文开始>
// Remove one blank line(\n\n).
<原文结束>

# <翻译开始>
// 删除一个空行（\n\n）
# <翻译结束>


<原文开始>
// avoid a single space at the end of a line.
<原文结束>

# <翻译开始>
// 避免在行尾出现单个空格。
# <翻译结束>


<原文开始>
// Formatted time string, like "2016-01-09 12:00:00".
<原文结束>

# <翻译开始>
// 格式化的时间字符串，如 "2016-01-09 12:00:00"。
# <翻译结束>


<原文开始>
// Formatted level string, like "DEBU", "ERRO", etc. Eg: ERRO
<原文结束>

# <翻译开始>
// 格式化的级别字符串，如 "DEBU", "ERRO" 等。例如：ERRO
# <翻译结束>


<原文开始>
// Current Logger object.
<原文结束>

# <翻译开始>
// 当前日志器对象。
# <翻译结束>

