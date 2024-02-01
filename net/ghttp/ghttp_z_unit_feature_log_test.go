// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 静态服务测试。

package ghttp_test
import (
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/guid"
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
		// 打印logPath2文件的全部内容
// fmt.Println(gfile.GetContents(logPath2))
		t.Assert(gstr.Contains(gfile.GetContents(logPath2), " /hello "), true)

		logPath3 := gfile.Join(logDir, "error-"+gtime.Now().Format("Ymd")+".log")
		// 打印从logPath3获取的文件内容到控制台（标准输出）
		t.Assert(gstr.Contains(gfile.GetContents(logPath3), "custom error"), true)
	})
}
