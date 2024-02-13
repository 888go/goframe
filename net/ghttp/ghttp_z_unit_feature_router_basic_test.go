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

func Test_Router_Basic1(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/:name", func(r *http类.Request) {
		r.Response.X写响应缓冲区("/:name")
	})
	s.X绑定("/:name/update", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Get别名("name"))
	})
	s.X绑定("/:name/:action", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Get别名("action"))
	})
	s.X绑定("/:name/*any", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Get别名("any"))
	})
	s.X绑定("/user/list/{field}.html", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Get别名("field"))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/john"), "")
		t.Assert(client.Get文本(ctx, "/john/update"), "john")
		t.Assert(client.Get文本(ctx, "/john/edit"), "edit")
		t.Assert(client.Get文本(ctx, "/user/list/100.html"), "100")
	})
}

func Test_Router_Basic2(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/{hash}", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Get别名("hash"))
	})
	s.X绑定("/{hash}.{type}", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Get别名("type"))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/data"), "data")
		t.Assert(client.Get文本(ctx, "/data.json"), "json")
	})
}

func Test_Router_Value(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取路由器Map副本()["hash"])
	})
	s.X绑定("/GetRouter", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取路由器值到泛型类("name", "john").String())
	})
	s.X绑定("/{hash}", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取路由器值到泛型类("hash").String())
	})
	s.X绑定("/{hash}.{type}", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取路由器值到泛型类("type").String())
	})
	s.X绑定("/{hash}.{type}.map", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取路由器Map副本()["type"])
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/"), "")
		t.Assert(client.Get文本(ctx, "/GetRouter"), "john")
		t.Assert(client.Get文本(ctx, "/data"), "data")
		t.Assert(client.Get文本(ctx, "/data.json"), "json")
		t.Assert(client.Get文本(ctx, "/data.json.map"), "json")
	})
}

// HTTP方法注册。
func Test_Router_Method(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("GET:/get", func(r *http类.Request) {

	})
	s.X绑定("POST:/post", func(r *http类.Request) {

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

// 路由器的额外字符 '/'
func Test_Router_ExtraChar(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/api", func(group *http类.RouterGroup) {
		group.X绑定GET("/test", func(r *http类.Request) {
			r.Response.X写响应缓冲区("test")
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/api/test"), "test")
		t.Assert(client.Get文本(ctx, "/api/test/"), "test")
		t.Assert(client.Get文本(ctx, "/api/test//"), "test")
		t.Assert(client.Get文本(ctx, "//api/test//"), "test")
		t.Assert(client.Get文本(ctx, "//api//test//"), "test")
		t.Assert(client.Get文本(ctx, "///api///test///"), "test")
	})
}

// 自定义状态处理器
func Test_Router_Status(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/200", func(r *http类.Request) {
		r.Response.X写响应缓冲区与HTTP状态码(200)
	})
	s.X绑定("/300", func(r *http类.Request) {
		r.Response.X写响应缓冲区与HTTP状态码(300)
	})
	s.X绑定("/400", func(r *http类.Request) {
		r.Response.X写响应缓冲区与HTTP状态码(400)
	})
	s.X绑定("/500", func(r *http类.Request) {
		r.Response.X写响应缓冲区与HTTP状态码(500)
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

		resp5, err := client.Get响应对象(ctx, "/404")
		t.AssertNil(err)
		defer resp5.X关闭()
		t.Assert(resp5.StatusCode, 404)
	})
}

func Test_Router_CustomStatusHandler(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/", func(r *http类.Request) {
		r.Response.X写响应缓冲区("hello")
	})
	s.X绑定状态码中间件(404, func(r *http类.Request) {
		r.Response.X写响应缓冲区("404 page")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "hello")
		resp, err := client.Get响应对象(ctx, "/ThisDoesNotExist")
		t.AssertNil(err)
		defer resp.X关闭()
		t.Assert(resp.StatusCode, 404)
		t.Assert(resp.X取响应文本(), "404 page")
	})
}

// 404未找到路由。
func Test_Router_404(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/", func(r *http类.Request) {
		r.Response.X写响应缓冲区("hello")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "hello")
		resp, err := client.Get响应对象(ctx, "/ThisDoesNotExist")
		t.AssertNil(err)
		defer resp.X关闭()
		t.Assert(resp.StatusCode, 404)
	})
}

func Test_Router_Priority(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/admin", func(r *http类.Request) {
		r.Response.X写响应缓冲区("admin")
	})
	s.X绑定("/admin-{page}", func(r *http类.Request) {
		r.Response.X写响应缓冲区("admin-{page}")
	})
	s.X绑定("/admin-goods", func(r *http类.Request) {
		r.Response.X写响应缓冲区("admin-goods")
	})
	s.X绑定("/admin-goods-{page}", func(r *http类.Request) {
		r.Response.X写响应缓冲区("admin-goods-{page}")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/admin"), "admin")
		t.Assert(client.Get文本(ctx, "/admin-1"), "admin-{page}")
		t.Assert(client.Get文本(ctx, "/admin-goods"), "admin-goods")
		t.Assert(client.Get文本(ctx, "/admin-goods-2"), "admin-goods-{page}")
	})
}
