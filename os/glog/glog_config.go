// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package glog

import (
	"context"
	"io"
)

// SetConfig 设置默认日志器的配置。
func SetConfig(config Config) error {
	return defaultLogger.SetConfig(config)
}

// SetConfigWithMap 通过map为默认日志器设置配置。
func SetConfigWithMap(m map[string]interface{}) error {
	return defaultLogger.SetConfigWithMap(m)
}

// SetPath 设置文件日志的目录路径。
func SetPath(path string) error {
	return defaultLogger.SetPath(path)
}

// GetPath 返回用于文件日志记录的日志目录路径。
// 如果未设置目录路径，则返回空字符串。
func GetPath() string {
	return defaultLogger.GetPath()
}

// SetFile 设置文件日志的文件名`pattern`。
// 在`pattern`中可以使用日期时间模式，例如：access-{Ymd}.log。
// 默认的文件名模式是：Y-m-d.log，例如：2018-01-01.log
// 这段Go语言代码注释翻译成中文为：
// 设置文件日志的文件名为 `pattern`。
// 可以在 `pattern` 中使用日期时间格式化字符串，例如：access-{Ymd}.log（表示按年月日生成不同文件）。
// 默认的文件名格式是：Y-m-d.log，例如：2018-01-01.log
func SetFile(pattern string) {
	defaultLogger.SetFile(pattern)
}

// SetLevel 设置默认的日志级别。
func SetLevel(level int) {
	defaultLogger.SetLevel(level)
}

// GetLevel 返回默认的日志级别值。
func GetLevel() int {
	return defaultLogger.GetLevel()
}

// SetWriter 设置自定义的日志 `writer` 用于日志记录。
// `writer` 对象应实现 io.Writer 接口。
// 开发者可以使用自定义的日志 `writer` 将日志输出重定向到其他服务，
// 例如：kafka、mysql、mongodb 等等。
func SetWriter(writer io.Writer) {
	defaultLogger.SetWriter(writer)
}

// GetWriter 返回一个自定义的writer对象，该对象实现了io.Writer接口。
// 如果未设置自定义writer，则返回nil。
func GetWriter() io.Writer {
	return defaultLogger.GetWriter()
}

// SetDebug 用于启用/禁用默认 defaultLogger 的调试级别。
// 调试级别默认是启用状态。
func SetDebug(debug bool) {
	defaultLogger.SetDebug(debug)
}

// SetAsync 用于启用/禁用默认默认Logger的异步日志输出功能。
func SetAsync(enabled bool) {
	defaultLogger.SetAsync(enabled)
}

// SetStdoutPrint 设置是否将日志内容输出到标准输出（stdout），默认情况下为true。
func SetStdoutPrint(enabled bool) {
	defaultLogger.SetStdoutPrint(enabled)
}

// SetHeaderPrint 设置是否输出日志内容的头部，默认为true。
func SetHeaderPrint(enabled bool) {
	defaultLogger.SetHeaderPrint(enabled)
}

// SetPrefix 设置每个日志内容的前缀字符串。
// 前缀是头部的一部分，这意味着如果关闭了头部输出，则不会输出任何前缀。
func SetPrefix(prefix string) {
	defaultLogger.SetPrefix(prefix)
}

// SetFlags 设置日志输出功能的额外标志。
func SetFlags(flags int) {
	defaultLogger.SetFlags(flags)
}

// GetFlags 返回默认日志器的标志。
func GetFlags() int {
	return defaultLogger.GetFlags()
}

// SetCtxKeys 为 defaultLogger 设置上下文键。这些键用于从上下文中检索值并将其打印到日志内容中。
//
// 注意，多次调用此函数会覆盖之前设置的上下文键。
func SetCtxKeys(keys ...interface{}) {
	defaultLogger.SetCtxKeys(keys...)
}

// GetCtxKeys 获取并返回用于日志记录的上下文键。
func GetCtxKeys() []interface{} {
	return defaultLogger.GetCtxKeys()
}

// PrintStack 打印调用栈，
// 可选参数 `skip` 指定了从终点开始需要跳过的堆栈偏移量。
func PrintStack(ctx context.Context, skip ...int) {
	defaultLogger.PrintStack(ctx, skip...)
}

// GetStack 返回调用堆栈的内容，
// 可选参数 `skip` 指定了从终点开始跳过的堆栈偏移量。
func GetStack(skip ...int) string {
	return defaultLogger.GetStack(skip...)
}

// SetStack 启用/禁用失败日志输出中的堆栈跟踪功能。
func SetStack(enabled bool) {
	defaultLogger.SetStack(enabled)
}

// SetLevelStr 通过级别字符串设置日志记录级别。
func SetLevelStr(levelStr string) error {
	return defaultLogger.SetLevelStr(levelStr)
}

// SetLevelPrefix 为指定等级设置前缀字符串。
func SetLevelPrefix(level int, prefix string) {
	defaultLogger.SetLevelPrefix(level, prefix)
}

// SetLevelPrefixes 为默认日志器设置等级到前缀字符串的映射。
func SetLevelPrefixes(prefixes map[int]string) {
	defaultLogger.SetLevelPrefixes(prefixes)
}

// GetLevelPrefix 返回指定级别的前缀字符串。
func GetLevelPrefix(level int) string {
	return defaultLogger.GetLevelPrefix(level)
}

// SetHandlers 为默认的 defaultLogger 设置日志处理器。
func SetHandlers(handlers ...Handler) {
	defaultLogger.SetHandlers(handlers...)
}

// SetWriterColorEnable 设置文件日志记录使用颜色
func SetWriterColorEnable(enabled bool) {
	defaultLogger.SetWriterColorEnable(enabled)
}
