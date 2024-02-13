// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

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
		元数据类.Meta `method:"get" summary:"Test summary" tags:"Test"`
		Age        int
		Name       string
	}
	type TestRes struct {
		Id   int
		Age  int
		Name string
	}
	s := g.Http类(uid类.X生成())
	s.X设置APISwaggerUI路径("/swagger")
	s.X设置APIOpenApiUI路径("/api.json")
	s.Use别名(http类.MiddlewareHandlerResponse)
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
		}, 错误类.X创建("error")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/test?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18,"Name":"john"}}`)
		t.Assert(c.Get文本(ctx, "/test/error"), `{"code":50,"message":"error","data":{"Id":1,"Age":0,"Name":""}}`)

		t.Assert(文本类.X是否包含(c.Get文本(ctx, "/swagger/"), `API Reference`), true)
		t.Assert(文本类.X是否包含(c.Get文本(ctx, "/api.json"), `/test/error`), true)
	})
}

func Test_OpenApi_Multiple_Methods_Swagger(t *testing.T) {
	type TestReq struct {
		元数据类.Meta `method:"get,post" summary:"Test summary" tags:"Test"`
		Age        int
		Name       string
	}
	type TestRes struct {
		Id   int
		Age  int
		Name string
	}
	s := g.Http类(uid类.X生成())
	s.X设置APISwaggerUI路径("/swagger")
	s.X设置APIOpenApiUI路径("/api.json")
	s.Use别名(http类.MiddlewareHandlerResponse)
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
		}, 错误类.X创建("error")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
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

		// 只适用于GET和POST方法。
		t.Assert(c.Get文本(ctx, "/test?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18,"Name":"john"}}`)
		t.Assert(c.Get文本(ctx, "/test/error"), `{"code":50,"message":"error","data":{"Id":1,"Age":0,"Name":""}}`)
		t.Assert(c.Post文本(ctx, "/test?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18,"Name":"john"}}`)
		t.Assert(c.Post文本(ctx, "/test/error"), `{"code":50,"message":"error","data":{"Id":1,"Age":0,"Name":""}}`)

		// 对其他方法无效。
		t.Assert(c.Put文本(ctx, "/test?age=18&name=john"), `{"code":65,"message":"Not Found","data":null}`)
		t.Assert(c.Put文本(ctx, "/test/error"), `{"code":65,"message":"Not Found","data":null}`)

		t.Assert(文本类.X是否包含(c.Get文本(ctx, "/swagger/"), `API Reference`), true)
		t.Assert(文本类.X是否包含(c.Get文本(ctx, "/api.json"), `/test/error`), true)
	})
}

func Test_OpenApi_Method_All_Swagger(t *testing.T) {
	type TestReq struct {
		元数据类.Meta `method:"all" summary:"Test summary" tags:"Test"`
		Age        int
		Name       string
	}
	type TestRes struct {
		Id   int
		Age  int
		Name string
	}
	s := g.Http类(uid类.X生成())
	s.X设置APISwaggerUI路径("/swagger")
	s.X设置APIOpenApiUI路径("/api.json")
	s.Use别名(http类.MiddlewareHandlerResponse)
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
		}, 错误类.X创建("error")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
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

		t.Assert(文本类.X是否包含(c.Get文本(ctx, "/swagger/"), `API Reference`), true)
		t.Assert(文本类.X是否包含(c.Get文本(ctx, "/api.json"), `/test/error`), true)
	})
}
