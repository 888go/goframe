// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package g

import (
	"context"
	"io"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/util/gutil"
)

// Go 创建一个新的异步 goroutine 函数，并指定了恢复函数。
//
// 参数 `recoverFunc` 在 `goroutineFunc` 执行过程中发生任何 panic 时被调用。
// 如果 `recoverFunc` 为 nil，它将忽略 `goroutineFunc` 中的 panic，且不会向父 goroutine 抛出 panic。
//
// 但是要注意，如果 `recoverFunc` 自身也抛出 panic，这个 panic 将会被抛给父 goroutine。 md5:3820395064a9e843
func Go(
	ctx context.Context,
	goroutineFunc func(ctx context.Context),
	recoverFunc func(ctx context.Context, exception error),
) {
	gutil.Go(ctx, goroutineFunc, recoverFunc)
}

// NewVar返回一个gvar.Var。 md5:384cbf5dc86048e6
func NewVar(i interface{}, safe ...bool) *Var {
	return gvar.New(i, safe...)
}

// Wait 是对 ghttp.Wait 的别名，它会阻塞直到所有的网络服务器停止运行。在多服务器情况下，这通常会被使用。 md5:efbd28e068404766
func Wait() {
	ghttp.Wait()
}

// Listen 是 gproc.Listen 的别名，用于处理接收到的信号，并自动调用已注册的信号处理函数。
// 它会阻塞直到接收到关闭信号，并且所有已注册的关闭处理函数执行完毕。 md5:566b27b7da3e69b8
func Listen() {
	gproc.Listen()
}

// Dump 将变量以更易读的格式输出到标准输出。 md5:e19cf0b857ffd1c6
func Dump(values ...interface{}) {
	gutil.Dump(values...)
}

// DumpTo 将变量 `values` 作为字符串写入到 `writer` 中，提供更易人工阅读的格式. md5:68fd8fc9ea0dfc4b
func DumpTo(writer io.Writer, value interface{}, option gutil.DumpOption) {
	gutil.DumpTo(writer, value, option)
}

// DumpWithType 类似于 Dump，但带有类型信息。同时参阅 Dump。 md5:faabab79589d38a3
func DumpWithType(values ...interface{}) {
	gutil.DumpWithType(values...)
}

// DumpWithOption 函数将变量 `values` 以更易于人工阅读的字符串形式返回。 md5:99fec3f0f209dcf7
func DumpWithOption(value interface{}, option gutil.DumpOption) {
	gutil.DumpWithOption(value, option)
}

// DumpJson 将 JSON 内容以美化的方式输出到标准输出。 md5:9f4c95e099395360
func DumpJson(value any) {
	gutil.DumpJson(value)
}

// Throw抛出一个异常，可以被TryCatch函数捕获。 md5:ade0a15f9238635b
func Throw(exception interface{}) {
	gutil.Throw(exception)
}

// Try 使用内部的 panic...recover 实现 try...catch 逻辑。如果发生任何异常，它会返回错误；否则返回 nil。 md5:7c85aa857bb16fca
func Try(ctx context.Context, try func(ctx context.Context)) (err error) {
	return gutil.Try(ctx, try)
}

// TryCatch 使用内部的 panic...recover 机制来实现 try...catch... 的逻辑。
// 如果发生任何异常，它会自动调用函数 `catch` 并将异常作为错误传递。
//
// 但请注意，如果函数 `catch` 本身也引发了 panic，当前的 goroutine 会触发 panic。 md5:c9fae3297a82421f
func TryCatch(ctx context.Context, try func(ctx context.Context), catch func(ctx context.Context, exception error)) {
	gutil.TryCatch(ctx, try, catch)
}

// IsNil 检查给定的 `value` 是否为 nil。
// 参数 `traceSource` 用于在 `value` 为指向指针的指针类型时，追踪到源变量。当 `traceSource` 为真且源为 nil 时，它会返回 nil。
// 请注意，该功能可能使用反射，这可能稍微影响性能。 md5:1a86ab3bedd8914d
func IsNil(value interface{}, traceSource ...bool) bool {
	return empty.IsNil(value, traceSource...)
}

// IsEmpty 检查给定的 `value` 是否为空。
// 如果 `value` 等于：0，nil，false，""，切片、映射或通道的长度为0，则返回 true。
// 否则，返回 true。
//
// 参数 `traceSource` 用于追踪，如果给定的 `value` 是指向指针的指针类型，它会追踪到源变量。当 `traceSource` 为 true 时，如果源为空，将返回 true。
// 注意，它可能使用反射特性，这可能会影响性能。 md5:7262defa92ce37cb
func IsEmpty(value interface{}, traceSource ...bool) bool {
	return empty.IsEmpty(value, traceSource...)
}

// RequestFromCtx 从上下文中检索并返回Request对象。 md5:c247eac3d031fb2b
func RequestFromCtx(ctx context.Context) *ghttp.Request {
	return ghttp.RequestFromCtx(ctx)
}
