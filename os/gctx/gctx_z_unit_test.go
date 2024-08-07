// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 上下文类_test

import (
	"context"
	"testing"

	gctx "github.com/888go/goframe/os/gctx"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_New(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := gctx.X创建()
		t.AssertNE(ctx, nil)
		t.AssertNE(gctx.X取上下文id(ctx), "")
	})
}

func Test_WithCtx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.WithValue(context.TODO(), "TEST", 1)
		ctx = gctx.X创建并从上下文(ctx)
		t.AssertNE(gctx.X取上下文id(ctx), "")
		t.Assert(ctx.Value("TEST"), 1)
	})
}

func Test_SetInitCtx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.WithValue(context.TODO(), "TEST", 1)
		gctx.X设置初始化上下文(ctx)
		t.AssertNE(gctx.X取初始化上下文(), "")
		t.Assert(gctx.X取初始化上下文().Value("TEST"), 1)
	})
}
