// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
)

func Test_OTEL_TraceID(t *testing.T) {
	var (
		traceId string
	)
	s := g.Server(guid.S())
	s.BindHandler("/", func(r *ghttp.Request) {
		traceId = gtrace.GetTraceID(r.Context())
		r.Response.Write(r.GetUrl())
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
		client := g.Client()
		client.SetBrowserMode(true)
		client.SetPrefix(prefix)
		res, err := client.Get(ctx, "/")
		t.AssertNil(err)
		defer res.Close()

		t.Assert(res.Header.Get("Trace-Id"), traceId)
	})
}
