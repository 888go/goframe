// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

import (
	"fmt"
	"github.com/888go/goframe/encoding/gxml"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/os/gview"
	"net/http"
	"strings"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

func Test_Response_ServeFile(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/ServeFile", func(r *http类.Request) {
		filePath := r.X取查询参数到泛型类("filePath")
		r.Response.X发送文件(filePath.String())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		srcPath := 单元测试类.DataPath("upload", "file1.txt")
		t.Assert(client.Get文本(ctx, "/ServeFile", "filePath=file1.txt"), "Not Found")

		t.Assert(
			client.Get文本(ctx, "/ServeFile", "filePath="+srcPath),
			"file1.txt: This file is for uploading unit test case.")

		t.Assert(
			strings.Contains(
				client.Get文本(ctx, "/ServeFile", "filePath=files/server.key"),
				"BEGIN RSA PRIVATE KEY"),
			true)
	})
}

func Test_Response_ServeFileDownload(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/ServeFileDownload", func(r *http类.Request) {
		filePath := r.X取查询参数到泛型类("filePath")
		r.Response.X下载文件(filePath.String())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		srcPath := 单元测试类.DataPath("upload", "file1.txt")
		t.Assert(client.Get文本(ctx, "/ServeFileDownload", "filePath=file1.txt"), "Not Found")

		t.Assert(
			client.Get文本(ctx, "/ServeFileDownload", "filePath="+srcPath),
			"file1.txt: This file is for uploading unit test case.")

		t.Assert(
			strings.Contains(
				client.Get文本(ctx, "/ServeFileDownload", "filePath=files/server.key"),
				"BEGIN RSA PRIVATE KEY"),
			true)
	})
}

func Test_Response_Redirect(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/", func(r *http类.Request) {
		r.Response.X写响应缓冲区("RedirectResult")
	})
	s.X绑定("/RedirectTo", func(r *http类.Request) {
		r.Response.X重定向("/")
	})
	s.X绑定("/RedirectTo301", func(r *http类.Request) {
		r.Response.X重定向("/", http.StatusMovedPermanently)
	})
	s.X绑定("/RedirectBack", func(r *http类.Request) {
		r.Response.X重定向到来源页面()
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/RedirectTo"), "RedirectResult")
		t.Assert(client.Get文本(ctx, "/RedirectTo301"), "RedirectResult")
		t.Assert(client.X设置协议头("Referer", "/").Get文本(ctx, "/RedirectBack"), "RedirectResult")
	})
}

func Test_Response_Buffer(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/Buffer", func(r *http类.Request) {
		name := r.X取查询参数到泛型类("name").X取字节集()
		r.Response.X设置缓冲区字节集(name)
		buffer := r.Response.X取缓冲区字节集()
		r.Response.X清空缓冲区()
		r.Response.X写响应缓冲区(buffer)
	})
	s.X绑定("/BufferString", func(r *http类.Request) {
		name := r.X取查询参数到泛型类("name").X取字节集()
		r.Response.X设置缓冲区字节集(name)
		bufferString := r.Response.X取缓冲区文本()
		r.Response.X清空缓冲区()
		r.Response.X写响应缓冲区(bufferString)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/Buffer", "name=john"), []byte("john"))
		t.Assert(client.Get文本(ctx, "/BufferString", "name=john"), "john")
	})
}

func Test_Response_WriteTpl(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		v := 模板类.New(单元测试类.DataPath("template", "basic"))
		s := g.Http类(uid类.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/", func(r *http类.Request) {
			err := r.Response.X输出到模板文件("noexist.html", g.Map{
				"name": "john",
			})
			t.AssertNE(err, nil)
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.AssertNE(client.Get文本(ctx, "/"), "Name:john")
	})
}

func Test_Response_WriteTplDefault(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		v := 模板类.New()
		v.SetDefaultFile(单元测试类.DataPath("template", "basic", "index.html"))
		s := g.Http类(uid类.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/", func(r *http类.Request) {
			err := r.Response.X输出到默认模板文件(g.Map{"name": "john"})
			t.AssertNil(err)
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Name:john")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		v := 模板类.New()
		v.SetDefaultFile(单元测试类.DataPath("template", "basic", "noexit.html"))
		s := g.Http类(uid类.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/", func(r *http类.Request) {
			err := r.Response.X输出到默认模板文件(g.Map{"name": "john"})
			t.AssertNil(err)
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.AssertNE(client.Get文本(ctx, "/"), "Name:john")
	})
}

