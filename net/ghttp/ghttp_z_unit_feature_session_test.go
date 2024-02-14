// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

import (
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

func Test_Session_Cookie(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/set", func(r *http类.X请求) {
		r.Session.X设置值(r.Get别名("k").String(), r.Get别名("v").String())
	})
	s.X绑定("/get", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区(r.Session.Get(r.Get别名("k").String()))
	})
	s.X绑定("/remove", func(r *http类.X请求) {
		r.Session.Remove(r.Get别名("k").String())
	})
	s.X绑定("/clear", func(r *http类.X请求) {
		r.Session.RemoveAll()
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X启用浏览器模式(true)
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		r1, e1 := client.Get响应对象(ctx, "/set?k=key1&v=100")
		if r1 != nil {
			defer r1.X关闭()
		}
		t.Assert(e1, nil)
		t.Assert(r1.X取响应文本(), "")

		t.Assert(client.Get文本(ctx, "/set?k=key2&v=200"), "")

		t.Assert(client.Get文本(ctx, "/get?k=key1"), "100")
		t.Assert(client.Get文本(ctx, "/get?k=key2"), "200")
		t.Assert(client.Get文本(ctx, "/get?k=key3"), "")
		t.Assert(client.Get文本(ctx, "/remove?k=key1"), "")
		t.Assert(client.Get文本(ctx, "/remove?k=key3"), "")
		t.Assert(client.Get文本(ctx, "/remove?k=key4"), "")
		t.Assert(client.Get文本(ctx, "/get?k=key1"), "")
		t.Assert(client.Get文本(ctx, "/get?k=key2"), "200")
		t.Assert(client.Get文本(ctx, "/clear"), "")
		t.Assert(client.Get文本(ctx, "/get?k=key2"), "")
	})
}

func Test_Session_Header(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/set", func(r *http类.X请求) {
		r.Session.X设置值(r.Get别名("k").String(), r.Get别名("v").String())
	})
	s.X绑定("/get", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区(r.Session.Get(r.Get别名("k").String()))
	})
	s.X绑定("/remove", func(r *http类.X请求) {
		r.Session.Remove(r.Get别名("k").String())
	})
	s.X绑定("/clear", func(r *http类.X请求) {
		r.Session.RemoveAll()
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		response, e1 := client.Get响应对象(ctx, "/set?k=key1&v=100")
		if response != nil {
			defer response.X关闭()
		}
		sessionId := response.X取Cookie(s.X取SessionID名称())
		t.Assert(e1, nil)
		t.AssertNE(sessionId, nil)
		t.Assert(response.X取响应文本(), "")

		client.X设置协议头(s.X取SessionID名称(), sessionId)

		t.Assert(client.Get文本(ctx, "/set?k=key2&v=200"), "")

		t.Assert(client.Get文本(ctx, "/get?k=key1"), "100")
		t.Assert(client.Get文本(ctx, "/get?k=key2"), "200")
		t.Assert(client.Get文本(ctx, "/get?k=key3"), "")
		t.Assert(client.Get文本(ctx, "/remove?k=key1"), "")
		t.Assert(client.Get文本(ctx, "/remove?k=key3"), "")
		t.Assert(client.Get文本(ctx, "/remove?k=key4"), "")
		t.Assert(client.Get文本(ctx, "/get?k=key1"), "")
		t.Assert(client.Get文本(ctx, "/get?k=key2"), "200")
		t.Assert(client.Get文本(ctx, "/clear"), "")
		t.Assert(client.Get文本(ctx, "/get?k=key2"), "")
	})
}

func Test_Session_StorageFile(t *testing.T) {
	sessionId := ""
	s := g.Http类(uid类.X生成())
	s.X绑定("/set", func(r *http类.X请求) {
		r.Session.X设置值(r.Get别名("k").String(), r.Get别名("v").String())
		r.X响应.X写响应缓冲区(r.Get别名("k").String(), "=", r.Get别名("v").String())
	})
	s.X绑定("/get", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区(r.Session.Get(r.Get别名("k").String()))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		response, e1 := client.Get响应对象(ctx, "/set?k=key&v=100")
		if response != nil {
			defer response.X关闭()
		}
		sessionId = response.X取Cookie(s.X取SessionID名称())
		t.Assert(e1, nil)
		t.AssertNE(sessionId, nil)
		t.Assert(response.X取响应文本(), "key=100")
	})
	time.Sleep(time.Second)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		client.X设置协议头(s.X取SessionID名称(), sessionId)
		t.Assert(client.Get文本(ctx, "/get?k=key"), "100")
		t.Assert(client.Get文本(ctx, "/get?k=key1"), "")
	})
}

