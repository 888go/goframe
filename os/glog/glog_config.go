// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package glog

import (
	"context"
	"io"
)

// SetConfig 为默认日志器设置配置。. md5:5f388c6c2d28724c
func SetConfig(config Config) error {
	return defaultLogger.SetConfig(config)
}

// 使用映射为defaultLogger设置配置。. md5:195834b40cfce621
func SetConfigWithMap(m map[string]interface{}) error {
	return defaultLogger.SetConfigWithMap(m)
}

// SetPath 设置文件日志的目录路径。. md5:817e6d2802241584
func SetPath(path string) error {
	return defaultLogger.SetPath(path)
}

// GetPath 返回用于文件日志记录的目录路径。
// 如果未设置目录路径，它将返回空字符串。
// md5:f69da996992ffd9e
func GetPath() string {
	return defaultLogger.GetPath()
}

// SetFile 设置文件日志的文件名模式为 `pattern`。
// 在 `pattern` 中可以使用日期时间模式，例如：access-YYYYMMDD.log。
// 默认的文件名模式为：Y-m-d.log，例如：2018-01-01.log
// md5:03b3a973ce783b24
func SetFile(pattern string) {
	defaultLogger.SetFile(pattern)
}

// SetLevel 设置默认的日志级别。. md5:ab428766cae30d65
func SetLevel(level int) {
	defaultLogger.SetLevel(level)
}

// GetLevel 返回默认的日志级别值。. md5:5ae9c76de12ac5c9
func GetLevel() int {
	return defaultLogger.GetLevel()
}

// SetWriter 设置自定义的日志记录器`writer`。
// `writer`对象应实现io.Writer接口。
// 开发者可以使用自定义的`writer`将日志输出重定向到其他服务，如：kafka、mysql、mongodb等。
// md5:50799f5f4e6280ea
func SetWriter(writer io.Writer) {
	defaultLogger.SetWriter(writer)
}

// GetWriter 返回自定义的 writer 对象，它实现了 io.Writer 接口。
// 如果没有设置自定义 writer，它将返回 nil。
// md5:37b3d820b5547c91
func GetWriter() io.Writer {
	return defaultLogger.GetWriter()
}

// SetDebug 启用/禁用默认 defaultLogger 的调试级别。默认情况下，调试级别已启用。
// md5:335740e0731656b1
func SetDebug(debug bool) {
	defaultLogger.SetDebug(debug)
}

// SetAsync 用于启用或禁用默认默认日志记录器的异步日志输出功能。. md5:2659994b118a6211
func SetAsync(enabled bool) {
	defaultLogger.SetAsync(enabled)
}

// SetStdoutPrint 设置是否将日志内容输出到stdout，默认为true。. md5:f68d5bea1f6372d1
func SetStdoutPrint(enabled bool) {
	defaultLogger.SetStdoutPrint(enabled)
}

// SetHeaderPrint 设置日志输出的头部是否打印，默认为 true。. md5:3e71cb67564384cc
func SetHeaderPrint(enabled bool) {
	defaultLogger.SetHeaderPrint(enabled)
}

// SetPrefix 设置日志内容的前缀字符串。
// 前缀是日志头的一部分，如果关闭了头部输出，就不会显示前缀。
// md5:31d8e3c101c1eea6
func SetPrefix(prefix string) {
	defaultLogger.SetPrefix(prefix)
}

// SetFlags 为日志输出功能设置额外的标志。. md5:40253d4ed662de77
func SetFlags(flags int) {
	defaultLogger.SetFlags(flags)
}

// GetFlags 返回默认日志器defaultLogger的标志。. md5:ca4e9bc9260b114a
func GetFlags() int {
	return defaultLogger.GetFlags()
}

// SetCtxKeys 为 defaultLogger 设置上下文键。这些键用于从上下文中检索值，并将它们打印到日志内容中。
// 
// 注意，多次调用此函数会覆盖之前设置的上下文键。
// md5:8a4710a99dd19d4d
func SetCtxKeys(keys ...interface{}) {
	defaultLogger.SetCtxKeys(keys...)
}

// GetCtxKeys检索并返回日志记录的上下文键。. md5:1e780871ada8c59c
func GetCtxKeys() []interface{} {
	return defaultLogger.GetCtxKeys()
}

// PrintStack 打印调用堆栈，
// 可选参数 `skip` 用于指定从堆栈终点开始忽略的偏移量。
// md5:ef6cd40820765783
func PrintStack(ctx context.Context, skip ...int) {
	defaultLogger.PrintStack(ctx, skip...)
}

// GetStack 返回调用者栈的内容，
// 可选参数 `skip` 指定从终点开始要跳过的栈偏移量。
// md5:13592be3061e779d
func GetStack(skip ...int) string {
	return defaultLogger.GetStack(skip...)
}

// SetStack 启用/禁用失败日志输出中的堆栈功能。. md5:3c80a664fff650de
func SetStack(enabled bool) {
	defaultLogger.SetStack(enabled)
}

// SetLevelStr 通过级别字符串设置日志级别。. md5:53cbbdf23584340e
func SetLevelStr(levelStr string) error {
	return defaultLogger.SetLevelStr(levelStr)
}

// SetLevelPrefix 为指定的日志级别设置前缀字符串。. md5:a2b7a43af150bcb7
func SetLevelPrefix(level int, prefix string) {
	defaultLogger.SetLevelPrefix(level, prefix)
}

// SetLevelPrefixes 设置默认logger的级别到前缀字符串的映射。. md5:5d55474512859352
func SetLevelPrefixes(prefixes map[int]string) {
	defaultLogger.SetLevelPrefixes(prefixes)
}

// GetLevelPrefix 返回指定级别的前缀字符串。. md5:339b86b4f84d6049
func GetLevelPrefix(level int) string {
	return defaultLogger.GetLevelPrefix(level)
}

// SetHandlers 设置默认日志记录器defaultLogger的处理程序。. md5:637a5dc8ccc73e8f
func SetHandlers(handlers ...Handler) {
	defaultLogger.SetHandlers(handlers...)
}

// SetWriterColorEnable 设置文件日志记录使用颜色. md5:3ff25da59932914b
func SetWriterColorEnable(enabled bool) {
	defaultLogger.SetWriterColorEnable(enabled)
}
