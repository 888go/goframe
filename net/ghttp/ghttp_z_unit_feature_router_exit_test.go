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

func Test_Router_Exit(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定HookMap("/*", map[ghttp.Hook名称]ghttp.HandlerFunc{
		ghttp.HookBeforeServe:  func(r *ghttp.Request) { r.X响应.X写响应缓冲区("1") },
		ghttp.HookAfterServe:   func(r *ghttp.Request) { r.X响应.X写响应缓冲区("2") },
		ghttp.HookBeforeOutput: func(r *ghttp.Request) { r.X响应.X写响应缓冲区("3") },
		ghttp.HookAfterOutput:  func(r *ghttp.Request) { r.X响应.X写响应缓冲区("4") },
	})
	s.X绑定("/test/test", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("test-start")
		r.X退出当前()
		r.X响应.X写响应缓冲区("test-end")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "123")
		t.Assert(client.Get文本(ctx, "/test/test"), "1test-start23")
	})
}

func Test_Router_ExitHook(t *testing.T) {
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
			r.X退出Hook()
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
		t.Assert(client.Get文本(ctx, "/priority/show"), "3show")
	})
}

func Test_Router_ExitAll(t *testing.T) {
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
			r.X退出全部()
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
		t.Assert(client.Get文本(ctx, "/priority/show"), "3")
	})
}
