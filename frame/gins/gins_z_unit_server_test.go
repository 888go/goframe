// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gins_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/888go/goframe/frame/gins"
	"github.com/888go/goframe/internal/instance"
	ghttp "github.com/888go/goframe/net/ghttp"
	gcfg "github.com/888go/goframe/os/gcfg"
	gctx "github.com/888go/goframe/os/gctx"
	gfile "github.com/888go/goframe/os/gfile"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_Server(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			path                = gcfg.DefaultConfigFileName
			serverConfigContent = gtest.DataContent("server", "config.yaml")
			err                 = gfile.PutContents(path, serverConfigContent)
		)
		t.AssertNil(err)
		defer gfile.Remove(path)

		instance.Clear()
		defer instance.Clear()

		s := gins.Server("tempByInstanceName")
		s.BindHandler("/", func(r *ghttp.Request) {
			r.Response.Write("hello")
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		time.Sleep(100 * time.Millisecond)

		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
		client := gins.HttpClient()
		client.SetPrefix(prefix)
		t.Assert(client.GetContent(gctx.New(), "/"), "hello")
	})
}
