// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func Test_BindMiddleware_Basic1(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/test/test", func(r *http类.Request) {
		r.Response.X写响应缓冲区("test")
	})
	s.X绑定全局中间件("/test", func(r *http类.Request) {
		r.Response.X写响应缓冲区("1")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("2")
	}, func(r *http类.Request) {
		r.Response.X写响应缓冲区("3")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("4")
	})
	s.X绑定全局中间件("/test/:name", func(r *http类.Request) {
		r.Response.X写响应缓冲区("5")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("6")
	}, func(r *http类.Request) {
		r.Response.X写响应缓冲区("7")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("8")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/test"), "1342")
		t.Assert(client.Get文本(ctx, "/test/test"), "57test86")
	})
}

func Test_BindMiddleware_Basic2(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/test/test", func(r *http类.Request) {
		r.Response.X写响应缓冲区("test")
	})
	s.X绑定全局中间件("/*", func(r *http类.Request) {
		r.Response.X写响应缓冲区("1")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("2")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "12")
		t.Assert(client.Get文本(ctx, "/test"), "12")
		t.Assert(client.Get文本(ctx, "/test/test"), "1test2")
	})
}

func Test_BindMiddleware_Basic3(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/test/test", func(r *http类.Request) {
		r.Response.X写响应缓冲区("test")
	})
	s.X绑定全局中间件("PUT:/test", func(r *http类.Request) {
		r.Response.X写响应缓冲区("1")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("2")
	}, func(r *http类.Request) {
		r.Response.X写响应缓冲区("3")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("4")
	})
	s.X绑定全局中间件("POST:/test/:name", func(r *http类.Request) {
		r.Response.X写响应缓冲区("5")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("6")
	}, func(r *http类.Request) {
		r.Response.X写响应缓冲区("7")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("8")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/test"), "Not Found")
		t.Assert(client.Put文本(ctx, "/test"), "1342")
		t.Assert(client.Post文本(ctx, "/test"), "Not Found")
		t.Assert(client.Get文本(ctx, "/test/test"), "test")
		t.Assert(client.Put文本(ctx, "/test/test"), "test")
		t.Assert(client.Post文本(ctx, "/test/test"), "57test86")
	})
}

func Test_BindMiddleware_Basic4(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定中间件(func(r *http类.Request) {
			r.Response.X写响应缓冲区("1")
			r.Middleware.Next()
		})
		group.X绑定中间件(func(r *http类.Request) {
			r.Middleware.Next()
			r.Response.X写响应缓冲区("2")
		})
		group.X绑定所有类型("/test", func(r *http类.Request) {
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

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/test"), "1test2")
		t.Assert(client.Put文本(ctx, "/test/none"), "Not Found")
	})
}

func Test_Middleware_With_Static(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定中间件(func(r *http类.Request) {
			r.Response.X写响应缓冲区("1")
			r.Middleware.Next()
			r.Response.X写响应缓冲区("2")
		})
		group.X绑定所有类型("/user/list", func(r *http类.Request) {
			r.Response.X写响应缓冲区("list")
		})
	})
	s.SetDumpRouterMap(false)
	s.X设置静态文件根目录(单元测试类.DataPath("static1"))
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "index")
		t.Assert(client.Get文本(ctx, "/test.html"), "test")
		t.Assert(client.Get文本(ctx, "/none"), "Not Found")
		t.Assert(client.Get文本(ctx, "/user/list"), "1list2")
	})
}

func Test_Middleware_Status(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定中间件(func(r *http类.Request) {
			r.Middleware.Next()
			r.Response.X写覆盖响应缓冲区(r.Response.Status)
		})
		group.X绑定所有类型("/user/list", func(r *http类.Request) {
			r.Response.X写响应缓冲区("list")
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/user/list"), "200")

		resp, err := client.Get响应对象(ctx, "/")
		t.AssertNil(err)
		defer resp.X关闭()
		t.Assert(resp.StatusCode, 404)
	})
}

