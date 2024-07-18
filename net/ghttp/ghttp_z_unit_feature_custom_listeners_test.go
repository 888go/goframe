// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp_test

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
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
