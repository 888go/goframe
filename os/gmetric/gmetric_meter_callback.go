// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。 md5:c14c707c81272457

package gmetric

// CallbackItem 是已注册的全局回调项。 md5:afee0a9b376754c0
type CallbackItem struct {
	Callback    Callback           // Global callback.
	Metrics     []ObservableMetric // 在特定指标上回调。 md5:1b8f919fb6f66516
	MeterOption MeterOption        // MeterOption 是度量器所持有的选项。 md5:7cbdb72b93713b5d
	Provider    Provider           // Provider 是回调项绑定到的 Provider。 md5:5ad4c5696a74f008
}

var (
	// Registered callbacks.
	globalCallbackItems = make([]CallbackItem, 0)
)

// RegisterCallback 在特定指标上注册回调。
// 回调与特定组件和版本绑定，当关联的指标被读取时会被调用。
// 同一组件和版本上的多个回调将按照它们注册的顺序被调用。 md5:a7b0f2e948a5cd42
func (meter *localMeter) RegisterCallback(callback Callback, observableMetrics ...ObservableMetric) error {
	if len(observableMetrics) == 0 {
		return nil
	}
	globalCallbackItems = append(globalCallbackItems, CallbackItem{
		Callback:    callback,
		Metrics:     observableMetrics,
		MeterOption: meter.MeterOption,
	})
	return nil
}

// MustRegisterCallback 类似于 RegisterCallback，但是如果发生任何错误，它会直接 panic。 md5:41b35f310c8c461d
func (meter *localMeter) MustRegisterCallback(callback Callback, observableMetrics ...ObservableMetric) {
	err := meter.RegisterCallback(callback, observableMetrics...)
	if err != nil {
		panic(err)
	}
}

// GetRegisteredCallbacks 获取并返回已注册的全局回调函数。
// 如果返回回调函数，会截断回调函数切片。 md5:22e1858dfd047cb0
func GetRegisteredCallbacks() []CallbackItem {
	items := globalCallbackItems
	globalCallbackItems = globalCallbackItems[:0]
	return items
}
