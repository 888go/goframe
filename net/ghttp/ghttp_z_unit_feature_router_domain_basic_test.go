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
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

func Test_Router_DomainBasic(t *testing.T) {
	s := g.Http类(uid类.X生成())
	d := s.X创建域名路由("localhost, local")
	d.X绑定("/:name", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("/:name")
	})
	d.X绑定("/:name/update", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区(r.Get别名("name"))
	})
	d.X绑定("/:name/:action", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区(r.Get别名("action"))
	})
	d.X绑定("/:name/*any", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区(r.Get别名("any"))
	})
	d.X绑定("/user/list/{field}.html", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区(r.Get别名("field"))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/john"), "Not Found")
		t.Assert(client.Get文本(ctx, "/john/update"), "Not Found")
		t.Assert(client.Get文本(ctx, "/john/edit"), "Not Found")
		t.Assert(client.Get文本(ctx, "/user/list/100.html"), "Not Found")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://localhost:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/john"), "")
		t.Assert(client.Get文本(ctx, "/john/update"), "john")
		t.Assert(client.Get文本(ctx, "/john/edit"), "edit")
		t.Assert(client.Get文本(ctx, "/user/list/100.html"), "100")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://local:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/john"), "")
		t.Assert(client.Get文本(ctx, "/john/update"), "john")
		t.Assert(client.Get文本(ctx, "/john/edit"), "edit")
		t.Assert(client.Get文本(ctx, "/user/list/100.html"), "100")
	})
}

func Test_Router_DomainMethod(t *testing.T) {
	s := g.Http类(uid类.X生成())
	d := s.X创建域名路由("localhost, local")
	d.X绑定("GET:/get", func(r *http类.X请求) {

	})
	d.X绑定("POST:/post", func(r *http类.X请求) {

	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		resp1, err := client.Get响应对象(ctx, "/get")
		t.AssertNil(err)
		defer resp1.X关闭()
		t.Assert(resp1.StatusCode, 404)

		resp2, err := client.Post响应对象(ctx, "/get")
		t.AssertNil(err)
		defer resp2.X关闭()
		t.Assert(resp2.StatusCode, 404)

		resp3, err := client.Get响应对象(ctx, "/post")
		t.AssertNil(err)
		defer resp3.X关闭()
		t.Assert(resp3.StatusCode, 404)

		resp4, err := client.Post响应对象(ctx, "/post")
		t.AssertNil(err)
		defer resp4.X关闭()
		t.Assert(resp4.StatusCode, 404)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://localhost:%d", s.X取已监听端口()))

		resp1, err := client.Get响应对象(ctx, "/get")
		t.AssertNil(err)
		defer resp1.X关闭()
		t.Assert(resp1.StatusCode, 200)

		resp2, err := client.Post响应对象(ctx, "/get")
		t.AssertNil(err)
		defer resp2.X关闭()
		t.Assert(resp2.StatusCode, 404)

		resp3, err := client.Get响应对象(ctx, "/post")
		t.AssertNil(err)
		defer resp3.X关闭()
		t.Assert(resp3.StatusCode, 404)

		resp4, err := client.Post响应对象(ctx, "/post")
		t.AssertNil(err)
		defer resp4.X关闭()
		t.Assert(resp4.StatusCode, 200)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://local:%d", s.X取已监听端口()))

		resp1, err := client.Get响应对象(ctx, "/get")
		t.AssertNil(err)
		defer resp1.X关闭()
		t.Assert(resp1.StatusCode, 200)

		resp2, err := client.Post响应对象(ctx, "/get")
		t.AssertNil(err)
		defer resp2.X关闭()
		t.Assert(resp2.StatusCode, 404)

		resp3, err := client.Get响应对象(ctx, "/post")
		t.AssertNil(err)
		defer resp3.X关闭()
		t.Assert(resp3.StatusCode, 404)

		resp4, err := client.Post响应对象(ctx, "/post")
		t.AssertNil(err)
		defer resp4.X关闭()
		t.Assert(resp4.StatusCode, 200)
	})
}

