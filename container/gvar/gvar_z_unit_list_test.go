// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gvar_test

import (
	"testing"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestVar_ListItemValues_Map(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		listMap := g.List{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 99},
			g.Map{"id": 3, "score": 99},
		}
		t.Assert(gvar.New(listMap).ListItemValues("id"), g.Slice{1, 2, 3})
		t.Assert(gvar.New(listMap).ListItemValues("score"), g.Slice{100, 99, 99})
	})
	gtest.C(t, func(t *gtest.T) {
		listMap := g.List{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": nil},
			g.Map{"id": 3, "score": 0},
		}
		t.Assert(gvar.New(listMap).ListItemValues("id"), g.Slice{1, 2, 3})
		t.Assert(gvar.New(listMap).ListItemValues("score"), g.Slice{100, nil, 0})
	})
}

func TestVar_ListItemValues_Struct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			Id    int
			Score float64
		}
		listStruct := g.Slice{
			T{1, 100},
			T{2, 99},
			T{3, 0},
		}
		t.Assert(gvar.New(listStruct).ListItemValues("Id"), g.Slice{1, 2, 3})
		t.Assert(gvar.New(listStruct).ListItemValues("Score"), g.Slice{100, 99, 0})
	})
	// Pointer items.
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			Id    int
			Score float64
		}
		listStruct := g.Slice{
			&T{1, 100},
			&T{2, 99},
			&T{3, 0},
		}
		t.Assert(gvar.New(listStruct).ListItemValues("Id"), g.Slice{1, 2, 3})
		t.Assert(gvar.New(listStruct).ListItemValues("Score"), g.Slice{100, 99, 0})
	})
	// Nil element value.
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			Id    int
			Score interface{}
		}
		listStruct := g.Slice{
			T{1, 100},
			T{2, nil},
			T{3, 0},
		}
		t.Assert(gvar.New(listStruct).ListItemValues("Id"), g.Slice{1, 2, 3})
		t.Assert(gvar.New(listStruct).ListItemValues("Score"), g.Slice{100, nil, 0})
	})
}

func TestVar_ListItemValuesUnique(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		listMap := g.List{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 100},
			g.Map{"id": 3, "score": 100},
			g.Map{"id": 4, "score": 100},
			g.Map{"id": 5, "score": 100},
		}
		t.Assert(gvar.New(listMap).ListItemValuesUnique("id"), g.Slice{1, 2, 3, 4, 5})
		t.Assert(gvar.New(listMap).ListItemValuesUnique("score"), g.Slice{100})
	})
	gtest.C(t, func(t *gtest.T) {
		listMap := g.List{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 100},
			g.Map{"id": 3, "score": 100},
			g.Map{"id": 4, "score": 100},
			g.Map{"id": 5, "score": 99},
		}
		t.Assert(gvar.New(listMap).ListItemValuesUnique("id"), g.Slice{1, 2, 3, 4, 5})
		t.Assert(gvar.New(listMap).ListItemValuesUnique("score"), g.Slice{100, 99})
	})
	gtest.C(t, func(t *gtest.T) {
		listMap := g.List{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 100},
			g.Map{"id": 3, "score": 0},
			g.Map{"id": 4, "score": 100},
			g.Map{"id": 5, "score": 99},
		}
		t.Assert(gvar.New(listMap).ListItemValuesUnique("id"), g.Slice{1, 2, 3, 4, 5})
		t.Assert(gvar.New(listMap).ListItemValuesUnique("score"), g.Slice{100, 0, 99})
	})
}
