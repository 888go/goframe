// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package provider
import (
	"context"
	
	"go.opentelemetry.io/otel/trace"
	
	"github.com/888go/goframe/internal/tracing"
	)
// IDGenerator 是一个追踪 ID 生成器。
type IDGenerator struct{}

// NewIDGenerator 返回一个新的 IDGenerator。
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

// NewIDs 创建并返回一个新的跟踪ID和跨度ID。
func (id *IDGenerator) NewIDs(ctx context.Context) (traceID trace.TraceID, spanID trace.SpanID) {
	return tracing.NewIDs()
}

// NewSpanID 返回一个新的跟踪片段（span）的ID，该片段属于具有traceID的跟踪。
func (id *IDGenerator) NewSpanID(ctx context.Context, traceID trace.TraceID) (spanID trace.SpanID) {
	return tracing.NewSpanID()
}
