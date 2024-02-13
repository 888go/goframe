// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 静态服务测试。

package http类_test

import (
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/guid"
)

func Test_Static_ServerRoot(t *testing.T) {
	// SetServerRoot 使用绝对路径
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		path := fmt.Sprintf(`%s/ghttp/static/test/%d`, 文件类.X取临时目录(), s.X取已监听端口())
		defer 文件类.X删除(path)
		文件类.X写入文本(path+"/index.htm", "index")
		s.X设置静态文件根目录(path)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "index")
		t.Assert(client.Get文本(ctx, "/index.htm"), "index")
	})

	// SetServerRoot 使用相对路径
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		path := fmt.Sprintf(`static/test/%d`, s.X取已监听端口())
		defer 文件类.X删除(path)
		文件类.X写入文本(path+"/index.htm", "index")
		s.X设置静态文件根目录(path)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "index")
		t.Assert(client.Get文本(ctx, "/index.htm"), "index")
	})
}

func Test_Static_ServerRoot_Security(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		s.X设置静态文件根目录(单元测试类.DataPath("static1"))
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "index")
		t.Assert(client.Get文本(ctx, "/index.htm"), "Not Found")
		t.Assert(client.Get文本(ctx, "/index.html"), "index")
		t.Assert(client.Get文本(ctx, "/test.html"), "test")
		t.Assert(client.Get文本(ctx, "/../main.html"), "Not Found")
		t.Assert(client.Get文本(ctx, "/..%2Fmain.html"), "Not Found")
	})
}

func Test_Static_Folder_Forbidden(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		path := fmt.Sprintf(`%s/ghttp/static/test/%d`, 文件类.X取临时目录(), s.X取已监听端口())
		defer 文件类.X删除(path)
		文件类.X写入文本(path+"/test.html", "test")
		s.X设置静态文件根目录(path)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Forbidden")
		t.Assert(client.Get文本(ctx, "/index.html"), "Not Found")
		t.Assert(client.Get文本(ctx, "/test.html"), "test")
	})
}

func Test_Static_IndexFolder(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		path := fmt.Sprintf(`%s/ghttp/static/test/%d`, 文件类.X取临时目录(), s.X取已监听端口())
		defer 文件类.X删除(path)
		文件类.X写入文本(path+"/test.html", "test")
		s.X设置静态文件是否列出子文件(true)
		s.X设置静态文件根目录(path)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.AssertNE(client.Get文本(ctx, "/"), "Forbidden")
		t.AssertNE(文本类.X查找(client.Get文本(ctx, "/"), `<a href="/test.html"`), -1)
		t.Assert(client.Get文本(ctx, "/index.html"), "Not Found")
		t.Assert(client.Get文本(ctx, "/test.html"), "test")
	})
}

func Test_Static_IndexFiles1(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		path := fmt.Sprintf(`%s/ghttp/static/test/%d`, 文件类.X取临时目录(), s.X取已监听端口())
		defer 文件类.X删除(path)
		文件类.X写入文本(path+"/index.html", "index")
		文件类.X写入文本(path+"/test.html", "test")
		s.X设置静态文件根目录(path)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "index")
		t.Assert(client.Get文本(ctx, "/index.html"), "index")
		t.Assert(client.Get文本(ctx, "/test.html"), "test")
	})
}

func Test_Static_IndexFiles2(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		path := fmt.Sprintf(`%s/ghttp/static/test/%d`, 文件类.X取临时目录(), s.X取已监听端口())
		defer 文件类.X删除(path)
		文件类.X写入文本(path+"/test.html", "test")
		s.X设置静态文件索引([]string{"index.html", "test.html"})
		s.X设置静态文件根目录(path)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "test")
		t.Assert(client.Get文本(ctx, "/index.html"), "Not Found")
		t.Assert(client.Get文本(ctx, "/test.html"), "test")
	})
}

