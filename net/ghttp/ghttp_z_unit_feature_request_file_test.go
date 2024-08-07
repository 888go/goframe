// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/888go/goframe/internal/json"

	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gfile "github.com/888go/goframe/os/gfile"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
	gmeta "github.com/888go/goframe/util/gmeta"
	guid "github.com/888go/goframe/util/guid"
)

func Test_Params_File_Single(t *testing.T) {
	dstDirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
	s := g.Http类(guid.X生成())
	s.X绑定("/upload/single", func(r *ghttp.Request) {
		file := r.X取上传文件对象("file")
		if file == nil {
			r.X响应.X写响应缓冲区并退出("upload file cannot be empty")
		}

		if name, err := file.X保存(dstDirPath, r.Get别名("randomlyRename").X取布尔()); err == nil {
			r.X响应.X写响应缓冲区并退出(name)
		}
		r.X响应.X写响应缓冲区并退出("upload failed")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// normal name
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath := gtest.DataPath("upload", "file1.txt")
		dstPath := gfile.X路径生成(dstDirPath, "file1.txt")
		content := client.Post文本(ctx, "/upload/single", g.Map{
			"file": "@file:" + srcPath,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt")
		t.Assert(gfile.X读文本(dstPath), gfile.X读文本(srcPath))
	})
	// randomly rename.
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath := gtest.DataPath("upload", "file2.txt")
		content := client.Post文本(ctx, "/upload/single", g.Map{
			"file":           "@file:" + srcPath,
			"randomlyRename": true,
		})
		dstPath := gfile.X路径生成(dstDirPath, content)
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(gfile.X读文本(dstPath), gfile.X读文本(srcPath))
	})
}

