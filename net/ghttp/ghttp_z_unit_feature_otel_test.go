// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/net/gtrace"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func Test_OTEL_TraceID(t *testing.T) {
	var (
		traceId string
	)
	s := g.Http类(guid.X生成())
	s.X绑定("/", func(r *ghttp.Request) {
		traceId = gtrace.GetTraceID(r.Context别名())
		r.X响应.X写响应缓冲区(r.X取URL())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
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
