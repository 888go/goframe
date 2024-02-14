// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

import (
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

func Test_Middleware_CORS1(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/api.v2", func(group *http类.X分组路由) {
		group.X绑定中间件(MiddlewareCORS)
		group.X绑定POST("/user/list", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区("list")
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		// Common Checks.
		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/api.v2"), "Not Found")

		// GET 请求未匹配到任何路由。
		resp, err := client.Get响应对象(ctx, "/api.v2/user/list")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 0)
		t.Assert(resp.StatusCode, 404)
		resp.X关闭()

		// POST 请求与路由匹配，并且经过 CORS 中间件处理。
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/api.v2", func(group *http类.X分组路由) {
		group.X绑定中间件(MiddlewareCORS)
		group.X绑定GET("/user/list/{type}", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区(r.Get别名("type"))
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/api.v2", func(group *http类.X分组路由) {
		group.X绑定中间件(http类.X中间件跨域函数)
		group.X绑定GET("/user/list/{type}", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区(r.Get别名("type"))
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		client.X设置协议头("Access-Control-Request-Method", "POST")
		resp, err := client.Get响应对象(ctx, "/api.v2/user/list/1")
		t.AssertNil(err)
		resp.X关闭()
	})
}