func Test_Middleware_Hook_With_Static(t *testing.T) {
	s := g.Http类(uid类.X生成())
	a := 数组类.X创建(true)
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定Hook("/*", http类.HookBeforeServe, func(r *http类.Request) {
			a.Append别名(1)
			fmt.Println("HookBeforeServe")
			r.Response.X写响应缓冲区("a")
		})
		group.X绑定Hook("/*", http类.HookAfterServe, func(r *http类.Request) {
			a.Append别名(1)
			fmt.Println("HookAfterServe")
			r.Response.X写响应缓冲区("b")
		})
		group.X绑定中间件(func(r *http类.Request) {
			r.Response.X写响应缓冲区("1")
			r.Middleware.Next()
			r.Response.X写响应缓冲区("2")
		})
		group.X绑定所有类型("/user/list", func(r *http类.Request) {
			r.Response.X写响应缓冲区("list")
		})
	})
	s.SetDumpRouterMap(false)
	s.X设置静态文件根目录(单元测试类.DataPath("static1"))
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		// 由于长度断言有时会失败，因此我在此为了调试目的添加了 time.Sleep。

		t.Assert(client.Get文本(ctx, "/"), "index")
		time.Sleep(100 * time.Millisecond)
		t.Assert(a.X取长度(), 2)

		t.Assert(client.Get文本(ctx, "/test.html"), "test")
		time.Sleep(100 * time.Millisecond)
		t.Assert(a.X取长度(), 4)

		t.Assert(client.Get文本(ctx, "/none"), "ab")
		time.Sleep(100 * time.Millisecond)
		t.Assert(a.X取长度(), 6)

		t.Assert(client.Get文本(ctx, "/user/list"), "a1list2b")
		time.Sleep(100 * time.Millisecond)
		t.Assert(a.X取长度(), 8)
	})
}

func Test_BindMiddleware_Status(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/test/test", func(r *http类.Request) {
		r.Response.X写响应缓冲区("test")
	})
	s.X绑定全局中间件("/test/*any", func(r *http类.Request) {
		r.Middleware.Next()
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/test"), "Not Found")
		t.Assert(client.Get文本(ctx, "/test/test"), "test")
		t.Assert(client.Get文本(ctx, "/test/test/test"), "Not Found")
	})
}

func Test_BindMiddlewareDefault_Basic1(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/test/test", func(r *http类.Request) {
		r.Response.X写响应缓冲区("test")
	})
	s.X绑定全局默认中间件(func(r *http类.Request) {
		r.Response.X写响应缓冲区("1")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("2")
	})
	s.X绑定全局默认中间件(func(r *http类.Request) {
		r.Response.X写响应缓冲区("3")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("4")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "1342")
		t.Assert(client.Get文本(ctx, "/test/test"), "13test42")
	})
}

func Test_BindMiddlewareDefault_Basic2(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("PUT:/test/test", func(r *http类.Request) {
		r.Response.X写响应缓冲区("test")
	})
	s.X绑定全局默认中间件(func(r *http类.Request) {
		r.Response.X写响应缓冲区("1")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("2")
	})
	s.X绑定全局默认中间件(func(r *http类.Request) {
		r.Response.X写响应缓冲区("3")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("4")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "1342")
		t.Assert(client.Put文本(ctx, "/"), "1342")
		t.Assert(client.Get文本(ctx, "/test/test"), "1342")
		t.Assert(client.Put文本(ctx, "/test/test"), "13test42")
	})
}

func Test_BindMiddlewareDefault_Basic3(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/test/test", func(r *http类.Request) {
		r.Response.X写响应缓冲区("test")
	})
	s.X绑定全局默认中间件(func(r *http类.Request) {
		r.Response.X写响应缓冲区("1")
		r.Middleware.Next()
	})
	s.X绑定全局默认中间件(func(r *http类.Request) {
		r.Middleware.Next()
		r.Response.X写响应缓冲区("2")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "12")
		t.Assert(client.Get文本(ctx, "/test/test"), "1test2")
	})
}

func Test_BindMiddlewareDefault_Basic4(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/test/test", func(r *http类.Request) {
		r.Response.X写响应缓冲区("test")
	})
	s.X绑定全局默认中间件(func(r *http类.Request) {
		r.Middleware.Next()
		r.Response.X写响应缓冲区("1")
	})
	s.X绑定全局默认中间件(func(r *http类.Request) {
		r.Response.X写响应缓冲区("2")
		r.Middleware.Next()
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "21")
		t.Assert(client.Get文本(ctx, "/test/test"), "2test1")
	})
}

func Test_BindMiddlewareDefault_Basic5(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/test/test", func(r *http类.Request) {
		r.Response.X写响应缓冲区("test")
	})
	s.X绑定全局默认中间件(func(r *http类.Request) {
		r.Response.X写响应缓冲区("1")
		r.Middleware.Next()
	})
	s.X绑定全局默认中间件(func(r *http类.Request) {
		r.Response.X写响应缓冲区("2")
		r.Middleware.Next()
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "12")
		t.Assert(client.Get文本(ctx, "/test/test"), "12test")
	})
}

func Test_BindMiddlewareDefault_Status(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/test/test", func(r *http类.Request) {
		r.Response.X写响应缓冲区("test")
	})
	s.X绑定全局默认中间件(func(r *http类.Request) {
		r.Middleware.Next()
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/test/test"), "test")
	})
}

