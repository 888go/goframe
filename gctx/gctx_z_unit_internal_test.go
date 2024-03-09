// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 上下文类_test

import (
	"context"
	"testing"
	"time"
	
	"github.com/888go/goframe/gctx"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_NeverDone(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx, _ := context.WithDeadline(上下文类.X创建(), time.Now().Add(time.Hour))
		t.AssertNE(ctx, nil)
		t.AssertNE(ctx.Done(), nil)
		t.Assert(ctx.Err(), nil)

		tm, ok := ctx.Deadline()
		t.AssertNE(tm, time.Time{})
		t.Assert(ok, true)

		ctx = 上下文类.NeverDone(ctx)
		t.AssertNE(ctx, nil)
		t.Assert(ctx.Done(), nil)
		t.Assert(ctx.Err(), nil)

		tm, ok = ctx.Deadline()
		t.Assert(tm, time.Time{})
		t.Assert(ok, false)
	})
}
