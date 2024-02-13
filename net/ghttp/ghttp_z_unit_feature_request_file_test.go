// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
	
	"github.com/888go/goframe/internal/json"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gmeta"
	"github.com/888go/goframe/util/guid"
)

func Test_Params_File_Single(t *testing.T) {
	dstDirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
	s := g.Http类(uid类.X生成())
	s.X绑定("/upload/single", func(r *http类.Request) {
		file := r.X取上传文件对象("file")
		if file == nil {
			r.Response.X写响应缓冲区并退出("upload file cannot be empty")
		}

		if name, err := file.X保存(dstDirPath, r.Get别名("randomlyRename").X取布尔()); err == nil {
			r.Response.X写响应缓冲区并退出(name)
		}
		r.Response.X写响应缓冲区并退出("upload failed")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// normal name
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath := 单元测试类.DataPath("upload", "file1.txt")
		dstPath := 文件类.X路径生成(dstDirPath, "file1.txt")
		content := client.Post文本(ctx, "/upload/single", g.Map{
			"file": "@file:" + srcPath,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt")
		t.Assert(文件类.X读文本(dstPath), 文件类.X读文本(srcPath))
	})
	// randomly rename.
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath := 单元测试类.DataPath("upload", "file2.txt")
		content := client.Post文本(ctx, "/upload/single", g.Map{
			"file":           "@file:" + srcPath,
			"randomlyRename": true,
		})
		dstPath := 文件类.X路径生成(dstDirPath, content)
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(文件类.X读文本(dstPath), 文件类.X读文本(srcPath))
	})
}