type ObjectMiddleware struct{}

func (o *ObjectMiddleware) Init(r *http类.Request) {
	r.Response.X写响应缓冲区("100")
}

func (o *ObjectMiddleware) Shut(r *http类.Request) {
	r.Response.X写响应缓冲区("200")
}

func (o *ObjectMiddleware) Index(r *http类.Request) {
	r.Response.X写响应缓冲区("Object Index")
}

func (o *ObjectMiddleware) Show(r *http类.Request) {
	r.Response.X写响应缓冲区("Object Show")
}

func (o *ObjectMiddleware) Info(r *http类.Request) {
	r.Response.X写响应缓冲区("Object Info")
}

func Test_BindMiddlewareDefault_Basic6(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定对象("/", new(ObjectMiddleware))
	s.X绑定全局默认中间件(func(r *http类.Request) {
		r.Response.X写响应缓冲区("1")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("2")
	})
	s.X绑定全局默认中间件(func(r *http类.Request) {
		r.Response.X写响应缓冲区("3")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("4")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "13100Object Index20042")
		t.Assert(client.Get文本(ctx, "/init"), "1342")
		t.Assert(client.Get文本(ctx, "/shut"), "1342")
		t.Assert(client.Get文本(ctx, "/index"), "13100Object Index20042")
		t.Assert(client.Get文本(ctx, "/show"), "13100Object Show20042")
		t.Assert(client.Get文本(ctx, "/none-exist"), "1342")
	})
}

func Test_Hook_Middleware_Basic1(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/test/test", func(r *http类.Request) {
		r.Response.X写响应缓冲区("test")
	})
	s.X绑定Hook("/*", http类.HookBeforeServe, func(r *http类.Request) {
		r.Response.X写响应缓冲区("a")
	})
	s.X绑定Hook("/*", http类.HookAfterServe, func(r *http类.Request) {
		r.Response.X写响应缓冲区("b")
	})
	s.X绑定Hook("/*", http类.HookBeforeServe, func(r *http类.Request) {
		r.Response.X写响应缓冲区("c")
	})
	s.X绑定Hook("/*", http类.HookAfterServe, func(r *http类.Request) {
		r.Response.X写响应缓冲区("d")
	})
	s.X绑定全局默认中间件(func(r *http类.Request) {
		r.Response.X写响应缓冲区("1")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("2")
	})
	s.X绑定全局默认中间件(func(r *http类.Request) {
		r.Response.X写响应缓冲区("3")
		r.Middleware.Next()
		r.Response.X写响应缓冲区("4")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "ac1342bd")
		t.Assert(client.Get文本(ctx, "/test/test"), "ac13test42bd")
	})
}

func MiddlewareAuth(r *http类.Request) {
	token := r.Get别名("token").String()
	if token == "123456" {
		r.Middleware.Next()
	} else {
		r.Response.X写响应缓冲区与HTTP状态码(http.StatusForbidden)
	}
}

func MiddlewareCORS(r *http类.Request) {
	r.Response.X跨域请求全允许()
	r.Middleware.Next()
}

func Test_Middleware_CORSAndAuth(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.Use别名(MiddlewareCORS)
	s.X创建分组路由("/api.v2", func(group *http类.RouterGroup) {
		group.X绑定中间件(MiddlewareAuth)
		group.X绑定POST("/user/list", func(r *http类.Request) {
			r.Response.X写响应缓冲区("list")
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
		// Auth Checks.
		t.Assert(client.Post文本(ctx, "/api.v2/user/list"), "Forbidden")
		t.Assert(client.Post文本(ctx, "/api.v2/user/list", "token=123456"), "list")
		// CORS Checks.
		resp, err := client.Post响应对象(ctx, "/api.v2/user/list", "token=123456")
		t.AssertNil(err)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 1)
		t.Assert(resp.Header["Access-Control-Allow-Headers"][0], "Origin,Content-Type,Accept,User-Agent,Cookie,Authorization,X-Auth-Token,X-Requested-With")
		t.Assert(resp.Header["Access-Control-Allow-Methods"][0], "GET,PUT,POST,DELETE,PATCH,HEAD,CONNECT,OPTIONS,TRACE")
		t.Assert(resp.Header["Access-Control-Allow-Origin"][0], "*")
		t.Assert(resp.Header["Access-Control-Max-Age"][0], "3628800")
		resp.X关闭()
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.X设置协议头("Access-Control-Request-Headers", "GF,GoFrame").Get文本(ctx, "/"), "Not Found")
		t.Assert(client.X设置协议头("Origin", "GoFrame").Get文本(ctx, "/"), "Not Found")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.X设置协议头("Referer", "Referer").Post文本(ctx, "/"), "Not Found")
	})
}

