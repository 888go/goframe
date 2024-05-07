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

func TestStrSet_Var(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var s 集合类.StrSet
		s.X加入("1", "1", "2")
		s.X加入([]string{"3", "4"}...)
		t.Assert(s.X取数量(), 4)
		t.AssertIN("1", s.X取集合切片())
		t.AssertIN("2", s.X取集合切片())
		t.AssertIN("3", s.X取集合切片())
		t.AssertIN("4", s.X取集合切片())
		t.AssertNI("0", s.X取集合切片())
		t.Assert(s.X是否存在("4"), true)
		t.Assert(s.X是否存在("5"), false)
		s.X删除("1")
		t.Assert(s.X取数量(), 3)
		s.X清空()
		t.Assert(s.X取数量(), 0)
	})
}

func TestStrSet_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建文本()
		s.X加入("1", "1", "2")
		s.X加入([]string{"3", "4"}...)
		t.Assert(s.X取数量(), 4)
		t.AssertIN("1", s.X取集合切片())
		t.AssertIN("2", s.X取集合切片())
		t.AssertIN("3", s.X取集合切片())
		t.AssertIN("4", s.X取集合切片())
		t.AssertNI("0", s.X取集合切片())
		t.Assert(s.X是否存在("4"), true)
		t.Assert(s.X是否存在("5"), false)
		s.X删除("1")
		t.Assert(s.X取数量(), 3)
		s.X清空()
		t.Assert(s.X取数量(), 0)
	})
}

func TestStrSet_ContainsI(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建文本()
		s.X加入("a", "b", "C")
		t.Assert(s.X是否存在("A"), false)
		t.Assert(s.X是否存在("a"), true)
		t.Assert(s.X是否存在并忽略大小写("A"), true)
		t.Assert(s.X是否存在并忽略大小写("d"), false)
	})
}

func TestStrSet_Iterator(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建文本()
		s.X加入("1", "2", "3")
		t.Assert(s.X取数量(), 3)

		a1 := garray.New(true)
		a2 := garray.New(true)
		s.X遍历(func(v string) bool {
			a1.Append("1")
			return false
		})
		s.X遍历(func(v string) bool {
			a2.Append("1")
			return true
		})
		t.Assert(a1.Len(), 1)
		t.Assert(a2.Len(), 3)
	})
}

func TestStrSet_LockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建文本()
		s.X加入("1", "2", "3")
		t.Assert(s.X取数量(), 3)
		s.X写锁定_函数(func(m map[string]struct{}) {
			delete(m, "1")
		})
		t.Assert(s.X取数量(), 2)
		s.X读锁定_函数(func(m map[string]struct{}) {
			t.Assert(m, map[string]struct{}{
				"3": {},
				"2": {},
			})
		})
	})
}

func TestStrSet_Equal(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建文本()
		s2 := 集合类.X创建文本()
		s3 := 集合类.X创建文本()
		s4 := 集合类.X创建文本()
		s1.X加入("1", "2", "3")
		s2.X加入("1", "2", "3")
		s3.X加入("1", "2", "3", "4")
		s4.X加入("4", "5", "6")
		t.Assert(s1.X是否相等(s2), true)
		t.Assert(s1.X是否相等(s3), false)
		t.Assert(s1.X是否相等(s4), false)
		s5 := s1
		t.Assert(s1.X是否相等(s5), true)
	})
}

func TestStrSet_IsSubsetOf(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建文本()
		s2 := 集合类.X创建文本()
		s3 := 集合类.X创建文本()
		s1.X加入("1", "2")
		s2.X加入("1", "2", "3")
		s3.X加入("1", "2", "3", "4")
		t.Assert(s1.X是否为子集(s2), true)
		t.Assert(s2.X是否为子集(s3), true)
		t.Assert(s1.X是否为子集(s3), true)
		t.Assert(s2.X是否为子集(s1), false)
		t.Assert(s3.X是否为子集(s2), false)

		s4 := s1
		t.Assert(s1.X是否为子集(s4), true)
	})
}

func TestStrSet_Union(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建文本()
		s2 := 集合类.X创建文本()
		s1.X加入("1", "2")
		s2.X加入("3", "4")
		s3 := s1.X取并集(s2)
		t.Assert(s3.X是否存在("1"), true)
		t.Assert(s3.X是否存在("2"), true)
		t.Assert(s3.X是否存在("3"), true)
		t.Assert(s3.X是否存在("4"), true)
	})
}

func TestStrSet_Diff(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建文本()
		s2 := 集合类.X创建文本()
		s1.X加入("1", "2", "3")
		s2.X加入("3", "4", "5")
		s3 := s1.X取差集(s2)
		t.Assert(s3.X是否存在("1"), true)
		t.Assert(s3.X是否存在("2"), true)
		t.Assert(s3.X是否存在("3"), false)
		t.Assert(s3.X是否存在("4"), false)

		s4 := s1
		s5 := s1.X取差集(s2, s4)
		t.Assert(s5.X是否存在("1"), true)
		t.Assert(s5.X是否存在("2"), true)
		t.Assert(s5.X是否存在("3"), false)
		t.Assert(s5.X是否存在("4"), false)
	})
}

