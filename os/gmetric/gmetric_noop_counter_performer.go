// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

import "context"

// noopCounterPerformer 是一个实现了 CounterPerformer 接口但没有实际操作的实现者。 md5:6266c86e4f91ce59
type noopCounterPerformer struct{}

// newNoopCounterPerformer 创建并返回一个没有真正操作的CounterPerformer。 md5:1aa76ee5b4d4985d
func newNoopCounterPerformer() CounterPerformer {
	return noopCounterPerformer{}
}

// Inc 将计数器增加1。 md5:81a1f6bd9f2e53ca
// ff:
// noopCounterPerformer:
// ctx:
// option:
func (noopCounterPerformer) Inc(ctx context.Context, option ...Option) {}

// Add将给定的值添加到计数器中。如果值小于0，则会引发恐慌。 md5:d0e547634cd9c23f
// ff:
// noopCounterPerformer:
// ctx:
// increment:
// option:
func (noopCounterPerformer) Add(ctx context.Context, increment float64, option ...Option) {}
