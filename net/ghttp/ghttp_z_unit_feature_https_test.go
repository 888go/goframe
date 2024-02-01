// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp_test
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
