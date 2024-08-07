// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 日志类

import (
	"context"
	"io"
)

// X设置配置项 为默认日志器设置配置。 md5:5f388c6c2d28724c
func X设置配置项(配置项 Config) error {
	return defaultLogger.X设置配置项(配置项)
}

// 使用映射为defaultLogger设置配置。 md5:195834b40cfce621
func X设置配置Map(m map[string]interface{}) error {
	return defaultLogger.X设置配置Map(m)
}

// X设置文件路径 设置文件日志的目录路径。 md5:817e6d2802241584
func X设置文件路径(文件路径 string) error {
	return defaultLogger.X设置文件路径(文件路径)
}

// X取文件路径 返回用于文件日志记录的目录路径。
// 如果未设置目录路径，它将返回空字符串。
// md5:f69da996992ffd9e
func X取文件路径() string {
	return defaultLogger.X取文件路径()
}

// X设置文件名格式 设置文件日志的文件名模式为 `pattern`。
// 在 `pattern` 中可以使用日期时间模式，例如：access-YYYYMMDD.log。
// 默认的文件名模式为：Y-m-d.log，例如：2018-01-01.log
// md5:03b3a973ce783b24
func X设置文件名格式(文件名格式 string) {
	defaultLogger.X设置文件名格式(文件名格式)
}

// X设置级别 设置默认的日志级别。 md5:ab428766cae30d65
func X设置级别(级别 int) {
	defaultLogger.X设置级别(级别)
}

// X取级别 返回默认的日志级别值。 md5:5ae9c76de12ac5c9
func X取级别() int {
	return defaultLogger.X取级别()
}

// X设置Writer 设置自定义的日志记录器`writer`。
// `writer`对象应实现io.Writer接口。
// 开发者可以使用自定义的`writer`将日志输出重定向到其他服务，如：kafka、mysql、mongodb等。
// md5:50799f5f4e6280ea
func X设置Writer(writer io.Writer) {
	defaultLogger.X设置Writer(writer)
}

// X取Writer 返回自定义的 writer 对象，它实现了 io.Writer 接口。
// 如果没有设置自定义 writer，它将返回 nil。
// md5:37b3d820b5547c91
func X取Writer() io.Writer {
	return defaultLogger.X取Writer()
}

// X设置debug 启用/禁用默认 defaultLogger 的调试级别。默认情况下，调试级别已启用。
// md5:335740e0731656b1
func X设置debug(开启 bool) {
	defaultLogger.X设置debug(开启)
}

// X设置异步输出 用于启用或禁用默认默认日志记录器的异步日志输出功能。 md5:2659994b118a6211
func X设置异步输出(开启 bool) {
	defaultLogger.X设置异步输出(开启)
}

// X设置是否同时输出到终端 设置是否将日志内容输出到stdout，默认为true。 md5:f68d5bea1f6372d1
func X设置是否同时输出到终端(开启 bool) {
	defaultLogger.X设置是否同时输出到终端(开启)
}

// X设置是否输出头信息 设置日志输出的头部是否打印，默认为 true。 md5:3e71cb67564384cc
func X设置是否输出头信息(开启 bool) {
	defaultLogger.X设置是否输出头信息(开启)
}

// X设置前缀 设置日志内容的前缀字符串。
// 前缀是日志头的一部分，如果关闭了头部输出，就不会显示前缀。
// md5:31d8e3c101c1eea6
func X设置前缀(前缀 string) {
	defaultLogger.X设置前缀(前缀)
}

// X设置额外标识 为日志输出功能设置额外的标志。 md5:40253d4ed662de77
func X设置额外标识(标识 int) {
	defaultLogger.X设置额外标识(标识)
}

// X取标识 返回默认日志器defaultLogger的标志。 md5:ca4e9bc9260b114a
func X取标识() int {
	return defaultLogger.X取标识()
}

// X设置上下文名称 为 defaultLogger 设置上下文键。这些键用于从上下文中检索值，并将它们打印到日志内容中。
// 
// 注意，多次调用此函数会覆盖之前设置的上下文键。
// md5:8a4710a99dd19d4d
func X设置上下文名称(名称 ...interface{}) {
	defaultLogger.X设置上下文名称(名称...)
}

// X取上下文名称检索并返回日志记录的上下文键。 md5:1e780871ada8c59c
func X取上下文名称() []interface{} {
	return defaultLogger.X取上下文名称()
}

// X输出堆栈信息 打印调用堆栈，
// 可选参数 `skip` 用于指定从堆栈终点开始忽略的偏移量。
// md5:ef6cd40820765783
func X输出堆栈信息(上下文 context.Context, 偏移量 ...int) {
	defaultLogger.PrintStack(上下文, 偏移量...)
}

// X取堆栈信息 返回调用者栈的内容，
// 可选参数 `skip` 指定从终点开始要跳过的栈偏移量。
// md5:13592be3061e779d
func X取堆栈信息(偏移量 ...int) string {
	return defaultLogger.X取堆栈信息(偏移量...)
}

// X设置堆栈跟踪 启用/禁用失败日志输出中的堆栈功能。 md5:3c80a664fff650de
func X设置堆栈跟踪(开启 bool) {
	defaultLogger.X设置堆栈跟踪(开启)
}

// X设置文本级别 通过级别字符串设置日志级别。 md5:53cbbdf23584340e
func X设置文本级别(级别 string) error {
	return defaultLogger.X设置文本级别(级别)
}

// X设置级别前缀 为指定的日志级别设置前缀字符串。 md5:a2b7a43af150bcb7
func X设置级别前缀(级别 int, 前缀 string) {
	defaultLogger.X设置级别前缀(级别, 前缀)
}

// X设置级别前缀Map 设置默认logger的级别到前缀字符串的映射。 md5:5d55474512859352
func X设置级别前缀Map(前缀Map map[int]string) {
	defaultLogger.X设置级别前缀Map(前缀Map)
}

// X取级别前缀 返回指定级别的前缀字符串。 md5:339b86b4f84d6049
func X取级别前缀(级别 int) string {
	return defaultLogger.X取级别前缀(级别)
}

// X设置中间件 设置默认日志记录器defaultLogger的处理程序。 md5:637a5dc8ccc73e8f
func X设置中间件(处理函数 ...Handler) {
	defaultLogger.X设置中间件(处理函数...)
}

// X设置文件是否输出颜色 设置文件日志记录使用颜色. md5:3ff25da59932914b
func X设置文件是否输出颜色(开启 bool) {
	defaultLogger.X设置文件是否输出颜色(开启)
}
