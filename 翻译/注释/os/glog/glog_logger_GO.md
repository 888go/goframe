
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
// Logger is the struct for logging management.
<原文结束>

# <翻译开始>
// Logger 是用于日志管理的结构体。
# <翻译结束>


<原文开始>
// Parent logger, if it is not empty, it means the logger is used in chaining function.
<原文结束>

# <翻译开始>
// 父级日志器，如果非空，则表示该日志器在链式函数中使用。
# <翻译结束>







<原文开始>
// Print full file name and line number: /a/b/c/d.go:23.
<原文结束>

# <翻译开始>
// 打印完整的文件名和行号：/a/b/c/d.go:23.
# <翻译结束>


<原文开始>
// Print final file name element and line number: d.go:23. overrides F_FILE_LONG.
<原文结束>

# <翻译开始>
// 打印文件名的最后一个元素和行号：d.go:23。该选项覆盖了 F_FILE_LONG。
# <翻译结束>


<原文开始>
// Print the date in the local time zone: 2009-01-23.
<原文结束>

# <翻译开始>
// 在本地时区打印日期：2009-01-23。
# <翻译结束>


<原文开始>
// Print the time in the local time zone: 01:23:23.
<原文结束>

# <翻译开始>
// 在本地时区打印时间：01:23:23
// ```go
// 下面是示例代码（假设）
// package main
// import (
//     "fmt"
//     "time"
// )
// func main() {
    // 获取当前时间
//     now := time.Now()
    // 格式化为 HH:mm:ss 的格式
//     formattedTime := now.Format("15:04:05")
    // 打印在本地时区的时间：01:23:23
//     fmt.Println(formattedTime)
// }
// 这段注释描述的是要在本地时区打印时间的意图，但实际代码需要获取当前时间并进行格式化以匹配指定格式。
# <翻译结束>


<原文开始>
// Print the time with milliseconds in the local time zone: 01:23:23.675.
<原文结束>

# <翻译开始>
// 在本地时区打印带有毫秒的时间：01:23:23.675。
# <翻译结束>


<原文开始>
// Print Caller function name and package: main.main
<原文结束>

# <翻译开始>
// 打印调用函数名称和包名：main.main
# <翻译结束>


<原文开始>
// New creates and returns a custom logger.
<原文结束>

# <翻译开始>
// New 创建并返回一个自定义日志器。
# <翻译结束>


<原文开始>
// NewWithWriter creates and returns a custom logger with io.Writer.
<原文结束>

# <翻译开始>
// NewWithWriter 通过 io.Writer 创建并返回一个自定义的日志记录器。
# <翻译结束>


<原文开始>
// Clone returns a new logger, which a `shallow copy` of the current logger.
// Note that the attribute `config` of the cloned one is the shallow copy of current one.
<原文结束>

# <翻译开始>
// Clone 返回一个新的日志器，它是当前日志器的“浅复制”。
// 注意，克隆后的日志器其 `config` 属性是对当前日志器该属性的浅复制。
# <翻译结束>


<原文开始>
// getFilePath returns the logging file path.
// The logging file name must have extension name of "log".
<原文结束>

# <翻译开始>
// getFilePath 返回日志文件的路径。
// 日志文件名必须包含 ".log" 扩展名。
# <翻译结束>


<原文开始>
// Content containing "{}" in the file name is formatted using gtime.
<原文结束>

# <翻译开始>
// 文件名中包含 "{}" 的内容将使用 gtime 进行格式化。
# <翻译结束>


<原文开始>
// print prints `s` to defined writer, logging file or passed `std`.
<原文结束>

# <翻译开始>
// print 将 `s` 打印输出到预定义的 writer（写入器）、日志文件或传入的 `std`。
# <翻译结束>


<原文开始>
	// Lazy initialize for rotation feature.
	// It uses atomic reading operation to enhance the performance checking.
	// It here uses CAP for performance and concurrent safety.
	// It just initializes once for each logger.
<原文结束>

# <翻译开始>
// 为日志旋转功能进行惰性初始化。
// 使用原子读取操作以提升性能检查的效率。
// 这里使用了CAP以保证性能和并发安全性。
// 对于每个日志器，仅初始化一次。
# <翻译结束>

















