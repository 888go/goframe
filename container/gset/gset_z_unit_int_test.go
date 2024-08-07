// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// go test *.go

package 集合类_test

import (
	"strings"
	"sync"
	"testing"
	"time"

	garray "github.com/888go/goframe/container/garray"
	gset "github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
)

func TestIntSet_Var(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var s gset.IntSet
		s.X加入(1, 1, 2)
		s.X加入([]int{3, 4}...)
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

func TestIntSet_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := gset.X创建整数()
		s.X加入(1, 1, 2)
		s.X加入([]int{3, 4}...)
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

func TestIntSet_Iterator(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := gset.X创建整数()
		s.X加入(1, 2, 3)
		t.Assert(s.X取数量(), 3)

		a1 := garray.X创建(true)
		a2 := garray.X创建(true)
		s.X遍历(func(v int) bool {
			a1.Append别名(1)
			return false
		})
		s.X遍历(func(v int) bool {
			a2.Append别名(1)
			return true
		})
		t.Assert(a1.X取长度(), 1)
		t.Assert(a2.X取长度(), 3)
	})
}

func TestIntSet_LockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := gset.X创建整数()
		s.X加入(1, 2, 3)
		t.Assert(s.X取数量(), 3)
		s.X写锁定_函数(func(m map[int]struct{}) {
			delete(m, 1)
		})
		t.Assert(s.X取数量(), 2)
		s.X读锁定_函数(func(m map[int]struct{}) {
			t.Assert(m, map[int]struct{}{
				3: {},
				2: {},
			})
		})
	})
}

func TestIntSet_Equal(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := gset.X创建整数()
		s2 := gset.X创建整数()
		s3 := gset.X创建整数()
		s4 := gset.X创建整数()
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

func TestIntSet_IsSubsetOf(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := gset.X创建整数()
		s2 := gset.X创建整数()
		s3 := gset.X创建整数()
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

func TestIntSet_Union(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := gset.X创建整数()
		s2 := gset.X创建整数()
		s1.X加入(1, 2)
		s2.X加入(3, 4)
		s3 := s1.X取并集(s2)
		t.Assert(s3.X是否存在(1), true)
		t.Assert(s3.X是否存在(2), true)
		t.Assert(s3.X是否存在(3), true)
		t.Assert(s3.X是否存在(4), true)
	})
}

func TestIntSet_Diff(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := gset.X创建整数()
		s2 := gset.X创建整数()
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

func TestIntSet_Intersect(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := gset.X创建整数()
		s2 := gset.X创建整数()
		s1.X加入(1, 2, 3)
		s2.X加入(3, 4, 5)
		s3 := s1.X取交集(s2)
		t.Assert(s3.X是否存在(1), false)
		t.Assert(s3.X是否存在(2), false)
		t.Assert(s3.X是否存在(3), true)
		t.Assert(s3.X是否存在(4), false)
	})
}

func TestIntSet_Complement(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := gset.X创建整数()
		s2 := gset.X创建整数()
		s1.X加入(1, 2, 3)
		s2.X加入(3, 4, 5)
		s3 := s1.X取补集(s2)
		t.Assert(s3.X是否存在(1), false)
		t.Assert(s3.X是否存在(2), false)
		t.Assert(s3.X是否存在(4), true)
		t.Assert(s3.X是否存在(5), true)
	})
}

func TestIntSet_Size(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := gset.X创建整数(true)
		s1.X加入(1, 2, 3)
		t.Assert(s1.X取数量(), 3)

	})

}

func TestIntSet_Merge(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := gset.X创建整数()
		s2 := gset.X创建整数()
		s1.X加入(1, 2, 3)
		s2.X加入(3, 4, 5)
		s3 := s1.X合并(s2)
		t.Assert(s3.X是否存在(1), true)
		t.Assert(s3.X是否存在(5), true)
		t.Assert(s3.X是否存在(6), false)
	})
}

func TestIntSet_Join(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := gset.X创建整数()
		t.Assert(s1.X取集合文本(","), "")
		s1.X加入(1, 2, 3)
		s3 := s1.X取集合文本(",")
		t.Assert(strings.Contains(s3, "1"), true)
		t.Assert(strings.Contains(s3, "2"), true)
		t.Assert(strings.Contains(s3, "3"), true)
	})
}

func TestIntSet_String(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := gset.X创建整数()
		s1.X加入(1, 2, 3)
		s3 := s1.String()
		t.Assert(strings.Contains(s3, "["), true)
		t.Assert(strings.Contains(s3, "]"), true)
		t.Assert(strings.Contains(s3, "1"), true)
		t.Assert(strings.Contains(s3, "2"), true)
		t.Assert(strings.Contains(s3, "3"), true)
		s1 = nil
		t.Assert(s1.String(), "")
	})
}

func TestIntSet_Sum(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := gset.X创建整数()
		s1.X加入(1, 2, 3)
		s2 := gset.X创建整数()
		s2.X加入(5, 6, 7)
		t.Assert(s2.X求和(), 18)

	})

}

func TestIntSet_Pop(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := gset.X创建整数()
		t.Assert(s.X出栈(), 0)
		s.X加入(4, 2, 3)
		t.Assert(s.X取数量(), 3)
		t.AssertIN(s.X出栈(), []int{4, 2, 3})
		t.AssertIN(s.X出栈(), []int{4, 2, 3})
		t.Assert(s.X取数量(), 1)
	})
}