func MiddlewareScope1(r *http类.Request) {
	r.Response.X写响应缓冲区("a")
	r.Middleware.Next()
	r.Response.X写响应缓冲区("b")
}

func MiddlewareScope2(r *http类.Request) {
	r.Response.X写响应缓冲区("c")
	r.Middleware.Next()
	r.Response.X写响应缓冲区("d")
}

func MiddlewareScope3(r *http类.Request) {
	r.Response.X写响应缓冲区("e")
	r.Middleware.Next()
	r.Response.X写响应缓冲区("f")
}

func Test_Middleware_Scope(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定中间件(MiddlewareScope1)
		group.X绑定所有类型("/scope1", func(r *http类.Request) {
			r.Response.X写响应缓冲区("1")
		})
		group.X创建分组路由("/", func(group *http类.RouterGroup) {
			group.X绑定中间件(MiddlewareScope2)
			group.X绑定所有类型("/scope2", func(r *http类.Request) {
				r.Response.X写响应缓冲区("2")
			})
		})
		group.X创建分组路由("/", func(group *http类.RouterGroup) {
			group.X绑定中间件(MiddlewareScope3)
			group.X绑定所有类型("/scope3", func(r *http类.Request) {
				r.Response.X写响应缓冲区("3")
			})
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/scope1"), "a1b")
		t.Assert(client.Get文本(ctx, "/scope2"), "ac2db")
		t.Assert(client.Get文本(ctx, "/scope3"), "ae3fb")
	})
}

func Test_Middleware_Panic(t *testing.T) {
	s := g.Http类(uid类.X生成())
	i := 0
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X创建分组路由("/", func(group *http类.RouterGroup) {
			group.X绑定中间件(func(r *http类.Request) {
				i++
				panic("error")
				// r.Middleware.Next() 表示调用下一个中间件。
			}, func(r *http类.Request) {
				i++
				r.Middleware.Next()
			})
			group.X绑定所有类型("/", func(r *http类.Request) {
				r.Response.X写响应缓冲区(i)
			})
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "exception recovered: error")
	})
}

func Test_Middleware_JsonBody(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定中间件(http类.X中间件函数_验证JSON格式请求体)
		group.X绑定所有类型("/", func(r *http类.Request) {
			r.Response.X写响应缓冲区("hello")
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "hello")
		t.Assert(client.Put文本(ctx, "/"), "hello")
		t.Assert(client.Put文本(ctx, "/", `{"name":"john"}`), "hello")
		t.Assert(client.Put文本(ctx, "/", `{"name":}`), "the request body content should be JSON format")
	})
}

func Test_MiddlewareHandlerResponse(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定中间件(http类.MiddlewareHandlerResponse)
		group.X绑定GET("/403", func(r *http类.Request) {
			r.Response.X写响应缓冲区与HTTP状态码(http.StatusForbidden, "")
		})
		group.X绑定GET("/default", func(r *http类.Request) {
			r.Response.X写响应缓冲区与HTTP状态码(http.StatusInternalServerError, "")
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		rsp, err := client.Get响应对象(ctx, "/403")
		t.AssertNil(err)
		t.Assert(rsp.StatusCode, http.StatusForbidden)
		rsp, err = client.Get响应对象(ctx, "/default")
		t.AssertNil(err)
		t.Assert(rsp.StatusCode, http.StatusInternalServerError)
	})
}

func Test_MiddlewareHandlerGzipResponse(t *testing.T) {
	tp := testTracerProvider{}
	otel.SetTracerProvider(&tp)
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定GET("/default", func(r *http类.Request) {
			var buffer strings.Builder
			gzipWriter := gzip.NewWriter(&buffer)
			defer gzipWriter.Close()
			_, _ = gzipWriter.Write([]byte("hello"))
			// 设置响应头，表明内容使用 gzip 压缩
			r.Response.Header().Set("Content-Encoding", "gzip")
			r.Response.Header().Set("Content-Type", "text/plain")
			r.Response.Header().Set("Content-Length", fmt.Sprint(buffer.Len()))
			// 写入压缩后的内容
			r.Response.X写响应缓冲区(buffer.String())
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		rsp, err := client.Get响应对象(ctx, "/default")
		t.AssertNil(err)
		t.Assert(rsp.StatusCode, http.StatusOK)
	})
}

type testTracerProvider struct{}

var _ trace.TracerProvider = &testTracerProvider{}

func (*testTracerProvider) Tracer(_ string, _ ...trace.TracerOption) trace.Tracer {
	return trace.NewNoopTracerProvider().Tracer("")
}
