// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp_test

import (
	"context"
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gmeta"
	"github.com/888go/goframe/util/guid"
)

func Test_OpenApi_Swagger(t *testing.T) {
	type TestReq struct {
		gmeta.Meta `method:"get" summary:"Test summary" tags:"Test"`
		Age        int
		Name       string
	}
	type TestRes struct {
		Id   int
		Age  int
		Name string
	}
	s := g.Server(guid.S())
	s.SetSwaggerPath("/swagger")
	s.SetOpenApiPath("/api.json")
	s.Use(ghttp.MiddlewareHandlerResponse)
	s.BindHandler("/test", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{
			Id:   1,
			Age:  req.Age,
			Name: req.Name,
		}, nil
	})
	s.BindHandler("/test/error", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{
			Id:   1,
			Age:  req.Age,
			Name: req.Name,
		}, gerror.New("error")
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(c.GetContent(ctx, "/test?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18,"Name":"john"}}`)
		t.Assert(c.GetContent(ctx, "/test/error"), `{"code":50,"message":"error","data":{"Id":1,"Age":0,"Name":""}}`)

		t.Assert(gstr.Contains(c.GetContent(ctx, "/swagger/"), `API Reference`), true)
		t.Assert(gstr.Contains(c.GetContent(ctx, "/api.json"), `/test/error`), true)
	})
}

func Test_OpenApi_Multiple_Methods_Swagger(t *testing.T) {
	type TestReq struct {
		gmeta.Meta `method:"get,post" summary:"Test summary" tags:"Test"`
		Age        int
		Name       string
	}
	type TestRes struct {
		Id   int
		Age  int
		Name string
	}
	s := g.Server(guid.S())
	s.SetSwaggerPath("/swagger")
	s.SetOpenApiPath("/api.json")
	s.Use(ghttp.MiddlewareHandlerResponse)
	s.BindHandler("/test", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{
			Id:   1,
			Age:  req.Age,
			Name: req.Name,
		}, nil
	})
	s.BindHandler("/test/error", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{
			Id:   1,
			Age:  req.Age,
			Name: req.Name,
		}, gerror.New("error")
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		openapi := s.GetOpenApi()
		t.AssertNE(openapi.Paths["/test"].Get, nil)
		t.AssertNE(openapi.Paths["/test"].Post, nil)
		t.AssertNE(openapi.Paths["/test/error"].Get, nil)
		t.AssertNE(openapi.Paths["/test/error"].Post, nil)

		t.Assert(len(openapi.Paths["/test"].Get.Parameters), 2)
		t.Assert(len(openapi.Paths["/test/error"].Get.Parameters), 2)
		t.Assert(len(openapi.Components.Schemas.Get(`github.com.888go.goframe.net.ghttp_test.TestReq`).Value.Properties.Map()), 2)

		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		// 只适用于GET和POST方法。
		t.Assert(c.GetContent(ctx, "/test?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18,"Name":"john"}}`)
		t.Assert(c.GetContent(ctx, "/test/error"), `{"code":50,"message":"error","data":{"Id":1,"Age":0,"Name":""}}`)
		t.Assert(c.PostContent(ctx, "/test?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18,"Name":"john"}}`)
		t.Assert(c.PostContent(ctx, "/test/error"), `{"code":50,"message":"error","data":{"Id":1,"Age":0,"Name":""}}`)

		// 对其他方法无效。
		t.Assert(c.PutContent(ctx, "/test?age=18&name=john"), `{"code":65,"message":"Not Found","data":null}`)
		t.Assert(c.PutContent(ctx, "/test/error"), `{"code":65,"message":"Not Found","data":null}`)

		t.Assert(gstr.Contains(c.GetContent(ctx, "/swagger/"), `API Reference`), true)
		t.Assert(gstr.Contains(c.GetContent(ctx, "/api.json"), `/test/error`), true)
	})
}

func Test_OpenApi_Method_All_Swagger(t *testing.T) {
	type TestReq struct {
		gmeta.Meta `method:"all" summary:"Test summary" tags:"Test"`
		Age        int
		Name       string
	}
	type TestRes struct {
		Id   int
		Age  int
		Name string
	}
	s := g.Server(guid.S())
	s.SetSwaggerPath("/swagger")
	s.SetOpenApiPath("/api.json")
	s.Use(ghttp.MiddlewareHandlerResponse)
	s.BindHandler("/test", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{
			Id:   1,
			Age:  req.Age,
			Name: req.Name,
		}, nil
	})
	s.BindHandler("/test/error", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{
			Id:   1,
			Age:  req.Age,
			Name: req.Name,
		}, gerror.New("error")
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		openapi := s.GetOpenApi()
		t.AssertNE(openapi.Paths["/test"].Get, nil)
		t.AssertNE(openapi.Paths["/test"].Post, nil)
		t.AssertNE(openapi.Paths["/test"].Delete, nil)
		t.AssertNE(openapi.Paths["/test/error"].Get, nil)
		t.AssertNE(openapi.Paths["/test/error"].Post, nil)
		t.AssertNE(openapi.Paths["/test/error"].Delete, nil)

		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(c.GetContent(ctx, "/test?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18,"Name":"john"}}`)
		t.Assert(c.GetContent(ctx, "/test/error"), `{"code":50,"message":"error","data":{"Id":1,"Age":0,"Name":""}}`)
		t.Assert(c.PostContent(ctx, "/test?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18,"Name":"john"}}`)
		t.Assert(c.PostContent(ctx, "/test/error"), `{"code":50,"message":"error","data":{"Id":1,"Age":0,"Name":""}}`)

		t.Assert(gstr.Contains(c.GetContent(ctx, "/swagger/"), `API Reference`), true)
		t.Assert(gstr.Contains(c.GetContent(ctx, "/api.json"), `/test/error`), true)
	})
}