func Test_Router_DomainStatus(t *testing.T) {
	s := g.Http类(uid类.X生成())
	d := s.X创建域名路由("localhost, local")
	d.X绑定("/200", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区与HTTP状态码(200)
	})
	d.X绑定("/300", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区与HTTP状态码(300)
	})
	d.X绑定("/400", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区与HTTP状态码(400)
	})
	d.X绑定("/500", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区与HTTP状态码(500)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		resp1, err := client.Get响应对象(ctx, "/200")
		t.AssertNil(err)
		defer resp1.X关闭()
		t.Assert(resp1.StatusCode, 404)

		resp2, err := client.Get响应对象(ctx, "/300")
		t.AssertNil(err)
		defer resp2.X关闭()
		t.Assert(resp2.StatusCode, 404)

		resp3, err := client.Get响应对象(ctx, "/400")
		t.AssertNil(err)
		defer resp3.X关闭()
		t.Assert(resp3.StatusCode, 404)

		resp4, err := client.Get响应对象(ctx, "/500")
		t.AssertNil(err)
		defer resp4.X关闭()
		t.Assert(resp4.StatusCode, 404)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://localhost:%d", s.X取已监听端口()))

		resp1, err := client.Get响应对象(ctx, "/200")
		t.AssertNil(err)
		defer resp1.X关闭()
		t.Assert(resp1.StatusCode, 200)

		resp2, err := client.Get响应对象(ctx, "/300")
		t.AssertNil(err)
		defer resp2.X关闭()
		t.Assert(resp2.StatusCode, 300)

		resp3, err := client.Get响应对象(ctx, "/400")
		t.AssertNil(err)
		defer resp3.X关闭()
		t.Assert(resp3.StatusCode, 400)

		resp4, err := client.Get响应对象(ctx, "/500")
		t.AssertNil(err)
		defer resp4.X关闭()
		t.Assert(resp4.StatusCode, 500)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://local:%d", s.X取已监听端口()))

		resp1, err := client.Get响应对象(ctx, "/200")
		t.AssertNil(err)
		defer resp1.X关闭()
		t.Assert(resp1.StatusCode, 200)

		resp2, err := client.Get响应对象(ctx, "/300")
		t.AssertNil(err)
		defer resp2.X关闭()
		t.Assert(resp2.StatusCode, 300)

		resp3, err := client.Get响应对象(ctx, "/400")
		t.AssertNil(err)
		defer resp3.X关闭()
		t.Assert(resp3.StatusCode, 400)

		resp4, err := client.Get响应对象(ctx, "/500")
		t.AssertNil(err)
		defer resp4.X关闭()
		t.Assert(resp4.StatusCode, 500)
	})
}

func Test_Router_DomainCustomStatusHandler(t *testing.T) {
	s := g.Http类(uid类.X生成())
	d := s.X创建域名路由("localhost, local")
	d.X绑定("/", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("hello")
	})
	d.X绑定状态码中间件(404, func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("404 page")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/ThisDoesNotExist"), "Not Found")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://localhost:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "hello")
		t.Assert(client.Get文本(ctx, "/ThisDoesNotExist"), "404 page")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://local:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "hello")
		t.Assert(client.Get文本(ctx, "/ThisDoesNotExist"), "404 page")
	})
}

func Test_Router_Domain404(t *testing.T) {
	s := g.Http类(uid类.X生成())
	d := s.X创建域名路由("localhost, local")
	d.X绑定("/", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("hello")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://localhost:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "hello")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://local:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "hello")
	})
}

func Test_Router_DomainGroup(t *testing.T) {
	s := g.Http类(uid类.X生成())
	d := s.X创建域名路由("localhost, local")
	d.X创建分组路由("/", func(group *http类.X分组路由) {
		group.X创建分组路由("/app", func(group *http类.X分组路由) {
			group.X绑定GET("/{table}/list/{page}.html", func(r *http类.X请求) {
				intlog.Print(r.Context别名(), "/{table}/list/{page}.html")
				r.X响应.X写响应缓冲区(r.Get别名("table"), "&", r.Get别名("page"))
			})
			group.X绑定GET("/order/info/{order_id}", func(r *http类.X请求) {
				intlog.Print(r.Context别名(), "/order/info/{order_id}")
				r.X响应.X写响应缓冲区(r.Get别名("order_id"))
			})
			group.X绑定DELETE("/comment/{id}", func(r *http类.X请求) {
				intlog.Print(r.Context别名(), "/comment/{id}")
				r.X响应.X写响应缓冲区(r.Get别名("id"))
			})
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client1 := g.X网页类()
		client1.X设置url前缀(fmt.Sprintf("http://local:%d", s.X取已监听端口()))

		client2 := g.X网页类()
		client2.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client1.Get文本(ctx, "/app/t/list/2.html"), "t&2")
		t.Assert(client2.Get文本(ctx, "/app/t/list/2.html"), "Not Found")

		t.Assert(client1.Get文本(ctx, "/app/order/info/2"), "2")
		t.Assert(client2.Get文本(ctx, "/app/order/info/2"), "Not Found")

		t.Assert(client1.Get文本(ctx, "/app/comment/20"), "Not Found")
		t.Assert(client2.Get文本(ctx, "/app/comment/20"), "Not Found")

		t.Assert(client1.Delete文本(ctx, "/app/comment/20"), "20")
		t.Assert(client2.Delete文本(ctx, "/app/comment/20"), "Not Found")
	})
}
