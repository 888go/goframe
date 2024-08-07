// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	gxml "github.com/888go/goframe/encoding/gxml"
	"github.com/888go/goframe/internal/json"
	gview "github.com/888go/goframe/os/gview"

	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func Test_Response_ServeFile(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/ServeFile", func(r *ghttp.Request) {
		filePath := r.X取查询参数到泛型类("filePath")
		r.X响应.X发送文件(filePath.String())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		srcPath := gtest.DataPath("upload", "file1.txt")
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
	s := g.Http类(guid.X生成())
	s.X绑定("/ServeFileDownload", func(r *ghttp.Request) {
		filePath := r.X取查询参数到泛型类("filePath")
		r.X响应.X下载文件(filePath.String())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		srcPath := gtest.DataPath("upload", "file1.txt")
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
	s := g.Http类(guid.X生成())
	s.X绑定("/", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("RedirectResult")
	})
	s.X绑定("/RedirectTo", func(r *ghttp.Request) {
		r.X响应.X重定向("/")
	})
	s.X绑定("/RedirectTo301", func(r *ghttp.Request) {
		r.X响应.X重定向("/", http.StatusMovedPermanently)
	})
	s.X绑定("/RedirectBack", func(r *ghttp.Request) {
		r.X响应.X重定向到来源页面()
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/RedirectTo"), "RedirectResult")
		t.Assert(client.Get文本(ctx, "/RedirectTo301"), "RedirectResult")
		t.Assert(client.X设置协议头("Referer", "/").Get文本(ctx, "/RedirectBack"), "RedirectResult")
	})
}

func Test_Response_Buffer(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/Buffer", func(r *ghttp.Request) {
		name := r.X取查询参数到泛型类("name").X取字节集()
		r.X响应.SetBuffer(name)
		buffer := r.X响应.Buffer()
		r.X响应.ClearBuffer()
		r.X响应.X写响应缓冲区(buffer)
	})
	s.X绑定("/BufferString", func(r *ghttp.Request) {
		name := r.X取查询参数到泛型类("name").X取字节集()
		r.X响应.SetBuffer(name)
		bufferString := r.X响应.BufferString()
		r.X响应.ClearBuffer()
		r.X响应.X写响应缓冲区(bufferString)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/Buffer", "name=john"), []byte("john"))
		t.Assert(client.Get文本(ctx, "/BufferString", "name=john"), "john")
	})
}

func Test_Response_WriteTpl(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gview.New(gtest.DataPath("template", "basic"))
		s := g.Http类(guid.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/", func(r *ghttp.Request) {
			err := r.X响应.X输出到模板文件("noexist.html", g.Map{
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
	gtest.C(t, func(t *gtest.T) {
		v := gview.New()
		v.SetDefaultFile(gtest.DataPath("template", "basic", "index.html"))
		s := g.Http类(guid.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/", func(r *ghttp.Request) {
			err := r.X响应.X输出到默认模板文件(g.Map{"name": "john"})
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
	gtest.C(t, func(t *gtest.T) {
		v := gview.New()
		v.SetDefaultFile(gtest.DataPath("template", "basic", "noexit.html"))
		s := g.Http类(guid.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/", func(r *ghttp.Request) {
			err := r.X响应.X输出到默认模板文件(g.Map{"name": "john"})
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
	gtest.C(t, func(t *gtest.T) {
		v := gview.New()
		v.SetDefaultFile(gtest.DataPath("template", "basic", "index.html"))
		s := g.Http类(guid.X生成())
		s.X设置默认模板对象(v)
		s.X绑定("/", func(r *ghttp.Request) {
			res, err := r.X响应.X解析默认模板文件(g.Map{"name": "john"})
			t.AssertNil(err)
			r.X响应.X写响应缓冲区(res)
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
	s := g.Http类(guid.X生成())
	s.X绑定("/", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区()
	})
	s.X绑定("/WriteOverExit", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("WriteOverExit")
		r.X响应.X写覆盖响应缓冲区并退出("")
	})
	s.X绑定("/WritefExit", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区并退出与格式化("%s", "WritefExit")
	})
	s.X绑定("/Writeln", func(r *ghttp.Request) {
		name := r.X取查询参数到泛型类("name")
		r.X响应.X写响应缓冲区并换行(name)
	})
	s.X绑定("/WritelnNil", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区并换行()
	})
	s.X绑定("/Writefln", func(r *ghttp.Request) {
		name := r.X取查询参数到泛型类("name")
		r.X响应.X写响应缓冲区并格式化与换行("%s", name)
	})
	s.X绑定("/WriteJson", func(r *ghttp.Request) {
		m := map[string]string{"name": "john"}
		if bytes, err := json.Marshal(m); err == nil {
			r.X响应.X写响应缓冲区JSON(bytes)
		}
	})
	s.X绑定("/WriteJsonP", func(r *ghttp.Request) {
		m := map[string]string{"name": "john"}
		if bytes, err := json.Marshal(m); err == nil {
			r.X响应.X写响应缓冲区JSONP(bytes)
		}
	})
	s.X绑定("/WriteJsonPWithStruct", func(r *ghttp.Request) {
		user := User{"john"}
		r.X响应.X写响应缓冲区JSONP(user)
	})
	s.X绑定("/WriteXml", func(r *ghttp.Request) {
		m := map[string]interface{}{"name": "john"}
		if bytes, err := gxml.Encode(m); err == nil {
			r.X响应.X写响应缓冲区XML(bytes)
		}
	})
	s.X绑定("/WriteXmlWithStruct", func(r *ghttp.Request) {
		user := User{"john"}
		r.X响应.X写响应缓冲区XML(user)
	})

	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
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
