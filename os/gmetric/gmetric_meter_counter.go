// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

// localCounter是实现接口Counter的本地实例。 md5:65f564de04eb1400
type localCounter struct {
	Metric
	MeterOption
	MetricOption
	CounterPerformer
}

var (
		// 检查是否实现了MetricInitializer接口。 md5:87bf7f014f5d05df
	_ MetricInitializer = (*localCounter)(nil)
		// 检查对于接口PerformerExporter的实现。 md5:7bc09f325273ded9
	_ PerformerExporter = (*localCounter)(nil)
)

// Counter 创建并返回一个新的 Counter.. md5:84e33be2f1339329
func (meter *localMeter) Counter(name string, option MetricOption) (Counter, error) {
	m, err := meter.newMetric(MetricTypeCounter, name, option)
	if err != nil {
		return nil, err
	}
	counter := &localCounter{
		Metric:           m,
		MeterOption:      meter.MeterOption,
		MetricOption:     option,
		CounterPerformer: newNoopCounterPerformer(),
	}
	if globalProvider != nil {
		if err = counter.Init(globalProvider); err != nil {
			return nil, err
		}
	}
	allMetrics = append(allMetrics, counter)
	return counter, nil
}

// MustCounter 创建并返回一个新的计数器。
// 如果发生任何错误，它将引发恐慌。
// md5:5a16e08ea093036c
func (meter *localMeter) MustCounter(name string, option MetricOption) Counter {
	m, err := meter.Counter(name, option)
	if err != nil {
		panic(err)
	}
	return m
}

// Init 在创建Provider时初始化Metric。 md5:a46b2bb4d31aa7d0
func (l *localCounter) Init(provider Provider) (err error) {
	if _, ok := l.CounterPerformer.(noopCounterPerformer); !ok {
		// already initialized.
		return
	}
	l.CounterPerformer, err = provider.MeterPerformer(l.MeterOption).CounterPerformer(
		l.Info().Name(),
		l.MetricOption,
	)
	return
}

// Performer 实现了 PerformerExporter 接口，该接口用于导出 Metric 的内部 Performer。
// 这通常被指标实现所使用。
// md5:e521fc985b9a53e2
func (l *localCounter) Performer() any {
	return l.CounterPerformer
}
