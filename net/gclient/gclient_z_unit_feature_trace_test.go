// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 网页类_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
	
	"go.opentelemetry.io/otel"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/tracing"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

type CustomProvider struct {
	*sdkTrace.TracerProvider
}

func NewCustomProvider() *CustomProvider {
	return &CustomProvider{
		TracerProvider: sdkTrace.NewTracerProvider(
			sdkTrace.WithIDGenerator(NewCustomIDGenerator()),
		),
	}
}

type CustomIDGenerator struct{}

func NewCustomIDGenerator() *CustomIDGenerator {
	return &CustomIDGenerator{}
}

func (id *CustomIDGenerator) NewIDs(ctx context.Context) (traceID trace.TraceID, spanID trace.SpanID) {
	return tracing.NewIDs()
}

func (id *CustomIDGenerator) NewSpanID(ctx context.Context, traceID trace.TraceID) (spanID trace.SpanID) {
	return tracing.NewSpanID()
}

func TestClient_CustomProvider(t *testing.T) {
	provider := otel.GetTracerProvider()
	defer otel.SetTracerProvider(provider)

	otel.SetTracerProvider(NewCustomProvider())

	s := g.Http类(uid类.X生成())
	s.X绑定("/hello", func(r *http类.Request) {
		r.Response.WriteHeader(200)
		r.Response.X写响应缓冲区JSON(g.Map{"field": "test_for_response_body"})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		url := fmt.Sprintf("127.0.0.1:%d/hello", s.X取已监听端口())
		resp, err := c.X请求响应对象(ctx, http.MethodGet, url)
		t.AssertNil(err)
		t.AssertNE(resp, nil)
		t.Assert(resp.X取响应文本(), "{\"field\":\"test_for_response_body\"}")
	})
}
