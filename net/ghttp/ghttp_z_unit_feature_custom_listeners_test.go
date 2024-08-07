// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtcp "github.com/888go/goframe/net/gtcp"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
	guid "github.com/888go/goframe/util/guid"
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
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定GET("/test", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("test")
		})
	})
	err := s.X设置自定义监听器(ln1)
	gtest.AssertNil(err)

	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(
			gstr.X过滤首尾符并含空白(c.Get文本(ctx, "/test")),
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
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定GET("/test", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("test")
		})
	})

	err := s.X设置自定义监听器(ln1, ln2)
	gtest.AssertNil(err)

	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", p1))

		t.Assert(
			gstr.X过滤首尾符并含空白(c.Get文本(ctx, "/test")),
			"test",
		)

		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", p2))

		t.Assert(
			gstr.X过滤首尾符并含空白(c.Get文本(ctx, "/test")),
			"test",
		)
	})
}

func Test_SetWrongCustomListeners(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Http类(guid.X生成())
		s.X创建分组路由("/", func(group *ghttp.X分组路由) {
			group.X绑定GET("/test", func(r *ghttp.Request) {
				r.X响应.X写响应缓冲区("test")
			})
		})
		err := s.X设置自定义监听器(nil)
		t.AssertNQ(err, nil)
	})
}
