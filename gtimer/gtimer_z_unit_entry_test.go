// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Job Operations

package gtimer_test

import (
	"context"
	"testing"
	"time"
	
	"github.com/gogf/gf/v2/container/garray"
	"github.com/888go/goframe/gtimer"
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
