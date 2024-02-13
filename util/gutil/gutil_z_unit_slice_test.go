// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类_test

import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gutil"
)

func Test_SliceCopy(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Slice别名{
			"K1", "v1", "K2", "v2",
		}
		s1 := 工具类.SliceCopy(s)
		t.Assert(s, s1)
	})
}

func Test_SliceDelete(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Slice别名{
			"K1", "v1", "K2", "v2",
		}
		t.Assert(工具类.SliceDelete(s, 0), g.Slice别名{
			"v1", "K2", "v2",
		})
		t.Assert(工具类.SliceDelete(s, 5), s)
	})
}

func Test_SliceToMap(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Slice别名{
			"K1", "v1", "K2", "v2",
		}
		m := 工具类.SliceToMap(s)
		t.Assert(len(m), 2)
		t.Assert(m, g.Map{
			"K1": "v1",
			"K2": "v2",
		})

		m1 := 工具类.SliceToMap(&s)
		t.Assert(len(m1), 2)
		t.Assert(m1, g.Map{
			"K1": "v1",
			"K2": "v2",
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Slice别名{
			"K1", "v1", "K2",
		}
		m := 工具类.SliceToMap(s)
		t.Assert(len(m), 0)
		t.Assert(m, nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		m := 工具类.SliceToMap(1)
		t.Assert(len(m), 0)
		t.Assert(m, nil)
	})
}

func Test_SliceToMapWithColumnAsKey(t *testing.T) {
	m1 := g.Map{"K1": "v1", "K2": 1}
	m2 := g.Map{"K1": "v2", "K2": 2}
	s := g.Slice别名{m1, m2}
	单元测试类.C(t, func(t *单元测试类.T) {
		m := 工具类.SliceToMapWithColumnAsKey(s, "K1")
		t.Assert(m, g.MapAnyAny{
			"v1": m1,
			"v2": m2,
		})

		n := 工具类.SliceToMapWithColumnAsKey(&s, "K1")
		t.Assert(n, g.MapAnyAny{
			"v1": m1,
			"v2": m2,
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		m := 工具类.SliceToMapWithColumnAsKey(s, "K2")
		t.Assert(m, g.MapAnyAny{
			1: m1,
			2: m2,
		})
	})
}
