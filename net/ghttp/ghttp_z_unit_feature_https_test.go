// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

import (
	_ "github.com/888go/goframe/net/ghttp/testdata/https/packed"
	
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/net/gtcp"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/guid"
)

func Test_HTTPS_Basic(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定GET("/test", func(r *http类.Request) {
			r.Response.X写响应缓冲区("test")
		})
	})
	s.X启用HTTPS(
		单元测试类.DataPath("https", "files", "server.crt"),
		单元测试类.DataPath("https", "files", "server.key"),
	)
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	// HTTP
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.AssertIN(文本类.X过滤首尾符并含空白(c.Get文本(ctx, "/")), g.Slice别名{"", "Client sent an HTTP request to an HTTPS server."})
		t.AssertIN(文本类.X过滤首尾符并含空白(c.Get文本(ctx, "/test")), g.Slice别名{"", "Client sent an HTTP request to an HTTPS server."})
	})
	// HTTPS
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("https://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Get文本(ctx, "/"), "Not Found")
		t.Assert(c.Get文本(ctx, "/test"), "test")
	})
}

func Test_HTTPS_Resource(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定GET("/test", func(r *http类.Request) {
			r.Response.X写响应缓冲区("test")
		})
	})
	s.X启用HTTPS(
		文件类.X路径生成("files", "server.crt"),
		文件类.X路径生成("files", "server.key"),
	)
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	// HTTP
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.AssertIN(文本类.X过滤首尾符并含空白(c.Get文本(ctx, "/")), g.Slice别名{"", "Client sent an HTTP request to an HTTPS server."})
		t.AssertIN(文本类.X过滤首尾符并含空白(c.Get文本(ctx, "/test")), g.Slice别名{"", "Client sent an HTTP request to an HTTPS server."})
	})
	// HTTPS
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("https://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Get文本(ctx, "/"), "Not Found")
		t.Assert(c.Get文本(ctx, "/test"), "test")
	})
}

func Test_HTTPS_HTTP_Basic(t *testing.T) {
	var (
		portHttp, _  = tcp类.GetFreePort()
		portHttps, _ = tcp类.GetFreePort()
	)
	s := g.Http类(时间类.X取文本时间戳纳秒())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定GET("/test", func(r *http类.Request) {
			r.Response.X写响应缓冲区("test")
		})
	})
	s.X启用HTTPS(
		单元测试类.DataPath("https", "files", "server.crt"),
		单元测试类.DataPath("https", "files", "server.key"),
	)
	s.X设置监听端口(portHttp)
	s.X设置HTTPS监听端口(portHttps)
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	// HTTP
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", portHttp))
		t.Assert(c.Get文本(ctx, "/"), "Not Found")
		t.Assert(c.Get文本(ctx, "/test"), "test")
	})
	// HTTPS
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("https://127.0.0.1:%d", portHttps))
		t.Assert(c.Get文本(ctx, "/"), "Not Found")
		t.Assert(c.Get文本(ctx, "/test"), "test")
	})
}
