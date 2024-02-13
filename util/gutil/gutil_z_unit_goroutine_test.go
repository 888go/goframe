// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类_test

import (
	"context"
	"sync"
	"testing"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gutil"
)

func Test_Go(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			wg    = sync.WaitGroup{}
			array = 数组类.NewArray别名(true)
		)
		wg.Add(1)
		工具类.Go(ctx, func(ctx context.Context) {
			defer wg.Done()
			array.Append别名(1)
		}, nil)
		wg.Wait()
		t.Assert(array.X取长度(), 1)
	})
	// recover
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			wg    = sync.WaitGroup{}
			array = 数组类.NewArray别名(true)
		)
		wg.Add(1)
		工具类.Go(ctx, func(ctx context.Context) {
			defer wg.Done()
			panic("error")
			array.Append别名(1)
		}, nil)
		wg.Wait()
		t.Assert(array.X取长度(), 0)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			wg    = sync.WaitGroup{}
			array = 数组类.NewArray别名(true)
		)
		wg.Add(1)
		工具类.Go(ctx, func(ctx context.Context) {
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