func TestIntSet_Pops(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := gset.X创建整数()
		s.X加入(1, 4, 2, 3)
		t.Assert(s.X取数量(), 4)
		t.Assert(s.X出栈多个(0), nil)
		t.AssertIN(s.X出栈多个(1), []int{1, 4, 2, 3})
		t.Assert(s.X取数量(), 3)
		a := s.X出栈多个(2)
		t.Assert(len(a), 2)
		t.AssertIN(a, []int{1, 4, 2, 3})
		t.Assert(s.X取数量(), 1)
	})

	gtest.C(t, func(t *gtest.T) {
		s := gset.X创建整数(true)
		a := []int{1, 2, 3, 4}
		s.X加入(a...)
		t.Assert(s.X取数量(), 4)
		t.Assert(s.X出栈多个(-2), nil)
		t.AssertIN(s.X出栈多个(-1), a)
	})
}

func TestIntSet_AddIfNotExist(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := gset.X创建整数(true)
		s.X加入(1)
		t.Assert(s.X是否存在(1), true)
		t.Assert(s.X加入值并跳过已存在(1), false)
		t.Assert(s.X加入值并跳过已存在(2), true)
		t.Assert(s.X是否存在(2), true)
		t.Assert(s.X加入值并跳过已存在(2), false)
		t.Assert(s.X是否存在(2), true)
	})
	gtest.C(t, func(t *gtest.T) {
		s := gset.IntSet{}
		t.Assert(s.X加入值并跳过已存在(1), true)
	})
}

func TestIntSet_AddIfNotExistFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := gset.X创建整数(true)
		s.X加入(1)
		t.Assert(s.X是否存在(1), true)
		t.Assert(s.X是否存在(2), false)
		t.Assert(s.X加入值并跳过已存在_函数(2, func() bool { return false }), false)
		t.Assert(s.X是否存在(2), false)
		t.Assert(s.X加入值并跳过已存在_函数(2, func() bool { return true }), true)
		t.Assert(s.X是否存在(2), true)
		t.Assert(s.X加入值并跳过已存在_函数(2, func() bool { return true }), false)
		t.Assert(s.X是否存在(2), true)
	})
	gtest.C(t, func(t *gtest.T) {
		s := gset.X创建整数(true)
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
		s := gset.IntSet{}
		t.Assert(s.X加入值并跳过已存在_函数(1, func() bool { return true }), true)
	})
}

func TestIntSet_AddIfNotExistFuncLock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := gset.X创建整数(true)
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
		s := gset.IntSet{}
		t.Assert(s.X加入值并跳过已存在_并发安全函数(1, func() bool { return true }), true)
	})
}

func TestIntSet_Json(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := []int{1, 3, 2, 4}
		a1 := gset.X创建整数并按值(s1)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(len(b1), len(b2))
		t.Assert(err1, err2)

		a2 := gset.X创建整数()
		err2 = json.UnmarshalUseNumber(b2, &a2)
		t.Assert(err2, nil)
		t.Assert(a2.X是否存在(1), true)
		t.Assert(a2.X是否存在(2), true)
		t.Assert(a2.X是否存在(3), true)
		t.Assert(a2.X是否存在(4), true)
		t.Assert(a2.X是否存在(5), false)

		var a3 gset.IntSet
		err := json.UnmarshalUseNumber(b2, &a3)
		t.AssertNil(err)
		t.Assert(a2.X是否存在(1), true)
		t.Assert(a2.X是否存在(2), true)
		t.Assert(a2.X是否存在(3), true)
		t.Assert(a2.X是否存在(4), true)
		t.Assert(a2.X是否存在(5), false)
	})
}

func TestIntSet_Walk(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var set gset.IntSet
		set.X加入(g.SliceInt别名{1, 2}...)
		set.X遍历修改(func(item int) int {
			return item + 10
		})
		t.Assert(set.X取数量(), 2)
		t.Assert(set.X是否存在(11), true)
		t.Assert(set.X是否存在(12), true)
	})
}

func TestIntSet_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Set  *gset.IntSet
	}
	// JSON
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name": "john",
			"set":  []byte(`[1,2,3]`),
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Set.X取数量(), 3)
		t.Assert(v.Set.X是否存在(1), true)
		t.Assert(v.Set.X是否存在(2), true)
		t.Assert(v.Set.X是否存在(3), true)
		t.Assert(v.Set.X是否存在(4), false)
	})
	// Map
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name": "john",
			"set":  g.Slice别名{1, 2, 3},
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Set.X取数量(), 3)
		t.Assert(v.Set.X是否存在(1), true)
		t.Assert(v.Set.X是否存在(2), true)
		t.Assert(v.Set.X是否存在(3), true)
		t.Assert(v.Set.X是否存在(4), false)
	})
}

func TestIntSet_DeepCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		set := gset.X创建整数()
		set.X加入(1, 2, 3)

		copySet := set.DeepCopy().(*gset.IntSet)
		copySet.X加入(4)
		t.AssertNE(set.X取数量(), copySet.X取数量())
		t.AssertNE(set.String(), copySet.String())

		set = nil
		t.AssertNil(set.DeepCopy())
	})
}
