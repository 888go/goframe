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

type ObjectRest2 struct{}

func (o *ObjectRest2) Init(r *ghttp.Request) {
	r.X响应.X写响应缓冲区("1")
}

func (o *ObjectRest2) Shut(r *ghttp.Request) {
	r.X响应.X写响应缓冲区("2")
}

func (o *ObjectRest2) Get(r *ghttp.Request) {
	r.X响应.X写响应缓冲区("Object Get", r.Get别名("id"))
}

func (o *ObjectRest2) Put(r *ghttp.Request) {
	r.X响应.X写响应缓冲区("Object Put", r.Get别名("id"))
}

func (o *ObjectRest2) Post(r *ghttp.Request) {
	r.X响应.X写响应缓冲区("Object Post", r.Get别名("id"))
}

func (o *ObjectRest2) Delete(r *ghttp.Request) {
	r.X响应.X写响应缓冲区("Object Delete", r.Get别名("id"))
}

func Test_Router_ObjectRest_Id(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定RESTfulAPI对象("/object/:id", new(ObjectRest2))
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/object/99"), "1Object Get992")
		t.Assert(client.Put文本(ctx, "/object/99"), "1Object Put992")
		t.Assert(client.Post文本(ctx, "/object/99"), "1Object Post992")
		t.Assert(client.Delete文本(ctx, "/object/99"), "1Object Delete992")
	})
}
