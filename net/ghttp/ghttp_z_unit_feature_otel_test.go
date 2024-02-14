// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

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
	s := g.Http类(uid类.X生成())
	s.X绑定("/", func(r *http类.X请求) {
		traceId = gtrace.GetTraceID(r.Context别名())
		r.X响应.X写响应缓冲区(r.X取URL())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X启用浏览器模式(true)
		client.X设置url前缀(prefix)
		res, err := client.Get响应对象(ctx, "/")
		t.AssertNil(err)
		defer res.X关闭()

		t.Assert(res.Header.Get("Trace-Id"), traceId)
	})
}
