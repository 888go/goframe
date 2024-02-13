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
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/guid"
)

func Test_Log(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		logDir := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		s := g.Http类(uid类.X生成())
		s.X绑定("/hello", func(r *http类.Request) {
			r.Response.X写响应缓冲区("hello")
		})
		s.X绑定("/error", func(r *http类.Request) {
			panic("custom error")
		})
		s.X设置日志存储目录(logDir)
		s.X设置日志开启访客记录(true)
		s.X设置日志开启错误记录(true)
		s.X设置日志开启输出到CMD(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		defer 文件类.X删除(logDir)
		time.Sleep(100 * time.Millisecond)
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/hello"), "hello")
		t.Assert(client.Get文本(ctx, "/error"), "exception recovered: custom error")

		var (
			logPath1 = 文件类.X路径生成(logDir, 时间类.X创建并按当前时间().X取格式文本("Y-m-d")+".log")
			content  = 文件类.X读文本(logPath1)
		)
		t.Assert(文本类.X是否包含(content, "http server started listening on"), true)
		t.Assert(文本类.X是否包含(content, "HANDLER"), true)

		logPath2 := 文件类.X路径生成(logDir, "access-"+时间类.X创建并按当前时间().X取格式文本("Ymd")+".log")
		// 打印logPath2文件的全部内容
// fmt.Println(gfile.GetContents(logPath2))
		t.Assert(文本类.X是否包含(文件类.X读文本(logPath2), " /hello "), true)

		logPath3 := 文件类.X路径生成(logDir, "error-"+时间类.X创建并按当前时间().X取格式文本("Ymd")+".log")
		// 打印从logPath3获取的文件内容到控制台（标准输出）
		t.Assert(文本类.X是否包含(文件类.X读文本(logPath3), "custom error"), true)
	})
}
