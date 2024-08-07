// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package empty_test

import (
	"testing"
	"time"

	gvar "github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/empty"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
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
	gtest.C(t, func(t *gtest.T) {
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
		t.Assert(empty.IsEmpty(gconv.X取整数(tmpT1)), true)
		t.Assert(empty.IsEmpty(gconv.X取整数8位(tmpT1)), true)
		t.Assert(empty.IsEmpty(gconv.X取整数16位(tmpT1)), true)
		t.Assert(empty.IsEmpty(gconv.X取整数32位(tmpT1)), true)
		t.Assert(empty.IsEmpty(gconv.X取整数64位(tmpT1)), true)
		t.Assert(empty.IsEmpty(gconv.X取正整数64位(tmpT1)), true)
		t.Assert(empty.IsEmpty(gconv.X取正整数(tmpT1)), true)
		t.Assert(empty.IsEmpty(gconv.X取正整数16位(tmpT1)), true)
		t.Assert(empty.IsEmpty(gconv.X取正整数32位(tmpT1)), true)
		t.Assert(empty.IsEmpty(gconv.X取正整数64位(tmpT1)), true)
		t.Assert(empty.IsEmpty(gconv.X取小数32位(tmpT1)), true)
		t.Assert(empty.IsEmpty(gconv.X取小数64位(tmpT1)), true)
		t.Assert(empty.IsEmpty(false), true)
		t.Assert(empty.IsEmpty([]byte("")), true)
		t.Assert(empty.IsEmpty(""), true)
		t.Assert(empty.IsEmpty(g.Map{}), true)
		t.Assert(empty.IsEmpty(g.Slice别名{}), true)
		t.Assert(empty.IsEmpty(g.X切片{}), true)
		t.Assert(empty.IsEmpty(tmpT2), true)
		t.Assert(empty.IsEmpty(tmpT3), true)
		t.Assert(empty.IsEmpty(tmpT3), true)
		t.Assert(empty.IsEmpty(tmpT4), true)
		t.Assert(empty.IsEmpty(tmpT5), true)
		t.Assert(empty.IsEmpty(tmpT6), true)
		t.Assert(empty.IsEmpty(tmpT7), true)
		t.Assert(empty.IsEmpty(tmpT8), true)

		// false
		t.Assert(empty.IsEmpty(gconv.X取整数(tmpF1)), false)
		t.Assert(empty.IsEmpty(gconv.X取整数8位(tmpF1)), false)
		t.Assert(empty.IsEmpty(gconv.X取整数16位(tmpF1)), false)
		t.Assert(empty.IsEmpty(gconv.X取整数32位(tmpF1)), false)
		t.Assert(empty.IsEmpty(gconv.X取整数64位(tmpF1)), false)
		t.Assert(empty.IsEmpty(gconv.X取正整数(tmpF1)), false)
		t.Assert(empty.IsEmpty(gconv.X取正整数8位(tmpF1)), false)
		t.Assert(empty.IsEmpty(gconv.X取正整数16位(tmpF1)), false)
		t.Assert(empty.IsEmpty(gconv.X取正整数32位(tmpF1)), false)
		t.Assert(empty.IsEmpty(gconv.X取正整数64位(tmpF1)), false)
		t.Assert(empty.IsEmpty(gconv.X取小数32位(tmpF1)), false)
		t.Assert(empty.IsEmpty(gconv.X取小数64位(tmpF1)), false)
		t.Assert(empty.IsEmpty(true), false)
		t.Assert(empty.IsEmpty(tmpT1), false)
		t.Assert(empty.IsEmpty([]byte("1")), false)
		t.Assert(empty.IsEmpty(g.Map{"a": 1}), false)
		t.Assert(empty.IsEmpty(g.Slice别名{"1"}), false)
		t.Assert(empty.IsEmpty(g.X切片{"1"}), false)
		t.Assert(empty.IsEmpty(tmpF2), false)
		t.Assert(empty.IsEmpty(tmpF3), false)
		t.Assert(empty.IsEmpty(tmpF4), false)
		t.Assert(empty.IsEmpty(tmpF5), false)
		t.Assert(empty.IsEmpty(tmpF6), false)
	})
}

func TestIsNil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(empty.IsNil(nil), true)
	})
	gtest.C(t, func(t *gtest.T) {
		var i int
		t.Assert(empty.IsNil(i), false)
	})
	gtest.C(t, func(t *gtest.T) {
		var i *int
		t.Assert(empty.IsNil(i), true)
	})
	gtest.C(t, func(t *gtest.T) {
		var i *int
		t.Assert(empty.IsNil(&i), false)
		t.Assert(empty.IsNil(&i, true), true)
	})
}

type Issue3362St struct {
	time.Time
}

func Test_Issue3362(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type A struct {
			Issue3362 *Issue3362St `json:"issue,omitempty"`
		}
		m := gvar.X创建(
			&A{},
		).X取Map(
			gvar.MapOption{
				OmitEmpty: true,
			},
		)
		t.Assert(m, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		var i int
		t.Assert(empty.IsNil(i), false)
	})
	gtest.C(t, func(t *gtest.T) {
		var i *int
		t.Assert(empty.IsNil(i), true)
	})
	gtest.C(t, func(t *gtest.T) {
		var i *int
		t.Assert(empty.IsNil(&i), false)
		t.Assert(empty.IsNil(&i, true), true)
	})
}
