// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。 md5:c14c707c81272457

package gmetric

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gregex"
)

// localMetric 实现了 Metric 接口。 md5:1d7e92821badcf91
type localMetric struct {
	MetricInfo
}

// newMetric 创建并返回一个实现Metric接口的对象。 md5:566502ff9b514701
func (meter *localMeter) newMetric(
	metricType MetricType, metricName string, metricOption MetricOption,
) (Metric, error) {
	if metricName == "" {
		return nil, gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`error creating %s metric while given name is empty, option: %s`,
			metricType, gjson.MustEncodeString(metricOption),
		)
	}
	if !gregex.IsMatchString(MetricNamePattern, metricName) {
		return nil, gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid metric name "%s", should match regular expression pattern "%s"`,
			metricName, MetricNamePattern,
		)
	}
	return &localMetric{
		MetricInfo: meter.newMetricInfo(metricType, metricName, metricOption),
	}, nil
}

// Info 返回一个Metric的基本信息。 md5:d521e5fdeb6e591f
func (l *localMetric) Info() MetricInfo {
	return l.MetricInfo
}
