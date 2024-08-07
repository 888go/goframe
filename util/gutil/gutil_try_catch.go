// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 工具类

import (
	"context"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
)

// X异常输出 抛出一个异常，这个异常可以被 TryCatch 结构或 recover 语句捕获。 md5:44e69b1d2fded57f
func X异常输出(消息 interface{}) {
	panic(消息)
}

// X异常捕捉 使用内部的 panic...recover 实现 try...catch 逻辑。如果发生任何异常，它会返回错误；否则返回 nil。
// md5:7c85aa857bb16fca
func X异常捕捉(上下文 context.Context, 处理函数 func(上下文 context.Context)) (错误 error) {
	if 处理函数 == nil {
		return
	}
	defer func() {
		if exception := recover(); exception != nil {
			if v, ok := exception.(error); ok && gerror.X判断是否带堆栈(v) {
				错误 = v
			} else {
				错误 = gerror.X创建错误码并格式化(gcode.CodeInternalPanic, "%+v", exception)
			}
		}
	}()
	处理函数(上下文)
	return
}

// X异常捕捉并带异常处理 实现了类似于 `try...catch...` 的异常处理机制，利用内部的 `panic...recover` 逻辑。
// 当在 `try` 块中发生任何异常时，它会自动调用函数 `catch` 并将异常传递给错误参数。
// 如果 `catch` 函数为 nil，则忽略 `try` 中的 panic，不会向父 goroutine 抛出 panic。
//
// 但是，需要注意的是，如果 `catch` 函数自身也抛出 panic，当前 goroutine 将会 panic。
// md5:6cbe568aa0940b07
func X异常捕捉并带异常处理(上下文 context.Context, 处理函数 func(上下文 context.Context), 异常处理函数 func(上下文 context.Context, 错误 error)) {
	if 处理函数 == nil {
		return
	}
	if exception := X异常捕捉(上下文, 处理函数); exception != nil && 异常处理函数 != nil {
		异常处理函数(上下文, exception)
	}
}
