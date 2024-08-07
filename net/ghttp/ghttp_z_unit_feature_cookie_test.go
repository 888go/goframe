// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func Test_Cookie(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/set", func(r *ghttp.Request) {
		r.Cookie.X设置值(r.Get别名("k").String(), r.Get别名("v").String())
	})
	s.X绑定("/get", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区(r.Cookie.X取值(r.Get别名("k").String()))
	})
	s.X绑定("/remove", func(r *ghttp.Request) {
		r.Cookie.X删除值(r.Get别名("k").String())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
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
	})
}

func Test_SetHttpCookie(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/set", func(r *ghttp.Request) {
		r.Cookie.X设置httpcookie(&http.Cookie{
			Name:  r.Get别名("k").String(),
			Value: r.Get别名("v").String(),
		})
	})
	s.X绑定("/get", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区(r.Cookie.X取值(r.Get别名("k").String()))
	})
	s.X绑定("/remove", func(r *ghttp.Request) {
		r.Cookie.X删除值(r.Get别名("k").String())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
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
		// 测试断言：获取"/get?k=key2"的响应状态为200
		// 测试断言：获取"/get?k=key3"的响应为空字符串
		// 测试断言：获取"/remove?k=key1"的响应为空字符串
		// 测试断言：获取"/remove?k=key3"的响应为空字符串
		// 测试断言：获取"/remove?k=key4"的响应为空字符串
		// 测试断言：获取"/get?k=key1"的响应为空字符串
		// 测试断言：获取"/get?k=key2"的响应状态为200
		// md5:fa4c58c1c55bab25
	})
}

func Test_CookieOptionsDefault(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/test", func(r *ghttp.Request) {
		r.Cookie.X设置值(r.Get别名("k").String(), r.Get别名("v").String())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X启用浏览器模式(true)
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		r1, e1 := client.Get响应对象(ctx, "/test?k=key1&v=100")
		if r1 != nil {
			defer r1.X关闭()
		}

		t.Assert(e1, nil)
		t.Assert(r1.X取响应文本(), "")

		parts := strings.Split(r1.Header.Get("Set-Cookie"), "; ")

		t.AssertIN(len(parts), []int{3, 4}) //github.com/golang/go/commit/542693e00529fbb4248fac614ece68b127a5ec4d. md5:6f46e9f2afc803ab
	})
}

func Test_CookieOptions(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X设置配置项Map(g.Map{
		"cookieSameSite": "lax",
		"cookieSecure":   true,
		"cookieHttpOnly": true,
	})
	s.X绑定("/test", func(r *ghttp.Request) {
		r.Cookie.X设置值(r.Get别名("k").String(), r.Get别名("v").String())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X启用浏览器模式(true)
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		r1, e1 := client.Get响应对象(ctx, "/test?k=key1&v=100")
		if r1 != nil {
			defer r1.X关闭()
		}

		t.Assert(e1, nil)
		t.Assert(r1.X取响应文本(), "")

		parts := strings.Split(r1.Header.Get("Set-Cookie"), "; ")

		t.AssertEQ(len(parts), 6)
		t.Assert(parts[3], "HttpOnly")
		t.Assert(parts[4], "Secure")
		t.Assert(parts[5], "SameSite=Lax")
	})
}
