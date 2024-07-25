// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package provider

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/gogf/gf/v2/internal/tracing"
)

// IDGenerator 是一个追踪ID生成器。 md5:72df652f0bf82307
type IDGenerator struct{}

// NewIDGenerator 返回一个新的 IDGenerator。 md5:f4445c039fa564db
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

// NewIDs 生成并返回一个新的 trace 和 span ID。 md5:1b5d0dc93c6f38ff
func (id *IDGenerator) NewIDs(ctx context.Context) (traceID trace.TraceID, spanID trace.SpanID) {
	return tracing.NewIDs()
}

// NewSpanID 返回一个新的span在trace中的ID，带有traceID。 md5:1fd29c38ad04372b
func (id *IDGenerator) NewSpanID(ctx context.Context, traceID trace.TraceID) (spanID trace.SpanID) {
	return tracing.NewSpanID()
}
