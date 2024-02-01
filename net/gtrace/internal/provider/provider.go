// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package provider
import (
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	)
type TracerProvider struct {
	*sdkTrace.TracerProvider
}

// New 函数返回一个新创建并已配置的 TracerProvider，该实例尚未关联 SpanProcessor。
//
// 默认情况下，返回的 TracerProvider 已按照以下配置进行设置：
// - 使用 ParentBased(AlwaysSample) 采样器；
// - 使用基于 Unix 纳秒时间戳和随机数生成的 IDGenerator；
// - 使用资源默认值 resource.Default()；
// - 使用默认的 SpanLimits。
//
// 通过传入的 opts 参数可用于覆盖这些默认值，并相应地配置返回的 TracerProvider。
func New() *TracerProvider {
	return &TracerProvider{
		TracerProvider: sdkTrace.NewTracerProvider(
			sdkTrace.WithIDGenerator(NewIDGenerator()),
		),
	}
}
