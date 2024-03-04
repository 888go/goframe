
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
// SetConfig set configurations for the defaultLogger.
<原文结束>

# <翻译开始>
// SetConfig 设置默认日志器的配置。
# <翻译结束>


<原文开始>
// SetConfigWithMap set configurations with map for the defaultLogger.
<原文结束>

# <翻译开始>
// SetConfigWithMap 通过map为默认日志器设置配置。
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
// SetLevel sets the default logging level.
<原文结束>

# <翻译开始>
// SetLevel 设置默认的日志级别。
# <翻译结束>


<原文开始>
// GetLevel returns the default logging level value.
<原文结束>

# <翻译开始>
// GetLevel 返回默认的日志级别值。
# <翻译结束>


<原文开始>
// SetWriter sets the customized logging `writer` for logging.
// The `writer` object should implements the io.Writer interface.
// Developer can use customized logging `writer` to redirect logging output to another service,
// eg: kafka, mysql, mongodb, etc.
<原文结束>

# <翻译开始>
// SetWriter 设置自定义的日志 `writer` 用于日志记录。
// `writer` 对象应实现 io.Writer 接口。
// 开发者可以使用自定义的日志 `writer` 将日志输出重定向到其他服务，
// 例如：kafka、mysql、mongodb 等等。
# <翻译结束>


<原文开始>
// GetWriter returns the customized writer object, which implements the io.Writer interface.
// It returns nil if no customized writer set.
<原文结束>

# <翻译开始>
// GetWriter 返回一个自定义的writer对象，该对象实现了io.Writer接口。
// 如果未设置自定义writer，则返回nil。
# <翻译结束>


<原文开始>
// SetDebug enables/disables the debug level for default defaultLogger.
// The debug level is enabled in default.
<原文结束>

# <翻译开始>
// SetDebug 用于启用/禁用默认 defaultLogger 的调试级别。
// 调试级别默认是启用状态。
# <翻译结束>


<原文开始>
// SetAsync enables/disables async logging output feature for default defaultLogger.
<原文结束>

# <翻译开始>
// SetAsync 用于启用/禁用默认默认Logger的异步日志输出功能。
# <翻译结束>


<原文开始>
// SetStdoutPrint sets whether ouptput the logging contents to stdout, which is true in default.
<原文结束>

# <翻译开始>
// SetStdoutPrint 设置是否将日志内容输出到标准输出（stdout），默认情况下为true。
# <翻译结束>


<原文开始>
// SetHeaderPrint sets whether output header of the logging contents, which is true in default.
<原文结束>

# <翻译开始>
// SetHeaderPrint 设置是否输出日志内容的头部，默认为true。
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
// SetFlags sets extra flags for logging output features.
<原文结束>

# <翻译开始>
// SetFlags 设置日志输出功能的额外标志。
# <翻译结束>


<原文开始>
// GetFlags returns the flags of defaultLogger.
<原文结束>

# <翻译开始>
// GetFlags 返回默认日志器的标志。
# <翻译结束>


<原文开始>
// SetCtxKeys sets the context keys for defaultLogger. The keys is used for retrieving values
// from context and printing them to logging content.
//
// Note that multiple calls of this function will overwrite the previous set context keys.
<原文结束>

# <翻译开始>
// SetCtxKeys 为 defaultLogger 设置上下文键。这些键用于从上下文中检索值并将其打印到日志内容中。
//
// 注意，多次调用此函数会覆盖之前设置的上下文键。
# <翻译结束>


<原文开始>
// GetCtxKeys retrieves and returns the context keys for logging.
<原文结束>

# <翻译开始>
// GetCtxKeys 获取并返回用于日志记录的上下文键。
# <翻译结束>


<原文开始>
// PrintStack prints the caller stack,
// the optional parameter `skip` specify the skipped stack offset from the end point.
<原文结束>

# <翻译开始>
// PrintStack 打印调用栈，
// 可选参数 `skip` 指定了从终点开始需要跳过的堆栈偏移量。
# <翻译结束>


<原文开始>
// GetStack returns the caller stack content,
// the optional parameter `skip` specify the skipped stack offset from the end point.
<原文结束>

# <翻译开始>
// GetStack 返回调用堆栈的内容，
// 可选参数 `skip` 指定了从终点开始跳过的堆栈偏移量。
# <翻译结束>


<原文开始>
// SetStack enables/disables the stack feature in failure logging outputs.
<原文结束>

# <翻译开始>
// SetStack 启用/禁用失败日志输出中的堆栈跟踪功能。
# <翻译结束>


<原文开始>
// SetLevelStr sets the logging level by level string.
<原文结束>

# <翻译开始>
// SetLevelStr 通过级别字符串设置日志记录级别。
# <翻译结束>


<原文开始>
// SetLevelPrefix sets the prefix string for specified level.
<原文结束>

# <翻译开始>
// SetLevelPrefix 为指定等级设置前缀字符串。
# <翻译结束>


<原文开始>
// SetLevelPrefixes sets the level to prefix string mapping for the defaultLogger.
<原文结束>

# <翻译开始>
// SetLevelPrefixes 为默认日志器设置等级到前缀字符串的映射。
# <翻译结束>


<原文开始>
// GetLevelPrefix returns the prefix string for specified level.
<原文结束>

# <翻译开始>
// GetLevelPrefix 返回指定级别的前缀字符串。
# <翻译结束>


<原文开始>
// SetHandlers sets the logging handlers for default defaultLogger.
<原文结束>

# <翻译开始>
// SetHandlers 为默认的 defaultLogger 设置日志处理器。
# <翻译结束>


<原文开始>
// SetWriterColorEnable sets the file logging with color
<原文结束>

# <翻译开始>
// SetWriterColorEnable 设置文件日志记录使用颜色
# <翻译结束>

