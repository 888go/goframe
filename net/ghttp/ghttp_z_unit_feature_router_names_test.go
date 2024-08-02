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
	r.Response.Write("Object Show Name")
}

func Test_NameToUri_FullName(t *testing.T) {
	s := g.Server(guid.S())
	s.SetNameToUriType(ghttp.UriTypeFullName)
	s.BindObject("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.GetContent(ctx, "/NamesObject/ShowName"), "Object Show Name")
	})
}

func Test_NameToUri_AllLower(t *testing.T) {
	s := g.Server(guid.S())
	s.SetNameToUriType(ghttp.UriTypeAllLower)
	s.BindObject("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.GetContent(ctx, "/namesobject/showname"), "Object Show Name")
	})
}

func Test_NameToUri_Camel(t *testing.T) {
	s := g.Server(guid.S())
	s.SetNameToUriType(ghttp.UriTypeCamel)
	s.BindObject("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.GetContent(ctx, "/namesObject/showName"), "Object Show Name")
	})
}

func Test_NameToUri_Default(t *testing.T) {
	s := g.Server(guid.S())
	s.SetNameToUriType(ghttp.UriTypeDefault)
	s.BindObject("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.GetContent(ctx, "/names-object/show-name"), "Object Show Name")
	})
}
