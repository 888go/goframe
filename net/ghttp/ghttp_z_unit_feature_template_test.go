// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 静态服务测试。 md5:2105c089651008de

package http类_test

import (
	"fmt"
	"testing"
	"time"

	ghtml "github.com/888go/goframe/encoding/ghtml"
	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gview "github.com/888go/goframe/os/gview"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func Test_Template_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gview.New(gtest.DataPath("template", "basic"))
		s := g.Http类(guid.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/", func(r *ghttp.Request) {
			err := r.X响应.X输出到模板文件("index.html", g.Map{
				"name": "john",
			})
			t.AssertNil(err)
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Name:john")
		t.Assert(client.Get文本(ctx, "/"), "Name:john")
	})
}

func Test_Template_Encode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gview.New(gtest.DataPath("template", "basic"))
		v.SetAutoEncode(true)
		s := g.Http类(guid.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/", func(r *ghttp.Request) {
			err := r.X响应.X输出到模板文件("index.html", g.Map{
				"name": "john",
			})
			t.AssertNil(err)
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Name:john")
		t.Assert(client.Get文本(ctx, "/"), "Name:john")
	})
}

func Test_Template_Layout1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gview.New(gtest.DataPath("template", "layout1"))
		s := g.Http类(guid.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/layout", func(r *ghttp.Request) {
			err := r.X响应.X输出到模板文件("layout.html", g.Map{
				"mainTpl": "main/main1.html",
			})
			t.AssertNil(err)
		})
		s.X绑定("/nil", func(r *ghttp.Request) {
			err := r.X响应.X输出到模板文件("layout.html", nil)
			t.AssertNil(err)
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/layout"), "123")
		t.Assert(client.Get文本(ctx, "/nil"), "123")
	})
}

func Test_Template_Layout2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gview.New(gtest.DataPath("template", "layout2"))
		s := g.Http类(guid.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/main1", func(r *ghttp.Request) {
			err := r.X响应.X输出到模板文件("layout.html", g.Map{
				"mainTpl": "main/main1.html",
			})
			t.AssertNil(err)
		})
		s.X绑定("/main2", func(r *ghttp.Request) {
			err := r.X响应.X输出到模板文件("layout.html", g.Map{
				"mainTpl": "main/main2.html",
			})
			t.AssertNil(err)
		})
		s.X绑定("/nil", func(r *ghttp.Request) {
			err := r.X响应.X输出到模板文件("layout.html", nil)
			t.AssertNil(err)
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/main1"), "a1b")
		t.Assert(client.Get文本(ctx, "/main2"), "a2b")
		t.Assert(client.Get文本(ctx, "/nil"), "ab")
	})
}

func Test_Template_BuildInVarRequest(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Http类(guid.X生成())
		s.X绑定("/:table/test", func(r *ghttp.Request) {
			err := r.X响应.X输出文本模板("{{.Request.table}}")
			t.AssertNil(err)
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/user/test"), "user")
		t.Assert(client.Get文本(ctx, "/order/test"), "order")
	})
}

func Test_Template_XSS(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gview.New()
		v.SetAutoEncode(true)
		c := "<br>"
		s := g.Http类(guid.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/", func(r *ghttp.Request) {
			err := r.X响应.X输出文本模板("{{if eq 1 1}}{{.v}}{{end}}", g.Map{
				"v": c,
			})
			t.AssertNil(err)
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), ghtml.X编码(c))
	})
}
