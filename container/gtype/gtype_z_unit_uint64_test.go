// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 安全变量类_test

import (
	"math"
	"sync"
	"testing"

	gtype "github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/internal/json"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
)

type Temp struct {
	Name string
	Age  int
}

func Test_Uint64(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var wg sync.WaitGroup
		addTimes := 1000
		i := gtype.NewUint64(0)
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
		i1 := gtype.NewUint64()
		t.AssertEQ(i1.X取值(), uint64(0))

		i2 := gtype.NewUint64(11)
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
		i := gtype.NewUint64(math.MaxUint64)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.X取值())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := gtype.NewUint64()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.X取值(), i)
	})
}

func Test_Uint64_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *gtype.Uint64
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
