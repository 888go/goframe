// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gins_test

import (
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/gins"
	"github.com/888go/goframe/internal/instance"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/os/gcfg"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/test/gtest"
)

func Test_Server(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			path                = 配置类.X默认配置文件名称
			serverConfigContent = 单元测试类.DataContent("server", "config.yaml")
			err                 = 文件类.X写入文本(path, serverConfigContent)
		)
		t.AssertNil(err)
		defer 文件类.X删除(path)

		instance.Clear()
		defer instance.Clear()

		s := gins.Server("tempByInstanceName")
		s.X绑定("/", func(r *http类.Request) {
			r.Response.X写响应缓冲区("hello")
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()

		time.Sleep(100 * time.Millisecond)

		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := gins.HttpClient()
		client.X设置url前缀(prefix)
		t.Assert(client.Get文本(上下文类.X创建(), "/"), "hello")
	})
}
