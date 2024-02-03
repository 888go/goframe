// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gutil

import (
	"context"
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
	if goroutineFunc == nil {
		return
	}
	go TryCatch(ctx, goroutineFunc, recoverFunc)
}
