// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

// GlobalProvider 用于创建Meter和Metric的入口。GlobalProvider 只有一个用于Meter创建的函数，设计目的是方便使用。
// md5:25431281857b5d86
type GlobalProvider interface {
	// Meter 使用MeterOption创建并返回Meter。 md5:d1e5c4bca1f6e03c
	Meter(option MeterOption) Meter
}

// Meter 存放了创建各种指标功能的函数。 md5:ecfd04354d47531f
type Meter interface {
	// Counter 创建并返回一个新的 Counter.. md5:84e33be2f1339329
	Counter(name string, option MetricOption) (Counter, error)

	// UpDownCounter 创建并返回一个新的UpDownCounter.. md5:17cdda79297f292f
	UpDownCounter(name string, option MetricOption) (UpDownCounter, error)

	// Histogram 创建并返回一个新的 Histogram.. md5:8a66ea5ba65143f0
	Histogram(name string, option MetricOption) (Histogram, error)

	// ObservableCounter 创建并返回一个新的 ObservableCounter。 md5:1fb1055edede2f1e
	ObservableCounter(name string, option MetricOption) (ObservableCounter, error)

	// ObservableUpDownCounter 创建并返回一个新的ObservableUpDownCounter。 md5:a7f48b253e6c2099
	ObservableUpDownCounter(name string, option MetricOption) (ObservableUpDownCounter, error)

	// ObservableGauge 创建并返回一个新的可观察计量表。 md5:406f093f6a405dd4
	ObservableGauge(name string, option MetricOption) (ObservableGauge, error)

// MustCounter 创建并返回一个新的计数器。
// 如果发生任何错误，它将引发恐慌。
// md5:8ca39b864372ccfe
	MustCounter(name string, option MetricOption) Counter

// MustUpDownCounter 创建并返回一个新的UpDownCounter。
// 如果发生任何错误，它将引发恐慌。
// md5:9bb6fb57771f0266
	MustUpDownCounter(name string, option MetricOption) UpDownCounter

// MustHistogram 创建并返回一个新的直方图。
// 如果发生任何错误，它将引发恐慌。
// md5:fc31a9bb5a94fd34
	MustHistogram(name string, option MetricOption) Histogram

// MustObservableCounter 创建并返回一个新的可观察计数器。如果发生任何错误，它将 panic。
// md5:d12041c97b0c5aa2
	MustObservableCounter(name string, option MetricOption) ObservableCounter

// MustObservableUpDownCounter 创建并返回一个新的 ObservableUpDownCounter。
// 如果发生任何错误，它将引发 panic。
// md5:04565499baba4d44
	MustObservableUpDownCounter(name string, option MetricOption) ObservableUpDownCounter

// MustObservableGauge 创建并返回一个新的 ObservableGauge。
// 如果发生任何错误，它将引发恐慌。
// md5:f45a16dcd373e219
	MustObservableGauge(name string, option MetricOption) ObservableGauge

// RegisterCallback 在某些指标上注册回调函数。
// 回调函数与特定的组件和版本绑定，当关联的指标被读取时会被调用。
// 同一个组件和版本可以注册多个回调函数，它们将按照注册的顺序被调用。
// md5:89a5acee144aeb40
	RegisterCallback(callback Callback, canBeCallbackMetrics ...ObservableMetric) error

	// MustRegisterCallback 类似于 RegisterCallback，但是如果发生任何错误，它会直接 panic。 md5:41b35f310c8c461d
	MustRegisterCallback(callback Callback, canBeCallbackMetrics ...ObservableMetric)
}

type localGlobalProvider struct {
}

var (
	// globalProvider 是用于全局使用的提供者。 md5:6345c47ee4cbbaf6
	globalProvider Provider
)

// GetGlobalProvider 获取GetGlobalProvider实例。 md5:88505f26db856ed4
// ff:
func GetGlobalProvider() GlobalProvider {
	return &localGlobalProvider{}
}

// SetGlobalProvider 将 `provider` 注册为全局提供者，
// 这意味着后续创建的指标将基于全局提供者。
// md5:ca9b6936745e89a3
// ff:
// provider:
func SetGlobalProvider(provider Provider) {
	globalProvider = provider
}

// Meter 使用MeterOption创建并返回Meter。 md5:d1e5c4bca1f6e03c
// ff:
// l:
// option:
func (l *localGlobalProvider) Meter(option MeterOption) Meter {
	return newMeter(option)
}