<原文开始>
// doFinalPrint outputs the logging content according configuration.
<原文结束>

# <翻译开始>
// doFinalPrint 根据配置输出日志内容。
# <翻译结束>







<原文开始>
// Output content to disk file.
<原文结束>

# <翻译开始>
// 将内容输出到磁盘文件。
# <翻译结束>












<原文开始>
// printToWriter writes buffer to writer.
<原文结束>

# <翻译开始>
// printToWriter 将缓冲区内容写入到writer中。
# <翻译结束>


<原文开始>
// printToStdout outputs logging content to stdout.
<原文结束>

# <翻译开始>
// printToStdout 将日志内容输出到 stdout（标准输出）。
# <翻译结束>


<原文开始>
		// This will lose color in Windows os system. DO NOT USE.
		// if _, err := os.Stdout.Write(input.getRealBuffer(true).Bytes()); err != nil {
<原文结束>

# <翻译开始>
// 这将在Windows操作系统中丢失颜色。请勿使用。
// 如果 _, err := os.Stdout.Write(input.getRealBuffer(true).Bytes()); 出现错误err，则不为nil {
# <翻译结束>


<原文开始>
// This will print color in Windows os system.
<原文结束>

# <翻译开始>
// 这将在Windows操作系统中打印颜色。
# <翻译结束>


<原文开始>
// printToFile outputs logging content to disk file.
<原文结束>

# <翻译开始>
// printToFile 将日志内容输出到磁盘文件。
# <翻译结束>







<原文开始>
// Logging content outputting to disk file.
<原文结束>

# <翻译开始>
// 将日志内容输出到磁盘文件。
# <翻译结束>


<原文开始>
// createFpInPool retrieves and returns a file pointer from file pool.
<原文结束>

# <翻译开始>
// createFpInPool 从文件池中获取并返回一个文件指针。
# <翻译结束>


<原文开始>
// getFpFromPool retrieves and returns a file pointer from file pool.
<原文结束>

# <翻译开始>
// getFpFromPool 从文件池中获取并返回一个文件指针。
# <翻译结束>


<原文开始>
// printStd prints content `s` without stack.
<原文结束>

# <翻译开始>
// printStd 不带堆栈地打印内容`s`。
# <翻译结束>


<原文开始>
// printStd prints content `s` with stack check.
<原文结束>

# <翻译开始>
// printStd 在进行堆栈检查的情况下打印内容`s`。
# <翻译结束>


<原文开始>
// In matter of sequence, do not use stderr here, but use the same stdout.
<原文结束>

# <翻译开始>
// 在顺序输出方面，此处不要使用 stderr，而应使用相同的 stdout。
# <翻译结束>


<原文开始>
// format formats `values` using fmt.Sprintf.
<原文结束>

# <翻译开始>
// format 使用 fmt.Sprintf 对 `values` 进行格式化。
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
// Whether filter framework error stacks.
<原文结束>

# <翻译开始>
// 是否过滤框架错误堆栈。
# <翻译结束>












<原文开始>
// Logger configuration.
<原文结束>

# <翻译开始>
// 日志器配置。
# <翻译结束>


<原文开始>
// Caller path and Fn name.
<原文结束>

# <翻译开始>
// 调用路径和函数名称。
# <翻译结束>


<原文开始>
// Convert value to string.
<原文结束>

# <翻译开始>
// 将值转换为字符串。
# <翻译结束>


<原文开始>
// Allow output to stdout?
<原文结束>

# <翻译开始>
// 是否允许输出到标准输出（stdout）？
# <翻译结束>


<原文开始>
// Used custom writer.
<原文结束>

# <翻译开始>
// 使用了自定义的写入器。
# <翻译结束>


<原文开始>
// Output to custom writer.
<原文结束>

# <翻译开始>
// 将输出发送到自定义写入器。
# <翻译结束>


<原文开始>
// Rotation file size checks.
<原文结束>

# <翻译开始>
// 旋转文件大小检查。
# <翻译结束>

