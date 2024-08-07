// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 定时类

import (
	"context"
	"testing"
	"time"

	garray "github.com/888go/goframe/container/garray"
	gtest "github.com/888go/goframe/test/gtest"
)

func TestTimer_Proceed(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := garray.X创建(true)
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
	gtest.C(t, func(t *gtest.T) {
		array := garray.X创建(true)
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
	gtest.C(t, func(t *gtest.T) {
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
	gtest.C(t, func(t *gtest.T) {
		var (
			size  = 1000000
			array = garray.X创建整数并按范围(0, size, 1)
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
