// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类_test

import (
	"context"
	"reflect"
	"testing"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gutil"
)

var (
	ctx = context.TODO()
)

func Test_Try(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := `gutil Try test`
		t.Assert(工具类.X异常捕捉(ctx, func(ctx context.Context) {
			panic(s)
		}), s)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		s := `gutil Try test`
		t.Assert(工具类.X异常捕捉(ctx, func(ctx context.Context) {
			panic(错误类.X创建(s))
		}), s)
	})
}

func Test_TryCatch(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		工具类.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
			panic("gutil TryCatch test")
		}, nil)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		工具类.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
			panic("gutil TryCatch test")

		}, func(ctx context.Context, err error) {
			t.Assert(err, "gutil TryCatch test")
		})
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		工具类.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
			panic(错误类.X创建("gutil TryCatch test"))

		}, func(ctx context.Context, err error) {
			t.Assert(err, "gutil TryCatch test")
		})
	})
}

func Test_IsEmpty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(工具类.X是否为空(1), false)
	})
}

func Test_Throw(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			t.Assert(recover(), "gutil Throw test")
		}()

		工具类.X异常输出("gutil Throw test")
	})
}

func Test_Keys(t *testing.T) {
	// not support int
	单元测试类.C(t, func(t *单元测试类.T) {
		var val int = 1
		keys := 工具类.X取所有名称(reflect.ValueOf(val))
		t.AssertEQ(len(keys), 0)
	})
	// map
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		keys := 工具类.X取所有名称(&map[int]int{
			1: 10,
			2: 20,
		})
		t.AssertIN("1", keys)
		t.AssertIN("2", keys)
	})
	// *struct
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			A string
			B int
		}
		keys := 工具类.X取所有名称(new(T))
		t.Assert(keys, g.SliceStr别名{"A", "B"})
	})
	// *struct nil
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			A string
			B int
		}
		var pointer *T
		keys := 工具类.X取所有名称(pointer)
		t.Assert(keys, g.SliceStr别名{"A", "B"})
	})
	// **struct nil
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			A string
			B int
		}
		var pointer *T
		keys := 工具类.X取所有名称(&pointer)
		t.Assert(keys, g.SliceStr别名{"A", "B"})
	})
}

func Test_Values(t *testing.T) {
	// not support int
	单元测试类.C(t, func(t *单元测试类.T) {
		var val int = 1
		keys := 工具类.X取所有值(reflect.ValueOf(val))
		t.AssertEQ(len(keys), 0)
	})
	// map
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		keys := 工具类.X取所有值(&map[int]int{
			1: 10,
			2: 20,
		})
		t.AssertIN(10, keys)
		t.AssertIN(20, keys)
	})
	// struct
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			A string
			B int
		}
		keys := 工具类.X取所有值(T{
			A: "1",
			B: 2,
		})
		t.Assert(keys, g.Slice别名{"1", 2})
	})
}

func TestListToMapByKey(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		listMap := []map[string]interface{}{
			{"key1": 1, "key2": 2},
			{"key3": 3, "key4": 4},
		}
		t.Assert(工具类.ListToMapByKey(listMap, "key1"), "{\"1\":{\"key1\":1,\"key2\":2}}")
	})
}

func Test_GetOrDefaultStr(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(工具类.X取文本值或取默认值("a", "b"), "b")
		t.Assert(工具类.X取文本值或取默认值("a", "b", "c"), "b")
		t.Assert(工具类.X取文本值或取默认值("a"), "a")
	})
}

func Test_GetOrDefaultAny(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(工具类.X取值或取默认值("a", "b"), "b")
		t.Assert(工具类.X取值或取默认值("a", "b", "c"), "b")
		t.Assert(工具类.X取值或取默认值("a"), "a")
	})
}
