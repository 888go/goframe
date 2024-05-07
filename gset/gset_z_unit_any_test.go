// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// go test *.go

package 集合类_test

import (
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/888go/goframe/gset"
	"github.com/888go/goframe/gset/internal/json"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

func TestSet_Var(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var s 集合类.Set
		s.X加入(1, 1, 2)
		s.X加入([]interface{}{3, 4}...)
		t.Assert(s.X取数量(), 4)
		t.AssertIN(1, s.X取集合切片())
		t.AssertIN(2, s.X取集合切片())
		t.AssertIN(3, s.X取集合切片())
		t.AssertIN(4, s.X取集合切片())
		t.AssertNI(0, s.X取集合切片())
		t.Assert(s.X是否存在(4), true)
		t.Assert(s.X是否存在(5), false)
		s.X删除(1)
		t.Assert(s.X取数量(), 3)
		s.X清空()
		t.Assert(s.X取数量(), 0)
	})
}

func TestSet_New(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建()
		s.X加入(1, 1, 2)
		s.X加入([]interface{}{3, 4}...)
		t.Assert(s.X取数量(), 4)
		t.AssertIN(1, s.X取集合切片())
		t.AssertIN(2, s.X取集合切片())
		t.AssertIN(3, s.X取集合切片())
		t.AssertIN(4, s.X取集合切片())
		t.AssertNI(0, s.X取集合切片())
		t.Assert(s.X是否存在(4), true)
		t.Assert(s.X是否存在(5), false)
		s.X删除(1)
		t.Assert(s.X取数量(), 3)
		s.X清空()
		t.Assert(s.X取数量(), 0)
	})
}

func TestSet_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.NewSet别名()
		s.X加入(1, 1, 2)
		s.X加入([]interface{}{3, 4}...)
		t.Assert(s.X取数量(), 4)
		t.AssertIN(1, s.X取集合切片())
		t.AssertIN(2, s.X取集合切片())
		t.AssertIN(3, s.X取集合切片())
		t.AssertIN(4, s.X取集合切片())
		t.AssertNI(0, s.X取集合切片())
		t.Assert(s.X是否存在(4), true)
		t.Assert(s.X是否存在(5), false)
		s.X删除(1)
		t.Assert(s.X取数量(), 3)
		s.X清空()
		t.Assert(s.X取数量(), 0)
	})
}

func TestSet_Iterator(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.NewSet别名()
		s.X加入(1, 2, 3)
		t.Assert(s.X取数量(), 3)

		a1 := garray.New(true)
		a2 := garray.New(true)
		s.X遍历(func(v interface{}) bool {
			a1.Append(1)
			return false
		})
		s.X遍历(func(v interface{}) bool {
			a2.Append(1)
			return true
		})
		t.Assert(a1.Len(), 1)
		t.Assert(a2.Len(), 3)
	})
}

func TestSet_LockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.NewSet别名()
		s.X加入(1, 2, 3)
		t.Assert(s.X取数量(), 3)
		s.X写锁定_函数(func(m map[interface{}]struct{}) {
			delete(m, 1)
		})
		t.Assert(s.X取数量(), 2)
		s.X读锁定_函数(func(m map[interface{}]struct{}) {
			t.Assert(m, map[interface{}]struct{}{
				3: {},
				2: {},
			})
		})
	})
}

func TestSet_Equal(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.NewSet别名()
		s2 := 集合类.NewSet别名()
		s3 := 集合类.NewSet别名()
		s4 := 集合类.NewSet别名()
		s1.X加入(1, 2, 3)
		s2.X加入(1, 2, 3)
		s3.X加入(1, 2, 3, 4)
		s4.X加入(4, 5, 6)
		t.Assert(s1.X是否相等(s2), true)
		t.Assert(s1.X是否相等(s3), false)
		t.Assert(s1.X是否相等(s4), false)
		s5 := s1
		t.Assert(s1.X是否相等(s5), true)
	})
}

func TestSet_IsSubsetOf(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.NewSet别名()
		s2 := 集合类.NewSet别名()
		s3 := 集合类.NewSet别名()
		s1.X加入(1, 2)
		s2.X加入(1, 2, 3)
		s3.X加入(1, 2, 3, 4)
		t.Assert(s1.X是否为子集(s2), true)
		t.Assert(s2.X是否为子集(s3), true)
		t.Assert(s1.X是否为子集(s3), true)
		t.Assert(s2.X是否为子集(s1), false)
		t.Assert(s3.X是否为子集(s2), false)

		s4 := s1
		t.Assert(s1.X是否为子集(s4), true)
	})
}

func TestSet_Union(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.NewSet别名()
		s2 := 集合类.NewSet别名()
		s1.X加入(1, 2)
		s2.X加入(3, 4)
		s3 := s1.X取并集(s2)
		t.Assert(s3.X是否存在(1), true)
		t.Assert(s3.X是否存在(2), true)
		t.Assert(s3.X是否存在(3), true)
		t.Assert(s3.X是否存在(4), true)
	})
}

