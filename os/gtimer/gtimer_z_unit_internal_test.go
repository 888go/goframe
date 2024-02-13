// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 定时类

import (
	"context"
	"testing"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/test/gtest"
)

func TestTimer_Proceed(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建(true)
		timer := X创建(TimerOptions{
			Interval: time.Hour,
		})
		timer.X加入循环任务(ctx, 10000*time.Hour, func(ctx context.Context) {
			array.Append别名(1)
		})
		timer.proceed(10001)
		time.Sleep(10 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		timer.proceed(20001)
		time.Sleep(10 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建(true)
		timer := X创建(TimerOptions{
			Interval: time.Millisecond * 100,
		})
		timer.X加入循环任务(ctx, 10000*time.Hour, func(ctx context.Context) {
			array.Append别名(1)
		})
		ticks := int64((10000 * time.Hour) / (time.Millisecond * 100))
		timer.proceed(ticks + 1)
		time.Sleep(10 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		timer.proceed(2*ticks + 1)
		time.Sleep(10 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func TestTimer_PriorityQueue(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		queue := newPriorityQueue()
		queue.Push(1, 1)
		queue.Push(4, 4)
		queue.Push(5, 5)
		queue.Push(2, 2)
		queue.Push(3, 3)
		t.Assert(queue.Pop(), 1)
		t.Assert(queue.Pop(), 2)
		t.Assert(queue.Pop(), 3)
		t.Assert(queue.Pop(), 4)
		t.Assert(queue.Pop(), 5)
	})
}

func TestTimer_PriorityQueue_FirstOneInArrayIsTheLeast(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			size  = 1000000
			array = 数组类.X创建整数并按范围(0, size, 1)
		)
		array.X随机排序()
		queue := newPriorityQueue()
		array.X遍历(func(k int, v int) bool {
			queue.Push(v, int64(v))
			return true
		})
		for i := 0; i < size; i++ {
			t.Assert(queue.Pop(), i)
			t.Assert(queue.heap.array[0].priority, i+1)
		}
	})
}