func Test_Params_File_CustomName(t *testing.T) {
	dstDirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
	s := g.Http类(uid类.X生成())
	s.X绑定("/upload/single", func(r *http类.Request) {
		file := r.X取上传文件对象("file")
		if file == nil {
			r.Response.X写响应缓冲区并退出("upload file cannot be empty")
		}
		file.Filename = "my.txt"
		if name, err := file.X保存(dstDirPath, r.Get别名("randomlyRename").X取布尔()); err == nil {
			r.Response.X写响应缓冲区并退出(name)
		}
		r.Response.X写响应缓冲区并退出("upload failed")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath := 单元测试类.DataPath("upload", "file1.txt")
		dstPath := 文件类.X路径生成(dstDirPath, "my.txt")
		content := client.Post文本(ctx, "/upload/single", g.Map{
			"file": "@file:" + srcPath,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "my.txt")
		t.Assert(文件类.X读文本(dstPath), 文件类.X读文本(srcPath))
	})
}

func Test_Params_File_Batch(t *testing.T) {
	dstDirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
	s := g.Http类(uid类.X生成())
	s.X绑定("/upload/batch", func(r *http类.Request) {
		files := r.X取上传文件数组对象("file")
		if files == nil {
			r.Response.X写响应缓冲区并退出("upload file cannot be empty")
		}
		if names, err := files.X保存(dstDirPath, r.Get别名("randomlyRename").X取布尔()); err == nil {
			r.Response.X写响应缓冲区并退出(文本类.X连接(names, ","))
		}
		r.Response.X写响应缓冲区并退出("upload failed")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// normal name
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath1 := 单元测试类.DataPath("upload", "file1.txt")
		srcPath2 := 单元测试类.DataPath("upload", "file2.txt")
		dstPath1 := 文件类.X路径生成(dstDirPath, "file1.txt")
		dstPath2 := 文件类.X路径生成(dstDirPath, "file2.txt")
		content := client.Post文本(ctx, "/upload/batch", g.Map{
			"file[0]": "@file:" + srcPath1,
			"file[1]": "@file:" + srcPath2,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt,file2.txt")
		t.Assert(文件类.X读文本(dstPath1), 文件类.X读文本(srcPath1))
		t.Assert(文件类.X读文本(dstPath2), 文件类.X读文本(srcPath2))
	})
	// randomly rename.
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath1 := 单元测试类.DataPath("upload", "file1.txt")
		srcPath2 := 单元测试类.DataPath("upload", "file2.txt")
		content := client.Post文本(ctx, "/upload/batch", g.Map{
			"file[0]":        "@file:" + srcPath1,
			"file[1]":        "@file:" + srcPath2,
			"randomlyRename": true,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")

		array := 文本类.X分割并忽略空值(content, ",")
		t.Assert(len(array), 2)
		dstPath1 := 文件类.X路径生成(dstDirPath, array[0])
		dstPath2 := 文件类.X路径生成(dstDirPath, array[1])
		t.Assert(文件类.X读文本(dstPath1), 文件类.X读文本(srcPath1))
		t.Assert(文件类.X读文本(dstPath2), 文件类.X读文本(srcPath2))
	})
}

func Test_Params_Strict_Route_File_Single_Ptr_Attrr(t *testing.T) {
	type Req struct {
		元数据类.Meta `method:"post" mime:"multipart/form-data"`
		File       *http类.UploadFile `type:"file"`
	}
	type Res struct{}

	dstDirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
	s := g.Http类(uid类.X生成())
	s.X绑定("/upload/single", func(ctx context.Context, req *Req) (res *Res, err error) {
		var (
			r    = g.Http类上下文取请求对象(ctx)
			file = req.File
		)
		if file == nil {
			r.Response.X写响应缓冲区并退出("upload file cannot be empty")
		}
		name, err := file.X保存(dstDirPath)
		if err != nil {
			r.Response.X写响应缓冲区并退出(err)
		}
		r.Response.X写响应缓冲区并退出(name)
		return
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// normal name
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath := 单元测试类.DataPath("upload", "file1.txt")
		dstPath := 文件类.X路径生成(dstDirPath, "file1.txt")
		content := client.Post文本(ctx, "/upload/single", g.Map{
			"file": "@file:" + srcPath,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt")
		t.Assert(文件类.X读文本(dstPath), 文件类.X读文本(srcPath))
	})
}

func Test_Params_Strict_Route_File_Single_Struct_Attr(t *testing.T) {
	type Req struct {
		元数据类.Meta `method:"post" mime:"multipart/form-data"`
		File       http类.UploadFile `type:"file"`
	}
	type Res struct{}

	dstDirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
	s := g.Http类(uid类.X生成())
	s.X绑定("/upload/single", func(ctx context.Context, req *Req) (res *Res, err error) {
		var (
			r    = g.Http类上下文取请求对象(ctx)
			file = req.File
		)
		name, err := file.X保存(dstDirPath)
		if err != nil {
			r.Response.X写响应缓冲区并退出(err)
		}
		r.Response.X写响应缓冲区并退出(name)
		return
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// normal name
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath := 单元测试类.DataPath("upload", "file1.txt")
		dstPath := 文件类.X路径生成(dstDirPath, "file1.txt")
		content := client.Post文本(ctx, "/upload/single", g.Map{
			"file": "@file:" + srcPath,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt")
		t.Assert(文件类.X读文本(dstPath), 文件类.X读文本(srcPath))
	})
}

func Test_Params_File_Upload_Required(t *testing.T) {
	type Req struct {
		元数据类.Meta `method:"post" mime:"multipart/form-data"`
		File       *http类.UploadFile `type:"file" v:"required#upload file is required"`
	}
	type Res struct{}

	s := g.Http类(uid类.X生成())
	s.Use别名(http类.MiddlewareHandlerResponse)
	s.X绑定("/upload/required", func(ctx context.Context, req *Req) (res *Res, err error) {
		return
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// file is empty
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		content := client.Post文本(ctx, "/upload/required")
		t.Assert(content, `{"code":51,"message":"upload file is required","data":null}`)
	})
}

func Test_Params_File_MarshalJSON(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/upload/single", func(r *http类.Request) {
		file := r.X取上传文件对象("file")
		if file == nil {
			r.Response.X写响应缓冲区并退出("upload file cannot be empty")
		}

		if bytes, err := json.Marshal(file); err != nil {
			r.Response.X写响应缓冲区并退出(err)
		} else {
			r.Response.X写响应缓冲区并退出(bytes)
		}
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// normal name
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath := 单元测试类.DataPath("upload", "file1.txt")
		content := client.Post文本(ctx, "/upload/single", g.Map{
			"file": "@file:" + srcPath,
		})
		t.Assert(strings.Contains(content, "file1.txt"), true)
	})
}

// 批量上传时仅选择一个文件
func Test_Params_Strict_Route_File_Batch_Up_One(t *testing.T) {
	type Req struct {
		元数据类.Meta `method:"post" mime:"multipart/form-data"`
		Files      http类.UploadFiles `type:"file"`
	}
	type Res struct{}

	dstDirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
	s := g.Http类(uid类.X生成())
	s.X绑定("/upload/batch", func(ctx context.Context, req *Req) (res *Res, err error) {
		var (
			r     = g.Http类上下文取请求对象(ctx)
			files = req.Files
		)
		if len(files) == 0 {
			r.Response.X写响应缓冲区并退出("upload file cannot be empty")
		}
		names, err := files.X保存(dstDirPath)
		if err != nil {
			r.Response.X写响应缓冲区并退出(err)
		}
		r.Response.X写响应缓冲区并退出(文本类.X连接(names, ","))
		return
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// normal name
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath := 单元测试类.DataPath("upload", "file1.txt")
		dstPath := 文件类.X路径生成(dstDirPath, "file1.txt")
		content := client.Post文本(ctx, "/upload/batch", g.Map{
			"files": "@file:" + srcPath,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt")
		t.Assert(文件类.X读文本(dstPath), 文件类.X读文本(srcPath))
	})
}

// 批量上传时选择多个文件
func Test_Params_Strict_Route_File_Batch_Up_Multiple(t *testing.T) {
	type Req struct {
		元数据类.Meta `method:"post" mime:"multipart/form-data"`
		Files      http类.UploadFiles `type:"file"`
	}
	type Res struct{}

	dstDirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
	s := g.Http类(uid类.X生成())
	s.X绑定("/upload/batch", func(ctx context.Context, req *Req) (res *Res, err error) {
		var (
			r     = g.Http类上下文取请求对象(ctx)
			files = req.Files
		)
		if len(files) == 0 {
			r.Response.X写响应缓冲区并退出("upload file cannot be empty")
		}
		names, err := files.X保存(dstDirPath)
		if err != nil {
			r.Response.X写响应缓冲区并退出(err)
		}
		r.Response.X写响应缓冲区并退出(文本类.X连接(names, ","))
		return
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(100 * time.Millisecond)
	// normal name
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		srcPath1 := 单元测试类.DataPath("upload", "file1.txt")
		srcPath2 := 单元测试类.DataPath("upload", "file2.txt")
		dstPath1 := 文件类.X路径生成(dstDirPath, "file1.txt")
		dstPath2 := 文件类.X路径生成(dstDirPath, "file2.txt")
		content := client.Post文本(ctx, "/upload/batch",
			"files=@file:"+srcPath1+
				"&files=@file:"+srcPath2,
		)
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt,file2.txt")
		t.Assert(文件类.X读文本(dstPath1), 文件类.X读文本(srcPath1))
		t.Assert(文件类.X读文本(dstPath2), 文件类.X读文本(srcPath2))
	})
}
