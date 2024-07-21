// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package provider

import (
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
)

type TracerProvider struct {
	*sdkTrace.TracerProvider
}

		// New 返回一个新的并配置好的 TracerProvider，它没有 SpanProcessor。
		// 
		// 默认情况下，返回的 TracerProvider 配置如下：
		// - 一个基于父span（AlwaysSample）的采样器；
		// - 一个基于Unix纳秒时间戳和随机数的ID生成器；
		// - 资源.Default()资源；
		// - 默认的Span限制。
		// 
		// 传递给 opts 的参数将用于覆盖这些默认值，并适当地配置返回的 TracerProvider。
		// md5:92a14af244d0cf0e
func New() *TracerProvider {
	return &TracerProvider{
		TracerProvider: sdkTrace.NewTracerProvider(
			sdkTrace.WithIDGenerator(NewIDGenerator()),
		),
	}
}
