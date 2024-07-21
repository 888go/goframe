// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

// localObservableUpDownCounter是实现ObservableUpDownCounter接口的本地实现。 md5:be0cdcbf24029cf7
type localObservableUpDownCounter struct {
	Metric
	MeterOption
	MetricOption
	ObservableUpDownCounterPerformer
}

var (
	// 检查是否实现了MetricInitializer接口。 md5:87bf7f014f5d05df
	_ MetricInitializer = (*localObservableUpDownCounter)(nil)
	// 检查对于接口PerformerExporter的实现。 md5:7bc09f325273ded9
	_ PerformerExporter = (*localObservableUpDownCounter)(nil)
)

// ObservableUpDownCounter 创建并返回一个新的ObservableUpDownCounter。 md5:a7f48b253e6c2099
// ff:
// meter:
// name:
// option:
// ObservableUpDownCounter:
func (meter *localMeter) ObservableUpDownCounter(name string, option MetricOption) (ObservableUpDownCounter, error) {
	m, err := meter.newMetric(MetricTypeObservableUpDownCounter, name, option)
	if err != nil {
		return nil, err
	}
	observableUpDownCounter := &localObservableUpDownCounter{
		Metric:                           m,
		MeterOption:                      meter.MeterOption,
		MetricOption:                     option,
		ObservableUpDownCounterPerformer: newNoopObservableUpDownCounterPerformer(),
	}
	if globalProvider != nil {
		if err = observableUpDownCounter.Init(globalProvider); err != nil {
			return nil, err
		}
	}
	allMetrics = append(allMetrics, observableUpDownCounter)
	return observableUpDownCounter, nil
}

// MustObservableUpDownCounter 创建并返回一个新的 ObservableUpDownCounter。
// 如果发生任何错误，它将 panic。
// md5:2c1790e420456ea1
// ff:
// meter:
// name:
// option:
func (meter *localMeter) MustObservableUpDownCounter(name string, option MetricOption) ObservableUpDownCounter {
	m, err := meter.ObservableCounter(name, option)
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
func (l *localObservableUpDownCounter) Init(provider Provider) (err error) {
	if _, ok := l.ObservableUpDownCounterPerformer.(noopObservableUpDownCounterPerformer); !ok {
		// already initialized.
		return
	}
	l.ObservableUpDownCounterPerformer, err = provider.MeterPerformer(l.MeterOption).ObservableUpDownCounterPerformer(
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
func (l *localObservableUpDownCounter) Performer() any {
	return l.ObservableUpDownCounterPerformer
}
