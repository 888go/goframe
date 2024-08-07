// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"net"
	"net/http"

	"github.com/888go/goframe"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gmetric"
	gstr "github.com/888go/goframe/text/gstr"
)

type localMetricManager struct {
	HttpServerRequestActive        gmetric.UpDownCounter
	HttpServerRequestTotal         gmetric.Counter
	HttpServerRequestDuration      gmetric.Histogram
	HttpServerRequestDurationTotal gmetric.Counter
	HttpServerRequestBodySize      gmetric.Counter
	HttpServerResponseBodySize     gmetric.Counter
}

const (
	metricAttrKeyServerAddress          = "server.address"
	metricAttrKeyServerPort             = "server.port"
	metricAttrKeyHttpRoute              = "http.route"
	metricAttrKeyUrlSchema              = "url.schema"
	metricAttrKeyHttpRequestMethod      = "http.request.method"
	metricAttrKeyErrorCode              = "error.code"
	metricAttrKeyHttpResponseStatusCode = "http.response.status_code"
	metricAttrKeyNetworkProtocolVersion = "network.protocol.version"
)

var (
		// metricManager是用于HTTP服务器指标的。 md5:c022dc479dce1fbb
	metricManager = newMetricManager()
)

func newMetricManager() *localMetricManager {
	meter := gmetric.GetGlobalProvider().Meter(gmetric.MeterOption{
		Instrument:        instrumentName,
		InstrumentVersion: gf.VERSION,
	})
	mm := &localMetricManager{
		HttpServerRequestDuration: meter.MustHistogram(
			"http.server.request.duration",
			gmetric.MetricOption{
				Help:       "Measures the duration of inbound request.",
				Unit:       "ms",
				Attributes: gmetric.Attributes{},
				Buckets: []float64{
					1,
					5,
					10,
					25,
					50,
					75,
					100,
					250,
					500,
					750,
					1000,
					2500,
					5000,
					7500,
					10000,
					30000,
					60000,
				},
			},
		),
		HttpServerRequestTotal: meter.MustCounter(
			"http.server.request.total",
			gmetric.MetricOption{
				Help:       "Total processed request number.",
				Unit:       "",
				Attributes: gmetric.Attributes{},
			},
		),
		HttpServerRequestActive: meter.MustUpDownCounter(
			"http.server.request.active",
			gmetric.MetricOption{
				Help:       "Number of active server requests.",
				Unit:       "",
				Attributes: gmetric.Attributes{},
			},
		),
		HttpServerRequestDurationTotal: meter.MustCounter(
			"http.server.request.duration_total",
			gmetric.MetricOption{
				Help:       "Total execution duration of request.",
				Unit:       "ms",
				Attributes: gmetric.Attributes{},
			},
		),
		HttpServerRequestBodySize: meter.MustCounter(
			"http.server.request.body_size",
			gmetric.MetricOption{
				Help:       "Incoming request bytes total.",
				Unit:       "bytes",
				Attributes: gmetric.Attributes{},
			},
		),
		HttpServerResponseBodySize: meter.MustCounter(
			"http.server.response.body_size",
			gmetric.MetricOption{
				Help:       "Response bytes total.",
				Unit:       "bytes",
				Attributes: gmetric.Attributes{},
			},
		),
	}
	return mm
}

func (m *localMetricManager) GetMetricOptionForRequestDurationByMap(attrMap gmetric.AttributeMap) gmetric.Option {
	return gmetric.Option{
		Attributes: attrMap.Pick(
			metricAttrKeyServerAddress,
			metricAttrKeyServerPort,
		),
	}
}

func (m *localMetricManager) GetMetricOptionForRequest(r *Request) gmetric.Option {
	attrMap := m.GetMetricAttributeMap(r)
	return m.GetMetricOptionForRequestByMap(attrMap)
}

func (m *localMetricManager) GetMetricOptionForRequestByMap(attrMap gmetric.AttributeMap) gmetric.Option {
	return gmetric.Option{
		Attributes: attrMap.Pick(
			metricAttrKeyServerAddress,
			metricAttrKeyServerPort,
			metricAttrKeyHttpRoute,
			metricAttrKeyUrlSchema,
			metricAttrKeyHttpRequestMethod,
			metricAttrKeyNetworkProtocolVersion,
		),
	}
}