func Test_Params_File_CustomName(t *testing.T) {
	dstDirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
	s := g.Http类(guid.X生成())
	s.X绑定("/upload/single", func(r *ghttp.Request) {
		file := r.X取上传文件对象("file")
		if file == nil {
			r.X响应.X写响应缓冲区并退出("upload file cannot be empty")
		}
		file.Filename = "my.txt"
		if name, err := file.X保存(dstDirPath, r.Get别名("randomlyRename").X取布尔()); err == nil {
			r.X响应.X写响应缓冲区并退出(name)
		}
		r.X响应.X写响应缓冲区并退出("upload failed")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath := gtest.DataPath("upload", "file1.txt")
		dstPath := gfile.X路径生成(dstDirPath, "my.txt")
		content := client.Post文本(ctx, "/upload/single", g.Map{
			"file": "@file:" + srcPath,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "my.txt")
		t.Assert(gfile.X读文本(dstPath), gfile.X读文本(srcPath))
	})
}

func Test_Params_File_Batch(t *testing.T) {
	dstDirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
	s := g.Http类(guid.X生成())
	s.X绑定("/upload/batch", func(r *ghttp.Request) {
		files := r.X取上传文件切片对象("file")
		if files == nil {
			r.X响应.X写响应缓冲区并退出("upload file cannot be empty")
		}
		if names, err := files.X保存(dstDirPath, r.Get别名("randomlyRename").X取布尔()); err == nil {
			r.X响应.X写响应缓冲区并退出(gstr.X连接(names, ","))
		}
		r.X响应.X写响应缓冲区并退出("upload failed")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// normal name
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath1 := gtest.DataPath("upload", "file1.txt")
		srcPath2 := gtest.DataPath("upload", "file2.txt")
		dstPath1 := gfile.X路径生成(dstDirPath, "file1.txt")
		dstPath2 := gfile.X路径生成(dstDirPath, "file2.txt")
		content := client.Post文本(ctx, "/upload/batch", g.Map{
			"file[0]": "@file:" + srcPath1,
			"file[1]": "@file:" + srcPath2,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt,file2.txt")
		t.Assert(gfile.X读文本(dstPath1), gfile.X读文本(srcPath1))
		t.Assert(gfile.X读文本(dstPath2), gfile.X读文本(srcPath2))
	})
	// randomly rename.
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath1 := gtest.DataPath("upload", "file1.txt")
		srcPath2 := gtest.DataPath("upload", "file2.txt")
		content := client.Post文本(ctx, "/upload/batch", g.Map{
			"file[0]":        "@file:" + srcPath1,
			"file[1]":        "@file:" + srcPath2,
			"randomlyRename": true,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")

		array := gstr.X分割并忽略空值(content, ",")
		t.Assert(len(array), 2)
		dstPath1 := gfile.X路径生成(dstDirPath, array[0])
		dstPath2 := gfile.X路径生成(dstDirPath, array[1])
		t.Assert(gfile.X读文本(dstPath1), gfile.X读文本(srcPath1))
		t.Assert(gfile.X读文本(dstPath2), gfile.X读文本(srcPath2))
	})
}

func Test_Params_Strict_Route_File_Single_Ptr_Attrr(t *testing.T) {
	type Req struct {
		gmeta.Meta `method:"post" mime:"multipart/form-data"`
		File       *ghttp.UploadFile `type:"file"`
	}
	type Res struct{}

	dstDirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
	s := g.Http类(guid.X生成())
	s.X绑定("/upload/single", func(ctx context.Context, req *Req) (res *Res, err error) {
		var (
			r    = g.Http类上下文取请求对象(ctx)
			file = req.File
		)
		if file == nil {
			r.X响应.X写响应缓冲区并退出("upload file cannot be empty")
		}
		name, err := file.X保存(dstDirPath)
		if err != nil {
			r.X响应.X写响应缓冲区并退出(err)
		}
		r.X响应.X写响应缓冲区并退出(name)
		return
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// normal name
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath := gtest.DataPath("upload", "file1.txt")
		dstPath := gfile.X路径生成(dstDirPath, "file1.txt")
		content := client.Post文本(ctx, "/upload/single", g.Map{
			"file": "@file:" + srcPath,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt")
		t.Assert(gfile.X读文本(dstPath), gfile.X读文本(srcPath))
	})
}

func Test_Params_Strict_Route_File_Single_Struct_Attr(t *testing.T) {
	type Req struct {
		gmeta.Meta `method:"post" mime:"multipart/form-data"`
		File       ghttp.UploadFile `type:"file"`
	}
	type Res struct{}

	dstDirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
	s := g.Http类(guid.X生成())
	s.X绑定("/upload/single", func(ctx context.Context, req *Req) (res *Res, err error) {
		var (
			r    = g.Http类上下文取请求对象(ctx)
			file = req.File
		)
		name, err := file.X保存(dstDirPath)
		if err != nil {
			r.X响应.X写响应缓冲区并退出(err)
		}
		r.X响应.X写响应缓冲区并退出(name)
		return
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// normal name
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath := gtest.DataPath("upload", "file1.txt")
		dstPath := gfile.X路径生成(dstDirPath, "file1.txt")
		content := client.Post文本(ctx, "/upload/single", g.Map{
			"file": "@file:" + srcPath,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt")
		t.Assert(gfile.X读文本(dstPath), gfile.X读文本(srcPath))
	})
}

func Test_Params_File_Upload_Required(t *testing.T) {
	type Req struct {
		gmeta.Meta `method:"post" mime:"multipart/form-data"`
		File       *ghttp.UploadFile `type:"file" v:"required#upload file is required"`
	}
	type Res struct{}

	s := g.Http类(guid.X生成())
	s.Use别名(ghttp.MiddlewareHandlerResponse)
	s.X绑定("/upload/required", func(ctx context.Context, req *Req) (res *Res, err error) {
		return
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// file is empty
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		content := client.Post文本(ctx, "/upload/required")
		t.Assert(content, `{"code":51,"message":"upload file is required","data":null}`)
	})
}

func Test_Params_File_MarshalJSON(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/upload/single", func(r *ghttp.Request) {
		file := r.X取上传文件对象("file")
		if file == nil {
			r.X响应.X写响应缓冲区并退出("upload file cannot be empty")
		}

		if bytes, err := json.Marshal(file); err != nil {
			r.X响应.X写响应缓冲区并退出(err)
		} else {
			r.X响应.X写响应缓冲区并退出(bytes)
		}
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// normal name
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath := gtest.DataPath("upload", "file1.txt")
		content := client.Post文本(ctx, "/upload/single", g.Map{
			"file": "@file:" + srcPath,
		})
		t.Assert(strings.Contains(content, "file1.txt"), true)
	})
}

// 批量上传时只选择一个文件. md5:e662751ee0c7888c
func Test_Params_Strict_Route_File_Batch_Up_One(t *testing.T) {
	type Req struct {
		gmeta.Meta `method:"post" mime:"multipart/form-data"`
		Files      ghttp.UploadFiles `type:"file"`
	}
	type Res struct{}

	dstDirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
	s := g.Http类(guid.X生成())
	s.X绑定("/upload/batch", func(ctx context.Context, req *Req) (res *Res, err error) {
		var (
			r     = g.Http类上下文取请求对象(ctx)
			files = req.Files
		)
		if len(files) == 0 {
			r.X响应.X写响应缓冲区并退出("upload file cannot be empty")
		}
		names, err := files.X保存(dstDirPath)
		if err != nil {
			r.X响应.X写响应缓冲区并退出(err)
		}
		r.X响应.X写响应缓冲区并退出(gstr.X连接(names, ","))
		return
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// normal name
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath := gtest.DataPath("upload", "file1.txt")
		dstPath := gfile.X路径生成(dstDirPath, "file1.txt")
		content := client.Post文本(ctx, "/upload/batch", g.Map{
			"files": "@file:" + srcPath,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt")
		t.Assert(gfile.X读文本(dstPath), gfile.X读文本(srcPath))
	})
}

// 批量上传时选择多个文件. md5:b38cd6734ce32c09
func Test_Params_Strict_Route_File_Batch_Up_Multiple(t *testing.T) {
	type Req struct {
		gmeta.Meta `method:"post" mime:"multipart/form-data"`
		Files      ghttp.UploadFiles `type:"file"`
	}
	type Res struct{}

	dstDirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
	s := g.Http类(guid.X生成())
	s.X绑定("/upload/batch", func(ctx context.Context, req *Req) (res *Res, err error) {
		var (
			r     = g.Http类上下文取请求对象(ctx)
			files = req.Files
		)
		if len(files) == 0 {
			r.X响应.X写响应缓冲区并退出("upload file cannot be empty")
		}
		names, err := files.X保存(dstDirPath)
		if err != nil {
			r.X响应.X写响应缓冲区并退出(err)
		}
		r.X响应.X写响应缓冲区并退出(gstr.X连接(names, ","))
		return
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// normal name
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath1 := gtest.DataPath("upload", "file1.txt")
		srcPath2 := gtest.DataPath("upload", "file2.txt")
		dstPath1 := gfile.X路径生成(dstDirPath, "file1.txt")
		dstPath2 := gfile.X路径生成(dstDirPath, "file2.txt")
		content := client.Post文本(ctx, "/upload/batch",
			"files=@file:"+srcPath1+
				"&files=@file:"+srcPath2,
		)
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt,file2.txt")
		t.Assert(gfile.X读文本(dstPath1), gfile.X读文本(srcPath1))
		t.Assert(gfile.X读文本(dstPath2), gfile.X读文本(srcPath2))
	})
}
