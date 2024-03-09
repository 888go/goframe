// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 安全变量类_test

import (
	"math"
	"testing"
	
	"github.com/888go/goframe/gtype"
	"github.com/888go/goframe/gtype/internal/json"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

func Test_Float32(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		i := 安全变量类.NewFloat32(0)
		iClone := i.Clone()
		t.AssertEQ(iClone.X设置值(0.1), float32(0))
		t.AssertEQ(iClone.X取值(), float32(0.1))

		// empty param test
		i1 := 安全变量类.NewFloat32()
		t.AssertEQ(i1.X取值(), float32(0))

		i2 := 安全变量类.NewFloat32(1.23)
		t.AssertEQ(i2.Add(3.21), float32(4.44))
		t.AssertEQ(i2.Cas(4.45, 5.55), false)
		t.AssertEQ(i2.Cas(4.44, 5.55), true)
		t.AssertEQ(i2.String(), "5.55")

		copyVal := i2.DeepCopy()
		i2.X设置值(float32(6.66))
		t.AssertNE(copyVal, iClone.X取值())
		i2 = nil
		copyVal = i2.DeepCopy()
		t.AssertNil(copyVal)
	})
}

func Test_Float32_JSON(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := float32(math.MaxFloat32)
		i := 安全变量类.NewFloat32(v)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.X取值())

		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := 安全变量类.NewFloat32()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.X取值(), v)
	})
}

func Test_Float32_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *安全变量类.Float32
	}
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "123.456",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.X取值(), "123.456")
	})
}
