// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gctx

import (
	"context"
	"time"
)

// neverDoneCtx 永远不会完成。 md5:9cd0926cf01acafc
type neverDoneCtx struct {
	context.Context
}

// Done 禁止从父上下文中关闭 done。 md5:6ee3971853766306
func (*neverDoneCtx) Done() <-chan struct{} {
	return nil
}

// Deadline 禁止从父上下文中继承截止期限。 md5:b0a8033fcfdd0483
func (*neverDoneCtx) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

// Err 限制了从父上下文中完成的上下文。 md5:605f4a25a7f54817
func (c *neverDoneCtx) Err() error {
	return nil
}

// NeverDone 包装并返回一个永远不会完成的新上下文对象，这禁止手动完成上下文，使得上下文可以传递给异步的 Goroutine。
//
// 请注意，这不会影响父上下文的关闭（取消），因为它只是父上下文的包装器，只影响下一个上下文处理。 md5:38b63a322c5449a9
func NeverDone(ctx context.Context) context.Context {
	return &neverDoneCtx{ctx}
}
