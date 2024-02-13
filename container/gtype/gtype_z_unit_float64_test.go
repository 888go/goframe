// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 安全变量类_test

import (
	"math"
	"testing"
	
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_Float64(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		i := 安全变量类.NewFloat64(0)
		iClone := i.Clone()
		t.AssertEQ(iClone.X设置值(0.1), float64(0))
		t.AssertEQ(iClone.X取值(), float64(0.1))
		// empty param test
		i1 := 安全变量类.NewFloat64()
		t.AssertEQ(i1.X取值(), float64(0))

		i2 := 安全变量类.NewFloat64(1.1)
		t.AssertEQ(i2.Add(3.3), 4.4)
		t.AssertEQ(i2.Cas(4.5, 5.5), false)
		t.AssertEQ(i2.Cas(4.4, 5.5), true)
		t.AssertEQ(i2.String(), "5.5")

		copyVal := i2.DeepCopy()
		i2.X设置值(6.6)
		t.AssertNE(copyVal, iClone.X取值())
		i2 = nil
		copyVal = i2.DeepCopy()
		t.AssertNil(copyVal)
	})
}

func Test_Float64_JSON(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		v := math.MaxFloat64
		i := 安全变量类.NewFloat64(v)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.X取值())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := 安全变量类.NewFloat64()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.X取值(), v)
	})
}

func Test_Float64_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *安全变量类.Float64
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		var v *V
		err := 转换类.Struct(map[string]interface{}{
			"name": "john",
			"var":  "123.456",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.X取值(), "123.456")
	})
}
