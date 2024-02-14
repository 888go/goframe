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

func Test_Router_Group_Hook1(t *testing.T) {
	s := g.Http类(uid类.X生成())
	group := s.X创建分组路由("/api")
	group.X绑定GET("/handler", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("1")
	})
	group.X绑定所有类型("/handler", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("0")
	}, http类.HookBeforeServe)
	group.X绑定所有类型("/handler", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("2")
	}, http类.HookAfterServe)

	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/api/handler"), "012")
		t.Assert(client.Post文本(ctx, "/api/handler"), "02")
		t.Assert(client.Get文本(ctx, "/api/ThisDoesNotExist"), "Not Found")
	})
}

func Test_Router_Group_Hook2(t *testing.T) {
	s := g.Http类(uid类.X生成())
	group := s.X创建分组路由("/api")
	group.X绑定GET("/handler", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("1")
	})
	group.X绑定GET("/*", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("0")
	}, http类.HookBeforeServe)
	group.X绑定GET("/*", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("2")
	}, http类.HookAfterServe)

	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/api/handler"), "012")
		t.Assert(client.Post文本(ctx, "/api/handler"), "Not Found")
		t.Assert(client.Get文本(ctx, "/api/ThisDoesNotExist"), "02")
		t.Assert(client.Post文本(ctx, "/api/ThisDoesNotExist"), "Not Found")
	})
}
