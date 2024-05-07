// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 泛型类_test

import (
	"testing"

	"github.com/888go/goframe/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestVars(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var vs = 泛型类.Vars{
			泛型类.X创建(1),
			泛型类.X创建(2),
			泛型类.X创建(3),
		}
		t.AssertEQ(vs.X取文本切片(), []string{"1", "2", "3"})
		t.AssertEQ(vs.X取any切片(), []interface{}{1, 2, 3})
		t.AssertEQ(vs.X取小数32位切片(), []float32{1, 2, 3})
		t.AssertEQ(vs.X取小数64位切片(), []float64{1, 2, 3})
		t.AssertEQ(vs.X取整数切片(), []int{1, 2, 3})
		t.AssertEQ(vs.X取整数8位切片(), []int8{1, 2, 3})
		t.AssertEQ(vs.X取整数16位切片(), []int16{1, 2, 3})
		t.AssertEQ(vs.X取整数32位切片(), []int32{1, 2, 3})
		t.AssertEQ(vs.X取整数64位切片(), []int64{1, 2, 3})
		t.AssertEQ(vs.X取正整数切片(), []uint{1, 2, 3})
		t.AssertEQ(vs.X取正整数8位切片(), []uint8{1, 2, 3})
		t.AssertEQ(vs.X取正整数16位切片(), []uint16{1, 2, 3})
		t.AssertEQ(vs.X取正整数32位切片(), []uint32{1, 2, 3})
		t.AssertEQ(vs.X取正整数64位切片(), []uint64{1, 2, 3})
	})
}

func TestVars_Scan(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id   int
			Name string
		}
		var vs = 泛型类.Vars{
			泛型类.X创建(g.Map{"id": 1, "name": "john"}),
			泛型类.X创建(g.Map{"id": 2, "name": "smith"}),
		}
		var users []User
		err := vs.X取结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, 1)
		t.Assert(users[0].Name, "john")
		t.Assert(users[1].Id, 2)
		t.Assert(users[1].Name, "smith")
	})
}
