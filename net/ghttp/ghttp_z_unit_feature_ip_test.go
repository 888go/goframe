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
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

func TestRequest_GetRemoteIp(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		s.X绑定("/", func(r *http类.Request) {
			r.Response.X写响应缓冲区(r.X取远程IP地址())
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()

		time.Sleep(100 * time.Millisecond)

		clientV4 := g.X网页类()
		clientV4.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		clientV6 := g.X网页类()
		clientV6.X设置url前缀(fmt.Sprintf("http://[::1]:%d", s.X取已监听端口()))

		t.Assert(clientV4.Get文本(ctx, "/"), "127.0.0.1")
		t.Assert(clientV6.Get文本(ctx, "/"), "::1")
	})
}
