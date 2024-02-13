// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 循环链表类_test

import (
	"container/ring"
	"testing"
	
	"github.com/888go/goframe/container/gring"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
)

type Student struct {
	position int
	name     string
	upgrade  bool
}

func TestRing_Val(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		//定义cap 为3的ring类型数据
		r := 循环链表类.New(3, true)
		//分别给3个元素初始化赋值
		r.Put(&Student{1, "jimmy", true})
		r.Put(&Student{2, "tom", true})
		r.Put(&Student{3, "alon", false})

		//元素取值并判断和预设值是否相等
		t.Assert(r.X取值().(*Student).name, "jimmy")
		//从当前位置往后移两个元素
		r.Move(2)
		t.Assert(r.X取值().(*Student).name, "alon")
		//更新元素值
		//测试 value == nil
		r.X设置值(nil)
		t.Assert(r.X取值(), nil)
		//测试value != nil
		r.X设置值(&Student{3, "jack", true})
	})
}

func TestRing_CapLen(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		r := 循环链表类.New(10)
		t.Assert(r.Cap(), 10)
		t.Assert(r.Len(), 0)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		r := 循环链表类.New(10)
		r.Put("goframe")
		//cap长度 10
		t.Assert(r.Cap(), 10)
		//已有数据项 1
		t.Assert(r.Len(), 1)
	})
}

func TestRing_Position(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		r := 循环链表类.New(2)
		r.Put(1)
		r.Put(2)
		//往后移动1个元素
		r.Next()
		t.Assert(r.X取值(), 2)
		//往前移动1个元素
		r.Prev()
		t.Assert(r.X取值(), 1)

	})
}

func TestRing_Link(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		r := 循环链表类.New(3)
		r.Put(1)
		r.Put(2)
		r.Put(3)
		s := 循环链表类.New(2)
		s.Put("a")
		s.Put("b")

		rs := r.Link(s)
		t.Assert(rs.Move(2).X取值(), "b")
	})
}

func TestRing_Unlink(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		r := 循环链表类.New(5)
		for i := 1; i <= 5; i++ {
			r.Put(i)
		}
		t.Assert(r.X取值(), 1)
		// 1 2 3 4
		// 删除当前位置往后的2个数据，返回被删除的数据
		// 重新计算s len
		s := r.Unlink(2) // 2 3
		t.Assert(s.X取值(), 2)
		t.Assert(s.Len(), 2)
	})
}

func TestRing_Slice(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		ringLen := 5
		r := 循环链表类.New(ringLen)
		for i := 0; i < ringLen; i++ {
			r.Put(i + 1)
		}
		r.Move(2)              // 3
		array := r.SliceNext() // [3 4 5 1 2]
		t.Assert(array[0], 3)
		t.Assert(len(array), 5)

		//判断array是否等于[3 4 5 1 2]
		ra := []int{3, 4, 5, 1, 2}
		t.Assert(ra, array)

		//第3个元素设为nil
		r.X设置值(nil)
		array2 := r.SliceNext() //[4 5 1 2]
		//返回当前位置往后不为空的元素数组，长度为4
		t.Assert(array2, g.Slice别名{nil, 4, 5, 1, 2})

		array3 := r.SlicePrev() //[2 1 5 4]
		t.Assert(array3, g.Slice别名{nil, 2, 1, 5, 4})

		s := 循环链表类.New(ringLen)
		for i := 0; i < ringLen; i++ {
			s.Put(i + 1)
		}
		array4 := s.SlicePrev() // []
		t.Assert(array4, g.Slice别名{1, 5, 4, 3, 2})
	})
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题链接：https://github.com/gogf/gf/issues/1394
// 翻译成中文：
// 引用了GitHub上gogf/gf项目的一个问题：https://github.com/gogf/gf/issues/1394
func Test_Issue1394(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		// gring.
		gRing := 循环链表类.New(10)
		for i := 0; i < 10; i++ {
			gRing.Put(i)
		}
		t.Logf("the length:%d", gRing.Len())
		gRingResult := gRing.Unlink(6)
		for i := 0; i < 10; i++ {
			t.Log(gRing.X取值())
			gRing = gRing.Next()
		}
		t.Logf("the ring length:%d", gRing.Len())
		t.Logf("the result length:%d", gRingResult.Len())

		// stdring
		stdRing := ring.New(10)
		for i := 0; i < 10; i++ {
			stdRing.Value = i
			stdRing = stdRing.Next()
		}
		t.Logf("the length:%d", stdRing.Len())
		stdRingResult := stdRing.Unlink(6)
		for i := 0; i < 10; i++ {
			t.Log(stdRing.Value)
			stdRing = stdRing.Next()
		}
		t.Logf("the ring length:%d", stdRing.Len())
		t.Logf("the result length:%d", stdRingResult.Len())

		// Assertion.
		t.Assert(gRing.Len(), stdRing.Len())
		t.Assert(gRingResult.Len(), stdRingResult.Len())

		for i := 0; i < 10; i++ {
			t.Assert(stdRing.Value, gRing.X取值())
			stdRing = stdRing.Next()
			gRing = gRing.Next()
		}
	})

}

func TestRing_RLockIteratorNext(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		r := 循环链表类.New(10)
		for i := 0; i < 10; i++ {
			r.X设置值(i).Next()
		}

		iterVal := 0
		r.RLockIteratorNext(func(value interface{}) bool {
			if value.(int) == 0 {
				iterVal = value.(int)
				return false
			}
			return true
		})

		t.Assert(iterVal, 0)
	})
}

func TestRing_RLockIteratorPrev(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		r := 循环链表类.New(10)
		for i := 0; i < 10; i++ {
			r.X设置值(i).Next()
		}

		iterVal := 0
		r.RLockIteratorPrev(func(value interface{}) bool {
			if value.(int) == 0 {
				iterVal = value.(int)
				return false
			}
			return true
		})

		t.Assert(iterVal, 0)
	})
}
