// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// Job Operations

package gtimer_test

import (
	"context"
	"testing"
	"time"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestJob_Start_Stop_Close(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.New()
		array := garray.New(true)
		job := timer.Add(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)
		job.Stop()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)
		job.Start()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 2)
		job.Close()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 2)

		t.Assert(job.Status(), gtimer.StatusClosed)
	})
}

func TestJob_Singleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.New()
		array := garray.New(true)
		job := timer.Add(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
			time.Sleep(10 * time.Second)
		})
		t.Assert(job.IsSingleton(), false)
		job.SetSingleton(true)
		t.Assert(job.IsSingleton(), true)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)

		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestJob_SingletonQuick(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.New(gtimer.TimerOptions{
			Quick: true,
		})
		array := garray.New(true)
		job := timer.Add(ctx, 5*time.Second, func(ctx context.Context) {
			array.Append(1)
			time.Sleep(10 * time.Second)
		})
		t.Assert(job.IsSingleton(), false)
		job.SetSingleton(true)
		t.Assert(job.IsSingleton(), true)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)

		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestJob_SetTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.New()
		array := garray.New(true)
		job := timer.Add(ctx, 200*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		job.SetTimes(2)
		//job.IsSingleton()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestJob_Run(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.New()
		array := garray.New(true)
		job := timer.Add(ctx, 1000*time.Millisecond, func(ctx context.Context) {
			array.Append(1)
		})
		job.Job()(ctx)
		t.Assert(array.Len(), 1)
	})
}
