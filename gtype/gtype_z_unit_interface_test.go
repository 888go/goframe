// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 安全变量类_test

import (
	"testing"
	
	"github.com/888go/goframe/gtype"
	"github.com/888go/goframe/gtype/internal/json"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

func Test_Interface(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t1 := Temp{Name: "gf", Age: 18}
		t2 := Temp{Name: "gf", Age: 19}
		i := 安全变量类.New(t1)
		iClone := i.Clone()
		t.AssertEQ(iClone.X设置值(t2), t1)
		t.AssertEQ(iClone.X取值().(Temp), t2)

		// empty param test
		i1 := 安全变量类.New()
		t.AssertEQ(i1.X取值(), nil)

		i2 := 安全变量类.New("gf")
		t.AssertEQ(i2.String(), "gf")
		copyVal := i2.DeepCopy()
		i2.X设置值("goframe")
		t.AssertNE(copyVal, iClone.X取值())
		i2 = nil
		copyVal = i2.DeepCopy()
		t.AssertNil(copyVal)
	})
}

func Test_Interface_JSON(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := "i love gf"
		i := 安全变量类.New(s)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.X取值())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := 安全变量类.New()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.X取值(), s)
	})
}

func Test_Interface_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *安全变量类.Interface
	}
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "123",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.X取值(), "123")
	})
}
