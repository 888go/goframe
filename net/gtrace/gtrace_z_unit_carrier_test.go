// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtrace_test

import (
	"context"
	"testing"
	
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	
	"github.com/888go/goframe/net/gtrace"
	"github.com/888go/goframe/test/gtest"
)

const (
	traceIDStr = "4bf92f3577b34da6a3ce929d0e0e4736"
	spanIDStr  = "00f067aa0ba902b7"
)

var (
	traceID = mustTraceIDFromHex(traceIDStr)
	spanID  = mustSpanIDFromHex(spanIDStr)
)

func mustTraceIDFromHex(s string) (t trace.TraceID) {
	var err error
	t, err = trace.TraceIDFromHex(s)
	if err != nil {
		panic(err)
	}
	return
}

func mustSpanIDFromHex(s string) (t trace.SpanID) {
	var err error
	t, err = trace.SpanIDFromHex(s)
	if err != nil {
		panic(err)
	}
	return
}

func TestNewCarrier(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		ctx := trace.ContextWithRemoteSpanContext(context.Background(), trace.NewSpanContext(trace.SpanContextConfig{
			TraceID:    traceID,
			SpanID:     spanID,
			TraceFlags: trace.FlagsSampled,
		}))
		sc := trace.SpanContextFromContext(ctx)
		t.Assert(sc.TraceID().String(), traceID.String())
		t.Assert(sc.SpanID().String(), "00f067aa0ba902b7")

		ctx, _ = otel.Tracer("").Start(ctx, "inject")
		carrier1 := gtrace.NewCarrier()
		otel.GetTextMapPropagator().Inject(ctx, carrier1)

		ctx = otel.GetTextMapPropagator().Extract(ctx, carrier1)
		gotSc := trace.SpanContextFromContext(ctx)
		t.Assert(gotSc.TraceID().String(), traceID.String())
		// 新建的span在内部创建，因此SpanID是不同的。
	})
}