func TestStrSet_Intersect(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建文本()
		s2 := 集合类.X创建文本()
		s1.X加入("1", "2", "3")
		s2.X加入("3", "4", "5")
		s3 := s1.X取交集(s2)
		t.Assert(s3.X是否存在("1"), false)
		t.Assert(s3.X是否存在("2"), false)
		t.Assert(s3.X是否存在("3"), true)
		t.Assert(s3.X是否存在("4"), false)
	})
}

func TestStrSet_Complement(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建文本()
		s2 := 集合类.X创建文本()
		s1.X加入("1", "2", "3")
		s2.X加入("3", "4", "5")
		s3 := s1.X取补集(s2)
		t.Assert(s3.X是否存在("1"), false)
		t.Assert(s3.X是否存在("2"), false)
		t.Assert(s3.X是否存在("4"), true)
		t.Assert(s3.X是否存在("5"), true)
	})
}

func TestNewIntSetFrom(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建整数并按值([]int{1, 2, 3, 4})
		s2 := 集合类.X创建整数并按值([]int{5, 6, 7, 8})
		t.Assert(s1.X是否存在(3), true)
		t.Assert(s1.X是否存在(5), false)
		t.Assert(s2.X是否存在(3), false)
		t.Assert(s2.X是否存在(5), true)
	})
}

func TestStrSet_Merge(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建文本()
		s2 := 集合类.X创建文本()
		s1.X加入("1", "2", "3")
		s2.X加入("3", "4", "5")
		s3 := s1.X合并(s2)
		t.Assert(s3.X是否存在("1"), true)
		t.Assert(s3.X是否存在("6"), false)
		t.Assert(s3.X是否存在("4"), true)
		t.Assert(s3.X是否存在("5"), true)
	})
}

func TestNewStrSetFrom(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建文本并按值([]string{"a", "b", "c"}, true)
		t.Assert(s1.X是否存在("b"), true)
		t.Assert(s1.X是否存在("d"), false)
	})
}

func TestStrSet_Join(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建文本并按值([]string{"a", "b", "c"}, true)
		str1 := s1.X取集合文本(",")
		t.Assert(strings.Contains(str1, "b"), true)
		t.Assert(strings.Contains(str1, "d"), false)
	})

	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建文本()
		t.Assert(s1.X取集合文本(","), "")
		s1.X加入("a", `"b"`, `\c`)
		str1 := s1.X取集合文本(",")
		t.Assert(strings.Contains(str1, `"b"`), true)
		t.Assert(strings.Contains(str1, `\c`), true)
		t.Assert(strings.Contains(str1, `a`), true)
	})
}

func TestStrSet_String(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建文本并按值([]string{"a", "b", "c"}, true)
		str1 := s1.String()
		t.Assert(strings.Contains(str1, "b"), true)
		t.Assert(strings.Contains(str1, "d"), false)
		s1 = nil
		t.Assert(s1.String(), "")
	})

	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建(true)
		s1.X加入("a", "a2", "b", "c")
		str1 := s1.String()
		t.Assert(strings.Contains(str1, "["), true)
		t.Assert(strings.Contains(str1, "]"), true)
		t.Assert(strings.Contains(str1, "a2"), true)
	})
}

func TestStrSet_Sum(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建文本并按值([]string{"a", "b", "c"}, true)
		s2 := 集合类.X创建整数并按值([]int{2, 3, 4}, true)
		t.Assert(s1.X求和(), 0)
		t.Assert(s2.X求和(), 9)
	})
}

func TestStrSet_Size(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建文本并按值([]string{"a", "b", "c"}, true)
		t.Assert(s1.X取数量(), 3)

	})
}

func TestStrSet_Remove(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := 集合类.X创建文本并按值([]string{"a", "b", "c"}, true)
		s1.X删除("b")
		t.Assert(s1.X是否存在("b"), false)
		t.Assert(s1.X是否存在("c"), true)
	})
}

func TestStrSet_Pop(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := []string{"a", "b", "c", "d"}
		s := 集合类.X创建文本并按值(a, true)
		t.Assert(s.X取数量(), 4)
		t.AssertIN(s.X出栈(), a)
		t.Assert(s.X取数量(), 3)
		t.AssertIN(s.X出栈(), a)
		t.Assert(s.X取数量(), 2)

		s1 := 集合类.StrSet{}
		t.Assert(s1.X出栈(), "")
	})
}

func TestStrSet_Pops(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := []string{"a", "b", "c", "d"}
		s := 集合类.X创建文本并按值(a, true)
		array := s.X出栈多个(2)
		t.Assert(len(array), 2)
		t.Assert(s.X取数量(), 2)
		t.AssertIN(array, a)
		t.Assert(s.X出栈多个(0), nil)
		t.AssertIN(s.X出栈多个(2), a)
		t.Assert(s.X取数量(), 0)
	})

	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建文本(true)
		a := []string{"1", "2", "3", "4"}
		s.X加入(a...)
		t.Assert(s.X取数量(), 4)
		t.Assert(s.X出栈多个(-2), nil)
		t.AssertIN(s.X出栈多个(-1), a)
	})
}

