// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

// localMetricInstrument 实现了 MetricInstrument 接口。 md5:36ddf74b71077f27
type localInstrumentInfo struct {
	name    string
	version string
}

// newInstrumentInfo 创建并返回一个 MetricInstrument。 md5:6fedce2350c59b80
func (meter *localMeter) newInstrumentInfo() InstrumentInfo {
	return &localInstrumentInfo{
		name:    meter.Instrument,
		version: meter.InstrumentVersion,
	}
}

// Name返回度量的仪器名称。 md5:a45a5bfd2d0859c1
func (l *localInstrumentInfo) Name() string {
	return l.name
}

// Version 返回指标的仪器版本。 md5:de1a566af4fc0f5a
func (l *localInstrumentInfo) Version() string {
	return l.version
}
