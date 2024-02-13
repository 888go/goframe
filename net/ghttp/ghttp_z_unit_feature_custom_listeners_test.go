// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

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
		p1, _ = tcp类.GetFreePort()
		if ln1 == nil {
			ln1, _ = net.Listen("tcp", fmt.Sprintf(":%d", p1))
		}
		if ln1 != nil {
			break
		}
	}
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定GET("/test", func(r *http类.Request) {
			r.Response.X写响应缓冲区("test")
		})
	})
	err := s.X设置自定义监听器(ln1)
	单元测试类.AssertNil(err)

	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(
			文本类.X过滤首尾符并含空白(c.Get文本(ctx, "/test")),
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
		p1, _ = tcp类.GetFreePort()
		p2, _ = tcp类.GetFreePort()
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
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定GET("/test", func(r *http类.Request) {
			r.Response.X写响应缓冲区("test")
		})
	})

	err := s.X设置自定义监听器(ln1, ln2)
	单元测试类.AssertNil(err)

	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", p1))

		t.Assert(
			文本类.X过滤首尾符并含空白(c.Get文本(ctx, "/test")),
			"test",
		)

		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", p2))

		t.Assert(
			文本类.X过滤首尾符并含空白(c.Get文本(ctx, "/test")),
			"test",
		)
	})
}

func Test_SetWrongCustomListeners(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		s.X创建分组路由("/", func(group *http类.RouterGroup) {
			group.X绑定GET("/test", func(r *http类.Request) {
				r.Response.X写响应缓冲区("test")
			})
		})
		err := s.X设置自定义监听器(nil)
		t.AssertNQ(err, nil)
	})
}
