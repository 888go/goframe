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

type ObjectRest struct{}

func (o *ObjectRest) Init(r *http类.Request) {
	r.Response.X写响应缓冲区("1")
}

func (o *ObjectRest) Shut(r *http类.Request) {
	r.Response.X写响应缓冲区("2")
}

func (o *ObjectRest) Get(r *http类.Request) {
	r.Response.X写响应缓冲区("Object Get")
}

func (o *ObjectRest) Put(r *http类.Request) {
	r.Response.X写响应缓冲区("Object Put")
}

func (o *ObjectRest) Post(r *http类.Request) {
	r.Response.X写响应缓冲区("Object Post")
}

func (o *ObjectRest) Delete(r *http类.Request) {
	r.Response.X写响应缓冲区("Object Delete")
}

func (o *ObjectRest) Patch(r *http类.Request) {
	r.Response.X写响应缓冲区("Object Patch")
}

func (o *ObjectRest) Options(r *http类.Request) {
	r.Response.X写响应缓冲区("Object Options")
}

func (o *ObjectRest) Head(r *http类.Request) {
	r.Response.Header().Set("head-ok", "1")
}

func Test_Router_ObjectRest(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定RESTfulAPI对象("/", new(ObjectRest))
	s.X绑定RESTfulAPI对象("/{.struct}/{.method}", new(ObjectRest))
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "1Object Get2")
		t.Assert(client.Put文本(ctx, "/"), "1Object Put2")
		t.Assert(client.Post文本(ctx, "/"), "1Object Post2")
		t.Assert(client.Delete文本(ctx, "/"), "1Object Delete2")
		t.Assert(client.Patch文本(ctx, "/"), "1Object Patch2")
		t.Assert(client.Options文本(ctx, "/"), "1Object Options2")
		resp1, err := client.Head响应对象(ctx, "/")
		if err == nil {
			defer resp1.X关闭()
		}
		t.AssertNil(err)
		t.Assert(resp1.Header.Get("head-ok"), "1")

		t.Assert(client.Get文本(ctx, "/object-rest/get"), "1Object Get2")
		t.Assert(client.Put文本(ctx, "/object-rest/put"), "1Object Put2")
		t.Assert(client.Post文本(ctx, "/object-rest/post"), "1Object Post2")
		t.Assert(client.Delete文本(ctx, "/object-rest/delete"), "1Object Delete2")
		t.Assert(client.Patch文本(ctx, "/object-rest/patch"), "1Object Patch2")
		t.Assert(client.Options文本(ctx, "/object-rest/options"), "1Object Options2")
		resp2, err := client.Head响应对象(ctx, "/object-rest/head")
		if err == nil {
			defer resp2.X关闭()
		}
		t.AssertNil(err)
		t.Assert(resp2.Header.Get("head-ok"), "1")

		t.Assert(client.Get文本(ctx, "/none-exist"), "Not Found")
	})
}
