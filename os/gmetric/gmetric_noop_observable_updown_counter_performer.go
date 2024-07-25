// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。 md5:c14c707c81272457

package gmetric

// noopObservableUpDownCounterPerformer 是一个实现了 ObservableUpDownCounterPerformer 接口的实现者，
// 但其中并不包含实际操作。 md5:8c3502d99cc1c3e5
type noopObservableUpDownCounterPerformer struct{}

// newNoopObservableUpDownCounterPerformer 创建并返回一个无实际操作的 ObservableUpDownCounterPerformer. md5:fdae3e6ba842dbe2
func newNoopObservableUpDownCounterPerformer() ObservableUpDownCounterPerformer {
	return noopObservableUpDownCounterPerformer{}
}

func (noopObservableUpDownCounterPerformer) observable() {}
