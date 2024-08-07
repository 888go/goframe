// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	garray "github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func Test_BindMiddleware_Basic1(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/test/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("test")
	})
	s.X绑定全局中间件("/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("1")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("2")
	}, func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("3")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("4")
	})
	s.X绑定全局中间件("/test/:name", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("5")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("6")
	}, func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("7")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("8")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/test"), "1342")
		t.Assert(client.Get文本(ctx, "/test/test"), "57test86")
	})
}

func Test_BindMiddleware_Basic2(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/test/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("test")
	})
	s.X绑定全局中间件("/*", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("1")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("2")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "12")
		t.Assert(client.Get文本(ctx, "/test"), "12")
		t.Assert(client.Get文本(ctx, "/test/test"), "1test2")
	})
}

func Test_BindMiddleware_Basic3(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/test/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("test")
	})
	s.X绑定全局中间件("PUT:/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("1")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("2")
	}, func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("3")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("4")
	})
	s.X绑定全局中间件("POST:/test/:name", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("5")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("6")
	}, func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("7")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("8")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
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
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定中间件(func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("1")
			r.X中间件管理器.Next()
		})
		group.X绑定中间件(func(r *ghttp.Request) {
			r.X中间件管理器.Next()
			r.X响应.X写响应缓冲区("2")
		})
		group.X绑定所有类型("/test", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("test")
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/test"), "1test2")
		t.Assert(client.Put文本(ctx, "/test/none"), "Not Found")
	})
}

func Test_Middleware_With_Static(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定中间件(func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("1")
			r.X中间件管理器.Next()
			r.X响应.X写响应缓冲区("2")
		})
		group.X绑定所有类型("/user/list", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("list")
		})
	})
	s.SetDumpRouterMap(false)
	s.X设置静态文件根目录(gtest.DataPath("static1"))
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "index")
		t.Assert(client.Get文本(ctx, "/test.html"), "test")
		t.Assert(client.Get文本(ctx, "/none"), "Not Found")
		t.Assert(client.Get文本(ctx, "/user/list"), "1list2")
	})
}

func Test_Middleware_Status(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定中间件(func(r *ghttp.Request) {
			r.X中间件管理器.Next()
			r.X响应.X写覆盖响应缓冲区(r.X响应.Status)
		})
		group.X绑定所有类型("/user/list", func(r *ghttp.Request) {
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

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/user/list"), "200")

		resp, err := client.Get响应对象(ctx, "/")
		t.AssertNil(err)
		defer resp.X关闭()
		t.Assert(resp.StatusCode, 404)
	})
}

func Test_Middleware_Hook_With_Static(t *testing.T) {
	s := g.Http类(guid.X生成())
	a := garray.X创建(true)
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定Hook("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
			a.Append别名(1)
			fmt.Println("HookBeforeServe")
			r.X响应.X写响应缓冲区("a")
		})
		group.X绑定Hook("/*", ghttp.HookAfterServe, func(r *ghttp.Request) {
			a.Append别名(1)
			fmt.Println("HookAfterServe")
			r.X响应.X写响应缓冲区("b")
		})
		group.X绑定中间件(func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("1")
			r.X中间件管理器.Next()
			r.X响应.X写响应缓冲区("2")
		})
		group.X绑定所有类型("/user/list", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("list")
		})
	})
	s.SetDumpRouterMap(false)
	s.X设置静态文件根目录(gtest.DataPath("static1"))
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

				// 长度断言有时会失败，所以为了调试目的我在这里添加了time.Sleep。 md5:9a9315edc424405c

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
	s := g.Http类(guid.X生成())
	s.X绑定("/test/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("test")
	})
	s.X绑定全局中间件("/test/*any", func(r *ghttp.Request) {
		r.X中间件管理器.Next()
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/test"), "Not Found")
		t.Assert(client.Get文本(ctx, "/test/test"), "test")
		t.Assert(client.Get文本(ctx, "/test/test/test"), "Not Found")
	})
}

