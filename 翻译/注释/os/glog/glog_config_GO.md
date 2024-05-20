
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
// SetConfig set configurations for the defaultLogger.
<原文结束>

# <翻译开始>
// SetConfig 为默认日志器设置配置。. md5:5f388c6c2d28724c
# <翻译结束>


<原文开始>
// SetConfigWithMap set configurations with map for the defaultLogger.
<原文结束>

# <翻译开始>
// 使用映射为defaultLogger设置配置。. md5:195834b40cfce621
# <翻译结束>


<原文开始>
// SetPath sets the directory path for file logging.
<原文结束>

# <翻译开始>
// SetPath 设置文件日志的目录路径。. md5:817e6d2802241584
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
// SetLevel sets the default logging level.
<原文结束>

# <翻译开始>
// SetLevel 设置默认的日志级别。. md5:ab428766cae30d65
# <翻译结束>


<原文开始>
// GetLevel returns the default logging level value.
<原文结束>

# <翻译开始>
// GetLevel 返回默认的日志级别值。. md5:5ae9c76de12ac5c9
# <翻译结束>


<原文开始>
// SetWriter sets the customized logging `writer` for logging.
// The `writer` object should implements the io.Writer interface.
// Developer can use customized logging `writer` to redirect logging output to another service,
// eg: kafka, mysql, mongodb, etc.
<原文结束>

# <翻译开始>
// SetWriter 设置自定义的日志记录器`writer`。
// `writer`对象应实现io.Writer接口。
// 开发者可以使用自定义的`writer`将日志输出重定向到其他服务，如：kafka、mysql、mongodb等。
// md5:50799f5f4e6280ea
# <翻译结束>


<原文开始>
// GetWriter returns the customized writer object, which implements the io.Writer interface.
// It returns nil if no customized writer set.
<原文结束>

# <翻译开始>
// GetWriter 返回自定义的 writer 对象，它实现了 io.Writer 接口。
// 如果没有设置自定义 writer，它将返回 nil。
// md5:37b3d820b5547c91
# <翻译结束>


<原文开始>
// SetDebug enables/disables the debug level for default defaultLogger.
// The debug level is enabled in default.
<原文结束>

# <翻译开始>
// SetDebug 启用/禁用默认 defaultLogger 的调试级别。默认情况下，调试级别已启用。
// md5:335740e0731656b1
# <翻译结束>


<原文开始>
// SetAsync enables/disables async logging output feature for default defaultLogger.
<原文结束>

# <翻译开始>
// SetAsync 用于启用或禁用默认默认日志记录器的异步日志输出功能。. md5:2659994b118a6211
# <翻译结束>


<原文开始>
// SetStdoutPrint sets whether ouptput the logging contents to stdout, which is true in default.
<原文结束>

# <翻译开始>
// SetStdoutPrint 设置是否将日志内容输出到stdout，默认为true。. md5:f68d5bea1f6372d1
# <翻译结束>


<原文开始>
// SetHeaderPrint sets whether output header of the logging contents, which is true in default.
<原文结束>

# <翻译开始>
// SetHeaderPrint 设置日志输出的头部是否打印，默认为 true。. md5:3e71cb67564384cc
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
// SetFlags sets extra flags for logging output features.
<原文结束>

# <翻译开始>
// SetFlags 为日志输出功能设置额外的标志。. md5:40253d4ed662de77
# <翻译结束>


<原文开始>
// GetFlags returns the flags of defaultLogger.
<原文结束>

# <翻译开始>
// GetFlags 返回默认日志器defaultLogger的标志。. md5:ca4e9bc9260b114a
# <翻译结束>


<原文开始>
// SetCtxKeys sets the context keys for defaultLogger. The keys is used for retrieving values
// from context and printing them to logging content.
//
// Note that multiple calls of this function will overwrite the previous set context keys.
<原文结束>

# <翻译开始>
// SetCtxKeys 为 defaultLogger 设置上下文键。这些键用于从上下文中检索值，并将它们打印到日志内容中。
// 
// 注意，多次调用此函数会覆盖之前设置的上下文键。
// md5:8a4710a99dd19d4d
# <翻译结束>


<原文开始>
// GetCtxKeys retrieves and returns the context keys for logging.
<原文结束>

# <翻译开始>
// GetCtxKeys检索并返回日志记录的上下文键。. md5:1e780871ada8c59c
# <翻译结束>


<原文开始>
// PrintStack prints the caller stack,
// the optional parameter `skip` specify the skipped stack offset from the end point.
<原文结束>

# <翻译开始>
// PrintStack 打印调用堆栈，
// 可选参数 `skip` 用于指定从堆栈终点开始忽略的偏移量。
// md5:ef6cd40820765783
# <翻译结束>


<原文开始>
// GetStack returns the caller stack content,
// the optional parameter `skip` specify the skipped stack offset from the end point.
<原文结束>

# <翻译开始>
// GetStack 返回调用者栈的内容，
// 可选参数 `skip` 指定从终点开始要跳过的栈偏移量。
// md5:13592be3061e779d
# <翻译结束>


<原文开始>
// SetStack enables/disables the stack feature in failure logging outputs.
<原文结束>

# <翻译开始>
// SetStack 启用/禁用失败日志输出中的堆栈功能。. md5:3c80a664fff650de
# <翻译结束>


<原文开始>
// SetLevelStr sets the logging level by level string.
<原文结束>

# <翻译开始>
// SetLevelStr 通过级别字符串设置日志级别。. md5:53cbbdf23584340e
# <翻译结束>


<原文开始>
// SetLevelPrefix sets the prefix string for specified level.
<原文结束>

# <翻译开始>
// SetLevelPrefix 为指定的日志级别设置前缀字符串。. md5:a2b7a43af150bcb7
# <翻译结束>


<原文开始>
// SetLevelPrefixes sets the level to prefix string mapping for the defaultLogger.
<原文结束>

# <翻译开始>
// SetLevelPrefixes 设置默认logger的级别到前缀字符串的映射。. md5:5d55474512859352
# <翻译结束>


<原文开始>
// GetLevelPrefix returns the prefix string for specified level.
<原文结束>

# <翻译开始>
// GetLevelPrefix 返回指定级别的前缀字符串。. md5:339b86b4f84d6049
# <翻译结束>


<原文开始>
// SetHandlers sets the logging handlers for default defaultLogger.
<原文结束>

# <翻译开始>
// SetHandlers 设置默认日志记录器defaultLogger的处理程序。. md5:637a5dc8ccc73e8f
# <翻译结束>


<原文开始>
// SetWriterColorEnable sets the file logging with color
<原文结束>

# <翻译开始>
// SetWriterColorEnable 设置文件日志记录使用颜色. md5:3ff25da59932914b
# <翻译结束>

