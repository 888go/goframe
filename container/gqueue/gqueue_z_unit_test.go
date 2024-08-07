// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package 队列类_test

import (
	"testing"
	"time"

	gqueue "github.com/888go/goframe/container/gqueue"
	gtest "github.com/888go/goframe/test/gtest"
)

func TestQueue_Len(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			maxNum   = 100
			maxTries = 100
		)
		for n := 10; n < maxTries; n++ {
			q1 := gqueue.X创建(maxNum)
			for i := 0; i < maxNum; i++ {
				q1.X入栈(i)
			}
			t.Assert(q1.X取长度(), maxNum)
			t.Assert(q1.Size弃用(), maxNum)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			maxNum   = 100
			maxTries = 100
		)
		for n := 10; n < maxTries; n++ {
			q1 := gqueue.X创建()
			for i := 0; i < maxNum; i++ {
				q1.X入栈(i)
			}
			t.AssertLE(q1.X取长度(), maxNum)
			t.AssertLE(q1.Size弃用(), maxNum)
		}
	})
}

func TestQueue_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		q := gqueue.X创建()
		for i := 0; i < 100; i++ {
			q.X入栈(i)
		}
		t.Assert(q.X出栈(), 0)
		t.Assert(q.X出栈(), 1)
	})
}

func TestQueue_Pop(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		q1 := gqueue.X创建()
		q1.X入栈(1)
		q1.X入栈(2)
		q1.X入栈(3)
		q1.X入栈(4)
		i1 := q1.X出栈()
		t.Assert(i1, 1)
	})
}

func TestQueue_Close(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		q1 := gqueue.X创建()
		q1.X入栈(1)
		q1.X入栈(2)
		// wait sync to channel
		time.Sleep(10 * time.Millisecond)
		t.Assert(q1.X取长度(), 2)
		q1.X关闭()
	})
	gtest.C(t, func(t *gtest.T) {
		q1 := gqueue.X创建(2)
		q1.X入栈(1)
		q1.X入栈(2)
		// wait sync to channel
		time.Sleep(10 * time.Millisecond)
		t.Assert(q1.X取长度(), 2)
		q1.X关闭()
	})
}

func Test_Issue2509(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		q := gqueue.X创建()
		q.X入栈(1)
		q.X入栈(2)
		q.X入栈(3)
		t.AssertLE(q.X取长度(), 3)
		t.Assert(<-q.C, 1)
		t.AssertLE(q.X取长度(), 2)
		t.Assert(<-q.C, 2)
		t.AssertLE(q.X取长度(), 1)
		t.Assert(<-q.C, 3)
		t.Assert(q.X取长度(), 0)
	})
}
