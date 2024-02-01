// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gutil
import (
	"context"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	)
// Throw 抛出一个异常，该异常可以被 TryCatch 或 recover 捕获。
func Throw(exception interface{}) {
	panic(exception)
}

// Try 使用内部 panic...recover 实现 try...逻辑。
// 如果发生任何异常，它将返回错误，否则返回 nil。
func Try(ctx context.Context, try func(ctx context.Context)) (err error) {
	if try == nil {
		return
	}
	defer func() {
		if exception := recover(); exception != nil {
			if v, ok := exception.(error); ok && gerror.HasStack(v) {
				err = v
			} else {
				err = gerror.NewCodef(gcode.CodeInternalPanic, "%+v", exception)
			}
		}
	}()
	try(ctx)
	return
}

// TryCatch 实现了类似 `try...catch...` 的逻辑，通过内部使用 `panic...recover`。
// 如果发生任何异常，它会自动调用函数 `catch` 并将异常作为 error 传递给 `catch` 函数。
// 若传入的 `catch` 为 nil，则忽略来自 `try` 的 panic，并且不会向父级 goroutine 抛出 panic。
//
// 但是请注意，如果函数 `catch` 本身也抛出了 panic，则当前 goroutine 将会 panic。
func TryCatch(ctx context.Context, try func(ctx context.Context), catch func(ctx context.Context, exception error)) {
	if try == nil {
		return
	}
	if exception := Try(ctx, try); exception != nil && catch != nil {
		catch(ctx, exception)
	}
}