func Test_Response_ParseTplDefault(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		v := 模板类.New()
		v.SetDefaultFile(单元测试类.DataPath("template", "basic", "index.html"))
		s := g.Http类(uid类.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/", func(r *http类.Request) {
			res, err := r.Response.X解析默认模板文件(g.Map{"name": "john"})
			t.AssertNil(err)
			r.Response.X写响应缓冲区(res)
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "Name:john")
	})
}

func Test_Response_Write(t *testing.T) {
	type User struct {
		Name string `json:"name"`
	}
	s := g.Http类(uid类.X生成())
	s.X绑定("/", func(r *http类.Request) {
		r.Response.X写响应缓冲区()
	})
	s.X绑定("/WriteOverExit", func(r *http类.Request) {
		r.Response.X写响应缓冲区("WriteOverExit")
		r.Response.X写覆盖响应缓冲区并退出("")
	})
	s.X绑定("/WritefExit", func(r *http类.Request) {
		r.Response.X写响应缓冲区并退出与格式化("%s", "WritefExit")
	})
	s.X绑定("/Writeln", func(r *http类.Request) {
		name := r.X取查询参数到泛型类("name")
		r.Response.X写响应缓冲区并换行(name)
	})
	s.X绑定("/WritelnNil", func(r *http类.Request) {
		r.Response.X写响应缓冲区并换行()
	})
	s.X绑定("/Writefln", func(r *http类.Request) {
		name := r.X取查询参数到泛型类("name")
		r.Response.X写响应缓冲区并格式化与换行("%s", name)
	})
	s.X绑定("/WriteJson", func(r *http类.Request) {
		m := map[string]string{"name": "john"}
		if bytes, err := json.Marshal(m); err == nil {
			r.Response.X写响应缓冲区JSON(bytes)
		}
	})
	s.X绑定("/WriteJsonP", func(r *http类.Request) {
		m := map[string]string{"name": "john"}
		if bytes, err := json.Marshal(m); err == nil {
			r.Response.X写响应缓冲区JSONP(bytes)
		}
	})
	s.X绑定("/WriteJsonPWithStruct", func(r *http类.Request) {
		user := User{"john"}
		r.Response.X写响应缓冲区JSONP(user)
	})
	s.X绑定("/WriteXml", func(r *http类.Request) {
		m := map[string]interface{}{"name": "john"}
		if bytes, err := xml类.Encode(m); err == nil {
			r.Response.X写响应缓冲区XML(bytes)
		}
	})
	s.X绑定("/WriteXmlWithStruct", func(r *http类.Request) {
		user := User{"john"}
		r.Response.X写响应缓冲区XML(user)
	})

	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), "")
		t.Assert(client.Get文本(ctx, "/WriteOverExit"), "")
		t.Assert(client.Get文本(ctx, "/WritefExit"), "WritefExit")
		t.Assert(client.Get文本(ctx, "/Writeln"), "\n")
		t.Assert(client.Get文本(ctx, "/WritelnNil"), "\n")
		t.Assert(client.Get文本(ctx, "/Writeln", "name=john"), "john\n")
		t.Assert(client.Get文本(ctx, "/Writefln", "name=john"), "john\n")
		t.Assert(client.Get文本(ctx, "/WriteJson"), "{\"name\":\"john\"}")
		t.Assert(client.Get文本(ctx, "/WriteJsonP"), "{\"name\":\"john\"}")
		t.Assert(client.Get文本(ctx, "/WriteJsonPWithStruct"), "{\"name\":\"john\"}")
		t.Assert(client.Get文本(ctx, "/WriteJsonPWithStruct", "callback=callback"),
			"callback({\"name\":\"john\"})")
		t.Assert(client.Get文本(ctx, "/WriteXml"), "<name>john</name>")
		t.Assert(client.Get文本(ctx, "/WriteXmlWithStruct"), "<name>john</name>")
	})
}
