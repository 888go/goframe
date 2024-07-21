// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

import "context"

// noopUpDownCounterPerformer 是一个实现了 CounterPerformer 接口但没有实际操作的实现者。 md5:072081413ac6be4f
type noopUpDownCounterPerformer struct{}

// newNoopUpDownCounterPerformer 创建并返回一个没有真正操作的CounterPerformer。 md5:e7e1846c3b34c796
func newNoopUpDownCounterPerformer() UpDownCounterPerformer {
	return noopUpDownCounterPerformer{}
}

// Inc 将计数器增加1。 md5:81a1f6bd9f2e53ca
// ff:
// noopUpDownCounterPerformer:
// ctx:
// option:
func (noopUpDownCounterPerformer) Inc(ctx context.Context, option ...Option) {}

// Dec 将计数器减1。 md5:a57e133a130dce44
// ff:
// noopUpDownCounterPerformer:
// ctx:
// option:
func (noopUpDownCounterPerformer) Dec(ctx context.Context, option ...Option) {}

// Add 将给定的值添加到计数器中。 md5:52aba4021b7f61b2
// ff:
// noopUpDownCounterPerformer:
// ctx:
// increment:
// option:
func (noopUpDownCounterPerformer) Add(ctx context.Context, increment float64, option ...Option) {}
