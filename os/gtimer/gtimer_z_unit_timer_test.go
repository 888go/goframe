// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Timer Operations

package 定时类_test

import (
	"context"
	"testing"
	"time"

	garray "github.com/888go/goframe/container/garray"
	gtimer "github.com/888go/goframe/os/gtimer"
	gtest "github.com/888go/goframe/test/gtest"
)

func TestTimer_Add_Close(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
				// 打印 "start" 和当前时间。 md5:0798a4ad01b9b7d4
		timer.X加入循环任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
						// 这行Go语言注释的中文翻译是："打印job1和当前时间（使用time.Now()函数）"。它表示在控制台上输出字符串"job1"后面跟着当前的时间戳。 md5:35e2e54687424448
			array.Append别名(1)
		})
		timer.X加入循环任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
						// 这行Go语言注释的中文翻译是："打印'job2'和当前时间"。它使用了"fmt.Println"函数来输出字符串"job2"以及系统当前的时间。 md5:1130c8baf731a6bb
			array.Append别名(1)
		})
		timer.X加入循环任务(ctx, 400*time.Millisecond, func(ctx context.Context) {
						// 注释中的 `fmt.Println` 是一个函数调用，用于输出文本到标准输出（通常是终端或控制台），紧跟的是它将打印的两个参数："job3" 和 `time.Now()`。其中 `time.Now()` 是一个函数，用于获取当前的日期和时间。但由于这行代码被注释了，所以在程序运行时不会执行实际的打印操作。 md5:b27572b5b4091706
			array.Append别名(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 5)
		timer.X关闭任务()
		time.Sleep(250 * time.Millisecond)
		fixedLength := array.X取长度()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), fixedLength)
	})
}

func TestTimer_Start_Stop_Close(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		timer.X加入循环任务(ctx, 1000*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		t.Assert(array.X取长度(), 0)
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		timer.X暂停工作()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		timer.X开始工作()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
		timer.X关闭任务()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func TestJob_Reset(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		job := timer.X加入单例循环任务(ctx, 500*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(300 * time.Millisecond)
		job.X重置任务()
		time.Sleep(300 * time.Millisecond)
		job.X重置任务()
		time.Sleep(300 * time.Millisecond)
		job.X重置任务()
		time.Sleep(600 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestTimer_AddSingleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		timer.X加入单例循环任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
			time.Sleep(10 * time.Second)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 1)

		time.Sleep(500 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestTimer_AddSingletonWithQuick(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建(gtimer.TimerOptions{
			Interval: 100 * time.Millisecond,
			Quick:    true,
		})
		array := garray.X创建(true)
		timer.X加入单例循环任务(ctx, 5*time.Second, func(ctx context.Context) {
			array.Append别名(1)
			time.Sleep(10 * time.Second)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 1)

		time.Sleep(500 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestTimer_AddSingletonWithoutQuick(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建(gtimer.TimerOptions{
			Interval: 100 * time.Millisecond,
			Quick:    false,
		})
		array := garray.X创建(true)
		timer.X加入单例循环任务(ctx, 5*time.Second, func(ctx context.Context) {
			array.Append别名(1)
			time.Sleep(10 * time.Second)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 0)

		time.Sleep(500 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
	})
}

func TestTimer_AddOnce(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		timer.X加入单次任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		timer.X加入单次任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
		timer.X关闭任务()
		time.Sleep(250 * time.Millisecond)
		fixedLength := array.X取长度()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), fixedLength)
	})
}

func TestTimer_AddTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		timer.X加入指定次数任务(ctx, 200*time.Millisecond, 2, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func TestTimer_DelayAdd(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		timer.X延时加入循环任务(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestTimer_DelayAddJob(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		timer.X延时加入详细循环任务(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		}, false, 100, gtimer.StatusReady)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestTimer_DelayAddSingleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		timer.X延时加入单例循环任务(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
			time.Sleep(10 * time.Second)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 0)

		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestTimer_DelayAddOnce(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		timer.X延时加入单次任务(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 0)

		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 1)

		time.Sleep(500 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestTimer_DelayAddTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		timer.X延时加入指定次数任务(ctx, 200*time.Millisecond, 500*time.Millisecond, 2, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 0)

		time.Sleep(600 * time.Millisecond)
		t.Assert(array.X取长度(), 1)

		time.Sleep(600 * time.Millisecond)
		t.Assert(array.X取长度(), 2)

		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func TestTimer_AddLessThanInterval(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建(gtimer.TimerOptions{
			Interval: 100 * time.Millisecond,
		})
		array := garray.X创建(true)
		timer.X加入循环任务(ctx, 20*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(50 * time.Millisecond)
		t.Assert(array.X取长度(), 0)

		time.Sleep(110 * time.Millisecond)
		t.Assert(array.X取长度(), 1)

		time.Sleep(110 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func TestTimer_AddLeveledJob1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		timer.X延时加入循环任务(ctx, 1000*time.Millisecond, 1000*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(1500 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
		time.Sleep(1300 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestTimer_Exit(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		timer.X加入循环任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
			gtimer.X退出()
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}
