// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 安全变量类_test

import (
	"sync"
	"testing"
	
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_Byte(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var wg sync.WaitGroup
		addTimes := 127
		i := 安全变量类.NewByte(byte(0))
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
		i1 := 安全变量类.NewByte()
		t.AssertEQ(i1.X取值(), byte(0))

		i2 := 安全变量类.NewByte(byte(64))
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
	单元测试类.C(t, func(t *单元测试类.T) {
		i := 安全变量类.NewByte(49)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.X取值())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)
	})
	// Unmarshal
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		i := 安全变量类.NewByte()
		err = json.UnmarshalUseNumber([]byte("49"), &i)
		t.AssertNil(err)
		t.Assert(i.X取值(), "49")
	})
}

func Test_Byte_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *安全变量类.Byte
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		var v *V
		err := 转换类.Struct(map[string]interface{}{
			"name": "john",
			"var":  "2",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.X取值(), "2")
	})
}
