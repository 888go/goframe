
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
// Logger is the struct for logging management.
<原文结束>

# <翻译开始>
// Logger 是用于日志管理的结构体。 md5:c338807dca943ae3
# <翻译结束>


<原文开始>
// Parent logger, if it is not empty, it means the logger is used in chaining function.
<原文结束>

# <翻译开始>
// 父日志记录器，如果非空，表示该日志记录器用于链式调用功能。 md5:9efa0349702d1c2c
# <翻译结束>


<原文开始>
// Print logging content asynchronously。
<原文结束>

# <翻译开始>
// 异步打印日志内容。 md5:aae4973631499f41
# <翻译结束>


<原文开始>
// Print full file name and line number: /a/b/c/d.go:23.
<原文结束>

# <翻译开始>
// 打印完整的文件名和行号：/a/b/c/d.go:23。 md5:b0d552af751b5a59
# <翻译结束>


<原文开始>
// Print final file name element and line number: d.go:23. overrides F_FILE_LONG.
<原文结束>

# <翻译开始>
// 打印最终的文件名元素和行号：d.go:23。会覆盖F_FILE_LONG设置。 md5:b2f804c31e821aba
# <翻译结束>


<原文开始>
// Print the date in the local time zone: 2009-01-23.
<原文结束>

# <翻译开始>
// 在本地时区打印日期：2009-01-23。 md5:d7310c166de59388
# <翻译结束>


<原文开始>
// Print the time in the local time zone: 01:23:23.
<原文结束>

# <翻译开始>
// 在本地时区打印时间：01:23:23。 md5:547ab7e84bd2a146
# <翻译结束>


<原文开始>
// Print the time with milliseconds in the local time zone: 01:23:23.675.
<原文结束>

# <翻译开始>
// 在本地时区打印带有毫秒的时间：01:23:23.675。 md5:063a060ef145f39d
# <翻译结束>


<原文开始>
// Print Caller function name and package: main.main
<原文结束>

# <翻译开始>
// 打印调用者函数名称和包：main.main. md5:1d729cf0b4dce445
# <翻译结束>


<原文开始>
// New creates and returns a custom logger.
<原文结束>

# <翻译开始>
// New 创建并返回一个自定义的日志器。 md5:0205650422cdd95e
# <翻译结束>


<原文开始>
// NewWithWriter creates and returns a custom logger with io.Writer.
<原文结束>

# <翻译开始>
// NewWithWriter 创建并返回一个具有 io.Writer 的自定义日志器。 md5:51edfcbd62ded572
# <翻译结束>


<原文开始>
// Clone returns a new logger, which a `shallow copy` of the current logger.
// Note that the attribute `config` of the cloned one is the shallow copy of current one.
<原文结束>

# <翻译开始>
// Clone 返回一个新的记录器，它是当前记录器的`浅拷贝`。
// 注意，克隆体的`config`属性是对当前记录器配置的浅拷贝。
// md5:c70ded0c6903f4be
# <翻译结束>


<原文开始>
// getFilePath returns the logging file path.
// The logging file name must have extension name of "log".
<原文结束>

# <翻译开始>
// getFilePath 返回日志文件的路径。
// 日志文件的名称必须带有".log"扩展名。
// md5:f3fb565d6de6de8e
# <翻译结束>


<原文开始>
// Content containing "{}" in the file name is formatted using gtime.
<原文结束>

# <翻译开始>
// 文件名中包含“{}”的内容将使用gtime进行格式化。 md5:9502dc758bde7fca
# <翻译结束>


<原文开始>
// print prints `s` to defined writer, logging file or passed `std`.
<原文结束>

# <翻译开始>
// print 将`s`打印到已定义的writer（日志文件）或传递的`std`。 md5:2368d31e4b600609
# <翻译结束>


<原文开始>
	// Lazy initialize for rotation feature.
	// It uses atomic reading operation to enhance the performance checking.
	// It here uses CAP for performance and concurrent safety.
	// It just initializes once for each logger.
<原文结束>

# <翻译开始>
	// 延迟初始化旋转特性。
	// 它采用原子读取操作来增强性能检查。
	// 此处使用CAP以确保性能和并发安全性。
	// 每个日志器仅初始化一次。
	// md5:1562dbed8f576bc2
# <翻译结束>


<原文开始>
// Caller path and Fn name.
<原文结束>

