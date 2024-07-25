// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 静态服务测试。 md5:2105c089651008de

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
)

func TestRequest_GetRemoteIp(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.BindHandler("/", func(r *ghttp.Request) {
			r.Response.Write(r.GetRemoteIp())
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		time.Sleep(100 * time.Millisecond)

		clientV4 := g.Client()
		clientV4.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		clientV6 := g.Client()
		clientV6.SetPrefix(fmt.Sprintf("http://[::1]:%d", s.GetListenedPort()))

		t.Assert(clientV4.GetContent(ctx, "/"), "127.0.0.1")
		t.Assert(clientV6.GetContent(ctx, "/"), "::1")
	})
}
