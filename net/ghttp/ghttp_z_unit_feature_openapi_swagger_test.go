// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
	gmeta "github.com/888go/goframe/util/gmeta"
	guid "github.com/888go/goframe/util/guid"
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
	s := g.Http类(guid.X生成())
	s.X设置APISwaggerUI路径("/swagger")
	s.X设置APIOpenApiUI路径("/api.json")
	s.Use别名(ghttp.MiddlewareHandlerResponse)
	s.X绑定("/test", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{
			Id:   1,
			Age:  req.Age,
			Name: req.Name,
		}, nil
	})
	s.X绑定("/test/error", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{
			Id:   1,
			Age:  req.Age,
			Name: req.Name,
		}, gerror.X创建("error")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/test?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18,"Name":"john"}}`)
		t.Assert(c.Get文本(ctx, "/test/error"), `{"code":50,"message":"error","data":{"Id":1,"Age":0,"Name":""}}`)

		t.Assert(gstr.X是否包含(c.Get文本(ctx, "/swagger/"), `API Reference`), true)
		t.Assert(gstr.X是否包含(c.Get文本(ctx, "/api.json"), `/test/error`), true)
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
	s := g.Http类(guid.X生成())
	s.X设置APISwaggerUI路径("/swagger")
	s.X设置APIOpenApiUI路径("/api.json")
	s.Use别名(ghttp.MiddlewareHandlerResponse)
	s.X绑定("/test", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{
			Id:   1,
			Age:  req.Age,
			Name: req.Name,
		}, nil
	})
	s.X绑定("/test/error", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{
			Id:   1,
			Age:  req.Age,
			Name: req.Name,
		}, gerror.X创建("error")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		openapi := s.X取OpenApi对象()
		t.AssertNE(openapi.Paths["/test"].Get, nil)
		t.AssertNE(openapi.Paths["/test"].Post, nil)
		t.AssertNE(openapi.Paths["/test/error"].Get, nil)
		t.AssertNE(openapi.Paths["/test/error"].Post, nil)

		t.Assert(len(openapi.Paths["/test"].Get.Parameters), 2)
		t.Assert(len(openapi.Paths["/test/error"].Get.Parameters), 2)
		t.Assert(len(openapi.Components.Schemas.Get(`github.com.888go.goframe.net.ghttp_test.TestReq`).Value.Properties.Map()), 2)

		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

				// 仅限于GET和POST方法。 md5:7c61bc58a2e1a657
		t.Assert(c.Get文本(ctx, "/test?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18,"Name":"john"}}`)
		t.Assert(c.Get文本(ctx, "/test/error"), `{"code":50,"message":"error","data":{"Id":1,"Age":0,"Name":""}}`)
		t.Assert(c.Post文本(ctx, "/test?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18,"Name":"john"}}`)
		t.Assert(c.Post文本(ctx, "/test/error"), `{"code":50,"message":"error","data":{"Id":1,"Age":0,"Name":""}}`)

				// 不适用于其他方法。 md5:4ae430a7ad6e3ac9
		t.Assert(c.Put文本(ctx, "/test?age=18&name=john"), `{"code":65,"message":"Not Found","data":null}`)
		t.Assert(c.Put文本(ctx, "/test/error"), `{"code":65,"message":"Not Found","data":null}`)

		t.Assert(gstr.X是否包含(c.Get文本(ctx, "/swagger/"), `API Reference`), true)
		t.Assert(gstr.X是否包含(c.Get文本(ctx, "/api.json"), `/test/error`), true)
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
	s := g.Http类(guid.X生成())
	s.X设置APISwaggerUI路径("/swagger")
	s.X设置APIOpenApiUI路径("/api.json")
	s.Use别名(ghttp.MiddlewareHandlerResponse)
	s.X绑定("/test", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{
			Id:   1,
			Age:  req.Age,
			Name: req.Name,
		}, nil
	})
	s.X绑定("/test/error", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{
			Id:   1,
			Age:  req.Age,
			Name: req.Name,
		}, gerror.X创建("error")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		openapi := s.X取OpenApi对象()
		t.AssertNE(openapi.Paths["/test"].Get, nil)
		t.AssertNE(openapi.Paths["/test"].Post, nil)
		t.AssertNE(openapi.Paths["/test"].Delete, nil)
		t.AssertNE(openapi.Paths["/test/error"].Get, nil)
		t.AssertNE(openapi.Paths["/test/error"].Post, nil)
		t.AssertNE(openapi.Paths["/test/error"].Delete, nil)

		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/test?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18,"Name":"john"}}`)
		t.Assert(c.Get文本(ctx, "/test/error"), `{"code":50,"message":"error","data":{"Id":1,"Age":0,"Name":""}}`)
		t.Assert(c.Post文本(ctx, "/test?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18,"Name":"john"}}`)
		t.Assert(c.Post文本(ctx, "/test/error"), `{"code":50,"message":"error","data":{"Id":1,"Age":0,"Name":""}}`)

		t.Assert(gstr.X是否包含(c.Get文本(ctx, "/swagger/"), `API Reference`), true)
		t.Assert(gstr.X是否包含(c.Get文本(ctx, "/api.json"), `/test/error`), true)
	})
}
