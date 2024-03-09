// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类

import (
	"context"
	
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// Throw 抛出一个异常，该异常可以被 TryCatch 或 recover 捕获。
func X异常输出(消息 interface{}) {
	panic(消息)
}

// Try 使用内部 panic...recover 实现 try...逻辑。
// 如果发生任何异常，它将返回错误，否则返回 nil。
func X异常捕捉(上下文 context.Context, 处理函数 func(上下文 context.Context)) (错误 error) {
	if 处理函数 == nil {
		return
	}
	defer func() {
		if exception := recover(); exception != nil {
			if v, ok := exception.(error); ok && gerror.HasStack(v) {
				错误 = v
			} else {
				错误 = gerror.NewCodef(gcode.CodeInternalPanic, "%+v", exception)
			}
		}
	}()
	处理函数(上下文)
	return
}

// TryCatch 实现了类似 `try...catch...` 的逻辑，通过内部使用 `panic...recover`。
// 如果发生任何异常，它会自动调用函数 `catch` 并将异常作为 error 传递给 `catch` 函数。
// 若传入的 `catch` 为 nil，则忽略来自 `try` 的 panic，并且不会向父级 goroutine 抛出 panic。
//
// 但是请注意，如果函数 `catch` 本身也抛出了 panic，则当前 goroutine 将会 panic。
func X异常捕捉并带异常处理(上下文 context.Context, 处理函数 func(上下文 context.Context), 异常处理函数 func(上下文 context.Context, 错误 error)) {
	if 处理函数 == nil {
		return
	}
	if exception := X异常捕捉(上下文, 处理函数); exception != nil && 异常处理函数 != nil {
		异常处理函数(上下文, exception)
	}
}
