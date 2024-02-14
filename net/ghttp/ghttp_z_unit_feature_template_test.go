// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 静态服务测试。

package http类_test

import (
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/encoding/ghtml"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/os/gview"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

func Test_Template_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		v := 模板类.New(单元测试类.DataPath("template", "basic"))
		s := g.Http类(uid类.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/", func(r *http类.X请求) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		v := 模板类.New(单元测试类.DataPath("template", "basic"))
		v.SetAutoEncode(true)
		s := g.Http类(uid类.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/", func(r *http类.X请求) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		v := 模板类.New(单元测试类.DataPath("template", "layout1"))
		s := g.Http类(uid类.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/layout", func(r *http类.X请求) {
			err := r.X响应.X输出到模板文件("layout.html", g.Map{
				"mainTpl": "main/main1.html",
			})
			t.AssertNil(err)
		})
		s.X绑定("/nil", func(r *http类.X请求) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		v := 模板类.New(单元测试类.DataPath("template", "layout2"))
		s := g.Http类(uid类.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/main1", func(r *http类.X请求) {
			err := r.X响应.X输出到模板文件("layout.html", g.Map{
				"mainTpl": "main/main1.html",
			})
			t.AssertNil(err)
		})
		s.X绑定("/main2", func(r *http类.X请求) {
			err := r.X响应.X输出到模板文件("layout.html", g.Map{
				"mainTpl": "main/main2.html",
			})
			t.AssertNil(err)
		})
		s.X绑定("/nil", func(r *http类.X请求) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		s.X绑定("/:table/test", func(r *http类.X请求) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		v := 模板类.New()
		v.SetAutoEncode(true)
		c := "<br>"
		s := g.Http类(uid类.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/", func(r *http类.X请求) {
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

		t.Assert(client.Get文本(ctx, "/"), html类.X编码(c))
	})
}
