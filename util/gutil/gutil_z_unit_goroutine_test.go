// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 工具类_test

import (
	"context"
	"sync"
	"testing"

	garray "github.com/888go/goframe/container/garray"
	gtest "github.com/888go/goframe/test/gtest"
	gutil "github.com/888go/goframe/util/gutil"
)

func Test_Go(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			wg    = sync.WaitGroup{}
			array = garray.NewArray别名(true)
		)
		wg.Add(1)
		gutil.Go(ctx, func(ctx context.Context) {
			defer wg.Done()
			array.Append别名(1)
		}, nil)
		wg.Wait()
		t.Assert(array.X取长度(), 1)
	})
	// recover
	gtest.C(t, func(t *gtest.T) {
		var (
			wg    = sync.WaitGroup{}
			array = garray.NewArray别名(true)
		)
		wg.Add(1)
		gutil.Go(ctx, func(ctx context.Context) {
			defer wg.Done()
			panic("error")
			array.Append别名(1)
		}, nil)
		wg.Wait()
		t.Assert(array.X取长度(), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			wg    = sync.WaitGroup{}
			array = garray.NewArray别名(true)
		)
		wg.Add(1)
		gutil.Go(ctx, func(ctx context.Context) {
			panic("error")
		}, func(ctx context.Context, exception error) {
			defer wg.Done()
			array.Append别名(exception)
		})
		wg.Wait()
		t.Assert(array.X取长度(), 1)
		t.Assert(array.X取值(0), "error")
	})
}
