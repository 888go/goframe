// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gctx
import (
	"context"
	"time"
	)
// neverDoneCtx 永不结束
type neverDoneCtx struct {
	context.Context
}

// Done 从父上下文中禁止使用已完成的上下文。
func (*neverDoneCtx) Done() <-chan struct{} {
	return nil
}

// Deadline 禁止从父上下文中继承截止时间（deadline），即设置当前上下文不具有来自父上下文的截止时间。
func (*neverDoneCtx) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

// Err用于禁止从父上下文中获取完成信号。
func (c *neverDoneCtx) Err() error {
	return nil
}

// NeverDone 包装并返回一个新的上下文对象，该对象将永不完成，
// 这样可以禁止手动完成上下文，以使上下文能够传播到异步goroutine中。
func NeverDone(ctx context.Context) context.Context {
	return &neverDoneCtx{ctx}
}