func TestSet_Diff(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.NewSet别名()
		s2 := 集合类.NewSet别名()
		s1.X加入(1, 2, 3)
		s2.X加入(3, 4, 5)
		s3 := s1.X取差集(s2)
		t.Assert(s3.X是否存在(1), true)
		t.Assert(s3.X是否存在(2), true)
		t.Assert(s3.X是否存在(3), false)
		t.Assert(s3.X是否存在(4), false)

		s4 := s1
		s5 := s1.X取差集(s2, s4)
		t.Assert(s5.X是否存在(1), true)
		t.Assert(s5.X是否存在(2), true)
		t.Assert(s5.X是否存在(3), false)
		t.Assert(s5.X是否存在(4), false)
	})
}

func TestSet_Intersect(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.NewSet别名()
		s2 := 集合类.NewSet别名()
		s1.X加入(1, 2, 3)
		s2.X加入(3, 4, 5)
		s3 := s1.X取交集(s2)
		t.Assert(s3.X是否存在(1), false)
		t.Assert(s3.X是否存在(2), false)
		t.Assert(s3.X是否存在(3), true)
		t.Assert(s3.X是否存在(4), false)
	})
}

func TestSet_Complement(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.NewSet别名()
		s2 := 集合类.NewSet别名()
		s1.X加入(1, 2, 3)
		s2.X加入(3, 4, 5)
		s3 := s1.X取补集(s2)
		t.Assert(s3.X是否存在(1), false)
		t.Assert(s3.X是否存在(2), false)
		t.Assert(s3.X是否存在(4), true)
		t.Assert(s3.X是否存在(5), true)
	})
}

func TestNewFrom(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建并按值("a")
		s2 := 集合类.X创建并按值("b", false)
		s3 := 集合类.X创建并按值(3, true)
		s4 := 集合类.X创建并按值([]string{"s1", "s2"}, true)
		t.Assert(s1.X是否存在("a"), true)
		t.Assert(s2.X是否存在("b"), true)
		t.Assert(s3.X是否存在(3), true)
		t.Assert(s4.X是否存在("s1"), true)
		t.Assert(s4.X是否存在("s3"), false)

	})
}

func TestNew(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建()
		s1.X加入("a", 2)
		s2 := 集合类.X创建(true)
		s2.X加入("b", 3)
		t.Assert(s1.X是否存在("a"), true)

	})
}

func TestSet_Join(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建(true)
		s1.X加入("a", "a1", "b", "c")
		str1 := s1.X取集合文本(",")
		t.Assert(strings.Contains(str1, "a1"), true)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建(true)
		s1.X加入("a", `"b"`, `\c`)
		str1 := s1.X取集合文本(",")
		t.Assert(strings.Contains(str1, `"b"`), true)
		t.Assert(strings.Contains(str1, `\c`), true)
		t.Assert(strings.Contains(str1, `a`), true)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.Set{}
		t.Assert(s1.X取集合文本(","), "")
	})
}

func TestSet_String(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建(true)
		s1.X加入("a", "a2", "b", "c")
		str1 := s1.String()
		t.Assert(strings.Contains(str1, "["), true)
		t.Assert(strings.Contains(str1, "]"), true)
		t.Assert(strings.Contains(str1, "a2"), true)

		s1 = nil
		t.Assert(s1.String(), "")

		s2 := 集合类.X创建()
		s2.X加入(1)
		t.Assert(s2.String(), "[1]")
	})
}

func TestSet_Merge(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建(true)
		s2 := 集合类.X创建(true)
		s1.X加入("a", "a2", "b", "c")
		s2.X加入("b", "b1", "e", "f")
		ss := s1.X合并(s2)
		t.Assert(ss.X是否存在("a2"), true)
		t.Assert(ss.X是否存在("b1"), true)

	})
}

func TestSet_Sum(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建(true)
		s1.X加入(1, 2, 3, 4)
		t.Assert(s1.X求和(), int(10))

	})
}

func TestSet_Pop(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建(true)
		t.Assert(s.X出栈(), nil)
		s.X加入(1, 2, 3, 4)
		t.Assert(s.X取数量(), 4)
		t.AssertIN(s.X出栈(), []int{1, 2, 3, 4})
		t.Assert(s.X取数量(), 3)
	})
}

func TestSet_Pops(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建(true)
		s.X加入(1, 2, 3, 4)
		t.Assert(s.X取数量(), 4)
		t.Assert(s.X出栈多个(0), nil)
		t.AssertIN(s.X出栈多个(1), []int{1, 2, 3, 4})
		t.Assert(s.X取数量(), 3)
		a := s.X出栈多个(6)
		t.Assert(len(a), 3)
		t.AssertIN(a, []int{1, 2, 3, 4})
		t.Assert(s.X取数量(), 0)
	})

	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建(true)
		a := []interface{}{1, 2, 3, 4}
		s.X加入(a...)
		t.Assert(s.X取数量(), 4)
		t.Assert(s.X出栈多个(-2), nil)
		t.AssertIN(s.X出栈多个(-1), a)
	})
}

