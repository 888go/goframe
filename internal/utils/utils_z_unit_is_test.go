// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package utils_test

import (
	"testing"

	gvar "github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/utils"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

func TestVar_IsNil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsNil(0), false)
		t.Assert(utils.IsNil(nil), true)
		t.Assert(utils.IsNil(g.Map{}), false)
		t.Assert(utils.IsNil(g.Slice别名{}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsNil(1), false)
		t.Assert(utils.IsNil(0.1), false)
		t.Assert(utils.IsNil(g.Map{"k": "v"}), false)
		t.Assert(utils.IsNil(g.Slice别名{0}), false)
	})
}

func TestVar_IsEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsEmpty(0), true)
		t.Assert(utils.IsEmpty(nil), true)
		t.Assert(utils.IsEmpty(g.Map{}), true)
		t.Assert(utils.IsEmpty(g.Slice别名{}), true)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsEmpty(1), false)
		t.Assert(utils.IsEmpty(0.1), false)
		t.Assert(utils.IsEmpty(g.Map{"k": "v"}), false)
		t.Assert(utils.IsEmpty(g.Slice别名{0}), false)
	})
}

func TestVar_IsInt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsInt(0), true)
		t.Assert(utils.IsInt(nil), false)
		t.Assert(utils.IsInt(g.Map{}), false)
		t.Assert(utils.IsInt(g.Slice别名{}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsInt(1), true)
		t.Assert(utils.IsInt(-1), true)
		t.Assert(utils.IsInt(0.1), false)
		t.Assert(utils.IsInt(g.Map{"k": "v"}), false)
		t.Assert(utils.IsInt(g.Slice别名{0}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsInt(int8(1)), true)
		t.Assert(utils.IsInt(uint8(1)), false)
	})
}

func TestVar_IsUint(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsUint(0), false)
		t.Assert(utils.IsUint(nil), false)
		t.Assert(utils.IsUint(g.Map{}), false)
		t.Assert(utils.IsUint(g.Slice别名{}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsUint(1), false)
		t.Assert(utils.IsUint(-1), false)
		t.Assert(utils.IsUint(0.1), false)
		t.Assert(utils.IsUint(g.Map{"k": "v"}), false)
		t.Assert(utils.IsUint(g.Slice别名{0}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsUint(int8(1)), false)
		t.Assert(utils.IsUint(uint8(1)), true)
	})
}

func TestVar_IsFloat(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsFloat(0), false)
		t.Assert(utils.IsFloat(nil), false)
		t.Assert(utils.IsFloat(g.Map{}), false)
		t.Assert(utils.IsFloat(g.Slice别名{}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsFloat(1), false)
		t.Assert(utils.IsFloat(-1), false)
		t.Assert(utils.IsFloat(0.1), true)
		t.Assert(utils.IsFloat(float64(1)), true)
		t.Assert(utils.IsFloat(g.Map{"k": "v"}), false)
		t.Assert(utils.IsFloat(g.Slice别名{0}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsFloat(int8(1)), false)
		t.Assert(utils.IsFloat(uint8(1)), false)
	})
}

func TestVar_IsSlice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsSlice(0), false)
		t.Assert(utils.IsSlice(nil), false)
		t.Assert(utils.IsSlice(g.Map{}), false)
		t.Assert(utils.IsSlice(g.Slice别名{}), true)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsSlice(1), false)
		t.Assert(utils.IsSlice(-1), false)
		t.Assert(utils.IsSlice(0.1), false)
		t.Assert(utils.IsSlice(float64(1)), false)
		t.Assert(utils.IsSlice(g.Map{"k": "v"}), false)
		t.Assert(utils.IsSlice(g.Slice别名{0}), true)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsSlice(int8(1)), false)
		t.Assert(utils.IsSlice(uint8(1)), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsSlice(gvar.X创建(gtime.X创建并按当前时间()).X是否为切片()), false)
	})
}

func TestVar_IsMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsMap(0), false)
		t.Assert(utils.IsMap(nil), false)
		t.Assert(utils.IsMap(g.Map{}), true)
		t.Assert(utils.IsMap(g.Slice别名{}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsMap(1), false)
		t.Assert(utils.IsMap(-1), false)
		t.Assert(utils.IsMap(0.1), false)
		t.Assert(utils.IsMap(float64(1)), false)
		t.Assert(utils.IsMap(g.Map{"k": "v"}), true)
		t.Assert(utils.IsMap(g.Slice别名{0}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsMap(int8(1)), false)
		t.Assert(utils.IsMap(uint8(1)), false)
	})
}

func TestVar_IsStruct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsStruct(0), false)
		t.Assert(utils.IsStruct(nil), false)
		t.Assert(utils.IsStruct(g.Map{}), false)
		t.Assert(utils.IsStruct(g.Slice别名{}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.IsStruct(1), false)
		t.Assert(utils.IsStruct(-1), false)
		t.Assert(utils.IsStruct(0.1), false)
		t.Assert(utils.IsStruct(float64(1)), false)
		t.Assert(utils.IsStruct(g.Map{"k": "v"}), false)
		t.Assert(utils.IsStruct(g.Slice别名{0}), false)
	})
	gtest.C(t, func(t *gtest.T) {
		a := &struct {
		}{}
		t.Assert(utils.IsStruct(a), true)
		t.Assert(utils.IsStruct(*a), true)
		t.Assert(utils.IsStruct(&a), true)
	})
}
