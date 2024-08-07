// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 静态服务测试。 md5:2105c089651008de

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

func Test_StatusHandler(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Http类(guid.X生成())
		s.X绑定状态码中间件Map(map[int]ghttp.HandlerFunc{
			404: func(r *ghttp.Request) { r.X响应.X写覆盖响应缓冲区("404") },
			502: func(r *ghttp.Request) { r.X响应.X写覆盖响应缓冲区("502") },
		})
		s.X绑定("/502", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区与HTTP状态码并退出(502)
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/404"), "404")
		t.Assert(client.Get文本(ctx, "/502"), "502")
	})
}

func Test_StatusHandler_Multi(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Http类(guid.X生成())
		s.X绑定状态码中间件(502, func(r *ghttp.Request) {
			r.X响应.X写覆盖响应缓冲区("1")
		})
		s.X绑定状态码中间件(502, func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("2")
		})
		s.X绑定("/502", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区与HTTP状态码并退出(502)
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/502"), "12")
	})
}
