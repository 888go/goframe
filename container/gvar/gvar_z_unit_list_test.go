// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 泛型类_test

import (
	"testing"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
)

func TestVar_ListItemValues_Map(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		listMap := g.Map数组{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 99},
			g.Map{"id": 3, "score": 99},
		}
		t.Assert(泛型类.X创建(listMap).X取结构数组或Map数组值("id"), g.Slice别名{1, 2, 3})
		t.Assert(泛型类.X创建(listMap).X取结构数组或Map数组值("score"), g.Slice别名{100, 99, 99})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		listMap := g.Map数组{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": nil},
			g.Map{"id": 3, "score": 0},
		}
		t.Assert(泛型类.X创建(listMap).X取结构数组或Map数组值("id"), g.Slice别名{1, 2, 3})
		t.Assert(泛型类.X创建(listMap).X取结构数组或Map数组值("score"), g.Slice别名{100, nil, 0})
	})
}

func TestVar_ListItemValues_Struct(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			Id    int
			Score float64
		}
		listStruct := g.Slice别名{
			T{1, 100},
			T{2, 99},
			T{3, 0},
		}
		t.Assert(泛型类.X创建(listStruct).X取结构数组或Map数组值("Id"), g.Slice别名{1, 2, 3})
		t.Assert(泛型类.X创建(listStruct).X取结构数组或Map数组值("Score"), g.Slice别名{100, 99, 0})
	})
	// Pointer items.
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			Id    int
			Score float64
		}
		listStruct := g.Slice别名{
			&T{1, 100},
			&T{2, 99},
			&T{3, 0},
		}
		t.Assert(泛型类.X创建(listStruct).X取结构数组或Map数组值("Id"), g.Slice别名{1, 2, 3})
		t.Assert(泛型类.X创建(listStruct).X取结构数组或Map数组值("Score"), g.Slice别名{100, 99, 0})
	})
	// 空元素值。
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			Id    int
			Score interface{}
		}
		listStruct := g.Slice别名{
			T{1, 100},
			T{2, nil},
			T{3, 0},
		}
		t.Assert(泛型类.X创建(listStruct).X取结构数组或Map数组值("Id"), g.Slice别名{1, 2, 3})
		t.Assert(泛型类.X创建(listStruct).X取结构数组或Map数组值("Score"), g.Slice别名{100, nil, 0})
	})
}

func TestVar_ListItemValuesUnique(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		listMap := g.Map数组{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 100},
			g.Map{"id": 3, "score": 100},
			g.Map{"id": 4, "score": 100},
			g.Map{"id": 5, "score": 100},
		}
		t.Assert(泛型类.X创建(listMap).X取结构数组或Map数组值并去重("id"), g.Slice别名{1, 2, 3, 4, 5})
		t.Assert(泛型类.X创建(listMap).X取结构数组或Map数组值并去重("score"), g.Slice别名{100})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		listMap := g.Map数组{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 100},
			g.Map{"id": 3, "score": 100},
			g.Map{"id": 4, "score": 100},
			g.Map{"id": 5, "score": 99},
		}
		t.Assert(泛型类.X创建(listMap).X取结构数组或Map数组值并去重("id"), g.Slice别名{1, 2, 3, 4, 5})
		t.Assert(泛型类.X创建(listMap).X取结构数组或Map数组值并去重("score"), g.Slice别名{100, 99})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		listMap := g.Map数组{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 100},
			g.Map{"id": 3, "score": 0},
			g.Map{"id": 4, "score": 100},
			g.Map{"id": 5, "score": 99},
		}
		t.Assert(泛型类.X创建(listMap).X取结构数组或Map数组值并去重("id"), g.Slice别名{1, 2, 3, 4, 5})
		t.Assert(泛型类.X创建(listMap).X取结构数组或Map数组值并去重("score"), g.Slice别名{100, 0, 99})
	})
}
