// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

func Test_Cookie(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/set", func(r *http类.X请求) {
		r.Cookie.X设置值(r.Get别名("k").String(), r.Get别名("v").String())
	})
	s.X绑定("/get", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区(r.Cookie.X取值(r.Get别名("k").String()))
	})
	s.X绑定("/remove", func(r *http类.X请求) {
		r.Cookie.X删除值(r.Get别名("k").String())
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
	})
}

func Test_SetHttpCookie(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/set", func(r *http类.X请求) {
		r.Cookie.X设置httpcookie(&http.Cookie{
			Name:  r.Get别名("k").String(),
			Value: r.Get别名("v").String(),
		})
	})
	s.X绑定("/get", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区(r.Cookie.X取值(r.Get别名("k").String()))
	})
	s.X绑定("/remove", func(r *http类.X请求) {
		r.Cookie.X删除值(r.Get别名("k").String())
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
// t.Assert(client.GetContent(ctx, "/get?k=key2"), "200")
// 中文注释：断言通过上下文ctx，使用client获取"/get?k=key2"的请求内容，并判断其响应状态码应为"200"
// t.Assert(client.GetContent(ctx, "/get?k=key3"), "")
// 中文注释：断言通过上下文ctx，使用client获取"/get?k=key3"的请求内容，并判断其响应内容为空字符串
// t.Assert(client.GetContent(ctx, "/remove?k=key1"), "")
// 中文注释：断言通过上下文ctx，使用client发送"/remove?k=key1"的删除请求后，判断响应内容为空字符串
// t.Assert(client.GetContent(ctx, "/remove?k=key3"), "")
// 中文注释：断言通过上下文ctx，使用client发送"/remove?k=key3"的删除请求后，判断响应内容为空字符串
// t.Assert(client.GetContent(ctx, "/remove?k=key4"), "")
// 中文注释：断言通过上下文ctx，使用client发送"/remove?k=key4"的删除请求后，判断响应内容为空字符串
// t.Assert(client.GetContent(ctx, "/get?k=key1"), "")
// 中文注释：断言通过上下文ctx，使用client获取已删除的"/get?k=key1"的请求内容，判断响应内容为空字符串
// t.Assert(client.GetContent(ctx, "/get?k=key2"), "200")
// 中文注释：再次断言通过上下文ctx，使用client获取"/get?k=key2"的请求内容，并判断其响应状态码仍为"200"
// 上述代码是在进行HTTP接口测试，通过Assert方法验证请求和响应的结果是否符合预期。
	})
}

func Test_CookieOptionsDefault(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/test", func(r *http类.X请求) {
		r.Cookie.X设置值(r.Get别名("k").String(), r.Get别名("v").String())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
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

		t.AssertIN(len(parts), []int{3, 4}) // For go < 1.16 cookie always output "SameSite", see: https:// 这是Go语言标准库GitHub仓库中的一次提交记录，对应的commit（提交）哈希值为542693e00529fbb4248fac614ece68b127a5ec4d。
// 由于没有提供具体的代码片段，这里无法给出详细的代码注释翻译。通常这种形式的引用用于指向Go语言源码在GitHub上的某一次具体提交，该提交可能包含对某个功能的修复、优化或新功能的添加等。若需要了解这次提交的具体内容，可前往GitHub对应仓库查看此次commit的详细信息和改动内容。
	})
}

func Test_CookieOptions(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X设置配置项Map(g.Map{
		"cookieSameSite": "lax",
		"cookieSecure":   true,
		"cookieHttpOnly": true,
	})
	s.X绑定("/test", func(r *http类.X请求) {
		r.Cookie.X设置值(r.Get别名("k").String(), r.Get别名("v").String())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
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
