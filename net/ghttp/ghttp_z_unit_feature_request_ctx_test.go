// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
	
	"github.com/888go/goframe/encoding/gbase64"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

func Test_Request_IsFileRequest(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		s.X创建分组路由("/", func(group *http类.RouterGroup) {
			group.X绑定所有类型("/", func(r *http类.Request) {
				r.Response.X写响应缓冲区(r.X是否为文件请求())
			})
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()

		time.Sleep(100 * time.Millisecond)

		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/"), false)
	})
}

func Test_Request_IsAjaxRequest(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		s.X创建分组路由("/", func(group *http类.RouterGroup) {
			group.X绑定所有类型("/", func(r *http类.Request) {
				r.Response.X写响应缓冲区(r.X是否为AJAX请求())
			})
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()

		time.Sleep(100 * time.Millisecond)

		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/"), false)
	})
}

func Test_Request_GetClientIp(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		s.X创建分组路由("/", func(group *http类.RouterGroup) {
			group.X绑定所有类型("/", func(r *http类.Request) {
				r.Response.X写响应缓冲区(r.X取客户端IP地址())
			})
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()

		time.Sleep(100 * time.Millisecond)

		c := g.X网页类()
		c.X设置协议头("X-Forwarded-For", "192.168.0.1")
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/"), "192.168.0.1")
	})
}

func Test_Request_GetUrl(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		s.X创建分组路由("/", func(group *http类.RouterGroup) {
			group.X绑定所有类型("/", func(r *http类.Request) {
				r.Response.X写响应缓冲区(r.X取URL())
			})
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()

		time.Sleep(100 * time.Millisecond)

		c := g.X网页类()
		c.X设置协议头("X-Forwarded-Proto", "https")
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/"), fmt.Sprintf("https://127.0.0.1:%d/", s.X取已监听端口()))
	})
}

func Test_Request_GetReferer(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		s.X创建分组路由("/", func(group *http类.RouterGroup) {
			group.X绑定所有类型("/", func(r *http类.Request) {
				r.Response.X写响应缓冲区(r.X取引用来源URL())
			})
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()

		time.Sleep(100 * time.Millisecond)

		c := g.X网页类()
		c.X设置协议头("Referer", "Referer")
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/"), "Referer")
	})
}

func Test_Request_GetServeHandler(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		s.X创建分组路由("/", func(group *http类.RouterGroup) {
			group.X绑定所有类型("/", func(r *http类.Request) {
				r.Response.X写响应缓冲区(r.X取路由解析对象() != nil)
			})
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()

		time.Sleep(100 * time.Millisecond)

		c := g.X网页类()
		c.X设置协议头("Referer", "Referer")
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/"), true)
	})
}

func Test_Request_BasicAuth(t *testing.T) {
	const (
		user      = "root"
		pass      = "123456"
		wrongPass = "12345"
	)

	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定所有类型("/auth1", func(r *http类.Request) {
			r.X账号密码认证(user, pass, "tips")
		})
		group.X绑定所有类型("/auth2", func(r *http类.Request) {
			r.X账号密码认证(user, pass)
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		rsp, err := c.Get响应对象(ctx, "/auth1")
		t.AssertNil(err)
		t.Assert(rsp.Header.Get("WWW-Authenticate"), "Basic realm=\"tips\"")
		t.Assert(rsp.StatusCode, http.StatusUnauthorized)

		rsp, err = c.X设置协议头("Authorization", user+pass).Get响应对象(ctx, "/auth1")
		t.AssertNil(err)
		t.Assert(rsp.StatusCode, http.StatusForbidden)

		rsp, err = c.X设置协议头("Authorization", "Test "+user+pass).Get响应对象(ctx, "/auth1")
		t.AssertNil(err)
		t.Assert(rsp.StatusCode, http.StatusForbidden)

		rsp, err = c.X设置协议头("Authorization", "Basic "+user+pass).Get响应对象(ctx, "/auth1")
		t.AssertNil(err)
		t.Assert(rsp.StatusCode, http.StatusForbidden)

		rsp, err = c.X设置协议头("Authorization", "Basic "+编码base64类.X文本编码(user+pass)).Get响应对象(ctx, "/auth1")
		t.AssertNil(err)
		t.Assert(rsp.StatusCode, http.StatusForbidden)

		rsp, err = c.X设置协议头("Authorization", "Basic "+编码base64类.X文本编码(user+":"+wrongPass)).Get响应对象(ctx, "/auth1")
		t.AssertNil(err)
		t.Assert(rsp.StatusCode, http.StatusUnauthorized)

		rsp, err = c.X账号密码(user, pass).Get响应对象(ctx, "/auth1")
		t.AssertNil(err)
		t.Assert(rsp.StatusCode, http.StatusOK)

		rsp, err = c.Get响应对象(ctx, "/auth2")
		t.AssertNil(err)
		t.Assert(rsp.Header.Get("WWW-Authenticate"), "Basic realm=\"Need Login\"")
		t.Assert(rsp.StatusCode, http.StatusUnauthorized)
	})
}

