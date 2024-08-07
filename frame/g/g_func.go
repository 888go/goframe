// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package g

import (
	"context"
	"io"

	gvar "github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/internal/empty"
	ghttp "github.com/888go/goframe/net/ghttp"
	gproc "github.com/888go/goframe/os/gproc"
	gutil "github.com/888go/goframe/util/gutil"
)

// Go 创建一个新的异步 goroutine 函数，并指定了恢复函数。
//
// 参数 `recoverFunc` 在 `goroutineFunc` 执行过程中发生任何 panic 时被调用。
// 如果 `recoverFunc` 为 nil，它将忽略 `goroutineFunc` 中的 panic，且不会向父 goroutine 抛出 panic。
//
// 但是要注意，如果 `recoverFunc` 自身也抛出 panic，这个 panic 将会被抛给父 goroutine。
// md5:3820395064a9e843
func Go(
	ctx context.Context,
	goroutineFunc func(ctx context.Context),
	recoverFunc func(ctx context.Context, exception error),
) {
	gutil.Go(ctx, goroutineFunc, recoverFunc)
}

// X泛型类返回一个gvar.Var。 md5:384cbf5dc86048e6
func X泛型类(值 interface{}, 并发安全 ...bool) *Var {
	return gvar.X创建(值, 并发安全...)
}

// Http类等待所有服务完成 是对 ghttp.Http类等待所有服务完成 的别名，它会阻塞直到所有的网络服务器停止运行。在多服务器情况下，这通常会被使用。
// md5:efbd28e068404766
func Http类等待所有服务完成() {
	ghttp.X等待所有服务完成()
}

// Listen 是 gproc.Listen 的别名，用于处理接收到的信号，并自动调用已注册的信号处理函数。
// 它会阻塞直到接收到关闭信号，并且所有已注册的关闭处理函数执行完毕。
// md5:566b27b7da3e69b8
func Listen() {
	gproc.Listen()
}

// X调试输出 将变量以更易读的格式输出到标准输出。 md5:e19cf0b857ffd1c6
func X调试输出(值s ...interface{}) {
	gutil.X调试输出(值s...)
}

// X调试输出到Writer 将变量 `values` 作为字符串写入到 `writer` 中，提供更易人工阅读的格式. md5:68fd8fc9ea0dfc4b
func X调试输出到Writer(writer io.Writer, 值 interface{}, 选项 gutil.DumpOption) {
	gutil.X调试输出到Writer(writer, 值, 选项)
}

// X调试输出并带类型 类似于 Dump，但带有类型信息。同时参阅 Dump。
// md5:faabab79589d38a3
func X调试输出并带类型(值s ...interface{}) {
	gutil.X调试输出并带类型(值s...)
}

// X调试输出并带选项 函数将变量 `values` 以更易于人工阅读的字符串形式返回。 md5:99fec3f0f209dcf7
func X调试输出并带选项(值s interface{}, 选项 gutil.DumpOption) {
	gutil.X调试输出并带选项(值s, 选项)
}

// X调试输出json 将 JSON 内容以美化的方式输出到标准输出。 md5:9f4c95e099395360
func X调试输出json(value any) {
	gutil.X调试输出json(value)
}

// X异常输出抛出一个异常，可以被TryCatch函数捕获。 md5:ade0a15f9238635b
func X异常输出(消息 interface{}) {
	gutil.X异常输出(消息)
}

// X异常捕捉 使用内部的 panic...recover 实现 try...catch 逻辑。如果发生任何异常，它会返回错误；否则返回 nil。
// md5:7c85aa857bb16fca
func X异常捕捉(上下文 context.Context, 处理函数 func(上下文 context.Context)) (错误 error) {
	return gutil.X异常捕捉(上下文, 处理函数)
}

// X异常捕捉并带异常处理 使用内部的 panic...recover 机制来实现 try...catch... 的逻辑。
// 如果发生任何异常，它会自动调用函数 `catch` 并将异常作为错误传递。
//
// 但请注意，如果函数 `catch` 本身也引发了 panic，当前的 goroutine 会触发 panic。
// md5:c9fae3297a82421f
func X异常捕捉并带异常处理(上下文 context.Context, 处理函数 func(上下文 context.Context), 异常处理函数 func(上下文 context.Context, 错误 error)) {
	gutil.X异常捕捉并带异常处理(上下文, 处理函数, 异常处理函数)
}

// X是否为Nil 检查给定的 `value` 是否为 nil。
// 参数 `traceSource` 用于在 `value` 为指向指针的指针类型时，追踪到源变量。当 `traceSource` 为真且源为 nil 时，它会返回 nil。
// 请注意，该功能可能使用反射，这可能稍微影响性能。
// md5:1a86ab3bedd8914d
func X是否为Nil(value interface{}, traceSource ...bool) bool {
	return empty.IsNil(value, traceSource...)
}

// X是否为空 检查给定的 `value` 是否为空。
// 如果 `value` 等于：0，nil，false，""，切片、映射或通道的长度为0，则返回 true。
// 否则，返回 true。
// 
// 参数 `traceSource` 用于追踪，如果给定的 `value` 是指向指针的指针类型，它会追踪到源变量。当 `traceSource` 为 true 时，如果源为空，将返回 true。
// 注意，它可能使用反射特性，这可能会影响性能。
// md5:7262defa92ce37cb
func X是否为空(值 interface{}, 追踪到源变量 ...bool) bool {
	return empty.IsEmpty(值, 追踪到源变量...)
}

// Http类上下文取请求对象 从上下文中检索并返回Request对象。 md5:c247eac3d031fb2b
func Http类上下文取请求对象(上下文 context.Context) *ghttp.Request {
	return ghttp.X从上下文取请求对象(上下文)
}
