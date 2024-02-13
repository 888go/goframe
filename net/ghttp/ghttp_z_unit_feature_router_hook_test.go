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

func Test_Router_Hook_Basic(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定HookMap("/*", map[http类.HookName]http类.HandlerFunc{
		http类.HookBeforeServe:  func(r *http类.Request) { r.Response.X写响应缓冲区("1") },
		http类.HookAfterServe:   func(r *http类.Request) { r.Response.X写响应缓冲区("2") },
		http类.HookBeforeOutput: func(r *http类.Request) { r.Response.X写响应缓冲区("3") },
		http类.HookAfterOutput:  func(r *http类.Request) { r.Response.X写响应缓冲区("4") },
	})
	s.X绑定("/test/test", func(r *http类.Request) {
		r.Response.X写响应缓冲区("test")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "123")
		t.Assert(client.Get文本(ctx, "/test/test"), "1test23")
	})
}

func Test_Router_Hook_Fuzzy_Router(t *testing.T) {
	s := g.Http类(uid类.X生成())
	i := 1000
	pattern1 := "/:name/info"
	s.X绑定HookMap(pattern1, map[http类.HookName]http类.HandlerFunc{
		http类.HookBeforeServe: func(r *http类.Request) {
			r.X设置自定义参数("uid", i)
			i++
		},
	})
	s.X绑定(pattern1, func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Get别名("uid"))
	})

	pattern2 := "/{object}/list/{page}.java"
	s.X绑定HookMap(pattern2, map[http类.HookName]http类.HandlerFunc{
		http类.HookBeforeOutput: func(r *http类.Request) {
			r.Response.X设置缓冲区字节集([]byte(
				fmt.Sprint(r.Get别名("object"), "&", r.Get别名("page"), "&", i),
			))
		},
	})
	s.X绑定(pattern2, func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Router.Uri)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/john"), "Not Found")
		t.Assert(client.Get文本(ctx, "/john/info"), "1000")
		t.Assert(client.Get文本(ctx, "/john/info"), "1001")
		t.Assert(client.Get文本(ctx, "/john/list/1.java"), "john&1&1002")
		t.Assert(client.Get文本(ctx, "/john/list/2.java"), "john&2&1002")
	})
}

func Test_Router_Hook_Priority(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/priority/show", func(r *http类.Request) {
		r.Response.X写响应缓冲区("show")
	})

	s.X绑定HookMap("/priority/:name", map[http类.HookName]http类.HandlerFunc{
		http类.HookBeforeServe: func(r *http类.Request) {
			r.Response.X写响应缓冲区("1")
		},
	})
	s.X绑定HookMap("/priority/*any", map[http类.HookName]http类.HandlerFunc{
		http类.HookBeforeServe: func(r *http类.Request) {
			r.Response.X写响应缓冲区("2")
		},
	})
	s.X绑定HookMap("/priority/show", map[http类.HookName]http类.HandlerFunc{
		http类.HookBeforeServe: func(r *http类.Request) {
			r.Response.X写响应缓冲区("3")
		},
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/priority/show"), "312show")
		t.Assert(client.Get文本(ctx, "/priority/any/any"), "2")
		t.Assert(client.Get文本(ctx, "/priority/name"), "12")
	})
}

func Test_Router_Hook_Multi(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/multi-hook", func(r *http类.Request) {
		r.Response.X写响应缓冲区("show")
	})

	s.X绑定HookMap("/multi-hook", map[http类.HookName]http类.HandlerFunc{
		http类.HookBeforeServe: func(r *http类.Request) {
			r.Response.X写响应缓冲区("1")
		},
	})
	s.X绑定HookMap("/multi-hook", map[http类.HookName]http类.HandlerFunc{
		http类.HookBeforeServe: func(r *http类.Request) {
			r.Response.X写响应缓冲区("2")
		},
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/multi-hook"), "12show")
	})
}

func Test_Router_Hook_ExitAll(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/test", func(r *http类.Request) {
		r.Response.X写响应缓冲区("test")
	})
	s.X创建分组路由("/hook", func(group *http类.RouterGroup) {
		group.X绑定中间件(func(r *http类.Request) {
			r.Response.X写响应缓冲区("1")
			r.Middleware.Next()
		})
		group.X绑定所有类型("/test", func(r *http类.Request) {
			r.Response.X写响应缓冲区("2")
		})
	})

	s.X绑定Hook("/hook/*", http类.HookBeforeServe, func(r *http类.Request) {
		r.Response.X写响应缓冲区("hook")
		r.X退出全部()
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/test"), "test")
		t.Assert(client.Get文本(ctx, "/hook/test"), "hook")
	})
}
