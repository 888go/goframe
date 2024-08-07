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

func Test_Router_Hook_Basic(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定HookMap("/*", map[ghttp.Hook名称]ghttp.HandlerFunc{
		ghttp.HookBeforeServe:  func(r *ghttp.Request) { r.X响应.X写响应缓冲区("1") },
		ghttp.HookAfterServe:   func(r *ghttp.Request) { r.X响应.X写响应缓冲区("2") },
		ghttp.HookBeforeOutput: func(r *ghttp.Request) { r.X响应.X写响应缓冲区("3") },
		ghttp.HookAfterOutput:  func(r *ghttp.Request) { r.X响应.X写响应缓冲区("4") },
	})
	s.X绑定("/test/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("test")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "123")
		t.Assert(client.Get文本(ctx, "/test/test"), "1test23")
	})
}

func Test_Router_Hook_Fuzzy_Router(t *testing.T) {
	s := g.Http类(guid.X生成())
	i := 1000
	pattern1 := "/:name/info"
	s.X绑定HookMap(pattern1, map[ghttp.Hook名称]ghttp.HandlerFunc{
		ghttp.HookBeforeServe: func(r *ghttp.Request) {
			r.X设置自定义参数("uid", i)
			i++
		},
	})
	s.X绑定(pattern1, func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区(r.Get别名("uid"))
	})

	pattern2 := "/{object}/list/{page}.java"
	s.X绑定HookMap(pattern2, map[ghttp.Hook名称]ghttp.HandlerFunc{
		ghttp.HookBeforeOutput: func(r *ghttp.Request) {
			r.X响应.SetBuffer([]byte(
				fmt.Sprint(r.Get别名("object"), "&", r.Get别名("page"), "&", i),
			))
		},
	})
	s.X绑定(pattern2, func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区(r.X路由.Uri)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
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
	s := g.Http类(guid.X生成())
	s.X绑定("/priority/show", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("show")
	})

	s.X绑定HookMap("/priority/:name", map[ghttp.Hook名称]ghttp.HandlerFunc{
		ghttp.HookBeforeServe: func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("1")
		},
	})
	s.X绑定HookMap("/priority/*any", map[ghttp.Hook名称]ghttp.HandlerFunc{
		ghttp.HookBeforeServe: func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("2")
		},
	})
	s.X绑定HookMap("/priority/show", map[ghttp.Hook名称]ghttp.HandlerFunc{
		ghttp.HookBeforeServe: func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("3")
		},
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/priority/show"), "312show")
		t.Assert(client.Get文本(ctx, "/priority/any/any"), "2")
		t.Assert(client.Get文本(ctx, "/priority/name"), "12")
	})
}

func Test_Router_Hook_Multi(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/multi-hook", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("show")
	})

	s.X绑定HookMap("/multi-hook", map[ghttp.Hook名称]ghttp.HandlerFunc{
		ghttp.HookBeforeServe: func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("1")
		},
	})
	s.X绑定HookMap("/multi-hook", map[ghttp.Hook名称]ghttp.HandlerFunc{
		ghttp.HookBeforeServe: func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("2")
		},
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/multi-hook"), "12show")
	})
}

func Test_Router_Hook_ExitAll(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("test")
	})
	s.X创建分组路由("/hook", func(group *ghttp.X分组路由) {
		group.X绑定中间件(func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("1")
			r.X中间件管理器.Next()
		})
		group.X绑定所有类型("/test", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("2")
		})
	})

	s.X绑定Hook("/hook/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("hook")
		r.X退出全部()
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/test"), "test")
		t.Assert(client.Get文本(ctx, "/hook/test"), "hook")
	})
}
