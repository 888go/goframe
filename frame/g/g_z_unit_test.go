// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package g_test

import (
	"context"
	"os"
	"sync"
	"testing"

	garray "github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
	gutil "github.com/888go/goframe/util/gutil"
)

var (
	ctx = context.TODO()
)

func Test_NewVar(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(1).X取整数(), 1)
		t.Assert(g.X泛型类(1, true).X取整数(), 1)
	})
}

func Test_Dump(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		g.X调试输出("GoFrame")
	})
}

func Test_DumpTo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		g.X调试输出到Writer(os.Stdout, "GoFrame", gutil.DumpOption{})
	})
}

func Test_DumpWithType(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		g.X调试输出并带类型("GoFrame", 123)
	})
}

func Test_DumpWithOption(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		g.X调试输出并带选项("GoFrame", gutil.DumpOption{})
	})
}

func Test_Try(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		g.X异常捕捉(ctx, func(ctx context.Context) {
			g.X调试输出("GoFrame")
		})
	})
}

func Test_TryCatch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		g.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
			g.X调试输出("GoFrame")
		}, func(ctx context.Context, exception error) {
			g.X调试输出(exception)
		})
	})
	gtest.C(t, func(t *gtest.T) {
		g.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
			g.X异常输出("GoFrame")
		}, func(ctx context.Context, exception error) {
			t.Assert(exception.Error(), "GoFrame")
		})
	})
}

func Test_IsNil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X是否为Nil(nil), true)
		t.Assert(g.X是否为Nil(0), false)
		t.Assert(g.X是否为Nil("GoFrame"), false)
	})
}

func Test_IsEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X是否为空(nil), true)
		t.Assert(g.X是否为空(0), true)
		t.Assert(g.X是否为空("GoFrame"), false)
	})
}

func Test_SetDebug(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		g.X设置debug(true)
	})
}

func Test_Object(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(g.X网页类(), nil)
		t.AssertNE(g.Http类(), nil)
		t.AssertNE(g.TCP类(), nil)
		t.AssertNE(g.UDP类(), nil)
		t.AssertNE(g.X模板类(), nil)
		t.AssertNE(g.X配置类(), nil)
		t.AssertNE(g.Cfg别名(), nil)
		t.AssertNE(g.X资源类(), nil)
		t.AssertNE(g.X多语言类(), nil)
		t.AssertNE(g.Res别名(), nil)
		t.AssertNE(g.X日志类(), nil)
		t.AssertNE(g.X效验类(), nil)
	})
}

func Test_Go(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			wg    = sync.WaitGroup{}
			array = garray.NewArray别名(true)
		)
		wg.Add(1)
		g.Go(context.Background(), func(ctx context.Context) {
			defer wg.Done()
			array.Append别名(1)
		}, nil)
		wg.Wait()
		t.Assert(array.X取长度(), 1)
	})
}
