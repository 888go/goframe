// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 协程类_test

import (
	"context"
	"sync"
	"testing"
	"time"

	garray "github.com/888go/goframe/container/garray"
	grpool "github.com/888go/goframe/os/grpool"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			err   error
			wg    = sync.WaitGroup{}
			array = garray.NewArray别名(true)
			size  = 100
		)
		wg.Add(size)
		for i := 0; i < size; i++ {
			err = grpool.Add(ctx, func(ctx context.Context) {
				array.Append别名(1)
				wg.Done()
			})
			t.AssertNil(err)
		}
		wg.Wait()

		time.Sleep(100 * time.Millisecond)

		t.Assert(array.X取长度(), size)
		t.Assert(grpool.Jobs(), 0)
		t.Assert(grpool.Size(), 0)
	})
}

func Test_Limit1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			wg    = sync.WaitGroup{}
			array = garray.NewArray别名(true)
			size  = 100
			pool  = grpool.New(10)
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
	gtest.C(t, func(t *gtest.T) {
		var (
			err   error
			wg    = sync.WaitGroup{}
			array = garray.NewArray别名(true)
			size  = 100
			pool  = grpool.New(1)
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
	gtest.C(t, func(t *gtest.T) {
		var (
			array = garray.NewArray别名(true)
			size  = 1000
			pool  = grpool.New(100)
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
	gtest.C(t, func(t *gtest.T) {
		var (
			err   error
			array = garray.NewArray别名(true)
		)
		err = grpool.AddWithRecover(ctx, func(ctx context.Context) {
			array.Append别名(1)
			panic(1)
		}, func(ctx context.Context, err error) {
			array.Append别名(1)
		})
		t.AssertNil(err)
		err = grpool.AddWithRecover(ctx, func(ctx context.Context) {
			panic(1)
			array.Append别名(1)
		}, nil)
		t.AssertNil(err)

		time.Sleep(500 * time.Millisecond)

		t.Assert(array.X取长度(), 2)
	})
}
