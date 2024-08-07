// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 日志类_test

import (
	"bytes"
	"context"
	"testing"

	garray "github.com/888go/goframe/container/garray"
	glog "github.com/888go/goframe/os/glog"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

var arrayForHandlerTest1 = garray.X创建文本()

func customHandler1(ctx context.Context, input *glog.HandlerInput) {
	arrayForHandlerTest1.Append别名(input.String(false))
	input.Next(ctx)
}

func TestLogger_SetHandlers1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		l := glog.X创建并按writer(w)
		l.X设置中间件(customHandler1)
		l.X设置上下文名称("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.X输出(ctx, 1, 2, 3)
		t.Assert(gstr.X统计次数(w.String(), "1234567890"), 1)
		t.Assert(gstr.X统计次数(w.String(), "abcdefg"), 1)
		t.Assert(gstr.X统计次数(w.String(), "1 2 3"), 1)

		t.Assert(arrayForHandlerTest1.X取长度(), 1)
		t.Assert(gstr.X统计次数(arrayForHandlerTest1.X取值(0), "1234567890"), 1)
		t.Assert(gstr.X统计次数(arrayForHandlerTest1.X取值(0), "abcdefg"), 1)
		t.Assert(gstr.X统计次数(arrayForHandlerTest1.X取值(0), "1 2 3"), 1)
	})
}

var arrayForHandlerTest2 = garray.X创建文本()

func customHandler2(ctx context.Context, input *glog.HandlerInput) {
	arrayForHandlerTest2.Append别名(input.String(false))
}

func TestLogger_SetHandlers2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		l := glog.X创建并按writer(w)
		l.X设置中间件(customHandler2)
		l.X设置上下文名称("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.X输出(ctx, 1, 2, 3)
		t.Assert(gstr.X统计次数(w.String(), "1234567890"), 0)
		t.Assert(gstr.X统计次数(w.String(), "abcdefg"), 0)
		t.Assert(gstr.X统计次数(w.String(), "1 2 3"), 0)

		t.Assert(arrayForHandlerTest2.X取长度(), 1)
		t.Assert(gstr.X统计次数(arrayForHandlerTest2.X取值(0), "1234567890"), 1)
		t.Assert(gstr.X统计次数(arrayForHandlerTest2.X取值(0), "abcdefg"), 1)
		t.Assert(gstr.X统计次数(arrayForHandlerTest2.X取值(0), "1 2 3"), 1)
	})
}

func TestLogger_SetHandlers_HandlerJson(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		l := glog.X创建并按writer(w)
		l.X设置中间件(glog.X中间件函数Json)
		l.X设置上下文名称("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.X输出DEBU(ctx, 1, 2, 3)
		t.Assert(gstr.X统计次数(w.String(), `"CtxStr":"1234567890, abcdefg"`), 1)
		t.Assert(gstr.X统计次数(w.String(), `"Content":"1 2 3"`), 1)
		t.Assert(gstr.X统计次数(w.String(), `"Level":"DEBU"`), 1)
	})
}

func TestLogger_SetHandlers_HandlerStructure(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		l := glog.X创建并按writer(w)
		l.X设置中间件(glog.X中间件函数文本结构化输出)
		l.X设置上下文名称("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.X输出DEBU(ctx, "debug", "uid", 1000)
		l.X输出INFO(ctx, "info", "' '", `"\n`)

		t.Assert(gstr.X统计次数(w.String(), "uid=1000"), 1)
		t.Assert(gstr.X统计次数(w.String(), "Content=debug"), 1)
		t.Assert(gstr.X统计次数(w.String(), `"' '"="\"\\n"`), 1)
		t.Assert(gstr.X统计次数(w.String(), `CtxStr="1234567890, abcdefg"`), 2)
	})
}

func Test_SetDefaultHandler(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		oldHandler := glog.X取默认中间件()
		glog.X设置默认中间件(func(ctx context.Context, in *glog.HandlerInput) {
			glog.X中间件函数Json(ctx, in)
		})
		defer glog.X设置默认中间件(oldHandler)

		w := bytes.NewBuffer(nil)
		l := glog.X创建并按writer(w)
		l.X设置上下文名称("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.X输出DEBU(ctx, 1, 2, 3)
		t.Assert(gstr.X统计次数(w.String(), "1234567890"), 1)
		t.Assert(gstr.X统计次数(w.String(), "abcdefg"), 1)
		t.Assert(gstr.X统计次数(w.String(), `"1 2 3"`), 1)
		t.Assert(gstr.X统计次数(w.String(), `"DEBU"`), 1)
	})
}
