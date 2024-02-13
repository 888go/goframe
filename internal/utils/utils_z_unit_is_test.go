// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package utils_test

import (
	"testing"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
)

func TestVar_IsNil(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.X是否为Nil(0), false)
		t.Assert(utils.X是否为Nil(nil), true)
		t.Assert(utils.X是否为Nil(g.Map{}), false)
		t.Assert(utils.X是否为Nil(g.Slice别名{}), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.X是否为Nil(1), false)
		t.Assert(utils.X是否为Nil(0.1), false)
		t.Assert(utils.X是否为Nil(g.Map{"k": "v"}), false)
		t.Assert(utils.X是否为Nil(g.Slice别名{0}), false)
	})
}

func TestVar_IsEmpty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsEmpty(0), true)
		t.Assert(utils.IsEmpty(nil), true)
		t.Assert(utils.IsEmpty(g.Map{}), true)
		t.Assert(utils.IsEmpty(g.Slice别名{}), true)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsEmpty(1), false)
		t.Assert(utils.IsEmpty(0.1), false)
		t.Assert(utils.IsEmpty(g.Map{"k": "v"}), false)
		t.Assert(utils.IsEmpty(g.Slice别名{0}), false)
	})
}

func TestVar_IsInt(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsInt(0), true)
		t.Assert(utils.IsInt(nil), false)
		t.Assert(utils.IsInt(g.Map{}), false)
		t.Assert(utils.IsInt(g.Slice别名{}), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsInt(1), true)
		t.Assert(utils.IsInt(-1), true)
		t.Assert(utils.IsInt(0.1), false)
		t.Assert(utils.IsInt(g.Map{"k": "v"}), false)
		t.Assert(utils.IsInt(g.Slice别名{0}), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsInt(int8(1)), true)
		t.Assert(utils.IsInt(uint8(1)), false)
	})
}

func TestVar_IsUint(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsUint(0), false)
		t.Assert(utils.IsUint(nil), false)
		t.Assert(utils.IsUint(g.Map{}), false)
		t.Assert(utils.IsUint(g.Slice别名{}), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsUint(1), false)
		t.Assert(utils.IsUint(-1), false)
		t.Assert(utils.IsUint(0.1), false)
		t.Assert(utils.IsUint(g.Map{"k": "v"}), false)
		t.Assert(utils.IsUint(g.Slice别名{0}), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsUint(int8(1)), false)
		t.Assert(utils.IsUint(uint8(1)), true)
	})
}

func TestVar_IsFloat(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsFloat(0), false)
		t.Assert(utils.IsFloat(nil), false)
		t.Assert(utils.IsFloat(g.Map{}), false)
		t.Assert(utils.IsFloat(g.Slice别名{}), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsFloat(1), false)
		t.Assert(utils.IsFloat(-1), false)
		t.Assert(utils.IsFloat(0.1), true)
		t.Assert(utils.IsFloat(float64(1)), true)
		t.Assert(utils.IsFloat(g.Map{"k": "v"}), false)
		t.Assert(utils.IsFloat(g.Slice别名{0}), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsFloat(int8(1)), false)
		t.Assert(utils.IsFloat(uint8(1)), false)
	})
}

func TestVar_IsSlice(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsSlice(0), false)
		t.Assert(utils.IsSlice(nil), false)
		t.Assert(utils.IsSlice(g.Map{}), false)
		t.Assert(utils.IsSlice(g.Slice别名{}), true)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsSlice(1), false)
		t.Assert(utils.IsSlice(-1), false)
		t.Assert(utils.IsSlice(0.1), false)
		t.Assert(utils.IsSlice(float64(1)), false)
		t.Assert(utils.IsSlice(g.Map{"k": "v"}), false)
		t.Assert(utils.IsSlice(g.Slice别名{0}), true)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsSlice(int8(1)), false)
		t.Assert(utils.IsSlice(uint8(1)), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsSlice(泛型类.X创建(时间类.X创建并按当前时间()).X是否为数组()), false)
	})
}

func TestVar_IsMap(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsMap(0), false)
		t.Assert(utils.IsMap(nil), false)
		t.Assert(utils.IsMap(g.Map{}), true)
		t.Assert(utils.IsMap(g.Slice别名{}), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsMap(1), false)
		t.Assert(utils.IsMap(-1), false)
		t.Assert(utils.IsMap(0.1), false)
		t.Assert(utils.IsMap(float64(1)), false)
		t.Assert(utils.IsMap(g.Map{"k": "v"}), true)
		t.Assert(utils.IsMap(g.Slice别名{0}), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsMap(int8(1)), false)
		t.Assert(utils.IsMap(uint8(1)), false)
	})
}

func TestVar_IsStruct(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsStruct(0), false)
		t.Assert(utils.IsStruct(nil), false)
		t.Assert(utils.IsStruct(g.Map{}), false)
		t.Assert(utils.IsStruct(g.Slice别名{}), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.IsStruct(1), false)
		t.Assert(utils.IsStruct(-1), false)
		t.Assert(utils.IsStruct(0.1), false)
		t.Assert(utils.IsStruct(float64(1)), false)
		t.Assert(utils.IsStruct(g.Map{"k": "v"}), false)
		t.Assert(utils.IsStruct(g.Slice别名{0}), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		a := &struct {
		}{}
		t.Assert(utils.IsStruct(a), true)
		t.Assert(utils.IsStruct(*a), true)
		t.Assert(utils.IsStruct(&a), true)
	})
}