func Test_Session_Custom_Id(t *testing.T) {
	var (
		sessionId = "1234567890"
		key       = "key"
		value     = "value"
		s         = g.Http类(uid类.X生成())
	)
	s.X绑定("/id", func(r *http类.X请求) {
		if err := r.Session.SetId(sessionId); err != nil {
			r.X响应.X写响应缓冲区并退出(err.Error())
		}
		if err := r.Session.X设置值(key, value); err != nil {
			r.X响应.X写响应缓冲区并退出(err.Error())
		}
		r.X响应.X写响应缓冲区并退出(r.Session.Id())
	})
	s.X绑定("/value", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区并退出(r.Session.Get(key))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		r, err := client.Get响应对象(ctx, "/id")
		t.AssertNil(err)
		defer r.X关闭()
		t.Assert(r.X取响应文本(), sessionId)
		t.Assert(r.X取Cookie(s.X取SessionID名称()), sessionId)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		client.X设置协议头(s.X取SessionID名称(), sessionId)
		t.Assert(client.Get文本(ctx, "/value"), value)
	})
}

func Test_Session_New_Id(t *testing.T) {
	var (
		sessionId     = "1234567890"
		newSessionId  = "0987654321"
		newSessionId2 = "abcdefghij"
		key           = "key"
		value         = "value"
		s             = g.Http类(uid类.X生成())
	)
	s.X绑定("/id", func(r *http类.X请求) {
		if err := r.Session.SetId(sessionId); err != nil {
			r.X响应.X写响应缓冲区并退出(err.Error())
		}
		if err := r.Session.X设置值(key, value); err != nil {
			r.X响应.X写响应缓冲区并退出(err.Error())
		}
		r.X响应.X写响应缓冲区并退出(r.Session.Id())
	})

	s.X绑定("/newIdBySession", func(r *http类.X请求) {
		// 在会话初始化之前使用
		if err := r.Session.SetId(newSessionId); err != nil {
			r.X响应.X写响应缓冲区并退出(err.Error())
		}
		if err := r.Session.X设置值(key, value); err != nil {
			r.X响应.X写响应缓冲区并退出(err.Error())
		}
		r.X响应.X写响应缓冲区并退出(r.Session.Id())
	})

	s.X绑定("/newIdByCookie", func(r *http类.X请求) {
		if err := r.Session.Remove("someKey"); err != nil {
			r.X响应.X写响应缓冲区并退出(err.Error())
		}

		r.Cookie.X设置SessionId到Cookie(newSessionId2)
// r.Response.WriteExit(r.Session.Id()) // 只修改cookie中的内容

		r.X响应.X写响应缓冲区并退出(newSessionId2)
	})

	s.X绑定("/value", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区并退出(r.Session.Get(key))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		r, err := client.Get响应对象(ctx, "/id")
		t.AssertNil(err)
		defer r.X关闭()
		t.Assert(r.X取响应文本(), sessionId)
		t.Assert(r.X取Cookie(s.X取SessionID名称()), sessionId)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		client.X设置协议头(s.X取SessionID名称(), sessionId)
		t.Assert(client.Get文本(ctx, "/value"), value)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		client.X设置协议头(s.X取SessionID名称(), sessionId)
		r, err := client.Get响应对象(ctx, "/newIdBySession")
		t.AssertNil(err)
		defer r.X关闭()
		t.Assert(r.X取响应文本(), newSessionId)
		t.Assert(r.X取Cookie(s.X取SessionID名称()), newSessionId)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		r, err := client.Get响应对象(ctx, "/newIdByCookie")
		client.X设置协议头(s.X取SessionID名称(), sessionId)
		t.AssertNil(err)
		defer r.X关闭()
		t.Assert(r.X取响应文本(), newSessionId2)
		t.Assert(r.X取Cookie(s.X取SessionID名称()), newSessionId2)
	})
}
