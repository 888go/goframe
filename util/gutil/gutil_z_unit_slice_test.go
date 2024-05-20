// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gutil_test

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gutil"
)

func Test_SliceCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Slice{
			"K1", "v1", "K2", "v2",
		}
		s1 := gutil.SliceCopy(s)
		t.Assert(s, s1)
	})
}

func Test_SliceDelete(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Slice{
			"K1", "v1", "K2", "v2",
		}
		t.Assert(gutil.SliceDelete(s, 0), g.Slice{
			"v1", "K2", "v2",
		})
		t.Assert(gutil.SliceDelete(s, 5), s)
	})
}

func Test_SliceToMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Slice{
			"K1", "v1", "K2", "v2",
		}
		m := gutil.SliceToMap(s)
		t.Assert(len(m), 2)
		t.Assert(m, g.Map{
			"K1": "v1",
			"K2": "v2",
		})

		m1 := gutil.SliceToMap(&s)
		t.Assert(len(m1), 2)
		t.Assert(m1, g.Map{
			"K1": "v1",
			"K2": "v2",
		})
	})
	gtest.C(t, func(t *gtest.T) {
		s := g.Slice{
			"K1", "v1", "K2",
		}
		m := gutil.SliceToMap(s)
		t.Assert(len(m), 0)
		t.Assert(m, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		m := gutil.SliceToMap(1)
		t.Assert(len(m), 0)
		t.Assert(m, nil)
	})
}

func Test_SliceToMapWithColumnAsKey(t *testing.T) {
	m1 := g.Map{"K1": "v1", "K2": 1}
	m2 := g.Map{"K1": "v2", "K2": 2}
	s := g.Slice{m1, m2}
	gtest.C(t, func(t *gtest.T) {
		m := gutil.SliceToMapWithColumnAsKey(s, "K1")
		t.Assert(m, g.MapAnyAny{
			"v1": m1,
			"v2": m2,
		})

		n := gutil.SliceToMapWithColumnAsKey(&s, "K1")
		t.Assert(n, g.MapAnyAny{
			"v1": m1,
			"v2": m2,
		})
	})
	gtest.C(t, func(t *gtest.T) {
		m := gutil.SliceToMapWithColumnAsKey(s, "K2")
		t.Assert(m, g.MapAnyAny{
			1: m1,
			2: m2,
		})
	})
}
