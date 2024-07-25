// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。 md5:c14c707c81272457

package gmetric

// noopObservableGaugePerformer是一个实现了ObservableGaugePerformer接口的实现者，但没有实际的操作。 md5:91261d11120ce302
type noopObservableGaugePerformer struct{}

// newNoopObservableGaugePerformer 创建并返回一个没有真正操作的ObservableGaugePerformer。 md5:bb41e3baf0049746
func newNoopObservableGaugePerformer() ObservableGaugePerformer {
	return noopObservableGaugePerformer{}
}

func (noopObservableGaugePerformer) observable() {}
