// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

type NamesObject struct{}

func (o *NamesObject) ShowName(r *ghttp.Request) {
	r.X响应.X写响应缓冲区("Object Show Name")
}

func Test_NameToUri_FullName(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.SetNameToUriType(ghttp.UriTypeFullName)
	s.X绑定对象("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X启用浏览器模式(true)
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.Get文本(ctx, "/NamesObject/ShowName"), "Object Show Name")
	})
}

func Test_NameToUri_AllLower(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.SetNameToUriType(ghttp.UriTypeAllLower)
	s.X绑定对象("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X启用浏览器模式(true)
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.Get文本(ctx, "/namesobject/showname"), "Object Show Name")
	})
}

func Test_NameToUri_Camel(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.SetNameToUriType(ghttp.UriTypeCamel)
	s.X绑定对象("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X启用浏览器模式(true)
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.Get文本(ctx, "/namesObject/showName"), "Object Show Name")
	})
}

func Test_NameToUri_Default(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.SetNameToUriType(ghttp.UriTypeDefault)
	s.X绑定对象("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X启用浏览器模式(true)
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.Get文本(ctx, "/names-object/show-name"), "Object Show Name")
	})
}
