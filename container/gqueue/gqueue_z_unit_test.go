// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package gqueue_test
import (
	"testing"
	"time"
	
	"github.com/888go/goframe/container/gqueue"
	"github.com/888go/goframe/test/gtest"
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
		time.Sleep(time.Millisecond)
		t.Assert(q1.Len(), 2)
		q1.Close()
	})
	gtest.C(t, func(t *gtest.T) {
		q1 := gqueue.New(2)
		q1.Push(1)
		q1.Push(2)
		time.Sleep(time.Millisecond)
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
