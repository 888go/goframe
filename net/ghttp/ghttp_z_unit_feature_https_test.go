// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp_test

import (
	_ "github.com/gogf/gf/v2/net/ghttp/testdata/https/packed"

	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
)

func Test_HTTPS_Basic(t *testing.T) {
	s := g.Server(guid.S())
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/test", func(r *ghttp.Request) {
			r.Response.Write("test")
		})
	})
	s.EnableHTTPS(
		gtest.DataPath("https", "files", "server.crt"),
		gtest.DataPath("https", "files", "server.key"),
	)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	// HTTP
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.AssertIN(gstr.Trim(c.GetContent(ctx, "/")), g.Slice{"", "Client sent an HTTP request to an HTTPS server."})
		t.AssertIN(gstr.Trim(c.GetContent(ctx, "/test")), g.Slice{"", "Client sent an HTTP request to an HTTPS server."})
	})
	// HTTPS
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("https://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(c.GetContent(ctx, "/"), "Not Found")
		t.Assert(c.GetContent(ctx, "/test"), "test")
	})
}

func Test_HTTPS_Resource(t *testing.T) {
	s := g.Server(guid.S())
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/test", func(r *ghttp.Request) {
			r.Response.Write("test")
		})
	})
	s.EnableHTTPS(
		gfile.Join("files", "server.crt"),
		gfile.Join("files", "server.key"),
	)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	// HTTP
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.AssertIN(gstr.Trim(c.GetContent(ctx, "/")), g.Slice{"", "Client sent an HTTP request to an HTTPS server."})
		t.AssertIN(gstr.Trim(c.GetContent(ctx, "/test")), g.Slice{"", "Client sent an HTTP request to an HTTPS server."})
	})
	// HTTPS
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("https://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(c.GetContent(ctx, "/"), "Not Found")
		t.Assert(c.GetContent(ctx, "/test"), "test")
	})
}

func Test_HTTPS_HTTP_Basic(t *testing.T) {
	var (
		portHttp, _  = gtcp.GetFreePort()
		portHttps, _ = gtcp.GetFreePort()
	)
	s := g.Server(gtime.TimestampNanoStr())
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/test", func(r *ghttp.Request) {
			r.Response.Write("test")
		})
	})
	s.EnableHTTPS(
		gtest.DataPath("https", "files", "server.crt"),
		gtest.DataPath("https", "files", "server.key"),
	)
	s.SetPort(portHttp)
	s.SetHTTPSPort(portHttps)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	// HTTP
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", portHttp))
		t.Assert(c.GetContent(ctx, "/"), "Not Found")
		t.Assert(c.GetContent(ctx, "/test"), "test")
	})
	// HTTPS
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("https://127.0.0.1:%d", portHttps))
		t.Assert(c.GetContent(ctx, "/"), "Not Found")
		t.Assert(c.GetContent(ctx, "/test"), "test")
	})
}
