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

func TestVar_IsNil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(0).X是否为Nil(), false)
		t.Assert(g.X泛型类(nil).X是否为Nil(), true)
		t.Assert(g.X泛型类(g.Map{}).X是否为Nil(), false)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为Nil(), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(1).X是否为Nil(), false)
		t.Assert(g.X泛型类(0.1).X是否为Nil(), false)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为Nil(), false)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为Nil(), false)
	})
}

func TestVar_IsEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(0).X是否为空(), true)
		t.Assert(g.X泛型类(nil).X是否为空(), true)
		t.Assert(g.X泛型类(g.Map{}).X是否为空(), true)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为空(), true)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(1).X是否为空(), false)
		t.Assert(g.X泛型类(0.1).X是否为空(), false)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为空(), false)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为空(), false)
	})
}

func TestVar_IsInt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(0).X是否为整数(), true)
		t.Assert(g.X泛型类(nil).X是否为整数(), false)
		t.Assert(g.X泛型类(g.Map{}).X是否为整数(), false)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为整数(), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(1).X是否为整数(), true)
		t.Assert(g.X泛型类(-1).X是否为整数(), true)
		t.Assert(g.X泛型类(0.1).X是否为整数(), false)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为整数(), false)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为整数(), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(int8(1)).X是否为整数(), true)
		t.Assert(g.X泛型类(uint8(1)).X是否为整数(), false)
	})
}

func TestVar_IsUint(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(0).X是否为正整数(), false)
		t.Assert(g.X泛型类(nil).X是否为正整数(), false)
		t.Assert(g.X泛型类(g.Map{}).X是否为正整数(), false)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为正整数(), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(1).X是否为正整数(), false)
		t.Assert(g.X泛型类(-1).X是否为正整数(), false)
		t.Assert(g.X泛型类(0.1).X是否为正整数(), false)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为正整数(), false)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为正整数(), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(int8(1)).X是否为正整数(), false)
		t.Assert(g.X泛型类(uint8(1)).X是否为正整数(), true)
	})
}

func TestVar_IsFloat(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(0).X是否为小数(), false)
		t.Assert(g.X泛型类(nil).X是否为小数(), false)
		t.Assert(g.X泛型类(g.Map{}).X是否为小数(), false)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为小数(), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(1).X是否为小数(), false)
		t.Assert(g.X泛型类(-1).X是否为小数(), false)
		t.Assert(g.X泛型类(0.1).X是否为小数(), true)
		t.Assert(g.X泛型类(float64(1)).X是否为小数(), true)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为小数(), false)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为小数(), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(int8(1)).X是否为小数(), false)
		t.Assert(g.X泛型类(uint8(1)).X是否为小数(), false)
	})
}

func TestVar_IsSlice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(0).X是否为切片(), false)
		t.Assert(g.X泛型类(nil).X是否为切片(), false)
		t.Assert(g.X泛型类(g.Map{}).X是否为切片(), false)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为切片(), true)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(1).X是否为切片(), false)
		t.Assert(g.X泛型类(-1).X是否为切片(), false)
		t.Assert(g.X泛型类(0.1).X是否为切片(), false)
		t.Assert(g.X泛型类(float64(1)).X是否为切片(), false)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为切片(), false)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为切片(), true)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(int8(1)).X是否为切片(), false)
		t.Assert(g.X泛型类(uint8(1)).X是否为切片(), false)
	})
}

func TestVar_IsMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(0).X是否为Map(), false)
		t.Assert(g.X泛型类(nil).X是否为Map(), false)
		t.Assert(g.X泛型类(g.Map{}).X是否为Map(), true)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为Map(), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(1).X是否为Map(), false)
		t.Assert(g.X泛型类(-1).X是否为Map(), false)
		t.Assert(g.X泛型类(0.1).X是否为Map(), false)
		t.Assert(g.X泛型类(float64(1)).X是否为Map(), false)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为Map(), true)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为Map(), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(int8(1)).X是否为Map(), false)
		t.Assert(g.X泛型类(uint8(1)).X是否为Map(), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gvar.X创建(gvar.X创建("asd")).X是否为Map(), false)
		t.Assert(gvar.X创建(&g.Map{"k": "v"}).X是否为Map(), true)
	})
}

func TestVar_IsStruct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(0).X是否为结构(), false)
		t.Assert(g.X泛型类(nil).X是否为结构(), false)
		t.Assert(g.X泛型类(g.Map{}).X是否为结构(), false)
		t.Assert(g.X泛型类(g.Slice别名{}).X是否为结构(), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.X泛型类(1).X是否为结构(), false)
		t.Assert(g.X泛型类(-1).X是否为结构(), false)
		t.Assert(g.X泛型类(0.1).X是否为结构(), false)
		t.Assert(g.X泛型类(float64(1)).X是否为结构(), false)
		t.Assert(g.X泛型类(g.Map{"k": "v"}).X是否为结构(), false)
		t.Assert(g.X泛型类(g.Slice别名{0}).X是否为结构(), false)
	})
	gtest.C(t, func(t *gtest.T) {
		a := &struct {
		}{}
		t.Assert(g.X泛型类(a).X是否为结构(), true)
		t.Assert(g.X泛型类(*a).X是否为结构(), true)
		t.Assert(g.X泛型类(&a).X是否为结构(), true)
	})
}
