// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 泛型类_test

import (
	"testing"

	gvar "github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
)

func TestVar_ListItemValues_Map(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		listMap := g.Map切片{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 99},
			g.Map{"id": 3, "score": 99},
		}
		t.Assert(gvar.X创建(listMap).X取结构切片或Map切片值("id"), g.Slice别名{1, 2, 3})
		t.Assert(gvar.X创建(listMap).X取结构切片或Map切片值("score"), g.Slice别名{100, 99, 99})
	})
	gtest.C(t, func(t *gtest.T) {
		listMap := g.Map切片{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": nil},
			g.Map{"id": 3, "score": 0},
		}
		t.Assert(gvar.X创建(listMap).X取结构切片或Map切片值("id"), g.Slice别名{1, 2, 3})
		t.Assert(gvar.X创建(listMap).X取结构切片或Map切片值("score"), g.Slice别名{100, nil, 0})
	})
}

func TestVar_ListItemValues_Struct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			Id    int
			Score float64
		}
		listStruct := g.Slice别名{
			T{1, 100},
			T{2, 99},
			T{3, 0},
		}
		t.Assert(gvar.X创建(listStruct).X取结构切片或Map切片值("Id"), g.Slice别名{1, 2, 3})
		t.Assert(gvar.X创建(listStruct).X取结构切片或Map切片值("Score"), g.Slice别名{100, 99, 0})
	})
	// Pointer items.
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			Id    int
			Score float64
		}
		listStruct := g.Slice别名{
			&T{1, 100},
			&T{2, 99},
			&T{3, 0},
		}
		t.Assert(gvar.X创建(listStruct).X取结构切片或Map切片值("Id"), g.Slice别名{1, 2, 3})
		t.Assert(gvar.X创建(listStruct).X取结构切片或Map切片值("Score"), g.Slice别名{100, 99, 0})
	})
	// Nil element value.
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			Id    int
			Score interface{}
		}
		listStruct := g.Slice别名{
			T{1, 100},
			T{2, nil},
			T{3, 0},
		}
		t.Assert(gvar.X创建(listStruct).X取结构切片或Map切片值("Id"), g.Slice别名{1, 2, 3})
		t.Assert(gvar.X创建(listStruct).X取结构切片或Map切片值("Score"), g.Slice别名{100, nil, 0})
	})
}

func TestVar_ListItemValuesUnique(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		listMap := g.Map切片{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 100},
			g.Map{"id": 3, "score": 100},
			g.Map{"id": 4, "score": 100},
			g.Map{"id": 5, "score": 100},
		}
		t.Assert(gvar.X创建(listMap).X取结构切片或Map切片值并去重("id"), g.Slice别名{1, 2, 3, 4, 5})
		t.Assert(gvar.X创建(listMap).X取结构切片或Map切片值并去重("score"), g.Slice别名{100})
	})
	gtest.C(t, func(t *gtest.T) {
		listMap := g.Map切片{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 100},
			g.Map{"id": 3, "score": 100},
			g.Map{"id": 4, "score": 100},
			g.Map{"id": 5, "score": 99},
		}
		t.Assert(gvar.X创建(listMap).X取结构切片或Map切片值并去重("id"), g.Slice别名{1, 2, 3, 4, 5})
		t.Assert(gvar.X创建(listMap).X取结构切片或Map切片值并去重("score"), g.Slice别名{100, 99})
	})
	gtest.C(t, func(t *gtest.T) {
		listMap := g.Map切片{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 100},
			g.Map{"id": 3, "score": 0},
			g.Map{"id": 4, "score": 100},
			g.Map{"id": 5, "score": 99},
		}
		t.Assert(gvar.X创建(listMap).X取结构切片或Map切片值并去重("id"), g.Slice别名{1, 2, 3, 4, 5})
		t.Assert(gvar.X创建(listMap).X取结构切片或Map切片值并去重("score"), g.Slice别名{100, 0, 99})
	})
}
