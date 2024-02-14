// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 静态服务测试。

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

func Test_StatusHandler(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		s.X绑定状态码中间件Map(map[int]http类.HandlerFunc{
			404: func(r *http类.X请求) { r.X响应.X写覆盖响应缓冲区("404") },
			502: func(r *http类.X请求) { r.X响应.X写覆盖响应缓冲区("502") },
		})
		s.X绑定("/502", func(r *http类.X请求) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		s.X绑定状态码中间件(502, func(r *http类.X请求) {
			r.X响应.X写覆盖响应缓冲区("1")
		})
		s.X绑定状态码中间件(502, func(r *http类.X请求) {
			r.X响应.X写响应缓冲区("2")
		})
		s.X绑定("/502", func(r *http类.X请求) {
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