func (m *localMetricManager) GetMetricOptionForResponseByMap(attrMap gmetric.AttributeMap) gmetric.Option {
	return gmetric.Option{
		Attributes: attrMap.Pick(
			metricAttrKeyServerAddress,
			metricAttrKeyServerPort,
			metricAttrKeyHttpRoute,
			metricAttrKeyUrlSchema,
			metricAttrKeyHttpRequestMethod,
			metricAttrKeyNetworkProtocolVersion,
			metricAttrKeyErrorCode,
			metricAttrKeyHttpResponseStatusCode,
		),
	}
}

func (m *localMetricManager) GetMetricAttributeMap(r *Request) gmetric.AttributeMap {
	var (
		serverAddress   string
		serverPort      string
		httpRoute       string
		protocolVersion string
		handler         = r.X取路由解析对象()
		localAddr       = r.Context别名().Value(http.LocalAddrContextKey)
		attrMap         = make(gmetric.AttributeMap)
	)
	serverAddress, serverPort = gstr.X分割2份(r.Host, ":")
	if localAddr != nil {
		_, serverPort = gstr.X分割2份(localAddr.(net.Addr).String(), ":")
	}
	if handler != nil && handler.Handler.X路由 != nil {
		httpRoute = handler.Handler.X路由.Uri
	} else {
		httpRoute = r.URL.Path
	}
	if array := gstr.X分割(r.Proto, "/"); len(array) > 1 {
		protocolVersion = array[1]
	}
	attrMap.Sets(gmetric.AttributeMap{
		metricAttrKeyServerAddress:          serverAddress,
		metricAttrKeyServerPort:             serverPort,
		metricAttrKeyHttpRoute:              httpRoute,
		metricAttrKeyUrlSchema:              r.GetSchema(),
		metricAttrKeyHttpRequestMethod:      r.Method,
		metricAttrKeyNetworkProtocolVersion: protocolVersion,
	})
	if r.LeaveTime != nil {
		var errCode int
		if err := r.X取错误信息(); err != nil {
			errCode = gerror.X取错误码(err).Code()
		}
		attrMap.Sets(gmetric.AttributeMap{
			metricAttrKeyErrorCode:              errCode,
			metricAttrKeyHttpResponseStatusCode: r.X响应.Status,
		})
	}
	return attrMap
}

func (s *X服务) handleMetricsBeforeRequest(r *Request) {
	if !gmetric.IsEnabled() {
		return
	}
	var (
		ctx           = r.Context别名()
		attrMap       = metricManager.GetMetricAttributeMap(r)
		requestOption = metricManager.GetMetricOptionForRequestByMap(attrMap)
	)
	metricManager.HttpServerRequestActive.Inc(
		ctx,
		requestOption,
	)
	metricManager.HttpServerRequestBodySize.Add(
		ctx,
		float64(r.ContentLength),
		requestOption,
	)
}

func (s *X服务) handleMetricsAfterRequestDone(r *Request) {
	if !gmetric.IsEnabled() {
		return
	}
	var (
		ctx             = r.Context别名()
		attrMap         = metricManager.GetMetricAttributeMap(r)
		durationMilli   = float64(r.LeaveTime.X取纳秒时长(r.EnterTime).Milliseconds())
		responseOption  = metricManager.GetMetricOptionForResponseByMap(attrMap)
		histogramOption = metricManager.GetMetricOptionForRequestDurationByMap(attrMap)
	)
	metricManager.HttpServerRequestTotal.Inc(ctx, responseOption)
	metricManager.HttpServerRequestActive.Dec(
		ctx,
		metricManager.GetMetricOptionForRequestByMap(attrMap),
	)
	metricManager.HttpServerResponseBodySize.Add(
		ctx,
		float64(r.X响应.BytesWritten()),
		responseOption,
	)
	metricManager.HttpServerRequestDurationTotal.Add(
		ctx,
		durationMilli,
		responseOption,
	)
	metricManager.HttpServerRequestDuration.Record(
		durationMilli,
		histogramOption,
	)
}
