// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"testing"
	"time"

	gurl "github.com/888go/goframe/encoding/gurl"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/httputil"
	ghttp "github.com/888go/goframe/net/ghttp"
	genv "github.com/888go/goframe/os/genv"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

var (
	ctx = context.TODO()
)

func init() {
	genv.X设置值("UNDER_TEST", "1")
}

func Test_GetUrl(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/url", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区(r.X取URL())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X启用浏览器模式(true)
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/url"), prefix+"/url")
	})
}

func Test_XUrlPath(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/test1", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区(`test1`)
	})
	s.X绑定("/test2", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区(`test2`)
	})
	s.X设置请求处理器(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set(ghttp.HeaderXUrlPath, "/test2")
		s.ServeHTTP(w, r)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/"), "test2")
		t.Assert(c.Get文本(ctx, "/test/test"), "test2")
		t.Assert(c.Get文本(ctx, "/test1"), "test2")
	})
}

func Test_GetListenedAddress(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区(`test`)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Get文本(ctx, "/"), "test")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(fmt.Sprintf(`:%d`, s.X取已监听端口()), s.X取已监听地址())
	})
}

func Test_GetListenedAddressWithHost(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区(`test`)
	})
	s.X设置监听地址("127.0.0.1:0")
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Get文本(ctx, "/"), "test")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(fmt.Sprintf(`127.0.0.1:%d`, s.X取已监听端口()), s.X取已监听地址())
	})
}

func Test_RoutePathParams(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/:param", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区(r.Get别名("param"), ",", r.Get别名("c"))
	})
	s.X设置监听地址("127.0.0.1:0")
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.X网页类()
		param := "net/http/get"
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Get文本(
			ctx,
			"/"+gurl.X编码(param)+"?a=1&b=2&c="+gurl.X编码(param)),
			"net/http/get,net/http/get",
		)
	})
}

func Test_BuildParams(t *testing.T) {
		// 正常情况和特殊情况. md5:e92447df0c42f0ae
	params := map[string]string{
		"val":   "12345678",
		"code1": "x&a=1", // for fix
		"code2": "x&a=111",
		"id":    "1+- ", // for fix
		"f":     "1#a=+- ",
		"v":     "",
		"n":     "null",
	}

	gtest.C(t, func(t *gtest.T) {
		res1 := httputil.BuildParams(params)
		vs, _ := url.ParseQuery(res1)
		t.Assert(len(params), len(vs))
		for k := range vs {
			vv := vs.Get(k)
			_, ok := params[k]
									// 检查没有多余的参数. md5:0a0f11b34d795824
			t.Assert(ok, true)
			// check equal
			t.AssertEQ(params[k], vv)
		}
	})
}

func Test_ServerSignal(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Log("skip windows")
		return
	}
	s := g.Http类(guid.X生成())
	s.X绑定("/", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("hello world")
	})
	gtest.Assert(s.X开始监听(), nil)
	g.Http类等待所有服务完成()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(s.X关闭当前服务(), nil)
	})
}
