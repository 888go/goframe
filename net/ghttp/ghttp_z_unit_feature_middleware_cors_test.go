// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp_test
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

		// GET 请求未匹配到任何路由。
		resp, err := client.Get(ctx, "/api.v2/user/list")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 0)
		t.Assert(resp.StatusCode, 404)
		resp.Close()

		// POST 请求与路由匹配，并且经过 CORS 中间件处理。
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
