// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 上下文类_test

import (
	"context"
	"testing"
	
	"github.com/888go/goframe/gctx"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_New(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := 上下文类.X创建()
		t.AssertNE(ctx, nil)
		t.AssertNE(上下文类.X取上下文id(ctx), "")
	})
}

func Test_WithCtx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.WithValue(context.TODO(), "TEST", 1)
		ctx = 上下文类.X创建并从上下文(ctx)
		t.AssertNE(上下文类.X取上下文id(ctx), "")
		t.Assert(ctx.Value("TEST"), 1)
	})
}

func Test_SetInitCtx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.WithValue(context.TODO(), "TEST", 1)
		上下文类.X设置初始化上下文(ctx)
		t.AssertNE(上下文类.X取初始化上下文(), "")
		t.Assert(上下文类.X取初始化上下文().Value("TEST"), 1)
	})
}
