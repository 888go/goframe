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

func Test_Router_Group_Group(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/api.v2", func(group *ghttp.X分组路由) {
		group.X绑定中间件(func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("1")
			r.X中间件管理器.Next()
			r.X响应.X写响应缓冲区("2")
		})
		group.X绑定GET("/test", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("test")
		})
		group.X创建分组路由("/order", func(group *ghttp.X分组路由) {
			group.X绑定GET("/list", func(r *ghttp.Request) {
				r.X响应.X写响应缓冲区("list")
			})
			group.X绑定PUT("/update", func(r *ghttp.Request) {
				r.X响应.X写响应缓冲区("update")
			})
		})
		group.X创建分组路由("/user", func(group *ghttp.X分组路由) {
			group.X绑定GET("/info", func(r *ghttp.Request) {
				r.X响应.X写响应缓冲区("info")
			})
			group.X绑定POST("/edit", func(r *ghttp.Request) {
				r.X响应.X写响应缓冲区("edit")
			})
			group.X绑定DELETE("/drop", func(r *ghttp.Request) {
				r.X响应.X写响应缓冲区("drop")
			})
		})
		group.X创建分组路由("/hook", func(group *ghttp.X分组路由) {
			group.X绑定Hook("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
				r.X响应.X写响应缓冲区("hook any")
			})
			group.X绑定Hook("/:name", ghttp.HookBeforeServe, func(r *ghttp.Request) {
				r.X响应.X写响应缓冲区("hook name")
			})
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
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
