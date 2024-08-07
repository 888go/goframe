// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Job Operations

package 定时类_test

import (
	"context"
	"testing"
	"time"

	garray "github.com/888go/goframe/container/garray"
	gtimer "github.com/888go/goframe/os/gtimer"
	gtest "github.com/888go/goframe/test/gtest"
)

func TestJob_Start_Stop_Close(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		job := timer.X加入循环任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		job.X暂停工作()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		job.X开始工作()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
		job.X关闭任务()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 2)

		t.Assert(job.X取任务状态(), gtimer.StatusClosed)
	})
}

func TestJob_Singleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		job := timer.X加入循环任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
			time.Sleep(10 * time.Second)
		})
		t.Assert(job.X是否单例模式(), false)
		job.X设置单例模式(true)
		t.Assert(job.X是否单例模式(), true)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 1)

		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestJob_SingletonQuick(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建(gtimer.TimerOptions{
			Quick: true,
		})
		array := garray.X创建(true)
		job := timer.X加入循环任务(ctx, 5*time.Second, func(ctx context.Context) {
			array.Append别名(1)
			time.Sleep(10 * time.Second)
		})
		t.Assert(job.X是否单例模式(), false)
		job.X设置单例模式(true)
		t.Assert(job.X是否单例模式(), true)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 1)

		time.Sleep(250 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestJob_SetTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		job := timer.X加入循环任务(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		job.X设置任务次数(2)
		//job.IsSingleton()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func TestJob_Run(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.X创建()
		array := garray.X创建(true)
		job := timer.X加入循环任务(ctx, 1000*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		job.X取任务函数()(ctx)
		t.Assert(array.X取长度(), 1)
	})
}
