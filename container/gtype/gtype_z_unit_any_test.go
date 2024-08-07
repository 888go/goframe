// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 安全变量类_test

import (
	"testing"

	gtype "github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/internal/json"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
)

func Test_Any(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t1 := Temp{Name: "gf", Age: 18}
		t2 := Temp{Name: "gf", Age: 19}
		i := gtype.New(t1)
		iClone := i.Clone()
		t.AssertEQ(iClone.X设置值(t2), t1)
		t.AssertEQ(iClone.X取值().(Temp), t2)

		// empty param test
		i1 := gtype.New()
		t.AssertEQ(i1.X取值(), nil)

		i2 := gtype.New("gf")
		t.AssertEQ(i2.String(), "gf")
		copyVal := i2.DeepCopy()
		i2.X设置值("goframe")
		t.AssertNE(copyVal, iClone.X取值())
		i2 = nil
		copyVal = i2.DeepCopy()
		t.AssertNil(copyVal)
	})
}

func Test_Any_JSON(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := "i love gf"
		i := gtype.New(s)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.X取值())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := gtype.New()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.X取值(), s)
	})
}

func Test_Any_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *gtype.Any
	}
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]any{
			"name": "john",
			"var":  "123",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.X取值(), "123")
	})
}
