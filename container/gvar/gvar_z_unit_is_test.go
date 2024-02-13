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

func TestVar_IsNil(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(0).X是否为Nil(), false)
		t.Assert(g.X泛型类(nil).X是否为Nil(), true)
		t.Assert(g.X泛型类(g.Map{}).X是否为Nil(), false)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为Nil(), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(1).X是否为Nil(), false)
		t.Assert(g.X泛型类(0.1).X是否为Nil(), false)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为Nil(), false)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为Nil(), false)
	})
}

func TestVar_IsEmpty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(0).X是否为空(), true)
		t.Assert(g.X泛型类(nil).X是否为空(), true)
		t.Assert(g.X泛型类(g.Map{}).X是否为空(), true)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为空(), true)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(1).X是否为空(), false)
		t.Assert(g.X泛型类(0.1).X是否为空(), false)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为空(), false)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为空(), false)
	})
}

func TestVar_IsInt(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(0).X是否为整数(), true)
		t.Assert(g.X泛型类(nil).X是否为整数(), false)
		t.Assert(g.X泛型类(g.Map{}).X是否为整数(), false)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为整数(), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(1).X是否为整数(), true)
		t.Assert(g.X泛型类(-1).X是否为整数(), true)
		t.Assert(g.X泛型类(0.1).X是否为整数(), false)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为整数(), false)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为整数(), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(int8(1)).X是否为整数(), true)
		t.Assert(g.X泛型类(uint8(1)).X是否为整数(), false)
	})
}

func TestVar_IsUint(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(0).X是否为正整数(), false)
		t.Assert(g.X泛型类(nil).X是否为正整数(), false)
		t.Assert(g.X泛型类(g.Map{}).X是否为正整数(), false)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为正整数(), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(1).X是否为正整数(), false)
		t.Assert(g.X泛型类(-1).X是否为正整数(), false)
		t.Assert(g.X泛型类(0.1).X是否为正整数(), false)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为正整数(), false)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为正整数(), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(int8(1)).X是否为正整数(), false)
		t.Assert(g.X泛型类(uint8(1)).X是否为正整数(), true)
	})
}

func TestVar_IsFloat(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(0).X是否为小数(), false)
		t.Assert(g.X泛型类(nil).X是否为小数(), false)
		t.Assert(g.X泛型类(g.Map{}).X是否为小数(), false)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为小数(), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(1).X是否为小数(), false)
		t.Assert(g.X泛型类(-1).X是否为小数(), false)
		t.Assert(g.X泛型类(0.1).X是否为小数(), true)
		t.Assert(g.X泛型类(float64(1)).X是否为小数(), true)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为小数(), false)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为小数(), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(int8(1)).X是否为小数(), false)
		t.Assert(g.X泛型类(uint8(1)).X是否为小数(), false)
	})
}

func TestVar_IsSlice(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(0).X是否为数组(), false)
		t.Assert(g.X泛型类(nil).X是否为数组(), false)
		t.Assert(g.X泛型类(g.Map{}).X是否为数组(), false)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为数组(), true)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(1).X是否为数组(), false)
		t.Assert(g.X泛型类(-1).X是否为数组(), false)
		t.Assert(g.X泛型类(0.1).X是否为数组(), false)
		t.Assert(g.X泛型类(float64(1)).X是否为数组(), false)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为数组(), false)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为数组(), true)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(int8(1)).X是否为数组(), false)
		t.Assert(g.X泛型类(uint8(1)).X是否为数组(), false)
	})
}

func TestVar_IsMap(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(0).X是否为Map(), false)
		t.Assert(g.X泛型类(nil).X是否为Map(), false)
		t.Assert(g.X泛型类(g.Map{}).X是否为Map(), true)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为Map(), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(1).X是否为Map(), false)
		t.Assert(g.X泛型类(-1).X是否为Map(), false)
		t.Assert(g.X泛型类(0.1).X是否为Map(), false)
		t.Assert(g.X泛型类(float64(1)).X是否为Map(), false)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为Map(), true)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为Map(), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(int8(1)).X是否为Map(), false)
		t.Assert(g.X泛型类(uint8(1)).X是否为Map(), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(泛型类.X创建(泛型类.X创建("asd")).X是否为Map(), false)
		t.Assert(泛型类.X创建(&g.Map{"k": "v"}).X是否为Map(), true)
	})
}

func TestVar_IsStruct(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(0).X是否为结构(), false)
		t.Assert(g.X泛型类(nil).X是否为结构(), false)
		t.Assert(g.X泛型类(g.Map{}).X是否为结构(), false)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为结构(), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X泛型类(1).X是否为结构(), false)
		t.Assert(g.X泛型类(-1).X是否为结构(), false)
		t.Assert(g.X泛型类(0.1).X是否为结构(), false)
		t.Assert(g.X泛型类(float64(1)).X是否为结构(), false)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为结构(), false)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为结构(), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		a := &struct {
		}{}
		t.Assert(g.X泛型类(a).X是否为结构(), true)
		t.Assert(g.X泛型类(*a).X是否为结构(), true)
		t.Assert(g.X泛型类(&a).X是否为结构(), true)
	})
}
