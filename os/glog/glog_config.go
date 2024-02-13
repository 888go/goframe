// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类

import (
	"context"
	"io"
)

// SetConfig 设置默认日志器的配置。
func X设置配置项(配置项 Config) error {
	return defaultLogger.X设置配置项(配置项)
}

// SetConfigWithMap 通过map为默认日志器设置配置。
func X设置配置Map(m map[string]interface{}) error {
	return defaultLogger.X设置配置Map(m)
}

// SetPath 设置文件日志的目录路径。
func X设置文件路径(文件路径 string) error {
	return defaultLogger.X设置文件路径(文件路径)
}

// GetPath 返回用于文件日志记录的日志目录路径。
// 如果未设置目录路径，则返回空字符串。
func X取文件路径() string {
	return defaultLogger.X取文件路径()
}

// SetFile 设置文件日志的文件名`pattern`。
// 在`pattern`中可以使用日期时间模式，例如：access-{Ymd}.log。
// 默认的文件名模式是：Y-m-d.log，例如：2018-01-01.log
// 这段Go语言代码注释翻译成中文为：
// 设置文件日志的文件名为 `pattern`。
// 可以在 `pattern` 中使用日期时间格式化字符串，例如：access-{Ymd}.log（表示按年月日生成不同文件）。
// 默认的文件名格式是：Y-m-d.log，例如：2018-01-01.log
func X设置文件名格式(文件名格式 string) {
	defaultLogger.X设置文件名格式(文件名格式)
}

// SetLevel 设置默认的日志级别。
func X设置级别(级别 int) {
	defaultLogger.X设置级别(级别)
}

// GetLevel 返回默认的日志级别值。
func X取级别() int {
	return defaultLogger.X取级别()
}

// SetWriter 设置自定义的日志 `writer` 用于日志记录。
// `writer` 对象应实现 io.Writer 接口。
// 开发者可以使用自定义的日志 `writer` 将日志输出重定向到其他服务，
// 例如：kafka、mysql、mongodb 等等。
func X设置Writer(writer io.Writer) {
	defaultLogger.X设置Writer(writer)
}

// GetWriter 返回一个自定义的writer对象，该对象实现了io.Writer接口。
// 如果未设置自定义writer，则返回nil。
func X取Writer() io.Writer {
	return defaultLogger.X取Writer()
}

// SetDebug 用于启用/禁用默认 defaultLogger 的调试级别。
// 调试级别默认是启用状态。
func X设置debug(开启 bool) {
	defaultLogger.X设置debug(开启)
}

// SetAsync 用于启用/禁用默认默认Logger的异步日志输出功能。
func X设置异步输出(开启 bool) {
	defaultLogger.X设置异步输出(开启)
}

// SetStdoutPrint 设置是否将日志内容输出到标准输出（stdout），默认情况下为true。
func X设置是否同时输出到终端(开启 bool) {
	defaultLogger.X设置是否同时输出到终端(开启)
}

// SetHeaderPrint 设置是否输出日志内容的头部，默认为true。
func X设置是否输出头信息(开启 bool) {
	defaultLogger.X设置是否输出头信息(开启)
}

// SetPrefix 设置每个日志内容的前缀字符串。
// 前缀是头部的一部分，这意味着如果关闭了头部输出，则不会输出任何前缀。
func X设置前缀(前缀 string) {
	defaultLogger.X设置前缀(前缀)
}

// SetFlags 设置日志输出功能的额外标志。
func X设置额外标识(标识 int) {
	defaultLogger.X设置额外标识(标识)
}

// GetFlags 返回默认日志器的标志。
func X取标识() int {
	return defaultLogger.X取标识()
}

// SetCtxKeys 为 defaultLogger 设置上下文键。这些键用于从上下文中检索值并将其打印到日志内容中。
//
// 注意，多次调用此函数会覆盖之前设置的上下文键。
func X设置上下文名称(名称 ...interface{}) {
	defaultLogger.X设置上下文名称(名称...)
}

// GetCtxKeys 获取并返回用于日志记录的上下文键。
func X取上下文名称() []interface{} {
	return defaultLogger.X取上下文名称()
}

// PrintStack 打印调用栈，
// 可选参数 `skip` 指定了从终点开始需要跳过的堆栈偏移量。
func X输出堆栈信息(上下文 context.Context, 偏移量 ...int) {
	defaultLogger.PrintStack(上下文, 偏移量...)
}

// GetStack 返回调用堆栈的内容，
// 可选参数 `skip` 指定了从终点开始跳过的堆栈偏移量。
func X取堆栈信息(偏移量 ...int) string {
	return defaultLogger.X取堆栈信息(偏移量...)
}

// SetStack 启用/禁用失败日志输出中的堆栈跟踪功能。
func X设置堆栈跟踪(开启 bool) {
	defaultLogger.X设置堆栈跟踪(开启)
}

// SetLevelStr 通过级别字符串设置日志记录级别。
func X设置文本级别(级别 string) error {
	return defaultLogger.X设置文本级别(级别)
}

// SetLevelPrefix 为指定等级设置前缀字符串。
func X设置级别前缀(级别 int, 前缀 string) {
	defaultLogger.X设置级别前缀(级别, 前缀)
}

// SetLevelPrefixes 为默认日志器设置等级到前缀字符串的映射。
func X设置级别前缀Map(前缀Map map[int]string) {
	defaultLogger.X设置级别前缀Map(前缀Map)
}

// GetLevelPrefix 返回指定级别的前缀字符串。
func X取级别前缀(级别 int) string {
	return defaultLogger.X取级别前缀(级别)
}

// SetHandlers 为默认的 defaultLogger 设置日志处理器。
func X设置中间件(处理函数 ...Handler) {
	defaultLogger.X设置中间件(处理函数...)
}

// SetWriterColorEnable 设置文件日志记录使用颜色
func X设置文件是否输出颜色(开启 bool) {
	defaultLogger.X设置文件是否输出颜色(开启)
}
