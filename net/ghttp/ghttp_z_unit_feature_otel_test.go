// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp_test
import (
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/net/gtrace"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
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
