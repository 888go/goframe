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

func Test_Bool(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		i := gtype.NewBool(true)
		iClone := i.Clone()
		t.AssertEQ(iClone.Set(false), true)
		t.AssertEQ(iClone.Val(), false)

		i1 := gtype.NewBool(false)
		iClone1 := i1.Clone()
		t.AssertEQ(iClone1.Set(true), false)
		t.AssertEQ(iClone1.Val(), true)

		t.AssertEQ(iClone1.Cas(false, true), false)
		t.AssertEQ(iClone1.String(), "true")
		t.AssertEQ(iClone1.Cas(true, false), true)
		t.AssertEQ(iClone1.String(), "false")

		copyVal := i1.DeepCopy()
		iClone.Set(true)
		t.AssertNE(copyVal, iClone.Val())
		iClone = nil
		copyVal = iClone.DeepCopy()
		t.AssertNil(copyVal)

		// empty param test
		i2 := gtype.NewBool()
		t.AssertEQ(i2.Val(), false)
	})
}

func Test_Bool_JSON(t *testing.T) {
	// Marshal
	gtest.C(t, func(t *gtest.T) {
		i := gtype.NewBool(true)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)
	})
	gtest.C(t, func(t *gtest.T) {
		i := gtype.NewBool(false)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)
	})
	// Unmarshal
	gtest.C(t, func(t *gtest.T) {
		var err error
		i := gtype.NewBool()
		err = json.UnmarshalUseNumber([]byte("true"), &i)
		t.AssertNil(err)
		t.Assert(i.Val(), true)
		err = json.UnmarshalUseNumber([]byte("false"), &i)
		t.AssertNil(err)
		t.Assert(i.Val(), false)
		err = json.UnmarshalUseNumber([]byte("1"), &i)
		t.AssertNil(err)
		t.Assert(i.Val(), true)
		err = json.UnmarshalUseNumber([]byte("0"), &i)
		t.AssertNil(err)
		t.Assert(i.Val(), false)
	})

	gtest.C(t, func(t *gtest.T) {
		i := gtype.NewBool(true)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := gtype.NewBool()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.Val(), i.Val())
	})
	gtest.C(t, func(t *gtest.T) {
		i := gtype.NewBool(false)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := gtype.NewBool()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.Val(), i.Val())
	})
}

func Test_Bool_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *gtype.Bool
	}
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "true",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), true)
	})
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "false",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), false)
	})
}