func Test_Request_SetCtx(t *testing.T) {
	type ctxKey string
	const testkey ctxKey = "test"
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定中间件(func(r *http类.Request) {
			ctx := context.WithValue(r.Context别名(), testkey, 1)
			r.X设置上下文对象(ctx)
			r.Middleware.Next()
		})
		group.X绑定所有类型("/", func(r *http类.Request) {
			r.Response.X写响应缓冲区(r.Context别名().Value(testkey))
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/"), "1")
	})
}

func Test_Request_GetCtx(t *testing.T) {
	type ctxKey string
	const testkey ctxKey = "test"
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定中间件(func(r *http类.Request) {
			ctx := context.WithValue(r.X取上下文对象(), testkey, 1)
			r.X设置上下文对象(ctx)
			r.Middleware.Next()
		})
		group.X绑定所有类型("/", func(r *http类.Request) {
			r.Response.X写响应缓冲区(r.Context别名().Value(testkey))
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/"), "1")
	})
}

func Test_Request_GetCtxVar(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定中间件(func(r *http类.Request) {
			r.Middleware.Next()
		})
		group.X绑定GET("/", func(r *http类.Request) {
			r.Response.X写响应缓冲区(r.X取上下文对象值("key", "val"))
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "val")
	})
}

func Test_Request_Form(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定所有类型("/", func(r *http类.Request) {
			r.X设置表单值("key", "val")
			r.Response.X写响应缓冲区(r.X取表单值到泛型类("key"))
		})
		group.X绑定所有类型("/useDef", func(r *http类.Request) {
			r.Response.X写响应缓冲区(r.X取表单值到泛型类("key", "defVal"))
		})
		group.X绑定所有类型("/GetFormMap", func(r *http类.Request) {
			r.Response.X写响应缓冲区(r.X取表单值到Map(map[string]interface{}{"key": "val"}))
		})
		group.X绑定所有类型("/GetFormMap1", func(r *http类.Request) {
			r.Response.X写响应缓冲区(r.X取表单值到Map(map[string]interface{}{"array": "val"}))
		})
		group.X绑定所有类型("/GetFormMapStrVar", func(r *http类.Request) {
			if r.Get别名("a") != nil {
				r.Response.X写响应缓冲区(r.X取表单值到Map泛型类()["a"])
			}
		})
		group.X绑定所有类型("/GetFormStruct", func(r *http类.Request) {
			var user User
			if err := r.X取表单值到结构(&user); err != nil {
				r.Response.X写响应缓冲区(err.Error())
			} else {
				r.Response.X写响应缓冲区(user.Name)
			}
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "val")
		t.Assert(client.Get文本(ctx, "/useDef"), "defVal")
		t.Assert(client.Post文本(ctx, "/GetFormMap"), "{\"key\":\"val\"}")
		t.Assert(client.Post文本(ctx, "/GetFormMap", "array[]=1&array[]=2"), "{\"key\":\"val\"}")
		t.Assert(client.Post文本(ctx, "/GetFormMap1", "array[]=1&array[]=2"), "{\"array\":[\"1\",\"2\"]}")
		t.Assert(client.Get文本(ctx, "/GetFormMapStrVar", "a=1&b=2"), nil)
		t.Assert(client.Post文本(ctx, "/GetFormMapStrVar", "a=1&b=2"), `1`)
		t.Assert(client.Post文本(ctx, "/GetFormStruct", g.Map{
			"id":   1,
			"name": "john",
		}), "john")
	})
}
