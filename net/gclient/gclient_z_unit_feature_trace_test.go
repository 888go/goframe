// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

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
	ghttp "github.com/888go/goframe/net/ghttp"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
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

	s := g.Server(guid.S())
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.WriteHeader(200)
		r.Response.WriteJson(g.Map{"field": "test_for_response_body"})
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		url := fmt.Sprintf("127.0.0.1:%d/hello", s.GetListenedPort())
		resp, err := c.DoRequest(ctx, http.MethodGet, url)
		t.AssertNil(err)
		t.AssertNE(resp, nil)
		t.Assert(resp.ReadAllString(), "{\"field\":\"test_for_response_body\"}")
	})
}
