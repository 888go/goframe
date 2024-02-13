// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package 队列类_test

import (
	"testing"
	"time"
	
	"github.com/888go/goframe/container/gqueue"
	"github.com/888go/goframe/test/gtest"
)

func TestQueue_Len(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			maxNum   = 100
			maxTries = 100
		)
		for n := 10; n < maxTries; n++ {
			q1 := 队列类.X创建(maxNum)
			for i := 0; i < maxNum; i++ {
				q1.X入栈(i)
			}
			t.Assert(q1.X取长度(), maxNum)
			t.Assert(q1.Size弃用(), maxNum)
		}
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			maxNum   = 100
			maxTries = 100
		)
		for n := 10; n < maxTries; n++ {
			q1 := 队列类.X创建()
			for i := 0; i < maxNum; i++ {
				q1.X入栈(i)
			}
			t.AssertLE(q1.X取长度(), maxNum)
			t.AssertLE(q1.Size弃用(), maxNum)
		}
	})
}

func TestQueue_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		q := 队列类.X创建()
		for i := 0; i < 100; i++ {
			q.X入栈(i)
		}
		t.Assert(q.X出栈(), 0)
		t.Assert(q.X出栈(), 1)
	})
}

func TestQueue_Pop(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		q1 := 队列类.X创建()
		q1.X入栈(1)
		q1.X入栈(2)
		q1.X入栈(3)
		q1.X入栈(4)
		i1 := q1.X出栈()
		t.Assert(i1, 1)
	})
}

func TestQueue_Close(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		q1 := 队列类.X创建()
		q1.X入栈(1)
		q1.X入栈(2)
		time.Sleep(time.Millisecond)
		t.Assert(q1.X取长度(), 2)
		q1.X关闭()
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		q1 := 队列类.X创建(2)
		q1.X入栈(1)
		q1.X入栈(2)
		time.Sleep(time.Millisecond)
		t.Assert(q1.X取长度(), 2)
		q1.X关闭()
	})
}

func Test_Issue2509(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		q := 队列类.X创建()
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
