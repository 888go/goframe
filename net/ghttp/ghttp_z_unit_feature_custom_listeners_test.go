// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp_test

import (
	"fmt"
	"net"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/net/gtcp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/guid"
)

func Test_SetSingleCustomListener(t *testing.T) {
	var (
		p1  int
		ln1 net.Listener
	)
	for i := 0; i < 1000; i++ {
		p1, _ = gtcp.GetFreePort()
		if ln1 == nil {
			ln1, _ = net.Listen("tcp", fmt.Sprintf(":%d", p1))
		}
		if ln1 != nil {
			break
		}
	}
	s := g.Server(guid.S())
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/test", func(r *ghttp.Request) {
			r.Response.Write("test")
		})
	})
	err := s.SetListener(ln1)
	gtest.AssertNil(err)

	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(
			gstr.Trim(c.GetContent(ctx, "/test")),
			"test",
		)
	})
}

func Test_SetMultipleCustomListeners(t *testing.T) {
	var (
		p1  int
		p2  int
		ln1 net.Listener
		ln2 net.Listener
	)
	for i := 0; i < 1000; i++ {
		p1, _ = gtcp.GetFreePort()
		p2, _ = gtcp.GetFreePort()
		if ln1 == nil {
			ln1, _ = net.Listen("tcp", fmt.Sprintf(":%d", p1))
		}
		if ln2 == nil {
			ln2, _ = net.Listen("tcp", fmt.Sprintf(":%d", p2))
		}
		if ln1 != nil && ln2 != nil {
			break
		}
	}
	s := g.Server(guid.S())
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/test", func(r *ghttp.Request) {
			r.Response.Write("test")
		})
	})

	err := s.SetListener(ln1, ln2)
	gtest.AssertNil(err)

	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p1))

		t.Assert(
			gstr.Trim(c.GetContent(ctx, "/test")),
			"test",
		)

		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p2))

		t.Assert(
			gstr.Trim(c.GetContent(ctx, "/test")),
			"test",
		)
	})
}

func Test_SetWrongCustomListeners(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.GET("/test", func(r *ghttp.Request) {
				r.Response.Write("test")
			})
		})
		err := s.SetListener(nil)
		t.AssertNQ(err, nil)
	})
}
