// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

// localHistogram 是接口 Histogram 的本地实现。 md5:85a2ceed0ede2ded
type localHistogram struct {
	Metric
	MeterOption
	MetricOption
	HistogramPerformer
}

var (
	// 检查是否实现了MetricInitializer接口。 md5:87bf7f014f5d05df
	_ MetricInitializer = (*localHistogram)(nil)
	// 检查对于接口PerformerExporter的实现。 md5:7bc09f325273ded9
	_ PerformerExporter = (*localHistogram)(nil)
)

// Histogram 创建并返回一个新的 Histogram.. md5:8a66ea5ba65143f0
// ff:
// meter:
// name:
// option:
// Histogram:
func (meter *localMeter) Histogram(name string, option MetricOption) (Histogram, error) {
	m, err := meter.newMetric(MetricTypeHistogram, name, option)
	if err != nil {
		return nil, err
	}
	histogram := &localHistogram{
		Metric:             m,
		MeterOption:        meter.MeterOption,
		MetricOption:       option,
		HistogramPerformer: newNoopHistogramPerformer(),
	}
	if globalProvider != nil {
		if err = histogram.Init(globalProvider); err != nil {
			return nil, err
		}
	}
	allMetrics = append(allMetrics, histogram)
	return histogram, nil
}

// MustHistogram 创建并返回一个新的Histogram。
// 如果发生任何错误，它将引发恐慌。
// md5:3716fed48bf43141
// ff:
// meter:
// name:
// option:
func (meter *localMeter) MustHistogram(name string, option MetricOption) Histogram {
	m, err := meter.Histogram(name, option)
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
func (l *localHistogram) Init(provider Provider) (err error) {
	if _, ok := l.HistogramPerformer.(noopHistogramPerformer); !ok {
		// already initialized.
		return
	}
	l.HistogramPerformer, err = provider.MeterPerformer(l.MeterOption).HistogramPerformer(
		l.Info().Name(),
		l.MetricOption,
	)
	return err
}

// Buckets 返回Histogram的桶数组。 md5:b0cdc9def1273944
// ff:
// l:
func (l *localHistogram) Buckets() []float64 {
	return l.MetricOption.Buckets
}

// Performer 实现了 PerformerExporter 接口，该接口用于导出 Metric 的内部 Performer。
// 这通常被指标实现所使用。
// md5:e521fc985b9a53e2
// ff:
// l:
func (l *localHistogram) Performer() any {
	return l.HistogramPerformer
}
