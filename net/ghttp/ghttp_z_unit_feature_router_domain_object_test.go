		// 版权归GoFrame作者(https:		//goframe.org)所有。保留所有权利。
		//
		// 本源代码形式受MIT许可证条款约束。
		// 如果未随本文件一同分发MIT许可证副本，
		// 您可以在https:		//github.com/gogf/gf处获取。
		// md5:a9832f33b234e3f3

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
)

type DomainObject struct{}

func (o *DomainObject) Init(r *ghttp.Request) {
	r.Response.Write("1")
}

func (o *DomainObject) Shut(r *ghttp.Request) {
	r.Response.Write("2")
}

func (o *DomainObject) Index(r *ghttp.Request) {
	r.Response.Write("Object Index")
}

func (o *DomainObject) Show(r *ghttp.Request) {
	r.Response.Write("Object Show")
}

func (o *DomainObject) Info(r *ghttp.Request) {
	r.Response.Write("Object Info")
}

func Test_Router_DomainObject1(t *testing.T) {
	s := g.Server(guid.S())
	s.Domain("localhost, local").BindObject("/", new(DomainObject))
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/init"), "Not Found")
		t.Assert(client.GetContent(ctx, "/shut"), "Not Found")
		t.Assert(client.GetContent(ctx, "/index"), "Not Found")
		t.Assert(client.GetContent(ctx, "/show"), "Not Found")
		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})

	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://localhost:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "1Object Index2")
		t.Assert(client.GetContent(ctx, "/init"), "Not Found")
		t.Assert(client.GetContent(ctx, "/shut"), "Not Found")
		t.Assert(client.GetContent(ctx, "/index"), "1Object Index2")
		t.Assert(client.GetContent(ctx, "/show"), "1Object Show2")
		t.Assert(client.GetContent(ctx, "/info"), "1Object Info2")
		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})

	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://local:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "1Object Index2")
		t.Assert(client.GetContent(ctx, "/init"), "Not Found")
		t.Assert(client.GetContent(ctx, "/shut"), "Not Found")
		t.Assert(client.GetContent(ctx, "/index"), "1Object Index2")
		t.Assert(client.GetContent(ctx, "/show"), "1Object Show2")
		t.Assert(client.GetContent(ctx, "/info"), "1Object Info2")
		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})
}

func Test_Router_DomainObject2(t *testing.T) {
	s := g.Server(guid.S())
	s.Domain("localhost, local").BindObject("/object", new(DomainObject), "Show, Info")
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/init"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/shut"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/index"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/show"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/info"), "Not Found")
		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://localhost:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/init"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/shut"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/index"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/show"), "1Object Show2")
		t.Assert(client.GetContent(ctx, "/object/info"), "1Object Info2")
		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://local:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/init"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/shut"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/index"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/show"), "1Object Show2")
		t.Assert(client.GetContent(ctx, "/object/info"), "1Object Info2")
		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})
}

func Test_Router_DomainObjectMethod(t *testing.T) {
	s := g.Server(guid.S())
	s.Domain("localhost, local").BindObjectMethod("/object-info", new(DomainObject), "Info")
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/init"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/shut"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/index"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/show"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/info"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object-info"), "Not Found")
		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://localhost:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/init"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/shut"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/index"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/show"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/info"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object-info"), "1Object Info2")
		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://local:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/init"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/shut"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/index"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/show"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/info"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object-info"), "1Object Info2")
		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})
}
