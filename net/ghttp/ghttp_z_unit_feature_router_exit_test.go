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

func Test_Router_Exit(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定HookMap("/*", map[http类.Hook名称]http类.HandlerFunc{
		http类.HookBeforeServe:  func(r *http类.X请求) { r.X响应.X写响应缓冲区("1") },
		http类.HookAfterServe:   func(r *http类.X请求) { r.X响应.X写响应缓冲区("2") },
		http类.HookBeforeOutput: func(r *http类.X请求) { r.X响应.X写响应缓冲区("3") },
		http类.HookAfterOutput:  func(r *http类.X请求) { r.X响应.X写响应缓冲区("4") },
	})
	s.X绑定("/test/test", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("test-start")
		r.X退出当前()
		r.X响应.X写响应缓冲区("test-end")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "123")
		t.Assert(client.Get文本(ctx, "/test/test"), "1test-start23")
	})
}

func Test_Router_ExitHook(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/priority/show", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("show")
	})

	s.X绑定HookMap("/priority/:name", map[http类.Hook名称]http类.HandlerFunc{
		http类.HookBeforeServe: func(r *http类.X请求) {
			r.X响应.X写响应缓冲区("1")
		},
	})
	s.X绑定HookMap("/priority/*any", map[http类.Hook名称]http类.HandlerFunc{
		http类.HookBeforeServe: func(r *http类.X请求) {
			r.X响应.X写响应缓冲区("2")
		},
	})
	s.X绑定HookMap("/priority/show", map[http类.Hook名称]http类.HandlerFunc{
		http类.HookBeforeServe: func(r *http类.X请求) {
			r.X响应.X写响应缓冲区("3")
			r.X退出Hook()
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
		t.Assert(client.Get文本(ctx, "/priority/show"), "3show")
	})
}

func Test_Router_ExitAll(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/priority/show", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("show")
	})

	s.X绑定HookMap("/priority/:name", map[http类.Hook名称]http类.HandlerFunc{
		http类.HookBeforeServe: func(r *http类.X请求) {
			r.X响应.X写响应缓冲区("1")
		},
	})
	s.X绑定HookMap("/priority/*any", map[http类.Hook名称]http类.HandlerFunc{
		http类.HookBeforeServe: func(r *http类.X请求) {
			r.X响应.X写响应缓冲区("2")
		},
	})
	s.X绑定HookMap("/priority/show", map[http类.Hook名称]http类.HandlerFunc{
		http类.HookBeforeServe: func(r *http类.X请求) {
			r.X响应.X写响应缓冲区("3")
			r.X退出全部()
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
		t.Assert(client.Get文本(ctx, "/priority/show"), "3")
	})
}
