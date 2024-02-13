// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类_test

import (
	"bytes"
	"context"
	"testing"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
)

var arrayForHandlerTest1 = 数组类.X创建文本()

func customHandler1(ctx context.Context, input *日志类.HandlerInput) {
	arrayForHandlerTest1.Append别名(input.String(false))
	input.Next(ctx)
}

func TestLogger_SetHandlers1(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		w := bytes.NewBuffer(nil)
		l := 日志类.X创建并按writer(w)
		l.X设置中间件(customHandler1)
		l.X设置上下文名称("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.X输出(ctx, 1, 2, 3)
		t.Assert(文本类.X统计次数(w.String(), "1234567890"), 1)
		t.Assert(文本类.X统计次数(w.String(), "abcdefg"), 1)
		t.Assert(文本类.X统计次数(w.String(), "1 2 3"), 1)

		t.Assert(arrayForHandlerTest1.X取长度(), 1)
		t.Assert(文本类.X统计次数(arrayForHandlerTest1.X取值(0), "1234567890"), 1)
		t.Assert(文本类.X统计次数(arrayForHandlerTest1.X取值(0), "abcdefg"), 1)
		t.Assert(文本类.X统计次数(arrayForHandlerTest1.X取值(0), "1 2 3"), 1)
	})
}

var arrayForHandlerTest2 = 数组类.X创建文本()

func customHandler2(ctx context.Context, input *日志类.HandlerInput) {
	arrayForHandlerTest2.Append别名(input.String(false))
}

func TestLogger_SetHandlers2(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		w := bytes.NewBuffer(nil)
		l := 日志类.X创建并按writer(w)
		l.X设置中间件(customHandler2)
		l.X设置上下文名称("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.X输出(ctx, 1, 2, 3)
		t.Assert(文本类.X统计次数(w.String(), "1234567890"), 0)
		t.Assert(文本类.X统计次数(w.String(), "abcdefg"), 0)
		t.Assert(文本类.X统计次数(w.String(), "1 2 3"), 0)

		t.Assert(arrayForHandlerTest2.X取长度(), 1)
		t.Assert(文本类.X统计次数(arrayForHandlerTest2.X取值(0), "1234567890"), 1)
		t.Assert(文本类.X统计次数(arrayForHandlerTest2.X取值(0), "abcdefg"), 1)
		t.Assert(文本类.X统计次数(arrayForHandlerTest2.X取值(0), "1 2 3"), 1)
	})
}

func TestLogger_SetHandlers_HandlerJson(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		w := bytes.NewBuffer(nil)
		l := 日志类.X创建并按writer(w)
		l.X设置中间件(日志类.X中间件函数Json)
		l.X设置上下文名称("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.X输出DEBU(ctx, 1, 2, 3)
		t.Assert(文本类.X统计次数(w.String(), "1234567890"), 1)
		t.Assert(文本类.X统计次数(w.String(), "abcdefg"), 1)
		t.Assert(文本类.X统计次数(w.String(), `"1 2 3"`), 1)
		t.Assert(文本类.X统计次数(w.String(), `"DEBU"`), 1)
	})
}

func TestLogger_SetHandlers_HandlerStructure(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		w := bytes.NewBuffer(nil)
		l := 日志类.X创建并按writer(w)
		l.X设置中间件(日志类.X中间件函数文本结构化输出)
		l.X设置上下文名称("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.X输出DEBU(ctx, "debug", "uid", 1000)
		l.X输出INFO(ctx, "info", "' '", `"\n`)

		t.Assert(文本类.X统计次数(w.String(), "uid=1000"), 1)
		t.Assert(文本类.X统计次数(w.String(), "Content=debug"), 1)
		t.Assert(文本类.X统计次数(w.String(), `"' '"="\"\\n"`), 1)
		t.Assert(文本类.X统计次数(w.String(), `CtxStr="1234567890, abcdefg"`), 2)
	})
}

func Test_SetDefaultHandler(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		oldHandler := 日志类.X取默认中间件()
		日志类.X设置默认中间件(func(ctx context.Context, in *日志类.HandlerInput) {
			日志类.X中间件函数Json(ctx, in)
		})
		defer 日志类.X设置默认中间件(oldHandler)

		w := bytes.NewBuffer(nil)
		l := 日志类.X创建并按writer(w)
		l.X设置上下文名称("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.X输出DEBU(ctx, 1, 2, 3)
		t.Assert(文本类.X统计次数(w.String(), "1234567890"), 1)
		t.Assert(文本类.X统计次数(w.String(), "abcdefg"), 1)
		t.Assert(文本类.X统计次数(w.String(), `"1 2 3"`), 1)
		t.Assert(文本类.X统计次数(w.String(), `"DEBU"`), 1)
	})
}
