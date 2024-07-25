// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gmetric_test

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/os/gmetric"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Counter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			meterOption = gmetric.MeterOption{
				Instrument:        "github.com/gogf/gf/example/metric/basic",
				InstrumentVersion: "v1.0",
			}
			metricName   = "goframe.metric.demo.counter"
			metricOption = gmetric.MetricOption{
				Help: "This is a simple demo for Counter usage",
				Unit: "%",
				Attributes: gmetric.Attributes{
					gmetric.NewAttribute("const_label_a", 1),
				},
			}
			meter   = gmetric.GetGlobalProvider().Meter(meterOption)
			counter = meter.MustCounter(metricName, metricOption)
		)
		t.Assert(counter.Info().Name(), metricName)
		t.Assert(counter.Info().Help(), metricOption.Help)
		t.Assert(counter.Info().Unit(), metricOption.Unit)
		t.Assert(counter.Info().Attributes(), metricOption.Attributes)
		t.Assert(counter.Info().Instrument().Name(), meterOption.Instrument)
		t.Assert(counter.Info().Instrument().Version(), meterOption.InstrumentVersion)
	})
}

func Test_Histogram(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			meterOption = gmetric.MeterOption{
				Instrument:        "github.com/gogf/gf/example/metric/basic",
				InstrumentVersion: "v1.0",
			}
			metricName   = "goframe.metric.demo.histogram"
			metricOption = gmetric.MetricOption{
				Help: "This is a simple demo for Histogram usage",
				Unit: "%",
				Attributes: gmetric.Attributes{
					gmetric.NewAttribute("const_label_a", 1),
				},
				Buckets: []float64{0, 10, 20, 50, 100, 500, 1000, 2000, 5000, 10000},
			}
			meter     = gmetric.GetGlobalProvider().Meter(meterOption)
			histogram = meter.MustHistogram(metricName, metricOption)
		)
		t.Assert(histogram.Info().Name(), metricName)
		t.Assert(histogram.Info().Help(), metricOption.Help)
		t.Assert(histogram.Info().Unit(), metricOption.Unit)
		t.Assert(histogram.Info().Attributes(), metricOption.Attributes)
		t.Assert(histogram.Info().Instrument().Name(), meterOption.Instrument)
		t.Assert(histogram.Info().Instrument().Version(), meterOption.InstrumentVersion)
		t.Assert(histogram.Buckets(), metricOption.Buckets)
	})
}

func Test_CommonAttributes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		commonAttributes := gmetric.CommonAttributes()
		t.AssertGT(len(commonAttributes), 1)
		fmt.Println(commonAttributes)
	})
}
