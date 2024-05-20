// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
)

func Test_Middleware_CORS1(t *testing.T) {
	s := g.Server(guid.S())
	s.Group("/api.v2", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.POST("/user/list", func(r *ghttp.Request) {
			r.Response.Write("list")
		})
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		// Common Checks.
		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/api.v2"), "Not Found")

		// GET请求没有指定任何路由。. md5:4bfb9b02af16686f
		resp, err := client.Get(ctx, "/api.v2/user/list")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 0)
		t.Assert(resp.StatusCode, 404)
		resp.Close()

		// POST请求与路由和CORS中间件匹配。. md5:b4101e02b3f9762b
		resp, err = client.Post(ctx, "/api.v2/user/list")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 1)
		t.Assert(resp.Header["Access-Control-Allow-Headers"][0], "Origin,Content-Type,Accept,User-Agent,Cookie,Authorization,X-Auth-Token,X-Requested-With")
		t.Assert(resp.Header["Access-Control-Allow-Methods"][0], "GET,PUT,POST,DELETE,PATCH,HEAD,CONNECT,OPTIONS,TRACE")
		t.Assert(resp.Header["Access-Control-Allow-Origin"][0], "*")
		t.Assert(resp.Header["Access-Control-Max-Age"][0], "3628800")
		resp.Close()
	})
	// OPTIONS GET
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		client.SetHeader("Access-Control-Request-Method", "GET")
		resp, err := client.Options(ctx, "/api.v2/user/list")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 0)
		t.Assert(resp.ReadAllString(), "Not Found")
		t.Assert(resp.StatusCode, 404)
		resp.Close()
	})
	// OPTIONS POST
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		client.SetHeader("Access-Control-Request-Method", "POST")
		resp, err := client.Options(ctx, "/api.v2/user/list")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 1)
		t.Assert(resp.StatusCode, 200)
		resp.Close()
	})
}

func Test_Middleware_CORS2(t *testing.T) {
	s := g.Server(guid.S())
	s.Group("/api.v2", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.GET("/user/list/{type}", func(r *ghttp.Request) {
			r.Response.Write(r.Get("type"))
		})
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		// Common Checks.
		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/api.v2"), "Not Found")
		// Get request.
		resp, err := client.Get(ctx, "/api.v2/user/list/1")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 1)
		t.Assert(resp.Header["Access-Control-Allow-Headers"][0], "Origin,Content-Type,Accept,User-Agent,Cookie,Authorization,X-Auth-Token,X-Requested-With")
		t.Assert(resp.Header["Access-Control-Allow-Methods"][0], "GET,PUT,POST,DELETE,PATCH,HEAD,CONNECT,OPTIONS,TRACE")
		t.Assert(resp.Header["Access-Control-Allow-Origin"][0], "*")
		t.Assert(resp.Header["Access-Control-Max-Age"][0], "3628800")
		t.Assert(resp.ReadAllString(), "1")
		resp.Close()
	})
	// OPTIONS GET None.
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		client.SetHeader("Access-Control-Request-Method", "GET")
		resp, err := client.Options(ctx, "/api.v2/user")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 0)
		t.Assert(resp.StatusCode, 404)
		resp.Close()
	})
	// OPTIONS GET
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		client.SetHeader("Access-Control-Request-Method", "GET")
		resp, err := client.Options(ctx, "/api.v2/user/list/1")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 1)
		t.Assert(resp.StatusCode, 200)
		resp.Close()
	})
	// OPTIONS POST
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		client.SetHeader("Access-Control-Request-Method", "POST")
		resp, err := client.Options(ctx, "/api.v2/user/list/1")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 0)
		t.Assert(resp.StatusCode, 404)
		resp.Close()
	})
}

func Test_Middleware_CORS3(t *testing.T) {
	s := g.Server(guid.S())
	s.Group("/api.v2", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareCORS)
		group.GET("/user/list/{type}", func(r *ghttp.Request) {
			r.Response.Write(r.Get("type"))
		})
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		client.SetHeader("Access-Control-Request-Method", "POST")
		resp, err := client.Get(ctx, "/api.v2/user/list/1")
		t.AssertNil(err)
		resp.Close()
	})
}
