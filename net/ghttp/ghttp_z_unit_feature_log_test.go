// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 静态服务测试。 md5:2105c089651008de

package http类_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gfile "github.com/888go/goframe/os/gfile"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
	guid "github.com/888go/goframe/util/guid"
)

func Test_Log(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		logDir := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		s := g.Http类(guid.X生成())
		s.X绑定("/hello", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区("hello")
		})
		s.X绑定("/error", func(r *ghttp.Request) {
			panic("custom error")
		})
		s.X设置日志存储目录(logDir)
		s.X设置日志开启访客记录(true)
		s.X设置日志开启错误记录(true)
		s.X设置日志开启输出到CMD(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		defer gfile.X删除(logDir)
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/hello"), "hello")
		t.Assert(client.Get文本(ctx, "/error"), "exception recovered: custom error")

		var (
			logPath1 = gfile.X路径生成(logDir, gtime.X创建并按当前时间().X取格式文本("Y-m-d")+".log")
			content  = gfile.X读文本(logPath1)
		)
		t.Assert(gstr.X是否包含(content, "http server started listening on"), true)
		t.Assert(gstr.X是否包含(content, "HANDLER"), true)

		logPath2 := gfile.X路径生成(logDir, "access-"+gtime.X创建并按当前时间().X取格式文本("Ymd")+".log")
				// 打印日志路径2中的内容（使用 gfile.GetContents 函数）. md5:bbe5d38122741545
		t.Assert(gstr.X是否包含(gfile.X读文本(logPath2), " /hello "), true)

		logPath3 := gfile.X路径生成(logDir, "error-"+gtime.X创建并按当前时间().X取格式文本("Ymd")+".log")
				// 打印日志路径3中的内容（使用gfile.GetContents函数）. md5:57930db9a93e7c59
		t.Assert(gstr.X是否包含(gfile.X读文本(logPath3), "custom error"), true)
	})
}
