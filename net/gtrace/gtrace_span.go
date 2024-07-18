// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtrace

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

// Span 是为了兼容性和扩展性而包装 trace.Span。 md5:4d80699ca0a21c89
type Span struct {
	trace.Span
}

// NewSpan 使用默认追踪器创建一个span。 md5:0b09fd499d941eba
// ff:
// ctx:
// spanName:
// opts:
func NewSpan(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, *Span) {
	ctx, span := NewTracer().Start(ctx, spanName, opts...)
	return ctx, &Span{
		Span: span,
	}
}