# <翻译开始>
// 调用者路径和函数名。 md5:92543c6732cddd2d
# <翻译结束>


<原文开始>
// Convert value to string.
<原文结束>

# <翻译开始>
// 将值转换为字符串。 md5:c172ad8437ce5ebf
# <翻译结束>


<原文开始>
// doFinalPrint outputs the logging content according configuration.
<原文结束>

# <翻译开始>
// doFinalPrint 根据配置输出日志内容。 md5:d442b45c82ee936e
# <翻译结束>


<原文开始>
// Allow output to stdout?
<原文结束>

# <翻译开始>
// 允许输出到标准输出（stdout）吗？. md5:0f9801ce26bfb625
# <翻译结束>


<原文开始>
// Output content to disk file.
<原文结束>

# <翻译开始>
// 将内容输出到磁盘文件。 md5:eac9c01dcbb64a4f
# <翻译结束>


<原文开始>
// Output to custom writer.
<原文结束>

# <翻译开始>
// 向自定义写入器输出。 md5:ef7c09b52f5d355b
# <翻译结束>


<原文开始>
// printToWriter writes buffer to writer.
<原文结束>

# <翻译开始>
// printToWriter 将缓冲区写入writer。 md5:5b62d7e007bf0275
# <翻译结束>


<原文开始>
// printToStdout outputs logging content to stdout.
<原文结束>

# <翻译开始>
// printToStdout 将日志内容输出到标准输出（stdout）。 md5:2fb96a5229ec7af7
# <翻译结束>


<原文开始>
		// This will lose color in Windows os system. DO NOT USE.
		// if _, err := os.Stdout.Write(input.getRealBuffer(true).Bytes()); err != nil {
<原文结束>

# <翻译开始>
		// 这将在Windows操作系统中丢失颜色信息。请勿使用。
		// if _, err := os.Stdout.Write(input.getRealBuffer(true).Bytes()); err != nil {
		// md5:29dd90df2339a223
# <翻译结束>


<原文开始>
// This will print color in Windows os system.
<原文结束>

# <翻译开始>
// 这将在Windows操作系统中打印颜色。 md5:c2abebac838c5747
# <翻译结束>


<原文开始>
// printToFile outputs logging content to disk file.
<原文结束>

# <翻译开始>
// printToFile 将日志内容输出到磁盘文件中。 md5:20d4379608c45b69
# <翻译结束>


<原文开始>
// Rotation file size checks.
<原文结束>

# <翻译开始>
// 旋转文件大小检查。 md5:82f7b948ac1657a5
# <翻译结束>


<原文开始>
// Logging content outputting to disk file.
<原文结束>

# <翻译开始>
// 将日志内容输出到磁盘文件。 md5:c3b494b8895acc38
# <翻译结束>


<原文开始>
// createFpInPool retrieves and returns a file pointer from file pool.
<原文结束>

# <翻译开始>
// createFpInPool 从文件池中检索并返回一个文件指针。 md5:4acfc3ca923fca99
# <翻译结束>


<原文开始>
// getFpFromPool retrieves and returns a file pointer from file pool.
<原文结束>

# <翻译开始>
// getFpFromPool 从文件池中检索并返回一个文件指针。 md5:e3cccf00ba4439d4
# <翻译结束>


<原文开始>
// printStd prints content `s` without stack.
<原文结束>

# <翻译开始>
// printStd 在不打印调用栈的情况下输出内容 `s`。 md5:f9e76daf3a16b514
# <翻译结束>


<原文开始>
// printErr prints content `s` with stack check.
<原文结束>

# <翻译开始>
// printErr 打印内容 `s`，并检查堆栈信息。 md5:d9883612e4e85bcb
# <翻译结束>


<原文开始>
// In matter of sequence, do not use stderr here, but use the same stdout.
<原文结束>

# <翻译开始>
// 从顺序上来说，这里不要使用stderr，而是要使用相同的stdout。 md5:1f1258ae1ca0856e
# <翻译结束>


<原文开始>
// format formats `values` using fmt.Sprintf.
<原文结束>

# <翻译开始>
// format 使用fmt.Sprintf格式化`values`。 md5:bd4bb80582842100
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
// Whether filter framework error stacks.
<原文结束>

# <翻译开始>
// 是否过滤框架错误堆栈。 md5:7cf82960065281e9
# <翻译结束>

