		// 版权归GoFrame作者(https:		//goframe.org)所有。保留所有权利。
		//
		// 本源代码形式受MIT许可证条款约束。
		// 如果未随本文件一同分发MIT许可证副本，
		// 您可以在https:		//github.com/gogf/gf处获取。
		// md5:a9832f33b234e3f3

package gutil_test

import (
	"context"
	"sync"
	"testing"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gutil"
)

func Test_Go(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			wg    = sync.WaitGroup{}
			array = garray.NewArray(true)
		)
		wg.Add(1)
		gutil.Go(ctx, func(ctx context.Context) {
			defer wg.Done()
			array.Append(1)
		}, nil)
		wg.Wait()
		t.Assert(array.Len(), 1)
	})
	// recover
	gtest.C(t, func(t *gtest.T) {
		var (
			wg    = sync.WaitGroup{}
			array = garray.NewArray(true)
		)
		wg.Add(1)
		gutil.Go(ctx, func(ctx context.Context) {
			defer wg.Done()
			panic("error")
			array.Append(1)
		}, nil)
		wg.Wait()
		t.Assert(array.Len(), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			wg    = sync.WaitGroup{}
			array = garray.NewArray(true)
		)
		wg.Add(1)
		gutil.Go(ctx, func(ctx context.Context) {
			panic("error")
		}, func(ctx context.Context, exception error) {
			defer wg.Done()
			array.Append(exception)
		})
		wg.Wait()
		t.Assert(array.Len(), 1)
		t.Assert(array.At(0), "error")
	})
}
