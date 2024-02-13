// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package 树形类_test

import (
	"fmt"
	"testing"
	
	"github.com/888go/goframe/container/gtree"
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gutil"
)

func getValue() interface{} {
	return 3
}

func Test_RedBlackTree_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := 树形类.NewRedBlackTree(工具类.X比较文本)
		m.X设置值("key1", "val1")
		t.Assert(m.Keys(), []interface{}{"key1"})

		t.Assert(m.Get("key1"), "val1")
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet("key2", "val2"), "val2")
		t.Assert(m.GetOrSet("key2", "val2"), "val2")
		t.Assert(m.SetIfNotExist("key2", "val2"), false)

		t.Assert(m.SetIfNotExist("key3", "val3"), true)

		t.Assert(m.Remove("key2"), "val2")
		t.Assert(m.Contains("key2"), false)

		t.AssertIN("key3", m.Keys())
		t.AssertIN("key1", m.Keys())
		t.AssertIN("val3", m.Values())
		t.AssertIN("val1", m.Values())

		m.Sets(map[interface{}]interface{}{"key3": "val3", "key1": "val1"})

		m.Flip()
		t.Assert(m.Map(), map[interface{}]interface{}{"val3": "key3", "val1": "key1"})

		m.Flip(工具类.X比较文本)
		t.Assert(m.Map(), map[interface{}]interface{}{"key3": "val3", "key1": "val1"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)

		m2 := 树形类.NewRedBlackTreeFrom(工具类.X比较文本, map[interface{}]interface{}{1: 1, "key1": "val1"})
		t.Assert(m2.Map(), map[interface{}]interface{}{1: 1, "key1": "val1"})
	})
}

func Test_RedBlackTree_Set_Fun(t *testing.T) {
// GetOrSetFunc 获取或设置（锁定或解锁）
	单元测试类.C(t, func(t *单元测试类.T) {
		m := 树形类.NewRedBlackTree(工具类.X比较文本)
		t.Assert(m.GetOrSetFunc("fun", getValue), 3)
		t.Assert(m.GetOrSetFunc("fun", getValue), 3)
		t.Assert(m.GetOrSetFuncLock("funlock", getValue), 3)
		t.Assert(m.GetOrSetFuncLock("funlock", getValue), 3)
		t.Assert(m.Get("funlock"), 3)
		t.Assert(m.Get("fun"), 3)
	})
// SetIfNotExistFunc 设置不存在时的处理函数，用于锁定或解锁
// 这个注释翻译可能不够准确，因为缺少了上下文。但从字面意思上看，该函数（SetIfNotExistFunc）可能是用来设置一个回调函数，当某个条件（例如资源）不存在时，执行锁定或解锁操作。但在实际应用中，请结合具体代码逻辑理解此注释含义。
	单元测试类.C(t, func(t *单元测试类.T) {
		m := 树形类.NewRedBlackTree(工具类.X比较文本)
		t.Assert(m.SetIfNotExistFunc("fun", getValue), true)
		t.Assert(m.SetIfNotExistFunc("fun", getValue), false)
		t.Assert(m.SetIfNotExistFuncLock("funlock", getValue), true)
		t.Assert(m.SetIfNotExistFuncLock("funlock", getValue), false)
		t.Assert(m.Get("funlock"), 3)
		t.Assert(m.Get("fun"), 3)
	})

}

func Test_RedBlackTree_Get_Set_Var(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := 树形类.NewRedBlackTree(工具类.X比较文本)
		t.AssertEQ(m.SetIfNotExist("key1", "val1"), true)
		t.AssertEQ(m.SetIfNotExist("key1", "val1"), false)
		t.AssertEQ(m.GetVarOrSet("key1", "val1"), 泛型类.X创建("val1", true))
		t.AssertEQ(m.GetVar("key1"), 泛型类.X创建("val1", true))
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		m := 树形类.NewRedBlackTree(工具类.X比较文本)
		t.AssertEQ(m.GetVarOrSetFunc("fun", getValue), 泛型类.X创建(3, true))
		t.AssertEQ(m.GetVarOrSetFunc("fun", getValue), 泛型类.X创建(3, true))
		t.AssertEQ(m.GetVarOrSetFuncLock("funlock", getValue), 泛型类.X创建(3, true))
		t.AssertEQ(m.GetVarOrSetFuncLock("funlock", getValue), 泛型类.X创建(3, true))
	})
}

func Test_RedBlackTree_Batch(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := 树形类.NewRedBlackTree(工具类.X比较文本)
		m.Sets(map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
		t.Assert(m.Map(), map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
		m.Removes([]interface{}{"key1", 1})
		t.Assert(m.Map(), map[interface{}]interface{}{"key2": "val2", "key3": "val3"})
	})
}

