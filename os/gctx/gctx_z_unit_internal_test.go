// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package 上下文类_test

import (
	"context"
	"testing"
	"time"

	gctx "github.com/888go/goframe/os/gctx"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_NeverDone(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx, _ := context.WithDeadline(gctx.New(), time.Now().Add(time.Hour))
		t.AssertNE(ctx, nil)
		t.AssertNE(ctx.Done(), nil)
		t.Assert(ctx.Err(), nil)

		tm, ok := ctx.Deadline()
		t.AssertNE(tm, time.Time{})
		t.Assert(ok, true)

		ctx = gctx.NeverDone(ctx)
		t.AssertNE(ctx, nil)
		t.Assert(ctx.Done(), nil)
		t.Assert(ctx.Err(), nil)

		tm, ok = ctx.Deadline()
		t.Assert(tm, time.Time{})
		t.Assert(ok, false)
	})
}
