// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package g_test

import (
	"context"
	"os"
	"sync"
	"testing"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gutil"
)

var (
	ctx = context.TODO()
)

func Test_NewVar(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(1).X取整数(), 1)
		t.Assert(g.X泛型类(1, true).X取整数(), 1)
	})
}

func Test_Dump(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		g.X调试输出("GoFrame")
	})
}

func Test_DumpTo(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		g.X调试输出到Writer(os.Stdout, "GoFrame", 工具类.DumpOption{})
	})
}

func Test_DumpWithType(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		g.X调试输出并带类型("GoFrame", 123)
	})
}

func Test_DumpWithOption(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		g.X调试输出并带选项("GoFrame", 工具类.DumpOption{})
	})
}

func Test_Try(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		g.X异常捕捉(ctx, func(ctx context.Context) {
			g.X调试输出("GoFrame")
		})
	})
}

func Test_TryCatch(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		g.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
			g.X调试输出("GoFrame")
		}, func(ctx context.Context, exception error) {
			g.X调试输出(exception)
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		g.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
			g.X异常输出("GoFrame")
		}, func(ctx context.Context, exception error) {
			t.Assert(exception.Error(), "GoFrame")
		})
	})
}

func Test_IsNil(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X是否为Nil(nil), true)
		t.Assert(g.X是否为Nil(0), false)
		t.Assert(g.X是否为Nil("GoFrame"), false)
	})
}

func Test_IsEmpty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X是否为空(nil), true)
		t.Assert(g.X是否为空(0), true)
		t.Assert(g.X是否为空("GoFrame"), false)
	})
}

func Test_SetDebug(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		g.X设置debug(true)
	})
}

func Test_Object(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			wg    = sync.WaitGroup{}
			array = 数组类.NewArray别名(true)
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
