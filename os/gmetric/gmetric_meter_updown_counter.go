// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

// localUpDownCounter 是 UpDownCounter 接口的本地实现。. md5:12f0deffd8bcf0dc
type localUpDownCounter struct {
	Metric
	MeterOption
	MetricOption
	UpDownCounterPerformer
}

var (
	// 检查是否实现了MetricInitializer接口。. md5:87bf7f014f5d05df
	_ MetricInitializer = (*localUpDownCounter)(nil)
	// 检查对于接口PerformerExporter的实现。. md5:7bc09f325273ded9
	_ PerformerExporter = (*localUpDownCounter)(nil)
)

//UpDownCounter 创建并返回一个新的Counter。. md5:0fd05c1ce07af34f
func (meter *localMeter) UpDownCounter(name string, option MetricOption) (UpDownCounter, error) {
	m, err := meter.newMetric(MetricTypeUpDownCounter, name, option)
	if err != nil {
		return nil, err
	}
	updownCounter := &localUpDownCounter{
		Metric:                 m,
		MeterOption:            meter.MeterOption,
		MetricOption:           option,
		UpDownCounterPerformer: newNoopUpDownCounterPerformer(),
	}
	if globalProvider != nil {
		if err = updownCounter.Init(globalProvider); err != nil {
			return nil, err
		}
	}
	allMetrics = append(allMetrics, updownCounter)
	return updownCounter, nil
}

// MustUpDownCounter 创建并返回一个新的计数器。
// 如果发生任何错误，它将引发恐慌。
// md5:02c439bc7eddaccb
func (meter *localMeter) MustUpDownCounter(name string, option MetricOption) UpDownCounter {
	m, err := meter.UpDownCounter(name, option)
	if err != nil {
		panic(err)
	}
	return m
}

// Init 在创建Provider时初始化Metric。. md5:a46b2bb4d31aa7d0
func (l *localUpDownCounter) Init(provider Provider) (err error) {
	if _, ok := l.UpDownCounterPerformer.(noopUpDownCounterPerformer); !ok {
		// already initialized.
		return
	}
	l.UpDownCounterPerformer, err = provider.MeterPerformer(l.MeterOption).UpDownCounterPerformer(
		l.Info().Name(),
		l.MetricOption,
	)
	return
}

// Performer 实现了 PerformerExporter 接口，该接口用于导出 Metric 的内部 Performer。
// 这通常被指标实现所使用。
// md5:e521fc985b9a53e2
func (l *localUpDownCounter) Performer() any {
	return l.UpDownCounterPerformer
}
