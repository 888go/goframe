// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp_test

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

func (o *ObjectRest2) Init(r *ghttp.Request) {
	r.Response.Write("1")
}

func (o *ObjectRest2) Shut(r *ghttp.Request) {
	r.Response.Write("2")
}

func (o *ObjectRest2) Get(r *ghttp.Request) {
	r.Response.Write("Object Get", r.Get("id"))
}

func (o *ObjectRest2) Put(r *ghttp.Request) {
	r.Response.Write("Object Put", r.Get("id"))
}

func (o *ObjectRest2) Post(r *ghttp.Request) {
	r.Response.Write("Object Post", r.Get("id"))
}

func (o *ObjectRest2) Delete(r *ghttp.Request) {
	r.Response.Write("Object Delete", r.Get("id"))
}

func Test_Router_ObjectRest_Id(t *testing.T) {
	s := g.Server(guid.S())
	s.BindObjectRest("/object/:id", new(ObjectRest2))
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/object/99"), "1Object Get992")
		t.Assert(client.PutContent(ctx, "/object/99"), "1Object Put992")
		t.Assert(client.PostContent(ctx, "/object/99"), "1Object Post992")
		t.Assert(client.DeleteContent(ctx, "/object/99"), "1Object Delete992")
	})
}
