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

type DomainObjectRest struct{}

func (o *DomainObjectRest) Init(r *ghttp.Request) {
	r.Response.Write("1")
}

func (o *DomainObjectRest) Shut(r *ghttp.Request) {
	r.Response.Write("2")
}

func (o *DomainObjectRest) Get(r *ghttp.Request) {
	r.Response.Write("Object Get")
}

func (o *DomainObjectRest) Put(r *ghttp.Request) {
	r.Response.Write("Object Put")
}

func (o *DomainObjectRest) Post(r *ghttp.Request) {
	r.Response.Write("Object Post")
}

func (o *DomainObjectRest) Delete(r *ghttp.Request) {
	r.Response.Write("Object Delete")
}

func (o *DomainObjectRest) Patch(r *ghttp.Request) {
	r.Response.Write("Object Patch")
}

func (o *DomainObjectRest) Options(r *ghttp.Request) {
	r.Response.Write("Object Options")
}

func (o *DomainObjectRest) Head(r *ghttp.Request) {
	r.Response.Header().Set("head-ok", "1")
}

func Test_Router_DomainObjectRest(t *testing.T) {
	s := g.Server(guid.S())
	d := s.Domain("localhost, local")
	d.BindObjectRest("/", new(DomainObjectRest))
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.PutContent(ctx, "/"), "Not Found")
		t.Assert(client.PostContent(ctx, "/"), "Not Found")
		t.Assert(client.DeleteContent(ctx, "/"), "Not Found")
		t.Assert(client.PatchContent(ctx, "/"), "Not Found")
		t.Assert(client.OptionsContent(ctx, "/"), "Not Found")
		resp1, err := client.Head(ctx, "/")
		if err == nil {
			defer resp1.Close()
		}
		t.AssertNil(err)
		t.Assert(resp1.Header.Get("head-ok"), "")
		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://localhost:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "1Object Get2")
		t.Assert(client.PutContent(ctx, "/"), "1Object Put2")
		t.Assert(client.PostContent(ctx, "/"), "1Object Post2")
		t.Assert(client.DeleteContent(ctx, "/"), "1Object Delete2")
		t.Assert(client.PatchContent(ctx, "/"), "1Object Patch2")
		t.Assert(client.OptionsContent(ctx, "/"), "1Object Options2")
		resp1, err := client.Head(ctx, "/")
		if err == nil {
			defer resp1.Close()
		}
		t.AssertNil(err)
		t.Assert(resp1.Header.Get("head-ok"), "1")
		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://local:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "1Object Get2")
		t.Assert(client.PutContent(ctx, "/"), "1Object Put2")
		t.Assert(client.PostContent(ctx, "/"), "1Object Post2")
		t.Assert(client.DeleteContent(ctx, "/"), "1Object Delete2")
		t.Assert(client.PatchContent(ctx, "/"), "1Object Patch2")
		t.Assert(client.OptionsContent(ctx, "/"), "1Object Options2")
		resp1, err := client.Head(ctx, "/")
		if err == nil {
			defer resp1.Close()
		}
		t.AssertNil(err)
		t.Assert(resp1.Header.Get("head-ok"), "1")
		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})
}
