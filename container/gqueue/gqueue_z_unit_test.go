// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。. md5:81db3d7bd1ed4da8

package gqueue_test

import (
	"testing"
	"time"

	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestQueue_Len(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			maxNum   = 100
			maxTries = 100
		)
		for n := 10; n < maxTries; n++ {
			q1 := gqueue.New(maxNum)
			for i := 0; i < maxNum; i++ {
				q1.Push(i)
			}
			t.Assert(q1.Len(), maxNum)
			t.Assert(q1.Size(), maxNum)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			maxNum   = 100
			maxTries = 100
		)
		for n := 10; n < maxTries; n++ {
			q1 := gqueue.New()
			for i := 0; i < maxNum; i++ {
				q1.Push(i)
			}
			t.AssertLE(q1.Len(), maxNum)
			t.AssertLE(q1.Size(), maxNum)
		}
	})
}

func TestQueue_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		q := gqueue.New()
		for i := 0; i < 100; i++ {
			q.Push(i)
		}
		t.Assert(q.Pop(), 0)
		t.Assert(q.Pop(), 1)
	})
}

func TestQueue_Pop(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		q1 := gqueue.New()
		q1.Push(1)
		q1.Push(2)
		q1.Push(3)
		q1.Push(4)
		i1 := q1.Pop()
		t.Assert(i1, 1)
	})
}

func TestQueue_Close(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		q1 := gqueue.New()
		q1.Push(1)
		q1.Push(2)
		// wait sync to channel
		time.Sleep(10 * time.Millisecond)
		t.Assert(q1.Len(), 2)
		q1.Close()
	})
	gtest.C(t, func(t *gtest.T) {
		q1 := gqueue.New(2)
		q1.Push(1)
		q1.Push(2)
		// wait sync to channel
		time.Sleep(10 * time.Millisecond)
		t.Assert(q1.Len(), 2)
		q1.Close()
	})
}

func Test_Issue2509(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		q := gqueue.New()
		q.Push(1)
		q.Push(2)
		q.Push(3)
		t.AssertLE(q.Len(), 3)
		t.Assert(<-q.C, 1)
		t.AssertLE(q.Len(), 2)
		t.Assert(<-q.C, 2)
		t.AssertLE(q.Len(), 1)
		t.Assert(<-q.C, 3)
		t.Assert(q.Len(), 0)
	})
}