func TestStrSet_AddIfNotExist(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建文本(true)
		s.X加入("1")
		t.Assert(s.X是否存在("1"), true)
		t.Assert(s.X加入值并跳过已存在("1"), false)
		t.Assert(s.X加入值并跳过已存在("2"), true)
		t.Assert(s.X是否存在("2"), true)
		t.Assert(s.X加入值并跳过已存在("2"), false)
		t.Assert(s.X是否存在("2"), true)
	})
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.StrSet{}
		t.Assert(s.X加入值并跳过已存在("1"), true)
	})
}

func TestStrSet_AddIfNotExistFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建文本(true)
		s.X加入("1")
		t.Assert(s.X是否存在("1"), true)
		t.Assert(s.X是否存在("2"), false)
		t.Assert(s.X加入值并跳过已存在_函数("2", func() bool { return false }), false)
		t.Assert(s.X是否存在("2"), false)
		t.Assert(s.X加入值并跳过已存在_函数("2", func() bool { return true }), true)
		t.Assert(s.X是否存在("2"), true)
		t.Assert(s.X加入值并跳过已存在_函数("2", func() bool { return true }), false)
		t.Assert(s.X是否存在("2"), true)
	})
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建文本(true)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			r := s.X加入值并跳过已存在_函数("1", func() bool {
				time.Sleep(100 * time.Millisecond)
				return true
			})
			t.Assert(r, false)
		}()
		s.X加入("1")
		wg.Wait()
	})
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.StrSet{}
		t.Assert(s.X加入值并跳过已存在_函数("1", func() bool { return true }), true)
	})
}

func TestStrSet_AddIfNotExistFuncLock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.X创建文本(true)
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			r := s.X加入值并跳过已存在_并发安全函数("1", func() bool {
				time.Sleep(500 * time.Millisecond)
				return true
			})
			t.Assert(r, true)
		}()
		time.Sleep(100 * time.Millisecond)
		go func() {
			defer wg.Done()
			r := s.X加入值并跳过已存在_并发安全函数("1", func() bool {
				return true
			})
			t.Assert(r, false)
		}()
		wg.Wait()
	})
	gtest.C(t, func(t *gtest.T) {
		s := 集合类.StrSet{}
		t.Assert(s.X加入值并跳过已存在_并发安全函数("1", func() bool { return true }), true)
	})
}

func TestStrSet_Json(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := []string{"a", "b", "d", "c"}
		a1 := 集合类.X创建文本并按值(s1)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(len(b1), len(b2))
		t.Assert(err1, err2)

		a2 := 集合类.X创建文本()
		err2 = json.UnmarshalUseNumber(b2, &a2)
		t.Assert(err2, nil)
		t.Assert(a2.X是否存在("a"), true)
		t.Assert(a2.X是否存在("b"), true)
		t.Assert(a2.X是否存在("c"), true)
		t.Assert(a2.X是否存在("d"), true)
		t.Assert(a2.X是否存在("e"), false)

		var a3 集合类.StrSet
		err := json.UnmarshalUseNumber(b2, &a3)
		t.AssertNil(err)
		t.Assert(a3.X是否存在("a"), true)
		t.Assert(a3.X是否存在("b"), true)
		t.Assert(a3.X是否存在("c"), true)
		t.Assert(a3.X是否存在("d"), true)
		t.Assert(a3.X是否存在("e"), false)
	})
}

func TestStrSet_Walk(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			set    集合类.StrSet
			names  = g.SliceStr{"user", "user_detail"}
			prefix = "gf_"
		)
		set.X加入(names...)
		// 为给定的表名添加前缀。
		set.X遍历修改(func(item string) string {
			return prefix + item
		})
		t.Assert(set.X取数量(), 2)
		t.Assert(set.X是否存在("gf_user"), true)
		t.Assert(set.X是否存在("gf_user_detail"), true)
	})
}

func TestStrSet_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Set  *集合类.StrSet
	}
	// JSON
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name": "john",
			"set":  []byte(`["1","2","3"]`),
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Set.X取数量(), 3)
		t.Assert(v.Set.X是否存在("1"), true)
		t.Assert(v.Set.X是否存在("2"), true)
		t.Assert(v.Set.X是否存在("3"), true)
		t.Assert(v.Set.X是否存在("4"), false)
	})
	// Map
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name": "john",
			"set":  g.SliceStr{"1", "2", "3"},
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Set.X取数量(), 3)
		t.Assert(v.Set.X是否存在("1"), true)
		t.Assert(v.Set.X是否存在("2"), true)
		t.Assert(v.Set.X是否存在("3"), true)
		t.Assert(v.Set.X是否存在("4"), false)
	})
}

func TestStrSet_DeepCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		set := 集合类.X创建文本()
		set.X加入("1", "2", "3")

		copySet := set.DeepCopy().(*集合类.StrSet)
		copySet.X加入("4")
		t.AssertNE(set.X取数量(), copySet.X取数量())
		t.AssertNE(set.String(), copySet.String())

		set = nil
		t.AssertNil(set.DeepCopy())
	})
}
