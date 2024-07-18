// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

// localMeter 是实现Meter接口的本地版本。 md5:e59e8a96222d14c4
type localMeter struct {
	MeterOption
}

// MeterOption 是用于创建Meter的创建选项。 md5:aeec4cb6c20611f7
type MeterOption struct {
// Instrument 是将此 Metric 绑定到全局度量提供程序的指标名称。这是指标的可选配置。
// md5:c68d90eb2e2a5738
	Instrument string

// InstrumentVersion 是用于将此指标绑定到全局 MeterProvider 的仪器版本。
// 这是指标的一个可选配置。
// md5:9e0f63bfaddf4047
	InstrumentVersion string

// Attributes 保存了计量器（Meter）中所有指标的常量键值对描述元数据。
// 这是一个可选的仪表配置。
// md5:b57d2f1c17951d62
	Attributes Attributes
}

// newMeter 创建并返回一个实现了Meter接口的对象。 md5:f52639f31640714f
func newMeter(option MeterOption) Meter {
	return &localMeter{
		MeterOption: option,
	}
}

// Performer 为 Meter 创建并返回表演者。 md5:78f41b5ebd242dd8
// ff:
// meter:
func (meter *localMeter) Performer() MeterPerformer {
	if globalProvider == nil {
		return nil
	}
	return globalProvider.MeterPerformer(meter.MeterOption)
}
