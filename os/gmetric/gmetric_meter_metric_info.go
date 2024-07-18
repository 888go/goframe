// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

import "fmt"

// localMetricInfo 实现了 MetricInfo 接口。 md5:0d4cae443266be67
type localMetricInfo struct {
	MetricType
	MetricOption
	InstrumentInfo
	MetricName string
}

// newMetricInfo 创建并返回一个 MetricInfo 实例。 md5:a8bf142202bb8a34
func (meter *localMeter) newMetricInfo(
	metricType MetricType, metricName string, metricOption MetricOption,
) MetricInfo {
	return &localMetricInfo{
		MetricName:     metricName,
		MetricType:     metricType,
		MetricOption:   metricOption,
		InstrumentInfo: meter.newInstrumentInfo(),
	}
}

// Name 返回度量指标的名称。 md5:d3300c9b003bf9e9
// ff:
// l:
func (l *localMetricInfo) Name() string {
	return l.MetricName
}

// Help 返回度量指标的帮助描述。 md5:af48acbd53473872
// ff:
// l:
func (l *localMetricInfo) Help() string {
	return l.MetricOption.Help
}

// Unit 返回度量的单位名称。 md5:985c39bd14c3f833
// ff:
// l:
func (l *localMetricInfo) Unit() string {
	return l.MetricOption.Unit
}

// Type返回度量的类型。 md5:5ceeadaee6d55192
// ff:
// l:
func (l *localMetricInfo) Type() MetricType {
	return l.MetricType
}

// Attributes 返回度量指标的常量属性切片。 md5:ca84de54457cab11
// ff:
// l:
func (l *localMetricInfo) Attributes() Attributes {
	return l.MetricOption.Attributes
}

// Instrument 返回指标的乐器信息。 md5:5024719a6126d731
// ff:
// l:
func (l *localMetricInfo) Instrument() InstrumentInfo {
	return l.InstrumentInfo
}

// ff:
// l:
func (l *localMetricInfo) Key() string {
	if l.Instrument().Name() != "" && l.Instrument().Version() != "" {
		return fmt.Sprintf(
			`%s@%s:%s`,
			l.Instrument().Name(),
			l.Instrument().Version(),
			l.Name(),
		)
	}
	if l.Instrument().Name() != "" && l.Instrument().Version() == "" {
		return fmt.Sprintf(
			`%s:%s`,
			l.Instrument().Name(),
			l.Name(),
		)
	}
	return l.Name()
}
