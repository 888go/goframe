// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 协程类_test

import (
	"context"
	"sync"
	"testing"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/os/grpool"
	"github.com/888go/goframe/test/gtest"
)

func Test_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			err   error
			wg    = sync.WaitGroup{}
			array = 数组类.NewArray别名(true)
			size  = 100
		)
		wg.Add(size)
		for i := 0; i < size; i++ {
			err = 协程类.Add(ctx, func(ctx context.Context) {
				array.Append别名(1)
				wg.Done()
			})
			t.AssertNil(err)
		}
		wg.Wait()

		time.Sleep(100 * time.Millisecond)

		t.Assert(array.X取长度(), size)
		t.Assert(协程类.Jobs(), 0)
		t.Assert(协程类.Size(), 0)
	})
}

func Test_Limit1(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			wg    = sync.WaitGroup{}
			array = 数组类.NewArray别名(true)
			size  = 100
			pool  = 协程类.New(10)
		)
		wg.Add(size)
		for i := 0; i < size; i++ {
			pool.Add(ctx, func(ctx context.Context) {
				array.Append别名(1)
				wg.Done()
			})
		}
		wg.Wait()
		t.Assert(array.X取长度(), size)
	})
}

func Test_Limit2(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			err   error
			wg    = sync.WaitGroup{}
			array = 数组类.NewArray别名(true)
			size  = 100
			pool  = 协程类.New(1)
		)
		wg.Add(size)
		for i := 0; i < size; i++ {
			err = pool.Add(ctx, func(ctx context.Context) {
				defer wg.Done()
				array.Append别名(1)
			})
			t.AssertNil(err)
		}
		wg.Wait()
		t.Assert(array.X取长度(), size)
	})
}

func Test_Limit3(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			array = 数组类.NewArray别名(true)
			size  = 1000
			pool  = 协程类.New(100)
		)
		t.Assert(pool.Cap(), 100)
		for i := 0; i < size; i++ {
			pool.Add(ctx, func(ctx context.Context) {
				array.Append别名(1)
				time.Sleep(2 * time.Second)
			})
		}
		time.Sleep(time.Second)
		t.Assert(pool.Size(), 100)
		t.Assert(pool.Jobs(), 900)
		t.Assert(array.X取长度(), 100)
		pool.Close()
		time.Sleep(2 * time.Second)
		t.Assert(pool.Size(), 0)
		t.Assert(pool.Jobs(), 900)
		t.Assert(array.X取长度(), 100)
		t.Assert(pool.IsClosed(), true)
		t.AssertNE(pool.Add(ctx, func(ctx context.Context) {}), nil)
	})
}

func Test_AddWithRecover(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			err   error
			array = 数组类.NewArray别名(true)
		)
		err = 协程类.AddWithRecover(ctx, func(ctx context.Context) {
			array.Append别名(1)
			panic(1)
		}, func(ctx context.Context, err error) {
			array.Append别名(1)
		})
		t.AssertNil(err)
		err = 协程类.AddWithRecover(ctx, func(ctx context.Context) {
			panic(1)
			array.Append别名(1)
		}, nil)
		t.AssertNil(err)

		time.Sleep(500 * time.Millisecond)

		t.Assert(array.X取长度(), 2)
	})
}
