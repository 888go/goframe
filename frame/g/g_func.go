// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package g
import (
	"context"
	"io"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/os/gproc"
	"github.com/888go/goframe/util/gutil"
	)
// Go 创建了一个具有指定恢复函数的新的异步 goroutine 函数。
//
// 参数 `recoverFunc` 在执行 `goroutineFunc` 期间发生任何 panic 时被调用。
// 如果 `recoverFunc` 被赋予 nil，它将忽略来自 `goroutineFunc` 的 panic，并且不会向父级 goroutine 抛出 panic。
//
// 但是请注意，如果 `recoverFunc` 也抛出了 panic，这样的 panic 将会被抛给父级 goroutine。
func Go(
	ctx context.Context,
	goroutineFunc func(ctx context.Context),
	recoverFunc func(ctx context.Context, exception error),
) {
	gutil.Go(ctx, goroutineFunc, recoverFunc)
}

// NewVar 返回一个 gvar.Var。
func NewVar(i interface{}, safe ...bool) *Var {
	return gvar.New(i, safe...)
}

// Wait 是 ghttp.Wait 的别名，它会阻塞直到所有 web 服务器关闭。
// 在多服务器场景中，它经常被使用。
func Wait() {
	ghttp.Wait()
}

// Listen 是 gproc.Listen 的别名，用于处理接收到的信号并自动调用已注册的信号处理器函数。
// 它会阻塞直到接收到关闭信号且所有已注册的关闭处理器执行完毕。
func Listen() {
	gproc.Listen()
}

// Dump 将一个变量以更易于人工阅读的方式输出到标准输出（stdout）
func Dump(values ...interface{}) {
	gutil.Dump(values...)
}

// DumpTo 将变量 `values` 转换为字符串并写入到 `writer` 中，以更易于人工阅读的方式
func DumpTo(writer io.Writer, value interface{}, option gutil.DumpOption) {
	gutil.DumpTo(writer, value, option)
}

// DumpWithType 的行为类似于 Dump，但会包含类型信息。
// 也可参考 Dump。
func DumpWithType(values ...interface{}) {
	gutil.DumpWithType(values...)
}

// DumpWithOption 使用自定义选项返回变量 `values`，将其格式化为更易读的字符串形式。
func DumpWithOption(value interface{}, option gutil.DumpOption) {
	gutil.DumpWithOption(value, option)
}

// DumpJson 将 JSON 内容格式化输出到标准输出（stdout）。
func DumpJson(jsonContent string) {
	gutil.DumpJson(jsonContent)
}

// Throw 抛出一个异常，该异常可以被 TryCatch 函数捕获。
func Throw(exception interface{}) {
	gutil.Throw(exception)
}

// Try 使用内部 panic...recover 实现 try...逻辑。
// 如果发生任何异常，它将返回错误，否则返回 nil。
func Try(ctx context.Context, try func(ctx context.Context)) (err error) {
	return gutil.Try(ctx, try)
}

// TryCatch 通过内部 panic...recover 实现了类似 try...catch... 的错误处理逻辑。
// 当出现任何异常时，它会自动调用函数 `catch` 并将异常作为 error 参数传递。
//
// 但是请注意，如果函数 `catch` 本身也抛出了 panic，则当前的 goroutine 将同样触发 panic。
func TryCatch(ctx context.Context, try func(ctx context.Context), catch func(ctx context.Context, exception error)) {
	gutil.TryCatch(ctx, try, catch)
}

// IsNil 检查给定的 `value` 是否为 nil。
// 参数 `traceSource` 用于在 `value` 类型为指向指针的指针时，追踪到源变量。如果当 `traceSource` 为真时源变量为 nil，则返回 nil。
// 注意，此函数可能使用 reflect 特性，这会对性能产生轻微影响。
func IsNil(value interface{}, traceSource ...bool) bool {
	return empty.IsNil(value, traceSource...)
}

// IsEmpty 检查给定的 `value` 是否为空。
// 当 `value` 为以下情形时返回 true：0、nil、false、""、slice/map/chan 的长度为 0。
// 否则返回 false。
// 参数 `traceSource` 用于在 `value` 类型为指针且指向另一个指针时，追踪到源变量。
// 当 `traceSource` 为 true 时，如果源变量为空，则返回 true。
// 注意，这可能使用 reflect 特性，会对性能造成轻微影响。
func IsEmpty(value interface{}, traceSource ...bool) bool {
	return empty.IsEmpty(value, traceSource...)
}

// RequestFromCtx 从 context 中检索并返回 Request 对象。
func RequestFromCtx(ctx context.Context) *ghttp.Request {
	return ghttp.RequestFromCtx(ctx)
}
