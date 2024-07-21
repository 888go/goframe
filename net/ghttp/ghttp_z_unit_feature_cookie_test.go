// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
)

func Test_Cookie(t *testing.T) {
	s := g.Server(guid.S())
	s.BindHandler("/set", func(r *ghttp.Request) {
		r.Cookie.Set(r.Get("k").String(), r.Get("v").String())
	})
	s.BindHandler("/get", func(r *ghttp.Request) {
		r.Response.Write(r.Cookie.Get(r.Get("k").String()))
	})
	s.BindHandler("/remove", func(r *ghttp.Request) {
		r.Cookie.Remove(r.Get("k").String())
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		r1, e1 := client.Get(ctx, "/set?k=key1&v=100")
		if r1 != nil {
			defer r1.Close()
		}

		t.Assert(e1, nil)
		t.Assert(r1.ReadAllString(), "")

		t.Assert(client.GetContent(ctx, "/set?k=key2&v=200"), "")

		t.Assert(client.GetContent(ctx, "/get?k=key1"), "100")
		t.Assert(client.GetContent(ctx, "/get?k=key2"), "200")
		t.Assert(client.GetContent(ctx, "/get?k=key3"), "")
		t.Assert(client.GetContent(ctx, "/remove?k=key1"), "")
		t.Assert(client.GetContent(ctx, "/remove?k=key3"), "")
		t.Assert(client.GetContent(ctx, "/remove?k=key4"), "")
		t.Assert(client.GetContent(ctx, "/get?k=key1"), "")
		t.Assert(client.GetContent(ctx, "/get?k=key2"), "200")
	})
}

func Test_SetHttpCookie(t *testing.T) {
	s := g.Server(guid.S())
	s.BindHandler("/set", func(r *ghttp.Request) {
		r.Cookie.SetHttpCookie(&http.Cookie{
			Name:  r.Get("k").String(),
			Value: r.Get("v").String(),
		})
	})
	s.BindHandler("/get", func(r *ghttp.Request) {
		r.Response.Write(r.Cookie.Get(r.Get("k").String()))
	})
	s.BindHandler("/remove", func(r *ghttp.Request) {
		r.Cookie.Remove(r.Get("k").String())
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		r1, e1 := client.Get(ctx, "/set?k=key1&v=100")
		if r1 != nil {
			defer r1.Close()
		}
		t.Assert(e1, nil)
		t.Assert(r1.ReadAllString(), "")

		t.Assert(client.GetContent(ctx, "/set?k=key2&v=200"), "")

		t.Assert(client.GetContent(ctx, "/get?k=key1"), "100")
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
	s := g.Server(guid.S())
	s.BindHandler("/test", func(r *ghttp.Request) {
		r.Cookie.Set(r.Get("k").String(), r.Get("v").String())
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		r1, e1 := client.Get(ctx, "/test?k=key1&v=100")
		if r1 != nil {
			defer r1.Close()
		}

		t.Assert(e1, nil)
		t.Assert(r1.ReadAllString(), "")

		parts := strings.Split(r1.Header.Get("Set-Cookie"), "; ")

		t.AssertIN(len(parts), []int{3, 4}) // 对于 Go < 1.16 版本，Cookie 始终会输出 "SameSite"，详情请参阅：https://github.com/golang/go/commit/542693e00529fbb4248fac614ece68b127a5ec4d. md5:6f46e9f2afc803ab
	})
}

func Test_CookieOptions(t *testing.T) {
	s := g.Server(guid.S())
	s.SetConfigWithMap(g.Map{
		"cookieSameSite": "lax",
		"cookieSecure":   true,
		"cookieHttpOnly": true,
	})
	s.BindHandler("/test", func(r *ghttp.Request) {
		r.Cookie.Set(r.Get("k").String(), r.Get("v").String())
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		r1, e1 := client.Get(ctx, "/test?k=key1&v=100")
		if r1 != nil {
			defer r1.Close()
		}

		t.Assert(e1, nil)
		t.Assert(r1.ReadAllString(), "")

		parts := strings.Split(r1.Header.Get("Set-Cookie"), "; ")

		t.AssertEQ(len(parts), 6)
		t.Assert(parts[3], "HttpOnly")
		t.Assert(parts[4], "Secure")
		t.Assert(parts[5], "SameSite=Lax")
	})
}
