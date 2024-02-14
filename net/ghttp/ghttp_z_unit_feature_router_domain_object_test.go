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

type DomainObject struct{}

func (o *DomainObject) Init(r *http类.X请求) {
	r.X响应.X写响应缓冲区("1")
}

func (o *DomainObject) Shut(r *http类.X请求) {
	r.X响应.X写响应缓冲区("2")
}

func (o *DomainObject) Index(r *http类.X请求) {
	r.X响应.X写响应缓冲区("Object Index")
}

func (o *DomainObject) Show(r *http类.X请求) {
	r.X响应.X写响应缓冲区("Object Show")
}

func (o *DomainObject) Info(r *http类.X请求) {
	r.X响应.X写响应缓冲区("Object Info")
}

func Test_Router_DomainObject1(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建域名路由("localhost, local").X绑定对象("/", new(DomainObject))
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/init"), "Not Found")
		t.Assert(client.Get文本(ctx, "/shut"), "Not Found")
		t.Assert(client.Get文本(ctx, "/index"), "Not Found")
		t.Assert(client.Get文本(ctx, "/show"), "Not Found")
		t.Assert(client.Get文本(ctx, "/none-exist"), "Not Found")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://localhost:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "1Object Index2")
		t.Assert(client.Get文本(ctx, "/init"), "Not Found")
		t.Assert(client.Get文本(ctx, "/shut"), "Not Found")
		t.Assert(client.Get文本(ctx, "/index"), "1Object Index2")
		t.Assert(client.Get文本(ctx, "/show"), "1Object Show2")
		t.Assert(client.Get文本(ctx, "/info"), "1Object Info2")
		t.Assert(client.Get文本(ctx, "/none-exist"), "Not Found")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://local:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "1Object Index2")
		t.Assert(client.Get文本(ctx, "/init"), "Not Found")
		t.Assert(client.Get文本(ctx, "/shut"), "Not Found")
		t.Assert(client.Get文本(ctx, "/index"), "1Object Index2")
		t.Assert(client.Get文本(ctx, "/show"), "1Object Show2")
		t.Assert(client.Get文本(ctx, "/info"), "1Object Info2")
		t.Assert(client.Get文本(ctx, "/none-exist"), "Not Found")
	})
}

func Test_Router_DomainObject2(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建域名路由("localhost, local").X绑定对象("/object", new(DomainObject), "Show, Info")
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/init"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/shut"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/index"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/show"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/info"), "Not Found")
		t.Assert(client.Get文本(ctx, "/none-exist"), "Not Found")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://localhost:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/init"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/shut"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/index"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/show"), "1Object Show2")
		t.Assert(client.Get文本(ctx, "/object/info"), "1Object Info2")
		t.Assert(client.Get文本(ctx, "/none-exist"), "Not Found")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://local:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/init"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/shut"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/index"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/show"), "1Object Show2")
		t.Assert(client.Get文本(ctx, "/object/info"), "1Object Info2")
		t.Assert(client.Get文本(ctx, "/none-exist"), "Not Found")
	})
}

func Test_Router_DomainObjectMethod(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建域名路由("localhost, local").X绑定对象方法("/object-info", new(DomainObject), "Info")
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/init"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/shut"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/index"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/show"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/info"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object-info"), "Not Found")
		t.Assert(client.Get文本(ctx, "/none-exist"), "Not Found")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://localhost:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/init"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/shut"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/index"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/show"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/info"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object-info"), "1Object Info2")
		t.Assert(client.Get文本(ctx, "/none-exist"), "Not Found")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://local:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/init"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/shut"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/index"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/show"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object/info"), "Not Found")
		t.Assert(client.Get文本(ctx, "/object-info"), "1Object Info2")
		t.Assert(client.Get文本(ctx, "/none-exist"), "Not Found")
	})
}
