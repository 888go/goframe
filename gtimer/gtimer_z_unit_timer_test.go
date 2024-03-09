// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Timer Operations

package 定时类_test

import (
	"context"
	"testing"
	"time"
	
	"github.com/gogf/gf/v2/container/garray"
	"github.com/888go/goframe/gtimer"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestTimer_Add_Close(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建()
		array := garray.New(true)
		// 输出 "start" 及当前时间（用 time.Now() 获取）
// ```go
//fmt.Println("start", time.Now())
		timer.X加入循环任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			// 输出 "job1" 及当前时间，使用 fmt.Println() 函数实现
			array.Append(1)
		})
		timer.X加入循环任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
// 输出 "job2" 及当前时间（用 time.Now() 函数获取）到标准输出（通常是终端或控制台）
			array.Append(1)
		})
		timer.X加入循环任务(ctx, 400*time.Millisecond, func(ctx context.Context) {
// 打印输出 "job3" 及 当前时间，使用 fmt.Println() 函数
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 2)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 5)
		timer.X关闭任务()
		time.Sleep(250 * time.Millisecond)
		fixedLength := array.Len()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), fixedLength)
	})
}

func TestTimer_Start_Stop_Close(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建()
		array := garray.New(true)
		timer.X加入循环任务(ctx, 1000*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		t.Assert(array.Len(), 0)
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 1)
		timer.X暂停工作()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 1)
		timer.X开始工作()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 2)
		timer.X关闭任务()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestJob_Reset(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建()
		array := garray.New(true)
		job := timer.X加入单例循环任务(ctx, 500*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(300 * time.Millisecond)
		job.X重置任务()
		time.Sleep(300 * time.Millisecond)
		job.X重置任务()
		time.Sleep(300 * time.Millisecond)
		job.X重置任务()
		time.Sleep(600 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_AddSingleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建()
		array := garray.New(true)
		timer.X加入单例循环任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
			time.Sleep(10 * time.Second)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)

		time.Sleep(500 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_AddSingletonWithQuick(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建(定时类.TimerOptions{
			Interval: 100 * time.Millisecond,
			Quick:    true,
		})
		array := garray.New(true)
		timer.X加入单例循环任务(ctx, 5*time.Second, func(ctx context.Context) {
			array.Append(1)
			time.Sleep(10 * time.Second)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)

		time.Sleep(500 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_AddSingletonWithoutQuick(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建(定时类.TimerOptions{
			Interval: 100 * time.Millisecond,
			Quick:    false,
		})
		array := garray.New(true)
		timer.X加入单例循环任务(ctx, 5*time.Second, func(ctx context.Context) {
			array.Append(1)
			time.Sleep(10 * time.Second)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 0)

		time.Sleep(500 * time.Millisecond)
		t.Assert(array.Len(), 0)
	})
}

func TestTimer_AddOnce(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建()
		array := garray.New(true)
		timer.X加入单次任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		timer.X加入单次任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 2)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 2)
		timer.X关闭任务()
		time.Sleep(250 * time.Millisecond)
		fixedLength := array.Len()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), fixedLength)
	})
}

func TestTimer_AddTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建()
		array := garray.New(true)
		timer.X加入指定次数任务(ctx, 200*time.Millisecond, 2, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestTimer_DelayAdd(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建()
		array := garray.New(true)
		timer.X延时加入循环任务(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_DelayAddJob(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建()
		array := garray.New(true)
		timer.X延时加入详细循环任务(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		}, false, 100, 定时类.StatusReady)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_DelayAddSingleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建()
		array := garray.New(true)
		timer.X延时加入单例循环任务(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
			time.Sleep(10 * time.Second)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 0)

		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_DelayAddOnce(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建()
		array := garray.New(true)
		timer.X延时加入单次任务(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 0)

		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)

		time.Sleep(500 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_DelayAddTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建()
		array := garray.New(true)
		timer.X延时加入指定次数任务(ctx, 200*time.Millisecond, 500*time.Millisecond, 2, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 0)

		time.Sleep(600 * time.Millisecond)
		t.Assert(array.Len(), 1)

		time.Sleep(600 * time.Millisecond)
		t.Assert(array.Len(), 2)

		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestTimer_AddLessThanInterval(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建(定时类.TimerOptions{
			Interval: 100 * time.Millisecond,
		})
		array := garray.New(true)
		timer.X加入循环任务(ctx, 20*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(50 * time.Millisecond)
		t.Assert(array.Len(), 0)

		time.Sleep(110 * time.Millisecond)
		t.Assert(array.Len(), 1)

		time.Sleep(110 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestTimer_AddLeveledJob1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建()
		array := garray.New(true)
		timer.X延时加入循环任务(ctx, 1000*time.Millisecond, 1000*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(1500 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(1300 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_Exit(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := 定时类.X创建()
		array := garray.New(true)
		timer.X加入循环任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
			定时类.X退出()
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}
