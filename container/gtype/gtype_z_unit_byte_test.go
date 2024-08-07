// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 安全变量类_test

import (
	"sync"
	"testing"

	gtype "github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/internal/json"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
)

func Test_Byte(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var wg sync.WaitGroup
		addTimes := 127
		i := gtype.NewByte(byte(0))
		iClone := i.Clone()
		t.AssertEQ(iClone.X设置值(byte(1)), byte(0))
		t.AssertEQ(iClone.X取值(), byte(1))
		for index := 0; index < addTimes; index++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				i.Add(1)
			}()
		}
		wg.Wait()
		t.AssertEQ(byte(addTimes), i.X取值())

		// empty param test
		i1 := gtype.NewByte()
		t.AssertEQ(i1.X取值(), byte(0))

		i2 := gtype.NewByte(byte(64))
		t.AssertEQ(i2.String(), "64")
		t.AssertEQ(i2.Cas(byte(63), byte(65)), false)
		t.AssertEQ(i2.Cas(byte(64), byte(65)), true)

		copyVal := i2.DeepCopy()
		i2.X设置值(byte(65))
		t.AssertNE(copyVal, iClone.X取值())
		i2 = nil
		copyVal = i2.DeepCopy()
		t.AssertNil(copyVal)
	})
}

func Test_Byte_JSON(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		i := gtype.NewByte(49)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.X取值())
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
		t.Assert(i.X取值(), "49")
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
		t.Assert(v.Var.X取值(), "2")
	})
}
