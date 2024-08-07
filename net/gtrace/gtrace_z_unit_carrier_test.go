// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtrace_test

import (
	"context"
	"testing"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/888go/goframe/net/gtrace"
	gtest "github.com/888go/goframe/test/gtest"
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
	gtest.C(t, func(t *gtest.T) {
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
				// 新的span在内部创建，所以SpanID会不同。 md5:91b3c3206e0ba72d
	})
}
