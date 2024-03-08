// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gvar_test

import (
	"testing"
	
	"github.com/888go/goframe/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestVars(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var vs = gvar.Vars{
			gvar.New(1),
			gvar.New(2),
			gvar.New(3),
		}
		t.AssertEQ(vs.Strings(), []string{"1", "2", "3"})
		t.AssertEQ(vs.Interfaces(), []interface{}{1, 2, 3})
		t.AssertEQ(vs.Float32s(), []float32{1, 2, 3})
		t.AssertEQ(vs.Float64s(), []float64{1, 2, 3})
		t.AssertEQ(vs.Ints(), []int{1, 2, 3})
		t.AssertEQ(vs.Int8s(), []int8{1, 2, 3})
		t.AssertEQ(vs.Int16s(), []int16{1, 2, 3})
		t.AssertEQ(vs.Int32s(), []int32{1, 2, 3})
		t.AssertEQ(vs.Int64s(), []int64{1, 2, 3})
		t.AssertEQ(vs.Uints(), []uint{1, 2, 3})
		t.AssertEQ(vs.Uint8s(), []uint8{1, 2, 3})
		t.AssertEQ(vs.Uint16s(), []uint16{1, 2, 3})
		t.AssertEQ(vs.Uint32s(), []uint32{1, 2, 3})
		t.AssertEQ(vs.Uint64s(), []uint64{1, 2, 3})
	})
}

func TestVars_Scan(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id   int
			Name string
		}
		var vs = gvar.Vars{
			gvar.New(g.Map{"id": 1, "name": "john"}),
			gvar.New(g.Map{"id": 2, "name": "smith"}),
		}
		var users []User
		err := vs.Scan(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, 1)
		t.Assert(users[0].Name, "john")
		t.Assert(users[1].Id, 2)
		t.Assert(users[1].Name, "smith")
	})
}