func Test_Static_AddSearchPath1(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		path1 := fmt.Sprintf(`%s/ghttp/static/test/%d`, 文件类.X取临时目录(), s.X取已监听端口())
		path2 := fmt.Sprintf(`%s/ghttp/static/test/%d/%d`, 文件类.X取临时目录(), s.X取已监听端口(), s.X取已监听端口())
		defer 文件类.X删除(path1)
		defer 文件类.X删除(path2)
		文件类.X写入文本(path2+"/test.html", "test")
		s.X设置静态文件根目录(path1)
		s.X静态文件添加额外搜索目录(path2)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Forbidden")
		t.Assert(client.Get文本(ctx, "/test.html"), "test")
	})
}

func Test_Static_AddSearchPath2(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		path1 := fmt.Sprintf(`%s/ghttp/static/test/%d`, 文件类.X取临时目录(), s.X取已监听端口())
		path2 := fmt.Sprintf(`%s/ghttp/static/test/%d/%d`, 文件类.X取临时目录(), s.X取已监听端口(), s.X取已监听端口())
		defer 文件类.X删除(path1)
		defer 文件类.X删除(path2)
		文件类.X写入文本(path1+"/test.html", "test1")
		文件类.X写入文本(path2+"/test.html", "test2")
		s.X设置静态文件根目录(path1)
		s.X静态文件添加额外搜索目录(path2)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Forbidden")
		t.Assert(client.Get文本(ctx, "/test.html"), "test1")
	})
}

func Test_Static_AddStaticPath(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		path1 := fmt.Sprintf(`%s/ghttp/static/test/%d`, 文件类.X取临时目录(), s.X取已监听端口())
		path2 := fmt.Sprintf(`%s/ghttp/static/test/%d/%d`, 文件类.X取临时目录(), s.X取已监听端口(), s.X取已监听端口())
		defer 文件类.X删除(path1)
		defer 文件类.X删除(path2)
		文件类.X写入文本(path1+"/test.html", "test1")
		文件类.X写入文本(path2+"/test.html", "test2")
		s.X设置静态文件根目录(path1)
		s.X静态文件添加目录映射("/my-test", path2)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Forbidden")
		t.Assert(client.Get文本(ctx, "/test.html"), "test1")
		t.Assert(client.Get文本(ctx, "/my-test/test.html"), "test2")
	})
}

func Test_Static_AddStaticPath_Priority(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		path1 := fmt.Sprintf(`%s/ghttp/static/test/%d/test`, 文件类.X取临时目录(), s.X取已监听端口())
		path2 := fmt.Sprintf(`%s/ghttp/static/test/%d/%d/test`, 文件类.X取临时目录(), s.X取已监听端口(), s.X取已监听端口())
		defer 文件类.X删除(path1)
		defer 文件类.X删除(path2)
		文件类.X写入文本(path1+"/test.html", "test1")
		文件类.X写入文本(path2+"/test.html", "test2")
		s.X设置静态文件根目录(path1)
		s.X静态文件添加目录映射("/test", path2)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Forbidden")
		t.Assert(client.Get文本(ctx, "/test.html"), "test1")
		t.Assert(client.Get文本(ctx, "/test/test.html"), "test2")
	})
}

func Test_Static_Rewrite(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		path := fmt.Sprintf(`%s/ghttp/static/test/%d`, 文件类.X取临时目录(), s.X取已监听端口())
		defer 文件类.X删除(path)
		文件类.X写入文本(path+"/test1.html", "test1")
		文件类.X写入文本(path+"/test2.html", "test2")
		s.X设置静态文件根目录(path)
		s.X设置路由URI重写规则("/test.html", "/test1.html")
		s.X设置路由URI重写规则Map(g.MapStrStr{
			"/my-test1": "/test1.html",
			"/my-test2": "/test2.html",
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Forbidden")
		t.Assert(client.Get文本(ctx, "/test.html"), "test1")
		t.Assert(client.Get文本(ctx, "/test1.html"), "test1")
		t.Assert(client.Get文本(ctx, "/test2.html"), "test2")
		t.Assert(client.Get文本(ctx, "/my-test1"), "test1")
		t.Assert(client.Get文本(ctx, "/my-test2"), "test2")
	})
}
