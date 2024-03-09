// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类_test

import (
	"context"
	"reflect"
	"testing"
	
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/888go/goframe/gutil"
)

var (
	ctx = context.TODO()
)

func Test_Try(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := `gutil Try test`
		t.Assert(工具类.X异常捕捉(ctx, func(ctx context.Context) {
			panic(s)
		}), s)
	})
	gtest.C(t, func(t *gtest.T) {
		s := `gutil Try test`
		t.Assert(工具类.X异常捕捉(ctx, func(ctx context.Context) {
			panic(gerror.New(s))
		}), s)
	})
}

func Test_TryCatch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		工具类.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
			panic("gutil TryCatch test")
		}, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		工具类.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
			panic("gutil TryCatch test")

		}, func(ctx context.Context, err error) {
			t.Assert(err, "gutil TryCatch test")
		})
	})

	gtest.C(t, func(t *gtest.T) {
		工具类.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
			panic(gerror.New("gutil TryCatch test"))

		}, func(ctx context.Context, err error) {
			t.Assert(err, "gutil TryCatch test")
		})
	})
}

func Test_IsEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X是否为空(1), false)
	})
}

func Test_Throw(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		defer func() {
			t.Assert(recover(), "gutil Throw test")
		}()

		工具类.X异常输出("gutil Throw test")
	})
}

func Test_Keys(t *testing.T) {
	// not support int
	gtest.C(t, func(t *gtest.T) {
		var val int = 1
		keys := 工具类.X取所有名称(reflect.ValueOf(val))
		t.AssertEQ(len(keys), 0)
	})
	// map
	gtest.C(t, func(t *gtest.T) {
		keys := 工具类.X取所有名称(map[int]int{
			1: 10,
			2: 20,
		})
		t.AssertIN("1", keys)
		t.AssertIN("2", keys)

		strKeys := 工具类.X取所有名称(map[string]interface{}{
			"key1": 1,
			"key2": 2,
		})
		t.AssertIN("key1", strKeys)
		t.AssertIN("key2", strKeys)
	})
	// *map
	gtest.C(t, func(t *gtest.T) {
		keys := 工具类.X取所有名称(&map[int]int{
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
		keys := 工具类.X取所有名称(new(T))
		t.Assert(keys, g.SliceStr{"A", "B"})
	})
	// *struct nil
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			A string
			B int
		}
		var pointer *T
		keys := 工具类.X取所有名称(pointer)
		t.Assert(keys, g.SliceStr{"A", "B"})
	})
	// **struct nil
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			A string
			B int
		}
		var pointer *T
		keys := 工具类.X取所有名称(&pointer)
		t.Assert(keys, g.SliceStr{"A", "B"})
	})
}

func Test_Values(t *testing.T) {
	// not support int
	gtest.C(t, func(t *gtest.T) {
		var val int = 1
		keys := 工具类.X取所有值(reflect.ValueOf(val))
		t.AssertEQ(len(keys), 0)
	})
	// map
	gtest.C(t, func(t *gtest.T) {
		values := 工具类.X取所有值(map[int]int{
			1: 10,
			2: 20,
		})
		t.AssertIN(10, values)
		t.AssertIN(20, values)

		values = 工具类.X取所有值(map[string]interface{}{
			"key1": 10,
			"key2": 20,
		})
		t.AssertIN(10, values)
		t.AssertIN(20, values)
	})
	// *map
	gtest.C(t, func(t *gtest.T) {
		keys := 工具类.X取所有值(&map[int]int{
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
		keys := 工具类.X取所有值(T{
			A: "1",
			B: 2,
		})
		t.Assert(keys, g.Slice{"1", 2})
	})
}

func TestListToMapByKey(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		listMap := []map[string]interface{}{
			{"key1": 1, "key2": 2},
			{"key3": 3, "key4": 4},
		}
		t.Assert(工具类.ListToMapByKey(listMap, "key1"), "{\"1\":{\"key1\":1,\"key2\":2}}")
	})
}

func Test_GetOrDefaultStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X取文本值或取默认值("a", "b"), "b")
		t.Assert(工具类.X取文本值或取默认值("a", "b", "c"), "b")
		t.Assert(工具类.X取文本值或取默认值("a"), "a")
	})
}

func Test_GetOrDefaultAny(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X取值或取默认值("a", "b"), "b")
		t.Assert(工具类.X取值或取默认值("a", "b", "c"), "b")
		t.Assert(工具类.X取值或取默认值("a"), "a")
	})
}
