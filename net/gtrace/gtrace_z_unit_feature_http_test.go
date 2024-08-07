// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtrace_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/net/gtrace"
	gctx "github.com/888go/goframe/os/gctx"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_Client_Server_Tracing(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p := 8888
		s := g.Http类(p)
		s.X绑定("/", func(r *ghttp.Request) {
			ctx := r.Context别名()
			g.X日志类().X输出(ctx, "GetTraceID:", gtrace.GetTraceID(ctx))
			r.X响应.X写响应缓冲区(gtrace.GetTraceID(ctx))
		})
		s.X设置监听端口(p)
		s.SetDumpRouterMap(false)
		t.AssertNil(s.X开始监听())
		defer s.X关闭当前服务()

		time.Sleep(100 * time.Millisecond)

		ctx := gctx.X创建()

		prefix := fmt.Sprintf("http://127.0.0.1:%d", p)
		client := g.X网页类()
		client.X设置url前缀(prefix)
		t.Assert(gtrace.IsUsingDefaultProvider(), true)
		t.Assert(client.Get文本(ctx, "/"), gtrace.GetTraceID(ctx))
		t.Assert(client.Get文本(ctx, "/"), gctx.X取上下文id(ctx))
	})
}

func Test_WithTraceID(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p := 8889
		s := g.Http类(p)
		s.X绑定("/", func(r *ghttp.Request) {
			ctx := r.Context别名()
			r.X响应.X写响应缓冲区(gtrace.GetTraceID(ctx))
		})
		s.X设置监听端口(p)
		s.SetDumpRouterMap(false)
		t.AssertNil(s.X开始监听())
		defer s.X关闭当前服务()

		time.Sleep(100 * time.Millisecond)

		ctx, err := gtrace.WithTraceID(context.TODO(), traceID.String())
		t.AssertNil(err)

		prefix := fmt.Sprintf("http://127.0.0.1:%d", p)
		client := g.X网页类()
		client.X设置url前缀(prefix)
		t.Assert(gtrace.IsUsingDefaultProvider(), true)
		t.Assert(client.Get文本(ctx, "/"), gtrace.GetTraceID(ctx))
		t.Assert(client.Get文本(ctx, "/"), traceIDStr)
	})
}
