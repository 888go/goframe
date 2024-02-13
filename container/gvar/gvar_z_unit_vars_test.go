// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 泛型类_test

import (
	"testing"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
)

func TestVars(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var vs = 泛型类.Vars{
			泛型类.X创建(1),
			泛型类.X创建(2),
			泛型类.X创建(3),
		}
		t.AssertEQ(vs.X取文本数组(), []string{"1", "2", "3"})
		t.AssertEQ(vs.X取any数组(), []interface{}{1, 2, 3})
		t.AssertEQ(vs.X取小数32位数组(), []float32{1, 2, 3})
		t.AssertEQ(vs.X取小数64位数组(), []float64{1, 2, 3})
		t.AssertEQ(vs.X取整数数组(), []int{1, 2, 3})
		t.AssertEQ(vs.X取整数8位数组(), []int8{1, 2, 3})
		t.AssertEQ(vs.X取整数16位数组(), []int16{1, 2, 3})
		t.AssertEQ(vs.X取整数32位数组(), []int32{1, 2, 3})
		t.AssertEQ(vs.X取整数64位数组(), []int64{1, 2, 3})
		t.AssertEQ(vs.X取正整数数组(), []uint{1, 2, 3})
		t.AssertEQ(vs.X取正整数8位数组(), []uint8{1, 2, 3})
		t.AssertEQ(vs.X取正整数16位数组(), []uint16{1, 2, 3})
		t.AssertEQ(vs.X取正整数32位数组(), []uint32{1, 2, 3})
		t.AssertEQ(vs.X取正整数64位数组(), []uint64{1, 2, 3})
	})
}

func TestVars_Scan(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
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
