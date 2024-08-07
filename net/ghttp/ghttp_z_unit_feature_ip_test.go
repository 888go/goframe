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
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func TestRequest_GetRemoteIp(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Http类(guid.X生成())
		s.X绑定("/", func(r *ghttp.Request) {
			r.X响应.X写响应缓冲区(r.X取远程IP地址())
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
