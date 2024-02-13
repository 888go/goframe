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

type NamesObject struct{}

func (o *NamesObject) ShowName(r *http类.Request) {
	r.Response.X写响应缓冲区("Object Show Name")
}

func Test_NameToUri_FullName(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.SetNameToUriType(http类.UriTypeFullName)
	s.X绑定对象("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X启用浏览器模式(true)
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.Get文本(ctx, "/NamesObject/ShowName"), "Object Show Name")
	})
}

func Test_NameToUri_AllLower(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.SetNameToUriType(http类.UriTypeAllLower)
	s.X绑定对象("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X启用浏览器模式(true)
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.Get文本(ctx, "/namesobject/showname"), "Object Show Name")
	})
}

func Test_NameToUri_Camel(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.SetNameToUriType(http类.UriTypeCamel)
	s.X绑定对象("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X启用浏览器模式(true)
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.Get文本(ctx, "/namesObject/showName"), "Object Show Name")
	})
}

func Test_NameToUri_Default(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.SetNameToUriType(http类.UriTypeDefault)
	s.X绑定对象("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X启用浏览器模式(true)
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.Get文本(ctx, "/names-object/show-name"), "Object Show Name")
	})
}
