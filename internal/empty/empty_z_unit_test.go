// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package empty_test

import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

type TestInt int

type TestString string

type TestPerson interface {
	Say() string
}

type TestWoman struct {
}

func (woman TestWoman) Say() string {
	return "nice"
}

func TestIsEmpty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		tmpT1 := "0"
		tmpT2 := func() {}
		tmpT2 = nil
		tmpT3 := make(chan int)
		var (
			tmpT4 TestPerson  = nil
			tmpT5 *TestPerson = nil
			tmpT6 TestPerson  = TestWoman{}
			tmpT7 TestInt     = 0
			tmpT8 TestString  = ""
		)
		tmpF1 := "1"
		tmpF2 := func(a string) string { return "1" }
		tmpF3 := make(chan int, 1)
		tmpF3 <- 1
		var (
			tmpF4 TestPerson = &TestWoman{}
			tmpF5 TestInt    = 1
			tmpF6 TestString = "1"
		)

		// true
		t.Assert(empty.IsEmpty(nil), true)
		t.Assert(empty.IsEmpty(0), true)
		t.Assert(empty.IsEmpty(转换类.X取整数(tmpT1)), true)
		t.Assert(empty.IsEmpty(转换类.X取整数8位(tmpT1)), true)
		t.Assert(empty.IsEmpty(转换类.X取整数16位(tmpT1)), true)
		t.Assert(empty.IsEmpty(转换类.X取整数32位(tmpT1)), true)
		t.Assert(empty.IsEmpty(转换类.X取整数64位(tmpT1)), true)
		t.Assert(empty.IsEmpty(转换类.X取正整数64位(tmpT1)), true)
		t.Assert(empty.IsEmpty(转换类.X取正整数(tmpT1)), true)
		t.Assert(empty.IsEmpty(转换类.X取正整数16位(tmpT1)), true)
		t.Assert(empty.IsEmpty(转换类.X取正整数32位(tmpT1)), true)
		t.Assert(empty.IsEmpty(转换类.X取正整数64位(tmpT1)), true)
		t.Assert(empty.IsEmpty(转换类.X取小数32位(tmpT1)), true)
		t.Assert(empty.IsEmpty(转换类.X取小数64位(tmpT1)), true)
		t.Assert(empty.IsEmpty(false), true)
		t.Assert(empty.IsEmpty([]byte("")), true)
		t.Assert(empty.IsEmpty(""), true)
		t.Assert(empty.IsEmpty(g.Map{}), true)
		t.Assert(empty.IsEmpty(g.Slice别名{}), true)
		t.Assert(empty.IsEmpty(g.X数组{}), true)
		t.Assert(empty.IsEmpty(tmpT2), true)
		t.Assert(empty.IsEmpty(tmpT3), true)
		t.Assert(empty.IsEmpty(tmpT3), true)
		t.Assert(empty.IsEmpty(tmpT4), true)
		t.Assert(empty.IsEmpty(tmpT5), true)
		t.Assert(empty.IsEmpty(tmpT6), true)
		t.Assert(empty.IsEmpty(tmpT7), true)
		t.Assert(empty.IsEmpty(tmpT8), true)

		// false
		t.Assert(empty.IsEmpty(转换类.X取整数(tmpF1)), false)
		t.Assert(empty.IsEmpty(转换类.X取整数8位(tmpF1)), false)
		t.Assert(empty.IsEmpty(转换类.X取整数16位(tmpF1)), false)
		t.Assert(empty.IsEmpty(转换类.X取整数32位(tmpF1)), false)
		t.Assert(empty.IsEmpty(转换类.X取整数64位(tmpF1)), false)
		t.Assert(empty.IsEmpty(转换类.X取正整数(tmpF1)), false)
		t.Assert(empty.IsEmpty(转换类.X取正整数8位(tmpF1)), false)
		t.Assert(empty.IsEmpty(转换类.X取正整数16位(tmpF1)), false)
		t.Assert(empty.IsEmpty(转换类.X取正整数32位(tmpF1)), false)
		t.Assert(empty.IsEmpty(转换类.X取正整数64位(tmpF1)), false)
		t.Assert(empty.IsEmpty(转换类.X取小数32位(tmpF1)), false)
		t.Assert(empty.IsEmpty(转换类.X取小数64位(tmpF1)), false)
		t.Assert(empty.IsEmpty(true), false)
		t.Assert(empty.IsEmpty(tmpT1), false)
		t.Assert(empty.IsEmpty([]byte("1")), false)
		t.Assert(empty.IsEmpty(g.Map{"a": 1}), false)
		t.Assert(empty.IsEmpty(g.Slice别名{"1"}), false)
		t.Assert(empty.IsEmpty(g.X数组{"1"}), false)
		t.Assert(empty.IsEmpty(tmpF2), false)
		t.Assert(empty.IsEmpty(tmpF3), false)
		t.Assert(empty.IsEmpty(tmpF4), false)
		t.Assert(empty.IsEmpty(tmpF5), false)
		t.Assert(empty.IsEmpty(tmpF6), false)
	})
}

func TestIsNil(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(empty.X是否为Nil(nil), true)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var i int
		t.Assert(empty.X是否为Nil(i), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var i *int
		t.Assert(empty.X是否为Nil(i), true)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var i *int
		t.Assert(empty.X是否为Nil(&i), false)
		t.Assert(empty.X是否为Nil(&i, true), true)
	})
}
