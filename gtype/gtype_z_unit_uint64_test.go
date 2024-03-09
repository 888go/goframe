// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 安全变量类_test

import (
	"math"
	"sync"
	"testing"
	
	"github.com/888go/goframe/gtype"
	"github.com/888go/goframe/gtype/internal/json"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

type Temp struct {
	Name string
	Age  int
}

func Test_Uint64(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var wg sync.WaitGroup
		addTimes := 1000
		i := 安全变量类.NewUint64(0)
		iClone := i.Clone()
		t.AssertEQ(iClone.X设置值(1), uint64(0))
		t.AssertEQ(iClone.X取值(), uint64(1))
		for index := 0; index < addTimes; index++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				i.Add(1)
			}()
		}
		wg.Wait()
		t.AssertEQ(uint64(addTimes), i.X取值())

		// empty param test
		i1 := 安全变量类.NewUint64()
		t.AssertEQ(i1.X取值(), uint64(0))

		i2 := 安全变量类.NewUint64(11)
		t.AssertEQ(i2.Add(1), uint64(12))
		t.AssertEQ(i2.Cas(11, 13), false)
		t.AssertEQ(i2.Cas(12, 13), true)
		t.AssertEQ(i2.String(), "13")

		copyVal := i2.DeepCopy()
		i2.X设置值(14)
		t.AssertNE(copyVal, iClone.X取值())
		i2 = nil
		copyVal = i2.DeepCopy()
		t.AssertNil(copyVal)
	})
}

func Test_Uint64_JSON(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		i := 安全变量类.NewUint64(math.MaxUint64)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.X取值())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := 安全变量类.NewUint64()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.X取值(), i)
	})
}

func Test_Uint64_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *安全变量类.Uint64
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
