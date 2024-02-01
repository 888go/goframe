// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtrace
import (
	"context"
	
	"go.opentelemetry.io/otel/trace"
	)
// Span 为兼容性和扩展性而封装了 trace.Span。
type Span struct {
	trace.Span
}

// NewSpan 使用默认的追踪器创建一个跨度。
func NewSpan(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, *Span) {
	ctx, span := NewTracer().Start(ctx, spanName, opts...)
	return ctx, &Span{
		Span: span,
	}
}
