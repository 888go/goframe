// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

// localObservableGauge是ObservableGauge接口的本地实现。 md5:7efef87b9cf8f215
type localObservableGauge struct {
	Metric
	MeterOption
	MetricOption
	ObservableGaugePerformer
}

var (
	// 检查是否实现了MetricInitializer接口。 md5:87bf7f014f5d05df
	_ MetricInitializer = (*localObservableGauge)(nil)
	// 检查对于接口PerformerExporter的实现。 md5:7bc09f325273ded9
	_ PerformerExporter = (*localObservableGauge)(nil)
)

// ObservableGauge 创建并返回一个新的可观察计量表。 md5:406f093f6a405dd4
// ff:
// meter:
// name:
// option:
// ObservableGauge:
func (meter *localMeter) ObservableGauge(name string, option MetricOption) (ObservableGauge, error) {
	m, err := meter.newMetric(MetricTypeObservableGauge, name, option)
	if err != nil {
		return nil, err
	}
	observableGauge := &localObservableGauge{
		Metric:                   m,
		MeterOption:              meter.MeterOption,
		MetricOption:             option,
		ObservableGaugePerformer: newNoopObservableGaugePerformer(),
	}
	if globalProvider != nil {
		if err = observableGauge.Init(globalProvider); err != nil {
			return nil, err
		}
	}
	allMetrics = append(allMetrics, observableGauge)
	return observableGauge, nil
}

// MustObservableGauge 创建并返回一个新的 ObservableGauge。
// 如果发生任何错误，它将引发 panic。
// md5:fba8400cc344af01
// ff:
// meter:
// name:
// option:
func (meter *localMeter) MustObservableGauge(name string, option MetricOption) ObservableGauge {
	m, err := meter.ObservableGauge(name, option)
	if err != nil {
		panic(err)
	}
	return m
}

// Init 在创建Provider时初始化Metric。 md5:a46b2bb4d31aa7d0
// ff:
// l:
// provider:
// err:
func (l *localObservableGauge) Init(provider Provider) (err error) {
	if _, ok := l.ObservableGaugePerformer.(noopObservableGaugePerformer); !ok {
		// already initialized.
		return
	}
	l.ObservableGaugePerformer, err = provider.MeterPerformer(l.MeterOption).ObservableGaugePerformer(
		l.Info().Name(),
		l.MetricOption,
	)
	return err
}

// Performer 实现了 PerformerExporter 接口，该接口用于导出 Metric 的内部 Performer。
// 这通常被指标实现所使用。
// md5:e521fc985b9a53e2
// ff:
// l:
func (l *localObservableGauge) Performer() any {
	return l.ObservableGaugePerformer
}
