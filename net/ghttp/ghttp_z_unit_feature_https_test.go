// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	_ "github.com/888go/goframe/net/ghttp/testdata/https/packed"

	"fmt"
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtcp "github.com/888go/goframe/net/gtcp"
	gfile "github.com/888go/goframe/os/gfile"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
	guid "github.com/888go/goframe/util/guid"
)

func Test_HTTPS_Basic(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定GET("/test", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("test")
		})
	})
	s.X启用HTTPS(
		gtest.DataPath("https", "files", "server.crt"),
		gtest.DataPath("https", "files", "server.key"),
	)
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	// HTTP
	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.AssertIN(gstr.X过滤首尾符并含空白(c.Get文本(ctx, "/")), g.Slice别名{"", "Client sent an HTTP request to an HTTPS server."})
		t.AssertIN(gstr.X过滤首尾符并含空白(c.Get文本(ctx, "/test")), g.Slice别名{"", "Client sent an HTTP request to an HTTPS server."})
	})
	// HTTPS
	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("https://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Get文本(ctx, "/"), "Not Found")
		t.Assert(c.Get文本(ctx, "/test"), "test")
	})
}

func Test_HTTPS_Resource(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定GET("/test", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("test")
		})
	})
	s.X启用HTTPS(
		gfile.X路径生成("files", "server.crt"),
		gfile.X路径生成("files", "server.key"),
	)
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	// HTTP
	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.AssertIN(gstr.X过滤首尾符并含空白(c.Get文本(ctx, "/")), g.Slice别名{"", "Client sent an HTTP request to an HTTPS server."})
		t.AssertIN(gstr.X过滤首尾符并含空白(c.Get文本(ctx, "/test")), g.Slice别名{"", "Client sent an HTTP request to an HTTPS server."})
	})
	// HTTPS
	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("https://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Get文本(ctx, "/"), "Not Found")
		t.Assert(c.Get文本(ctx, "/test"), "test")
	})
}

func Test_HTTPS_HTTP_Basic(t *testing.T) {
	var (
		portHttp, _  = gtcp.GetFreePort()
		portHttps, _ = gtcp.GetFreePort()
	)
	s := g.Http类(gtime.X取文本时间戳纳秒())
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定GET("/test", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("test")
		})
	})
	s.X启用HTTPS(
		gtest.DataPath("https", "files", "server.crt"),
		gtest.DataPath("https", "files", "server.key"),
	)
	s.X设置监听端口(portHttp)
	s.X设置HTTPS监听端口(portHttps)
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	// HTTP
	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", portHttp))
		t.Assert(c.Get文本(ctx, "/"), "Not Found")
		t.Assert(c.Get文本(ctx, "/test"), "test")
	})
	// HTTPS
	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("https://127.0.0.1:%d", portHttps))
		t.Assert(c.Get文本(ctx, "/"), "Not Found")
		t.Assert(c.Get文本(ctx, "/test"), "test")
	})
}