func Test_BindMiddlewareDefault_Basic1(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/test/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("test")
	})
	s.X绑定全局默认中间件(func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("1")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("2")
	})
	s.X绑定全局默认中间件(func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("3")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("4")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "1342")
		t.Assert(client.Get文本(ctx, "/test/test"), "13test42")
	})
}

func Test_BindMiddlewareDefault_Basic2(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("PUT:/test/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("test")
	})
	s.X绑定全局默认中间件(func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("1")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("2")
	})
	s.X绑定全局默认中间件(func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("3")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("4")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "1342")
		t.Assert(client.Put文本(ctx, "/"), "1342")
		t.Assert(client.Get文本(ctx, "/test/test"), "1342")
		t.Assert(client.Put文本(ctx, "/test/test"), "13test42")
	})
}

func Test_BindMiddlewareDefault_Basic3(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/test/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("test")
	})
	s.X绑定全局默认中间件(func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("1")
		r.X中间件管理器.Next()
	})
	s.X绑定全局默认中间件(func(r *ghttp.Request) {
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("2")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "12")
		t.Assert(client.Get文本(ctx, "/test/test"), "1test2")
	})
}

func Test_BindMiddlewareDefault_Basic4(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/test/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("test")
	})
	s.X绑定全局默认中间件(func(r *ghttp.Request) {
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("1")
	})
	s.X绑定全局默认中间件(func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("2")
		r.X中间件管理器.Next()
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "21")
		t.Assert(client.Get文本(ctx, "/test/test"), "2test1")
	})
}

func Test_BindMiddlewareDefault_Basic5(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/test/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("test")
	})
	s.X绑定全局默认中间件(func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("1")
		r.X中间件管理器.Next()
	})
	s.X绑定全局默认中间件(func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("2")
		r.X中间件管理器.Next()
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "12")
		t.Assert(client.Get文本(ctx, "/test/test"), "12test")
	})
}

func Test_BindMiddlewareDefault_Status(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/test/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("test")
	})
	s.X绑定全局默认中间件(func(r *ghttp.Request) {
		r.X中间件管理器.Next()
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/test/test"), "test")
	})
}

type ObjectMiddleware struct{}

func (o *ObjectMiddleware) Init(r *ghttp.Request) {
	r.X响应.X写响应缓冲区("100")
}

func (o *ObjectMiddleware) Shut(r *ghttp.Request) {
	r.X响应.X写响应缓冲区("200")
}

func (o *ObjectMiddleware) Index(r *ghttp.Request) {
	r.X响应.X写响应缓冲区("Object Index")
}

func (o *ObjectMiddleware) Show(r *ghttp.Request) {
	r.X响应.X写响应缓冲区("Object Show")
}

func (o *ObjectMiddleware) Info(r *ghttp.Request) {
	r.X响应.X写响应缓冲区("Object Info")
}

func Test_BindMiddlewareDefault_Basic6(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定对象("/", new(ObjectMiddleware))
	s.X绑定全局默认中间件(func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("1")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("2")
	})
	s.X绑定全局默认中间件(func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("3")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("4")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
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
	s := g.Http类(guid.X生成())
	s.X绑定("/test/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("test")
	})
	s.X绑定Hook("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("a")
	})
	s.X绑定Hook("/*", ghttp.HookAfterServe, func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("b")
	})
	s.X绑定Hook("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("c")
	})
	s.X绑定Hook("/*", ghttp.HookAfterServe, func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("d")
	})
	s.X绑定全局默认中间件(func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("1")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("2")
	})
	s.X绑定全局默认中间件(func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("3")
		r.X中间件管理器.Next()
		r.X响应.X写响应缓冲区("4")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "ac1342bd")
		t.Assert(client.Get文本(ctx, "/test/test"), "ac13test42bd")
	})
}

func MiddlewareAuth(r *ghttp.Request) {
	token := r.Get别名("token").String()
	if token == "123456" {
		r.X中间件管理器.Next()
	} else {
		r.X响应.X写响应缓冲区与HTTP状态码(http.StatusForbidden)
	}
}

func MiddlewareCORS(r *ghttp.Request) {
	r.X响应.X跨域请求全允许()
	r.X中间件管理器.Next()
}

