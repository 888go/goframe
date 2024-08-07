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

func Test_Router_Group_Hook1(t *testing.T) {
	s := g.Http类(guid.X生成())
	group := s.X创建分组路由("/api")
	group.X绑定GET("/handler", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("1")
	})
	group.X绑定所有类型("/handler", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("0")
	}, ghttp.HookBeforeServe)
	group.X绑定所有类型("/handler", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("2")
	}, ghttp.HookAfterServe)

	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/api/handler"), "012")
		t.Assert(client.Post文本(ctx, "/api/handler"), "02")
		t.Assert(client.Get文本(ctx, "/api/ThisDoesNotExist"), "Not Found")
	})
}

func Test_Router_Group_Hook2(t *testing.T) {
	s := g.Http类(guid.X生成())
	group := s.X创建分组路由("/api")
	group.X绑定GET("/handler", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("1")
	})
	group.X绑定GET("/*", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("0")
	}, ghttp.HookBeforeServe)
	group.X绑定GET("/*", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("2")
	}, ghttp.HookAfterServe)

	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/api/handler"), "012")
		t.Assert(client.Post文本(ctx, "/api/handler"), "Not Found")
		t.Assert(client.Get文本(ctx, "/api/ThisDoesNotExist"), "02")
		t.Assert(client.Post文本(ctx, "/api/ThisDoesNotExist"), "Not Found")
	})
}