func TestSet_Json(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := []interface{}{"a", "b", "d", "c"}
		a1 := 集合类.X创建并按值(s1)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(len(b1), len(b2))
		t.Assert(err1, err2)

		a2 := 集合类.X创建()
		err2 = json.UnmarshalUseNumber(b2, &a2)
		t.Assert(err2, nil)
		t.Assert(a2.X是否存在("a"), true)
		t.Assert(a2.X是否存在("b"), true)
		t.Assert(a2.X是否存在("c"), true)
		t.Assert(a2.X是否存在("d"), true)
		t.Assert(a2.X是否存在("e"), false)

		var a3 集合类.Set
		err := json.UnmarshalUseNumber(b2, &a3)
		t.AssertNil(err)
		t.Assert(a3.X是否存在("a"), true)
		t.Assert(a3.X是否存在("b"), true)
		t.Assert(a3.X是否存在("c"), true)
		t.Assert(a3.X是否存在("d"), true)
		t.Assert(a3.X是否存在("e"), false)
	})
}

func TestSet_AddIfNotExist(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建(true)
		s.X加入(1)
		t.Assert(s.X是否存在(1), true)
		t.Assert(s.X加入值并跳过已存在(1), false)
		t.Assert(s.X加入值并跳过已存在(2), true)
		t.Assert(s.X是否存在(2), true)
		t.Assert(s.X加入值并跳过已存在(2), false)
		t.Assert(s.X加入值并跳过已存在(nil), false)
		t.Assert(s.X是否存在(2), true)
	})
}

func TestSet_AddIfNotExistFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建(true)
		s.X加入(1)
		t.Assert(s.X是否存在(1), true)
		t.Assert(s.X是否存在(2), false)
		t.Assert(s.X加入值并跳过已存在_函数(2, func() bool { return false }), false)
		t.Assert(s.X是否存在(2), false)
		t.Assert(s.X加入值并跳过已存在_函数(2, func() bool { return true }), true)
		t.Assert(s.X是否存在(2), true)
		t.Assert(s.X加入值并跳过已存在_函数(2, func() bool { return true }), false)
		t.Assert(s.X是否存在(2), true)
		t.Assert(s.X加入值并跳过已存在_函数(nil, func() bool { return false }), false)
	})
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建(true)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			r := s.X加入值并跳过已存在_函数(1, func() bool {
				time.Sleep(100 * time.Millisecond)
				return true
			})
			t.Assert(r, false)
		}()
		s.X加入(1)
		wg.Wait()
	})
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.Set{}
		t.Assert(s.X加入值并跳过已存在_函数(1, func() bool { return true }), true)
	})
}

func TestSet_Walk(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var set 集合类.Set
		set.X加入(g.Slice{1, 2}...)
		set.X遍历修改(func(item interface{}) interface{} {
			return gconv.Int(item) + 10
		})
		t.Assert(set.X取数量(), 2)
		t.Assert(set.X是否存在(11), true)
		t.Assert(set.X是否存在(12), true)
	})
}

func TestSet_AddIfNotExistFuncLock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建(true)
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			r := s.X加入值并跳过已存在_并发安全函数(1, func() bool {
				time.Sleep(500 * time.Millisecond)
				return true
			})
			t.Assert(r, true)
		}()
		time.Sleep(100 * time.Millisecond)
		go func() {
			defer wg.Done()
			r := s.X加入值并跳过已存在_并发安全函数(1, func() bool {
				return true
			})
			t.Assert(r, false)
		}()
		wg.Wait()
	})
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建(true)
		t.Assert(s.X加入值并跳过已存在_并发安全函数(nil, func() bool { return true }), false)
		s1 := 集合类.Set{}
		t.Assert(s1.X加入值并跳过已存在_并发安全函数(1, func() bool { return true }), true)
	})
}

func TestSet_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Set  *集合类.Set
	}
	// JSON
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name": "john",
			"set":  []byte(`["k1","k2","k3"]`),
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Set.X取数量(), 3)
		t.Assert(v.Set.X是否存在("k1"), true)
		t.Assert(v.Set.X是否存在("k2"), true)
		t.Assert(v.Set.X是否存在("k3"), true)
		t.Assert(v.Set.X是否存在("k4"), false)
	})
	// Map
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name": "john",
			"set":  g.Slice{"k1", "k2", "k3"},
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Set.X取数量(), 3)
		t.Assert(v.Set.X是否存在("k1"), true)
		t.Assert(v.Set.X是否存在("k2"), true)
		t.Assert(v.Set.X是否存在("k3"), true)
		t.Assert(v.Set.X是否存在("k4"), false)
	})
}

func TestSet_DeepCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		set := 集合类.X创建()
		set.X加入(1, 2, 3)

		copySet := set.DeepCopy().(*集合类.Set)
		copySet.X加入(4)
		t.AssertNE(set.X取数量(), copySet.X取数量())
		t.AssertNE(set.String(), copySet.String())

		set = nil
		t.AssertNil(set.DeepCopy())
	})
}
