// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Job Operations

package 定时类_test

import (
	"context"
	"testing"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/os/gtimer"
	"github.com/888go/goframe/test/gtest"
)

func TestJob_Start_Stop_Close(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timer := 定时类.X创建()
		array := 数组类.X创建(true)
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

		t.Assert(job.X取任务状态(), 定时类.StatusClosed)
	})
}

func TestJob_Singleton(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timer := 定时类.X创建()
		array := 数组类.X创建(true)
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
	单元测试类.C(t, func(t *单元测试类.T) {
		timer := 定时类.X创建(定时类.TimerOptions{
			Quick: true,
		})
		array := 数组类.X创建(true)
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
	单元测试类.C(t, func(t *单元测试类.T) {
		timer := 定时类.X创建()
		array := 数组类.X创建(true)
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
	单元测试类.C(t, func(t *单元测试类.T) {
		timer := 定时类.X创建()
		array := 数组类.X创建(true)
		job := timer.X加入循环任务(ctx, 1000*time.Millisecond, func(ctx context.Context) {
			array.Append别名(1)
		})
		job.X取任务函数()(ctx)
		t.Assert(array.X取长度(), 1)
	})
}
