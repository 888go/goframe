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

type ObjectRest2 struct{}

func (o *ObjectRest2) Init(r *http类.Request) {
	r.Response.X写响应缓冲区("1")
}

func (o *ObjectRest2) Shut(r *http类.Request) {
	r.Response.X写响应缓冲区("2")
}

func (o *ObjectRest2) Get(r *http类.Request) {
	r.Response.X写响应缓冲区("Object Get", r.Get别名("id"))
}

func (o *ObjectRest2) Put(r *http类.Request) {
	r.Response.X写响应缓冲区("Object Put", r.Get别名("id"))
}

func (o *ObjectRest2) Post(r *http类.Request) {
	r.Response.X写响应缓冲区("Object Post", r.Get别名("id"))
}

func (o *ObjectRest2) Delete(r *http类.Request) {
	r.Response.X写响应缓冲区("Object Delete", r.Get别名("id"))
}

func Test_Router_ObjectRest_Id(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定RESTfulAPI对象("/object/:id", new(ObjectRest2))
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/object/99"), "1Object Get992")
		t.Assert(client.Put文本(ctx, "/object/99"), "1Object Put992")
		t.Assert(client.Post文本(ctx, "/object/99"), "1Object Post992")
		t.Assert(client.Delete文本(ctx, "/object/99"), "1Object Delete992")
	})
}
