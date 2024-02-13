// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package functions

package 定时类_test

import (
	"context"
	"testing"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/os/gtimer"
	"github.com/888go/goframe/test/gtest"
)

var (
	ctx = context.TODO()
)

func TestSetTimeout(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建(true)
		定时类.SetTimeout别名(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestSetInterval(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建(true)
		定时类.SetInterval别名(ctx, 300*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 3)
	})
}

func TestAddEntry(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建(true)
		定时类.X加入详细循环任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		}, false, 2, 定时类.StatusReady)
		time.Sleep(1100 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func TestAddSingleton(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建(true)
		定时类.X加入单例循环任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
			time.Sleep(10000 * time.Millisecond)
		})
		time.Sleep(1100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestAddTimes(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建(true)
		定时类.X加入指定次数任务(ctx, 200*time.Millisecond, 2, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func TestDelayAdd(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建(true)
		定时类.X延时加入循环任务(ctx, 500*time.Millisecond, 500*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(600 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
		time.Sleep(600 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestDelayAddEntry(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建(true)
		定时类.X延时加入详细循环任务(ctx, 500*time.Millisecond, 500*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		}, false, 2, 定时类.StatusReady)
		time.Sleep(500 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
		time.Sleep(2000 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func TestDelayAddSingleton(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建(true)
		定时类.X延时加入单例循环任务(ctx, 500*time.Millisecond, 500*time.Millisecond, func(ctx context.Context) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建(true)
		定时类.X延时加入单次任务(ctx, 1000*time.Millisecond, 2000*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(2000 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
		time.Sleep(2000 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestDelayAddTimes(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建(true)
		定时类.X延时加入指定次数任务(ctx, 500*time.Millisecond, 500*time.Millisecond, 2, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
		time.Sleep(1500 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}
