// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类_test

import (
	"context"
	"sync"
	"testing"
	
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/888go/goframe/gutil"
)

func Test_Go(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			wg    = sync.WaitGroup{}
			array = garray.NewArray(true)
		)
		wg.Add(1)
		工具类.Go(ctx, func(ctx context.Context) {
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
		工具类.Go(ctx, func(ctx context.Context) {
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
		工具类.Go(ctx, func(ctx context.Context) {
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
