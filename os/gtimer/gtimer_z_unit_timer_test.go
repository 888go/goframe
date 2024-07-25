// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Timer Operations

package gtimer_test

import (
	"context"
	"testing"
	"time"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestTimer_Add_Close(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.New()
		array := garray.New(true)
				// 打印 "start" 和当前时间。 md5:0798a4ad01b9b7d4
		timer.Add(ctx, 200*time.Millisecond, func(ctx context.Context) {
						// 这行Go语言注释的中文翻译是："打印job1和当前时间（使用time.Now()函数）"。它表示在控制台上输出字符串"job1"后面跟着当前的时间戳。 md5:35e2e54687424448
			array.Append(1)
		})
		timer.Add(ctx, 200*time.Millisecond, func(ctx context.Context) {
						// 这行Go语言注释的中文翻译是："打印'job2'和当前时间"。它使用了"fmt.Println"函数来输出字符串"job2"以及系统当前的时间。 md5:1130c8baf731a6bb
			array.Append(1)
		})
		timer.Add(ctx, 400*time.Millisecond, func(ctx context.Context) {
						// 注释中的 `fmt.Println` 是一个函数调用，用于输出文本到标准输出（通常是终端或控制台），紧跟的是它将打印的两个参数："job3" 和 `time.Now()`。其中 `time.Now()` 是一个函数，用于获取当前的日期和时间。但由于这行代码被注释了，所以在程序运行时不会执行实际的打印操作。 md5:b27572b5b4091706
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 2)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 5)
		timer.Close()
		time.Sleep(250 * time.Millisecond)
		fixedLength := array.Len()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), fixedLength)
	})
}

func TestTimer_Start_Stop_Close(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.New()
		array := garray.New(true)
		timer.Add(ctx, 1000*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		t.Assert(array.Len(), 0)
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 1)
		timer.Stop()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 1)
		timer.Start()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 2)
		timer.Close()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestJob_Reset(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.New()
		array := garray.New(true)
		job := timer.AddSingleton(ctx, 500*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(300 * time.Millisecond)
		job.Reset()
		time.Sleep(300 * time.Millisecond)
		job.Reset()
		time.Sleep(300 * time.Millisecond)
		job.Reset()
		time.Sleep(600 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_AddSingleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.New()
		array := garray.New(true)
		timer.AddSingleton(ctx, 200*time.Millisecond, func(ctx context.Context) {
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
		timer := gtimer.New(gtimer.TimerOptions{
			Interval: 100 * time.Millisecond,
			Quick:    true,
		})
		array := garray.New(true)
		timer.AddSingleton(ctx, 5*time.Second, func(ctx context.Context) {
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
		timer := gtimer.New(gtimer.TimerOptions{
			Interval: 100 * time.Millisecond,
			Quick:    false,
		})
		array := garray.New(true)
		timer.AddSingleton(ctx, 5*time.Second, func(ctx context.Context) {
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
		timer := gtimer.New()
		array := garray.New(true)
		timer.AddOnce(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		timer.AddOnce(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 2)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 2)
		timer.Close()
		time.Sleep(250 * time.Millisecond)
		fixedLength := array.Len()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), fixedLength)
	})
}

func TestTimer_AddTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.New()
		array := garray.New(true)
		timer.AddTimes(ctx, 200*time.Millisecond, 2, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestTimer_DelayAdd(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.New()
		array := garray.New(true)
		timer.DelayAdd(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
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
		timer := gtimer.New()
		array := garray.New(true)
		timer.DelayAddEntry(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		}, false, 100, gtimer.StatusReady)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_DelayAddSingleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.New()
		array := garray.New(true)
		timer.DelayAddSingleton(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
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
		timer := gtimer.New()
		array := garray.New(true)
		timer.DelayAddOnce(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
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
		timer := gtimer.New()
		array := garray.New(true)
		timer.DelayAddTimes(ctx, 200*time.Millisecond, 500*time.Millisecond, 2, func(ctx context.Context) {
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
		timer := gtimer.New(gtimer.TimerOptions{
			Interval: 100 * time.Millisecond,
		})
		array := garray.New(true)
		timer.Add(ctx, 20*time.Millisecond, func(ctx context.Context) {
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
		timer := gtimer.New()
		array := garray.New(true)
		timer.DelayAdd(ctx, 1000*time.Millisecond, 1000*time.Millisecond, func(ctx context.Context) {
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
		timer := gtimer.New()
		array := garray.New(true)
		timer.Add(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
			gtimer.Exit()
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}
