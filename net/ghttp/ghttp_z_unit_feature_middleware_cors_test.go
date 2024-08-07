// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func Test_Middleware_CORS1(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/api.v2", func(group *ghttp.X分组路由) {
		group.X绑定中间件(MiddlewareCORS)
		group.X绑定POST("/user/list", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("list")
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		// Common Checks.
		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/api.v2"), "Not Found")

				// GET请求没有指定任何路由。 md5:4bfb9b02af16686f
		resp, err := client.Get响应对象(ctx, "/api.v2/user/list")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 0)
		t.Assert(resp.StatusCode, 404)
		resp.X关闭()

				// POST请求与路由和CORS中间件匹配。 md5:b4101e02b3f9762b
		resp, err = client.Post响应对象(ctx, "/api.v2/user/list")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 1)
		t.Assert(resp.Header["Access-Control-Allow-Headers"][0], "Origin,Content-Type,Accept,User-Agent,Cookie,Authorization,X-Auth-Token,X-Requested-With")
		t.Assert(resp.Header["Access-Control-Allow-Methods"][0], "GET,PUT,POST,DELETE,PATCH,HEAD,CONNECT,OPTIONS,TRACE")
		t.Assert(resp.Header["Access-Control-Allow-Origin"][0], "*")
		t.Assert(resp.Header["Access-Control-Max-Age"][0], "3628800")
		resp.X关闭()
	})
	// OPTIONS GET
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		client.X设置协议头("Access-Control-Request-Method", "GET")
		resp, err := client.Options响应对象(ctx, "/api.v2/user/list")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 0)
		t.Assert(resp.X取响应文本(), "Not Found")
		t.Assert(resp.StatusCode, 404)
		resp.X关闭()
	})
	// OPTIONS POST
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		client.X设置协议头("Access-Control-Request-Method", "POST")
		resp, err := client.Options响应对象(ctx, "/api.v2/user/list")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 1)
		t.Assert(resp.StatusCode, 200)
		resp.X关闭()
	})
}

func Test_Middleware_CORS2(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/api.v2", func(group *ghttp.X分组路由) {
		group.X绑定中间件(MiddlewareCORS)
		group.X绑定GET("/user/list/{type}", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区(r.Get别名("type"))
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		// Common Checks.
		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/api.v2"), "Not Found")
		// Get request.
		resp, err := client.Get响应对象(ctx, "/api.v2/user/list/1")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 1)
		t.Assert(resp.Header["Access-Control-Allow-Headers"][0], "Origin,Content-Type,Accept,User-Agent,Cookie,Authorization,X-Auth-Token,X-Requested-With")
		t.Assert(resp.Header["Access-Control-Allow-Methods"][0], "GET,PUT,POST,DELETE,PATCH,HEAD,CONNECT,OPTIONS,TRACE")
		t.Assert(resp.Header["Access-Control-Allow-Origin"][0], "*")
		t.Assert(resp.Header["Access-Control-Max-Age"][0], "3628800")
		t.Assert(resp.X取响应文本(), "1")
		resp.X关闭()
	})
	// OPTIONS GET None.
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		client.X设置协议头("Access-Control-Request-Method", "GET")
		resp, err := client.Options响应对象(ctx, "/api.v2/user")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 0)
		t.Assert(resp.StatusCode, 404)
		resp.X关闭()
	})
	// OPTIONS GET
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		client.X设置协议头("Access-Control-Request-Method", "GET")
		resp, err := client.Options响应对象(ctx, "/api.v2/user/list/1")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 1)
		t.Assert(resp.StatusCode, 200)
		resp.X关闭()
	})
	// OPTIONS POST
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		client.X设置协议头("Access-Control-Request-Method", "POST")
		resp, err := client.Options响应对象(ctx, "/api.v2/user/list/1")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 0)
		t.Assert(resp.StatusCode, 404)
		resp.X关闭()
	})
}

func Test_Middleware_CORS3(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/api.v2", func(group *ghttp.X分组路由) {
		group.X绑定中间件(ghttp.X中间件跨域函数)
		group.X绑定GET("/user/list/{type}", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区(r.Get别名("type"))
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		client.X设置协议头("Access-Control-Request-Method", "POST")
		resp, err := client.Get响应对象(ctx, "/api.v2/user/list/1")
		t.AssertNil(err)
		resp.X关闭()
	})
}