func Test_Middleware_CORSAndAuth(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.Use别名(MiddlewareCORS)
	s.X创建分组路由("/api.v2", func(group *ghttp.X分组路由) {
		group.X绑定中间件(MiddlewareAuth)
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
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.X设置协议头("Access-Control-Request-Headers", "GF,GoFrame").Get文本(ctx, "/"), "Not Found")
		t.Assert(client.X设置协议头("Origin", "GoFrame").Get文本(ctx, "/"), "Not Found")
	})
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.X设置协议头("Referer", "Referer").Post文本(ctx, "/"), "Not Found")
	})
}

func MiddlewareScope1(r *ghttp.Request) {
	r.X响应.X写响应缓冲区("a")
	r.X中间件管理器.Next()
	r.X响应.X写响应缓冲区("b")
}

func MiddlewareScope2(r *ghttp.Request) {
	r.X响应.X写响应缓冲区("c")
	r.X中间件管理器.Next()
	r.X响应.X写响应缓冲区("d")
}

func MiddlewareScope3(r *ghttp.Request) {
	r.X响应.X写响应缓冲区("e")
	r.X中间件管理器.Next()
	r.X响应.X写响应缓冲区("f")
}

func Test_Middleware_Scope(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定中间件(MiddlewareScope1)
		group.X绑定所有类型("/scope1", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("1")
		})
		group.X创建分组路由("/", func(group *ghttp.X分组路由) {
			group.X绑定中间件(MiddlewareScope2)
			group.X绑定所有类型("/scope2", func(r *ghttp.Request) {
				r.X响应.X写响应缓冲区("2")
			})
		})
		group.X创建分组路由("/", func(group *ghttp.X分组路由) {
			group.X绑定中间件(MiddlewareScope3)
			group.X绑定所有类型("/scope3", func(r *ghttp.Request) {
				r.X响应.X写响应缓冲区("3")
			})
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/scope1"), "a1b")
		t.Assert(client.Get文本(ctx, "/scope2"), "ac2db")
		t.Assert(client.Get文本(ctx, "/scope3"), "ae3fb")
	})
}

func Test_Middleware_Panic(t *testing.T) {
	s := g.Http类(guid.X生成())
	i := 0
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X创建分组路由("/", func(group *ghttp.X分组路由) {
			group.X绑定中间件(func(r *ghttp.Request) {
				i++
				panic("error")
				// r.Middleware.Next()
			}, func(r *ghttp.Request) {
				i++
				r.X中间件管理器.Next()
			})
			group.X绑定所有类型("/", func(r *ghttp.Request) {
				r.X响应.X写响应缓冲区(i)
			})
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "exception recovered: error")
	})
}

func Test_Middleware_JsonBody(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定中间件(ghttp.X中间件函数_验证JSON格式请求体)
		group.X绑定所有类型("/", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("hello")
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "hello")
		t.Assert(client.Put文本(ctx, "/"), "hello")
		t.Assert(client.Put文本(ctx, "/", `{"name":"john"}`), "hello")
		t.Assert(client.Put文本(ctx, "/", `{"name":}`), "the request body content should be JSON format")
	})
}

func Test_MiddlewareHandlerResponse(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定中间件(ghttp.MiddlewareHandlerResponse)
		group.X绑定GET("/403", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区与HTTP状态码(http.StatusForbidden, "")
		})
		group.X绑定GET("/default", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区与HTTP状态码(http.StatusInternalServerError, "")
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
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
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定GET("/default", func(r *ghttp.Request) {
			var buffer strings.Builder
			gzipWriter := gzip.NewWriter(&buffer)
			defer gzipWriter.Close()
			_, _ = gzipWriter.Write([]byte("hello"))
			// 设置响应头，表明内容使用 gzip 压缩
			r.X响应.Header().Set("Content-Encoding", "gzip")
			r.X响应.Header().Set("Content-Type", "text/plain")
			r.X响应.Header().Set("Content-Length", fmt.Sprint(buffer.Len()))
			// 写入压缩后的内容
			r.X响应.X写响应缓冲区(buffer.String())
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
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
