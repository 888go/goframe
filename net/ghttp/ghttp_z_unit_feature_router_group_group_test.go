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

func Test_Router_Group_Group(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/api.v2", func(group *http类.X分组路由) {
		group.X绑定中间件(func(r *http类.X请求) {
			r.X响应.X写响应缓冲区("1")
			r.X中间件管理器.Next()
			r.X响应.X写响应缓冲区("2")
		})
		group.X绑定GET("/test", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区("test")
		})
		group.X创建分组路由("/order", func(group *http类.X分组路由) {
			group.X绑定GET("/list", func(r *http类.X请求) {
				r.X响应.X写响应缓冲区("list")
			})
			group.X绑定PUT("/update", func(r *http类.X请求) {
				r.X响应.X写响应缓冲区("update")
			})
		})
		group.X创建分组路由("/user", func(group *http类.X分组路由) {
			group.X绑定GET("/info", func(r *http类.X请求) {
				r.X响应.X写响应缓冲区("info")
			})
			group.X绑定POST("/edit", func(r *http类.X请求) {
				r.X响应.X写响应缓冲区("edit")
			})
			group.X绑定DELETE("/drop", func(r *http类.X请求) {
				r.X响应.X写响应缓冲区("drop")
			})
		})
		group.X创建分组路由("/hook", func(group *http类.X分组路由) {
			group.X绑定Hook("/*", http类.HookBeforeServe, func(r *http类.X请求) {
				r.X响应.X写响应缓冲区("hook any")
			})
			group.X绑定Hook("/:name", http类.HookBeforeServe, func(r *http类.X请求) {
				r.X响应.X写响应缓冲区("hook name")
			})
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/api.v2"), "Not Found")
		t.Assert(client.Get文本(ctx, "/api.v2/test"), "1test2")
		t.Assert(client.Get文本(ctx, "/api.v2/hook"), "hook any")
		t.Assert(client.Get文本(ctx, "/api.v2/hook/name"), "hook namehook any")
		t.Assert(client.Get文本(ctx, "/api.v2/hook/name/any"), "hook any")
		t.Assert(client.Get文本(ctx, "/api.v2/order/list"), "1list2")
		t.Assert(client.Get文本(ctx, "/api.v2/order/update"), "Not Found")
		t.Assert(client.Put文本(ctx, "/api.v2/order/update"), "1update2")
		t.Assert(client.Get文本(ctx, "/api.v2/user/drop"), "Not Found")
		t.Assert(client.Delete文本(ctx, "/api.v2/user/drop"), "1drop2")
		t.Assert(client.Get文本(ctx, "/api.v2/user/edit"), "Not Found")
		t.Assert(client.Post文本(ctx, "/api.v2/user/edit"), "1edit2")
		t.Assert(client.Get文本(ctx, "/api.v2/user/info"), "1info2")
	})
}
