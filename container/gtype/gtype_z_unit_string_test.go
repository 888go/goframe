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

func Test_String(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		i := gtype.NewString("abc")
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
		i1 := gtype.NewString()
		t.AssertEQ(i1.X取值(), "")
	})
}

func Test_String_JSON(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := "i love gf"
		i1 := gtype.NewString(s)
		b1, err1 := json.Marshal(i1)
		b2, err2 := json.Marshal(i1.X取值())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := gtype.NewString()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.X取值(), s)
	})
}

func Test_String_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *gtype.String
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