func Test_RedBlackTree_Iterator(t *testing.T) {
	keys := []string{"1", "key1", "key2", "key3", "key4"}
	keyLen := len(keys)
	index := 0

	expect := map[interface{}]interface{}{"key4": "val4", 1: 1, "key1": "val1", "key2": "val2", "key3": "val3"}
	m := 树形类.NewRedBlackTreeFrom(工具类.X比较文本, expect)

	单元测试类.C(t, func(t *单元测试类.T) {

		m.X遍历(func(k interface{}, v interface{}) bool {
			t.Assert(k, keys[index])
			index++
			t.Assert(expect[k], v)
			return true
		})

		m.IteratorDesc(func(k interface{}, v interface{}) bool {
			index--
			t.Assert(k, keys[index])
			t.Assert(expect[k], v)
			return true
		})
	})
	m.Print()
	// 断言返回值对遍历控制
	单元测试类.C(t, func(t *单元测试类.T) {
		i := 0
		j := 0
		m.X遍历(func(k interface{}, v interface{}) bool {
			i++
			return true
		})
		m.X遍历(func(k interface{}, v interface{}) bool {
			j++
			return false
		})
		t.Assert(i, keyLen)
		t.Assert(j, 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		i := 0
		j := 0
		m.IteratorDesc(func(k interface{}, v interface{}) bool {
			i++
			return true
		})
		m.IteratorDesc(func(k interface{}, v interface{}) bool {
			j++
			return false
		})
		t.Assert(i, keyLen)
		t.Assert(j, 1)
	})
}

func Test_RedBlackTree_IteratorFrom(t *testing.T) {
	m := make(map[interface{}]interface{})
	for i := 1; i <= 10; i++ {
		m[i] = i * 10
	}
	tree := 树形类.NewRedBlackTreeFrom(工具类.X比较整数, m)

	单元测试类.C(t, func(t *单元测试类.T) {
		n := 5
		tree.IteratorFrom(5, true, func(key, value interface{}) bool {
			t.Assert(n, key)
			t.Assert(n*10, value)
			n++
			return true
		})

		i := 5
		tree.IteratorAscFrom(5, true, func(key, value interface{}) bool {
			t.Assert(i, key)
			t.Assert(i*10, value)
			i++
			return true
		})

		j := 5
		tree.IteratorDescFrom(5, true, func(key, value interface{}) bool {
			t.Assert(j, key)
			t.Assert(j*10, value)
			j--
			return true
		})
	})
}

func Test_RedBlackTree_Clone(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		//clone 方法是深克隆
		m := 树形类.NewRedBlackTreeFrom(工具类.X比较文本, map[interface{}]interface{}{1: 1, "key1": "val1"})
		m_clone := m.Clone()
		m.Remove(1)
		//修改原 map,clone 后的 map 不影响
		t.AssertIN(1, m_clone.Keys())

		m_clone.Remove("key1")
		//修改clone map,原 map 不影响
		t.AssertIN("key1", m.Keys())
	})
}

func Test_RedBlackTree_LRNode(t *testing.T) {
	expect := map[interface{}]interface{}{"key4": "val4", "key1": "val1", "key2": "val2", "key3": "val3"}
	//safe
	单元测试类.C(t, func(t *单元测试类.T) {
		m := 树形类.NewRedBlackTreeFrom(工具类.X比较文本, expect)
		t.Assert(m.Left().Key, "key1")
		t.Assert(m.Right().Key, "key4")
	})
	//unsafe
	单元测试类.C(t, func(t *单元测试类.T) {
		m := 树形类.NewRedBlackTreeFrom(工具类.X比较文本, expect, true)
		t.Assert(m.Left().Key, "key1")
		t.Assert(m.Right().Key, "key4")
	})
}

func Test_RedBlackTree_CeilingFloor(t *testing.T) {
	expect := map[interface{}]interface{}{
		20: "val20",
		6:  "val6",
		10: "val10",
		12: "val12",
		1:  "val1",
		15: "val15",
		19: "val19",
		8:  "val8",
		4:  "val4"}
	//found and eq
	单元测试类.C(t, func(t *单元测试类.T) {
		m := 树形类.NewRedBlackTreeFrom(工具类.X比较整数, expect)
		c, cf := m.Ceiling(8)
		t.Assert(cf, true)
		t.Assert(c.Value, "val8")
		f, ff := m.Floor(20)
		t.Assert(ff, true)
		t.Assert(f.Value, "val20")
	})
	//found and neq
	单元测试类.C(t, func(t *单元测试类.T) {
		m := 树形类.NewRedBlackTreeFrom(工具类.X比较整数, expect)
		c, cf := m.Ceiling(9)
		t.Assert(cf, true)
		t.Assert(c.Value, "val10")
		f, ff := m.Floor(5)
		t.Assert(ff, true)
		t.Assert(f.Value, "val4")
	})
	//nofound
	单元测试类.C(t, func(t *单元测试类.T) {
		m := 树形类.NewRedBlackTreeFrom(工具类.X比较整数, expect)
		c, cf := m.Ceiling(21)
		t.Assert(cf, false)
		t.Assert(c, nil)
		f, ff := m.Floor(-1)
		t.Assert(ff, false)
		t.Assert(f, nil)
	})
}

func Test_RedBlackTree_Remove(t *testing.T) {
	m := 树形类.NewRedBlackTree(工具类.X比较整数)
	for i := 1; i <= 100; i++ {
		m.X设置值(i, fmt.Sprintf("val%d", i))
	}
	expect := m.Map()
	单元测试类.C(t, func(t *单元测试类.T) {
		for k, v := range expect {
			m1 := m.Clone()
			t.Assert(m1.Remove(k), v)
			t.Assert(m1.Remove(k), nil)
		}
	})
}
