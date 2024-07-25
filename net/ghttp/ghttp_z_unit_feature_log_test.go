// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 静态服务测试。 md5:2105c089651008de

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
)

func Test_Log(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		logDir := gfile.Temp(gtime.TimestampNanoStr())
		s := g.Server(guid.S())
		s.BindHandler("/hello", func(r *ghttp.Request) {
			r.Response.Write("hello")
		})
		s.BindHandler("/error", func(r *ghttp.Request) {
			panic("custom error")
		})
		s.SetLogPath(logDir)
		s.SetAccessLogEnabled(true)
		s.SetErrorLogEnabled(true)
		s.SetLogStdout(false)
		s.Start()
		defer s.Shutdown()
		defer gfile.Remove(logDir)
		time.Sleep(100 * time.Millisecond)
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/hello"), "hello")
		t.Assert(client.GetContent(ctx, "/error"), "exception recovered: custom error")

		var (
			logPath1 = gfile.Join(logDir, gtime.Now().Format("Y-m-d")+".log")
			content  = gfile.GetContents(logPath1)
		)
		t.Assert(gstr.Contains(content, "http server started listening on"), true)
		t.Assert(gstr.Contains(content, "HANDLER"), true)

		logPath2 := gfile.Join(logDir, "access-"+gtime.Now().Format("Ymd")+".log")
		// 打印日志路径2中的内容（使用 gfile.GetContents 函数）. md5:bbe5d38122741545
		t.Assert(gstr.Contains(gfile.GetContents(logPath2), " /hello "), true)

		logPath3 := gfile.Join(logDir, "error-"+gtime.Now().Format("Ymd")+".log")
		// 打印日志路径3中的内容（使用gfile.GetContents函数）. md5:57930db9a93e7c59
		t.Assert(gstr.Contains(gfile.GetContents(logPath3), "custom error"), true)
	})
}
