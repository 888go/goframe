// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtype_test

import (
	"sync"
	"testing"
	
	"github.com/888go/goframe/gtype"
	"github.com/888go/goframe/gtype/internal/json"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

func Test_Byte(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var wg sync.WaitGroup
		addTimes := 127
		i := gtype.NewByte(byte(0))
		iClone := i.Clone()
		t.AssertEQ(iClone.Set(byte(1)), byte(0))
		t.AssertEQ(iClone.Val(), byte(1))
		for index := 0; index < addTimes; index++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				i.Add(1)
			}()
		}
		wg.Wait()
		t.AssertEQ(byte(addTimes), i.Val())

		// empty param test
		i1 := gtype.NewByte()
		t.AssertEQ(i1.Val(), byte(0))

		i2 := gtype.NewByte(byte(64))
		t.AssertEQ(i2.String(), "64")
		t.AssertEQ(i2.Cas(byte(63), byte(65)), false)
		t.AssertEQ(i2.Cas(byte(64), byte(65)), true)

		copyVal := i2.DeepCopy()
		i2.Set(byte(65))
		t.AssertNE(copyVal, iClone.Val())
		i2 = nil
		copyVal = i2.DeepCopy()
		t.AssertNil(copyVal)
	})
}

func Test_Byte_JSON(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		i := gtype.NewByte(49)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)
	})
	// Unmarshal
	gtest.C(t, func(t *gtest.T) {
		var err error
		i := gtype.NewByte()
		err = json.UnmarshalUseNumber([]byte("49"), &i)
		t.AssertNil(err)
		t.Assert(i.Val(), "49")
	})
}

func Test_Byte_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *gtype.Byte
	}
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "2",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), "2")
	})
}
