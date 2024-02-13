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

func Test_Bytes(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		i := 安全变量类.NewBytes([]byte("abc"))
		iClone := i.Clone()
		t.AssertEQ(iClone.X设置值([]byte("123")), []byte("abc"))
		t.AssertEQ(iClone.X取值(), []byte("123"))

		// empty param test
		i1 := 安全变量类.NewBytes()
		t.AssertEQ(i1.X取值(), nil)

		i2 := 安全变量类.NewBytes([]byte("abc"))
		t.Assert(i2.String(), "abc")

		copyVal := i2.DeepCopy()
		i2.X设置值([]byte("def"))
		t.AssertNE(copyVal, iClone.X取值())
		i2 = nil
		copyVal = i2.DeepCopy()
		t.AssertNil(copyVal)
	})
}

func Test_Bytes_JSON(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		b := []byte("i love gf")
		i := 安全变量类.NewBytes(b)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.X取值())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := 安全变量类.NewBytes()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.X取值(), b)
	})
}

func Test_Bytes_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *安全变量类.Bytes
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
