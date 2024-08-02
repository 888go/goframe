// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// static service testing.

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
		// fmt.Println(gfile.GetContents(logPath2))
		t.Assert(gstr.Contains(gfile.GetContents(logPath2), " /hello "), true)

		logPath3 := gfile.Join(logDir, "error-"+gtime.Now().Format("Ymd")+".log")
		// fmt.Println(gfile.GetContents(logPath3))
		t.Assert(gstr.Contains(gfile.GetContents(logPath3), "custom error"), true)
	})
}
