// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gctx_test

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_New(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := gctx.New()
		t.AssertNE(ctx, nil)
		t.AssertNE(gctx.CtxId(ctx), "")
	})
}

func Test_WithCtx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.WithValue(context.TODO(), "TEST", 1)
		ctx = gctx.WithCtx(ctx)
		t.AssertNE(gctx.CtxId(ctx), "")
		t.Assert(ctx.Value("TEST"), 1)
	})
}

func Test_SetInitCtx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.WithValue(context.TODO(), "TEST", 1)
		gctx.SetInitCtx(ctx)
		t.AssertNE(gctx.GetInitCtx(), "")
		t.Assert(gctx.GetInitCtx().Value("TEST"), 1)
	})
}
