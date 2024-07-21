// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

// localObservableCounter 是接口 ObservableCounter 的本地实现。 md5:49a0950ea1c362dc
type localObservableCounter struct {
	Metric
	MeterOption
	MetricOption
	ObservableCounterPerformer
}

var (
	// 检查是否实现了MetricInitializer接口。 md5:87bf7f014f5d05df
	_ MetricInitializer = (*localObservableCounter)(nil)
	// 检查对于接口PerformerExporter的实现。 md5:7bc09f325273ded9
	_ PerformerExporter = (*localObservableCounter)(nil)
)

// ObservableCounter 创建并返回一个新的 ObservableCounter。 md5:1fb1055edede2f1e
// ff:
// meter:
// name:
// option:
// ObservableCounter:
func (meter *localMeter) ObservableCounter(name string, option MetricOption) (ObservableCounter, error) {
	m, err := meter.newMetric(MetricTypeObservableCounter, name, option)
	if err != nil {
		return nil, err
	}
	observableCounter := &localObservableCounter{
		Metric:                     m,
		MeterOption:                meter.MeterOption,
		MetricOption:               option,
		ObservableCounterPerformer: newNoopObservableCounterPerformer(),
	}
	if globalProvider != nil {
		if err = observableCounter.Init(globalProvider); err != nil {
			return nil, err
		}
	}
	allMetrics = append(allMetrics, observableCounter)
	return observableCounter, nil
}

// MustObservableCounter 创建并返回一个新的 ObservableCounter。
// 如果发生任何错误，它将引发恐慌。
// md5:0e100a900e418612
// ff:
// meter:
// name:
// option:
func (meter *localMeter) MustObservableCounter(name string, option MetricOption) ObservableCounter {
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
func (l *localObservableCounter) Init(provider Provider) (err error) {
	if _, ok := l.ObservableCounterPerformer.(noopObservableCounterPerformer); !ok {
		// already initialized.
		return
	}
	l.ObservableCounterPerformer, err = provider.MeterPerformer(l.MeterOption).ObservableCounterPerformer(
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
func (l *localObservableCounter) Performer() any {
	return l.ObservableCounterPerformer
}
