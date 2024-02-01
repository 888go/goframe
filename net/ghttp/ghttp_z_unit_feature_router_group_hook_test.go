// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp_test
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
	s := g.Server(guid.S())
	group := s.Group("/api")
	group.GET("/handler", func(r *ghttp.Request) {
		r.Response.Write("1")
	})
	group.ALL("/handler", func(r *ghttp.Request) {
		r.Response.Write("0")
	}, ghttp.HookBeforeServe)
	group.ALL("/handler", func(r *ghttp.Request) {
		r.Response.Write("2")
	}, ghttp.HookAfterServe)

	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(client.GetContent(ctx, "/api/handler"), "012")
		t.Assert(client.PostContent(ctx, "/api/handler"), "02")
		t.Assert(client.GetContent(ctx, "/api/ThisDoesNotExist"), "Not Found")
	})
}

func Test_Router_Group_Hook2(t *testing.T) {
	s := g.Server(guid.S())
	group := s.Group("/api")
	group.GET("/handler", func(r *ghttp.Request) {
		r.Response.Write("1")
	})
	group.GET("/*", func(r *ghttp.Request) {
		r.Response.Write("0")
	}, ghttp.HookBeforeServe)
	group.GET("/*", func(r *ghttp.Request) {
		r.Response.Write("2")
	}, ghttp.HookAfterServe)

	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(client.GetContent(ctx, "/api/handler"), "012")
		t.Assert(client.PostContent(ctx, "/api/handler"), "Not Found")
		t.Assert(client.GetContent(ctx, "/api/ThisDoesNotExist"), "02")
		t.Assert(client.PostContent(ctx, "/api/ThisDoesNotExist"), "Not Found")
	})
}
