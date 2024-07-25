// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

// noopObservableCounterPerformer 是一个实现了 ObservableCounterPerformer 接口的实现者，但其中没有任何实际操作。 md5:a91562bcf10391c7
type noopObservableCounterPerformer struct{}

// newNoopObservableCounterPerformer 创建并返回一个没有实际操作的 ObservableCounterPerformer。 md5:4076ea9366a1a5e7
func newNoopObservableCounterPerformer() ObservableCounterPerformer {
	return noopObservableCounterPerformer{}
}

func (noopObservableCounterPerformer) observable() {}
