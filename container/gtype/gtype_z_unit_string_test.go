// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 安全变量类_test

import (
	"testing"
	
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_String(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		i := 安全变量类.NewString("abc")
		iClone := i.Clone()
		t.AssertEQ(iClone.X设置值("123"), "abc")
		t.AssertEQ(iClone.X取值(), "123")
		t.AssertEQ(iClone.String(), "123")
		//
		copyVal := iClone.DeepCopy()
		iClone.X设置值("124")
		t.AssertNE(copyVal, iClone.X取值())
		iClone = nil
		copyVal = iClone.DeepCopy()
		t.AssertNil(copyVal)
		// empty param test
		i1 := 安全变量类.NewString()
		t.AssertEQ(i1.X取值(), "")
	})
}

func Test_String_JSON(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := "i love gf"
		i1 := 安全变量类.NewString(s)
		b1, err1 := json.Marshal(i1)
		b2, err2 := json.Marshal(i1.X取值())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := 安全变量类.NewString()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.X取值(), s)
	})
}

func Test_String_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *安全变量类.String
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		var v *V
		err := 转换类.Struct(map[string]interface{}{
			"name": "john",
			"var":  "123",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.X取值(), "123")
	})
}
