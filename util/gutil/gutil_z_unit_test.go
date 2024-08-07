// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 工具类_test

import (
	"context"
	"reflect"
	"testing"

	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
	gutil "github.com/888go/goframe/util/gutil"
)

var (
	ctx = context.TODO()
)

func Test_Try(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := `gutil Try test`
		t.Assert(gutil.X异常捕捉(ctx, func(ctx context.Context) {
			panic(s)
		}), s)
	})
	gtest.C(t, func(t *gtest.T) {
		s := `gutil Try test`
		t.Assert(gutil.X异常捕捉(ctx, func(ctx context.Context) {
			panic(gerror.X创建(s))
		}), s)
	})
}

func Test_TryCatch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		gutil.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
			panic("gutil TryCatch test")
		}, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		gutil.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
			panic("gutil TryCatch test")

		}, func(ctx context.Context, err error) {
			t.Assert(err, "gutil TryCatch test")
		})
	})

	gtest.C(t, func(t *gtest.T) {
		gutil.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
			panic(gerror.X创建("gutil TryCatch test"))

		}, func(ctx context.Context, err error) {
			t.Assert(err, "gutil TryCatch test")
		})
	})
}

func Test_Throw(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		defer func() {
			t.Assert(recover(), "gutil Throw test")
		}()

		gutil.X异常输出("gutil Throw test")
	})
}

func Test_Keys(t *testing.T) {
	// not support int
	gtest.C(t, func(t *gtest.T) {
		var val int = 1
		keys := gutil.X取所有名称(reflect.ValueOf(val))
		t.AssertEQ(len(keys), 0)
	})
	// map
	gtest.C(t, func(t *gtest.T) {
		keys := gutil.X取所有名称(map[int]int{
			1: 10,
			2: 20,
		})
		t.AssertIN("1", keys)
		t.AssertIN("2", keys)

		strKeys := gutil.X取所有名称(map[string]interface{}{
			"key1": 1,
			"key2": 2,
		})
		t.AssertIN("key1", strKeys)
		t.AssertIN("key2", strKeys)
	})
	// *map
	gtest.C(t, func(t *gtest.T) {
		keys := gutil.X取所有名称(&map[int]int{
			1: 10,
			2: 20,
		})
		t.AssertIN("1", keys)
		t.AssertIN("2", keys)
	})
	// *struct
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			A string
			B int
		}
		keys := gutil.X取所有名称(new(T))
		t.Assert(keys, g.SliceStr别名{"A", "B"})
	})
	// *struct nil
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			A string
			B int
		}
		var pointer *T
		keys := gutil.X取所有名称(pointer)
		t.Assert(keys, g.SliceStr别名{"A", "B"})
	})
	// **struct nil
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			A string
			B int
		}
		var pointer *T
		keys := gutil.X取所有名称(&pointer)
		t.Assert(keys, g.SliceStr别名{"A", "B"})
	})
}

func Test_Values(t *testing.T) {
	// not support int
	gtest.C(t, func(t *gtest.T) {
		var val int = 1
		keys := gutil.X取所有值(reflect.ValueOf(val))
		t.AssertEQ(len(keys), 0)
	})
	// map
	gtest.C(t, func(t *gtest.T) {
		values := gutil.X取所有值(map[int]int{
			1: 10,
			2: 20,
		})
		t.AssertIN(10, values)
		t.AssertIN(20, values)

		values = gutil.X取所有值(map[string]interface{}{
			"key1": 10,
			"key2": 20,
		})
		t.AssertIN(10, values)
		t.AssertIN(20, values)
	})
	// *map
	gtest.C(t, func(t *gtest.T) {
		keys := gutil.X取所有值(&map[int]int{
			1: 10,
			2: 20,
		})
		t.AssertIN(10, keys)
		t.AssertIN(20, keys)
	})
	// struct
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			A string
			B int
		}
		keys := gutil.X取所有值(T{
			A: "1",
			B: 2,
		})
		t.Assert(keys, g.Slice别名{"1", 2})
	})
}

func TestListToMapByKey(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		listMap := []map[string]interface{}{
			{"key1": 1, "key2": 2},
			{"key3": 3, "key4": 4},
		}
		t.Assert(gutil.ListToMapByKey(listMap, "key1"), "{\"1\":{\"key1\":1,\"key2\":2}}")
	})
}

func Test_GetOrDefaultStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gutil.X取文本值或取默认值("a", "b"), "b")
		t.Assert(gutil.X取文本值或取默认值("a", "b", "c"), "b")
		t.Assert(gutil.X取文本值或取默认值("a"), "a")
	})
}

func Test_GetOrDefaultAny(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gutil.X取值或取默认值("a", "b"), "b")
		t.Assert(gutil.X取值或取默认值("a", "b", "c"), "b")
		t.Assert(gutil.X取值或取默认值("a"), "a")
	})
}
