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
