// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Package functions

package 定时类_test

import (
	"context"
	"testing"
	"time"

	garray "github.com/888go/goframe/container/garray"
	gtimer "github.com/888go/goframe/os/gtimer"
	gtest "github.com/888go/goframe/test/gtest"
)

var (
	ctx = context.TODO()
)

func TestSetTimeout(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.X创建(true)
		gtimer.SetTimeout别名(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestSetInterval(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.X创建(true)
		gtimer.SetInterval别名(ctx, 300*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 3)
	})
}

func TestAddEntry(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.X创建(true)
		gtimer.X加入详细循环任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		}, false, 2, gtimer.StatusReady)
		time.Sleep(1100 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func TestAddSingleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.X创建(true)
		gtimer.X加入单例循环任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
			time.Sleep(10000 * time.Millisecond)
		})
		time.Sleep(1100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestAddTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.X创建(true)
		gtimer.X加入指定次数任务(ctx, 200*time.Millisecond, 2, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func TestDelayAdd(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.X创建(true)
		gtimer.X延时加入循环任务(ctx, 500*time.Millisecond, 500*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(600 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
		time.Sleep(600 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestDelayAddEntry(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.X创建(true)
		gtimer.X延时加入详细循环任务(ctx, 500*time.Millisecond, 500*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		}, false, 2, gtimer.StatusReady)
		time.Sleep(500 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
		time.Sleep(2000 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func TestDelayAddSingleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.X创建(true)
		gtimer.X延时加入单例循环任务(ctx, 500*time.Millisecond, 500*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
			time.Sleep(10000 * time.Millisecond)
		})
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestDelayAddOnce(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.X创建(true)
		gtimer.X延时加入单次任务(ctx, 1000*time.Millisecond, 2000*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(2000 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
		time.Sleep(2000 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestDelayAddTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.X创建(true)
		gtimer.X延时加入指定次数任务(ctx, 500*time.Millisecond, 500*time.Millisecond, 2, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
		time.Sleep(1500 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}
