// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package functions

package gtimer_test

import (
	"context"
	"testing"
	"time"
	
	"github.com/gogf/gf/v2/container/garray"
	"github.com/888go/goframe/gtimer"
	"github.com/gogf/gf/v2/test/gtest"
)

var (
	ctx = context.TODO()
)

func TestSetTimeout(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.New(true)
		gtimer.SetTimeout(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestSetInterval(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.New(true)
		gtimer.SetInterval(ctx, 300*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 3)
	})
}

func TestAddEntry(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.New(true)
		gtimer.AddEntry(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		}, false, 2, gtimer.StatusReady)
		time.Sleep(1100 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestAddSingleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.New(true)
		gtimer.AddSingleton(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
			time.Sleep(10000 * time.Millisecond)
		})
		time.Sleep(1100 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestAddTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.New(true)
		gtimer.AddTimes(ctx, 200*time.Millisecond, 2, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestDelayAdd(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.New(true)
		gtimer.DelayAdd(ctx, 500*time.Millisecond, 500*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(600 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(600 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestDelayAddEntry(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.New(true)
		gtimer.DelayAddEntry(ctx, 500*time.Millisecond, 500*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		}, false, 2, gtimer.StatusReady)
		time.Sleep(500 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(2000 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestDelayAddSingleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.New(true)
		gtimer.DelayAddSingleton(ctx, 500*time.Millisecond, 500*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
			time.Sleep(10000 * time.Millisecond)
		})
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestDelayAddOnce(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.New(true)
		gtimer.DelayAddOnce(ctx, 1000*time.Millisecond, 2000*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(2000 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(2000 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestDelayAddTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.New(true)
		gtimer.DelayAddTimes(ctx, 500*time.Millisecond, 500*time.Millisecond, 2, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(1500 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}
